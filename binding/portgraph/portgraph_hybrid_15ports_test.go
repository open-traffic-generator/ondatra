// 4 Dut, 4 Ate with 15 ports and 3 switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	t15dut1port1  = &ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port2  = &ConcretePort{Desc: "dut1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port3  = &ConcretePort{Desc: "dut1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port4  = &ConcretePort{Desc: "dut1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port5  = &ConcretePort{Desc: "dut1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port6  = &ConcretePort{Desc: "dut1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port7  = &ConcretePort{Desc: "dut1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port8  = &ConcretePort{Desc: "dut1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port9  = &ConcretePort{Desc: "dut1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port10 = &ConcretePort{Desc: "dut1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port11 = &ConcretePort{Desc: "dut1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port12 = &ConcretePort{Desc: "dut1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port13 = &ConcretePort{Desc: "dut1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port14 = &ConcretePort{Desc: "dut1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1port15 = &ConcretePort{Desc: "dut1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut1       = &ConcreteNode{Desc: "dut1", Ports: []*ConcretePort{t15dut1port1, t15dut1port2, t15dut1port3, t15dut1port4, t15dut1port5, t15dut1port6, t15dut1port7, t15dut1port8, t15dut1port9, t15dut1port10, t15dut1port11, t15dut1port12, t15dut1port13, t15dut1port14, t15dut1port15}, Attrs: map[string]string{"vendor": "CISCO"}}

	t15dut2port1  = &ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port2  = &ConcretePort{Desc: "dut2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port3  = &ConcretePort{Desc: "dut2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port4  = &ConcretePort{Desc: "dut2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port5  = &ConcretePort{Desc: "dut2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port6  = &ConcretePort{Desc: "dut2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port7  = &ConcretePort{Desc: "dut2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port8  = &ConcretePort{Desc: "dut2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port9  = &ConcretePort{Desc: "dut2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port10 = &ConcretePort{Desc: "dut2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port11 = &ConcretePort{Desc: "dut2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port12 = &ConcretePort{Desc: "dut2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port13 = &ConcretePort{Desc: "dut2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port14 = &ConcretePort{Desc: "dut2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut2port15 = &ConcretePort{Desc: "dut2:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15dut2 = &ConcreteNode{Desc: "dut2", Ports: []*ConcretePort{t15dut2port1, t15dut2port2, t15dut2port3, t15dut2port4, t15dut2port5, t15dut2port6, t15dut2port7, t15dut2port8, t15dut2port9, t15dut2port10, t15dut2port11, t15dut2port12, t15dut2port13, t15dut2port14, t15dut2port15}, Attrs: map[string]string{"vendor": "CISCO"}}

	t15dut3port1  = &ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port2  = &ConcretePort{Desc: "dut3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port3  = &ConcretePort{Desc: "dut3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port4  = &ConcretePort{Desc: "dut3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port5  = &ConcretePort{Desc: "dut3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port6  = &ConcretePort{Desc: "dut3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port7  = &ConcretePort{Desc: "dut3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port8  = &ConcretePort{Desc: "dut3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port9  = &ConcretePort{Desc: "dut3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port10 = &ConcretePort{Desc: "dut3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port11 = &ConcretePort{Desc: "dut3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port12 = &ConcretePort{Desc: "dut3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port13 = &ConcretePort{Desc: "dut3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port14 = &ConcretePort{Desc: "dut3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut3port15 = &ConcretePort{Desc: "dut3:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15dut3 = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t15dut3port1, t15dut3port2, t15dut3port3, t15dut3port4, t15dut3port5, t15dut3port6, t15dut3port7, t15dut3port8, t15dut3port9, t15dut3port10, t15dut3port11, t15dut3port12, t15dut3port13, t15dut3port14, t15dut3port15}, Attrs: map[string]string{"vendor": "CISCO"}}
	// dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t15dut3port1}, Attrs: map[string]string{"vendor": "CISCO"}}
	t15dut4port1  = &ConcretePort{Desc: "dut4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port2  = &ConcretePort{Desc: "dut4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port3  = &ConcretePort{Desc: "dut4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port4  = &ConcretePort{Desc: "dut4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port5  = &ConcretePort{Desc: "dut4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port6  = &ConcretePort{Desc: "dut4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port7  = &ConcretePort{Desc: "dut4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port8  = &ConcretePort{Desc: "dut4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port9  = &ConcretePort{Desc: "dut4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port10 = &ConcretePort{Desc: "dut4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port11 = &ConcretePort{Desc: "dut4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port12 = &ConcretePort{Desc: "dut4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port13 = &ConcretePort{Desc: "dut4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port14 = &ConcretePort{Desc: "dut4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15dut4port15 = &ConcretePort{Desc: "dut4:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15dut4       = &ConcreteNode{Desc: "dut4", Ports: []*ConcretePort{t15dut4port1, t15dut4port2, t15dut4port3, t15dut4port4, t15dut4port5, t15dut4port6, t15dut4port7, t15dut4port8, t15dut4port9, t15dut4port10, t15dut4port11, t15dut4port12, t15dut4port13, t15dut4port14, t15dut4port15}, Attrs: map[string]string{"vendor": "CISCO"}}
	t15ate1port1  = &ConcretePort{Desc: "ate1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port2  = &ConcretePort{Desc: "ate1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port3  = &ConcretePort{Desc: "ate1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port4  = &ConcretePort{Desc: "ate1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port5  = &ConcretePort{Desc: "ate1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port6  = &ConcretePort{Desc: "ate1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port7  = &ConcretePort{Desc: "ate1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port8  = &ConcretePort{Desc: "ate1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port9  = &ConcretePort{Desc: "ate1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port10 = &ConcretePort{Desc: "ate1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port11 = &ConcretePort{Desc: "ate1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port12 = &ConcretePort{Desc: "ate1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port13 = &ConcretePort{Desc: "ate1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port14 = &ConcretePort{Desc: "ate1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate1port15 = &ConcretePort{Desc: "ate1:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15ate1 = &ConcreteNode{Desc: "ate1", Ports: []*ConcretePort{t15ate1port1, t15ate1port2, t15ate1port3, t15ate1port4, t15ate1port5, t15ate1port6, t15ate1port7, t15ate1port8, t15ate1port9, t15ate1port10, t15ate1port11, t15ate1port12, t15ate1port13, t15ate1port14, t15ate1port15}, Attrs: map[string]string{"vendor": "TGEN"}}

	t15ate2port1  = &ConcretePort{Desc: "ate2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port2  = &ConcretePort{Desc: "ate2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port3  = &ConcretePort{Desc: "ate2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port4  = &ConcretePort{Desc: "ate2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port5  = &ConcretePort{Desc: "ate2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port6  = &ConcretePort{Desc: "ate2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port7  = &ConcretePort{Desc: "ate2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port8  = &ConcretePort{Desc: "ate2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port9  = &ConcretePort{Desc: "ate2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port10 = &ConcretePort{Desc: "ate2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port11 = &ConcretePort{Desc: "ate2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port12 = &ConcretePort{Desc: "ate2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port13 = &ConcretePort{Desc: "ate2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port14 = &ConcretePort{Desc: "ate2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate2port15 = &ConcretePort{Desc: "ate2:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15ate2 = &ConcreteNode{Desc: "ate2", Ports: []*ConcretePort{t15ate2port1, t15ate2port2, t15ate2port3, t15ate2port4, t15ate2port5, t15ate2port6, t15ate2port7, t15ate2port8, t15ate2port9, t15ate2port10, t15ate2port11, t15ate2port12, t15ate2port13, t15ate2port14, t15ate2port15}, Attrs: map[string]string{"vendor": "TGEN"}}

	t15ate3port1  = &ConcretePort{Desc: "ate3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port2  = &ConcretePort{Desc: "ate3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port3  = &ConcretePort{Desc: "ate3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port4  = &ConcretePort{Desc: "ate3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port5  = &ConcretePort{Desc: "ate3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port6  = &ConcretePort{Desc: "ate3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port7  = &ConcretePort{Desc: "ate3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port8  = &ConcretePort{Desc: "ate3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port9  = &ConcretePort{Desc: "ate3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port10 = &ConcretePort{Desc: "ate3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port11 = &ConcretePort{Desc: "ate3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port12 = &ConcretePort{Desc: "ate3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port13 = &ConcretePort{Desc: "ate3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port14 = &ConcretePort{Desc: "ate3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate3port15 = &ConcretePort{Desc: "ate3:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15ate3 = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t15ate3port1, t15ate3port2, t15ate3port3, t15ate3port4, t15ate3port5, t15ate3port6, t15ate3port7, t15ate3port8, t15ate3port9, t15ate3port10, t15ate3port11, t15ate3port12, t15ate3port13, t15ate3port14, t15ate3port15}, Attrs: map[string]string{"vendor": "TGEN"}}
	// ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t15ate3port1}, Attrs: map[string]string{"vendor": "TGEN"}}
	t15ate4port1  = &ConcretePort{Desc: "ate4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port2  = &ConcretePort{Desc: "ate4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port3  = &ConcretePort{Desc: "ate4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port4  = &ConcretePort{Desc: "ate4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port5  = &ConcretePort{Desc: "ate4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port6  = &ConcretePort{Desc: "ate4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port7  = &ConcretePort{Desc: "ate4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port8  = &ConcretePort{Desc: "ate4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port9  = &ConcretePort{Desc: "ate4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port10 = &ConcretePort{Desc: "ate4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port11 = &ConcretePort{Desc: "ate4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port12 = &ConcretePort{Desc: "ate4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port13 = &ConcretePort{Desc: "ate4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port14 = &ConcretePort{Desc: "ate4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15ate4port15 = &ConcretePort{Desc: "ate4:port15", Attrs: map[string]string{"speed": "S_400G"}}

	t15ate4 = &ConcreteNode{Desc: "ate4", Ports: []*ConcretePort{t15ate4port1, t15ate4port2, t15ate4port3, t15ate4port4, t15ate4port5, t15ate4port6, t15ate4port7, t15ate4port8, t15ate4port9, t15ate4port10, t15ate4port11, t15ate4port12, t15ate4port13, t15ate4port14, t15ate4port15}, Attrs: map[string]string{"vendor": "TGEN"}}

	t15switch1port1  = &ConcretePort{Desc: "switch1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port2  = &ConcretePort{Desc: "switch1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port3  = &ConcretePort{Desc: "switch1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port4  = &ConcretePort{Desc: "switch1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port5  = &ConcretePort{Desc: "switch1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port6  = &ConcretePort{Desc: "switch1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port7  = &ConcretePort{Desc: "switch1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port8  = &ConcretePort{Desc: "switch1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port9  = &ConcretePort{Desc: "switch1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port10 = &ConcretePort{Desc: "switch1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port11 = &ConcretePort{Desc: "switch1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port12 = &ConcretePort{Desc: "switch1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port13 = &ConcretePort{Desc: "switch1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port14 = &ConcretePort{Desc: "switch1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port15 = &ConcretePort{Desc: "switch1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port16 = &ConcretePort{Desc: "switch1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port17 = &ConcretePort{Desc: "switch1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port18 = &ConcretePort{Desc: "switch1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port19 = &ConcretePort{Desc: "switch1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port20 = &ConcretePort{Desc: "switch1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port21 = &ConcretePort{Desc: "switch1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port22 = &ConcretePort{Desc: "switch1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port23 = &ConcretePort{Desc: "switch1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port24 = &ConcretePort{Desc: "switch1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port25 = &ConcretePort{Desc: "switch1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port26 = &ConcretePort{Desc: "switch1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port27 = &ConcretePort{Desc: "switch1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port28 = &ConcretePort{Desc: "switch1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port29 = &ConcretePort{Desc: "switch1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port30 = &ConcretePort{Desc: "switch1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch1port31 = &ConcretePort{Desc: "switch1:port31", Attrs: map[string]string{"speed": "S_400G"}}

	t15switch1 = &ConcreteNode{Desc: "switch1", Ports: []*ConcretePort{t15switch1port1, t15switch1port2, t15switch1port3, t15switch1port4, t15switch1port5, t15switch1port6, t15switch1port7, t15switch1port8, t15switch1port9, t15switch1port10, t15switch1port11, t15switch1port12, t15switch1port13, t15switch1port14, t15switch1port15, t15switch1port16, t15switch1port17, t15switch1port18, t15switch1port19, t15switch1port20, t15switch1port21, t15switch1port22, t15switch1port23, t15switch1port24, t15switch1port25, t15switch1port26, t15switch1port27, t15switch1port28, t15switch1port29, t15switch1port30, t15switch1port31}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	t15switch2port1  = &ConcretePort{Desc: "switch2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port2  = &ConcretePort{Desc: "switch2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port3  = &ConcretePort{Desc: "switch2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port4  = &ConcretePort{Desc: "switch2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port5  = &ConcretePort{Desc: "switch2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port6  = &ConcretePort{Desc: "switch2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port7  = &ConcretePort{Desc: "switch2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port8  = &ConcretePort{Desc: "switch2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port9  = &ConcretePort{Desc: "switch2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port10 = &ConcretePort{Desc: "switch2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port11 = &ConcretePort{Desc: "switch2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port12 = &ConcretePort{Desc: "switch2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port13 = &ConcretePort{Desc: "switch2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port14 = &ConcretePort{Desc: "switch2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port15 = &ConcretePort{Desc: "switch2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port16 = &ConcretePort{Desc: "switch2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port17 = &ConcretePort{Desc: "switch2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port18 = &ConcretePort{Desc: "switch2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port19 = &ConcretePort{Desc: "switch2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port20 = &ConcretePort{Desc: "switch2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port21 = &ConcretePort{Desc: "switch2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port22 = &ConcretePort{Desc: "switch2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port23 = &ConcretePort{Desc: "switch2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port24 = &ConcretePort{Desc: "switch2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port25 = &ConcretePort{Desc: "switch2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port26 = &ConcretePort{Desc: "switch2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port27 = &ConcretePort{Desc: "switch2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port28 = &ConcretePort{Desc: "switch2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port29 = &ConcretePort{Desc: "switch2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port30 = &ConcretePort{Desc: "switch2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port31 = &ConcretePort{Desc: "switch2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch2port32 = &ConcretePort{Desc: "switch2:port32", Attrs: map[string]string{"speed": "S_400G"}}

	t15switch2 = &ConcreteNode{Desc: "switch2", Ports: []*ConcretePort{t15switch2port1, t15switch2port2, t15switch2port3, t15switch2port4, t15switch2port5, t15switch2port6, t15switch2port7, t15switch2port8, t15switch2port9, t15switch2port10, t15switch2port11, t15switch2port12, t15switch2port13, t15switch2port14, t15switch2port15, t15switch2port16, t15switch2port17, t15switch2port18, t15switch2port19, t15switch2port20, t15switch2port21, t15switch2port22, t15switch2port23, t15switch2port24, t15switch2port25, t15switch2port26, t15switch2port27, t15switch2port28, t15switch2port29, t15switch2port30, t15switch2port31, t15switch2port32}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	t15switch3port1  = &ConcretePort{Desc: "switch3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port2  = &ConcretePort{Desc: "switch3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port3  = &ConcretePort{Desc: "switch3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port4  = &ConcretePort{Desc: "switch3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port5  = &ConcretePort{Desc: "switch3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port6  = &ConcretePort{Desc: "switch3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port7  = &ConcretePort{Desc: "switch3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port8  = &ConcretePort{Desc: "switch3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port9  = &ConcretePort{Desc: "switch3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port10 = &ConcretePort{Desc: "switch3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port11 = &ConcretePort{Desc: "switch3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port12 = &ConcretePort{Desc: "switch3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port13 = &ConcretePort{Desc: "switch3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port14 = &ConcretePort{Desc: "switch3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port15 = &ConcretePort{Desc: "switch3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port16 = &ConcretePort{Desc: "switch3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port17 = &ConcretePort{Desc: "switch3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port18 = &ConcretePort{Desc: "switch3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port19 = &ConcretePort{Desc: "switch3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port20 = &ConcretePort{Desc: "switch3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port21 = &ConcretePort{Desc: "switch3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port22 = &ConcretePort{Desc: "switch3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port23 = &ConcretePort{Desc: "switch3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port24 = &ConcretePort{Desc: "switch3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port25 = &ConcretePort{Desc: "switch3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port26 = &ConcretePort{Desc: "switch3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port27 = &ConcretePort{Desc: "switch3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port28 = &ConcretePort{Desc: "switch3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port29 = &ConcretePort{Desc: "switch3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port30 = &ConcretePort{Desc: "switch3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t15switch3port31 = &ConcretePort{Desc: "switch3:port31", Attrs: map[string]string{"speed": "S_400G"}}

	t15switch3 = &ConcreteNode{Desc: "switch3", Ports: []*ConcretePort{t15switch3port1, t15switch3port2, t15switch3port3, t15switch3port4, t15switch3port5, t15switch3port6, t15switch3port7, t15switch3port8, t15switch3port9, t15switch3port10, t15switch3port11, t15switch3port12, t15switch3port13, t15switch3port14, t15switch3port15, t15switch3port16, t15switch3port17, t15switch3port18, t15switch3port19, t15switch3port20, t15switch3port21, t15switch3port22, t15switch3port23, t15switch3port24, t15switch3port25, t15switch3port26, t15switch3port27, t15switch3port28, t15switch3port29, t15switch3port30, t15switch3port31}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}
)

var superGraphTest15 = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		t15dut1,
		t15ate1,
		t15dut2,
		t15ate2,
		t15dut3,
		t15ate3,
		t15dut4,
		t15ate4,
		t15switch1, t15switch2, t15switch3},
	Edges: []*ConcreteEdge{
		{Src: t15dut1port1, Dst: t15ate1port1},
		{Src: t15dut1port2, Dst: t15ate1port2},
		{Src: t15dut1port3, Dst: t15ate1port3},
		{Src: t15dut1port4, Dst: t15ate1port4},
		{Src: t15dut1port5, Dst: t15ate1port5},
		{Src: t15dut1port6, Dst: t15ate1port6},
		{Src: t15dut1port7, Dst: t15ate1port7},
		{Src: t15dut1port8, Dst: t15ate1port8},
		{Src: t15dut1port9, Dst: t15ate1port9},
		{Src: t15dut1port10, Dst: t15ate1port10},
		{Src: t15dut1port11, Dst: t15switch1port1},
		{Src: t15dut1port12, Dst: t15switch1port2},
		{Src: t15dut1port13, Dst: t15switch1port3},
		{Src: t15dut1port14, Dst: t15switch1port4},
		{Src: t15dut1port15, Dst: t15switch1port5},

		{Src: t15switch1port6, Dst: t15ate1port11},
		{Src: t15switch1port7, Dst: t15ate1port12},
		{Src: t15switch1port8, Dst: t15ate1port13},
		{Src: t15switch1port9, Dst: t15ate1port14},
		{Src: t15switch1port10, Dst: t15ate1port15},

		{Src: t15switch1port31, Dst: t15switch2port31},

		{Src: t15dut2port1, Dst: t15ate2port1},
		{Src: t15dut2port2, Dst: t15ate2port2},
		{Src: t15dut2port3, Dst: t15ate2port3},
		{Src: t15dut2port4, Dst: t15ate2port4},
		{Src: t15dut2port5, Dst: t15ate2port5},
		{Src: t15dut2port6, Dst: t15ate2port6},
		{Src: t15dut2port7, Dst: t15ate2port7},
		{Src: t15dut2port8, Dst: t15ate2port8},
		{Src: t15dut2port9, Dst: t15ate2port9},
		{Src: t15dut2port10, Dst: t15ate2port10},
		{Src: t15dut2port11, Dst: t15switch1port11},
		{Src: t15dut2port12, Dst: t15switch1port12},
		{Src: t15dut2port13, Dst: t15switch1port13},
		{Src: t15dut2port14, Dst: t15switch1port14},
		{Src: t15dut2port15, Dst: t15switch1port15},

		{Src: t15switch1port16, Dst: t15ate2port11},
		{Src: t15switch1port17, Dst: t15ate2port12},
		{Src: t15switch1port18, Dst: t15ate2port13},
		{Src: t15switch1port19, Dst: t15ate2port14},
		{Src: t15switch1port20, Dst: t15ate2port15},

		{Src: t15dut3port1, Dst: t15switch2port1},
		{Src: t15dut3port2, Dst: t15switch2port2},
		{Src: t15dut3port3, Dst: t15switch2port3},
		{Src: t15dut3port4, Dst: t15switch2port4},
		{Src: t15dut3port5, Dst: t15switch2port5},
		{Src: t15dut3port6, Dst: t15switch2port6},
		{Src: t15dut3port7, Dst: t15switch2port7},
		{Src: t15dut3port8, Dst: t15switch2port8},
		{Src: t15dut3port9, Dst: t15switch2port9},
		{Src: t15dut3port10, Dst: t15switch2port10},
		{Src: t15switch2port11, Dst: t15ate3port1},
		{Src: t15switch2port12, Dst: t15ate3port2},
		{Src: t15switch2port13, Dst: t15ate3port3},
		{Src: t15switch2port14, Dst: t15ate3port4},
		{Src: t15switch2port15, Dst: t15ate3port5},
		{Src: t15switch2port16, Dst: t15ate3port6},
		{Src: t15switch2port17, Dst: t15ate3port7},
		{Src: t15switch2port18, Dst: t15ate3port8},
		{Src: t15switch2port19, Dst: t15ate3port9},
		{Src: t15switch2port20, Dst: t15ate3port10},
		{Src: t15dut4port11, Dst: t15switch2port21},
		{Src: t15dut4port12, Dst: t15switch2port22},
		{Src: t15dut4port13, Dst: t15switch2port23},
		{Src: t15dut4port14, Dst: t15switch2port24},
		{Src: t15dut4port15, Dst: t15switch2port25},

		{Src: t15switch2port26, Dst: t15ate4port11},
		{Src: t15switch2port27, Dst: t15ate4port12},
		{Src: t15switch2port28, Dst: t15ate4port13},
		{Src: t15switch2port29, Dst: t15ate4port14},
		{Src: t15switch2port30, Dst: t15ate4port15},

		{Src: t15switch2port32, Dst: t15switch3port31},

		{Src: t15dut3port11, Dst: t15switch3port1},
		{Src: t15dut3port12, Dst: t15switch3port2},
		{Src: t15dut3port13, Dst: t15switch3port3},
		{Src: t15dut3port14, Dst: t15switch3port4},
		{Src: t15dut3port15, Dst: t15switch3port5},

		{Src: t15switch3port6, Dst: t15ate3port11},
		{Src: t15switch3port7, Dst: t15ate3port12},
		{Src: t15switch3port8, Dst: t15ate3port13},
		{Src: t15switch3port9, Dst: t15ate3port14},
		{Src: t15switch3port10, Dst: t15ate3port15},

		{Src: t15dut4port1, Dst: t15switch3port21},
		{Src: t15dut4port2, Dst: t15switch3port22},
		{Src: t15dut4port3, Dst: t15switch3port23},
		{Src: t15dut4port4, Dst: t15switch3port24},
		{Src: t15dut4port5, Dst: t15switch3port25},
		{Src: t15dut4port6, Dst: t15switch3port26},
		{Src: t15dut4port7, Dst: t15switch3port27},
		{Src: t15dut4port8, Dst: t15switch3port28},
		{Src: t15dut4port9, Dst: t15switch3port29},
		{Src: t15dut4port10, Dst: t15switch3port30},
		{Src: t15switch3port11, Dst: t15ate4port1},
		{Src: t15switch3port12, Dst: t15ate4port2},
		{Src: t15switch3port13, Dst: t15ate4port3},
		{Src: t15switch3port14, Dst: t15ate4port4},
		{Src: t15switch3port15, Dst: t15ate4port5},
		{Src: t15switch3port16, Dst: t15ate4port6},
		{Src: t15switch3port17, Dst: t15ate4port7},
		{Src: t15switch3port18, Dst: t15ate4port8},
		{Src: t15switch3port19, Dst: t15ate4port9},
		{Src: t15switch3port20, Dst: t15ate4port10},
	},
}

var (
	abst15dut1port1  = &AbstractPort{Desc: "absdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port2  = &AbstractPort{Desc: "absdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port3  = &AbstractPort{Desc: "absdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port4  = &AbstractPort{Desc: "absdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port5  = &AbstractPort{Desc: "absdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port6  = &AbstractPort{Desc: "absdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port7  = &AbstractPort{Desc: "absdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port8  = &AbstractPort{Desc: "absdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port9  = &AbstractPort{Desc: "absdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port10 = &AbstractPort{Desc: "absdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port11 = &AbstractPort{Desc: "absdut1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port12 = &AbstractPort{Desc: "absdut1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port13 = &AbstractPort{Desc: "absdut1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port14 = &AbstractPort{Desc: "absdut1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1port15 = &AbstractPort{Desc: "absdut1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut1       = &AbstractNode{Desc: "absdut1", Ports: []*AbstractPort{abst15dut1port1, abst15dut1port2, abst15dut1port3, abst15dut1port4, abst15dut1port5, abst15dut1port6, abst15dut1port7, abst15dut1port8, abst15dut1port9, abst15dut1port10, abst15dut1port11, abst15dut1port12, abst15dut1port13, abst15dut1port14, abst15dut1port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst15dut2port1  = &AbstractPort{Desc: "absdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port2  = &AbstractPort{Desc: "absdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port3  = &AbstractPort{Desc: "absdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port4  = &AbstractPort{Desc: "absdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port5  = &AbstractPort{Desc: "absdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port6  = &AbstractPort{Desc: "absdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port7  = &AbstractPort{Desc: "absdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port8  = &AbstractPort{Desc: "absdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port9  = &AbstractPort{Desc: "absdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port10 = &AbstractPort{Desc: "absdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port11 = &AbstractPort{Desc: "absdut2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port12 = &AbstractPort{Desc: "absdut2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port13 = &AbstractPort{Desc: "absdut2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port14 = &AbstractPort{Desc: "absdut2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut2port15 = &AbstractPort{Desc: "absdut2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15dut2 = &AbstractNode{Desc: "absdut2", Ports: []*AbstractPort{abst15dut2port1, abst15dut2port2, abst15dut2port3, abst15dut2port4, abst15dut2port5, abst15dut2port6, abst15dut2port7, abst15dut2port8, abst15dut2port9, abst15dut2port10, abst15dut2port11, abst15dut2port12, abst15dut2port13, abst15dut2port14, abst15dut2port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst15dut3port1  = &AbstractPort{Desc: "absdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port2  = &AbstractPort{Desc: "absdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port3  = &AbstractPort{Desc: "absdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port4  = &AbstractPort{Desc: "absdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port5  = &AbstractPort{Desc: "absdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port6  = &AbstractPort{Desc: "absdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port7  = &AbstractPort{Desc: "absdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port8  = &AbstractPort{Desc: "absdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port9  = &AbstractPort{Desc: "absdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port10 = &AbstractPort{Desc: "absdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port11 = &AbstractPort{Desc: "absdut3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port12 = &AbstractPort{Desc: "absdut3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port13 = &AbstractPort{Desc: "absdut3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port14 = &AbstractPort{Desc: "absdut3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut3port15 = &AbstractPort{Desc: "absdut3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15dut3 = &AbstractNode{Desc: "absdut3", Ports: []*AbstractPort{abst15dut3port1, abst15dut3port2, abst15dut3port3, abst15dut3port4, abst15dut3port5, abst15dut3port6, abst15dut3port7, abst15dut3port8, abst15dut3port9, abst15dut3port10, abst15dut3port11, abst15dut3port12, abst15dut3port13, abst15dut3port14, abst15dut3port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst15dut4port1  = &AbstractPort{Desc: "absdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port2  = &AbstractPort{Desc: "absdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port3  = &AbstractPort{Desc: "absdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port4  = &AbstractPort{Desc: "absdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port5  = &AbstractPort{Desc: "absdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port6  = &AbstractPort{Desc: "absdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port7  = &AbstractPort{Desc: "absdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port8  = &AbstractPort{Desc: "absdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port9  = &AbstractPort{Desc: "absdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port10 = &AbstractPort{Desc: "absdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port11 = &AbstractPort{Desc: "absdut4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port12 = &AbstractPort{Desc: "absdut4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port13 = &AbstractPort{Desc: "absdut4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port14 = &AbstractPort{Desc: "absdut4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15dut4port15 = &AbstractPort{Desc: "absdut4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15dut4 = &AbstractNode{Desc: "absdut4", Ports: []*AbstractPort{abst15dut4port1, abst15dut4port2, abst15dut4port3, abst15dut4port4, abst15dut4port5, abst15dut4port6, abst15dut4port7, abst15dut4port8, abst15dut4port9, abst15dut4port10, abst15dut4port11, abst15dut4port12, abst15dut4port13, abst15dut4port14, abst15dut4port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst15ate1port1  = &AbstractPort{Desc: "absate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port2  = &AbstractPort{Desc: "absate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port3  = &AbstractPort{Desc: "absate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port4  = &AbstractPort{Desc: "absate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port5  = &AbstractPort{Desc: "absate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port6  = &AbstractPort{Desc: "absate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port7  = &AbstractPort{Desc: "absate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port8  = &AbstractPort{Desc: "absate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port9  = &AbstractPort{Desc: "absate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port10 = &AbstractPort{Desc: "absate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port11 = &AbstractPort{Desc: "absate1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port12 = &AbstractPort{Desc: "absate1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port13 = &AbstractPort{Desc: "absate1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port14 = &AbstractPort{Desc: "absate1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate1port15 = &AbstractPort{Desc: "absate1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15ate1 = &AbstractNode{Desc: "absate1", Ports: []*AbstractPort{abst15ate1port1, abst15ate1port2, abst15ate1port3, abst15ate1port4, abst15ate1port5, abst15ate1port6, abst15ate1port7, abst15ate1port8, abst15ate1port9, abst15ate1port10, abst15ate1port11, abst15ate1port12, abst15ate1port13, abst15ate1port14, abst15ate1port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst15ate2port1  = &AbstractPort{Desc: "absate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port2  = &AbstractPort{Desc: "absate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port3  = &AbstractPort{Desc: "absate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port4  = &AbstractPort{Desc: "absate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port5  = &AbstractPort{Desc: "absate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port6  = &AbstractPort{Desc: "absate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port7  = &AbstractPort{Desc: "absate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port8  = &AbstractPort{Desc: "absate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port9  = &AbstractPort{Desc: "absate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port10 = &AbstractPort{Desc: "absate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port11 = &AbstractPort{Desc: "absate2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port12 = &AbstractPort{Desc: "absate2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port13 = &AbstractPort{Desc: "absate2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port14 = &AbstractPort{Desc: "absate2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate2port15 = &AbstractPort{Desc: "absate2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15ate2 = &AbstractNode{Desc: "absate2", Ports: []*AbstractPort{abst15ate2port1, abst15ate2port2, abst15ate2port3, abst15ate2port4, abst15ate2port5, abst15ate2port6, abst15ate2port7, abst15ate2port8, abst15ate2port9, abst15ate2port10, abst15ate2port11, abst15ate2port12, abst15ate2port13, abst15ate2port14, abst15ate2port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst15ate3port1  = &AbstractPort{Desc: "absate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port2  = &AbstractPort{Desc: "absate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port3  = &AbstractPort{Desc: "absate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port4  = &AbstractPort{Desc: "absate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port5  = &AbstractPort{Desc: "absate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port6  = &AbstractPort{Desc: "absate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port7  = &AbstractPort{Desc: "absate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port8  = &AbstractPort{Desc: "absate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port9  = &AbstractPort{Desc: "absate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port10 = &AbstractPort{Desc: "absate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port11 = &AbstractPort{Desc: "absate3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port12 = &AbstractPort{Desc: "absate3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port13 = &AbstractPort{Desc: "absate3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port14 = &AbstractPort{Desc: "absate3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate3port15 = &AbstractPort{Desc: "absate3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15ate3 = &AbstractNode{Desc: "absate3", Ports: []*AbstractPort{abst15ate3port1, abst15ate3port2, abst15ate3port3, abst15ate3port4, abst15ate3port5, abst15ate3port6, abst15ate3port7, abst15ate3port8, abst15ate3port9, abst15ate3port10, abst15ate3port11, abst15ate3port12, abst15ate3port13, abst15ate3port14, abst15ate3port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst15ate4port1  = &AbstractPort{Desc: "absate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port2  = &AbstractPort{Desc: "absate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port3  = &AbstractPort{Desc: "absate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port4  = &AbstractPort{Desc: "absate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port5  = &AbstractPort{Desc: "absate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port6  = &AbstractPort{Desc: "absate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port7  = &AbstractPort{Desc: "absate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port8  = &AbstractPort{Desc: "absate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port9  = &AbstractPort{Desc: "absate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port10 = &AbstractPort{Desc: "absate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port11 = &AbstractPort{Desc: "absate4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port12 = &AbstractPort{Desc: "absate4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port13 = &AbstractPort{Desc: "absate4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port14 = &AbstractPort{Desc: "absate4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst15ate4port15 = &AbstractPort{Desc: "absate4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abst15ate4 = &AbstractNode{Desc: "absate4", Ports: []*AbstractPort{abst15ate4port1, abst15ate4port2, abst15ate4port3, abst15ate4port4, abst15ate4port5, abst15ate4port6, abst15ate4port7, abst15ate4port8, abst15ate4port9, abst15ate4port10, abst15ate4port11, abst15ate4port12, abst15ate4port13, abst15ate4port14, abst15ate4port15}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
)

var abstractGraph15 = &AbstractGraph{
	Nodes: []*AbstractNode{abst15dut1,
		abst15ate1,
		abst15dut2,
		abst15ate2,
		abst15dut3,
		abst15ate3,
		abst15dut4,
		abst15ate4},
	Edges: []*AbstractEdge{
		{abst15dut1port1, abst15ate1port1},
		{abst15dut1port2, abst15ate1port2},
		{abst15dut1port3, abst15ate1port3},
		{abst15dut1port4, abst15ate1port4},
		{abst15dut1port5, abst15ate1port5},
		{abst15dut1port6, abst15ate1port6},
		{abst15dut1port7, abst15ate1port7},
		{abst15dut1port8, abst15ate1port8},
		{abst15dut1port9, abst15ate1port9},
		{abst15dut1port10, abst15ate1port10},
		{abst15dut1port11, abst15ate1port11},
		{abst15dut1port12, abst15ate1port12},
		{abst15dut1port13, abst15ate1port13},
		{abst15dut1port14, abst15ate1port14},
		{abst15dut1port15, abst15ate1port15},

		{abst15dut2port1, abst15ate2port1},
		{abst15dut2port2, abst15ate2port2},
		{abst15dut2port3, abst15ate2port3},
		{abst15dut2port4, abst15ate2port4},
		{abst15dut2port5, abst15ate2port5},
		{abst15dut2port6, abst15ate2port6},
		{abst15dut2port7, abst15ate2port7},
		{abst15dut2port8, abst15ate2port8},
		{abst15dut2port9, abst15ate2port9},
		{abst15dut2port10, abst15ate2port10},
		{abst15dut2port11, abst15ate2port11},
		{abst15dut2port12, abst15ate2port12},
		{abst15dut2port13, abst15ate2port13},
		{abst15dut2port14, abst15ate2port14},
		{abst15dut2port15, abst15ate2port15},

		{abst15dut3port1, abst15ate3port1},
		{abst15dut3port2, abst15ate3port2},
		{abst15dut3port3, abst15ate3port3},
		{abst15dut3port4, abst15ate3port4},
		{abst15dut3port5, abst15ate3port5},
		{abst15dut3port6, abst15ate3port6},
		{abst15dut3port7, abst15ate3port7},
		{abst15dut3port8, abst15ate3port8},
		{abst15dut3port9, abst15ate3port9},
		{abst15dut3port10, abst15ate3port10},
		{abst15dut3port11, abst15ate3port11},
		{abst15dut3port12, abst15ate3port12},
		{abst15dut3port13, abst15ate3port13},
		{abst15dut3port14, abst15ate3port14},
		{abst15dut3port15, abst15ate3port15},

		{abst15dut4port1, abst15ate4port1},
		{abst15dut4port2, abst15ate4port2},
		{abst15dut4port3, abst15ate4port3},
		{abst15dut4port4, abst15ate4port4},
		{abst15dut4port5, abst15ate4port5},
		{abst15dut4port6, abst15ate4port6},
		{abst15dut4port7, abst15ate4port7},
		{abst15dut4port8, abst15ate4port8},
		{abst15dut4port9, abst15ate4port9},
		{abst15dut4port10, abst15ate4port10},
		{abst15dut4port11, abst15ate4port11},
		{abst15dut4port12, abst15ate4port12},
		{abst15dut4port13, abst15ate4port13},
		{abst15dut4port14, abst15ate4port14},
		{abst15dut4port15, abst15ate4port15}},
}

func TestSolveTestHybrid15(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraph15, superGraphTest15)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 8 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 120 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
