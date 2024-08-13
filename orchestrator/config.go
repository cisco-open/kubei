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

package orchestrator

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"

	apitypes "github.com/openclarity/openclarity/api/types"
	discovery "github.com/openclarity/openclarity/orchestrator/discoverer"
	assetscanprocessor "github.com/openclarity/openclarity/orchestrator/processor/assetscan"
	assetscanwatcher "github.com/openclarity/openclarity/orchestrator/watcher/assetscan"
	assetscanestimationwatcher "github.com/openclarity/openclarity/orchestrator/watcher/assetscanestimation"
	scanwatcher "github.com/openclarity/openclarity/orchestrator/watcher/scan"
	scanconfigwatcher "github.com/openclarity/openclarity/orchestrator/watcher/scanconfig"
	scanestimationwatcher "github.com/openclarity/openclarity/orchestrator/watcher/scanestimation"
)

const (
	DefaultEnvPrefix = "OPENCLARITY_ORCHESTRATOR"

	// Approximately half the polling delay to allow for more efficient
	// reconcile cascading e.g. reconciling scan config creates a scan, the
	// poller for scan is offset by 7 seconds so should pick up the new
	// scan after 7 seconds instead of the full poller time.
	DefaultControllerStartupDelay = 7 * time.Second
	DefaultProviderKind           = apitypes.AWS
	DefaultHealthCheckAddress     = ":8082"
)

type Config struct {
	ProviderKind apitypes.CloudProvider `json:"provider,omitempty" mapstructure:"provider"`

	APIServerAddress   string `json:"apiserver-address,omitempty" mapstructure:"apiserver_address"`
	HealthCheckAddress string `json:"healthcheck-address,omitempty" mapstructure:"healthcheck_address"`

	// The Orchestrator starts the Controller(s) in a sequence and the ControllerStartupDelay is used for waiting
	// before starting each Controller to avoid them hitting the API at the same time and allow one Controller
	// to pick up an event generated by the other without waiting until the next polling cycle.
	ControllerStartupDelay time.Duration `json:"controller-startup-delay,omitempty" mapstructure:"controller_startup_delay"`

	DiscoveryConfig                  discovery.Config                  `json:"discovery,omitempty" mapstructure:"discovery"`
	ScanConfigWatcherConfig          scanconfigwatcher.Config          `json:"scanconfig-watcher,omitempty" mapstructure:"scanconfig_watcher"`
	ScanWatcherConfig                scanwatcher.Config                `json:"scan-watcher,omitempty" mapstructure:"scan_watcher"`
	AssetScanWatcherConfig           assetscanwatcher.Config           `json:"assetscan-watcher,omitempty" mapstructure:"assetscan_watcher"`
	AssetScanEstimationWatcherConfig assetscanestimationwatcher.Config `json:"assetscan-estimation-watcher,omitempty" mapstructure:"assetscan_estimation_watcher"`
	ScanEstimationWatcherConfig      scanestimationwatcher.Config      `json:"scan-estimation-watcher,omitempty" mapstructure:"scan_estimation_watcher"`
	AssetScanProcessorConfig         assetscanprocessor.Config         `json:"assetscan-processor,omitempty" mapstructure:"assetscan_processor"`
}

