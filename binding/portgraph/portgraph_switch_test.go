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
    "fmt"
    "regexp"
    "strings"
    "testing"

    "golang.org/x/net/context"
)

// Initialize ConcretePorts and ConcreteNodes for the super graph.
// These are the ConcreteNodes and ConcretePorts that should be matched to in tests.
var (
    dutnode1port1  = &ConcretePort{Desc: "dutnode1:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode1port2  = &ConcretePort{Desc: "dutnode1:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode2port1  = &ConcretePort{Desc: "atenode2:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode2port2  = &ConcretePort{Desc: "atenode2:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode3port1  = &ConcretePort{Desc: "dutnode3:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode3port2  = &ConcretePort{Desc: "dutnode3:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode4port1  = &ConcretePort{Desc: "atenode4:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode4port2  = &ConcretePort{Desc: "atenode4:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode5port1  = &ConcretePort{Desc: "dutnode5:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode6port1  = &ConcretePort{Desc: "atenode6:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode7port1  = &ConcretePort{Desc: "dutnode7:port1"}
    atenode8port1  = &ConcretePort{Desc: "atenode8:port1"}
    dutnode9port1  = &ConcretePort{Desc: "dutnode9:port1", Attrs: map[string]string{"attr": "FOO"}}
    dutnode9port2  = &ConcretePort{Desc: "dutnode9:port2", Attrs: map[string]string{"attr": "BAR", "attr2": "1"}}
    dutnode9port3  = &ConcretePort{Desc: "dutnode9:port3", Attrs: map[string]string{"attr": "BAR", "attr2": "2"}}
    atenode10port1 = &ConcretePort{Desc: "atenode10:port1", Attrs: map[string]string{"attr": "FOO"}}
    atenode10port2 = &ConcretePort{Desc: "atenode10:port2", Attrs: map[string]string{"attr": "BAR"}}
    dutnode11port1 = &ConcretePort{Desc: "dutnode11:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode11port2 = &ConcretePort{Desc: "dutnode11:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode11port3 = &ConcretePort{Desc: "dutnode11:port3", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode11port4 = &ConcretePort{Desc: "dutnode11:port4", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode12port1 = &ConcretePort{Desc: "atenode12:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode12port2 = &ConcretePort{Desc: "atenode12:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode12port3 = &ConcretePort{Desc: "atenode12:port3", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode12port4 = &ConcretePort{Desc: "atenode12:port4", Attrs: map[string]string{"speed": "S_400GB"}}

    dutnode13port1 = &ConcretePort{Desc: "dutnode13:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode13port2 = &ConcretePort{Desc: "dutnode13:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode13port3 = &ConcretePort{Desc: "dutnode13:port3", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode13port4 = &ConcretePort{Desc: "dutnode13:port4", Attrs: map[string]string{"speed": "S_400GB"}}
    dutnode13port5 = &ConcretePort{Desc: "dutnode13:port5", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode14port1 = &ConcretePort{Desc: "atenode14:port1", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode14port2 = &ConcretePort{Desc: "atenode14:port2", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode14port3 = &ConcretePort{Desc: "atenode14:port3", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode14port4 = &ConcretePort{Desc: "atenode14:port4", Attrs: map[string]string{"speed": "S_400GB"}}
    atenode14port5 = &ConcretePort{Desc: "atenode14:port5", Attrs: map[string]string{"speed": "S_400GB"}}

    d1p1 = &ConcretePort{Desc: "d1:port1", Attrs: map[string]string{"speed": "S_100GB"}}
    d1p2 = &ConcretePort{Desc: "d1:port2", Attrs: map[string]string{"speed": "S_100GB"}}
    a1p1 = &ConcretePort{Desc: "a1:port1", Attrs: map[string]string{"speed": "S_100GB"}}
    a1p2 = &ConcretePort{Desc: "a1:port2", Attrs: map[string]string{"speed": "S_100GB"}}

    switchnode9port1  = &ConcretePort{Desc: "switchnode9:port1"}
    switchnode9port2  = &ConcretePort{Desc: "switchnode9:port2"}
    switchnode9port3  = &ConcretePort{Desc: "switchnode9:port3"}
    switchnode9port4  = &ConcretePort{Desc: "switchnode9:port4"}
    switchnode9port5  = &ConcretePort{Desc: "switchnode9:port5"}
    switchnode9port6  = &ConcretePort{Desc: "switchnode9:port6"}
    switchnode9port7  = &ConcretePort{Desc: "switchnode9:port7"}
    switchnode9port8  = &ConcretePort{Desc: "switchnode9:port8"}
    switchnode10port1 = &ConcretePort{Desc: "switchnode10:port1"}
    switchnode10port2 = &ConcretePort{Desc: "switchnode10:port2"}
    switchnode10port3 = &ConcretePort{Desc: "switchnode10:port3"}
    switchnode10port4 = &ConcretePort{Desc: "switchnode10:port4"}
    switchnode10port5 = &ConcretePort{Desc: "switchnode10:port5"}
    switchnode10port6 = &ConcretePort{Desc: "switchnode10:port6"}
    switchnode10port7 = &ConcretePort{Desc: "switchnode10:port7"}
    switchnode10port8 = &ConcretePort{Desc: "switchnode10:port8"}

    switchnode11port1 = &ConcretePort{Desc: "switchnode11:port1"}
    switchnode11port2 = &ConcretePort{Desc: "switchnode11:port2"}
    switchnode11port3 = &ConcretePort{Desc: "switchnode11:port3"}
    switchnode11port4 = &ConcretePort{Desc: "switchnode11:port4"}
    switchnode11port5 = &ConcretePort{Desc: "switchnode11:port5"}
    switchnode12port1 = &ConcretePort{Desc: "switchnode12:port1"}
    switchnode12port2 = &ConcretePort{Desc: "switchnode12:port2"}
    switchnode12port3 = &ConcretePort{Desc: "switchnode12:port3"}
    switchnode12port4 = &ConcretePort{Desc: "switchnode12:port4"}
    switchnode12port5 = &ConcretePort{Desc: "switchnode12:port5"}

    switchnode13port1 = &ConcretePort{Desc: "switchnode13:port1"}
    switchnode13port2 = &ConcretePort{Desc: "switchnode13:port2"}
    switchnode13port3 = &ConcretePort{Desc: "switchnode13:port3"}
    switchnode13port4 = &ConcretePort{Desc: "switchnode13:port4"}
    switchnode13port5 = &ConcretePort{Desc: "switchnode13:port5"}
    switchnode13port6 = &ConcretePort{Desc: "switchnode13:port6"}

    switchnode14port1 = &ConcretePort{Desc: "switchnode14:port1"}
    switchnode14port2 = &ConcretePort{Desc: "switchnode14:port2"}
    switchnode15port1 = &ConcretePort{Desc: "switchnode15:port1"}
    switchnode15port2 = &ConcretePort{Desc: "switchnode15:port2"}

    switchnode16port1 = &ConcretePort{Desc: "switchnode16:port1"}
    switchnode16port2 = &ConcretePort{Desc: "switchnode16:port2"}
    switchnode16port3 = &ConcretePort{Desc: "switchnode16:port3"}
    switchnode16port4 = &ConcretePort{Desc: "switchnode16:port4"}
    switchnode16port5 = &ConcretePort{Desc: "switchnode16:port5"}
    switchnode16port6 = &ConcretePort{Desc: "switchnode16:port6"}

    dutnode1  = &ConcreteNode{Desc: "dutnode1", Ports: []*ConcretePort{dutnode1port1, dutnode1port2}, Attrs: map[string]string{"vendor": "CISCO"}}
    atenode2  = &ConcreteNode{Desc: "atenode2", Ports: []*ConcretePort{atenode2port1, atenode2port2}, Attrs: map[string]string{"vendor": "TGEN"}}
    dutnode3  = &ConcreteNode{Desc: "dutnode3", Ports: []*ConcretePort{dutnode3port1, dutnode3port2}, Attrs: map[string]string{"vendor": "CISCO"}}
    atenode4  = &ConcreteNode{Desc: "atenode4", Ports: []*ConcretePort{atenode4port1, atenode4port2}, Attrs: map[string]string{"vendor": "TGEN"}}
    dutnode5  = &ConcreteNode{Desc: "dutnode5", Ports: []*ConcretePort{dutnode5port1}, Attrs: map[string]string{"vendor": "CISCO"}}
    atenode6  = &ConcreteNode{Desc: "atenode6", Ports: []*ConcretePort{atenode6port1}, Attrs: map[string]string{"vendor": "TGEN"}}
    dutnode7  = &ConcreteNode{Desc: "dutnode7", Ports: []*ConcretePort{dutnode7port1}, Attrs: map[string]string{"vendor": "CISCO"}}
    atenode8  = &ConcreteNode{Desc: "atenode8", Ports: []*ConcretePort{atenode8port1}, Attrs: map[string]string{"vendor": "TGEN"}}
    dutnode9  = &ConcreteNode{Desc: "dutnode9", Ports: []*ConcretePort{dutnode9port1, dutnode9port2, dutnode9port3}, Attrs: map[string]string{"attr": "multi1"}}
    atenode10 = &ConcreteNode{Desc: "atenode10", Ports: []*ConcretePort{atenode10port1, atenode10port2}, Attrs: map[string]string{"attr": "multi2"}}
    dutnode11 = &ConcreteNode{Desc: "dutnode11", Ports: []*ConcretePort{dutnode11port1, dutnode11port2, dutnode11port3, dutnode11port4}, Attrs: map[string]string{"vendor": "CISCO"}}
    atenode12 = &ConcreteNode{Desc: "atenode12", Ports: []*ConcretePort{atenode12port1, atenode12port2, atenode12port3, atenode12port4}, Attrs: map[string]string{"vendor": "TGEN"}}
    dutnode13 = &ConcreteNode{Desc: "dutnode13", Ports: []*ConcretePort{dutnode13port1, dutnode13port2, dutnode13port3, dutnode13port4, dutnode13port5}, Attrs: map[string]string{"vendor": "CISCO"}}
    atenode14 = &ConcreteNode{Desc: "atenode14", Ports: []*ConcretePort{atenode14port1, atenode14port2, atenode14port3, atenode14port4, atenode14port5}, Attrs: map[string]string{"vendor": "TGEN"}}

    d1 = &ConcreteNode{Desc: "d1", Ports: []*ConcretePort{d1p1, d1p2}, Attrs: map[string]string{"vendor": "CISCO"}}
    a1 = &ConcreteNode{Desc: "a1", Ports: []*ConcretePort{a1p1, a1p2}, Attrs: map[string]string{"vendor": "TGEN"}}

    switchnode9  = &ConcreteNode{Desc: "switchnode9", Ports: []*ConcretePort{switchnode9port1, switchnode9port2, switchnode9port3, switchnode9port4, switchnode9port5, switchnode9port6, switchnode9port7, switchnode9port8}, Attrs: map[string]string{"role": "Switch", "name": "sw1"}}
    switchnode10 = &ConcreteNode{Desc: "switchnode10", Ports: []*ConcretePort{switchnode10port1, switchnode10port2, switchnode10port3, switchnode10port4, switchnode10port5, switchnode10port6, switchnode10port7, switchnode10port8}, Attrs: map[string]string{"role": "Switch", "name": "sw2"}}
    switchnode11 = &ConcreteNode{Desc: "switchnode11", Ports: []*ConcretePort{switchnode11port1, switchnode11port2, switchnode11port3, switchnode11port4, switchnode11port5}, Attrs: map[string]string{"role": "Switch", "name": "sw3"}}
    switchnode12 = &ConcreteNode{Desc: "switchnode12", Ports: []*ConcretePort{switchnode12port1, switchnode12port2, switchnode12port3, switchnode12port4, switchnode12port5}, Attrs: map[string]string{"role": "Switch", "name": "sw4"}}

    switchnode13 = &ConcreteNode{Desc: "switchnode13", Ports: []*ConcretePort{switchnode13port1, switchnode13port2, switchnode13port3, switchnode13port4, switchnode13port5, switchnode13port6}, Attrs: map[string]string{"role": "Switch", "name": "sw5"}}
    switchnode14 = &ConcreteNode{Desc: "switchnode14", Ports: []*ConcretePort{switchnode14port1, switchnode14port2}, Attrs: map[string]string{"role": "Switch", "name": "sw6"}}
    switchnode15 = &ConcreteNode{Desc: "switchnode15", Ports: []*ConcretePort{switchnode15port1, switchnode15port2}, Attrs: map[string]string{"role": "Switch", "name": "sw7"}}
    switchnode16 = &ConcreteNode{Desc: "switchnode16", Ports: []*ConcretePort{switchnode16port1, switchnode16port2, switchnode16port3, switchnode16port4, switchnode16port5, switchnode16port6}, Attrs: map[string]string{"role": "Switch", "name": "sw8"}}
)

var superGraphTest = &ConcreteGraph{
    Desc: "super",
    Nodes: []*ConcreteNode{
        dutnode1, atenode2, dutnode3, atenode4, dutnode5, atenode6, dutnode7, atenode8, dutnode9, atenode10, dutnode11, atenode12, dutnode13, atenode14, switchnode9, switchnode10, switchnode11, switchnode12, switchnode13, switchnode14, switchnode15, switchnode16, d1, a1,
    },
    Edges: []*ConcreteEdge{
        {Src: dutnode1port1, Dst: switchnode9port1},
        {Src: switchnode9port2, Dst: atenode2port1},
        {Src: dutnode1port2, Dst: switchnode9port3},
        {Src: switchnode9port4, Dst: atenode2port2},

        {Src: dutnode3port1, Dst: switchnode10port1},
        {Src: switchnode10port2, Dst: atenode4port1},
        {Src: dutnode3port2, Dst: switchnode10port3},
        {Src: switchnode10port4, Dst: atenode4port2},
        {Src: dutnode5port1, Dst: switchnode10port5},
        {Src: switchnode10port6, Dst: atenode6port1},
        {Src: dutnode7port1, Dst: atenode8port1},

        {Src: dutnode9port1, Dst: atenode10port1},
        {Src: dutnode9port1, Dst: atenode10port2},
        {Src: dutnode9port2, Dst: atenode10port1},

        {Src: dutnode11port1, Dst: switchnode11port1},
        {Src: dutnode11port2, Dst: switchnode11port2},
        {Src: dutnode11port3, Dst: switchnode11port3},
        {Src: dutnode11port4, Dst: switchnode11port4},
        {Src: switchnode11port3, Dst: switchnode12port3},
        {Src: switchnode12port1, Dst: atenode12port1},
        {Src: switchnode12port2, Dst: atenode12port2},
        {Src: switchnode12port3, Dst: atenode12port3},
        {Src: switchnode12port4, Dst: atenode12port4},

        {Src: dutnode13port1, Dst: switchnode13port1},
        {Src: dutnode13port2, Dst: switchnode13port2},
        {Src: dutnode13port3, Dst: switchnode13port3},
        {Src: dutnode13port4, Dst: switchnode13port4},
        {Src: dutnode13port5, Dst: switchnode13port5},
        {Src: switchnode13port6, Dst: switchnode14port1},
        {Src: switchnode14port2, Dst: switchnode15port1},
        {Src: switchnode15port2, Dst: switchnode16port6},
        {Src: switchnode16port1, Dst: atenode14port1},
        {Src: switchnode16port2, Dst: atenode14port2},
        {Src: switchnode16port3, Dst: atenode14port3},
        {Src: switchnode16port4, Dst: atenode14port4},
        {Src: switchnode16port5, Dst: atenode14port5},

        {Src: d1p1, Dst: switchnode9port5},
        {Src: switchnode9port6, Dst: a1p1},
        {Src: d1p2, Dst: switchnode9port7},
        {Src: switchnode9port8, Dst: a1p2},
    },
}

// Setup abstract Nodes and Ports for testing.
var (

    // Two Nodes, interconnected via Switch
    abst1      = &AbstractNode{Desc: "abst1", Ports: []*AbstractPort{abst1port1, abst1port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
    abst2      = &AbstractNode{Desc: "abst2", Ports: []*AbstractPort{abst2port1, abst2port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
    abst1port1 = &AbstractPort{Desc: "abst1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst1port2 = &AbstractPort{Desc: "abst1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst2port1 = &AbstractPort{Desc: "abst2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst2port2 = &AbstractPort{Desc: "abst2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    // Four Nodes, interconnected via Switch
    abst3      = &AbstractNode{Desc: "abst3", Ports: []*AbstractPort{abst3port1, abst3port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
    abst4      = &AbstractNode{Desc: "abst4", Ports: []*AbstractPort{abst4port1, abst4port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
    abst5      = &AbstractNode{Desc: "abst5", Ports: []*AbstractPort{abst5port1}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
    abst6      = &AbstractNode{Desc: "abst6", Ports: []*AbstractPort{abst6port1}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
    abst3port1 = &AbstractPort{Desc: "abst3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst3port2 = &AbstractPort{Desc: "abst3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst4port1 = &AbstractPort{Desc: "abst4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst4port2 = &AbstractPort{Desc: "abst4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst5port1 = &AbstractPort{Desc: "abst5:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst6port1 = &AbstractPort{Desc: "abst6:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}

    abst7      = &AbstractNode{Desc: "abst7", Ports: []*AbstractPort{abst7port1, abst7port2, abst7port3, abst7port4}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
    abst8      = &AbstractNode{Desc: "abst8", Ports: []*AbstractPort{abst8port1, abst8port2, abst8port3, abst8port4}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
    abst7port1 = &AbstractPort{Desc: "abst7:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst7port2 = &AbstractPort{Desc: "abst7:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst7port3 = &AbstractPort{Desc: "abst7:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst7port4 = &AbstractPort{Desc: "abst7:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst8port1 = &AbstractPort{Desc: "abst8:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst8port2 = &AbstractPort{Desc: "abst8:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst8port3 = &AbstractPort{Desc: "abst8:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst8port4 = &AbstractPort{Desc: "abst8:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}

    abst9       = &AbstractNode{Desc: "abst9", Ports: []*AbstractPort{abst9port1, abst9port2, abst9port3, abst9port4, abst9port5}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
    abst10      = &AbstractNode{Desc: "abst10", Ports: []*AbstractPort{abst10port1, abst10port2, abst10port3, abst10port4, abst10port5}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
    abst9port1  = &AbstractPort{Desc: "abst9:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst9port2  = &AbstractPort{Desc: "abst9:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst9port3  = &AbstractPort{Desc: "abst9:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst9port4  = &AbstractPort{Desc: "abst9:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst9port5  = &AbstractPort{Desc: "abst9:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst10port1 = &AbstractPort{Desc: "abst10:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst10port2 = &AbstractPort{Desc: "abst10:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst10port3 = &AbstractPort{Desc: "abst10:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst10port4 = &AbstractPort{Desc: "abst10:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
    abst10port5 = &AbstractPort{Desc: "abst10:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}

    abst11      = &AbstractNode{Desc: "abst11", Ports: []*AbstractPort{abst11port1, abst11port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
    abst12      = &AbstractNode{Desc: "abst12", Ports: []*AbstractPort{abst12port1, abst12port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
    abst11port1 = &AbstractPort{Desc: "abst11:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100GB")}}
    abst11port2 = &AbstractPort{Desc: "abst11:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100GB")}}
    abst12port1 = &AbstractPort{Desc: "abst12:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100GB")}}
    abst12port2 = &AbstractPort{Desc: "abst12:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100GB")}}
)

func TestSolveTest(t *testing.T) {
    tests := []struct {
        desc            string
        graph           *AbstractGraph
        wantNodes       map[*AbstractNode]*ConcreteNode
        wantPorts       map[*AbstractPort]*ConcretePort
        wantSolvedPorts []*AbstractPort
    }{{
        desc: "Two nodes, interconnected via Switch",
        graph: &AbstractGraph{
            Desc:  "two nodes, interconnected via Switch",
            Nodes: []*AbstractNode{abst1, abst2},
            Edges: []*AbstractEdge{{abst1port1, abst2port1}, {abst1port2, abst2port2}},
        },
        wantNodes: map[*AbstractNode]*ConcreteNode{abst1: dutnode1, abst2: atenode2},
        wantPorts: map[*AbstractPort]*ConcretePort{
            abst1port1: dutnode1port1,
            abst1port2: dutnode1port2,
            abst2port1: atenode2port1,
            abst2port2: atenode2port2,
        },
    }, {
        desc: "Four nodes, with multiple connections via Switch",
        graph: &AbstractGraph{
            Desc:  "four nodes, with multiple connections via Switch",
            Nodes: []*AbstractNode{abst3, abst4, abst5, abst6},
            Edges: []*AbstractEdge{{abst3port1, abst4port1}, {abst3port2, abst6port1}, {abst4port2, abst5port1}},
        },
        wantNodes: map[*AbstractNode]*ConcreteNode{abst3: dutnode3, abst4: atenode4, abst5: dutnode5, abst6: atenode6},
        wantPorts: map[*AbstractPort]*ConcretePort{
            abst3port1: dutnode3port1,
            abst3port2: dutnode3port2,
            abst4port1: atenode4port1,
            abst4port2: atenode4port2,
            abst5port1: dutnode5port1,
            abst6port1: atenode6port1,
        },
    }, {
        desc: "Two nodes, interconnected via multiple Switch (between single switch)",
        graph: &AbstractGraph{
            Desc:  "Two nodes, interconnected via multiple Switch (between single switch)",
            Nodes: []*AbstractNode{abst7, abst8},
            Edges: []*AbstractEdge{{abst7port1, abst8port1}, {abst7port2, abst8port2}, {abst7port3, abst8port3}, {abst7port4, abst8port4}},
        },
        wantNodes: map[*AbstractNode]*ConcreteNode{abst7: dutnode11, abst8: atenode12},
        wantPorts: map[*AbstractPort]*ConcretePort{
            abst7port1: dutnode11port1,
            abst7port2: dutnode11port2,
            abst7port3: dutnode11port3,
            abst7port4: dutnode11port4,
            abst8port1: atenode12port1,
            abst8port2: atenode12port2,
            abst8port3: atenode12port3,
            abst8port4: atenode12port4,
        },
    }, {
        desc: "Two nodes, interconnected via multiple Switch between two switches",
        graph: &AbstractGraph{
            Desc:  "Two nodes, interconnected via multiple Switch between two switches",
            Nodes: []*AbstractNode{abst9, abst10},
            Edges: []*AbstractEdge{{abst9port1, abst10port1}, {abst9port2, abst10port2}, {abst9port3, abst10port3}, {abst9port4, abst10port4}, {abst9port5, abst10port5}},
        },
        wantNodes: map[*AbstractNode]*ConcreteNode{abst9: dutnode13, abst10: atenode14},
        wantPorts: map[*AbstractPort]*ConcretePort{
            abst9port1:  dutnode13port1,
            abst9port2:  dutnode13port2,
            abst9port3:  dutnode13port3,
            abst9port4:  dutnode13port4,
            abst9port5:  dutnode13port5,
            abst10port1: atenode14port1,
            abst10port2: atenode14port2,
            abst10port3: atenode14port3,
            abst10port4: atenode14port4,
            abst10port5: atenode14port5,
        },
    }, {
        desc: "Two nodes, interconnected via Switch on different speeds",
        graph: &AbstractGraph{
            Desc:  "two nodes, interconnected via Switch on different speed",
            Nodes: []*AbstractNode{abst11, abst12},
            Edges: []*AbstractEdge{{abst11port1, abst12port1}, {abst11port2, abst12port2}},
        },
        wantNodes: map[*AbstractNode]*ConcreteNode{abst11: d1, abst12: a1},
        wantPorts: map[*AbstractPort]*ConcretePort{
            abst11port1: d1p1,
            abst11port2: d1p2,
            abst12port1: a1p1,
            abst12port2: a1p2,
        },
    }}
    for _, tc := range tests {
        t.Run(tc.desc, func(t *testing.T) {
            a, err := Solve(context.Background(), tc.graph, superGraphTest)
            if err != nil {
                t.Fatalf("Solve got error %v, want nil", err)
            }
            if len(a.Node2Node) != len(tc.wantNodes) {
                t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), len(tc.wantNodes))
            }
            for abs, got := range a.Node2Node {
                if want, ok := tc.wantNodes[abs]; !ok {
                    t.Fatalf("Solve assiged node abstract %q to %q; name does not exist", abs.Desc, got.Desc)
                } else if got != want {
                    t.Errorf("Solve for node %q got %q, want %q", abs.Desc, got.Desc, want.Desc)
                }
            }
            if len(a.Port2Port) != (len(tc.wantPorts) + len(tc.wantSolvedPorts)) {
                t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), (len(tc.wantPorts) + len(tc.wantSolvedPorts)))
            }
            for port, got := range a.Port2Port {
                if want, ok := tc.wantPorts[port]; !ok {
                    ok2 := false
                    for _, p := range tc.wantSolvedPorts {
                        if p == port {
                            ok2 = true
                            break
                        }
                    }
                    if !ok2 {
                        t.Fatalf("Solve assigned port %q to %q; port does not exist", port.Desc, got.Desc)
                    }
                } else {
                    gotPrefix := strings.Split(got.Desc, ":")[0]
                    wantPrefix := strings.Split(want.Desc, ":")[0]
                    if gotPrefix != wantPrefix {
                        t.Errorf("Solve for port %q got %q, want %q", port.Desc, got.Desc, want.Desc)
                    }
                }
            }
        })
    }
}

func TestSolveNotSolvableTest(t *testing.T) {
    nodeDesc1 := "dut1"
    nodeDesc2 := "dut2"
    nodeDesc3 := "dut3"
    portDesc1 := "dut:port1"
    portDesc2 := "dut:port2"
    portDesc3 := "dut:port3"
    portDesc4 := "dut:port4"
    portDesc5 := "dut:port5"
    portDesc6 := "dut:port6"
    portDesc7 := "dut:port7"
    portDesc8 := "dut:port8"
    samePort := &AbstractPort{Desc: portDesc1, Constraints: map[string]PortConstraint{"attr": Equal("BAR"), "attr2": Equal("1")}}
    aPort1 := &AbstractPort{Desc: portDesc1}
    aPort2 := &AbstractPort{Desc: portDesc2}
    aPort3 := &AbstractPort{Desc: portDesc3}
    aPort4 := &AbstractPort{Desc: portDesc4}
    zPort1 := &AbstractPort{Desc: portDesc5}
    zPort2 := &AbstractPort{Desc: portDesc6}
    zPort3 := &AbstractPort{Desc: portDesc7}
    zPort4 := &AbstractPort{Desc: portDesc8}

    tests := []struct {
        desc           string
        graph          *AbstractGraph
        wantAssigned   map[string]string
        wantUnassigned []string
    }{{
        desc: "Node does not exist in super",
        graph: &AbstractGraph{
            Desc:  "Node does not exist",
            Nodes: []*AbstractNode{&AbstractNode{Desc: nodeDesc1, Constraints: map[string]NodeConstraint{"DOES NOT EXIST": Equal("???")}}}},
        wantUnassigned: []string{nodeDesc1},
    }, {
        desc: "Port does not exist in super",
        graph: &AbstractGraph{
            Desc: "Port does not exist",
            Nodes: []*AbstractNode{
                &AbstractNode{
                    Desc:        nodeDesc1,
                    Ports:       []*AbstractPort{&AbstractPort{Desc: portDesc1, Constraints: map[string]PortConstraint{"DOES NOT EXIST": Equal("???")}}},
                    Constraints: map[string]NodeConstraint{"attr": Equal("multi1")},
                },
            },
        },
        wantAssigned:   map[string]string{nodeDesc1: "dutnode9"},
        wantUnassigned: []string{portDesc1},
    }, {
        desc: "One Port does not exist in super",
        graph: &AbstractGraph{
            Desc: "One port does not exist",
            Nodes: []*AbstractNode{
                &AbstractNode{
                    Desc: nodeDesc1,
                    Ports: []*AbstractPort{
                        samePort,
                        &AbstractPort{Desc: portDesc2, Constraints: map[string]PortConstraint{"attr": SameAsPort(samePort), "attr2": Equal("2")}},
                        // portDesc3 must be least constrained in order to be solved for last.
                        &AbstractPort{Desc: portDesc3, Constraints: map[string]PortConstraint{"attr2": Equal("1")}},
                    },
                    Constraints: map[string]NodeConstraint{"attr": Equal("multi1")},
                },
            },
        },
        wantAssigned:   map[string]string{nodeDesc1: "dutnode9", portDesc1: "dutnode9:port2", portDesc2: "dutnode9:port3"},
        wantUnassigned: []string{portDesc3},
    }, {
        desc: "Not enough edges to satisfy",
        graph: &AbstractGraph{
            Desc: "Not enough edges to satisfy",
            Nodes: []*AbstractNode{
                &AbstractNode{
                    Desc:        nodeDesc1,
                    Constraints: map[string]NodeConstraint{"vendor": Equal("UNIQUE90")},
                    Ports:       []*AbstractPort{aPort1, aPort2, aPort3, aPort4},
                },
                &AbstractNode{Desc: nodeDesc2, Ports: []*AbstractPort{zPort1, zPort2}},
                &AbstractNode{Desc: nodeDesc3, Ports: []*AbstractPort{zPort3, zPort4}},
            },
            Edges: []*AbstractEdge{{aPort1, zPort1}, {aPort2, zPort2}, {aPort3, zPort3}, {aPort4, zPort4}},
        },
        wantUnassigned: []string{nodeDesc1, nodeDesc2, nodeDesc3, portDesc1, portDesc2, portDesc3, portDesc4, portDesc5, portDesc6, portDesc7, portDesc8},
    }}
    for _, tc := range tests {
        t.Run(tc.desc, func(t *testing.T) {
            a, err := Solve(context.Background(), tc.graph, superGraphTest)
            if a != nil {
                t.Errorf("Solve got assignment %v, want nil", a)
            }
            if err == nil {
                t.Errorf("Solve got nil error, want error")
            }
            solveErr, ok := err.(*solveError)
            if !ok {
                t.Fatal("Solve got not *SolveErr type err, want *SolveErr type")
            }
            errString := solveErr.String()
            for abs, con := range tc.wantAssigned {
                re := regexp.MustCompile(fmt.Sprintf("%s.*%s.*%s", abs, "assigned", con))
                if !re.MatchString(errString) {
                    t.Errorf("Solve got error %q, want error to report %q is assigned to %q", errString, abs, con)
                }
            }
            for _, unassigned := range tc.wantUnassigned {
                re := regexp.MustCompile(fmt.Sprintf("%s.*%s", unassigned, "not assigned"))
                if !re.MatchString(errString) {
                    t.Errorf("Solve got error %q, want error to report %q is not assigned", errString, unassigned)
                }
            }
        })
    }
}

func TestSolveCancelledSwitch(t *testing.T) {
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    requestGraph := &AbstractGraph{
        Desc:  "four nodes, with multiple connections via switch",
        Nodes: []*AbstractNode{abst3, abst4, abst5, abst6},
        Edges: []*AbstractEdge{{abst3port1, abst4port1}, {abst3port2, abs4port2}, {abst5port1, abst6port1}},
    }
    _, err := Solve(ctx, requestGraph, superGraphTest)
    if err == nil {
        t.Error("Solve got nil error, want error due to cancel")
    }
}

// Benchmark solve for 4 DUT
// Run via blaze test with --test_arg=--test.bench=. --nocache_test_results
func Benchmark4DutSolveSwitch(b *testing.B) {
    b.ReportAllocs()
    requestGraph := &AbstractGraph{
        Desc:  "four nodes, with multiple connections via switch",
        Nodes: []*AbstractNode{abst3, abst4, abst5, abst6},
        Edges: []*AbstractEdge{{abst3port1, abst4port1}, {abst3port2, abst4port2}, {abst5port1, abst6port1}},
    }
    for i := 0; i < b.N; i++ {
        Solve(context.Background(), requestGraph, superGraphTest)
    }
}


