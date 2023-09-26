// Copyright © 2021 Cisco Systems, Inc. and its affiliates.
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

package e2e

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"gotest.tools/assert"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"

	"github.com/openclarity/kubeclarity/api/client/models"
	"github.com/openclarity/kubeclarity/e2e/common"
	"github.com/openclarity/kubeclarity/shared/pkg/converter"
)

const (
	DirectoryAnalyzeOutputSBOMFile = "dir.sbom"
	ImageAnalyzeOutputSBOMFile     = "merged.sbom"
	TestImageName                  = "nginx:1.10"
	ApplicationName                = "test-app"

	DockerArchiveApplicationName = "test-app-docker-archive"
	DockerArchiveOutputSBOMFile  = "docker-archive.sbom"

	TestImageWithMissingSyftMetadata      = "docker.io/weaveworksdemos/front-end:sha-14254f9"
	MissingMetaImageAnalyzeOutputSBOMFile = "missingmeta.sbom"
	MissingMetaApplicationName            = "test-app-missingm"

	TestImageWithNoComponents              = "gcr.io/google_containers/pause@sha256:927d98197ec1141a368550822d18fa1c60bdae27b78b0c004f705f548c07814f"
	NoComponentsImageAnalyzeOutputSBOMFile = "nocomponents.sbom"
	NoComponentsApplicationName            = "test-app-nocompos"
)

