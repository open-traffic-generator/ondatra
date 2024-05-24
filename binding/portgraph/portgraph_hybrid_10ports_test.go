// Hybrid 4 duts, 4 ates , each 10 ports, 3 switches
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	t10dut1port1  = &ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port2  = &ConcretePort{Desc: "dut1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port3  = &ConcretePort{Desc: "dut1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port4  = &ConcretePort{Desc: "dut1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port5  = &ConcretePort{Desc: "dut1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port6  = &ConcretePort{Desc: "dut1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port7  = &ConcretePort{Desc: "dut1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port8  = &ConcretePort{Desc: "dut1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port9  = &ConcretePort{Desc: "dut1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1port10 = &ConcretePort{Desc: "dut1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut1       = &ConcreteNode{Desc: "dut1", Ports: []*ConcretePort{t10dut1port1, t10dut1port2, t10dut1port3, t10dut1port4, t10dut1port5, t10dut1port6, t10dut1port7, t10dut1port8, t10dut1port9, t10dut1port10}, Attrs: map[string]string{"vendor": "CISCO"}}
	t10dut2port1  = &ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port2  = &ConcretePort{Desc: "dut2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port3  = &ConcretePort{Desc: "dut2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port4  = &ConcretePort{Desc: "dut2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port5  = &ConcretePort{Desc: "dut2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port6  = &ConcretePort{Desc: "dut2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port7  = &ConcretePort{Desc: "dut2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port8  = &ConcretePort{Desc: "dut2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port9  = &ConcretePort{Desc: "dut2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2port10 = &ConcretePort{Desc: "dut2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut2       = &ConcreteNode{Desc: "dut2", Ports: []*ConcretePort{t10dut2port1, t10dut2port2, t10dut2port3, t10dut2port4, t10dut2port5, t10dut2port6, t10dut2port7, t10dut2port8, t10dut2port9, t10dut2port10}, Attrs: map[string]string{"vendor": "CISCO"}}
	t10dut3port1  = &ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port2  = &ConcretePort{Desc: "dut3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port3  = &ConcretePort{Desc: "dut3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port4  = &ConcretePort{Desc: "dut3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port5  = &ConcretePort{Desc: "dut3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port6  = &ConcretePort{Desc: "dut3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port7  = &ConcretePort{Desc: "dut3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port8  = &ConcretePort{Desc: "dut3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port9  = &ConcretePort{Desc: "dut3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3port10 = &ConcretePort{Desc: "dut3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t10dut3port1, t10dut3port2, t10dut3port3, t10dut3port4, t10dut3port5, t10dut3port6, t10dut3port7, t10dut3port8, t10dut3port9, t10dut3port10}, Attrs: map[string]string{"vendor": "CISCO"}}
	// dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t10dut3port1}, Attrs: map[string]string{"vendor": "CISCO"}}
	t10dut4port1  = &ConcretePort{Desc: "dut4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port2  = &ConcretePort{Desc: "dut4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port3  = &ConcretePort{Desc: "dut4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port4  = &ConcretePort{Desc: "dut4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port5  = &ConcretePort{Desc: "dut4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port6  = &ConcretePort{Desc: "dut4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port7  = &ConcretePort{Desc: "dut4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port8  = &ConcretePort{Desc: "dut4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port9  = &ConcretePort{Desc: "dut4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4port10 = &ConcretePort{Desc: "dut4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10dut4       = &ConcreteNode{Desc: "dut4", Ports: []*ConcretePort{t10dut4port1, t10dut4port2, t10dut4port3, t10dut4port4, t10dut4port5, t10dut4port6, t10dut4port7, t10dut4port8, t10dut4port9, t10dut4port10}, Attrs: map[string]string{"vendor": "CISCO"}}

	t10ate1port1  = &ConcretePort{Desc: "ate1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port2  = &ConcretePort{Desc: "ate1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port3  = &ConcretePort{Desc: "ate1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port4  = &ConcretePort{Desc: "ate1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port5  = &ConcretePort{Desc: "ate1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port6  = &ConcretePort{Desc: "ate1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port7  = &ConcretePort{Desc: "ate1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port8  = &ConcretePort{Desc: "ate1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port9  = &ConcretePort{Desc: "ate1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1port10 = &ConcretePort{Desc: "ate1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate1       = &ConcreteNode{Desc: "ate1", Ports: []*ConcretePort{t10ate1port1, t10ate1port2, t10ate1port3, t10ate1port4, t10ate1port5, t10ate1port6, t10ate1port7, t10ate1port8, t10ate1port9, t10ate1port10}, Attrs: map[string]string{"vendor": "TGEN"}}
	t10ate2port1  = &ConcretePort{Desc: "ate2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port2  = &ConcretePort{Desc: "ate2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port3  = &ConcretePort{Desc: "ate2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port4  = &ConcretePort{Desc: "ate2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port5  = &ConcretePort{Desc: "ate2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port6  = &ConcretePort{Desc: "ate2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port7  = &ConcretePort{Desc: "ate2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port8  = &ConcretePort{Desc: "ate2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port9  = &ConcretePort{Desc: "ate2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2port10 = &ConcretePort{Desc: "ate2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate2       = &ConcreteNode{Desc: "ate2", Ports: []*ConcretePort{t10ate2port1, t10ate2port2, t10ate2port3, t10ate2port4, t10ate2port5, t10ate2port6, t10ate2port7, t10ate2port8, t10ate2port9, t10ate2port10}, Attrs: map[string]string{"vendor": "TGEN"}}
	t10ate3port1  = &ConcretePort{Desc: "ate3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port2  = &ConcretePort{Desc: "ate3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port3  = &ConcretePort{Desc: "ate3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port4  = &ConcretePort{Desc: "ate3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port5  = &ConcretePort{Desc: "ate3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port6  = &ConcretePort{Desc: "ate3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port7  = &ConcretePort{Desc: "ate3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port8  = &ConcretePort{Desc: "ate3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port9  = &ConcretePort{Desc: "ate3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3port10 = &ConcretePort{Desc: "ate3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t10ate3port1, t10ate3port2, t10ate3port3, t10ate3port4, t10ate3port5, t10ate3port6, t10ate3port7, t10ate3port8, t10ate3port9, t10ate3port10}, Attrs: map[string]string{"vendor": "TGEN"}}
	// ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t10ate3port1}, Attrs: map[string]string{"vendor": "TGEN"}}
	t10ate4port1  = &ConcretePort{Desc: "ate4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port2  = &ConcretePort{Desc: "ate4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port3  = &ConcretePort{Desc: "ate4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port4  = &ConcretePort{Desc: "ate4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port5  = &ConcretePort{Desc: "ate4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port6  = &ConcretePort{Desc: "ate4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port7  = &ConcretePort{Desc: "ate4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port8  = &ConcretePort{Desc: "ate4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port9  = &ConcretePort{Desc: "ate4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4port10 = &ConcretePort{Desc: "ate4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10ate4       = &ConcreteNode{Desc: "ate4", Ports: []*ConcretePort{t10ate4port1, t10ate4port2, t10ate4port3, t10ate4port4, t10ate4port5, t10ate4port6, t10ate4port7, t10ate4port8, t10ate4port9, t10ate4port10}, Attrs: map[string]string{"vendor": "TGEN"}}

	t10switch1port1  = &ConcretePort{Desc: "switch1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port2  = &ConcretePort{Desc: "switch1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port3  = &ConcretePort{Desc: "switch1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port4  = &ConcretePort{Desc: "switch1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port5  = &ConcretePort{Desc: "switch1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port6  = &ConcretePort{Desc: "switch1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port7  = &ConcretePort{Desc: "switch1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port8  = &ConcretePort{Desc: "switch1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port9  = &ConcretePort{Desc: "switch1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port10 = &ConcretePort{Desc: "switch1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port11 = &ConcretePort{Desc: "switch1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port12 = &ConcretePort{Desc: "switch1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port13 = &ConcretePort{Desc: "switch1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port14 = &ConcretePort{Desc: "switch1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port15 = &ConcretePort{Desc: "switch1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port16 = &ConcretePort{Desc: "switch1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port17 = &ConcretePort{Desc: "switch1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port18 = &ConcretePort{Desc: "switch1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port19 = &ConcretePort{Desc: "switch1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port20 = &ConcretePort{Desc: "switch1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1port21 = &ConcretePort{Desc: "switch1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch1       = &ConcreteNode{Desc: "switch1", Ports: []*ConcretePort{t10switch1port1, t10switch1port2, t10switch1port3, t10switch1port4, t10switch1port5, t10switch1port6, t10switch1port7, t10switch1port8, t10switch1port9, t10switch1port10, t10switch1port11, t10switch1port12, t10switch1port13, t10switch1port14, t10switch1port15, t10switch1port16, t10switch1port17, t10switch1port18, t10switch1port19, t10switch1port20, t10switch1port21}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	t10switch2port1  = &ConcretePort{Desc: "switch2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port2  = &ConcretePort{Desc: "switch2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port3  = &ConcretePort{Desc: "switch2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port4  = &ConcretePort{Desc: "switch2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port5  = &ConcretePort{Desc: "switch2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port6  = &ConcretePort{Desc: "switch2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port7  = &ConcretePort{Desc: "switch2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port8  = &ConcretePort{Desc: "switch2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port9  = &ConcretePort{Desc: "switch2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port10 = &ConcretePort{Desc: "switch2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port11 = &ConcretePort{Desc: "switch2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port12 = &ConcretePort{Desc: "switch2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port13 = &ConcretePort{Desc: "switch2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port14 = &ConcretePort{Desc: "switch2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port15 = &ConcretePort{Desc: "switch2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port16 = &ConcretePort{Desc: "switch2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port17 = &ConcretePort{Desc: "switch2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port18 = &ConcretePort{Desc: "switch2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port19 = &ConcretePort{Desc: "switch2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port20 = &ConcretePort{Desc: "switch2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port21 = &ConcretePort{Desc: "switch2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2port22 = &ConcretePort{Desc: "switch2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch2       = &ConcreteNode{Desc: "switch2", Ports: []*ConcretePort{t10switch2port1, t10switch2port2, t10switch2port3, t10switch2port4, t10switch2port5, t10switch2port6, t10switch2port7, t10switch2port8, t10switch2port9, t10switch2port10, t10switch2port11, t10switch2port12, t10switch2port13, t10switch2port14, t10switch2port15, t10switch2port16, t10switch2port17, t10switch2port18, t10switch2port19, t10switch2port20, t10switch2port21, t10switch2port22}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	t10switch3port1  = &ConcretePort{Desc: "switch3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port2  = &ConcretePort{Desc: "switch3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port3  = &ConcretePort{Desc: "switch3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port4  = &ConcretePort{Desc: "switch3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port5  = &ConcretePort{Desc: "switch3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port6  = &ConcretePort{Desc: "switch3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port7  = &ConcretePort{Desc: "switch3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port8  = &ConcretePort{Desc: "switch3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port9  = &ConcretePort{Desc: "switch3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port10 = &ConcretePort{Desc: "switch3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port11 = &ConcretePort{Desc: "switch3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port12 = &ConcretePort{Desc: "switch3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port13 = &ConcretePort{Desc: "switch3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port14 = &ConcretePort{Desc: "switch3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port15 = &ConcretePort{Desc: "switch3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port16 = &ConcretePort{Desc: "switch3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port17 = &ConcretePort{Desc: "switch3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port18 = &ConcretePort{Desc: "switch3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port19 = &ConcretePort{Desc: "switch3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port20 = &ConcretePort{Desc: "switch3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3port21 = &ConcretePort{Desc: "switch3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t10switch3       = &ConcreteNode{Desc: "switch3", Ports: []*ConcretePort{t10switch3port1, t10switch3port2, t10switch3port3, t10switch3port4, t10switch3port5, t10switch3port6, t10switch3port7, t10switch3port8, t10switch3port9, t10switch3port10, t10switch3port11, t10switch3port12, t10switch3port13, t10switch3port14, t10switch3port15, t10switch3port16, t10switch3port17, t10switch3port18, t10switch3port19, t10switch3port20, t10switch3port21}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}
)

var superGraphTest10 = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		t10dut1,
		t10ate1,
		t10dut2,
		t10ate2,
		t10dut3,
		t10ate3,
		t10dut4,
		t10ate4,
		t10switch1, t10switch2, t10switch3},
	Edges: []*ConcreteEdge{
		{Src: t10dut1port1, Dst: t10ate1port1},
		{Src: t10dut1port2, Dst: t10ate1port2},
		{Src: t10dut1port3, Dst: t10ate1port3},
		{Src: t10dut1port4, Dst: t10ate1port4},
		{Src: t10dut1port5, Dst: t10ate1port5},
		{Src: t10dut1port6, Dst: t10switch1port1},
		{Src: t10dut1port7, Dst: t10switch1port2},
		{Src: t10dut1port8, Dst: t10switch1port3},
		{Src: t10dut1port9, Dst: t10switch1port4},
		{Src: t10dut1port10, Dst: t10switch1port5},
		{Src: t10switch1port6, Dst: t10ate1port6},
		{Src: t10switch1port7, Dst: t10ate1port7},
		{Src: t10switch1port8, Dst: t10ate1port8},
		{Src: t10switch1port9, Dst: t10ate1port9},
		{Src: t10switch1port10, Dst: t10ate1port10},
		{Src: t10switch1port11, Dst: t10switch2port21},

		{Src: t10dut2port1, Dst: t10ate2port1},
		{Src: t10dut2port2, Dst: t10ate2port2},
		{Src: t10dut2port3, Dst: t10ate2port3},
		{Src: t10dut2port4, Dst: t10ate2port4},
		{Src: t10dut2port5, Dst: t10ate2port5},
		{Src: t10dut2port6, Dst: t10switch1port12},
		{Src: t10dut2port7, Dst: t10switch1port13},
		{Src: t10dut2port8, Dst: t10switch1port14},
		{Src: t10dut2port9, Dst: t10switch1port15},
		{Src: t10dut2port10, Dst: t10switch1port16},
		{Src: t10switch1port17, Dst: t10ate2port6},
		{Src: t10switch1port18, Dst: t10ate2port7},
		{Src: t10switch1port19, Dst: t10ate2port8},
		{Src: t10switch1port20, Dst: t10ate2port9},
		{Src: t10switch1port21, Dst: t10ate2port10},

		{Src: t10dut3port1, Dst: t10switch2port1},
		{Src: t10dut3port2, Dst: t10switch2port2},
		{Src: t10dut3port3, Dst: t10switch2port3},
		{Src: t10dut3port4, Dst: t10switch2port4},
		{Src: t10dut3port5, Dst: t10switch2port5},
		{Src: t10dut3port6, Dst: t10switch3port1},
		{Src: t10dut3port7, Dst: t10switch3port2},
		{Src: t10dut3port8, Dst: t10switch3port3},
		{Src: t10dut3port9, Dst: t10switch3port4},
		{Src: t10dut3port10, Dst: t10switch3port5},
		{Src: t10switch2port6, Dst: t10ate3port1},
		{Src: t10switch2port7, Dst: t10ate3port2},
		{Src: t10switch2port8, Dst: t10ate3port3},
		{Src: t10switch2port9, Dst: t10ate3port4},
		{Src: t10switch2port10, Dst: t10ate3port5},
		{Src: t10switch3port6, Dst: t10ate3port6},
		{Src: t10switch3port7, Dst: t10ate3port7},
		{Src: t10switch3port8, Dst: t10ate3port8},
		{Src: t10switch3port9, Dst: t10ate3port9},
		{Src: t10switch3port10, Dst: t10ate3port10},
		{Src: t10switch2port22, Dst: t10switch3port21},

		{Src: t10dut4port1, Dst: t10switch2port11},
		{Src: t10dut4port2, Dst: t10switch2port12},
		{Src: t10dut4port3, Dst: t10switch2port13},
		{Src: t10dut4port4, Dst: t10switch2port14},
		{Src: t10dut4port5, Dst: t10switch2port15},
		{Src: t10dut4port6, Dst: t10switch3port11},
		{Src: t10dut4port7, Dst: t10switch3port12},
		{Src: t10dut4port8, Dst: t10switch3port13},
		{Src: t10dut4port9, Dst: t10switch3port14},
		{Src: t10dut4port10, Dst: t10switch3port15},
		{Src: t10switch2port16, Dst: t10ate4port1},
		{Src: t10switch2port17, Dst: t10ate4port2},
		{Src: t10switch2port18, Dst: t10ate4port3},
		{Src: t10switch2port19, Dst: t10ate4port4},
		{Src: t10switch2port20, Dst: t10ate4port5},
		{Src: t10switch3port16, Dst: t10ate4port6},
		{Src: t10switch3port17, Dst: t10ate4port7},
		{Src: t10switch3port18, Dst: t10ate4port8},
		{Src: t10switch3port19, Dst: t10ate4port9},
		{Src: t10switch3port20, Dst: t10ate4port10},
	},
}

var (
	abst10dut1port1  = &AbstractPort{Desc: "absdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port2  = &AbstractPort{Desc: "absdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port3  = &AbstractPort{Desc: "absdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port4  = &AbstractPort{Desc: "absdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port5  = &AbstractPort{Desc: "absdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port6  = &AbstractPort{Desc: "absdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port7  = &AbstractPort{Desc: "absdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port8  = &AbstractPort{Desc: "absdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port9  = &AbstractPort{Desc: "absdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1port10 = &AbstractPort{Desc: "absdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut1       = &AbstractNode{Desc: "absdut1", Ports: []*AbstractPort{abst10dut1port1, abst10dut1port2, abst10dut1port3, abst10dut1port4, abst10dut1port5, abst10dut1port6, abst10dut1port7, abst10dut1port8, abst10dut1port9, abst10dut1port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst10dut2port1  = &AbstractPort{Desc: "absdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port2  = &AbstractPort{Desc: "absdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port3  = &AbstractPort{Desc: "absdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port4  = &AbstractPort{Desc: "absdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port5  = &AbstractPort{Desc: "absdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port6  = &AbstractPort{Desc: "absdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port7  = &AbstractPort{Desc: "absdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port8  = &AbstractPort{Desc: "absdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port9  = &AbstractPort{Desc: "absdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2port10 = &AbstractPort{Desc: "absdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut2       = &AbstractNode{Desc: "absdut2", Ports: []*AbstractPort{abst10dut2port1, abst10dut2port2, abst10dut2port3, abst10dut2port4, abst10dut2port5, abst10dut2port6, abst10dut2port7, abst10dut2port8, abst10dut2port9, abst10dut2port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst10dut3port1  = &AbstractPort{Desc: "absdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port2  = &AbstractPort{Desc: "absdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port3  = &AbstractPort{Desc: "absdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port4  = &AbstractPort{Desc: "absdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port5  = &AbstractPort{Desc: "absdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port6  = &AbstractPort{Desc: "absdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port7  = &AbstractPort{Desc: "absdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port8  = &AbstractPort{Desc: "absdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port9  = &AbstractPort{Desc: "absdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3port10 = &AbstractPort{Desc: "absdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut3       = &AbstractNode{Desc: "absdut3", Ports: []*AbstractPort{abst10dut3port1, abst10dut3port2, abst10dut3port3, abst10dut3port4, abst10dut3port5, abst10dut3port6, abst10dut3port7, abst10dut3port8, abst10dut3port9, abst10dut3port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst10dut4port1  = &AbstractPort{Desc: "absdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port2  = &AbstractPort{Desc: "absdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port3  = &AbstractPort{Desc: "absdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port4  = &AbstractPort{Desc: "absdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port5  = &AbstractPort{Desc: "absdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port6  = &AbstractPort{Desc: "absdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port7  = &AbstractPort{Desc: "absdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port8  = &AbstractPort{Desc: "absdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port9  = &AbstractPort{Desc: "absdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4port10 = &AbstractPort{Desc: "absdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10dut4       = &AbstractNode{Desc: "absdut4", Ports: []*AbstractPort{abst10dut4port1, abst10dut4port2, abst10dut4port3, abst10dut4port4, abst10dut4port5, abst10dut4port6, abst10dut4port7, abst10dut4port8, abst10dut4port9, abst10dut4port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst10ate1port1  = &AbstractPort{Desc: "absate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port2  = &AbstractPort{Desc: "absate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port3  = &AbstractPort{Desc: "absate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port4  = &AbstractPort{Desc: "absate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port5  = &AbstractPort{Desc: "absate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port6  = &AbstractPort{Desc: "absate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port7  = &AbstractPort{Desc: "absate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port8  = &AbstractPort{Desc: "absate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port9  = &AbstractPort{Desc: "absate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1port10 = &AbstractPort{Desc: "absate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate1       = &AbstractNode{Desc: "absate1", Ports: []*AbstractPort{abst10ate1port1, abst10ate1port2, abst10ate1port3, abst10ate1port4, abst10ate1port5, abst10ate1port6, abst10ate1port7, abst10ate1port8, abst10ate1port9, abst10ate1port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst10ate2port1  = &AbstractPort{Desc: "absate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port2  = &AbstractPort{Desc: "absate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port3  = &AbstractPort{Desc: "absate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port4  = &AbstractPort{Desc: "absate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port5  = &AbstractPort{Desc: "absate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port6  = &AbstractPort{Desc: "absate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port7  = &AbstractPort{Desc: "absate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port8  = &AbstractPort{Desc: "absate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port9  = &AbstractPort{Desc: "absate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2port10 = &AbstractPort{Desc: "absate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate2       = &AbstractNode{Desc: "absate2", Ports: []*AbstractPort{abst10ate2port1, abst10ate2port2, abst10ate2port3, abst10ate2port4, abst10ate2port5, abst10ate2port6, abst10ate2port7, abst10ate2port8, abst10ate2port9, abst10ate2port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst10ate3port1  = &AbstractPort{Desc: "absate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port2  = &AbstractPort{Desc: "absate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port3  = &AbstractPort{Desc: "absate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port4  = &AbstractPort{Desc: "absate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port5  = &AbstractPort{Desc: "absate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port6  = &AbstractPort{Desc: "absate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port7  = &AbstractPort{Desc: "absate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port8  = &AbstractPort{Desc: "absate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port9  = &AbstractPort{Desc: "absate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3port10 = &AbstractPort{Desc: "absate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate3       = &AbstractNode{Desc: "absate3", Ports: []*AbstractPort{abst10ate3port1, abst10ate3port2, abst10ate3port3, abst10ate3port4, abst10ate3port5, abst10ate3port6, abst10ate3port7, abst10ate3port8, abst10ate3port9, abst10ate3port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst10ate4port1  = &AbstractPort{Desc: "absate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port2  = &AbstractPort{Desc: "absate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port3  = &AbstractPort{Desc: "absate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port4  = &AbstractPort{Desc: "absate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port5  = &AbstractPort{Desc: "absate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port6  = &AbstractPort{Desc: "absate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port7  = &AbstractPort{Desc: "absate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port8  = &AbstractPort{Desc: "absate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port9  = &AbstractPort{Desc: "absate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4port10 = &AbstractPort{Desc: "absate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst10ate4       = &AbstractNode{Desc: "absate4", Ports: []*AbstractPort{abst10ate4port1, abst10ate4port2, abst10ate4port3, abst10ate4port4, abst10ate4port5, abst10ate4port6, abst10ate4port7, abst10ate4port8, abst10ate4port9, abst10ate4port10}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
)

var abstractGraph10 = &AbstractGraph{
	Nodes: []*AbstractNode{abst10dut1,
		abst10ate1,
		abst10dut2,
		abst10ate2,
		abst10dut3,
		abst10ate3,
		abst10dut4,
		abst10ate4},
	Edges: []*AbstractEdge{{abst10dut1port1, abst10ate1port1},
		{abst10dut1port2, abst10ate1port2},
		{abst10dut1port3, abst10ate1port3},
		{abst10dut1port4, abst10ate1port4},
		{abst10dut1port5, abst10ate1port5},
		{abst10dut1port6, abst10ate1port6},
		{abst10dut1port7, abst10ate1port7},
		{abst10dut1port8, abst10ate1port8},
		{abst10dut1port9, abst10ate1port9},
		{abst10dut1port10, abst10ate1port10},
		{abst10dut2port1, abst10ate2port1},
		{abst10dut2port2, abst10ate2port2},
		{abst10dut2port3, abst10ate2port3},
		{abst10dut2port4, abst10ate2port4},
		{abst10dut2port5, abst10ate2port5},
		{abst10dut2port6, abst10ate2port6},
		{abst10dut2port7, abst10ate2port7},
		{abst10dut2port8, abst10ate2port8},
		{abst10dut2port9, abst10ate2port9},
		{abst10dut2port10, abst10ate2port10},
		{abst10dut3port1, abst10ate3port1},
		{abst10dut3port2, abst10ate3port2},
		{abst10dut3port3, abst10ate3port3},
		{abst10dut3port4, abst10ate3port4},
		{abst10dut3port5, abst10ate3port5},
		{abst10dut3port6, abst10ate3port6},
		{abst10dut3port7, abst10ate3port7},
		{abst10dut3port8, abst10ate3port8},
		{abst10dut3port9, abst10ate3port9},
		{abst10dut3port10, abst10ate3port10},

		{abst10dut4port1, abst10ate4port1},
		{abst10dut4port2, abst10ate4port2},
		{abst10dut4port3, abst10ate4port3},
		{abst10dut4port4, abst10ate4port4},
		{abst10dut4port5, abst10ate4port5},
		{abst10dut4port6, abst10ate4port6},
		{abst10dut4port7, abst10ate4port7},
		{abst10dut4port8, abst10ate4port8},
		{abst10dut4port9, abst10ate4port9},
		{abst10dut4port10, abst10ate4port10}},
}

func TestSolveTestHybrid10(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraph10, superGraphTest10)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 8 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 80 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}