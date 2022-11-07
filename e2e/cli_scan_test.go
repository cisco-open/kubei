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
	"gotest.tools/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"

	"github.com/openclarity/kubeclarity/api/client/models"
	"github.com/openclarity/kubeclarity/e2e/common"
	"github.com/openclarity/kubeclarity/shared/pkg/formatter"
)

const (
	DirectoryAnalyzeOutputSBOMFile = "dir.sbom"
	ImageAnalyzeOutputSBOMFile     = "merged.sbom"
	TestImageName                  = "nginx:1.10"
	ApplicationName                = "test-app"
)

func TestCLIScan(t *testing.T) {
	t.Logf("Starting test CLI scan")

	stopCh := make(chan struct{})
	defer func() {
		stopCh <- struct{}{}
		time.Sleep(2 * time.Second)
	}()
	f1 := features.New("cli scan flow - analyze and scan").
		WithLabel("type", "cli").
		Assess("cli scan flow", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			// setup env
			t.Logf("setup env...")
			setupCLIScanTestEnv(stopCh)

			// create application
			t.Logf("create application...")
			appID := createApplication(t)

			// analyze dir
			t.Logf("analyze dir...")
			analyzeDir(t)
			validateAnalyzeDir(t)

			// analyze image with --merge-sbom directory sbom, and export to backend
			t.Logf("analyze image...")
			analyzeImage(t, DirectoryAnalyzeOutputSBOMFile, appID)
			validateAnalyzeImage(t)

			// scan merged sbom
			t.Logf("scan merged sbom...")
			scanSBOM(t, ImageAnalyzeOutputSBOMFile, appID)
			validateScanSBOM(t)

			// scan image
			t.Logf("scan image...")
			scanImage(t, TestImageName, appID)
			validateScanImage(t)

			return ctx
		}).Feature()

	// test features
	testenv.Test(t, f1)
}

func getCdxSbom(t *testing.T, fileName string) *cdx.BOM {
	t.Helper()
	sbom, err := os.ReadFile(fileName)
	assert.NilError(t, err)
	input := formatter.New(formatter.CycloneDXFormat, sbom)
	assert.NilError(t, input.Decode(formatter.CycloneDXFormat))
	return input.GetSBOM().(*cdx.BOM)
}

func validateAnalyzeDir(t *testing.T) {
	t.Helper()
	sbom := getCdxSbom(t, DirectoryAnalyzeOutputSBOMFile)
	assert.Assert(t, sbom != nil)
	assert.Assert(t, sbom.Components != nil)
	assert.Assert(t, len(*sbom.Components) > 0)
}

func validateAnalyzeImage(t *testing.T) {
	t.Helper()
	sbom := getCdxSbom(t, ImageAnalyzeOutputSBOMFile)
	assert.Assert(t, sbom != nil)
	// check generated sbom
	assert.Assert(t, sbom.Components != nil)
	assert.Assert(t, len(*sbom.Components) > 0)

	// check export to db
	packages := common.GetPackages(t, kubeclarityAPI)
	assert.Assert(t, *packages.Total > 0)

	appResources := common.GetApplicationResources(t, kubeclarityAPI)
	assert.Assert(t, *appResources.Total > 0)
}

func validateScanImage(t *testing.T) {
	t.Helper()
	vuls := common.GetVulnerabilities(t, kubeclarityAPI)
	assert.Assert(t, *vuls.Total > 0)

	appResources := common.GetApplicationResources(t, kubeclarityAPI)
	assert.Assert(t, appResources.Items[0].ResourceType == models.ResourceTypeIMAGE)

	cisDockerBenchmarkResults := common.GetCISDockerBenchmarkResults(t, kubeclarityAPI, appResources.Items[0].ID)
	assert.Assert(t, *cisDockerBenchmarkResults.Total > 0)
}

func validateScanSBOM(t *testing.T) {
	t.Helper()
	vuls := common.GetVulnerabilities(t, kubeclarityAPI)

	// TODO how to validate that vulnerabilities were added on top of scanned sbom vuls
	assert.Assert(t, *vuls.Total > 0)
}

func createApplication(t *testing.T) (appID string) {
	t.Helper()
	appType := models.ApplicationTypePOD
	res := common.PostApplications(t, kubeclarityAPI, &models.ApplicationInfo{
		Name: common.StringPtr(ApplicationName),
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
func analyzeImage(t *testing.T, inputSbom string, appID string) {
	t.Helper()
	assert.NilError(t, os.Setenv("BACKEND_HOST", "localhost:"+common.KubeClarityPortForwardHostPort))
	assert.NilError(t, os.Setenv("BACKEND_DISABLE_TLS", "true"))

	command := fmt.Sprintf("%v analyze %v --application-id %v --input-type image --merge-sbom %v -e -o %v", cliPath, TestImageName, appID, inputSbom, ImageAnalyzeOutputSBOMFile)

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("analyzeImage failed. failed to execute command. %v, %s", err, out)
	}
}

func scanSBOM(t *testing.T, inputSbom string, appID string) {
	t.Helper()
	assert.NilError(t, os.Setenv("BACKEND_HOST", "localhost:"+common.KubeClarityPortForwardHostPort))
	assert.NilError(t, os.Setenv("BACKEND_DISABLE_TLS", "true"))

	command := fmt.Sprintf("%v scan %v --application-id %v --input-type sbom -e", cliPath, inputSbom, appID)

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("scanSBOM failed. failed to execute command. %v, %s", err, out)
	}
}

func scanImage(t *testing.T, image string, appID string) {
	t.Helper()
	assert.NilError(t, os.Setenv("BACKEND_HOST", "localhost:"+common.KubeClarityPortForwardHostPort))
	assert.NilError(t, os.Setenv("BACKEND_DISABLE_TLS", "true"))

	command := fmt.Sprintf("%v scan %v --application-id %v --input-type image --cis-docker-benchmark-scan -e", cliPath, image, appID)

	cmd := exec.Command("/bin/sh", "-c", command)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("scanImage failed. failed to execute command. %v, %s", err, out)
	}
}

func setupCLIScanTestEnv(stopCh chan struct{}) {
	println("Set up cli scan test env...")

	println("port-forward to kubeclarity...")
	common.PortForwardToKubeClarity(stopCh)
}
