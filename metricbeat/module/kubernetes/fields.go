// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}
