// Copyright © 2022 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syft

import (
	"fmt"

	"github.com/anchore/syft/syft"
	"github.com/anchore/syft/syft/linux"
	syft_pkg "github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/pkg/cataloger"
	syft_sbom "github.com/anchore/syft/syft/sbom"
	"github.com/anchore/syft/syft/source"
	log "github.com/sirupsen/logrus"

	"github.com/openclarity/kubeclarity/shared/pkg/analyzer"
	"github.com/openclarity/kubeclarity/shared/pkg/config"
	"github.com/openclarity/kubeclarity/shared/pkg/formatter"
	"github.com/openclarity/kubeclarity/shared/pkg/job_manager"
	"github.com/openclarity/kubeclarity/shared/pkg/utils"
	"github.com/openclarity/kubeclarity/shared/pkg/utils/image_helper"
)

const AnalyzerName = "syft"

type Analyzer struct {
	name       string
	logger     *log.Entry
	config     config.SyftConfig
	resultChan chan job_manager.Result
	localImage bool
}

func New(conf *config.Config,
	logger *log.Entry,
	resultChan chan job_manager.Result) job_manager.Job {
	return &Analyzer{
		name:       AnalyzerName,
		logger:     logger.Dup().WithField("analyzer", AnalyzerName),
		config:     config.CreateSyftConfig(conf.Analyzer, conf.Registry),
		resultChan: resultChan,
		localImage: conf.LocalImageScan,
	}
}

func (a *Analyzer) Run(sourceType utils.SourceType, userInput string) error {
	src := utils.CreateSource(sourceType, userInput, a.localImage)
	a.logger.Infof("Called %s analyzer on source %s", a.name, src)
	s, _, err := source.New(src, a.config.RegistryOptions, []string{})
	if err != nil {
		return fmt.Errorf("failed to create source analyzer=%s: %v", a.name, err)
	}

	go func() {
		res := &analyzer.Results{}
		catalogerConfig := cataloger.Config{
			Search: cataloger.DefaultSearchConfig(),
		}
		catalogerConfig.Search.Scope = a.config.Scope

		p, _, d, err := syft.CatalogPackages(s, catalogerConfig)
		if err != nil {
			a.setError(res, fmt.Errorf("failed to write results: %v", err))
			return
		}

		output := formatter.New(formatter.SyftFormat, []byte{})
		sbom := generateSBOM(p, d, s)
		if err := output.SetSBOM(sbom); err != nil {
			a.setError(res, fmt.Errorf("failed to set SBOM struct in formatter: %v", err))
			return
		}
		if err := output.Encode(string(a.config.OutputFormat)); err != nil {
			a.setError(res, fmt.Errorf("failed to write results: %v", err))
			return
		}

		res = analyzer.CreateResults(output.GetSBOMBytes(), a.name, src, sourceType)
		// Syft uses ManifestDigest to fill version information in the case of an image.
		// We need RepoDigest as well which is not set by Syft if we using cycloneDX output.
		// Get the RepoDigest from image metadata and use it as SourceHash in the Result
		// that will be added to the component hash of metadata during the merge.
		if sourceType == utils.IMAGE {
			res.AppInfo.SourceHash = getImageHash(sbom, userInput)
		}
		a.logger.Infof("Sending successful results")
		a.resultChan <- res
	}()

	return nil
}

func generateSBOM(c *syft_pkg.Catalog, d *linux.Release, s *source.Source) syft_sbom.SBOM {
	return syft_sbom.SBOM{
		Artifacts: syft_sbom.Artifacts{
			PackageCatalog:    c,
			LinuxDistribution: d,
		},
		Source: s.Metadata,
	}
}

func (a *Analyzer) setError(res *analyzer.Results, err error) {
	res.Error = err
	a.logger.Error(res.Error)
	a.resultChan <- res
}

func getImageHash(sbom syft_sbom.SBOM, src string) string {
	return image_helper.GetHashFromRepoDigest(sbom.Source.ImageMetadata.RepoDigests, src)
}
