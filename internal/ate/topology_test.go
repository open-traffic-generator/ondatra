// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

func TestAddPorts(t *testing.T) {
	const (
		ixiaName            = "ixia1"
		chassisHost         = "192.168.1.1"
		portInstrumentation = "floating"
		firstPort           = "1/1"
		secondPort          = "2/2"
	)
	wantCfg := &ixconfig.Ixnetwork{
		Vport: []*ixconfig.Vport{{
			Name:     ixconfig.String(fmt.Sprintf("%s/1/1", ixiaName)),
			Location: ixconfig.String(fmt.Sprintf("%s;1;1", chassisHost)),
			L1Config: &ixconfig.VportL1Config{
				AresOneFourHundredGigLan: &ixconfig.VportAresOneFourHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusHundredGigLan: &ixconfig.VportNovusHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusTenGigLan: &ixconfig.VportNovusTenGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
			},
			ProtocolStack: &ixconfig.VportProtocolStack{
				Options: &ixconfig.VportProtocolStackOptions{
					McastSolicit: ixconfig.NumberUint32(8),
					RetransTime:  ixconfig.NumberUint32(7000),
				},
			},
		}, {
			Name:     ixconfig.String(fmt.Sprintf("%s/2/2", ixiaName)),
			Location: ixconfig.String(fmt.Sprintf("%s;2;2", chassisHost)),
			L1Config: &ixconfig.VportL1Config{
				AresOneFourHundredGigLan: &ixconfig.VportAresOneFourHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusHundredGigLan: &ixconfig.VportNovusHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusTenGigLan: &ixconfig.VportNovusTenGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
			},
			ProtocolStack: &ixconfig.VportProtocolStack{
				Options: &ixconfig.VportProtocolStackOptions{
					McastSolicit: ixconfig.NumberUint32(8),
					RetransTime:  ixconfig.NumberUint32(7000),
				},
			},
		}},
	}

	c := &IxiaCfgClient{
		name:        ixiaName,
		chassisHost: chassisHost,
		cfg:         &ixconfig.Ixnetwork{},
		ports:       make(map[string]*ixconfig.Vport),
	}

	top := &opb.Topology{
		// Invert order in function call to test alphabetization.
		Interfaces: []*opb.InterfaceConfig{
			{Link: &opb.InterfaceConfig_Port{secondPort}},
			// Test with multiple interfaces on a port.
			{Link: &opb.InterfaceConfig_Port{firstPort}},
			{Link: &opb.InterfaceConfig_Port{firstPort}},
		},
	}
	if err := c.addPorts(top); err != nil {
		t.Fatal(err)
	}
	if diff := jsonCfgDiff(t, wantCfg, c.cfg); diff != "" {
		t.Fatalf("addPorts: Incorrect config generated: diff (-want +got)\n%s", diff)
	}
	if diff := jsonCfgDiff(t, wantCfg.Vport[0], c.ports[firstPort]); diff != "" {
		t.Fatalf("addPorts: Incorrect config for port %s: diff (-want +got)\n%s", firstPort, diff)
	}
	if diff := jsonCfgDiff(t, wantCfg.Vport[1], c.ports[secondPort]); diff != "" {
		t.Fatalf("addPorts: Incorrect config for port %s: diff (-want +got)\n%s", secondPort, diff)
	}
}

func TestAddPortsErrors(t *testing.T) {
	c := &IxiaCfgClient{}

	tests := []struct {
		desc    string
		top     *opb.Topology
		wantErr string
	}{{
		desc: "port in two lags",
		top: &opb.Topology{
			Lags: []*opb.Lag{{
				Ports: []string{"1/1"},
			}, {
				Ports: []string{"1/1"},
			}},
		},
		wantErr: "belongs to two lags",
	}, {
		desc: "interface on port which belongs to lag",
		top: &opb.Topology{
			Lags:       []*opb.Lag{{Ports: []string{"1/1"}}},
			Interfaces: []*opb.InterfaceConfig{{Link: &opb.InterfaceConfig_Port{"1/1"}}},
		},
		wantErr: "already belongs to lag",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := c.addPorts(test.top)
			if err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Fatalf("addPorts: Expected error %q but got %v", test.wantErr, err)
			}
		})
	}
}

