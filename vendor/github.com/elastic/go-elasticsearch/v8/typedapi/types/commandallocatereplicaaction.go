// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// CommandAllocateReplicaAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/cluster/reroute/types.ts#L69-L76
type CommandAllocateReplicaAction struct {
	Index string `json:"index"`
	Node  string `json:"node"`
	Shard int    `json:"shard"`
}

// NewCommandAllocateReplicaAction returns a CommandAllocateReplicaAction.
func NewCommandAllocateReplicaAction() *CommandAllocateReplicaAction {
	r := &CommandAllocateReplicaAction{}

	return r
}
