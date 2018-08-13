// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datastore

import (
	"fmt"
	"github.com/ligato/cn-infra/logging"
)

// SimpleReport holds error/warning messages recorded during data collection /
// validation
type SimpleReport struct {
	Log  logging.Logger
	Data map[string][]string
}

func NewSimpleReport(log logging.Logger) *SimpleReport {
	return &SimpleReport{
		Log:  log,
		Data: make(map[string][]string),
	}
}

func (r *SimpleReport) LogErrAndAppendToNodeReport(nodeName string, errString string) {
	r.AppendToNodeReport(nodeName, errString)
	r.Log.Errorf(errString)
}

func (r *SimpleReport) AppendToNodeReport(nodeName string, errString string) {
	if r.Data[nodeName] == nil {
		r.Data[nodeName] = make([]string, 0)
	}
	r.Data[nodeName] = append(r.Data[nodeName], errString)
}

func (r *SimpleReport) Clear() {
	r.Data = make(map[string][]string)
}

func (r *SimpleReport) Print() {
	fmt.Println("Error SimpleReport:")
	fmt.Println("=============")
	for k, rl := range r.Data {
		fmt.Printf("Key: %s\n", k)
		for i, line := range rl {
			fmt.Printf("  %d: %s\n", i, line)
		}
		fmt.Println()
	}
}
