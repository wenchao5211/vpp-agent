//  Copyright (c) 2020 Pantheon.tech
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp1908

import (
	"context"
	"fmt"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/ifplugin/vppcalls"
	interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
)

// AddRdmaInterface adds new interface with RDMA driver.
func (h *InterfaceVppHandler) AddRdmaInterface(ctx context.Context, ifName string, rdmaLink *interfaces.RDMALink) (uint32, error) {
	return 0, fmt.Errorf("%w in VPP 19.08", vppcalls.ErrRdmaUnsupported)
}

// DeleteRdmaInterface removes interface with RDMA driver.
func (h *InterfaceVppHandler) DeleteRdmaInterface(ctx context.Context, ifName string, ifIdx uint32) error {
	return fmt.Errorf("%w in VPP 19.08", vppcalls.ErrRdmaUnsupported)
}

