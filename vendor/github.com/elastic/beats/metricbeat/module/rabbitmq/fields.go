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

package rabbitmq

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "rabbitmq", asset.ModuleFieldsPri, AssetRabbitmq); err != nil {
		panic(err)
	}
}

// AssetRabbitmq returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/rabbitmq.
func AssetRabbitmq() string {
	return "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"
}
