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


// Package sortmode
package sortmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/_types/sort.ts#L103-L112
type SortMode struct {
	Name string
}

var (
	Min = SortMode{"min"}

	Max = SortMode{"max"}

	Sum = SortMode{"sum"}

	Avg = SortMode{"avg"}

	Median = SortMode{"median"}
)

func (s SortMode) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SortMode) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "min":
		*s = Min
	case "max":
		*s = Max
	case "sum":
		*s = Sum
	case "avg":
		*s = Avg
	case "median":
		*s = Median
	default:
		*s = SortMode{string(text)}
	}

	return nil
}

func (s SortMode) String() string {
	return s.Name
}