func TestCLIScan(t *testing.T) {
	t.Logf("Starting test CLI scan")

	stopCh := make(chan struct{})
	defer func() {
		stopCh <- struct{}{}
		time.Sleep(2 * time.Second)
	}()
	cliScanFlowFunc := func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		// create application
		t.Logf("create application...")
		appID := createApplication(t, ApplicationName)

		// analyze dir
		t.Logf("analyze dir...")
		analyzeDir(t)
		validateAnalyzeDir(t)

		// analyze image with --merge-sbom directory sbom, and export to backend
		t.Logf("analyze image...")
		analyzeImage(t, DirectoryAnalyzeOutputSBOMFile, appID, TestImageName, "image", ImageAnalyzeOutputSBOMFile)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateAnalyzeImage(t, ImageAnalyzeOutputSBOMFile, appID)

		// scan merged sbom
		t.Logf("scan merged sbom...")
		scanSBOM(t, ImageAnalyzeOutputSBOMFile, appID)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateScanSBOM(t, appID)

		// scan image
		t.Logf("scan image...")
		scanImage(t, TestImageName, "image", appID)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateScanImage(t, appID, true)

		return ctx
	}
	cliScanFlowImageWithKnownBadMetadata := func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		// create application
		t.Logf("create application...")
		appID := createApplication(t, MissingMetaApplicationName)

		// analyze "bad" image
		t.Logf("analyze image...")
		analyzeImage(t, "", appID, TestImageWithMissingSyftMetadata, "image", MissingMetaImageAnalyzeOutputSBOMFile)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateAnalyzeImage(t, MissingMetaImageAnalyzeOutputSBOMFile, appID)

		// scan merged sbom
		t.Logf("scan merged sbom...")
		scanSBOM(t, MissingMetaImageAnalyzeOutputSBOMFile, appID)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateScanSBOM(t, appID)

		return ctx
	}
	cliScanFlowImageWithNoComponents := func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		// create application
		t.Logf("create application...")
		appID := createApplication(t, NoComponentsApplicationName)

		// analyze "bad" image
		t.Logf("analyze image...")
		analyzeImage(t, "", appID, TestImageWithNoComponents, "image", NoComponentsImageAnalyzeOutputSBOMFile)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)

		// check generated sbom is a valid cyclonedx even
		// though there is no components
		sbom := getCdxSbom(t, NoComponentsImageAnalyzeOutputSBOMFile)
		assert.Assert(t, sbom != nil)
		assert.Assert(t, sbom.Components != nil)
		assert.Assert(t, len(*sbom.Components) == 0)

		// scan sbom with no componenents
		t.Logf("scan merged sbom...")
		scanSBOM(t, NoComponentsImageAnalyzeOutputSBOMFile, appID)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)

		// validate scan results were sent to the backend but
		// no vuls were found because its got no components
		vuls := common.GetVulnerabilities(t, kubeclarityAPI, appID)
		assert.Assert(t, *vuls.Total == 0)

		return ctx
	}
	cliScanFlowDockerArchiveImage := func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		// create application
		t.Logf("create application...")
		appID := createApplication(t, DockerArchiveApplicationName)

		tmpDir, err := os.MkdirTemp("", "")
		if err != nil {
			t.Fatalf("unable to make temporary directory")
		}
		defer func() {
			err := os.RemoveAll(tmpDir)
			if err != nil {
				t.Logf("unable to remove temp directory: %v", err)
			}
		}()

		outputFile := filepath.Join(tmpDir, "image.tar")

		command := fmt.Sprintf("docker pull %s && docker image save %s -o %s", TestImageName, TestImageName, outputFile)
		cmd := exec.Command("/bin/sh", "-c", command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("docker save failed: %v, %s", err, out)
		}

		// analyze image and export to backend
		t.Logf("analyze image...")
		analyzeImage(t, "", appID, outputFile, "docker-archive", DockerArchiveOutputSBOMFile)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateAnalyzeImage(t, DockerArchiveOutputSBOMFile, appID)

		// scan image
		t.Logf("scan image...")
		scanImage(t, outputFile, "docker-archive", appID)
		time.Sleep(common.WaitForMaterializedViewRefreshSecond * time.Second)
		validateScanImage(t, appID, false)

		return ctx
	}
	f1 := features.New("cli scan flow - analyze and scan - default configuration").
		WithLabel("type", "cli").
		WithSetup("setup env", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			// setup env
			t.Logf("setup env...")
			setupCLIScanTestEnv(t, stopCh, "", "")

			return ctx
		}).
		Assess("cli scan flow", cliScanFlowFunc).
		Assess("cli scan flow - image with known bad metadata in cyclonedx", cliScanFlowImageWithKnownBadMetadata).
		Assess("cli scan flow - image with no components", cliScanFlowImageWithNoComponents).
		Assess("cli scan flow - docker archive image", cliScanFlowDockerArchiveImage).
		Feature()

	f2 := features.New("cli scan flow - analyze and scan - trivy").
		WithLabel("type", "cli").
		WithSetup("setup env", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			// setup env
			t.Logf("setup env...")
			setupCLIScanTestEnv(t, stopCh, "trivy", "trivy")

			return ctx
		}).
		Assess("cli scan flow", cliScanFlowFunc).
		Assess("cli scan flow - image with known bad metadata in cyclonedx", cliScanFlowImageWithKnownBadMetadata).
		Assess("cli scan flow - image with no components", cliScanFlowImageWithNoComponents).
		Assess("cli scan flow - docker archive image", cliScanFlowDockerArchiveImage).
		Feature()

	// test features
	testenv.Test(t, f1, f2)
}

func getCdxSbom(t *testing.T, fileName string) *cdx.BOM {
	t.Helper()
	sbom, err := converter.GetCycloneDXSBOMFromFile(fileName)
	assert.NilError(t, err)
	return sbom
}

func validateAnalyzeDir(t *testing.T) {
	t.Helper()
	sbom := getCdxSbom(t, DirectoryAnalyzeOutputSBOMFile)
	assert.Assert(t, sbom != nil)
	assert.Assert(t, sbom.Components != nil)
	assert.Assert(t, len(*sbom.Components) > 0)
}

func validateAnalyzeImage(t *testing.T, sbomFile, appID string) {
	t.Helper()
	sbom := getCdxSbom(t, sbomFile)
	assert.Assert(t, sbom != nil)
	// check generated sbom
	assert.Assert(t, sbom.Components != nil)
	assert.Assert(t, len(*sbom.Components) > 0)

	// check export to db
	packages := common.GetPackages(t, kubeclarityAPI, appID)
	assert.Assert(t, *packages.Total > 0)

	appResources := common.GetApplicationResources(t, kubeclarityAPI, appID)
	assert.Assert(t, *appResources.Total > 0)
}

