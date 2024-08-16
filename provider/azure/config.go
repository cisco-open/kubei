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

package azure

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

const (
	AMD64 = "x86_64"
	ARM64 = "arm64"

	DefaultEnvPrefix             = "OPENCLARITY_AZURE"
	DefaultScannerVMArchitecture = AMD64
)

var DefaultScannerVMSizeMapping = Mapping{
	AMD64: "Standard_D2s_v3",
	ARM64: "Standard_D2ps_v5",
}

var DefaultScannerImageSKUMapping = Mapping{
	AMD64: "20_04-lts-gen2",
	ARM64: "20_04-lts-arm64",
}

type AzurePublicKey string

func (a *AzurePublicKey) UnmarshalText(text []byte) error {
	if len(text) != 0 {
		publicKey, err := base64.StdEncoding.DecodeString(string(text))
		if err != nil {
			return fmt.Errorf("failed to decode azure scanner public key from base64: %w", err)
		}
		*a = AzurePublicKey(publicKey)
	}
	return nil
}

type Config struct {
	SubscriptionID              string         `mapstructure:"subscription_id"`
	ScannerLocation             string         `mapstructure:"scanner_location"`
	ScannerResourceGroup        string         `mapstructure:"scanner_resource_group"`
	ScannerSubnet               string         `mapstructure:"scanner_subnet_id"`
	ScannerPublicKey            AzurePublicKey `mapstructure:"scanner_public_key"`
	ScannerImagePublisher       string         `mapstructure:"scanner_image_publisher"`
	ScannerImageOffer           string         `mapstructure:"scanner_image_offer"`
	ScannerImageSKUMapping      Mapping        `mapstructure:"scanner_image_sku_mapping"`
	ScannerImageVersion         string         `mapstructure:"scanner_image_version"`
	ScannerSecurityGroup        string         `mapstructure:"scanner_security_group"`
	ScannerStorageAccountName   string         `mapstructure:"scanner_storage_account_name"`
	ScannerStorageContainerName string         `mapstructure:"scanner_storage_container_name"`
	ScannerVMSizeMapping        Mapping        `mapstructure:"scanner_vm_size_mapping"`
	ScannerVMArchitecture       string         `mapstructure:"scanner_vm_architecture"`
}

func NewConfig() (*Config, error) {
	// Avoid modifying the global instance
	v := viper.New()

	v.SetEnvPrefix(DefaultEnvPrefix)
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	_ = v.BindEnv("subscription_id")
	_ = v.BindEnv("scanner_location")
	_ = v.BindEnv("scanner_resource_group")
	_ = v.BindEnv("scanner_subnet_id")
	_ = v.BindEnv("scanner_public_key")
	_ = v.BindEnv("scanner_image_publisher")
	_ = v.BindEnv("scanner_image_offer")

	_ = v.BindEnv("scanner_image_sku_mapping")
	v.SetDefault("scanner_image_sku_mapping", DefaultScannerImageSKUMapping)

	_ = v.BindEnv("scanner_image_version")
	_ = v.BindEnv("scanner_security_group")
	_ = v.BindEnv("scanner_storage_account_name")
	_ = v.BindEnv("scanner_storage_container_name")

	_ = v.BindEnv("scanner_vm_size_mapping")
	v.SetDefault("scanner_vm_size_mapping", DefaultScannerVMSizeMapping)

	_ = v.BindEnv("scanner_vm_architecture")

	config := &Config{}
	if err := v.Unmarshal(&config, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc())); err != nil {
		return nil, fmt.Errorf("failed to parse provider configuration. Provider=Azure: %w", err)
	}
	return config, nil
}

// nolint:cyclop
func (c Config) Validate() error {
	if c.SubscriptionID == "" {
		return errors.New("parameter SubscriptionID must be provided")
	}

	if c.ScannerLocation == "" {
		return errors.New("parameter ScannerLocation must be provided")
	}

	if c.ScannerResourceGroup == "" {
		return errors.New("parameter ScannerResourceGroup must be provided")
	}

	if c.ScannerSubnet == "" {
		return errors.New("parameter ScannerSubnet must be provided")
	}

	if c.ScannerVMArchitecture == "" {
		return errors.New("parameter ScannerVmArchitecture must be provided")
	}

	if _, ok := c.ScannerVMSizeMapping[c.ScannerVMArchitecture]; !ok {
		return fmt.Errorf("failed to find vm size for architecture %s", c.ScannerVMArchitecture)
	}

	if c.ScannerImagePublisher == "" {
		return errors.New("parameter ScannerImagePublisher must be provided")
	}

	if c.ScannerImageOffer == "" {
		return errors.New("parameter ScannerImageOffer must be provided")
	}

	if _, ok := c.ScannerImageSKUMapping[c.ScannerVMArchitecture]; !ok {
		return fmt.Errorf("failed to find image sku for architecture %s", c.ScannerVMArchitecture)
	}

	if c.ScannerImageVersion == "" {
		return errors.New("parameter ScannerImageVersion must be provided")
	}

	if c.ScannerSecurityGroup == "" {
		return errors.New("parameter ScannerSecurityGroup must be provided")
	}

	if c.ScannerStorageAccountName == "" {
		return errors.New("parameter ScannerStorageAccountName must be provided")
	}

	if c.ScannerStorageContainerName == "" {
		return errors.New("parameter ScannerStorageContainerName must be provided")
	}

	return nil
}

type Mapping map[string]string

func (m *Mapping) UnmarshalText(text []byte) error {
	mapping := make(Mapping)
	items := strings.Split(string(text), ",")

	numOfParts := 2
	for _, item := range items {
		pair := strings.Split(item, ":")
		if len(pair) != numOfParts {
			continue
		}

		switch pair[0] {
		case AMD64, ARM64:
			mapping[pair[0]] = pair[1]
		default:
			return fmt.Errorf("unsupported architecture: %s", pair[0])
		}
	}
	*m = mapping

	return nil
}
