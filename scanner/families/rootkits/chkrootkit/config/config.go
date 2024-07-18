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

const (
	DefaultChkrootkitBinaryPath = "chkrootkit"
)

type Config struct {
	BinaryPath string `yaml:"binary_path" mapstructure:"binary_path" json:"binary_path"`
}

func (c *Config) GetBinaryPath() string {
	if c.BinaryPath != "" {
		return c.BinaryPath
	}

	return DefaultChkrootkitBinaryPath
}
