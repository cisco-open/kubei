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

package cmd

import (
	"fmt"
	"io"
	"os"

	dockle_config "github.com/Portshift/dockle/config"
	dockle_run "github.com/Portshift/dockle/pkg"
	dockle_types "github.com/Portshift/dockle/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/openclarity/kubeclarity/cli/pkg/config"
	_export "github.com/openclarity/kubeclarity/cli/pkg/scanner/export"
	"github.com/openclarity/kubeclarity/cli/pkg/scanner/filter"
	"github.com/openclarity/kubeclarity/cli/pkg/scanner/presenter"
	"github.com/openclarity/kubeclarity/cli/pkg/utils"
	sharedconfig "github.com/openclarity/kubeclarity/shared/pkg/config"
	"github.com/openclarity/kubeclarity/shared/pkg/converter"
	"github.com/openclarity/kubeclarity/shared/pkg/job_manager"
	sharedscanner "github.com/openclarity/kubeclarity/shared/pkg/scanner"
	"github.com/openclarity/kubeclarity/shared/pkg/scanner/job"
	sharedutils "github.com/openclarity/kubeclarity/shared/pkg/utils"
	cdx_helper "github.com/openclarity/kubeclarity/shared/pkg/utils/cyclonedx_helper"
	"github.com/openclarity/kubeclarity/shared/pkg/utils/image_helper"
)

// scanCmd represents the scan command.
var scanCmd = &cobra.Command{
	Use:   "scan [SOURCE]",
	Short: "Vulnerability scanner",
	Long:  `KubeClarity vulnerability scanner`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("an image/directory/file/sbom argument is required")
		}

		return cobra.MaximumNArgs(1)(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		vulnerabilityScanner(cmd, args)
	},
}

// nolint: gochecknoinits
func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringP(
		"output", "o", "",
		fmt.Sprintf("report output formatter, formats=%v", presenter.AvailableFormats),
	)
	scanCmd.Flags().StringP(
		"file", "f", "",
		"file to write the report output to (default is STDOUT)",
	)
	scanCmd.Flags().StringP("input-type", "i", "",
		fmt.Sprintf("set input type (input type can be %s,%s,%s,%s,%s,%s,%s default:%s)",
			sharedutils.SBOM, sharedutils.DIR, sharedutils.FILE, sharedutils.IMAGE, sharedutils.DOCKERARCHIVE, sharedutils.OCIARCHIVE, sharedutils.OCIDIR, sharedutils.IMAGE))
	scanCmd.Flags().String("application-id", "",
		"ID of a defined application to associate the exported vulnerability scan")
	scanCmd.Flags().BoolP("export", "e", false,
		"export vulnerability scan results to the backend")
	scanCmd.Flags().Bool("cis-docker-benchmark-scan", false,
		"enables CIS docker benchmark scan. (relevant only for image source type)")
	scanCmd.Flags().Bool("ignore-no-fix", false, "ignore vulnerabilities that have no fix available")
	scanCmd.Flags().StringSlice("ignore-vul", []string{}, "ignore list of vulnerabilities")
}

