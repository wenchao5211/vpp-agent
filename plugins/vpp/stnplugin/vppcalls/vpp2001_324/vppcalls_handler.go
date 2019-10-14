//  Copyright (c) 2019 Cisco and/or its affiliates.
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

package vpp2001_324

import (
	govppapi "git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/logging"

	vpp_stn "github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp2001_324/stn"
	"github.com/ligato/vpp-agent/plugins/vpp/ifplugin/ifaceidx"
	"github.com/ligato/vpp-agent/plugins/vpp/stnplugin/vppcalls"
)

func init() {
	var msgs []govppapi.Message
	msgs = append(msgs, vpp_stn.AllMessages()...)

	vppcalls.Versions["vpp2001_324"] = vppcalls.HandlerVersion{
		Msgs: msgs,
		New: func(
			ch govppapi.Channel, ifIdx ifaceidx.IfaceMetadataIndex, log logging.Logger,
		) vppcalls.StnVppAPI {
			return NewStnVppHandler(ch, ifIdx, log)
		},
	}
}

// StnVppHandler is accessor for STN-related vppcalls methods
type StnVppHandler struct {
	callsChannel govppapi.Channel
	ifIndexes    ifaceidx.IfaceMetadataIndex
	log          logging.Logger
}

// NewStnVppHandler creates new instance of STN vppcalls handler
func NewStnVppHandler(
	callsChan govppapi.Channel, ifIndexes ifaceidx.IfaceMetadataIndex, log logging.Logger,
) *StnVppHandler {
	return &StnVppHandler{
		callsChannel: callsChan,
		ifIndexes:    ifIndexes,
		log:          log,
	}
}