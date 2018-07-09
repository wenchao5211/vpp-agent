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

package nsplugin_test

import (
	"testing"

	"github.com/ligato/cn-infra/logging"
	"github.com/ligato/cn-infra/logging/logrus"
	"github.com/ligato/vpp-agent/plugins/linux/nsplugin"
	"github.com/ligato/vpp-agent/tests/linuxmock"
	. "github.com/onsi/gomega"
)

/* Linux namespace handler init and close */

// Test init function
func TestNsHandlerInit(t *testing.T) {
	plugin, _, _, _, _ := nsHandlerTestSetup(t)
	defer nsHandlerTestTeardown(plugin)

	// Base fields
	Expect(plugin).ToNot(BeNil())
	Expect(plugin.GetMicroserviceByLabel()).ToNot(BeNil())
	Expect(plugin.GetMicroserviceByLabel()).To(HaveLen(0))
	Expect(plugin.GetMicroserviceByID()).ToNot(BeNil())
	Expect(plugin.GetMicroserviceByID()).To(HaveLen(0))

	// todo test microservice tracker
}

/* Namespace handler Test Setup */

func nsHandlerTestSetup(t *testing.T) (*nsplugin.NsHandler, *linuxmock.IfNetlinkHandlerMock, *linuxmock.SystemMock,
	chan *nsplugin.MicroserviceCtx, chan *nsplugin.MicroserviceEvent) {
	RegisterTestingT(t)

	// Loggers
	pluginLog := logging.ForPlugin("linux-ns-handler-log", logrus.NewLogRegistry())
	pluginLog.SetLevel(logging.DebugLevel)
	// Handlers
	ifHandler := linuxmock.NewIfNetlinkHandlerMock()
	sysHandler := linuxmock.NewSystemMock()
	// Channels
	msChan := make(chan *nsplugin.MicroserviceCtx)
	ifNotif := make(chan *nsplugin.MicroserviceEvent)
	// Configurator
	plugin := &nsplugin.NsHandler{}
	err := plugin.Init(pluginLog, ifHandler, sysHandler, msChan, ifNotif)
	Expect(err).To(BeNil())

	return plugin, ifHandler, sysHandler, msChan, ifNotif
}

func nsHandlerTestTeardown(plugin *nsplugin.NsHandler) {
	Expect(plugin.Close()).To(Succeed())
}
