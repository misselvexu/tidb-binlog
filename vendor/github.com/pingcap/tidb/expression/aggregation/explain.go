// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package aggregation

import (
	"bytes"
	"fmt"
)

// ExplainAggFunc generates explain information for a aggregation function.
func ExplainAggFunc(agg *AggFuncDesc) string {
	buffer := bytes.NewBufferString(fmt.Sprintf("%s(", agg.Name))
	if agg.HasDistinct {
		buffer.WriteString("distinct ")
	}
	for i, arg := range agg.Args {
		buffer.WriteString(arg.ExplainInfo())
		if i+1 < len(agg.Args) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}