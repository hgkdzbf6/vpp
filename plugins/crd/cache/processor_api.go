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

package cache

import "github.com/contiv/vpp/plugins/crd/cache/telemetrymodel"

//Processor defines the API of the Node Processor which gos out and
//collects important information for the node and updates the cache with it.
//It also validates that all of the information for the nodes is correct.
type Processor interface {
	CollectNodeInfo(node *telemetrymodel.Node)
	ValidateNodeInfo()
}