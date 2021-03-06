// Copyright (c) 2020 Pantheon.tech
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

package e2e

import (
	"context"
	"os"
	"testing"

	. "github.com/onsi/gomega"

	"go.ligato.io/vpp-agent/v3/proto/ligato/kvscheduler"
	linux_interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/linux/interfaces"
	linux_l3 "go.ligato.io/vpp-agent/v3/proto/ligato/linux/l3"
	linux_namespace "go.ligato.io/vpp-agent/v3/proto/ligato/linux/namespace"
	vpp_interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
	vpp_l3 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/l3"
)

//
//  +----------------------------------------------------------------------+
//  | VPP                                                                  |
//  | +-------------------------------+  +-------------------------------+ |
//  | | VRF 1                         |  | VRF 2                         | |
//  | |        +-------------------+  |  |        +-------------------+  | |
//  | |        |  192.168.1.1/24   |  |  |        |  192.168.1.1./24  |  | |
//  | |        +---------+---------+  |  |        +----------+--------+  | |
//  | +------------------|------------+  +-------------------|-----------+ |
//  +--------------------|-----------------------------------|-------------+
//  Linux                | (TAP)                             | (TAP)
//    +------------------|------------+  +-------------------|-----------+
//    | VRF_1            |            |  | VRF_2             |           |
//    |        +---------+---------+  |  |         +---------+--------+  |
//    |        |  192.168.1.2/24   |  |  |         |  192.168.1.2/24  |  |
//    |        +-------------------+  |  |         +------------------+  |
//    +-------------------------------+  +-------------------------------+
//
func TestVRFsWithSameSubnets(t *testing.T) {
	if os.Getenv("TRAVIS") != "" {
		// VRFs are seemingly not supported on Ubuntu Xenial, which is used in Travis CI to run the tests.
		// TODO: remove `skip` once we upgrade to Ubuntu Bionic or newer
		t.Skip("skip for travis")
	}

	ctx := Setup(t)
	defer ctx.Teardown()

	const (
		vrf1ID        = 1
		vrf2ID        = 2
		vrf1Label     = "vrf-1"
		vrf2Label     = "vrf-2"
		vrfVppIP      = "192.168.1.1"
		vrfLinuxIP    = "192.168.1.2"
		vrfSubnetMask = "/24"
		tapNameSuffix = "-tap"
		msName        = "microservice1"
	)

	// TAP interfaces
	vrf1VppTap := &vpp_interfaces.Interface{
		Name:        vrf1Label + tapNameSuffix,
		Type:        vpp_interfaces.Interface_TAP,
		Enabled:     true,
		Vrf:         vrf1ID,
		IpAddresses: []string{vrfVppIP + vrfSubnetMask},
		Link: &vpp_interfaces.Interface_Tap{
			Tap: &vpp_interfaces.TapLink{
				Version:        2,
				ToMicroservice: msNamePrefix + msName,
			},
		},
	}
	vrf1LinuxTap := &linux_interfaces.Interface{
		Name:               vrf1Label + tapNameSuffix,
		Type:               linux_interfaces.Interface_TAP_TO_VPP,
		Enabled:            true,
		VrfMasterInterface: vrf1Label,
		IpAddresses:        []string{vrfLinuxIP + vrfSubnetMask},
		Link: &linux_interfaces.Interface_Tap{
			Tap: &linux_interfaces.TapLink{
				VppTapIfName: vrf1Label + tapNameSuffix,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}
	vrf2VppTap := &vpp_interfaces.Interface{
		Name:        vrf2Label + tapNameSuffix,
		Type:        vpp_interfaces.Interface_TAP,
		Enabled:     true,
		Vrf:         vrf2ID,
		IpAddresses: []string{vrfVppIP + vrfSubnetMask},
		Link: &vpp_interfaces.Interface_Tap{
			Tap: &vpp_interfaces.TapLink{
				Version:        2,
				ToMicroservice: msNamePrefix + msName,
			},
		},
	}
	vrf2LinuxTap := &linux_interfaces.Interface{
		Name:               vrf2Label + tapNameSuffix,
		Type:               linux_interfaces.Interface_TAP_TO_VPP,
		VrfMasterInterface: vrf2Label,
		Enabled:            true,
		IpAddresses:        []string{vrfLinuxIP + vrfSubnetMask},
		Link: &linux_interfaces.Interface_Tap{
			Tap: &linux_interfaces.TapLink{
				VppTapIfName: vrf2Label + tapNameSuffix,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}

	// VRFs
	vppVrf1 := &vpp_l3.VrfTable{
		Id:       vrf1ID,
		Label:    vrf1Label,
		Protocol: vpp_l3.VrfTable_IPV4,
	}
	vppVrf2 := &vpp_l3.VrfTable{
		Id:       vrf2ID,
		Label:    vrf2Label,
		Protocol: vpp_l3.VrfTable_IPV4,
	}
	linuxVrf1 := &linux_interfaces.Interface{
		Name:    vrf1Label,
		Type:    linux_interfaces.Interface_VRF_DEVICE,
		Enabled: true,
		Link: &linux_interfaces.Interface_VrfDev{
			VrfDev: &linux_interfaces.VrfDevLink{
				RoutingTable: vrf1ID,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}
	linuxVrf2 := &linux_interfaces.Interface{
		Name:    vrf2Label,
		Type:    linux_interfaces.Interface_VRF_DEVICE,
		Enabled: true,
		Link: &linux_interfaces.Interface_VrfDev{
			VrfDev: &linux_interfaces.VrfDevLink{
				RoutingTable: vrf2ID,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}

	ctx.StartMicroservice(msName)

	// configure everything in one resync
	err := ctx.GenericClient().ResyncConfig(
		vppVrf1, vppVrf2,
		linuxVrf1, linuxVrf2,
		vrf1VppTap, vrf1LinuxTap,
		vrf2VppTap, vrf2LinuxTap,
	)
	Expect(err).ToNot(HaveOccurred())

	Eventually(ctx.GetValueStateClb(vrf1LinuxTap)).Should(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf1VppTap)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf2LinuxTap)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf2VppTap)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(linuxVrf1)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(linuxVrf2)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vppVrf1)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vppVrf2)).To(Equal(kvscheduler.ValueState_CONFIGURED))

	// try to ping in both VRFs
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())

	Expect(ctx.AgentInSync()).To(BeTrue())

	// restart microservice
	ctx.StopMicroservice(msName)
	Eventually(ctx.GetValueStateClb(vrf1LinuxTap)).Should(Equal(kvscheduler.ValueState_PENDING))
	Eventually(ctx.GetValueStateClb(vrf2LinuxTap)).Should(Equal(kvscheduler.ValueState_PENDING))
	Expect(ctx.AgentInSync()).To(BeTrue())

	ctx.StartMicroservice(msName)
	Eventually(ctx.GetValueStateClb(vrf1LinuxTap)).Should(Equal(kvscheduler.ValueState_CONFIGURED))
	Eventually(ctx.GetValueStateClb(vrf2LinuxTap)).Should(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.AgentInSync()).To(BeTrue())

	// re-create Linux VRF1
	err = ctx.GenericClient().ChangeRequest().
		Delete(linuxVrf1).Send(context.Background())
	Expect(err).ToNot(HaveOccurred())
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).ToNot(Succeed())
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())

	err = ctx.GenericClient().ChangeRequest().Update(
		linuxVrf1,
	).Send(context.Background())
	Expect(err).ToNot(HaveOccurred())

	Eventually(ctx.PingFromMsClb(msName, vrfVppIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).Should(Succeed())
	Expect(ctx.PingFromMs(msName, vrfVppIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.AgentInSync()).To(BeTrue())
}

//
//  +--------------------------------------------------------------------------------+
//  | VPP                           inter-VRF RT:                                    |
//  | +-------------------------------+  .2.0/24   +-------------------------------+ |
//  | | VRF 1                         | ---------> | VRF 2                         | |
//  | |        +-------------------+  |            |        +-------------------+  | |
//  | |        |  192.168.1.1/24   |  |  .1.0/24   |        |  192.168.2.1./24  |  | |
//  | |        +---------+---------+  | <--------- |        +----------+--------+  | |
//  | +------------------|------------+            +-------------------|-----------+ |
//  +--------------------|---------------------------------------------|-------------+
//  Linux                | (TAP)                                       | (TAP)
//    +------------------|------------+            +-------------------|-----------+
//    | VRF_1            |            |            | VRF_2             |           |
//    |        +---------+---------+  |            |         +---------+--------+  |
//    |        |  192.168.1.2/24   |  |            |         |  192.168.2.2/24  |  |
//    |        +-------------------+  |            |         +------------------+  |
//    |                         ^     |            |                         ^     |
//    |                         |     |            |                         |     |
//    |    RT: 192.168.2.0/24 --+     |            |    RT: 192.168.1.0/24 --+     |
//    +-------------------------------+            +-------------------------------+
//
func TestVRFRoutes(t *testing.T) {
	if os.Getenv("TRAVIS") != "" {
		// VRFs are seemingly not supported on Ubuntu Xenial, which is used in Travis CI to run the tests.
		// TODO: remove `skip` once we upgrade to Ubuntu Bionic or newer
		t.Skip("skip for travis")
	}

	ctx := Setup(t)
	defer ctx.Teardown()

	const (
		vrf1ID        = 1
		vrf2ID        = 2
		vrf1Label     = "vrf-1"
		vrf2Label     = "vrf-2"
		vrfSubnetMask = "/24"
		vrf1Subnet    = "192.168.1.0" + vrfSubnetMask
		vrf1VppIP     = "192.168.1.1"
		vrf1LinuxIP   = "192.168.1.2"
		vrf2Subnet    = "192.168.2.0" + vrfSubnetMask
		vrf2VppIP     = "192.168.2.1"
		vrf2LinuxIP   = "192.168.2.2"
		tapNameSuffix = "-tap"
		msName        = "microservice1"
	)

	// TAP interfaces
	vrf1VppTap := &vpp_interfaces.Interface{
		Name:        vrf1Label + tapNameSuffix,
		Type:        vpp_interfaces.Interface_TAP,
		Enabled:     true,
		Vrf:         vrf1ID,
		IpAddresses: []string{vrf1VppIP + vrfSubnetMask},
		Link: &vpp_interfaces.Interface_Tap{
			Tap: &vpp_interfaces.TapLink{
				Version:        2,
				ToMicroservice: msNamePrefix + msName,
			},
		},
	}
	vrf1LinuxTap := &linux_interfaces.Interface{
		Name:               vrf1Label + tapNameSuffix,
		Type:               linux_interfaces.Interface_TAP_TO_VPP,
		Enabled:            true,
		VrfMasterInterface: vrf1Label,
		IpAddresses:        []string{vrf1LinuxIP + vrfSubnetMask},
		Link: &linux_interfaces.Interface_Tap{
			Tap: &linux_interfaces.TapLink{
				VppTapIfName: vrf1Label + tapNameSuffix,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}
	vrf2VppTap := &vpp_interfaces.Interface{
		Name:        vrf2Label + tapNameSuffix,
		Type:        vpp_interfaces.Interface_TAP,
		Enabled:     true,
		Vrf:         vrf2ID,
		IpAddresses: []string{vrf2VppIP + vrfSubnetMask},
		Link: &vpp_interfaces.Interface_Tap{
			Tap: &vpp_interfaces.TapLink{
				Version:        2,
				ToMicroservice: msNamePrefix + msName,
			},
		},
	}
	vrf2LinuxTap := &linux_interfaces.Interface{
		Name:               vrf2Label + tapNameSuffix,
		Type:               linux_interfaces.Interface_TAP_TO_VPP,
		VrfMasterInterface: vrf2Label,
		Enabled:            true,
		IpAddresses:        []string{vrf2LinuxIP + vrfSubnetMask},
		Link: &linux_interfaces.Interface_Tap{
			Tap: &linux_interfaces.TapLink{
				VppTapIfName: vrf2Label + tapNameSuffix,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}

	// VRFs
	vppVrf1 := &vpp_l3.VrfTable{
		Id:       vrf1ID,
		Label:    vrf1Label,
		Protocol: vpp_l3.VrfTable_IPV4,
	}
	vppVrf2 := &vpp_l3.VrfTable{
		Id:       vrf2ID,
		Label:    vrf2Label,
		Protocol: vpp_l3.VrfTable_IPV4,
	}
	linuxVrf1 := &linux_interfaces.Interface{
		Name:    vrf1Label,
		Type:    linux_interfaces.Interface_VRF_DEVICE,
		Enabled: true,
		Link: &linux_interfaces.Interface_VrfDev{
			VrfDev: &linux_interfaces.VrfDevLink{
				RoutingTable: vrf1ID,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}
	linuxVrf2 := &linux_interfaces.Interface{
		Name:    vrf2Label,
		Type:    linux_interfaces.Interface_VRF_DEVICE,
		Enabled: true,
		Link: &linux_interfaces.Interface_VrfDev{
			VrfDev: &linux_interfaces.VrfDevLink{
				RoutingTable: vrf2ID,
			},
		},
		Namespace: &linux_namespace.NetNamespace{
			Type:      linux_namespace.NetNamespace_MICROSERVICE,
			Reference: msNamePrefix + msName,
		},
	}

	// Routes
	vrf1VppRoute := &vpp_l3.Route{
		Type:        vpp_l3.Route_INTER_VRF,
		VrfId:       vrf1ID,
		DstNetwork:  vrf2Subnet,
		NextHopAddr: vrf2LinuxIP,
		ViaVrfId:    vrf2ID,
	}
	vrf2VppRoute := &vpp_l3.Route{
		Type:        vpp_l3.Route_INTER_VRF,
		VrfId:       vrf2ID,
		DstNetwork:  vrf1Subnet,
		NextHopAddr: vrf1LinuxIP,
		ViaVrfId:    vrf1ID,
	}
	vrf1LinuxRoute := &linux_l3.Route{
		OutgoingInterface: vrf1Label + tapNameSuffix,
		Scope:             linux_l3.Route_GLOBAL,
		DstNetwork:        vrf2Subnet,
		GwAddr:            vrf1VppIP,
	}
	vrf2LinuxRoute := &linux_l3.Route{
		OutgoingInterface: vrf2Label + tapNameSuffix,
		Scope:             linux_l3.Route_GLOBAL,
		DstNetwork:        vrf1Subnet,
		GwAddr:            vrf2VppIP,
	}

	ctx.StartMicroservice(msName)

	// configure everything in one resync
	err := ctx.GenericClient().ResyncConfig(
		vppVrf1, vppVrf2,
		linuxVrf1, linuxVrf2,
		vrf1VppTap, vrf1LinuxTap,
		vrf2VppTap, vrf2LinuxTap,
		vrf1VppRoute, vrf2VppRoute,
		vrf1LinuxRoute, vrf2LinuxRoute,
	)
	Expect(err).ToNot(HaveOccurred())

	Eventually(ctx.GetValueStateClb(vrf1LinuxTap)).Should(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf1VppTap)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf2LinuxTap)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf2VppTap)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf1VppRoute)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf2VppRoute)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf1LinuxRoute)).To(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.GetValueState(vrf2LinuxRoute)).To(Equal(kvscheduler.ValueState_CONFIGURED))

	// try to ping across VRFs
	Expect(ctx.PingFromMs(msName, vrf2LinuxIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.PingFromMs(msName, vrf1LinuxIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.AgentInSync()).To(BeTrue())

	// restart microservice
	ctx.StopMicroservice(msName)
	Eventually(ctx.GetValueStateClb(vrf1LinuxTap)).Should(Equal(kvscheduler.ValueState_PENDING))
	Eventually(ctx.GetValueStateClb(vrf2LinuxTap)).Should(Equal(kvscheduler.ValueState_PENDING))
	Expect(ctx.AgentInSync()).To(BeTrue())

	ctx.StartMicroservice(msName)
	Eventually(ctx.GetValueStateClb(vrf1LinuxTap)).Should(Equal(kvscheduler.ValueState_CONFIGURED))
	Eventually(ctx.GetValueStateClb(vrf2LinuxTap)).Should(Equal(kvscheduler.ValueState_CONFIGURED))
	Expect(ctx.PingFromMs(msName, vrf2LinuxIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.PingFromMs(msName, vrf1LinuxIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.AgentInSync()).To(BeTrue())

	// re-create Linux VRF1
	err = ctx.GenericClient().ChangeRequest().
		Delete(linuxVrf1).Send(context.Background())
	Expect(err).ToNot(HaveOccurred())
	Expect(ctx.PingFromMs(msName, vrf2LinuxIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).ToNot(Succeed())
	Expect(ctx.PingFromMs(msName, vrf1LinuxIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).ToNot(Succeed())

	err = ctx.GenericClient().ChangeRequest().Update(
		linuxVrf1,
	).Send(context.Background())
	Expect(err).ToNot(HaveOccurred())
	Eventually(ctx.PingFromMsClb(msName, vrf2LinuxIP, pingWithOutInterface(vrf1Label+tapNameSuffix))).Should(Succeed())
	Expect(ctx.PingFromMs(msName, vrf1LinuxIP, pingWithOutInterface(vrf2Label+tapNameSuffix))).To(Succeed())
	Expect(ctx.AgentInSync()).To(BeTrue())
}
