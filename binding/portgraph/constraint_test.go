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

package portgraph

import (
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		desc    string
		c       Constraint
		v       string
		b, want bool
	}{{
		desc: "Equal; match",
		c:    Equal("match"),
		v:    "match",
		b:    true,
		want: true,
	}, {
		desc: "Equal; match empty string",
		c:    Equal(""),
		v:    "",
		b:    true,
		want: true,
	}, {
		desc: "Equal; no match",
		c:    Equal("match"),
		v:    "not matching",
		b:    true,
		want: false,
	}, {
		desc: "Equal; no match on false",
		c:    Equal("match"),
		v:    "match",
		b:    false,
		want: false,
	}, {
		desc: "NotEqual; match",
		c:    NotEqual("don't match this"),
		v:    "not matching this",
		b:    true,
		want: true,
	}, {
		desc: "NotEqual; match empty string",
		c:    NotEqual("don't match this"),
		v:    "",
		b:    true,
		want: true,
	}, {
		desc: "NotEqual; no match",
		c:    NotEqual("match"),
		v:    "match",
		b:    true,
		want: false,
	}, {
		desc: "NotEqual; no match on false",
		c:    NotEqual("don't match this"),
		v:    "",
		b:    false,
		want: false,
	}, {
		desc: "Regex; match",
		c:    Regex(regexp.MustCompile(".*")),
		v:    "anything",
		b:    true,
		want: true,
	}, {
		desc: "Regex; match empty string",
		c:    Regex(regexp.MustCompile(".*")),
		v:    "",
		b:    true,
		want: true,
	}, {
		desc: "Regex; no match",
		c:    Regex(regexp.MustCompile("[0-9]+")),
		v:    "these are not numbers",
		b:    true,
		want: false,
	}, {
		desc: "Regex; no match on false",
		c:    Regex(regexp.MustCompile(".*")),
		v:    "",
		b:    false,
		want: false,
	}, {
		desc: "NotRegex; match",
		c:    NotRegex(regexp.MustCompile("[a-z]+")),
		v:    "1234",
		b:    true,
		want: true,
	}, {
		desc: "NotRegex; match empty string",
		c:    NotRegex(regexp.MustCompile("[a-z]+")),
		v:    "",
		b:    true,
		want: true,
	}, {
		desc: "NotRegex; no match",
		c:    NotRegex(regexp.MustCompile("[a-z]+")),
		v:    "these are not numbers",
		b:    true,
		want: false,
	}, {
		desc: "NotRegex; no match on false",
		c:    NotRegex(regexp.MustCompile("[a-z]+")),
		v:    "",
		b:    false,
		want: false,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			if m := tc.c.match(tc.v, tc.b); m != tc.want {
				t.Errorf("match() got %t, want %t", m, tc.want)
			}
		})
	}
}

func TestPortConstraint(t *testing.T) {
	dut1port1 := &AbstractPort{Desc: "dut1:port1"}
	dut1port2 := &AbstractPort{Desc: "dut1:port2"}
	node1port1 := &ConcretePort{Desc: "node1:port1", Attrs: map[string]string{"attr": "A"}}
	node1port2 := &ConcretePort{Desc: "node1:port2", Attrs: map[string]string{"attr": "B"}}
	abs2ConPort := map[*AbstractPort]*ConcretePort{
		dut1port1: node1port1,
		dut1port2: node1port2,
	}

	tests := []struct {
		desc  string
		match func(string, *AbstractPort, map[*AbstractPort]*ConcretePort) bool
		pc    PortConstraint
		p     *AbstractPort
		want  bool
	}{{
		desc:  "sameAsPort; match",
		match: SameAsPort(dut1port1).(*sameAsPort).match,
		p:     dut1port1,
		want:  true,
	}, {
		desc:  "sameAsPort; no match",
		match: SameAsPort(dut1port1).(*sameAsPort).match,
		p:     dut1port2,
		want:  false,
	}, {
		desc:  "notSameAsPort; match",
		match: NotSameAsPort(dut1port1).(*notSameAsPort).match,
		p:     dut1port2,
		want:  true,
	}, {
		desc:  "notSameAsPort; no match",
		match: NotSameAsPort(dut1port1).(*notSameAsPort).match,
		p:     dut1port1,
		want:  false,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			if m := tc.match("attr", tc.p, abs2ConPort); m != tc.want {
				t.Errorf("match() got %t, want %t", m, tc.want)
			}
		})
	}
}