func NewConfig() (*Config, error) {
	v := viper.NewWithOptions(
		viper.KeyDelimiter("."),
		viper.EnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")),
	)

	v.SetEnvPrefix(DefaultEnvPrefix)
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Orchestrator configuration
	_ = v.BindEnv("provider")
	v.SetDefault("provider", DefaultProviderKind)

	_ = v.BindEnv("apiserver_address")

	_ = v.BindEnv("healthcheck_address")
	v.SetDefault("healthcheck_address", DefaultHealthCheckAddress)

	_ = v.BindEnv("controller_startup_delay")
	v.SetDefault("controller_startup_delay", DefaultControllerStartupDelay)

	// Discovery Controller configuration
	_ = v.BindEnv("discovery.interval")
	v.SetDefault("discovery.interval", discovery.DefaultInterval)

	// ScanConfig Watcher Controller configuration
	_ = v.BindEnv("scanconfig_watcher.poll_period")
	v.SetDefault("scanconfig_watcher.poll_period", scanconfigwatcher.DefaultPollInterval)

	_ = v.BindEnv("scanconfig_watcher.reconcile_timeout")
	v.SetDefault("scanconfig_watcher.reconcile_timeout", scanconfigwatcher.DefaultReconcileTimeout)

	// Scan Watcher Controller configuration
	_ = v.BindEnv("scan_watcher.poll_period")
	v.SetDefault("scan_watcher.poll_period", scanwatcher.DefaultPollInterval)

	_ = v.BindEnv("scan_watcher.reconcile_timeout")
	v.SetDefault("scan_watcher.reconcile_timeout", scanwatcher.DefaultReconcileTimeout)

	_ = v.BindEnv("scan_watcher.scan_timeout")
	v.SetDefault("scan_watcher.scan_timeout", scanwatcher.DefaultScanTimeout)

	// AssetScan Watcher Controller configuration
	_ = v.BindEnv("assetscan_watcher.poll_period")
	v.SetDefault("assetscan_watcher.poll_period", assetscanwatcher.DefaultPollInterval)

	_ = v.BindEnv("assetscan_watcher.reconcile_timeout")
	v.SetDefault("assetscan_watcher.reconcile_timeout", assetscanwatcher.DefaultReconcileTimeout)

	_ = v.BindEnv("assetscan_watcher.abort_timeout")
	v.SetDefault("assetscan_watcher.abort_timeout", assetscanwatcher.DefaultAbortTimeout)

	_ = v.BindEnv("assetscan_watcher.delete_policy")
	v.SetDefault("assetscan_watcher.delete_policy", assetscanwatcher.DeleteJobPolicyAlways)

	// Scanner configuration
	_ = v.BindEnv("assetscan_watcher.scanner.apiserver_address")

	_ = v.BindEnv("assetscan_watcher.scanner.exploitsdb_address")

	_ = v.BindEnv("assetscan_watcher.scanner.trivy_server_address")

	_ = v.BindEnv("assetscan_watcher.scanner.trivy_scan_timeout")
	v.SetDefault("assetscan_watcher.scanner.trivy_scan_timeout", assetscanwatcher.DefaultTrivyScanTimeout)

	_ = v.BindEnv("assetscan_watcher.scanner.grype_server_address")

	_ = v.BindEnv("assetscan_watcher.scanner.grype_server_timeout")
	v.SetDefault("assetscan_watcher.scanner.grype_server_timeout", assetscanwatcher.DefaultGrypeServerTimeout)

	_ = v.BindEnv("assetscan_watcher.scanner.yara_rule_server_address")

	_ = v.BindEnv("assetscan_watcher.scanner.container_image")

	_ = v.BindEnv("assetscan_watcher.scanner.freshclam_mirror")

	// AssetScan Estimation Watcher Controller configuration
	_ = v.BindEnv("assetscan_estimation_watcher.poll_period")
	v.SetDefault("assetscan_estimation_watcher.poll_period", assetscanestimationwatcher.DefaultPollInterval)

	_ = v.BindEnv("assetscan_estimation_watcher.reconcile_timeout")
	v.SetDefault("assetscan_estimation_watcher.reconcile_timeout", assetscanestimationwatcher.DefaultReconcileTimeout)

	// Scan Estimation Watcher Controller configuration
	_ = v.BindEnv("scan_estimation_watcher.poll_period")
	v.SetDefault("scan_estimation_watcher.poll_period", scanestimationwatcher.DefaultPollInterval)

	_ = v.BindEnv("scan_estimation_watcher.reconcile_timeout")
	v.SetDefault("scan_estimation_watcher.reconcile_timeout", scanestimationwatcher.DefaultReconcileTimeout)

	_ = v.BindEnv("scan_estimation_watcher.estimation_timeout")
	v.SetDefault("scan_estimation_watcher.estimation_timeout", scanestimationwatcher.DefaultScanEstimationTimeout)

	// AssetScan Processor Controller configuration
	_ = v.BindEnv("assetscan_processor.poll_period")
	v.SetDefault("assetscan_processor.poll_period", assetscanprocessor.DefaultPollInterval)

	_ = v.BindEnv("assetscan_processor.reconcile_timeout")
	v.SetDefault("assetscan_processor.reconcile_timeout", assetscanprocessor.DefaultReconcileTimeout)

	decodeHooks := mapstructure.ComposeDecodeHookFunc(
		// TextUnmarshallerHookFunc is needed to decode custom types
		mapstructure.TextUnmarshallerHookFunc(),
		// Default decoders
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	)

	config := &Config{}
	if err := v.Unmarshal(config, viper.DecodeHook(decodeHooks)); err != nil {
		return nil, fmt.Errorf("failed to load Orchestrator configuration: %w", err)
	}

	return config, nil
}
