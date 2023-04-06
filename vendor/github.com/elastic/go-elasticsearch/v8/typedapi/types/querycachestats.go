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

// QueryCacheStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/_types/Stats.ts#L150-L159
type QueryCacheStats struct {
	CacheCount        int       `json:"cache_count"`
	CacheSize         int       `json:"cache_size"`
	Evictions         int       `json:"evictions"`
	HitCount          int       `json:"hit_count"`
	MemorySize        *ByteSize `json:"memory_size,omitempty"`
	MemorySizeInBytes int       `json:"memory_size_in_bytes"`
	MissCount         int       `json:"miss_count"`
	TotalCount        int       `json:"total_count"`
}

// NewQueryCacheStats returns a QueryCacheStats.
func NewQueryCacheStats() *QueryCacheStats {
	r := &QueryCacheStats{}

	return r
}
