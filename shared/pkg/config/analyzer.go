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

package config

import (
	"github.com/spf13/viper"

	"github.com/openclarity/kubeclarity/shared/v2/pkg/formatter"
)

// TODO: maybe we need to extend the unified config.
type Analyzer struct {
	OutputFormat string              `yaml:"output_format" mapstructure:"output_format"`
	AnalyzerList []string            `yaml:"analyzer_list" mapstructure:"analyzer_list"`
	Scope        string              `yaml:"scope" mapstructure:"scope"`
	TrivyConfig  AnalyzerTrivyConfig `yaml:"trivy_config" mapstructure:"trivy_config"`
}

const (
	AnalyzerList  = "ANALYZER_LIST"
	AnalyzerScope = "ANALYZER_SCOPE"
	OutputFormat  = "ANALYZER_OUTPUT_FORMAT"
)

func setAnalyzerConfigDefaults() {
	viper.SetDefault(AnalyzerList, []string{"syft", "gomod"})
	viper.SetDefault(AnalyzerScope, "squashed")
	viper.SetDefault(OutputFormat, formatter.CycloneDXJSONFormat)
}

func LoadAnalyzerConfig() *Analyzer {
	setAnalyzerConfigDefaults()
	return &Analyzer{
		OutputFormat: viper.GetString(OutputFormat),
		AnalyzerList: viper.GetStringSlice(AnalyzerList),
		Scope:        viper.GetString(AnalyzerScope),
		TrivyConfig:  LoadAnalyzerTrivyConfig(),
	}
}