// nolint:cyclop,gocognit
func vulnerabilityScanner(cmd *cobra.Command, args []string) {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		logger.Fatalf("Unable to get output filename: %v", err)
	}

	presenterConfig, err := presenter.CreateConfig(output)
	if err != nil {
		logger.Fatalf("Failed to validate presenter config: %v", err)
	}

	file, err := cmd.Flags().GetString("file")
	if err != nil {
		logger.Fatalf("Unable to get output filename: %v", err)
	}

	inputType, err := cmd.Flags().GetString("input-type")
	if err != nil {
		logger.Fatalf("Unable to get input type: %v", err)
	}

	sourceType, err := sharedutils.ValidateInputType(inputType)
	if err != nil {
		logger.Fatalf("Failed to validate input type: %v", err)
	}

	export, err := cmd.Flags().GetBool("export")
	if err != nil {
		logger.Fatalf("Unable to get export flag: %v", err)
	}

	appID, err := cmd.Flags().GetString("application-id")
	if err != nil {
		logger.Fatalf("Unable to get application ID: %v", err)
	}

	cisDockerBenchmarkEnabled, err := cmd.Flags().GetBool("cis-docker-benchmark-scan")
	if err != nil {
		logger.Fatalf("Unable to get cis-docker-benchmark-scan flag: %v", err)
	}

	var ignores filter.Ignores
	ignores.NoFix, err = cmd.Flags().GetBool("ignore-no-fix")
	if err != nil {
		logger.Fatalf("Unable to get ignore-no-fix flag: %v", err)
	}
	ignores.Vulnerabilities, err = cmd.Flags().GetStringSlice("ignore-vul")
	if err != nil {
		logger.Fatalf("Unable to get ignore-vul flag: %v", err)
	}

	manager := job_manager.New(appConfig.SharedConfig.Scanner.ScannersList, appConfig.SharedConfig, logger, job.Factory)
	results, err := manager.Run(sourceType, args[0])
	if err != nil {
		logger.Fatalf("Failed to run job manager: %v", err)
	}

	// Merge results
	mergedResults := sharedscanner.NewMergedResults()
	for name, result := range results {
		logger.Infof("Merging result from %q", name)
		if res, ok := result.(*sharedscanner.Results); ok {
			mergedResults = mergedResults.Merge(res)
		} else {
			logger.Errorf("Type assertion of result failed.")
		}
	}

	var hash string
	// nolint:exhaustive
	switch sourceType {
	case sharedutils.SBOM:
		bom, err := converter.GetCycloneDXSBOMFromFile(args[0])
		if err != nil {
			logger.Fatal(err)
		}

		bomMetaComponent := bom.Metadata.Component
		hash, err = cdx_helper.GetComponentHash(bomMetaComponent)
		if err != nil {
			logger.Fatalf("Unable to get hash from src BOM: %v", err)
		}
		// If the target and type of source are not defined, we will get them from SBOM.
		// For example in the case of dependency-track.
		mergedResults.SetName(bomMetaComponent.Name)
		mergedResults.SetType(cdx_helper.GetMetaComponentType(*bomMetaComponent))
		mergedResults.SetHash(hash)
	case sharedutils.IMAGE:
		// do nothing
		// grype set the fields of the source during scan
	default:
		hash, err = utils.GenerateHash(sourceType, args[0])
		if err != nil {
			logger.Fatalf("Failed to generate hash for source %s", args[0])
		}
		mergedResults.SetHash(hash)
	}

	writer, closer := getWriter(file)
	defer func() {
		if err := closer(); err != nil {
			log.Warnf("unable to close writer: %+v", err)
		}
	}()

	if len(ignores.Vulnerabilities) > 0 || ignores.NoFix {
		mergedResults = filter.FilterIgnoredVulnerabilities(mergedResults, ignores)
	}

	err = presenter.GetPresenter(presenterConfig, mergedResults).Present(writer)
	if err != nil {
		logger.Fatalf("Failed to present results: %v", err)
	}

	layerCommands, err := getLayerCommandsIfNeeded(sourceType, args[0], appConfig.SharedConfig)
	if err != nil {
		logger.Fatalf("Failed get layer commands. %v", err)
	}

	cisDockerBenchmarkResults, err := getCisDockerBenchmarkResultsIfNeeded(sourceType, args[0], appConfig, cisDockerBenchmarkEnabled)
	if err != nil {
		logger.Fatalf("Failed to get CIS Docker benchmark results: %v", err)
	}

	if export {
		logger.Infof("Exporting vulnerability scan results to the backend: %s", appConfig.Backend.Host)
		apiClient := utils.NewHTTPClient(appConfig.Backend)
		// TODO generate application ID
		if err := _export.Export(apiClient, mergedResults, layerCommands, cisDockerBenchmarkResults, appID); err != nil {
			logger.Errorf("Failed to export vulnerability scan results to the backend: %v", err)
		}
	}
}

func getWriter(filePath string) (io.Writer, func() error) {
	if filePath == "" {
		return os.Stdout, func() error { return nil }
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o666) // nolint:gomnd,gofumpt
	if err != nil {
		logger.Fatalf("Failed open file %s: %v", filePath, err)
	}

	logger.Infof("Writing results to %s", file.Name())

	return file, func() error {
		// nolint:wrapcheck
		return file.Close()
	}
}

func getLayerCommandsIfNeeded(sourceType sharedutils.SourceType, source string, sharedConf *sharedconfig.Config) ([]*image_helper.FsLayerCommand, error) {
	if sourceType != sharedutils.IMAGE {
		return nil, nil
	}
	layerCommands, err := image_helper.GetImageLayerCommands(source, sharedConf)
	if err != nil {
		return nil, fmt.Errorf("failed to get layer commands: %v", err)
	}

	return layerCommands, nil
}

func getCisDockerBenchmarkResultsIfNeeded(sourceType sharedutils.SourceType,
	source string, config *config.Config,
	needed bool,
) (dockle_types.AssessmentMap, error) {
	if sourceType != sharedutils.IMAGE || !needed {
		return nil, nil
	}

	assessmentMap, err := dockle_run.RunFromConfig(createDockleConfig(config, source))
	if err != nil {
		return nil, fmt.Errorf("failed to run dockle: %w", err)
	}

	logger.Infof("Image was scanned for CIS docker benchmark results.")
	logger.Debugf("assessmentMap=%+v", assessmentMap)

	return assessmentMap, nil
}

func createDockleConfig(config *config.Config, source string) *dockle_config.Config {
	var username, password string
	if len(config.SharedConfig.Registry.Auths) > 0 {
		username = config.SharedConfig.Registry.Auths[0].Username
		password = config.SharedConfig.Registry.Auths[0].Password
	}
	return &dockle_config.Config{
		Debug:     log.GetLevel() == log.DebugLevel,
		Timeout:   config.TimeOut,
		Username:  username,
		Password:  password,
		Insecure:  config.SharedConfig.Registry.SkipVerifyTLS,
		NonSSL:    config.SharedConfig.Registry.UseHTTP,
		ImageName: source,
	}
}
