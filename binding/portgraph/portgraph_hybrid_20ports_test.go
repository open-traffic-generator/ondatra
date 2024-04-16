// // 4 duts, 4 ates, 20 ports and 3 switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	t20dut1port1  = &ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port2  = &ConcretePort{Desc: "dut1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port3  = &ConcretePort{Desc: "dut1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port4  = &ConcretePort{Desc: "dut1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port5  = &ConcretePort{Desc: "dut1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port6  = &ConcretePort{Desc: "dut1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port7  = &ConcretePort{Desc: "dut1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port8  = &ConcretePort{Desc: "dut1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port9  = &ConcretePort{Desc: "dut1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port10 = &ConcretePort{Desc: "dut1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port11 = &ConcretePort{Desc: "dut1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port12 = &ConcretePort{Desc: "dut1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port13 = &ConcretePort{Desc: "dut1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port14 = &ConcretePort{Desc: "dut1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port15 = &ConcretePort{Desc: "dut1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port16 = &ConcretePort{Desc: "dut1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port17 = &ConcretePort{Desc: "dut1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port18 = &ConcretePort{Desc: "dut1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port19 = &ConcretePort{Desc: "dut1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1port20 = &ConcretePort{Desc: "dut1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut1       = &ConcreteNode{Desc: "dut1", Ports: []*ConcretePort{t20dut1port1, t20dut1port2, t20dut1port3, t20dut1port4, t20dut1port5, t20dut1port6, t20dut1port7, t20dut1port8, t20dut1port9, t20dut1port10, t20dut1port11, t20dut1port12, t20dut1port13, t20dut1port14, t20dut1port15, t20dut1port16, t20dut1port17, t20dut1port18, t20dut1port19, t20dut1port20}, Attrs: map[string]string{"vendor": "CISCO"}}

	t20dut2port1  = &ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port2  = &ConcretePort{Desc: "dut2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port3  = &ConcretePort{Desc: "dut2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port4  = &ConcretePort{Desc: "dut2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port5  = &ConcretePort{Desc: "dut2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port6  = &ConcretePort{Desc: "dut2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port7  = &ConcretePort{Desc: "dut2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port8  = &ConcretePort{Desc: "dut2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port9  = &ConcretePort{Desc: "dut2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port10 = &ConcretePort{Desc: "dut2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port11 = &ConcretePort{Desc: "dut2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port12 = &ConcretePort{Desc: "dut2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port13 = &ConcretePort{Desc: "dut2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port14 = &ConcretePort{Desc: "dut2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port15 = &ConcretePort{Desc: "dut2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port16 = &ConcretePort{Desc: "dut2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port17 = &ConcretePort{Desc: "dut2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port18 = &ConcretePort{Desc: "dut2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port19 = &ConcretePort{Desc: "dut2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2port20 = &ConcretePort{Desc: "dut2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut2       = &ConcreteNode{Desc: "dut2", Ports: []*ConcretePort{t20dut2port1, t20dut2port2, t20dut2port3, t20dut2port4, t20dut2port5, t20dut2port6, t20dut2port7, t20dut2port8, t20dut2port9, t20dut2port10, t20dut2port11, t20dut2port12, t20dut2port13, t20dut2port14, t20dut2port15, t20dut2port16, t20dut2port17, t20dut2port18, t20dut2port19, t20dut2port20}, Attrs: map[string]string{"vendor": "CISCO"}}

	t20dut3port1  = &ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port2  = &ConcretePort{Desc: "dut3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port3  = &ConcretePort{Desc: "dut3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port4  = &ConcretePort{Desc: "dut3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port5  = &ConcretePort{Desc: "dut3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port6  = &ConcretePort{Desc: "dut3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port7  = &ConcretePort{Desc: "dut3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port8  = &ConcretePort{Desc: "dut3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port9  = &ConcretePort{Desc: "dut3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port10 = &ConcretePort{Desc: "dut3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port11 = &ConcretePort{Desc: "dut3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port12 = &ConcretePort{Desc: "dut3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port13 = &ConcretePort{Desc: "dut3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port14 = &ConcretePort{Desc: "dut3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port15 = &ConcretePort{Desc: "dut3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port16 = &ConcretePort{Desc: "dut3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port17 = &ConcretePort{Desc: "dut3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port18 = &ConcretePort{Desc: "dut3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port19 = &ConcretePort{Desc: "dut3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3port20 = &ConcretePort{Desc: "dut3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t20dut3port1, t20dut3port2, t20dut3port3, t20dut3port4, t20dut3port5, t20dut3port6, t20dut3port7, t20dut3port8, t20dut3port9, t20dut3port10, t20dut3port11, t20dut3port12, t20dut3port13, t20dut3port14, t20dut3port15, t20dut3port16, t20dut3port17, t20dut3port18, t20dut3port19, t20dut3port20}, Attrs: map[string]string{"vendor": "CISCO"}}
	// dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t20dut3port1}, Attrs: map[string]string{"vendor": "CISCO"}}
	t20dut4port1  = &ConcretePort{Desc: "dut4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port2  = &ConcretePort{Desc: "dut4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port3  = &ConcretePort{Desc: "dut4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port4  = &ConcretePort{Desc: "dut4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port5  = &ConcretePort{Desc: "dut4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port6  = &ConcretePort{Desc: "dut4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port7  = &ConcretePort{Desc: "dut4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port8  = &ConcretePort{Desc: "dut4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port9  = &ConcretePort{Desc: "dut4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port10 = &ConcretePort{Desc: "dut4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port11 = &ConcretePort{Desc: "dut4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port12 = &ConcretePort{Desc: "dut4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port13 = &ConcretePort{Desc: "dut4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port14 = &ConcretePort{Desc: "dut4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port15 = &ConcretePort{Desc: "dut4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port16 = &ConcretePort{Desc: "dut4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port17 = &ConcretePort{Desc: "dut4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port18 = &ConcretePort{Desc: "dut4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port19 = &ConcretePort{Desc: "dut4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4port20 = &ConcretePort{Desc: "dut4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20dut4       = &ConcreteNode{Desc: "dut4", Ports: []*ConcretePort{t20dut4port1, t20dut4port2, t20dut4port3, t20dut4port4, t20dut4port5, t20dut4port6, t20dut4port7, t20dut4port8, t20dut4port9, t20dut4port10, t20dut4port11, t20dut4port12, t20dut4port13, t20dut4port14, t20dut4port15, t20dut4port16, t20dut4port17, t20dut4port18, t20dut4port19, t20dut4port20}, Attrs: map[string]string{"vendor": "CISCO"}}
	t20ate1port1  = &ConcretePort{Desc: "ate1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port2  = &ConcretePort{Desc: "ate1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port3  = &ConcretePort{Desc: "ate1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port4  = &ConcretePort{Desc: "ate1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port5  = &ConcretePort{Desc: "ate1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port6  = &ConcretePort{Desc: "ate1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port7  = &ConcretePort{Desc: "ate1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port8  = &ConcretePort{Desc: "ate1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port9  = &ConcretePort{Desc: "ate1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port10 = &ConcretePort{Desc: "ate1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port11 = &ConcretePort{Desc: "ate1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port12 = &ConcretePort{Desc: "ate1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port13 = &ConcretePort{Desc: "ate1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port14 = &ConcretePort{Desc: "ate1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port15 = &ConcretePort{Desc: "ate1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port16 = &ConcretePort{Desc: "ate1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port17 = &ConcretePort{Desc: "ate1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port18 = &ConcretePort{Desc: "ate1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port19 = &ConcretePort{Desc: "ate1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1port20 = &ConcretePort{Desc: "ate1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate1       = &ConcreteNode{Desc: "ate1", Ports: []*ConcretePort{t20ate1port1, t20ate1port2, t20ate1port3, t20ate1port4, t20ate1port5, t20ate1port6, t20ate1port7, t20ate1port8, t20ate1port9, t20ate1port10, t20ate1port11, t20ate1port12, t20ate1port13, t20ate1port14, t20ate1port15, t20ate1port16, t20ate1port17, t20ate1port18, t20ate1port19, t20ate1port20}, Attrs: map[string]string{"vendor": "TGEN"}}

	t20ate2port1  = &ConcretePort{Desc: "ate2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port2  = &ConcretePort{Desc: "ate2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port3  = &ConcretePort{Desc: "ate2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port4  = &ConcretePort{Desc: "ate2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port5  = &ConcretePort{Desc: "ate2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port6  = &ConcretePort{Desc: "ate2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port7  = &ConcretePort{Desc: "ate2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port8  = &ConcretePort{Desc: "ate2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port9  = &ConcretePort{Desc: "ate2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port10 = &ConcretePort{Desc: "ate2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port11 = &ConcretePort{Desc: "ate2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port12 = &ConcretePort{Desc: "ate2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port13 = &ConcretePort{Desc: "ate2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port14 = &ConcretePort{Desc: "ate2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port15 = &ConcretePort{Desc: "ate2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port16 = &ConcretePort{Desc: "ate2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port17 = &ConcretePort{Desc: "ate2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port18 = &ConcretePort{Desc: "ate2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port19 = &ConcretePort{Desc: "ate2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2port20 = &ConcretePort{Desc: "ate2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate2       = &ConcreteNode{Desc: "ate2", Ports: []*ConcretePort{t20ate2port1, t20ate2port2, t20ate2port3, t20ate2port4, t20ate2port5, t20ate2port6, t20ate2port7, t20ate2port8, t20ate2port9, t20ate2port10, t20ate2port11, t20ate2port12, t20ate2port13, t20ate2port14, t20ate2port15, t20ate2port16, t20ate2port17, t20ate2port18, t20ate2port19, t20ate2port20}, Attrs: map[string]string{"vendor": "TGEN"}}

	t20ate3port1  = &ConcretePort{Desc: "ate3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port2  = &ConcretePort{Desc: "ate3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port3  = &ConcretePort{Desc: "ate3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port4  = &ConcretePort{Desc: "ate3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port5  = &ConcretePort{Desc: "ate3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port6  = &ConcretePort{Desc: "ate3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port7  = &ConcretePort{Desc: "ate3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port8  = &ConcretePort{Desc: "ate3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port9  = &ConcretePort{Desc: "ate3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port10 = &ConcretePort{Desc: "ate3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port11 = &ConcretePort{Desc: "ate3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port12 = &ConcretePort{Desc: "ate3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port13 = &ConcretePort{Desc: "ate3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port14 = &ConcretePort{Desc: "ate3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port15 = &ConcretePort{Desc: "ate3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port16 = &ConcretePort{Desc: "ate3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port17 = &ConcretePort{Desc: "ate3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port18 = &ConcretePort{Desc: "ate3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port19 = &ConcretePort{Desc: "ate3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3port20 = &ConcretePort{Desc: "ate3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t20ate3port1, t20ate3port2, t20ate3port3, t20ate3port4, t20ate3port5, t20ate3port6, t20ate3port7, t20ate3port8, t20ate3port9, t20ate3port10, t20ate3port11, t20ate3port12, t20ate3port13, t20ate3port14, t20ate3port15, t20ate3port16, t20ate3port17, t20ate3port18, t20ate3port19, t20ate3port20}, Attrs: map[string]string{"vendor": "TGEN"}}
	// ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t20ate3port1}, Attrs: map[string]string{"vendor": "TGEN"}}
	t20ate4port1  = &ConcretePort{Desc: "ate4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port2  = &ConcretePort{Desc: "ate4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port3  = &ConcretePort{Desc: "ate4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port4  = &ConcretePort{Desc: "ate4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port5  = &ConcretePort{Desc: "ate4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port6  = &ConcretePort{Desc: "ate4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port7  = &ConcretePort{Desc: "ate4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port8  = &ConcretePort{Desc: "ate4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port9  = &ConcretePort{Desc: "ate4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port10 = &ConcretePort{Desc: "ate4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port11 = &ConcretePort{Desc: "ate4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port12 = &ConcretePort{Desc: "ate4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port13 = &ConcretePort{Desc: "ate4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port14 = &ConcretePort{Desc: "ate4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port15 = &ConcretePort{Desc: "ate4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port16 = &ConcretePort{Desc: "ate4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port17 = &ConcretePort{Desc: "ate4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port18 = &ConcretePort{Desc: "ate4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port19 = &ConcretePort{Desc: "ate4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4port20 = &ConcretePort{Desc: "ate4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20ate4       = &ConcreteNode{Desc: "ate4", Ports: []*ConcretePort{t20ate4port1, t20ate4port2, t20ate4port3, t20ate4port4, t20ate4port5, t20ate4port6, t20ate4port7, t20ate4port8, t20ate4port9, t20ate4port10, t20ate4port11, t20ate4port12, t20ate4port13, t20ate4port14, t20ate4port15, t20ate4port16, t20ate4port17, t20ate4port18, t20ate4port19, t20ate4port20}, Attrs: map[string]string{"vendor": "TGEN"}}

	t20switch1port1  = &ConcretePort{Desc: "switch1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port2  = &ConcretePort{Desc: "switch1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port3  = &ConcretePort{Desc: "switch1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port4  = &ConcretePort{Desc: "switch1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port5  = &ConcretePort{Desc: "switch1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port6  = &ConcretePort{Desc: "switch1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port7  = &ConcretePort{Desc: "switch1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port8  = &ConcretePort{Desc: "switch1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port9  = &ConcretePort{Desc: "switch1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port10 = &ConcretePort{Desc: "switch1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port11 = &ConcretePort{Desc: "switch1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port12 = &ConcretePort{Desc: "switch1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port13 = &ConcretePort{Desc: "switch1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port14 = &ConcretePort{Desc: "switch1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port15 = &ConcretePort{Desc: "switch1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port16 = &ConcretePort{Desc: "switch1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port17 = &ConcretePort{Desc: "switch1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port18 = &ConcretePort{Desc: "switch1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port19 = &ConcretePort{Desc: "switch1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port20 = &ConcretePort{Desc: "switch1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port21 = &ConcretePort{Desc: "switch1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port22 = &ConcretePort{Desc: "switch1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port23 = &ConcretePort{Desc: "switch1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port24 = &ConcretePort{Desc: "switch1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port25 = &ConcretePort{Desc: "switch1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port26 = &ConcretePort{Desc: "switch1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port27 = &ConcretePort{Desc: "switch1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port28 = &ConcretePort{Desc: "switch1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port29 = &ConcretePort{Desc: "switch1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port30 = &ConcretePort{Desc: "switch1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port31 = &ConcretePort{Desc: "switch1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port32 = &ConcretePort{Desc: "switch1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port33 = &ConcretePort{Desc: "switch1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port34 = &ConcretePort{Desc: "switch1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port35 = &ConcretePort{Desc: "switch1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port36 = &ConcretePort{Desc: "switch1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port37 = &ConcretePort{Desc: "switch1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port38 = &ConcretePort{Desc: "switch1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port39 = &ConcretePort{Desc: "switch1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port40 = &ConcretePort{Desc: "switch1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch1port41 = &ConcretePort{Desc: "switch1:port41", Attrs: map[string]string{"speed": "S_400G"}}

	t20switch1 = &ConcreteNode{Desc: "switch1", Ports: []*ConcretePort{t20switch1port1, t20switch1port2, t20switch1port3, t20switch1port4, t20switch1port5, t20switch1port6, t20switch1port7, t20switch1port8, t20switch1port9, t20switch1port10, t20switch1port11, t20switch1port12, t20switch1port13, t20switch1port14, t20switch1port15, t20switch1port16, t20switch1port17, t20switch1port18, t20switch1port19, t20switch1port20, t20switch1port21, t20switch1port22, t20switch1port23, t20switch1port24, t20switch1port25, t20switch1port26, t20switch1port27, t20switch1port28, t20switch1port29, t20switch1port30, t20switch1port31, t20switch1port32, t20switch1port33, t20switch1port34, t20switch1port35, t20switch1port36, t20switch1port37, t20switch1port38, t20switch1port39, t20switch1port40, t20switch1port41}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	t20switch2port1  = &ConcretePort{Desc: "switch2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port2  = &ConcretePort{Desc: "switch2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port3  = &ConcretePort{Desc: "switch2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port4  = &ConcretePort{Desc: "switch2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port5  = &ConcretePort{Desc: "switch2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port6  = &ConcretePort{Desc: "switch2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port7  = &ConcretePort{Desc: "switch2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port8  = &ConcretePort{Desc: "switch2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port9  = &ConcretePort{Desc: "switch2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port10 = &ConcretePort{Desc: "switch2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port11 = &ConcretePort{Desc: "switch2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port12 = &ConcretePort{Desc: "switch2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port13 = &ConcretePort{Desc: "switch2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port14 = &ConcretePort{Desc: "switch2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port15 = &ConcretePort{Desc: "switch2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port16 = &ConcretePort{Desc: "switch2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port17 = &ConcretePort{Desc: "switch2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port18 = &ConcretePort{Desc: "switch2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port19 = &ConcretePort{Desc: "switch2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port20 = &ConcretePort{Desc: "switch2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port21 = &ConcretePort{Desc: "switch2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port22 = &ConcretePort{Desc: "switch2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port23 = &ConcretePort{Desc: "switch2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port24 = &ConcretePort{Desc: "switch2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port25 = &ConcretePort{Desc: "switch2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port26 = &ConcretePort{Desc: "switch2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port27 = &ConcretePort{Desc: "switch2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port28 = &ConcretePort{Desc: "switch2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port29 = &ConcretePort{Desc: "switch2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port30 = &ConcretePort{Desc: "switch2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port31 = &ConcretePort{Desc: "switch2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port32 = &ConcretePort{Desc: "switch2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port33 = &ConcretePort{Desc: "switch2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port34 = &ConcretePort{Desc: "switch2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port35 = &ConcretePort{Desc: "switch2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port36 = &ConcretePort{Desc: "switch2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port37 = &ConcretePort{Desc: "switch2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port38 = &ConcretePort{Desc: "switch2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port39 = &ConcretePort{Desc: "switch2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port40 = &ConcretePort{Desc: "switch2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port41 = &ConcretePort{Desc: "switch2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2port42 = &ConcretePort{Desc: "switch2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch2       = &ConcreteNode{Desc: "switch2", Ports: []*ConcretePort{t20switch2port1, t20switch2port2, t20switch2port3, t20switch2port4, t20switch2port5, t20switch2port6, t20switch2port7, t20switch2port8, t20switch2port9, t20switch2port10, t20switch2port11, t20switch2port12, t20switch2port13, t20switch2port14, t20switch2port15, t20switch2port16, t20switch2port17, t20switch2port18, t20switch2port19, t20switch2port20, t20switch2port21, t20switch2port22, t20switch2port23, t20switch2port24, t20switch2port25, t20switch2port26, t20switch2port27, t20switch2port28, t20switch2port29, t20switch2port30, t20switch2port31, t20switch2port32, t20switch2port33, t20switch2port34, t20switch2port35, t20switch2port36, t20switch2port37, t20switch2port38, t20switch2port39, t20switch2port40, t20switch2port41, t20switch2port42}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	t20switch3port1  = &ConcretePort{Desc: "switch3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port2  = &ConcretePort{Desc: "switch3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port3  = &ConcretePort{Desc: "switch3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port4  = &ConcretePort{Desc: "switch3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port5  = &ConcretePort{Desc: "switch3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port6  = &ConcretePort{Desc: "switch3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port7  = &ConcretePort{Desc: "switch3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port8  = &ConcretePort{Desc: "switch3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port9  = &ConcretePort{Desc: "switch3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port10 = &ConcretePort{Desc: "switch3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port11 = &ConcretePort{Desc: "switch3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port12 = &ConcretePort{Desc: "switch3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port13 = &ConcretePort{Desc: "switch3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port14 = &ConcretePort{Desc: "switch3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port15 = &ConcretePort{Desc: "switch3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port16 = &ConcretePort{Desc: "switch3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port17 = &ConcretePort{Desc: "switch3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port18 = &ConcretePort{Desc: "switch3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port19 = &ConcretePort{Desc: "switch3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port20 = &ConcretePort{Desc: "switch3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port21 = &ConcretePort{Desc: "switch3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port22 = &ConcretePort{Desc: "switch3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port23 = &ConcretePort{Desc: "switch3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port24 = &ConcretePort{Desc: "switch3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port25 = &ConcretePort{Desc: "switch3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port26 = &ConcretePort{Desc: "switch3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port27 = &ConcretePort{Desc: "switch3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port28 = &ConcretePort{Desc: "switch3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port29 = &ConcretePort{Desc: "switch3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port30 = &ConcretePort{Desc: "switch3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port31 = &ConcretePort{Desc: "switch3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port32 = &ConcretePort{Desc: "switch3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port33 = &ConcretePort{Desc: "switch3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port34 = &ConcretePort{Desc: "switch3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port35 = &ConcretePort{Desc: "switch3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port36 = &ConcretePort{Desc: "switch3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port37 = &ConcretePort{Desc: "switch3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port38 = &ConcretePort{Desc: "switch3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port39 = &ConcretePort{Desc: "switch3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port40 = &ConcretePort{Desc: "switch3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t20switch3port41 = &ConcretePort{Desc: "switch3:port41", Attrs: map[string]string{"speed": "S_400G"}}

	t20switch3 = &ConcreteNode{Desc: "switch3", Ports: []*ConcretePort{t20switch3port1, t20switch3port2, t20switch3port3, t20switch3port4, t20switch3port5, t20switch3port6, t20switch3port7, t20switch3port8, t20switch3port9, t20switch3port10, t20switch3port11, t20switch3port12, t20switch3port13, t20switch3port14, t20switch3port15, t20switch3port16, t20switch3port17, t20switch3port18, t20switch3port19, t20switch3port20, t20switch3port21, t20switch3port22, t20switch3port23, t20switch3port24, t20switch3port25, t20switch3port26, t20switch3port27, t20switch3port28, t20switch3port29, t20switch3port30, t20switch3port31, t20switch3port32, t20switch3port33, t20switch3port34, t20switch3port35, t20switch3port36, t20switch3port37, t20switch3port38, t20switch3port39, t20switch3port40, t20switch3port41}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}
)

var superGraphTest20 = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		t20dut1,
		t20ate1,
		t20dut2,
		t20ate2,
		t20dut3,
		t20ate3,
		t20dut4,
		t20ate4,
		t20switch1, t20switch2, t20switch3},
	Edges: []*ConcreteEdge{
		{Src: t20dut1port1, Dst: t20ate1port1},
		{Src: t20dut1port2, Dst: t20ate1port2},
		{Src: t20dut1port3, Dst: t20ate1port3},
		{Src: t20dut1port4, Dst: t20ate1port4},
		{Src: t20dut1port5, Dst: t20ate1port5},
		{Src: t20dut1port6, Dst: t20ate1port6},
		{Src: t20dut1port7, Dst: t20ate1port7},
		{Src: t20dut1port8, Dst: t20ate1port8},
		{Src: t20dut1port9, Dst: t20ate1port9},
		{Src: t20dut1port10, Dst: t20ate1port10},
		{Src: t20dut1port11, Dst: t20switch1port1},
		{Src: t20dut1port12, Dst: t20switch1port2},
		{Src: t20dut1port13, Dst: t20switch1port3},
		{Src: t20dut1port14, Dst: t20switch1port4},
		{Src: t20dut1port15, Dst: t20switch1port5},
		{Src: t20dut1port16, Dst: t20switch1port6},
		{Src: t20dut1port17, Dst: t20switch1port7},
		{Src: t20dut1port18, Dst: t20switch1port8},
		{Src: t20dut1port19, Dst: t20switch1port9},
		{Src: t20dut1port20, Dst: t20switch1port10},
		{Src: t20switch1port11, Dst: t20ate1port11},
		{Src: t20switch1port12, Dst: t20ate1port12},
		{Src: t20switch1port13, Dst: t20ate1port13},
		{Src: t20switch1port14, Dst: t20ate1port14},
		{Src: t20switch1port15, Dst: t20ate1port15},
		{Src: t20switch1port16, Dst: t20ate1port16},
		{Src: t20switch1port17, Dst: t20ate1port17},
		{Src: t20switch1port18, Dst: t20ate1port18},
		{Src: t20switch1port19, Dst: t20ate1port19},
		{Src: t20switch1port20, Dst: t20ate1port20},
		{Src: t20switch1port41, Dst: t20switch2port41},

		{Src: t20dut2port1, Dst: t20ate2port1},
		{Src: t20dut2port2, Dst: t20ate2port2},
		{Src: t20dut2port3, Dst: t20ate2port3},
		{Src: t20dut2port4, Dst: t20ate2port4},
		{Src: t20dut2port5, Dst: t20ate2port5},
		{Src: t20dut2port6, Dst: t20ate2port6},
		{Src: t20dut2port7, Dst: t20ate2port7},
		{Src: t20dut2port8, Dst: t20ate2port8},
		{Src: t20dut2port9, Dst: t20ate2port9},
		{Src: t20dut2port10, Dst: t20ate2port10},
		{Src: t20dut2port11, Dst: t20switch1port21},
		{Src: t20dut2port12, Dst: t20switch1port22},
		{Src: t20dut2port13, Dst: t20switch1port23},
		{Src: t20dut2port14, Dst: t20switch1port24},
		{Src: t20dut2port15, Dst: t20switch1port25},
		{Src: t20dut2port16, Dst: t20switch1port26},
		{Src: t20dut2port17, Dst: t20switch1port27},
		{Src: t20dut2port18, Dst: t20switch1port28},
		{Src: t20dut2port19, Dst: t20switch1port29},
		{Src: t20dut2port20, Dst: t20switch1port30},
		{Src: t20switch1port31, Dst: t20ate2port11},
		{Src: t20switch1port32, Dst: t20ate2port12},
		{Src: t20switch1port33, Dst: t20ate2port13},
		{Src: t20switch1port34, Dst: t20ate2port14},
		{Src: t20switch1port35, Dst: t20ate2port15},
		{Src: t20switch1port36, Dst: t20ate2port16},
		{Src: t20switch1port37, Dst: t20ate2port17},
		{Src: t20switch1port38, Dst: t20ate2port18},
		{Src: t20switch1port39, Dst: t20ate2port19},
		{Src: t20switch1port40, Dst: t20ate2port20},

		{Src: t20dut3port1, Dst: t20switch2port1},
		{Src: t20dut3port2, Dst: t20switch2port2},
		{Src: t20dut3port3, Dst: t20switch2port3},
		{Src: t20dut3port4, Dst: t20switch2port4},
		{Src: t20dut3port5, Dst: t20switch2port5},
		{Src: t20dut3port6, Dst: t20switch2port6},
		{Src: t20dut3port7, Dst: t20switch2port7},
		{Src: t20dut3port8, Dst: t20switch2port8},
		{Src: t20dut3port9, Dst: t20switch2port9},
		{Src: t20dut3port10, Dst: t20switch2port10},
		{Src: t20switch2port11, Dst: t20ate3port1},
		{Src: t20switch2port12, Dst: t20ate3port2},
		{Src: t20switch2port13, Dst: t20ate3port3},
		{Src: t20switch2port14, Dst: t20ate3port4},
		{Src: t20switch2port15, Dst: t20ate3port5},
		{Src: t20switch2port16, Dst: t20ate3port6},
		{Src: t20switch2port17, Dst: t20ate3port7},
		{Src: t20switch2port18, Dst: t20ate3port8},
		{Src: t20switch2port19, Dst: t20ate3port9},
		{Src: t20switch2port20, Dst: t20ate3port10},
		{Src: t20dut4port11, Dst: t20switch2port21},
		{Src: t20dut4port12, Dst: t20switch2port22},
		{Src: t20dut4port13, Dst: t20switch2port23},
		{Src: t20dut4port14, Dst: t20switch2port24},
		{Src: t20dut4port15, Dst: t20switch2port25},
		{Src: t20dut4port16, Dst: t20switch2port26},
		{Src: t20dut4port17, Dst: t20switch2port27},
		{Src: t20dut4port18, Dst: t20switch2port28},
		{Src: t20dut4port19, Dst: t20switch2port29},
		{Src: t20dut4port20, Dst: t20switch2port30},
		{Src: t20switch2port31, Dst: t20ate4port11},
		{Src: t20switch2port32, Dst: t20ate4port12},
		{Src: t20switch2port33, Dst: t20ate4port13},
		{Src: t20switch2port34, Dst: t20ate4port14},
		{Src: t20switch2port35, Dst: t20ate4port15},
		{Src: t20switch2port36, Dst: t20ate4port16},
		{Src: t20switch2port37, Dst: t20ate4port17},
		{Src: t20switch2port38, Dst: t20ate4port18},
		{Src: t20switch2port39, Dst: t20ate4port19},
		{Src: t20switch2port40, Dst: t20ate4port20},
		{Src: t20switch2port42, Dst: t20switch3port41},

		{Src: t20dut3port11, Dst: t20switch3port1},
		{Src: t20dut3port12, Dst: t20switch3port2},
		{Src: t20dut3port13, Dst: t20switch3port3},
		{Src: t20dut3port14, Dst: t20switch3port4},
		{Src: t20dut3port15, Dst: t20switch3port5},
		{Src: t20dut3port16, Dst: t20switch3port6},
		{Src: t20dut3port17, Dst: t20switch3port7},
		{Src: t20dut3port18, Dst: t20switch3port8},
		{Src: t20dut3port19, Dst: t20switch3port9},
		{Src: t20dut3port20, Dst: t20switch3port10},
		{Src: t20switch3port11, Dst: t20ate3port11},
		{Src: t20switch3port12, Dst: t20ate3port12},
		{Src: t20switch3port13, Dst: t20ate3port13},
		{Src: t20switch3port14, Dst: t20ate3port14},
		{Src: t20switch3port15, Dst: t20ate3port15},
		{Src: t20switch3port16, Dst: t20ate3port16},
		{Src: t20switch3port17, Dst: t20ate3port17},
		{Src: t20switch3port18, Dst: t20ate3port18},
		{Src: t20switch3port19, Dst: t20ate3port19},
		{Src: t20switch3port20, Dst: t20ate3port20},
		{Src: t20dut4port1, Dst: t20switch3port21},
		{Src: t20dut4port2, Dst: t20switch3port22},
		{Src: t20dut4port3, Dst: t20switch3port23},
		{Src: t20dut4port4, Dst: t20switch3port24},
		{Src: t20dut4port5, Dst: t20switch3port25},
		{Src: t20dut4port6, Dst: t20switch3port26},
		{Src: t20dut4port7, Dst: t20switch3port27},
		{Src: t20dut4port8, Dst: t20switch3port28},
		{Src: t20dut4port9, Dst: t20switch3port29},
		{Src: t20dut4port10, Dst: t20switch3port30},
		{Src: t20switch3port31, Dst: t20ate4port1},
		{Src: t20switch3port32, Dst: t20ate4port2},
		{Src: t20switch3port33, Dst: t20ate4port3},
		{Src: t20switch3port34, Dst: t20ate4port4},
		{Src: t20switch3port35, Dst: t20ate4port5},
		{Src: t20switch3port36, Dst: t20ate4port6},
		{Src: t20switch3port37, Dst: t20ate4port7},
		{Src: t20switch3port38, Dst: t20ate4port8},
		{Src: t20switch3port39, Dst: t20ate4port9},
		{Src: t20switch3port40, Dst: t20ate4port10},
	},
}

var (
	abst20dut1port1  = &AbstractPort{Desc: "absdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port2  = &AbstractPort{Desc: "absdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port3  = &AbstractPort{Desc: "absdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port4  = &AbstractPort{Desc: "absdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port5  = &AbstractPort{Desc: "absdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port6  = &AbstractPort{Desc: "absdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port7  = &AbstractPort{Desc: "absdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port8  = &AbstractPort{Desc: "absdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port9  = &AbstractPort{Desc: "absdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port10 = &AbstractPort{Desc: "absdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port11 = &AbstractPort{Desc: "absdut1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port12 = &AbstractPort{Desc: "absdut1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port13 = &AbstractPort{Desc: "absdut1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port14 = &AbstractPort{Desc: "absdut1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port15 = &AbstractPort{Desc: "absdut1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port16 = &AbstractPort{Desc: "absdut1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port17 = &AbstractPort{Desc: "absdut1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port18 = &AbstractPort{Desc: "absdut1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port19 = &AbstractPort{Desc: "absdut1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1port20 = &AbstractPort{Desc: "absdut1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut1       = &AbstractNode{Desc: "absdut1", Ports: []*AbstractPort{abst20dut1port1, abst20dut1port2, abst20dut1port3, abst20dut1port4, abst20dut1port5, abst20dut1port6, abst20dut1port7, abst20dut1port8, abst20dut1port9, abst20dut1port10, abst20dut1port11, abst20dut1port12, abst20dut1port13, abst20dut1port14, abst20dut1port15, abst20dut1port16, abst20dut1port17, abst20dut1port18, abst20dut1port19, abst20dut1port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst20dut2port1  = &AbstractPort{Desc: "absdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port2  = &AbstractPort{Desc: "absdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port3  = &AbstractPort{Desc: "absdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port4  = &AbstractPort{Desc: "absdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port5  = &AbstractPort{Desc: "absdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port6  = &AbstractPort{Desc: "absdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port7  = &AbstractPort{Desc: "absdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port8  = &AbstractPort{Desc: "absdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port9  = &AbstractPort{Desc: "absdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port10 = &AbstractPort{Desc: "absdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port11 = &AbstractPort{Desc: "absdut2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port12 = &AbstractPort{Desc: "absdut2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port13 = &AbstractPort{Desc: "absdut2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port14 = &AbstractPort{Desc: "absdut2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port15 = &AbstractPort{Desc: "absdut2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port16 = &AbstractPort{Desc: "absdut2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port17 = &AbstractPort{Desc: "absdut2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port18 = &AbstractPort{Desc: "absdut2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port19 = &AbstractPort{Desc: "absdut2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2port20 = &AbstractPort{Desc: "absdut2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut2       = &AbstractNode{Desc: "absdut2", Ports: []*AbstractPort{abst20dut2port1, abst20dut2port2, abst20dut2port3, abst20dut2port4, abst20dut2port5, abst20dut2port6, abst20dut2port7, abst20dut2port8, abst20dut2port9, abst20dut2port10, abst20dut2port11, abst20dut2port12, abst20dut2port13, abst20dut2port14, abst20dut2port15, abst20dut2port16, abst20dut2port17, abst20dut2port18, abst20dut2port19, abst20dut2port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst20dut3port1  = &AbstractPort{Desc: "absdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port2  = &AbstractPort{Desc: "absdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port3  = &AbstractPort{Desc: "absdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port4  = &AbstractPort{Desc: "absdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port5  = &AbstractPort{Desc: "absdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port6  = &AbstractPort{Desc: "absdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port7  = &AbstractPort{Desc: "absdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port8  = &AbstractPort{Desc: "absdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port9  = &AbstractPort{Desc: "absdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port10 = &AbstractPort{Desc: "absdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port11 = &AbstractPort{Desc: "absdut3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port12 = &AbstractPort{Desc: "absdut3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port13 = &AbstractPort{Desc: "absdut3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port14 = &AbstractPort{Desc: "absdut3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port15 = &AbstractPort{Desc: "absdut3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port16 = &AbstractPort{Desc: "absdut3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port17 = &AbstractPort{Desc: "absdut3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port18 = &AbstractPort{Desc: "absdut3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port19 = &AbstractPort{Desc: "absdut3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3port20 = &AbstractPort{Desc: "absdut3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut3       = &AbstractNode{Desc: "absdut3", Ports: []*AbstractPort{abst20dut3port1, abst20dut3port2, abst20dut3port3, abst20dut3port4, abst20dut3port5, abst20dut3port6, abst20dut3port7, abst20dut3port8, abst20dut3port9, abst20dut3port10, abst20dut3port11, abst20dut3port12, abst20dut3port13, abst20dut3port14, abst20dut3port15, abst20dut3port16, abst20dut3port17, abst20dut3port18, abst20dut3port19, abst20dut3port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst20dut4port1  = &AbstractPort{Desc: "absdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port2  = &AbstractPort{Desc: "absdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port3  = &AbstractPort{Desc: "absdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port4  = &AbstractPort{Desc: "absdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port5  = &AbstractPort{Desc: "absdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port6  = &AbstractPort{Desc: "absdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port7  = &AbstractPort{Desc: "absdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port8  = &AbstractPort{Desc: "absdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port9  = &AbstractPort{Desc: "absdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port10 = &AbstractPort{Desc: "absdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port11 = &AbstractPort{Desc: "absdut4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port12 = &AbstractPort{Desc: "absdut4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port13 = &AbstractPort{Desc: "absdut4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port14 = &AbstractPort{Desc: "absdut4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port15 = &AbstractPort{Desc: "absdut4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port16 = &AbstractPort{Desc: "absdut4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port17 = &AbstractPort{Desc: "absdut4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port18 = &AbstractPort{Desc: "absdut4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port19 = &AbstractPort{Desc: "absdut4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4port20 = &AbstractPort{Desc: "absdut4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20dut4       = &AbstractNode{Desc: "absdut4", Ports: []*AbstractPort{abst20dut4port1, abst20dut4port2, abst20dut4port3, abst20dut4port4, abst20dut4port5, abst20dut4port6, abst20dut4port7, abst20dut4port8, abst20dut4port9, abst20dut4port10, abst20dut4port11, abst20dut4port12, abst20dut4port13, abst20dut4port14, abst20dut4port15, abst20dut4port16, abst20dut4port17, abst20dut4port18, abst20dut4port19, abst20dut4port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst20ate1port1  = &AbstractPort{Desc: "absate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port2  = &AbstractPort{Desc: "absate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port3  = &AbstractPort{Desc: "absate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port4  = &AbstractPort{Desc: "absate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port5  = &AbstractPort{Desc: "absate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port6  = &AbstractPort{Desc: "absate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port7  = &AbstractPort{Desc: "absate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port8  = &AbstractPort{Desc: "absate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port9  = &AbstractPort{Desc: "absate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port10 = &AbstractPort{Desc: "absate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port11 = &AbstractPort{Desc: "absate1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port12 = &AbstractPort{Desc: "absate1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port13 = &AbstractPort{Desc: "absate1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port14 = &AbstractPort{Desc: "absate1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port15 = &AbstractPort{Desc: "absate1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port16 = &AbstractPort{Desc: "absate1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port17 = &AbstractPort{Desc: "absate1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port18 = &AbstractPort{Desc: "absate1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port19 = &AbstractPort{Desc: "absate1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1port20 = &AbstractPort{Desc: "absate1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate1       = &AbstractNode{Desc: "absate1", Ports: []*AbstractPort{abst20ate1port1, abst20ate1port2, abst20ate1port3, abst20ate1port4, abst20ate1port5, abst20ate1port6, abst20ate1port7, abst20ate1port8, abst20ate1port9, abst20ate1port10, abst20ate1port11, abst20ate1port12, abst20ate1port13, abst20ate1port14, abst20ate1port15, abst20ate1port16, abst20ate1port17, abst20ate1port18, abst20ate1port19, abst20ate1port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst20ate2port1  = &AbstractPort{Desc: "absate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port2  = &AbstractPort{Desc: "absate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port3  = &AbstractPort{Desc: "absate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port4  = &AbstractPort{Desc: "absate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port5  = &AbstractPort{Desc: "absate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port6  = &AbstractPort{Desc: "absate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port7  = &AbstractPort{Desc: "absate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port8  = &AbstractPort{Desc: "absate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port9  = &AbstractPort{Desc: "absate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port10 = &AbstractPort{Desc: "absate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port11 = &AbstractPort{Desc: "absate2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port12 = &AbstractPort{Desc: "absate2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port13 = &AbstractPort{Desc: "absate2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port14 = &AbstractPort{Desc: "absate2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port15 = &AbstractPort{Desc: "absate2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port16 = &AbstractPort{Desc: "absate2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port17 = &AbstractPort{Desc: "absate2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port18 = &AbstractPort{Desc: "absate2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port19 = &AbstractPort{Desc: "absate2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2port20 = &AbstractPort{Desc: "absate2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate2       = &AbstractNode{Desc: "absate2", Ports: []*AbstractPort{abst20ate2port1, abst20ate2port2, abst20ate2port3, abst20ate2port4, abst20ate2port5, abst20ate2port6, abst20ate2port7, abst20ate2port8, abst20ate2port9, abst20ate2port10, abst20ate2port11, abst20ate2port12, abst20ate2port13, abst20ate2port14, abst20ate2port15, abst20ate2port16, abst20ate2port17, abst20ate2port18, abst20ate2port19, abst20ate2port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst20ate3port1  = &AbstractPort{Desc: "absate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port2  = &AbstractPort{Desc: "absate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port3  = &AbstractPort{Desc: "absate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port4  = &AbstractPort{Desc: "absate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port5  = &AbstractPort{Desc: "absate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port6  = &AbstractPort{Desc: "absate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port7  = &AbstractPort{Desc: "absate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port8  = &AbstractPort{Desc: "absate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port9  = &AbstractPort{Desc: "absate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port10 = &AbstractPort{Desc: "absate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port11 = &AbstractPort{Desc: "absate3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port12 = &AbstractPort{Desc: "absate3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port13 = &AbstractPort{Desc: "absate3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port14 = &AbstractPort{Desc: "absate3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port15 = &AbstractPort{Desc: "absate3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port16 = &AbstractPort{Desc: "absate3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port17 = &AbstractPort{Desc: "absate3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port18 = &AbstractPort{Desc: "absate3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port19 = &AbstractPort{Desc: "absate3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3port20 = &AbstractPort{Desc: "absate3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate3       = &AbstractNode{Desc: "absate3", Ports: []*AbstractPort{abst20ate3port1, abst20ate3port2, abst20ate3port3, abst20ate3port4, abst20ate3port5, abst20ate3port6, abst20ate3port7, abst20ate3port8, abst20ate3port9, abst20ate3port10, abst20ate3port11, abst20ate3port12, abst20ate3port13, abst20ate3port14, abst20ate3port15, abst20ate3port16, abst20ate3port17, abst20ate3port18, abst20ate3port19, abst20ate3port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst20ate4port1  = &AbstractPort{Desc: "absate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port2  = &AbstractPort{Desc: "absate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port3  = &AbstractPort{Desc: "absate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port4  = &AbstractPort{Desc: "absate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port5  = &AbstractPort{Desc: "absate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port6  = &AbstractPort{Desc: "absate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port7  = &AbstractPort{Desc: "absate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port8  = &AbstractPort{Desc: "absate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port9  = &AbstractPort{Desc: "absate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port10 = &AbstractPort{Desc: "absate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port11 = &AbstractPort{Desc: "absate4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port12 = &AbstractPort{Desc: "absate4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port13 = &AbstractPort{Desc: "absate4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port14 = &AbstractPort{Desc: "absate4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port15 = &AbstractPort{Desc: "absate4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port16 = &AbstractPort{Desc: "absate4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port17 = &AbstractPort{Desc: "absate4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port18 = &AbstractPort{Desc: "absate4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port19 = &AbstractPort{Desc: "absate4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4port20 = &AbstractPort{Desc: "absate4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst20ate4       = &AbstractNode{Desc: "absate4", Ports: []*AbstractPort{abst20ate4port1, abst20ate4port2, abst20ate4port3, abst20ate4port4, abst20ate4port5, abst20ate4port6, abst20ate4port7, abst20ate4port8, abst20ate4port9, abst20ate4port10, abst20ate4port11, abst20ate4port12, abst20ate4port13, abst20ate4port14, abst20ate4port15, abst20ate4port16, abst20ate4port17, abst20ate4port18, abst20ate4port19, abst20ate4port20}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
)

var abstractGraph20 = &AbstractGraph{
	Nodes: []*AbstractNode{abst20dut1,
		abst20ate1,
		abst20dut2,
		abst20ate2,
		abst20dut3,
		abst20ate3,
		abst20dut4,
		abst20ate4},
	Edges: []*AbstractEdge{
		{abst20dut1port1, abst20ate1port1},
		{abst20dut1port2, abst20ate1port2},
		{abst20dut1port3, abst20ate1port3},
		{abst20dut1port4, abst20ate1port4},
		{abst20dut1port5, abst20ate1port5},
		{abst20dut1port6, abst20ate1port6},
		{abst20dut1port7, abst20ate1port7},
		{abst20dut1port8, abst20ate1port8},
		{abst20dut1port9, abst20ate1port9},
		{abst20dut1port10, abst20ate1port10},
		{abst20dut1port11, abst20ate1port11},
		{abst20dut1port12, abst20ate1port12},
		{abst20dut1port13, abst20ate1port13},
		{abst20dut1port14, abst20ate1port14},
		{abst20dut1port15, abst20ate1port15},
		{abst20dut1port16, abst20ate1port16},
		{abst20dut1port17, abst20ate1port17},
		{abst20dut1port18, abst20ate1port18},
		{abst20dut1port19, abst20ate1port19},
		{abst20dut1port20, abst20ate1port20},

		{abst20dut2port1, abst20ate2port1},
		{abst20dut2port2, abst20ate2port2},
		{abst20dut2port3, abst20ate2port3},
		{abst20dut2port4, abst20ate2port4},
		{abst20dut2port5, abst20ate2port5},
		{abst20dut2port6, abst20ate2port6},
		{abst20dut2port7, abst20ate2port7},
		{abst20dut2port8, abst20ate2port8},
		{abst20dut2port9, abst20ate2port9},
		{abst20dut2port10, abst20ate2port10},
		{abst20dut2port11, abst20ate2port11},
		{abst20dut2port12, abst20ate2port12},
		{abst20dut2port13, abst20ate2port13},
		{abst20dut2port14, abst20ate2port14},
		{abst20dut2port15, abst20ate2port15},
		{abst20dut2port16, abst20ate2port16},
		{abst20dut2port17, abst20ate2port17},
		{abst20dut2port18, abst20ate2port18},
		{abst20dut2port19, abst20ate2port19},
		{abst20dut2port20, abst20ate2port20},
		{abst20dut3port1, abst20ate3port1},
		{abst20dut3port2, abst20ate3port2},
		{abst20dut3port3, abst20ate3port3},
		{abst20dut3port4, abst20ate3port4},
		{abst20dut3port5, abst20ate3port5},
		{abst20dut3port6, abst20ate3port6},
		{abst20dut3port7, abst20ate3port7},
		{abst20dut3port8, abst20ate3port8},
		{abst20dut3port9, abst20ate3port9},
		{abst20dut3port10, abst20ate3port10},
		{abst20dut3port11, abst20ate3port11},
		{abst20dut3port12, abst20ate3port12},
		{abst20dut3port13, abst20ate3port13},
		{abst20dut3port14, abst20ate3port14},
		{abst20dut3port15, abst20ate3port15},
		{abst20dut3port16, abst20ate3port16},
		{abst20dut3port17, abst20ate3port17},
		{abst20dut3port18, abst20ate3port18},
		{abst20dut3port19, abst20ate3port19},
		{abst20dut3port20, abst20ate3port20},

		{abst20dut4port1, abst20ate4port1},
		{abst20dut4port2, abst20ate4port2},
		{abst20dut4port3, abst20ate4port3},
		{abst20dut4port4, abst20ate4port4},
		{abst20dut4port5, abst20ate4port5},
		{abst20dut4port6, abst20ate4port6},
		{abst20dut4port7, abst20ate4port7},
		{abst20dut4port8, abst20ate4port8},
		{abst20dut4port9, abst20ate4port9},
		{abst20dut4port10, abst20ate4port10},
		{abst20dut4port11, abst20ate4port11},
		{abst20dut4port12, abst20ate4port12},
		{abst20dut4port13, abst20ate4port13},
		{abst20dut4port14, abst20ate4port14},
		{abst20dut4port15, abst20ate4port15},
		{abst20dut4port16, abst20ate4port16},
		{abst20dut4port17, abst20ate4port17},
		{abst20dut4port18, abst20ate4port18},
		{abst20dut4port19, abst20ate4port19},
		{abst20dut4port20, abst20ate4port20}},
}

func TestSolveTestHybrid20(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraph20, superGraphTest20)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 8 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 160 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