func TestAttrPortPair(t *testing.T) {
	absPort1 := &AbstractPort{Desc: "port1"}
	absPort2 := &AbstractPort{Desc: "port2"}
	absPort3 := &AbstractPort{Desc: "port3"}
	conPort1 := &ConcretePort{Desc: "port1", Attrs: map[string]string{"attr": "A"}}
	conPort2 := &ConcretePort{Desc: "port2", Attrs: map[string]string{"attr": "B"}}
	conPort3 := &ConcretePort{Desc: "port3"}
	abs2ConPort := map[*AbstractPort]*ConcretePort{
		absPort1: conPort1,
		absPort2: conPort2,
		absPort3: conPort3,
	}
	tests := []struct {
		desc           string
		p1, p2         *AbstractPort
		k              string
		wantV1, wantV2 string
		wantOk         bool
	}{{
		desc:   "ok",
		p1:     absPort1,
		p2:     absPort2,
		k:      "attr",
		wantV1: "A",
		wantV2: "B",
		wantOk: true,
	}, {
		desc:   "not ok; cp1 missing attr",
		p1:     absPort3,
		p2:     absPort2,
		k:      "attr",
		wantOk: false,
	}, {
		desc:   "not ok; cp2 missing attr",
		p1:     absPort1,
		p2:     absPort3,
		k:      "attr",
		wantOk: false,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			gotV1, gotV2, gotOk := attrsPortPair(tc.k, tc.p1, tc.p2, abs2ConPort)
			if tc.wantOk != gotOk {
				t.Fatalf("attrsPortPair() got %t, want %t", gotOk, tc.wantOk)
			}
			if tc.wantV1 != gotV1 {
				t.Errorf("attrsPortPair() got %s, want %s", gotV1, tc.wantV1)
			}
			if tc.wantV2 != gotV2 {
				t.Errorf("attrsPortPair() got %s, want %s", gotV2, tc.wantV2)
			}
		})
	}
}

func TestAttrNodePair(t *testing.T) {
	absNode1 := &AbstractNode{Desc: "abs1"}
	absNode2 := &AbstractNode{Desc: "abs2"}
	absNode3 := &AbstractNode{Desc: "abs3"}
	conNode1 := &ConcreteNode{Desc: "node1", Attrs: map[string]string{"attr": "A"}}
	conNode2 := &ConcreteNode{Desc: "node2", Attrs: map[string]string{"attr": "B"}}
	conNode3 := &ConcreteNode{Desc: "node3"}
	abs2ConNode := map[*AbstractNode]*ConcreteNode{
		absNode1: conNode1,
		absNode2: conNode2,
		absNode3: conNode3,
	}
	tests := []struct {
		desc           string
		p1, p2         *AbstractNode
		k              string
		wantV1, wantV2 string
		wantOk         bool
	}{{
		desc:   "ok",
		p1:     absNode1,
		p2:     absNode2,
		k:      "attr",
		wantV1: "A",
		wantV2: "B",
		wantOk: true,
	}, {
		desc:   "not ok; cp1 missing attr",
		p1:     absNode3,
		p2:     absNode2,
		k:      "attr",
		wantOk: false,
	}, {
		desc:   "not ok; cp2 missing attr",
		p1:     absNode1,
		p2:     absNode3,
		k:      "attr",
		wantOk: false,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			gotV1, gotV2, gotOk := attrsNodePair(tc.k, tc.p1, tc.p2, abs2ConNode)
			if tc.wantOk != gotOk {
				t.Fatalf("attrsNodePair() got %t, want %t", gotOk, tc.wantOk)
			}
			if tc.wantV1 != gotV1 {
				t.Errorf("attrsNodePair() got %s, want %s", gotV1, tc.wantV1)
			}
			if tc.wantV2 != gotV2 {
				t.Errorf("attrsNodePair() got %s, want %s", gotV2, tc.wantV2)
			}
		})
	}
}
