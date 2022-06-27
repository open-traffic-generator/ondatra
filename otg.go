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

package ondatra

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/debugger"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/internal/otg"
	"github.com/openconfig/ondatra/telemetry/otg/device"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// OTG provides the Open Traffic Generator API to an ATE.
type OTG struct {
	ate binding.ATE
}

func (o *OTG) String() string {
	return fmt.Sprintf("{ate: %s}", o.ate)
}

// NewConfig creates a new OTG config.
func (o *OTG) NewConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	debugger.ActionStarted(t, "Creating new config for %s", o.ate)
	cfg, err := otg.NewConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("NewConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

// PushConfig pushes config to the ATE.
func (o *OTG) PushConfig(t testing.TB, cfg gosnappi.Config) {
	t.Helper()
	debugger.ActionStarted(t, "Pushing config to %s", o.ate)
	warns, err := otg.PushConfig(context.Background(), o.ate, cfg)
	if err != nil {
		t.Fatalf("PushConfig(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("PushConfig(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// FetchConfig fetches config from the ATE.
func (o *OTG) FetchConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	debugger.ActionStarted(t, "Fetching config from %s", o.ate)
	cfg, err := otg.FetchConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("FetchConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

// StartProtocols starts protocols on the ATE.
func (o *OTG) StartProtocols(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Starting protocols on %s", o.ate)
	warns, err := otg.StartProtocols(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("StartProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StartProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StopProtocols stops protocols on the ATE.
func (o *OTG) StopProtocols(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Stopping protocols on %s", o.ate)
	warns, err := otg.StopProtocols(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StartTraffic starts traffic on the ATE.
func (o *OTG) StartTraffic(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Starting traffic on %s", o.ate)
	warns, err := otg.StartTraffic(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("StartTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StartTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StopTraffic stops traffic on the ATE.
func (o *OTG) StopTraffic(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Stopping traffic on %s", o.ate)
	warns, err := otg.StopTraffic(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("StopTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// WithdrawRoutes withdraw routes on the ATE.
func (o *OTG) WithdrawRoutes(t testing.TB, routes []string) {
	t.Helper()
	debugger.ActionStarted(t, "withdrawing routes for %v", o.ate)
	warns, err := otg.WithdrawRoutes(context.Background(), o.ate, routes)
	if err != nil {
		t.Fatalf("WithdrawRoutes(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("WithdrawRoutes(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// AdvertiseRoutes advertise routes on the ATE.
func (o *OTG) AdvertiseRoutes(t testing.TB, routes []string) {
	t.Helper()
	debugger.ActionStarted(t, "advertising routes for %v", o.ate)
	warns, err := otg.AdvertiseRoutes(context.Background(), o.ate, routes)
	if err != nil {
		t.Fatalf("AdvertiseRoutes(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("AdvertiseRoutes(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// UpLacpMember up lacp member ports on the ATE.
func (o *OTG) UpLacpMember(t testing.TB, lagMemberPorts []string) {
	t.Helper()
	debugger.ActionStarted(t, "UpLacpMember for %v", o.ate)
	warns, err := otg.UpLacpMember(context.Background(), o.ate, lagMemberPorts)
	if err != nil {
		t.Fatalf("UpLacpMember(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("UpLacpMember(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// DownLacpMember down lacp member ports on the ATE.
func (o *OTG) DownLacpMember(t testing.TB, lagMemberPorts []string) {
	t.Helper()
	debugger.ActionStarted(t, "DownLacpMember for %v", o.ate)
	warns, err := otg.DownLacpMember(context.Background(), o.ate, lagMemberPorts)
	if err != nil {
		t.Fatalf("DownLacpMember(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("DownLacpMember(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// Telemetry returns a telemetry path root for the device.
func (o *OTG) Telemetry() *device.DevicePath {
	root := device.DeviceRoot(o.ate.Name())
	root.PutCustomData(genutil.DefaultClientKey, func(ctx context.Context) (gpb.GNMIClient, error) {
		return otg.FetchGNMI(ctx, o.ate)
	})
	return root
}
