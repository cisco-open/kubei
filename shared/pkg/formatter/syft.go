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

package formatter

import (
	"bytes"
	"fmt"

	"github.com/anchore/syft/syft"
	syft_sbom "github.com/anchore/syft/syft/sbom"
)

const (
	SyftFormat = "json"
)

type Syft struct {
	name       string
	sbomBytes  []byte         // encoded BOM bytes
	sbomStruct syft_sbom.SBOM // decoded syft BOM struct
}

func newSyftFormatter(sbomBytes []byte) Formatter {
	return &Syft{
		name:       SyftFormat,
		sbomBytes:  sbomBytes,
		sbomStruct: syft_sbom.SBOM{},
	}
}

func (f *Syft) SetSBOM(bom interface{}) error {
	sbom, ok := bom.(syft_sbom.SBOM)
	if !ok {
		return fmt.Errorf("failed to set BOM in syft formatter")
	}
	f.sbomStruct = sbom

	return nil
}

func (f *Syft) GetSBOM() interface{} {
	return f.sbomStruct
}

func (f *Syft) GetSBOMBytes() []byte {
	return f.sbomBytes
}

func (f *Syft) Encode(format string) error {
	// TODO maybe check cyclonedx format as output
	sbomFormat := syft.FormatByName(format)
	var err error
	f.sbomBytes, err = syft.Encode(f.sbomStruct, sbomFormat)
	if err != nil {
		return fmt.Errorf("failed to encode SBOM: %v", err)
	}
	return nil
}

func (f *Syft) Decode(format string) error {
	reader := bytes.NewReader(f.sbomBytes)
	sbom, _, err := syft.Decode(reader)
	if err != nil {
		return fmt.Errorf("failed to decode SBOM: %w", err)
	}

	if err := f.SetSBOM(*sbom); err != nil {
		return fmt.Errorf("unable to set BOM data: %v", err)
	}

	return nil
}
