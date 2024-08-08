// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
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

package types

type Result struct {
	Misconfigurations []FlattenedMisconfiguration `json:"Misconfigurations"`
}

func NewResult() *Result {
	return &Result{}
}

func (r *Result) Merge(scan *ScannerResult) {
	if scan == nil {
		return
	}

	// Merge all misconfigurations
	for _, misconfiguration := range scan.Misconfigurations {
		r.Misconfigurations = append(r.Misconfigurations, FlattenedMisconfiguration{
			Misconfiguration: misconfiguration,
			ScannerName:      scan.ScannerName,
		})
	}
}