func TestAddLAGs(t *testing.T) {
	vport1 := &ixconfig.Vport{
		Name:     ixconfig.String("ixia1/1/1"),
		Location: ixconfig.String("192.168.1.1;1;1"),
	}
	vport2 := &ixconfig.Vport{
		Name:     ixconfig.String("ixia1/2/2"),
		Location: ixconfig.String("192.168.1.1;2;2"),
	}
	c := &IxiaCfgClient{
		name:        "ixia1",
		chassisHost: "192.168.1.1",
		cfg: &ixconfig.Ixnetwork{
			Vport: []*ixconfig.Vport{vport1, vport2},
		},
		ports: map[string]*ixconfig.Vport{
			"1/1": vport1,
			"2/2": vport2,
		},
		lags:     make(map[string]*ixconfig.Lag),
		lagPorts: make(map[*ixconfig.Lag][]*ixconfig.Vport),
	}
	updateXPaths(c.cfg)

	lags := []*opb.Lag{{
		Name:  "testLag",
		Ports: []string{"1/1", "2/2"},
	}}
	c.addLAGs(lags)

	wantCfg := &ixconfig.Ixnetwork{
		Vport: []*ixconfig.Vport{vport1, vport2},
		Lag: []*ixconfig.Lag{{
			Name: ixconfig.String("ixia1/testLag"),
			ProtocolStack: &ixconfig.LagProtocolStack{
				Enabled: ixconfig.MultivalueBool(true),
				Ethernet: []*ixconfig.LagEthernet{{
					Multiplier: ixconfig.NumberFloat64(1),
					Lagportlacp: []*ixconfig.LagLagportlacp{{
						Multiplier: ixconfig.NumberFloat64(1),
					}},
				}},
				Multiplier: ixconfig.NumberFloat64(1),
			},
			Vports: []string{"/vport[1]", "/vport[2]"},
		}},
	}
	if diff := jsonCfgDiff(t, wantCfg, c.cfg); diff != "" {
		t.Fatalf("addLAGs: Incorrect config generated: diff (-want +got)\n%s", diff)
	}
	if diff := jsonCfgDiff(t, wantCfg.Lag[0], c.lags["testLag"]); diff != "" {
		t.Fatalf("addLAGs: Incorrect config for lag: diff (-want +got)\n%s", diff)
	}
}

func TestAddTopologies(t *testing.T) {
	const (
		lagName = "testLag"
		port    = "1/1"
		ifName  = "someIntf"
	)
	vport := &ixconfig.Vport{
		Name:     ixconfig.String(port),
		Location: ixconfig.String("192.168.1.1;1;1"),
		L1Config: &ixconfig.VportL1Config{},
	}
	lag := &ixconfig.Lag{
		Name:   ixconfig.String(lagName),
		Vports: []string{port},
	}
	emptyCfgClient := func() *IxiaCfgClient {
		cfg := &ixconfig.Ixnetwork{
			Vport: []*ixconfig.Vport{vport},
			Lag:   []*ixconfig.Lag{lag},
		}
		updateXPaths(cfg)
		return &IxiaCfgClient{
			cfg:   cfg,
			ports: map[string]*ixconfig.Vport{port: vport},
			lags:  map[string]*ixconfig.Lag{lagName: lag},
			intfs: make(map[string]*intf),
		}
	}
	tests := []struct {
		desc                 string
		ifs                  []*opb.InterfaceConfig
		wantDg, wantNestedDg string
		wantLink             ixconfig.IxiaCfgNode
	}{{
		desc: "No LAG",
		ifs: []*opb.InterfaceConfig{{
			Name:     ifName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		}},
		wantDg:   fmt.Sprintf("Device Group on %s", ifName),
		wantLink: vport,
	}, {
		desc: "With LACP",
		ifs: []*opb.InterfaceConfig{{
			Name:       ifName,
			Link:       &opb.InterfaceConfig_Port{port},
			Ethernet:   &opb.EthernetConfig{Mtu: 1500},
			EnableLacp: true,
		}},
		wantDg:       fmt.Sprintf("LACP DeviceGroup on %s", ifName),
		wantNestedDg: fmt.Sprintf("Device Group on %s", ifName),
		wantLink:     vport,
	}, {
		desc: "With LAG",
		ifs: []*opb.InterfaceConfig{{
			Name:     ifName,
			Link:     &opb.InterfaceConfig_Lag{"testLag"},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		}},
		wantDg:   fmt.Sprintf("Device Group on %s", ifName),
		wantLink: lag,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := emptyCfgClient()
			c.addTopology(test.ifs)

			var gotDg, gotNestedDg string
			dg := c.cfg.Topology[0].DeviceGroup[0]
			gotDg = *(dg.Name)
			if len(dg.DeviceGroup) > 0 {
				gotNestedDg = *(dg.DeviceGroup[0].Name)
			}
			if gotDg != test.wantDg {
				t.Errorf("addTopology: did not find expected device group: got %q, want %q", gotDg, test.wantDg)
			}
			if gotNestedDg != test.wantNestedDg {
				t.Errorf("addTopology: did not find expected nested device group: got %q, want %q", gotNestedDg, test.wantNestedDg)
			}

			if c.intfs[ifName].deviceGroup == nil {
				t.Fatalf("addTopology: no mapped device group for interface %q", ifName)
			}
			if c.intfs[ifName].link != test.wantLink {
				t.Fatalf("addTopology: did not find expected vport for interface %q", ifName)
			}
		})
	}
}