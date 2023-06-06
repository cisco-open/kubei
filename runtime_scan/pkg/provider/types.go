// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
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

package provider

import (
	"context"

	"github.com/openclarity/vmclarity/api/models"
)

type Provider interface {
	// Kind returns models.CloudProvider
	Kind() models.CloudProvider

	Discoverer
	Scanner
}

type Discoverer interface {
	// DiscoverScopes returns a list of discovered models.Scopes
	DiscoverScopes(ctx context.Context) (*models.Scopes, error)
	// DiscoverTargets returns list of TargetType in ScanScopeType
	DiscoverTargets(ctx context.Context, scanScope *models.ScanScopeType) ([]models.TargetType, error)
}

type Scanner interface {
	// RunTargetScan is a non-blocking call which takes a ScanJobConfig and creates resources for performing Scan.
	// It may return FatalError or RetryableError to indicate if the error is permanent or transient.
	// It is expected to return RetryableError in case the resources are not ready or are still being created.
	// It must return nil if all the resources are created and ready.
	// It also must be idempotent.
	RunTargetScan(context.Context, *ScanJobConfig) error
	// RemoveTargetScan is a non-blocking call which takes a ScanJobConfig and remove resources created for Scan.
	// It may return FatalError or RetryableError to indicate if the error is permanent or transient.
	// It is expected to return RetryableError in case the resources are still being deleted.
	// It must return nil if all the resources are deleted.
	// It also must be idempotent.
	RemoveTargetScan(context.Context, *ScanJobConfig) error
}

type ScanMetadata struct {
	ScanID       string
	ScanResultID string
	TargetID     string
}

type ScanJobConfig struct {
	ScannerImage     string // Scanner Container Image to use containing the vmclarity-cli and tools
	ScannerCLIConfig string // Scanner CLI config yaml (families config yaml)
	VMClarityAddress string // The backend address for the scanner CLI to export too
	KeyPairName      string // The name of the key pair to set on the instance, ignored if not set, used mainly for debugging.
	Region           string // The region where the VMClarity server is deployed
	BlockDeviceName  string // The block device name used for attaching target volume to the scanner vm

	ScanMetadata
	models.ScannerInstanceCreationConfig
	models.Target
}