func validateScanImage(t *testing.T, appID string, cis bool) {
	t.Helper()
	vuls := common.GetVulnerabilities(t, kubeclarityAPI, appID)
	assert.Assert(t, *vuls.Total > 0)

	appResources := common.GetApplicationResources(t, kubeclarityAPI, appID)
	assert.Assert(t, appResources.Items[0].ResourceType == models.ResourceTypeIMAGE)

	if cis {
		cisDockerBenchmarkResults := common.GetCISDockerBenchmarkResults(t, kubeclarityAPI, appResources.Items[0].ID)
		assert.Assert(t, *cisDockerBenchmarkResults.Total > 0)
	}
}

func validateScanSBOM(t *testing.T, appID string) {
	t.Helper()
	vuls := common.GetVulnerabilities(t, kubeclarityAPI, appID)

	// TODO how to validate that vulnerabilities were added on top of scanned sbom vuls
	assert.Assert(t, *vuls.Total > 0)
}

func createApplication(t *testing.T, applicationName string) (appID string) {
	t.Helper()
	appType := models.ApplicationTypePOD
	res := common.PostApplications(t, kubeclarityAPI, &models.ApplicationInfo{
		Name: common.StringPtr(applicationName),
		Type: &appType,
	})

	appID = res.Payload.ID
	return
}

func analyzeDir(t *testing.T) {
	t.Helper()
	dirPath := filepath.Join(common.GetCurrentDir(), "vulnerable")

	command := cliPath + " analyze " + dirPath + " --input-type dir -o " + DirectoryAnalyzeOutputSBOMFile

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("analyzeDir failed. failed to execute command. %v, %s", err, out)
	}
}

var cliPath = filepath.Join(common.GetCurrentDir(), "kubeclarity-cli")

// analyze test image, merge inputSbom and export to backend
func analyzeImage(t *testing.T, inputSbom string, appID string, image string, inputType string, outputfile string) {
	command := fmt.Sprintf("%v analyze %v --application-id %v --input-type %s", cliPath, image, appID, inputType)

	if inputSbom != "" {
		command = fmt.Sprintf("%s --merge-sbom %v", command, inputSbom)
	}

	command = fmt.Sprintf("%s -e -o %v", command, outputfile)

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("analyzeImage failed. failed to execute command. %v, %s", err, out)
	}
}

func scanSBOM(t *testing.T, inputSbom string, appID string) {
	command := fmt.Sprintf("%v scan %v --application-id %v --input-type sbom -e", cliPath, inputSbom, appID)

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("scanSBOM failed. failed to execute command. %v, %s", err, out)
	}
}

func scanImage(t *testing.T, image string, inputType string, appID string) {

	command := fmt.Sprintf("%v scan %v --application-id %v --input-type %s --cis-docker-benchmark-scan -e", cliPath, image, appID, inputType)

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("scanImage failed. failed to execute command. %v, %s", err, out)
	}
}

func setupCLIScanTestEnv(t *testing.T, stopCh chan struct{}, analyzersList string, scannersList string) {
	println("Set up cli scan test env...")
	t.Helper()
	assert.NilError(t, os.Setenv("BACKEND_HOST", "localhost:"+common.KubeClarityPortForwardHostPort))
	assert.NilError(t, os.Setenv("BACKEND_DISABLE_TLS", "true"))
	if analyzersList != "" {
		assert.NilError(t, os.Setenv("ANALYZER_LIST", analyzersList))
	}
	if scannersList != "" {
		assert.NilError(t, os.Setenv("SCANNERS_LIST", scannersList))
	}

	println("port-forward to kubeclarity...")
	common.PortForwardToKubeClarity(stopCh)
}
