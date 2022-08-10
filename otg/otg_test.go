// Copyright 2022 Google LLC
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

package otg

import (
	"golang.org/x/net/context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
)

var (
	fakeSnappi = new(fakeGosnappi)
	fakeATE    = &fakebind.ATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Ports: map[string]*binding.Port{"port1": {}},
		}},
		DialOTGFn: func(context.Context) (gosnappi.GosnappiApi, error) {
			return fakeSnappi, nil
		},
	}
	otgAPI = &OTG{ate: fakeATE}
)

func TestNewConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	fakeSnappi.cfg = wantCfg
	gotCfg := otgAPI.NewConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("NewConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestFetchConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	fakeSnappi.cfg = wantCfg
	gotCfg := otgAPI.FetchConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("FetchConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestPushConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	wantCfg.Ports().Add().SetName("port1")
	otgAPI.PushConfig(t, wantCfg)
	if gotCfg := fakeSnappi.cfg; wantCfg != gotCfg {
		t.Errorf("PushConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestPushConfigBadPortName(t *testing.T) {
	cfg := gosnappi.NewConfig()
	cfg.Ports().Add().SetName("port3")
	gotErr := testt.ExpectFatal(t, func(t testing.TB) {
		otgAPI.PushConfig(t, cfg)
	})
	if wantErr := "not one of the ATE ports"; !strings.Contains(gotErr, wantErr) {
		t.Errorf("PushConfig got unexpected err %s, want err containing %s", gotErr, wantErr)
	}
}

func TestStartProtocols(t *testing.T) {
	fakeSnappi.protocol = nil
	otgAPI.StartProtocols(t)
	if got, want := fakeSnappi.protocol.state, gosnappi.ProtocolStateState.START; got != want {
		t.Errorf("StartProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStopProtocols(t *testing.T) {
	fakeSnappi.protocol = nil
	otgAPI.StopProtocols(t)
	if got, want := fakeSnappi.protocol.state, gosnappi.ProtocolStateState.STOP; got != want {
		t.Errorf("StopProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStartTraffic(t *testing.T) {
	fakeSnappi.transmit = nil
	otgAPI.StartTraffic(t)
	if got, want := fakeSnappi.transmit.state, gosnappi.TransmitStateState.START; got != want {
		t.Errorf("StartTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestStopTraffic(t *testing.T) {
	fakeSnappi.transmit = nil
	otgAPI.StopTraffic(t)
	if got, want := fakeSnappi.transmit.state, gosnappi.TransmitStateState.STOP; got != want {
		t.Errorf("StopTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestAdvertiseRoutes(t *testing.T) {
	fakeSnappi.route = nil
	wantNames := []string{"peer1", "peer2"}
	otgAPI.AdvertiseRoutes(t, wantNames)
	got := fakeSnappi.route
	if wantState := gosnappi.RouteStateState.ADVERTISE; got.state != wantState {
		t.Errorf("AdvertiseRoutes got unexpected route state %v, want %v", got.state, wantState)
	}
	if !cmp.Equal(got.names, wantNames) {
		t.Errorf("AdvertiseRoutes got unexpected route names %v, want %v", got.names, wantNames)
	}
}

func TestWithdrawRoutes(t *testing.T) {
	fakeSnappi.route = nil
	wantNames := []string{"peer1", "peer2"}
	otgAPI.WithdrawRoutes(t, wantNames)
	got := fakeSnappi.route
	if wantState := gosnappi.RouteStateState.WITHDRAW; got.state != wantState {
		t.Errorf("WithdrawRoutes got unexpected route state %v, want %v", got.state, wantState)
	}
	if !cmp.Equal(got.names, wantNames) {
		t.Errorf("WithdrawRoutes got unexpected route names %v, want %v", got.names, wantNames)
	}
}

func TestEnableLACPMembers(t *testing.T) {
	wantPorts := []string{"port1", "port2"}
	otgAPI.EnableLACPMembers(t, wantPorts)
	gotLACP := fakeSnappi.device.lacp
	if wantState := gosnappi.LacpMemberStateState.UP; gotLACP.State() != wantState {
		t.Errorf("EnableLACPMembers got unexpected LACP member state %v, want %v", gotLACP.State(), wantState)
	}
	if !cmp.Equal(gotLACP.LagMemberPortNames(), wantPorts) {
		t.Errorf("EnableLACPMembers got unexpected LACP member ports %v, want %v", gotLACP.LagMemberPortNames(), wantPorts)
	}
}

func TestDisableLACPMembers(t *testing.T) {
	wantPorts := []string{"port1", "port2"}
	otgAPI.DisableLACPMembers(t, wantPorts)
	gotLACP := fakeSnappi.device.lacp
	if wantState := gosnappi.LacpMemberStateState.DOWN; gotLACP.State() != wantState {
		t.Errorf("DisableLACPMembers got unexpected LACP member state %v, want %v", gotLACP.State(), wantState)
	}
	if !cmp.Equal(gotLACP.LagMemberPortNames(), wantPorts) {
		t.Errorf("DisableLACPMembers got unexpected LACP member ports %v, want %v", gotLACP.LagMemberPortNames(), wantPorts)
	}
}

type fakeGosnappi struct {
	gosnappi.GosnappiApi
	cfg      gosnappi.Config
	protocol *fakeProtocolState
	transmit *fakeTransmitState
	route    *fakeRouteState
	device   *fakeDeviceState
}

func (fg *fakeGosnappi) NewConfig() gosnappi.Config {
	return fg.cfg
}

func (fg *fakeGosnappi) GetConfig() (gosnappi.Config, error) {
	return fg.cfg, nil
}

func (fg *fakeGosnappi) SetConfig(cfg gosnappi.Config) (gosnappi.ResponseWarning, error) {
	fg.cfg = cfg
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewProtocolState() gosnappi.ProtocolState {
	return new(fakeProtocolState)
}

func (fg *fakeGosnappi) SetProtocolState(state gosnappi.ProtocolState) (gosnappi.ResponseWarning, error) {
	fg.protocol = state.(*fakeProtocolState)
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewTransmitState() gosnappi.TransmitState {
	return new(fakeTransmitState)
}

func (fg *fakeGosnappi) SetTransmitState(state gosnappi.TransmitState) (gosnappi.ResponseWarning, error) {
	fg.transmit = state.(*fakeTransmitState)
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewRouteState() gosnappi.RouteState {
	return new(fakeRouteState)
}

func (fg *fakeGosnappi) SetRouteState(state gosnappi.RouteState) (gosnappi.ResponseWarning, error) {
	fg.route = state.(*fakeRouteState)
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewDeviceState() gosnappi.DeviceState {
	return new(fakeDeviceState)
}

func (fg *fakeGosnappi) SetDeviceState(state gosnappi.DeviceState) (gosnappi.ResponseWarning, error) {
	fg.device = state.(*fakeDeviceState)
	return gosnappi.NewResponseWarning(), nil
}

type fakeProtocolState struct {
	gosnappi.ProtocolState
	state gosnappi.ProtocolStateStateEnum
}

func (fp *fakeProtocolState) SetState(state gosnappi.ProtocolStateStateEnum) gosnappi.ProtocolState {
	fp.state = state
	return fp
}

type fakeTransmitState struct {
	gosnappi.TransmitState
	state gosnappi.TransmitStateStateEnum
}

func (ft *fakeTransmitState) SetState(state gosnappi.TransmitStateStateEnum) gosnappi.TransmitState {
	ft.state = state
	return ft
}

type fakeRouteState struct {
	gosnappi.RouteState
	state gosnappi.RouteStateStateEnum
	names []string
}

func (fp *fakeRouteState) SetState(state gosnappi.RouteStateStateEnum) gosnappi.RouteState {
	fp.state = state
	return fp
}

func (fp *fakeRouteState) SetNames(names []string) gosnappi.RouteState {
	fp.names = names
	return fp
}

type fakeDeviceState struct {
	gosnappi.DeviceState
	lacp gosnappi.LacpMemberState
}

func (fd *fakeDeviceState) SetLacpMemberState(lacp gosnappi.LacpMemberState) gosnappi.DeviceState {
	fd.lacp = lacp
	return fd
}
