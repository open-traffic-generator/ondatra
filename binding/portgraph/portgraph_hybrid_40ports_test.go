// 4 duts, 4 ates, 40 ports each and 3 switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	t40dut1port1  = &ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port2  = &ConcretePort{Desc: "dut1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port3  = &ConcretePort{Desc: "dut1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port4  = &ConcretePort{Desc: "dut1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port5  = &ConcretePort{Desc: "dut1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port6  = &ConcretePort{Desc: "dut1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port7  = &ConcretePort{Desc: "dut1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port8  = &ConcretePort{Desc: "dut1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port9  = &ConcretePort{Desc: "dut1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port10 = &ConcretePort{Desc: "dut1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port11 = &ConcretePort{Desc: "dut1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port12 = &ConcretePort{Desc: "dut1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port13 = &ConcretePort{Desc: "dut1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port14 = &ConcretePort{Desc: "dut1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port15 = &ConcretePort{Desc: "dut1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port16 = &ConcretePort{Desc: "dut1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port17 = &ConcretePort{Desc: "dut1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port18 = &ConcretePort{Desc: "dut1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port19 = &ConcretePort{Desc: "dut1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port20 = &ConcretePort{Desc: "dut1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port21 = &ConcretePort{Desc: "dut1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port22 = &ConcretePort{Desc: "dut1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port23 = &ConcretePort{Desc: "dut1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port24 = &ConcretePort{Desc: "dut1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port25 = &ConcretePort{Desc: "dut1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port26 = &ConcretePort{Desc: "dut1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port27 = &ConcretePort{Desc: "dut1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port28 = &ConcretePort{Desc: "dut1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port29 = &ConcretePort{Desc: "dut1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port30 = &ConcretePort{Desc: "dut1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port31 = &ConcretePort{Desc: "dut1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port32 = &ConcretePort{Desc: "dut1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port33 = &ConcretePort{Desc: "dut1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port34 = &ConcretePort{Desc: "dut1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port35 = &ConcretePort{Desc: "dut1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port36 = &ConcretePort{Desc: "dut1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port37 = &ConcretePort{Desc: "dut1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port38 = &ConcretePort{Desc: "dut1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port39 = &ConcretePort{Desc: "dut1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1port40 = &ConcretePort{Desc: "dut1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut1       = &ConcreteNode{Desc: "dut1", Ports: []*ConcretePort{t40dut1port1, t40dut1port2, t40dut1port3, t40dut1port4, t40dut1port5, t40dut1port6, t40dut1port7, t40dut1port8, t40dut1port9, t40dut1port10, t40dut1port11, t40dut1port12, t40dut1port13, t40dut1port14, t40dut1port15, t40dut1port16, t40dut1port17, t40dut1port18, t40dut1port19, t40dut1port20, t40dut1port21, t40dut1port22, t40dut1port23, t40dut1port24, t40dut1port25, t40dut1port26, t40dut1port27, t40dut1port28, t40dut1port29, t40dut1port30, t40dut1port31, t40dut1port32, t40dut1port33, t40dut1port34, t40dut1port35, t40dut1port36, t40dut1port37, t40dut1port38, t40dut1port39, t40dut1port40}, Attrs: map[string]string{"vendor": "CISCO"}}

	t40dut2port1  = &ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port2  = &ConcretePort{Desc: "dut2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port3  = &ConcretePort{Desc: "dut2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port4  = &ConcretePort{Desc: "dut2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port5  = &ConcretePort{Desc: "dut2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port6  = &ConcretePort{Desc: "dut2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port7  = &ConcretePort{Desc: "dut2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port8  = &ConcretePort{Desc: "dut2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port9  = &ConcretePort{Desc: "dut2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port10 = &ConcretePort{Desc: "dut2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port11 = &ConcretePort{Desc: "dut2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port12 = &ConcretePort{Desc: "dut2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port13 = &ConcretePort{Desc: "dut2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port14 = &ConcretePort{Desc: "dut2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port15 = &ConcretePort{Desc: "dut2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port16 = &ConcretePort{Desc: "dut2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port17 = &ConcretePort{Desc: "dut2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port18 = &ConcretePort{Desc: "dut2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port19 = &ConcretePort{Desc: "dut2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port20 = &ConcretePort{Desc: "dut2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port21 = &ConcretePort{Desc: "dut2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port22 = &ConcretePort{Desc: "dut2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port23 = &ConcretePort{Desc: "dut2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port24 = &ConcretePort{Desc: "dut2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port25 = &ConcretePort{Desc: "dut2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port26 = &ConcretePort{Desc: "dut2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port27 = &ConcretePort{Desc: "dut2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port28 = &ConcretePort{Desc: "dut2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port29 = &ConcretePort{Desc: "dut2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port30 = &ConcretePort{Desc: "dut2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port31 = &ConcretePort{Desc: "dut2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port32 = &ConcretePort{Desc: "dut2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port33 = &ConcretePort{Desc: "dut2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port34 = &ConcretePort{Desc: "dut2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port35 = &ConcretePort{Desc: "dut2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port36 = &ConcretePort{Desc: "dut2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port37 = &ConcretePort{Desc: "dut2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port38 = &ConcretePort{Desc: "dut2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port39 = &ConcretePort{Desc: "dut2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2port40 = &ConcretePort{Desc: "dut2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut2       = &ConcreteNode{Desc: "dut2", Ports: []*ConcretePort{t40dut2port1, t40dut2port2, t40dut2port3, t40dut2port4, t40dut2port5, t40dut2port6, t40dut2port7, t40dut2port8, t40dut2port9, t40dut2port10, t40dut2port11, t40dut2port12, t40dut2port13, t40dut2port14, t40dut2port15, t40dut2port16, t40dut2port17, t40dut2port18, t40dut2port19, t40dut2port20, t40dut2port21, t40dut2port22, t40dut2port23, t40dut2port24, t40dut2port25, t40dut2port26, t40dut2port27, t40dut2port28, t40dut2port29, t40dut2port30, t40dut2port31, t40dut2port32, t40dut2port33, t40dut2port34, t40dut2port35, t40dut2port36, t40dut2port37, t40dut2port38, t40dut2port39, t40dut2port40}, Attrs: map[string]string{"vendor": "CISCO"}}

	t40dut3port1  = &ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port2  = &ConcretePort{Desc: "dut3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port3  = &ConcretePort{Desc: "dut3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port4  = &ConcretePort{Desc: "dut3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port5  = &ConcretePort{Desc: "dut3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port6  = &ConcretePort{Desc: "dut3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port7  = &ConcretePort{Desc: "dut3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port8  = &ConcretePort{Desc: "dut3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port9  = &ConcretePort{Desc: "dut3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port10 = &ConcretePort{Desc: "dut3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port11 = &ConcretePort{Desc: "dut3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port12 = &ConcretePort{Desc: "dut3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port13 = &ConcretePort{Desc: "dut3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port14 = &ConcretePort{Desc: "dut3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port15 = &ConcretePort{Desc: "dut3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port16 = &ConcretePort{Desc: "dut3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port17 = &ConcretePort{Desc: "dut3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port18 = &ConcretePort{Desc: "dut3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port19 = &ConcretePort{Desc: "dut3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port20 = &ConcretePort{Desc: "dut3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port21 = &ConcretePort{Desc: "dut3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port22 = &ConcretePort{Desc: "dut3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port23 = &ConcretePort{Desc: "dut3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port24 = &ConcretePort{Desc: "dut3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port25 = &ConcretePort{Desc: "dut3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port26 = &ConcretePort{Desc: "dut3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port27 = &ConcretePort{Desc: "dut3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port28 = &ConcretePort{Desc: "dut3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port29 = &ConcretePort{Desc: "dut3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port30 = &ConcretePort{Desc: "dut3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port31 = &ConcretePort{Desc: "dut3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port32 = &ConcretePort{Desc: "dut3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port33 = &ConcretePort{Desc: "dut3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port34 = &ConcretePort{Desc: "dut3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port35 = &ConcretePort{Desc: "dut3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port36 = &ConcretePort{Desc: "dut3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port37 = &ConcretePort{Desc: "dut3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port38 = &ConcretePort{Desc: "dut3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port39 = &ConcretePort{Desc: "dut3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3port40 = &ConcretePort{Desc: "dut3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{t40dut3port1, t40dut3port2, t40dut3port3, t40dut3port4, t40dut3port5, t40dut3port6, t40dut3port7, t40dut3port8, t40dut3port9, t40dut3port10, t40dut3port11, t40dut3port12, t40dut3port13, t40dut3port14, t40dut3port15, t40dut3port16, t40dut3port17, t40dut3port18, t40dut3port19, t40dut3port20, t40dut3port21, t40dut3port22, t40dut3port23, t40dut3port24, t40dut3port25, t40dut3port26, t40dut3port27, t40dut3port28, t40dut3port29, t40dut3port30, t40dut3port31, t40dut3port32, t40dut3port33, t40dut3port34, t40dut3port35, t40dut3port36, t40dut3port37, t40dut3port38, t40dut3port39, t40dut3port40}, Attrs: map[string]string{"vendor": "CISCO"}}

	t40dut4port1  = &ConcretePort{Desc: "dut4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port2  = &ConcretePort{Desc: "dut4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port3  = &ConcretePort{Desc: "dut4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port4  = &ConcretePort{Desc: "dut4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port5  = &ConcretePort{Desc: "dut4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port6  = &ConcretePort{Desc: "dut4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port7  = &ConcretePort{Desc: "dut4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port8  = &ConcretePort{Desc: "dut4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port9  = &ConcretePort{Desc: "dut4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port10 = &ConcretePort{Desc: "dut4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port11 = &ConcretePort{Desc: "dut4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port12 = &ConcretePort{Desc: "dut4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port13 = &ConcretePort{Desc: "dut4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port14 = &ConcretePort{Desc: "dut4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port15 = &ConcretePort{Desc: "dut4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port16 = &ConcretePort{Desc: "dut4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port17 = &ConcretePort{Desc: "dut4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port18 = &ConcretePort{Desc: "dut4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port19 = &ConcretePort{Desc: "dut4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port20 = &ConcretePort{Desc: "dut4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port21 = &ConcretePort{Desc: "dut4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port22 = &ConcretePort{Desc: "dut4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port23 = &ConcretePort{Desc: "dut4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port24 = &ConcretePort{Desc: "dut4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port25 = &ConcretePort{Desc: "dut4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port26 = &ConcretePort{Desc: "dut4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port27 = &ConcretePort{Desc: "dut4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port28 = &ConcretePort{Desc: "dut4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port29 = &ConcretePort{Desc: "dut4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port30 = &ConcretePort{Desc: "dut4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port31 = &ConcretePort{Desc: "dut4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port32 = &ConcretePort{Desc: "dut4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port33 = &ConcretePort{Desc: "dut4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port34 = &ConcretePort{Desc: "dut4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port35 = &ConcretePort{Desc: "dut4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port36 = &ConcretePort{Desc: "dut4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port37 = &ConcretePort{Desc: "dut4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port38 = &ConcretePort{Desc: "dut4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port39 = &ConcretePort{Desc: "dut4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4port40 = &ConcretePort{Desc: "dut4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40dut4       = &ConcreteNode{Desc: "dut4", Ports: []*ConcretePort{t40dut4port1, t40dut4port2, t40dut4port3, t40dut4port4, t40dut4port5, t40dut4port6, t40dut4port7, t40dut4port8, t40dut4port9, t40dut4port10, t40dut4port11, t40dut4port12, t40dut4port13, t40dut4port14, t40dut4port15, t40dut4port16, t40dut4port17, t40dut4port18, t40dut4port19, t40dut4port20, t40dut4port21, t40dut4port22, t40dut4port23, t40dut4port24, t40dut4port25, t40dut4port26, t40dut4port27, t40dut4port28, t40dut4port29, t40dut4port30, t40dut4port31, t40dut4port32, t40dut4port33, t40dut4port34, t40dut4port35, t40dut4port36, t40dut4port37, t40dut4port38, t40dut4port39, t40dut4port40}, Attrs: map[string]string{"vendor": "CISCO"}}

	t40ate1port1  = &ConcretePort{Desc: "ate1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port2  = &ConcretePort{Desc: "ate1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port3  = &ConcretePort{Desc: "ate1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port4  = &ConcretePort{Desc: "ate1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port5  = &ConcretePort{Desc: "ate1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port6  = &ConcretePort{Desc: "ate1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port7  = &ConcretePort{Desc: "ate1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port8  = &ConcretePort{Desc: "ate1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port9  = &ConcretePort{Desc: "ate1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port10 = &ConcretePort{Desc: "ate1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port11 = &ConcretePort{Desc: "ate1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port12 = &ConcretePort{Desc: "ate1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port13 = &ConcretePort{Desc: "ate1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port14 = &ConcretePort{Desc: "ate1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port15 = &ConcretePort{Desc: "ate1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port16 = &ConcretePort{Desc: "ate1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port17 = &ConcretePort{Desc: "ate1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port18 = &ConcretePort{Desc: "ate1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port19 = &ConcretePort{Desc: "ate1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port20 = &ConcretePort{Desc: "ate1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port21 = &ConcretePort{Desc: "ate1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port22 = &ConcretePort{Desc: "ate1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port23 = &ConcretePort{Desc: "ate1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port24 = &ConcretePort{Desc: "ate1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port25 = &ConcretePort{Desc: "ate1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port26 = &ConcretePort{Desc: "ate1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port27 = &ConcretePort{Desc: "ate1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port28 = &ConcretePort{Desc: "ate1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port29 = &ConcretePort{Desc: "ate1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port30 = &ConcretePort{Desc: "ate1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port31 = &ConcretePort{Desc: "ate1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port32 = &ConcretePort{Desc: "ate1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port33 = &ConcretePort{Desc: "ate1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port34 = &ConcretePort{Desc: "ate1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port35 = &ConcretePort{Desc: "ate1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port36 = &ConcretePort{Desc: "ate1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port37 = &ConcretePort{Desc: "ate1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port38 = &ConcretePort{Desc: "ate1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port39 = &ConcretePort{Desc: "ate1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1port40 = &ConcretePort{Desc: "ate1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate1       = &ConcreteNode{Desc: "ate1", Ports: []*ConcretePort{t40ate1port1, t40ate1port2, t40ate1port3, t40ate1port4, t40ate1port5, t40ate1port6, t40ate1port7, t40ate1port8, t40ate1port9, t40ate1port10, t40ate1port11, t40ate1port12, t40ate1port13, t40ate1port14, t40ate1port15, t40ate1port16, t40ate1port17, t40ate1port18, t40ate1port19, t40ate1port20, t40ate1port21, t40ate1port22, t40ate1port23, t40ate1port24, t40ate1port25, t40ate1port26, t40ate1port27, t40ate1port28, t40ate1port29, t40ate1port30, t40ate1port31, t40ate1port32, t40ate1port33, t40ate1port34, t40ate1port35, t40ate1port36, t40ate1port37, t40ate1port38, t40ate1port39, t40ate1port40}, Attrs: map[string]string{"vendor": "TGEN"}}

	t40ate2port1  = &ConcretePort{Desc: "ate2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port2  = &ConcretePort{Desc: "ate2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port3  = &ConcretePort{Desc: "ate2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port4  = &ConcretePort{Desc: "ate2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port5  = &ConcretePort{Desc: "ate2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port6  = &ConcretePort{Desc: "ate2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port7  = &ConcretePort{Desc: "ate2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port8  = &ConcretePort{Desc: "ate2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port9  = &ConcretePort{Desc: "ate2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port10 = &ConcretePort{Desc: "ate2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port11 = &ConcretePort{Desc: "ate2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port12 = &ConcretePort{Desc: "ate2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port13 = &ConcretePort{Desc: "ate2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port14 = &ConcretePort{Desc: "ate2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port15 = &ConcretePort{Desc: "ate2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port16 = &ConcretePort{Desc: "ate2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port17 = &ConcretePort{Desc: "ate2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port18 = &ConcretePort{Desc: "ate2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port19 = &ConcretePort{Desc: "ate2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port20 = &ConcretePort{Desc: "ate2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port21 = &ConcretePort{Desc: "ate2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port22 = &ConcretePort{Desc: "ate2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port23 = &ConcretePort{Desc: "ate2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port24 = &ConcretePort{Desc: "ate2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port25 = &ConcretePort{Desc: "ate2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port26 = &ConcretePort{Desc: "ate2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port27 = &ConcretePort{Desc: "ate2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port28 = &ConcretePort{Desc: "ate2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port29 = &ConcretePort{Desc: "ate2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port30 = &ConcretePort{Desc: "ate2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port31 = &ConcretePort{Desc: "ate2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port32 = &ConcretePort{Desc: "ate2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port33 = &ConcretePort{Desc: "ate2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port34 = &ConcretePort{Desc: "ate2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port35 = &ConcretePort{Desc: "ate2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port36 = &ConcretePort{Desc: "ate2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port37 = &ConcretePort{Desc: "ate2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port38 = &ConcretePort{Desc: "ate2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port39 = &ConcretePort{Desc: "ate2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2port40 = &ConcretePort{Desc: "ate2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate2       = &ConcreteNode{Desc: "ate2", Ports: []*ConcretePort{t40ate2port1, t40ate2port2, t40ate2port3, t40ate2port4, t40ate2port5, t40ate2port6, t40ate2port7, t40ate2port8, t40ate2port9, t40ate2port10, t40ate2port11, t40ate2port12, t40ate2port13, t40ate2port14, t40ate2port15, t40ate2port16, t40ate2port17, t40ate2port18, t40ate2port19, t40ate2port20, t40ate2port21, t40ate2port22, t40ate2port23, t40ate2port24, t40ate2port25, t40ate2port26, t40ate2port27, t40ate2port28, t40ate2port29, t40ate2port30, t40ate2port31, t40ate2port32, t40ate2port33, t40ate2port34, t40ate2port35, t40ate2port36, t40ate2port37, t40ate2port38, t40ate2port39, t40ate2port40}, Attrs: map[string]string{"vendor": "TGEN"}}

	t40ate3port1  = &ConcretePort{Desc: "ate3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port2  = &ConcretePort{Desc: "ate3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port3  = &ConcretePort{Desc: "ate3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port4  = &ConcretePort{Desc: "ate3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port5  = &ConcretePort{Desc: "ate3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port6  = &ConcretePort{Desc: "ate3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port7  = &ConcretePort{Desc: "ate3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port8  = &ConcretePort{Desc: "ate3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port9  = &ConcretePort{Desc: "ate3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port10 = &ConcretePort{Desc: "ate3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port11 = &ConcretePort{Desc: "ate3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port12 = &ConcretePort{Desc: "ate3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port13 = &ConcretePort{Desc: "ate3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port14 = &ConcretePort{Desc: "ate3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port15 = &ConcretePort{Desc: "ate3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port16 = &ConcretePort{Desc: "ate3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port17 = &ConcretePort{Desc: "ate3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port18 = &ConcretePort{Desc: "ate3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port19 = &ConcretePort{Desc: "ate3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port20 = &ConcretePort{Desc: "ate3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port21 = &ConcretePort{Desc: "ate3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port22 = &ConcretePort{Desc: "ate3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port23 = &ConcretePort{Desc: "ate3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port24 = &ConcretePort{Desc: "ate3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port25 = &ConcretePort{Desc: "ate3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port26 = &ConcretePort{Desc: "ate3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port27 = &ConcretePort{Desc: "ate3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port28 = &ConcretePort{Desc: "ate3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port29 = &ConcretePort{Desc: "ate3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port30 = &ConcretePort{Desc: "ate3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port31 = &ConcretePort{Desc: "ate3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port32 = &ConcretePort{Desc: "ate3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port33 = &ConcretePort{Desc: "ate3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port34 = &ConcretePort{Desc: "ate3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port35 = &ConcretePort{Desc: "ate3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port36 = &ConcretePort{Desc: "ate3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port37 = &ConcretePort{Desc: "ate3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port38 = &ConcretePort{Desc: "ate3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port39 = &ConcretePort{Desc: "ate3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3port40 = &ConcretePort{Desc: "ate3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{t40ate3port1, t40ate3port2, t40ate3port3, t40ate3port4, t40ate3port5, t40ate3port6, t40ate3port7, t40ate3port8, t40ate3port9, t40ate3port10, t40ate3port11, t40ate3port12, t40ate3port13, t40ate3port14, t40ate3port15, t40ate3port16, t40ate3port17, t40ate3port18, t40ate3port19, t40ate3port20, t40ate3port21, t40ate3port22, t40ate3port23, t40ate3port24, t40ate3port25, t40ate3port26, t40ate3port27, t40ate3port28, t40ate3port29, t40ate3port30, t40ate3port31, t40ate3port32, t40ate3port33, t40ate3port34, t40ate3port35, t40ate3port36, t40ate3port37, t40ate3port38, t40ate3port39, t40ate3port40}, Attrs: map[string]string{"vendor": "TGEN"}}

	t40ate4port1  = &ConcretePort{Desc: "ate4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port2  = &ConcretePort{Desc: "ate4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port3  = &ConcretePort{Desc: "ate4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port4  = &ConcretePort{Desc: "ate4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port5  = &ConcretePort{Desc: "ate4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port6  = &ConcretePort{Desc: "ate4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port7  = &ConcretePort{Desc: "ate4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port8  = &ConcretePort{Desc: "ate4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port9  = &ConcretePort{Desc: "ate4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port10 = &ConcretePort{Desc: "ate4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port11 = &ConcretePort{Desc: "ate4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port12 = &ConcretePort{Desc: "ate4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port13 = &ConcretePort{Desc: "ate4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port14 = &ConcretePort{Desc: "ate4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port15 = &ConcretePort{Desc: "ate4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port16 = &ConcretePort{Desc: "ate4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port17 = &ConcretePort{Desc: "ate4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port18 = &ConcretePort{Desc: "ate4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port19 = &ConcretePort{Desc: "ate4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port20 = &ConcretePort{Desc: "ate4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port21 = &ConcretePort{Desc: "ate4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port22 = &ConcretePort{Desc: "ate4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port23 = &ConcretePort{Desc: "ate4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port24 = &ConcretePort{Desc: "ate4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port25 = &ConcretePort{Desc: "ate4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port26 = &ConcretePort{Desc: "ate4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port27 = &ConcretePort{Desc: "ate4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port28 = &ConcretePort{Desc: "ate4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port29 = &ConcretePort{Desc: "ate4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port30 = &ConcretePort{Desc: "ate4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port31 = &ConcretePort{Desc: "ate4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port32 = &ConcretePort{Desc: "ate4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port33 = &ConcretePort{Desc: "ate4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port34 = &ConcretePort{Desc: "ate4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port35 = &ConcretePort{Desc: "ate4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port36 = &ConcretePort{Desc: "ate4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port37 = &ConcretePort{Desc: "ate4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port38 = &ConcretePort{Desc: "ate4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port39 = &ConcretePort{Desc: "ate4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4port40 = &ConcretePort{Desc: "ate4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40ate4       = &ConcreteNode{Desc: "ate4", Ports: []*ConcretePort{t40ate4port1, t40ate4port2, t40ate4port3, t40ate4port4, t40ate4port5, t40ate4port6, t40ate4port7, t40ate4port8, t40ate4port9, t40ate4port10, t40ate4port11, t40ate4port12, t40ate4port13, t40ate4port14, t40ate4port15, t40ate4port16, t40ate4port17, t40ate4port18, t40ate4port19, t40ate4port20, t40ate4port21, t40ate4port22, t40ate4port23, t40ate4port24, t40ate4port25, t40ate4port26, t40ate4port27, t40ate4port28, t40ate4port29, t40ate4port30, t40ate4port31, t40ate4port32, t40ate4port33, t40ate4port34, t40ate4port35, t40ate4port36, t40ate4port37, t40ate4port38, t40ate4port39, t40ate4port40}, Attrs: map[string]string{"vendor": "TGEN"}}

	t40switch1port1  = &ConcretePort{Desc: "switch1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port2  = &ConcretePort{Desc: "switch1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port3  = &ConcretePort{Desc: "switch1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port4  = &ConcretePort{Desc: "switch1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port5  = &ConcretePort{Desc: "switch1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port6  = &ConcretePort{Desc: "switch1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port7  = &ConcretePort{Desc: "switch1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port8  = &ConcretePort{Desc: "switch1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port9  = &ConcretePort{Desc: "switch1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port10 = &ConcretePort{Desc: "switch1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port11 = &ConcretePort{Desc: "switch1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port12 = &ConcretePort{Desc: "switch1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port13 = &ConcretePort{Desc: "switch1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port14 = &ConcretePort{Desc: "switch1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port15 = &ConcretePort{Desc: "switch1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port16 = &ConcretePort{Desc: "switch1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port17 = &ConcretePort{Desc: "switch1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port18 = &ConcretePort{Desc: "switch1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port19 = &ConcretePort{Desc: "switch1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port20 = &ConcretePort{Desc: "switch1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port21 = &ConcretePort{Desc: "switch1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port22 = &ConcretePort{Desc: "switch1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port23 = &ConcretePort{Desc: "switch1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port24 = &ConcretePort{Desc: "switch1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port25 = &ConcretePort{Desc: "switch1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port26 = &ConcretePort{Desc: "switch1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port27 = &ConcretePort{Desc: "switch1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port28 = &ConcretePort{Desc: "switch1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port29 = &ConcretePort{Desc: "switch1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port30 = &ConcretePort{Desc: "switch1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port31 = &ConcretePort{Desc: "switch1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port32 = &ConcretePort{Desc: "switch1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port33 = &ConcretePort{Desc: "switch1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port34 = &ConcretePort{Desc: "switch1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port35 = &ConcretePort{Desc: "switch1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port36 = &ConcretePort{Desc: "switch1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port37 = &ConcretePort{Desc: "switch1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port38 = &ConcretePort{Desc: "switch1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port39 = &ConcretePort{Desc: "switch1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port40 = &ConcretePort{Desc: "switch1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port41 = &ConcretePort{Desc: "switch1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port42 = &ConcretePort{Desc: "switch1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port43 = &ConcretePort{Desc: "switch1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port44 = &ConcretePort{Desc: "switch1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port45 = &ConcretePort{Desc: "switch1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port46 = &ConcretePort{Desc: "switch1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port47 = &ConcretePort{Desc: "switch1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port48 = &ConcretePort{Desc: "switch1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port49 = &ConcretePort{Desc: "switch1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port50 = &ConcretePort{Desc: "switch1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port51 = &ConcretePort{Desc: "switch1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port52 = &ConcretePort{Desc: "switch1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port53 = &ConcretePort{Desc: "switch1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port54 = &ConcretePort{Desc: "switch1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port55 = &ConcretePort{Desc: "switch1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port56 = &ConcretePort{Desc: "switch1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port57 = &ConcretePort{Desc: "switch1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port58 = &ConcretePort{Desc: "switch1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port59 = &ConcretePort{Desc: "switch1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port60 = &ConcretePort{Desc: "switch1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port61 = &ConcretePort{Desc: "switch1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port62 = &ConcretePort{Desc: "switch1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port63 = &ConcretePort{Desc: "switch1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port64 = &ConcretePort{Desc: "switch1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port65 = &ConcretePort{Desc: "switch1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port66 = &ConcretePort{Desc: "switch1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port67 = &ConcretePort{Desc: "switch1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port68 = &ConcretePort{Desc: "switch1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port69 = &ConcretePort{Desc: "switch1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port70 = &ConcretePort{Desc: "switch1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port71 = &ConcretePort{Desc: "switch1:port71", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port72 = &ConcretePort{Desc: "switch1:port72", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port73 = &ConcretePort{Desc: "switch1:port73", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port74 = &ConcretePort{Desc: "switch1:port74", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port75 = &ConcretePort{Desc: "switch1:port75", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port76 = &ConcretePort{Desc: "switch1:port76", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port77 = &ConcretePort{Desc: "switch1:port77", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port78 = &ConcretePort{Desc: "switch1:port78", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port79 = &ConcretePort{Desc: "switch1:port79", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port80 = &ConcretePort{Desc: "switch1:port80", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch1port81 = &ConcretePort{Desc: "switch1:port81", Attrs: map[string]string{"speed": "S_400G"}}

	t40switch1 = &ConcreteNode{Desc: "switch1", Ports: []*ConcretePort{t40switch1port1, t40switch1port2, t40switch1port3, t40switch1port4, t40switch1port5, t40switch1port6, t40switch1port7, t40switch1port8, t40switch1port9, t40switch1port10, t40switch1port11, t40switch1port12, t40switch1port13, t40switch1port14, t40switch1port15, t40switch1port16, t40switch1port17, t40switch1port18, t40switch1port19, t40switch1port20, t40switch1port21, t40switch1port22, t40switch1port23, t40switch1port24, t40switch1port25, t40switch1port26, t40switch1port27, t40switch1port28, t40switch1port29, t40switch1port30, t40switch1port31, t40switch1port32, t40switch1port33, t40switch1port34, t40switch1port35, t40switch1port36, t40switch1port37, t40switch1port38, t40switch1port39, t40switch1port40, t40switch1port41, t40switch1port42, t40switch1port43, t40switch1port44, t40switch1port45, t40switch1port46, t40switch1port47, t40switch1port48, t40switch1port49, t40switch1port50, t40switch1port51, t40switch1port52, t40switch1port53, t40switch1port54, t40switch1port55, t40switch1port56, t40switch1port57, t40switch1port58, t40switch1port59, t40switch1port60, t40switch1port61, t40switch1port62, t40switch1port63, t40switch1port64, t40switch1port65, t40switch1port66, t40switch1port67, t40switch1port68, t40switch1port69, t40switch1port70, t40switch1port71, t40switch1port72, t40switch1port73, t40switch1port74, t40switch1port75, t40switch1port76, t40switch1port77, t40switch1port78, t40switch1port79, t40switch1port80, t40switch1port81}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	t40switch2port1  = &ConcretePort{Desc: "switch2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port2  = &ConcretePort{Desc: "switch2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port3  = &ConcretePort{Desc: "switch2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port4  = &ConcretePort{Desc: "switch2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port5  = &ConcretePort{Desc: "switch2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port6  = &ConcretePort{Desc: "switch2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port7  = &ConcretePort{Desc: "switch2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port8  = &ConcretePort{Desc: "switch2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port9  = &ConcretePort{Desc: "switch2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port10 = &ConcretePort{Desc: "switch2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port11 = &ConcretePort{Desc: "switch2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port12 = &ConcretePort{Desc: "switch2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port13 = &ConcretePort{Desc: "switch2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port14 = &ConcretePort{Desc: "switch2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port15 = &ConcretePort{Desc: "switch2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port16 = &ConcretePort{Desc: "switch2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port17 = &ConcretePort{Desc: "switch2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port18 = &ConcretePort{Desc: "switch2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port19 = &ConcretePort{Desc: "switch2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port20 = &ConcretePort{Desc: "switch2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port21 = &ConcretePort{Desc: "switch2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port22 = &ConcretePort{Desc: "switch2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port23 = &ConcretePort{Desc: "switch2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port24 = &ConcretePort{Desc: "switch2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port25 = &ConcretePort{Desc: "switch2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port26 = &ConcretePort{Desc: "switch2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port27 = &ConcretePort{Desc: "switch2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port28 = &ConcretePort{Desc: "switch2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port29 = &ConcretePort{Desc: "switch2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port30 = &ConcretePort{Desc: "switch2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port31 = &ConcretePort{Desc: "switch2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port32 = &ConcretePort{Desc: "switch2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port33 = &ConcretePort{Desc: "switch2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port34 = &ConcretePort{Desc: "switch2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port35 = &ConcretePort{Desc: "switch2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port36 = &ConcretePort{Desc: "switch2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port37 = &ConcretePort{Desc: "switch2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port38 = &ConcretePort{Desc: "switch2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port39 = &ConcretePort{Desc: "switch2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port40 = &ConcretePort{Desc: "switch2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port41 = &ConcretePort{Desc: "switch2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port42 = &ConcretePort{Desc: "switch2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port43 = &ConcretePort{Desc: "switch2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port44 = &ConcretePort{Desc: "switch2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port45 = &ConcretePort{Desc: "switch2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port46 = &ConcretePort{Desc: "switch2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port47 = &ConcretePort{Desc: "switch2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port48 = &ConcretePort{Desc: "switch2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port49 = &ConcretePort{Desc: "switch2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port50 = &ConcretePort{Desc: "switch2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port51 = &ConcretePort{Desc: "switch2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port52 = &ConcretePort{Desc: "switch2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port53 = &ConcretePort{Desc: "switch2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port54 = &ConcretePort{Desc: "switch2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port55 = &ConcretePort{Desc: "switch2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port56 = &ConcretePort{Desc: "switch2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port57 = &ConcretePort{Desc: "switch2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port58 = &ConcretePort{Desc: "switch2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port59 = &ConcretePort{Desc: "switch2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port60 = &ConcretePort{Desc: "switch2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port61 = &ConcretePort{Desc: "switch2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port62 = &ConcretePort{Desc: "switch2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port63 = &ConcretePort{Desc: "switch2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port64 = &ConcretePort{Desc: "switch2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port65 = &ConcretePort{Desc: "switch2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port66 = &ConcretePort{Desc: "switch2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port67 = &ConcretePort{Desc: "switch2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port68 = &ConcretePort{Desc: "switch2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port69 = &ConcretePort{Desc: "switch2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port70 = &ConcretePort{Desc: "switch2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port71 = &ConcretePort{Desc: "switch2:port71", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port72 = &ConcretePort{Desc: "switch2:port72", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port73 = &ConcretePort{Desc: "switch2:port73", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port74 = &ConcretePort{Desc: "switch2:port74", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port75 = &ConcretePort{Desc: "switch2:port75", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port76 = &ConcretePort{Desc: "switch2:port76", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port77 = &ConcretePort{Desc: "switch2:port77", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port78 = &ConcretePort{Desc: "switch2:port78", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port79 = &ConcretePort{Desc: "switch2:port79", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port80 = &ConcretePort{Desc: "switch2:port80", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port81 = &ConcretePort{Desc: "switch2:port81", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch2port82 = &ConcretePort{Desc: "switch2:port82", Attrs: map[string]string{"speed": "S_400G"}}

	t40switch2 = &ConcreteNode{Desc: "switch2", Ports: []*ConcretePort{t40switch2port1, t40switch2port2, t40switch2port3, t40switch2port4, t40switch2port5, t40switch2port6, t40switch2port7, t40switch2port8, t40switch2port9, t40switch2port10, t40switch2port11, t40switch2port12, t40switch2port13, t40switch2port14, t40switch2port15, t40switch2port16, t40switch2port17, t40switch2port18, t40switch2port19, t40switch2port20, t40switch2port21, t40switch2port22, t40switch2port23, t40switch2port24, t40switch2port25, t40switch2port26, t40switch2port27, t40switch2port28, t40switch2port29, t40switch2port30, t40switch2port31, t40switch2port32, t40switch2port33, t40switch2port34, t40switch2port35, t40switch2port36, t40switch2port37, t40switch2port38, t40switch2port39, t40switch2port40, t40switch2port41, t40switch2port42, t40switch2port43, t40switch2port44, t40switch2port45, t40switch2port46, t40switch2port47, t40switch2port48, t40switch2port49, t40switch2port50, t40switch2port51, t40switch2port52, t40switch2port53, t40switch2port54, t40switch2port55, t40switch2port56, t40switch2port57, t40switch2port58, t40switch2port59, t40switch2port60, t40switch2port61, t40switch2port62, t40switch2port63, t40switch2port64, t40switch2port65, t40switch2port66, t40switch2port67, t40switch2port68, t40switch2port69, t40switch2port70, t40switch2port71, t40switch2port72, t40switch2port73, t40switch2port74, t40switch2port75, t40switch2port76, t40switch2port77, t40switch2port78, t40switch2port79, t40switch2port80, t40switch2port81, t40switch2port82}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	t40switch3port1  = &ConcretePort{Desc: "switch3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port2  = &ConcretePort{Desc: "switch3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port3  = &ConcretePort{Desc: "switch3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port4  = &ConcretePort{Desc: "switch3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port5  = &ConcretePort{Desc: "switch3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port6  = &ConcretePort{Desc: "switch3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port7  = &ConcretePort{Desc: "switch3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port8  = &ConcretePort{Desc: "switch3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port9  = &ConcretePort{Desc: "switch3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port10 = &ConcretePort{Desc: "switch3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port11 = &ConcretePort{Desc: "switch3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port12 = &ConcretePort{Desc: "switch3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port13 = &ConcretePort{Desc: "switch3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port14 = &ConcretePort{Desc: "switch3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port15 = &ConcretePort{Desc: "switch3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port16 = &ConcretePort{Desc: "switch3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port17 = &ConcretePort{Desc: "switch3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port18 = &ConcretePort{Desc: "switch3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port19 = &ConcretePort{Desc: "switch3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port20 = &ConcretePort{Desc: "switch3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port21 = &ConcretePort{Desc: "switch3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port22 = &ConcretePort{Desc: "switch3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port23 = &ConcretePort{Desc: "switch3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port24 = &ConcretePort{Desc: "switch3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port25 = &ConcretePort{Desc: "switch3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port26 = &ConcretePort{Desc: "switch3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port27 = &ConcretePort{Desc: "switch3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port28 = &ConcretePort{Desc: "switch3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port29 = &ConcretePort{Desc: "switch3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port30 = &ConcretePort{Desc: "switch3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port31 = &ConcretePort{Desc: "switch3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port32 = &ConcretePort{Desc: "switch3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port33 = &ConcretePort{Desc: "switch3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port34 = &ConcretePort{Desc: "switch3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port35 = &ConcretePort{Desc: "switch3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port36 = &ConcretePort{Desc: "switch3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port37 = &ConcretePort{Desc: "switch3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port38 = &ConcretePort{Desc: "switch3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port39 = &ConcretePort{Desc: "switch3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port40 = &ConcretePort{Desc: "switch3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port41 = &ConcretePort{Desc: "switch3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port42 = &ConcretePort{Desc: "switch3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port43 = &ConcretePort{Desc: "switch3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port44 = &ConcretePort{Desc: "switch3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port45 = &ConcretePort{Desc: "switch3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port46 = &ConcretePort{Desc: "switch3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port47 = &ConcretePort{Desc: "switch3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port48 = &ConcretePort{Desc: "switch3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port49 = &ConcretePort{Desc: "switch3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port50 = &ConcretePort{Desc: "switch3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port51 = &ConcretePort{Desc: "switch3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port52 = &ConcretePort{Desc: "switch3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port53 = &ConcretePort{Desc: "switch3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port54 = &ConcretePort{Desc: "switch3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port55 = &ConcretePort{Desc: "switch3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port56 = &ConcretePort{Desc: "switch3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port57 = &ConcretePort{Desc: "switch3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port58 = &ConcretePort{Desc: "switch3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port59 = &ConcretePort{Desc: "switch3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port60 = &ConcretePort{Desc: "switch3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port61 = &ConcretePort{Desc: "switch3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port62 = &ConcretePort{Desc: "switch3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port63 = &ConcretePort{Desc: "switch3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port64 = &ConcretePort{Desc: "switch3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port65 = &ConcretePort{Desc: "switch3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port66 = &ConcretePort{Desc: "switch3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port67 = &ConcretePort{Desc: "switch3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port68 = &ConcretePort{Desc: "switch3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port69 = &ConcretePort{Desc: "switch3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port70 = &ConcretePort{Desc: "switch3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port71 = &ConcretePort{Desc: "switch3:port71", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port72 = &ConcretePort{Desc: "switch3:port72", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port73 = &ConcretePort{Desc: "switch3:port73", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port74 = &ConcretePort{Desc: "switch3:port74", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port75 = &ConcretePort{Desc: "switch3:port75", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port76 = &ConcretePort{Desc: "switch3:port76", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port77 = &ConcretePort{Desc: "switch3:port77", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port78 = &ConcretePort{Desc: "switch3:port78", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port79 = &ConcretePort{Desc: "switch3:port79", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port80 = &ConcretePort{Desc: "switch3:port80", Attrs: map[string]string{"speed": "S_400G"}}
	t40switch3port81 = &ConcretePort{Desc: "switch3:port81", Attrs: map[string]string{"speed": "S_400G"}}

	t40switch3 = &ConcreteNode{Desc: "switch3", Ports: []*ConcretePort{t40switch3port1, t40switch3port2, t40switch3port3, t40switch3port4, t40switch3port5, t40switch3port6, t40switch3port7, t40switch3port8, t40switch3port9, t40switch3port10, t40switch3port11, t40switch3port12, t40switch3port13, t40switch3port14, t40switch3port15, t40switch3port16, t40switch3port17, t40switch3port18, t40switch3port19, t40switch3port20, t40switch3port21, t40switch3port22, t40switch3port23, t40switch3port24, t40switch3port25, t40switch3port26, t40switch3port27, t40switch3port28, t40switch3port29, t40switch3port30, t40switch3port31, t40switch3port32, t40switch3port33, t40switch3port34, t40switch3port35, t40switch3port36, t40switch3port37, t40switch3port38, t40switch3port39, t40switch3port40, t40switch3port41, t40switch3port42, t40switch3port43, t40switch3port44, t40switch3port45, t40switch3port46, t40switch3port47, t40switch3port48, t40switch3port49, t40switch3port50, t40switch3port51, t40switch3port52, t40switch3port53, t40switch3port54, t40switch3port55, t40switch3port56, t40switch3port57, t40switch3port58, t40switch3port59, t40switch3port60, t40switch3port61, t40switch3port62, t40switch3port63, t40switch3port64, t40switch3port65, t40switch3port66, t40switch3port67, t40switch3port68, t40switch3port69, t40switch3port70, t40switch3port71, t40switch3port72, t40switch3port73, t40switch3port74, t40switch3port75, t40switch3port76, t40switch3port77, t40switch3port78, t40switch3port79, t40switch3port80, t40switch3port81}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}
)

var superGraphTest40 = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		t40dut1,
		t40ate1,
		t40dut2,
		t40ate2,
		t40dut3,
		t40ate3,
		t40dut4,
		t40ate4,
		t40switch1, t40switch2, t40switch3},
	Edges: []*ConcreteEdge{
		{Src: t40dut1port1, Dst: t40ate1port1},
		{Src: t40dut1port2, Dst: t40ate1port2},
		{Src: t40dut1port3, Dst: t40ate1port3},
		{Src: t40dut1port4, Dst: t40ate1port4},
		{Src: t40dut1port5, Dst: t40ate1port5},
		{Src: t40dut1port6, Dst: t40ate1port6},
		{Src: t40dut1port7, Dst: t40ate1port7},
		{Src: t40dut1port8, Dst: t40ate1port8},
		{Src: t40dut1port9, Dst: t40ate1port9},
		{Src: t40dut1port10, Dst: t40ate1port10},
		{Src: t40dut1port11, Dst: t40ate1port11},
		{Src: t40dut1port12, Dst: t40ate1port12},
		{Src: t40dut1port13, Dst: t40ate1port13},
		{Src: t40dut1port14, Dst: t40ate1port14},
		{Src: t40dut1port15, Dst: t40ate1port15},
		{Src: t40dut1port16, Dst: t40ate1port16},
		{Src: t40dut1port17, Dst: t40ate1port17},
		{Src: t40dut1port18, Dst: t40ate1port18},
		{Src: t40dut1port19, Dst: t40ate1port19},
		{Src: t40dut1port20, Dst: t40ate1port20},
		{Src: t40dut1port21, Dst: t40switch1port1},
		{Src: t40dut1port22, Dst: t40switch1port2},
		{Src: t40dut1port23, Dst: t40switch1port3},
		{Src: t40dut1port24, Dst: t40switch1port4},
		{Src: t40dut1port25, Dst: t40switch1port5},
		{Src: t40dut1port26, Dst: t40switch1port6},
		{Src: t40dut1port27, Dst: t40switch1port7},
		{Src: t40dut1port28, Dst: t40switch1port8},
		{Src: t40dut1port29, Dst: t40switch1port9},
		{Src: t40dut1port30, Dst: t40switch1port10},
		{Src: t40dut1port31, Dst: t40switch1port11},
		{Src: t40dut1port32, Dst: t40switch1port12},
		{Src: t40dut1port33, Dst: t40switch1port13},
		{Src: t40dut1port34, Dst: t40switch1port14},
		{Src: t40dut1port35, Dst: t40switch1port15},
		{Src: t40dut1port36, Dst: t40switch1port16},
		{Src: t40dut1port37, Dst: t40switch1port17},
		{Src: t40dut1port38, Dst: t40switch1port18},
		{Src: t40dut1port39, Dst: t40switch1port19},
		{Src: t40dut1port40, Dst: t40switch1port20},

		{Src: t40switch1port21, Dst: t40ate1port21},
		{Src: t40switch1port22, Dst: t40ate1port22},
		{Src: t40switch1port23, Dst: t40ate1port23},
		{Src: t40switch1port24, Dst: t40ate1port24},
		{Src: t40switch1port25, Dst: t40ate1port25},
		{Src: t40switch1port26, Dst: t40ate1port26},
		{Src: t40switch1port27, Dst: t40ate1port27},
		{Src: t40switch1port28, Dst: t40ate1port28},
		{Src: t40switch1port29, Dst: t40ate1port29},
		{Src: t40switch1port30, Dst: t40ate1port30},
		{Src: t40switch1port31, Dst: t40ate1port31},
		{Src: t40switch1port32, Dst: t40ate1port32},
		{Src: t40switch1port33, Dst: t40ate1port33},
		{Src: t40switch1port34, Dst: t40ate1port34},
		{Src: t40switch1port35, Dst: t40ate1port35},
		{Src: t40switch1port36, Dst: t40ate1port36},
		{Src: t40switch1port37, Dst: t40ate1port37},
		{Src: t40switch1port38, Dst: t40ate1port38},
		{Src: t40switch1port39, Dst: t40ate1port39},
		{Src: t40switch1port40, Dst: t40ate1port40},
		{Src: t40switch1port81, Dst: t40switch2port81},

		{Src: t40dut2port1, Dst: t40ate2port1},
		{Src: t40dut2port2, Dst: t40ate2port2},
		{Src: t40dut2port3, Dst: t40ate2port3},
		{Src: t40dut2port4, Dst: t40ate2port4},
		{Src: t40dut2port5, Dst: t40ate2port5},
		{Src: t40dut2port6, Dst: t40ate2port6},
		{Src: t40dut2port7, Dst: t40ate2port7},
		{Src: t40dut2port8, Dst: t40ate2port8},
		{Src: t40dut2port9, Dst: t40ate2port9},
		{Src: t40dut2port10, Dst: t40ate2port10},
		{Src: t40dut2port11, Dst: t40ate2port11},
		{Src: t40dut2port12, Dst: t40ate2port12},
		{Src: t40dut2port13, Dst: t40ate2port13},
		{Src: t40dut2port14, Dst: t40ate2port14},
		{Src: t40dut2port15, Dst: t40ate2port15},
		{Src: t40dut2port16, Dst: t40ate2port16},
		{Src: t40dut2port17, Dst: t40ate2port17},
		{Src: t40dut2port18, Dst: t40ate2port18},
		{Src: t40dut2port19, Dst: t40ate2port19},
		{Src: t40dut2port20, Dst: t40ate2port20},

		{Src: t40dut2port21, Dst: t40switch1port41},
		{Src: t40dut2port22, Dst: t40switch1port42},
		{Src: t40dut2port23, Dst: t40switch1port43},
		{Src: t40dut2port24, Dst: t40switch1port44},
		{Src: t40dut2port25, Dst: t40switch1port45},
		{Src: t40dut2port26, Dst: t40switch1port46},
		{Src: t40dut2port27, Dst: t40switch1port47},
		{Src: t40dut2port28, Dst: t40switch1port48},
		{Src: t40dut2port29, Dst: t40switch1port49},
		{Src: t40dut2port30, Dst: t40switch1port50},
		{Src: t40dut2port31, Dst: t40switch1port51},
		{Src: t40dut2port32, Dst: t40switch1port52},
		{Src: t40dut2port33, Dst: t40switch1port53},
		{Src: t40dut2port34, Dst: t40switch1port54},
		{Src: t40dut2port35, Dst: t40switch1port55},
		{Src: t40dut2port36, Dst: t40switch1port56},
		{Src: t40dut2port37, Dst: t40switch1port57},
		{Src: t40dut2port38, Dst: t40switch1port58},
		{Src: t40dut2port39, Dst: t40switch1port59},
		{Src: t40dut2port40, Dst: t40switch1port60},

		{Src: t40switch1port61, Dst: t40ate2port21},
		{Src: t40switch1port62, Dst: t40ate2port22},
		{Src: t40switch1port63, Dst: t40ate2port23},
		{Src: t40switch1port64, Dst: t40ate2port24},
		{Src: t40switch1port65, Dst: t40ate2port25},
		{Src: t40switch1port66, Dst: t40ate2port26},
		{Src: t40switch1port67, Dst: t40ate2port27},
		{Src: t40switch1port68, Dst: t40ate2port28},
		{Src: t40switch1port69, Dst: t40ate2port29},
		{Src: t40switch1port70, Dst: t40ate2port30},
		{Src: t40switch1port71, Dst: t40ate2port31},
		{Src: t40switch1port72, Dst: t40ate2port32},
		{Src: t40switch1port73, Dst: t40ate2port33},
		{Src: t40switch1port74, Dst: t40ate2port34},
		{Src: t40switch1port75, Dst: t40ate2port35},
		{Src: t40switch1port76, Dst: t40ate2port36},
		{Src: t40switch1port77, Dst: t40ate2port37},
		{Src: t40switch1port78, Dst: t40ate2port38},
		{Src: t40switch1port79, Dst: t40ate2port39},
		{Src: t40switch1port80, Dst: t40ate2port40},

		{Src: t40dut3port1, Dst: t40switch2port1},
		{Src: t40dut3port2, Dst: t40switch2port2},
		{Src: t40dut3port3, Dst: t40switch2port3},
		{Src: t40dut3port4, Dst: t40switch2port4},
		{Src: t40dut3port5, Dst: t40switch2port5},
		{Src: t40dut3port6, Dst: t40switch2port6},
		{Src: t40dut3port7, Dst: t40switch2port7},
		{Src: t40dut3port8, Dst: t40switch2port8},
		{Src: t40dut3port9, Dst: t40switch2port9},
		{Src: t40dut3port10, Dst: t40switch2port10},
		{Src: t40dut3port11, Dst: t40switch2port11},
		{Src: t40dut3port12, Dst: t40switch2port12},
		{Src: t40dut3port13, Dst: t40switch2port13},
		{Src: t40dut3port14, Dst: t40switch2port14},
		{Src: t40dut3port15, Dst: t40switch2port15},
		{Src: t40dut3port16, Dst: t40switch2port16},
		{Src: t40dut3port17, Dst: t40switch2port17},
		{Src: t40dut3port18, Dst: t40switch2port18},
		{Src: t40dut3port19, Dst: t40switch2port19},
		{Src: t40dut3port20, Dst: t40switch2port20},

		{Src: t40switch2port21, Dst: t40ate3port1},
		{Src: t40switch2port22, Dst: t40ate3port2},
		{Src: t40switch2port23, Dst: t40ate3port3},
		{Src: t40switch2port24, Dst: t40ate3port4},
		{Src: t40switch2port25, Dst: t40ate3port5},
		{Src: t40switch2port26, Dst: t40ate3port6},
		{Src: t40switch2port27, Dst: t40ate3port7},
		{Src: t40switch2port28, Dst: t40ate3port8},
		{Src: t40switch2port29, Dst: t40ate3port9},
		{Src: t40switch2port30, Dst: t40ate3port10},
		{Src: t40switch2port31, Dst: t40ate3port11},
		{Src: t40switch2port32, Dst: t40ate3port12},
		{Src: t40switch2port33, Dst: t40ate3port13},
		{Src: t40switch2port34, Dst: t40ate3port14},
		{Src: t40switch2port35, Dst: t40ate3port15},
		{Src: t40switch2port36, Dst: t40ate3port16},
		{Src: t40switch2port37, Dst: t40ate3port17},
		{Src: t40switch2port38, Dst: t40ate3port18},
		{Src: t40switch2port39, Dst: t40ate3port19},
		{Src: t40switch2port40, Dst: t40ate3port20},

		{Src: t40dut4port1, Dst: t40switch2port41},
		{Src: t40dut4port2, Dst: t40switch2port42},
		{Src: t40dut4port3, Dst: t40switch2port43},
		{Src: t40dut4port4, Dst: t40switch2port44},
		{Src: t40dut4port5, Dst: t40switch2port45},
		{Src: t40dut4port6, Dst: t40switch2port46},
		{Src: t40dut4port7, Dst: t40switch2port47},
		{Src: t40dut4port8, Dst: t40switch2port48},
		{Src: t40dut4port9, Dst: t40switch2port49},
		{Src: t40dut4port10, Dst: t40switch2port50},
		{Src: t40dut4port11, Dst: t40switch2port51},
		{Src: t40dut4port12, Dst: t40switch2port52},
		{Src: t40dut4port13, Dst: t40switch2port53},
		{Src: t40dut4port14, Dst: t40switch2port54},
		{Src: t40dut4port15, Dst: t40switch2port55},
		{Src: t40dut4port16, Dst: t40switch2port56},
		{Src: t40dut4port17, Dst: t40switch2port57},
		{Src: t40dut4port18, Dst: t40switch2port58},
		{Src: t40dut4port19, Dst: t40switch2port59},
		{Src: t40dut4port20, Dst: t40switch2port60},

		{Src: t40switch2port61, Dst: t40ate4port1},
		{Src: t40switch2port62, Dst: t40ate4port2},
		{Src: t40switch2port63, Dst: t40ate4port3},
		{Src: t40switch2port64, Dst: t40ate4port4},
		{Src: t40switch2port65, Dst: t40ate4port5},
		{Src: t40switch2port66, Dst: t40ate4port6},
		{Src: t40switch2port67, Dst: t40ate4port7},
		{Src: t40switch2port68, Dst: t40ate4port8},
		{Src: t40switch2port69, Dst: t40ate4port9},
		{Src: t40switch2port70, Dst: t40ate4port10},
		{Src: t40switch2port71, Dst: t40ate4port11},
		{Src: t40switch2port72, Dst: t40ate4port12},
		{Src: t40switch2port73, Dst: t40ate4port13},
		{Src: t40switch2port74, Dst: t40ate4port14},
		{Src: t40switch2port75, Dst: t40ate4port15},
		{Src: t40switch2port76, Dst: t40ate4port16},
		{Src: t40switch2port77, Dst: t40ate4port17},
		{Src: t40switch2port78, Dst: t40ate4port18},
		{Src: t40switch2port79, Dst: t40ate4port19},
		{Src: t40switch2port80, Dst: t40ate4port20},
		{Src: t40switch2port82, Dst: t40switch3port81},

		{Src: t40dut3port21, Dst: t40switch3port1},
		{Src: t40dut3port22, Dst: t40switch3port2},
		{Src: t40dut3port23, Dst: t40switch3port3},
		{Src: t40dut3port24, Dst: t40switch3port4},
		{Src: t40dut3port25, Dst: t40switch3port5},
		{Src: t40dut3port26, Dst: t40switch3port6},
		{Src: t40dut3port27, Dst: t40switch3port7},
		{Src: t40dut3port28, Dst: t40switch3port8},
		{Src: t40dut3port29, Dst: t40switch3port9},
		{Src: t40dut3port30, Dst: t40switch3port10},
		{Src: t40dut3port31, Dst: t40switch3port11},
		{Src: t40dut3port32, Dst: t40switch3port12},
		{Src: t40dut3port33, Dst: t40switch3port13},
		{Src: t40dut3port34, Dst: t40switch3port14},
		{Src: t40dut3port35, Dst: t40switch3port15},
		{Src: t40dut3port36, Dst: t40switch3port16},
		{Src: t40dut3port37, Dst: t40switch3port17},
		{Src: t40dut3port38, Dst: t40switch3port18},
		{Src: t40dut3port39, Dst: t40switch3port19},
		{Src: t40dut3port40, Dst: t40switch3port20},

		{Src: t40switch3port21, Dst: t40ate3port21},
		{Src: t40switch3port22, Dst: t40ate3port22},
		{Src: t40switch3port23, Dst: t40ate3port23},
		{Src: t40switch3port24, Dst: t40ate3port24},
		{Src: t40switch3port25, Dst: t40ate3port25},
		{Src: t40switch3port26, Dst: t40ate3port26},
		{Src: t40switch3port27, Dst: t40ate3port27},
		{Src: t40switch3port28, Dst: t40ate3port28},
		{Src: t40switch3port29, Dst: t40ate3port29},
		{Src: t40switch3port30, Dst: t40ate3port30},
		{Src: t40switch3port31, Dst: t40ate3port31},
		{Src: t40switch3port32, Dst: t40ate3port32},
		{Src: t40switch3port33, Dst: t40ate3port33},
		{Src: t40switch3port34, Dst: t40ate3port34},
		{Src: t40switch3port35, Dst: t40ate3port35},
		{Src: t40switch3port36, Dst: t40ate3port36},
		{Src: t40switch3port37, Dst: t40ate3port37},
		{Src: t40switch3port38, Dst: t40ate3port38},
		{Src: t40switch3port39, Dst: t40ate3port39},
		{Src: t40switch3port40, Dst: t40ate3port40},

		{Src: t40dut4port21, Dst: t40switch3port41},
		{Src: t40dut4port22, Dst: t40switch3port42},
		{Src: t40dut4port23, Dst: t40switch3port43},
		{Src: t40dut4port24, Dst: t40switch3port44},
		{Src: t40dut4port25, Dst: t40switch3port45},
		{Src: t40dut4port26, Dst: t40switch3port46},
		{Src: t40dut4port27, Dst: t40switch3port47},
		{Src: t40dut4port28, Dst: t40switch3port48},
		{Src: t40dut4port29, Dst: t40switch3port49},
		{Src: t40dut4port30, Dst: t40switch3port50},
		{Src: t40dut4port31, Dst: t40switch3port51},
		{Src: t40dut4port32, Dst: t40switch3port52},
		{Src: t40dut4port33, Dst: t40switch3port53},
		{Src: t40dut4port34, Dst: t40switch3port54},
		{Src: t40dut4port35, Dst: t40switch3port55},
		{Src: t40dut4port36, Dst: t40switch3port56},
		{Src: t40dut4port37, Dst: t40switch3port57},
		{Src: t40dut4port38, Dst: t40switch3port58},
		{Src: t40dut4port39, Dst: t40switch3port59},
		{Src: t40dut4port40, Dst: t40switch3port60},

		{Src: t40switch3port61, Dst: t40ate4port21},
		{Src: t40switch3port62, Dst: t40ate4port22},
		{Src: t40switch3port63, Dst: t40ate4port23},
		{Src: t40switch3port64, Dst: t40ate4port24},
		{Src: t40switch3port65, Dst: t40ate4port25},
		{Src: t40switch3port66, Dst: t40ate4port26},
		{Src: t40switch3port67, Dst: t40ate4port27},
		{Src: t40switch3port68, Dst: t40ate4port28},
		{Src: t40switch3port69, Dst: t40ate4port29},
		{Src: t40switch3port70, Dst: t40ate4port30},
		{Src: t40switch3port71, Dst: t40ate4port31},
		{Src: t40switch3port72, Dst: t40ate4port32},
		{Src: t40switch3port73, Dst: t40ate4port33},
		{Src: t40switch3port74, Dst: t40ate4port34},
		{Src: t40switch3port75, Dst: t40ate4port35},
		{Src: t40switch3port76, Dst: t40ate4port36},
		{Src: t40switch3port77, Dst: t40ate4port37},
		{Src: t40switch3port78, Dst: t40ate4port38},
		{Src: t40switch3port79, Dst: t40ate4port39},
		{Src: t40switch3port80, Dst: t40ate4port40},
	},
}

var (
	abst40dut1port1  = &AbstractPort{Desc: "absdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port2  = &AbstractPort{Desc: "absdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port3  = &AbstractPort{Desc: "absdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port4  = &AbstractPort{Desc: "absdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port5  = &AbstractPort{Desc: "absdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port6  = &AbstractPort{Desc: "absdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port7  = &AbstractPort{Desc: "absdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port8  = &AbstractPort{Desc: "absdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port9  = &AbstractPort{Desc: "absdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port10 = &AbstractPort{Desc: "absdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port11 = &AbstractPort{Desc: "absdut1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port12 = &AbstractPort{Desc: "absdut1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port13 = &AbstractPort{Desc: "absdut1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port14 = &AbstractPort{Desc: "absdut1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port15 = &AbstractPort{Desc: "absdut1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port16 = &AbstractPort{Desc: "absdut1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port17 = &AbstractPort{Desc: "absdut1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port18 = &AbstractPort{Desc: "absdut1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port19 = &AbstractPort{Desc: "absdut1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port20 = &AbstractPort{Desc: "absdut1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port21 = &AbstractPort{Desc: "absdut1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port22 = &AbstractPort{Desc: "absdut1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port23 = &AbstractPort{Desc: "absdut1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port24 = &AbstractPort{Desc: "absdut1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port25 = &AbstractPort{Desc: "absdut1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port26 = &AbstractPort{Desc: "absdut1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port27 = &AbstractPort{Desc: "absdut1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port28 = &AbstractPort{Desc: "absdut1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port29 = &AbstractPort{Desc: "absdut1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port30 = &AbstractPort{Desc: "absdut1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port31 = &AbstractPort{Desc: "absdut1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port32 = &AbstractPort{Desc: "absdut1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port33 = &AbstractPort{Desc: "absdut1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port34 = &AbstractPort{Desc: "absdut1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port35 = &AbstractPort{Desc: "absdut1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port36 = &AbstractPort{Desc: "absdut1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port37 = &AbstractPort{Desc: "absdut1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port38 = &AbstractPort{Desc: "absdut1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port39 = &AbstractPort{Desc: "absdut1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1port40 = &AbstractPort{Desc: "absdut1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut1       = &AbstractNode{Desc: "absdut1", Ports: []*AbstractPort{abst40dut1port1, abst40dut1port2, abst40dut1port3, abst40dut1port4, abst40dut1port5, abst40dut1port6, abst40dut1port7, abst40dut1port8, abst40dut1port9, abst40dut1port10, abst40dut1port11, abst40dut1port12, abst40dut1port13, abst40dut1port14, abst40dut1port15, abst40dut1port16, abst40dut1port17, abst40dut1port18, abst40dut1port19, abst40dut1port20, abst40dut1port21, abst40dut1port22, abst40dut1port23, abst40dut1port24, abst40dut1port25, abst40dut1port26, abst40dut1port27, abst40dut1port28, abst40dut1port29, abst40dut1port30,
		abst40dut1port31, abst40dut1port32, abst40dut1port33, abst40dut1port34, abst40dut1port35, abst40dut1port36, abst40dut1port37, abst40dut1port38, abst40dut1port39, abst40dut1port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst40dut2port1  = &AbstractPort{Desc: "absdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port2  = &AbstractPort{Desc: "absdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port3  = &AbstractPort{Desc: "absdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port4  = &AbstractPort{Desc: "absdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port5  = &AbstractPort{Desc: "absdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port6  = &AbstractPort{Desc: "absdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port7  = &AbstractPort{Desc: "absdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port8  = &AbstractPort{Desc: "absdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port9  = &AbstractPort{Desc: "absdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port10 = &AbstractPort{Desc: "absdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port11 = &AbstractPort{Desc: "absdut2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port12 = &AbstractPort{Desc: "absdut2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port13 = &AbstractPort{Desc: "absdut2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port14 = &AbstractPort{Desc: "absdut2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port15 = &AbstractPort{Desc: "absdut2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port16 = &AbstractPort{Desc: "absdut2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port17 = &AbstractPort{Desc: "absdut2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port18 = &AbstractPort{Desc: "absdut2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port19 = &AbstractPort{Desc: "absdut2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port20 = &AbstractPort{Desc: "absdut2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port21 = &AbstractPort{Desc: "absdut2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port22 = &AbstractPort{Desc: "absdut2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port23 = &AbstractPort{Desc: "absdut2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port24 = &AbstractPort{Desc: "absdut2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port25 = &AbstractPort{Desc: "absdut2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port26 = &AbstractPort{Desc: "absdut2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port27 = &AbstractPort{Desc: "absdut2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port28 = &AbstractPort{Desc: "absdut2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port29 = &AbstractPort{Desc: "absdut2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port30 = &AbstractPort{Desc: "absdut2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port31 = &AbstractPort{Desc: "absdut2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port32 = &AbstractPort{Desc: "absdut2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port33 = &AbstractPort{Desc: "absdut2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port34 = &AbstractPort{Desc: "absdut2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port35 = &AbstractPort{Desc: "absdut2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port36 = &AbstractPort{Desc: "absdut2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port37 = &AbstractPort{Desc: "absdut2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port38 = &AbstractPort{Desc: "absdut2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port39 = &AbstractPort{Desc: "absdut2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2port40 = &AbstractPort{Desc: "absdut2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut2       = &AbstractNode{Desc: "absdut2", Ports: []*AbstractPort{abst40dut2port1, abst40dut2port2, abst40dut2port3, abst40dut2port4, abst40dut2port5, abst40dut2port6, abst40dut2port7, abst40dut2port8, abst40dut2port9, abst40dut2port10, abst40dut2port11, abst40dut2port12, abst40dut2port13, abst40dut2port14, abst40dut2port15, abst40dut2port16, abst40dut2port17, abst40dut2port18, abst40dut2port19, abst40dut2port20, abst40dut2port21, abst40dut2port22, abst40dut2port23, abst40dut2port24, abst40dut2port25, abst40dut2port26, abst40dut2port27, abst40dut2port28, abst40dut2port29, abst40dut2port30,
		abst40dut2port31, abst40dut2port32, abst40dut2port33, abst40dut2port34, abst40dut2port35, abst40dut2port36, abst40dut2port37, abst40dut2port38, abst40dut2port39, abst40dut2port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst40dut3port1  = &AbstractPort{Desc: "absdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port2  = &AbstractPort{Desc: "absdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port3  = &AbstractPort{Desc: "absdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port4  = &AbstractPort{Desc: "absdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port5  = &AbstractPort{Desc: "absdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port6  = &AbstractPort{Desc: "absdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port7  = &AbstractPort{Desc: "absdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port8  = &AbstractPort{Desc: "absdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port9  = &AbstractPort{Desc: "absdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port10 = &AbstractPort{Desc: "absdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port11 = &AbstractPort{Desc: "absdut3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port12 = &AbstractPort{Desc: "absdut3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port13 = &AbstractPort{Desc: "absdut3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port14 = &AbstractPort{Desc: "absdut3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port15 = &AbstractPort{Desc: "absdut3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port16 = &AbstractPort{Desc: "absdut3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port17 = &AbstractPort{Desc: "absdut3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port18 = &AbstractPort{Desc: "absdut3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port19 = &AbstractPort{Desc: "absdut3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port20 = &AbstractPort{Desc: "absdut3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port21 = &AbstractPort{Desc: "absdut3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port22 = &AbstractPort{Desc: "absdut3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port23 = &AbstractPort{Desc: "absdut3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port24 = &AbstractPort{Desc: "absdut3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port25 = &AbstractPort{Desc: "absdut3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port26 = &AbstractPort{Desc: "absdut3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port27 = &AbstractPort{Desc: "absdut3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port28 = &AbstractPort{Desc: "absdut3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port29 = &AbstractPort{Desc: "absdut3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port30 = &AbstractPort{Desc: "absdut3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port31 = &AbstractPort{Desc: "absdut3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port32 = &AbstractPort{Desc: "absdut3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port33 = &AbstractPort{Desc: "absdut3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port34 = &AbstractPort{Desc: "absdut3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port35 = &AbstractPort{Desc: "absdut3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port36 = &AbstractPort{Desc: "absdut3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port37 = &AbstractPort{Desc: "absdut3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port38 = &AbstractPort{Desc: "absdut3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port39 = &AbstractPort{Desc: "absdut3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3port40 = &AbstractPort{Desc: "absdut3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut3       = &AbstractNode{Desc: "absdut3", Ports: []*AbstractPort{abst40dut3port1, abst40dut3port2, abst40dut3port3, abst40dut3port4, abst40dut3port5, abst40dut3port6, abst40dut3port7, abst40dut3port8, abst40dut3port9, abst40dut3port10, abst40dut3port11, abst40dut3port12, abst40dut3port13, abst40dut3port14, abst40dut3port15, abst40dut3port16, abst40dut3port17, abst40dut3port18, abst40dut3port19, abst40dut3port20, abst40dut3port21, abst40dut3port22, abst40dut3port23, abst40dut3port24, abst40dut3port25, abst40dut3port26, abst40dut3port27, abst40dut3port28, abst40dut3port29, abst40dut3port30,
		abst40dut3port31, abst40dut3port32, abst40dut3port33, abst40dut3port34, abst40dut3port35, abst40dut3port36, abst40dut3port37, abst40dut3port38, abst40dut3port39, abst40dut3port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst40dut4port1  = &AbstractPort{Desc: "absdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port2  = &AbstractPort{Desc: "absdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port3  = &AbstractPort{Desc: "absdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port4  = &AbstractPort{Desc: "absdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port5  = &AbstractPort{Desc: "absdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port6  = &AbstractPort{Desc: "absdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port7  = &AbstractPort{Desc: "absdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port8  = &AbstractPort{Desc: "absdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port9  = &AbstractPort{Desc: "absdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port10 = &AbstractPort{Desc: "absdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port11 = &AbstractPort{Desc: "absdut4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port12 = &AbstractPort{Desc: "absdut4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port13 = &AbstractPort{Desc: "absdut4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port14 = &AbstractPort{Desc: "absdut4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port15 = &AbstractPort{Desc: "absdut4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port16 = &AbstractPort{Desc: "absdut4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port17 = &AbstractPort{Desc: "absdut4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port18 = &AbstractPort{Desc: "absdut4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port19 = &AbstractPort{Desc: "absdut4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port20 = &AbstractPort{Desc: "absdut4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port21 = &AbstractPort{Desc: "absdut4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port22 = &AbstractPort{Desc: "absdut4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port23 = &AbstractPort{Desc: "absdut4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port24 = &AbstractPort{Desc: "absdut4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port25 = &AbstractPort{Desc: "absdut4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port26 = &AbstractPort{Desc: "absdut4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port27 = &AbstractPort{Desc: "absdut4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port28 = &AbstractPort{Desc: "absdut4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port29 = &AbstractPort{Desc: "absdut4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port30 = &AbstractPort{Desc: "absdut4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port31 = &AbstractPort{Desc: "absdut4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port32 = &AbstractPort{Desc: "absdut4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port33 = &AbstractPort{Desc: "absdut4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port34 = &AbstractPort{Desc: "absdut4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port35 = &AbstractPort{Desc: "absdut4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port36 = &AbstractPort{Desc: "absdut4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port37 = &AbstractPort{Desc: "absdut4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port38 = &AbstractPort{Desc: "absdut4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port39 = &AbstractPort{Desc: "absdut4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4port40 = &AbstractPort{Desc: "absdut4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40dut4       = &AbstractNode{Desc: "absdut4", Ports: []*AbstractPort{abst40dut4port1, abst40dut4port2, abst40dut4port3, abst40dut4port4, abst40dut4port5, abst40dut4port6, abst40dut4port7, abst40dut4port8, abst40dut4port9, abst40dut4port10, abst40dut4port11, abst40dut4port12, abst40dut4port13, abst40dut4port14, abst40dut4port15, abst40dut4port16, abst40dut4port17, abst40dut4port18, abst40dut4port19, abst40dut4port20, abst40dut4port21, abst40dut4port22, abst40dut4port23, abst40dut4port24, abst40dut4port25, abst40dut4port26, abst40dut4port27, abst40dut4port28, abst40dut4port29, abst40dut4port30,
		abst40dut4port31, abst40dut4port32, abst40dut4port33, abst40dut4port34, abst40dut4port35, abst40dut4port36, abst40dut4port37, abst40dut4port38, abst40dut4port39, abst40dut4port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abst40ate1port1  = &AbstractPort{Desc: "absate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port2  = &AbstractPort{Desc: "absate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port3  = &AbstractPort{Desc: "absate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port4  = &AbstractPort{Desc: "absate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port5  = &AbstractPort{Desc: "absate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port6  = &AbstractPort{Desc: "absate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port7  = &AbstractPort{Desc: "absate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port8  = &AbstractPort{Desc: "absate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port9  = &AbstractPort{Desc: "absate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port10 = &AbstractPort{Desc: "absate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port11 = &AbstractPort{Desc: "absate1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port12 = &AbstractPort{Desc: "absate1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port13 = &AbstractPort{Desc: "absate1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port14 = &AbstractPort{Desc: "absate1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port15 = &AbstractPort{Desc: "absate1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port16 = &AbstractPort{Desc: "absate1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port17 = &AbstractPort{Desc: "absate1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port18 = &AbstractPort{Desc: "absate1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port19 = &AbstractPort{Desc: "absate1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port20 = &AbstractPort{Desc: "absate1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port21 = &AbstractPort{Desc: "absate1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port22 = &AbstractPort{Desc: "absate1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port23 = &AbstractPort{Desc: "absate1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port24 = &AbstractPort{Desc: "absate1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port25 = &AbstractPort{Desc: "absate1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port26 = &AbstractPort{Desc: "absate1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port27 = &AbstractPort{Desc: "absate1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port28 = &AbstractPort{Desc: "absate1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port29 = &AbstractPort{Desc: "absate1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port30 = &AbstractPort{Desc: "absate1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port31 = &AbstractPort{Desc: "absate1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port32 = &AbstractPort{Desc: "absate1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port33 = &AbstractPort{Desc: "absate1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port34 = &AbstractPort{Desc: "absate1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port35 = &AbstractPort{Desc: "absate1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port36 = &AbstractPort{Desc: "absate1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port37 = &AbstractPort{Desc: "absate1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port38 = &AbstractPort{Desc: "absate1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port39 = &AbstractPort{Desc: "absate1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1port40 = &AbstractPort{Desc: "absate1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate1       = &AbstractNode{Desc: "absate1", Ports: []*AbstractPort{abst40ate1port1, abst40ate1port2, abst40ate1port3, abst40ate1port4, abst40ate1port5, abst40ate1port6, abst40ate1port7, abst40ate1port8, abst40ate1port9, abst40ate1port10, abst40ate1port11, abst40ate1port12, abst40ate1port13, abst40ate1port14, abst40ate1port15, abst40ate1port16, abst40ate1port17, abst40ate1port18, abst40ate1port19, abst40ate1port20, abst40ate1port21, abst40ate1port22, abst40ate1port23, abst40ate1port24, abst40ate1port25, abst40ate1port26, abst40ate1port27, abst40ate1port28, abst40ate1port29, abst40ate1port30,
		abst40ate1port31, abst40ate1port32, abst40ate1port33, abst40ate1port34, abst40ate1port35, abst40ate1port36, abst40ate1port37, abst40ate1port38, abst40ate1port39, abst40ate1port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst40ate2port1  = &AbstractPort{Desc: "absate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port2  = &AbstractPort{Desc: "absate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port3  = &AbstractPort{Desc: "absate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port4  = &AbstractPort{Desc: "absate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port5  = &AbstractPort{Desc: "absate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port6  = &AbstractPort{Desc: "absate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port7  = &AbstractPort{Desc: "absate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port8  = &AbstractPort{Desc: "absate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port9  = &AbstractPort{Desc: "absate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port10 = &AbstractPort{Desc: "absate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port11 = &AbstractPort{Desc: "absate2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port12 = &AbstractPort{Desc: "absate2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port13 = &AbstractPort{Desc: "absate2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port14 = &AbstractPort{Desc: "absate2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port15 = &AbstractPort{Desc: "absate2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port16 = &AbstractPort{Desc: "absate2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port17 = &AbstractPort{Desc: "absate2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port18 = &AbstractPort{Desc: "absate2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port19 = &AbstractPort{Desc: "absate2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port20 = &AbstractPort{Desc: "absate2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port21 = &AbstractPort{Desc: "absate2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port22 = &AbstractPort{Desc: "absate2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port23 = &AbstractPort{Desc: "absate2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port24 = &AbstractPort{Desc: "absate2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port25 = &AbstractPort{Desc: "absate2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port26 = &AbstractPort{Desc: "absate2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port27 = &AbstractPort{Desc: "absate2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port28 = &AbstractPort{Desc: "absate2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port29 = &AbstractPort{Desc: "absate2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port30 = &AbstractPort{Desc: "absate2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port31 = &AbstractPort{Desc: "absate2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port32 = &AbstractPort{Desc: "absate2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port33 = &AbstractPort{Desc: "absate2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port34 = &AbstractPort{Desc: "absate2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port35 = &AbstractPort{Desc: "absate2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port36 = &AbstractPort{Desc: "absate2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port37 = &AbstractPort{Desc: "absate2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port38 = &AbstractPort{Desc: "absate2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port39 = &AbstractPort{Desc: "absate2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2port40 = &AbstractPort{Desc: "absate2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate2       = &AbstractNode{Desc: "absate2", Ports: []*AbstractPort{abst40ate2port1, abst40ate2port2, abst40ate2port3, abst40ate2port4, abst40ate2port5, abst40ate2port6, abst40ate2port7, abst40ate2port8, abst40ate2port9, abst40ate2port10, abst40ate2port11, abst40ate2port12, abst40ate2port13, abst40ate2port14, abst40ate2port15, abst40ate2port16, abst40ate2port17, abst40ate2port18, abst40ate2port19, abst40ate2port20, abst40ate2port21, abst40ate2port22, abst40ate2port23, abst40ate2port24, abst40ate2port25, abst40ate2port26, abst40ate2port27, abst40ate2port28, abst40ate2port29, abst40ate2port30,
		abst40ate2port31, abst40ate2port32, abst40ate2port33, abst40ate2port34, abst40ate2port35, abst40ate2port36, abst40ate2port37, abst40ate2port38, abst40ate2port39, abst40ate2port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst40ate3port1  = &AbstractPort{Desc: "absate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port2  = &AbstractPort{Desc: "absate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port3  = &AbstractPort{Desc: "absate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port4  = &AbstractPort{Desc: "absate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port5  = &AbstractPort{Desc: "absate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port6  = &AbstractPort{Desc: "absate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port7  = &AbstractPort{Desc: "absate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port8  = &AbstractPort{Desc: "absate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port9  = &AbstractPort{Desc: "absate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port10 = &AbstractPort{Desc: "absate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port11 = &AbstractPort{Desc: "absate3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port12 = &AbstractPort{Desc: "absate3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port13 = &AbstractPort{Desc: "absate3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port14 = &AbstractPort{Desc: "absate3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port15 = &AbstractPort{Desc: "absate3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port16 = &AbstractPort{Desc: "absate3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port17 = &AbstractPort{Desc: "absate3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port18 = &AbstractPort{Desc: "absate3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port19 = &AbstractPort{Desc: "absate3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port20 = &AbstractPort{Desc: "absate3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port21 = &AbstractPort{Desc: "absate3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port22 = &AbstractPort{Desc: "absate3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port23 = &AbstractPort{Desc: "absate3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port24 = &AbstractPort{Desc: "absate3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port25 = &AbstractPort{Desc: "absate3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port26 = &AbstractPort{Desc: "absate3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port27 = &AbstractPort{Desc: "absate3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port28 = &AbstractPort{Desc: "absate3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port29 = &AbstractPort{Desc: "absate3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port30 = &AbstractPort{Desc: "absate3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port31 = &AbstractPort{Desc: "absate3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port32 = &AbstractPort{Desc: "absate3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port33 = &AbstractPort{Desc: "absate3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port34 = &AbstractPort{Desc: "absate3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port35 = &AbstractPort{Desc: "absate3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port36 = &AbstractPort{Desc: "absate3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port37 = &AbstractPort{Desc: "absate3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port38 = &AbstractPort{Desc: "absate3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port39 = &AbstractPort{Desc: "absate3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3port40 = &AbstractPort{Desc: "absate3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate3       = &AbstractNode{Desc: "absate3", Ports: []*AbstractPort{abst40ate3port1, abst40ate3port2, abst40ate3port3, abst40ate3port4, abst40ate3port5, abst40ate3port6, abst40ate3port7, abst40ate3port8, abst40ate3port9, abst40ate3port10, abst40ate3port11, abst40ate3port12, abst40ate3port13, abst40ate3port14, abst40ate3port15, abst40ate3port16, abst40ate3port17, abst40ate3port18, abst40ate3port19, abst40ate3port20, abst40ate3port21, abst40ate3port22, abst40ate3port23, abst40ate3port24, abst40ate3port25, abst40ate3port26, abst40ate3port27, abst40ate3port28, abst40ate3port29, abst40ate3port30,
		abst40ate3port31, abst40ate3port32, abst40ate3port33, abst40ate3port34, abst40ate3port35, abst40ate3port36, abst40ate3port37, abst40ate3port38, abst40ate3port39, abst40ate3port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abst40ate4port1  = &AbstractPort{Desc: "absate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port2  = &AbstractPort{Desc: "absate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port3  = &AbstractPort{Desc: "absate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port4  = &AbstractPort{Desc: "absate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port5  = &AbstractPort{Desc: "absate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port6  = &AbstractPort{Desc: "absate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port7  = &AbstractPort{Desc: "absate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port8  = &AbstractPort{Desc: "absate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port9  = &AbstractPort{Desc: "absate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port10 = &AbstractPort{Desc: "absate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port11 = &AbstractPort{Desc: "absate4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port12 = &AbstractPort{Desc: "absate4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port13 = &AbstractPort{Desc: "absate4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port14 = &AbstractPort{Desc: "absate4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port15 = &AbstractPort{Desc: "absate4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port16 = &AbstractPort{Desc: "absate4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port17 = &AbstractPort{Desc: "absate4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port18 = &AbstractPort{Desc: "absate4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port19 = &AbstractPort{Desc: "absate4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port20 = &AbstractPort{Desc: "absate4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port21 = &AbstractPort{Desc: "absate4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port22 = &AbstractPort{Desc: "absate4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port23 = &AbstractPort{Desc: "absate4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port24 = &AbstractPort{Desc: "absate4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port25 = &AbstractPort{Desc: "absate4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port26 = &AbstractPort{Desc: "absate4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port27 = &AbstractPort{Desc: "absate4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port28 = &AbstractPort{Desc: "absate4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port29 = &AbstractPort{Desc: "absate4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port30 = &AbstractPort{Desc: "absate4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port31 = &AbstractPort{Desc: "absate4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port32 = &AbstractPort{Desc: "absate4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port33 = &AbstractPort{Desc: "absate4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port34 = &AbstractPort{Desc: "absate4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port35 = &AbstractPort{Desc: "absate4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port36 = &AbstractPort{Desc: "absate4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port37 = &AbstractPort{Desc: "absate4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port38 = &AbstractPort{Desc: "absate4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port39 = &AbstractPort{Desc: "absate4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4port40 = &AbstractPort{Desc: "absate4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abst40ate4       = &AbstractNode{Desc: "absate4", Ports: []*AbstractPort{abst40ate4port1, abst40ate4port2, abst40ate4port3, abst40ate4port4, abst40ate4port5, abst40ate4port6, abst40ate4port7, abst40ate4port8, abst40ate4port9, abst40ate4port10, abst40ate4port11, abst40ate4port12, abst40ate4port13, abst40ate4port14, abst40ate4port15, abst40ate4port16, abst40ate4port17, abst40ate4port18, abst40ate4port19, abst40ate4port20, abst40ate4port21, abst40ate4port22, abst40ate4port23, abst40ate4port24, abst40ate4port25, abst40ate4port26, abst40ate4port27, abst40ate4port28, abst40ate4port29, abst40ate4port30,
		abst40ate4port31, abst40ate4port32, abst40ate4port33, abst40ate4port34, abst40ate4port35, abst40ate4port36, abst40ate4port37, abst40ate4port38, abst40ate4port39, abst40ate4port40}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
)

var abstractGraph40 = &AbstractGraph{
	Nodes: []*AbstractNode{abst40dut1,
		abst40ate1,
		abst40dut2,
		abst40ate2,
		abst40dut3,
		abst40ate3,
		abst40dut4,
		abst40ate4},
	Edges: []*AbstractEdge{
		{abst40dut1port1, abst40ate1port1},
		{abst40dut1port2, abst40ate1port2},
		{abst40dut1port3, abst40ate1port3},
		{abst40dut1port4, abst40ate1port4},
		{abst40dut1port5, abst40ate1port5},
		{abst40dut1port6, abst40ate1port6},
		{abst40dut1port7, abst40ate1port7},
		{abst40dut1port8, abst40ate1port8},
		{abst40dut1port9, abst40ate1port9},
		{abst40dut1port10, abst40ate1port10},
		{abst40dut1port11, abst40ate1port11},
		{abst40dut1port12, abst40ate1port12},
		{abst40dut1port13, abst40ate1port13},
		{abst40dut1port14, abst40ate1port14},
		{abst40dut1port15, abst40ate1port15},
		{abst40dut1port16, abst40ate1port16},
		{abst40dut1port17, abst40ate1port17},
		{abst40dut1port18, abst40ate1port18},
		{abst40dut1port19, abst40ate1port19},
		{abst40dut1port20, abst40ate1port20},
		{abst40dut1port21, abst40ate1port21},
		{abst40dut1port22, abst40ate1port22},
		{abst40dut1port23, abst40ate1port23},
		{abst40dut1port24, abst40ate1port24},
		{abst40dut1port25, abst40ate1port25},
		{abst40dut1port26, abst40ate1port26},
		{abst40dut1port27, abst40ate1port27},
		{abst40dut1port28, abst40ate1port28},
		{abst40dut1port29, abst40ate1port29},
		{abst40dut1port30, abst40ate1port30},
		{abst40dut1port31, abst40ate1port31},
		{abst40dut1port32, abst40ate1port32},
		{abst40dut1port33, abst40ate1port33},
		{abst40dut1port34, abst40ate1port34},
		{abst40dut1port35, abst40ate1port35},
		{abst40dut1port36, abst40ate1port36},
		{abst40dut1port37, abst40ate1port37},
		{abst40dut1port38, abst40ate1port38},
		{abst40dut1port39, abst40ate1port39},
		{abst40dut1port40, abst40ate1port40},

		{abst40dut2port1, abst40ate2port1},
		{abst40dut2port2, abst40ate2port2},
		{abst40dut2port3, abst40ate2port3},
		{abst40dut2port4, abst40ate2port4},
		{abst40dut2port5, abst40ate2port5},
		{abst40dut2port6, abst40ate2port6},
		{abst40dut2port7, abst40ate2port7},
		{abst40dut2port8, abst40ate2port8},
		{abst40dut2port9, abst40ate2port9},
		{abst40dut2port10, abst40ate2port10},
		{abst40dut2port11, abst40ate2port11},
		{abst40dut2port12, abst40ate2port12},
		{abst40dut2port13, abst40ate2port13},
		{abst40dut2port14, abst40ate2port14},
		{abst40dut2port15, abst40ate2port15},
		{abst40dut2port16, abst40ate2port16},
		{abst40dut2port17, abst40ate2port17},
		{abst40dut2port18, abst40ate2port18},
		{abst40dut2port19, abst40ate2port19},
		{abst40dut2port20, abst40ate2port20},
		{abst40dut2port21, abst40ate2port21},
		{abst40dut2port22, abst40ate2port22},
		{abst40dut2port23, abst40ate2port23},
		{abst40dut2port24, abst40ate2port24},
		{abst40dut2port25, abst40ate2port25},
		{abst40dut2port26, abst40ate2port26},
		{abst40dut2port27, abst40ate2port27},
		{abst40dut2port28, abst40ate2port28},
		{abst40dut2port29, abst40ate2port29},
		{abst40dut2port30, abst40ate2port30},
		{abst40dut2port31, abst40ate2port31},
		{abst40dut2port32, abst40ate2port32},
		{abst40dut2port33, abst40ate2port33},
		{abst40dut2port34, abst40ate2port34},
		{abst40dut2port35, abst40ate2port35},
		{abst40dut2port36, abst40ate2port36},
		{abst40dut2port37, abst40ate2port37},
		{abst40dut2port38, abst40ate2port38},
		{abst40dut2port39, abst40ate2port39},
		{abst40dut2port40, abst40ate2port40},

		{abst40dut3port1, abst40ate3port1},
		{abst40dut3port2, abst40ate3port2},
		{abst40dut3port3, abst40ate3port3},
		{abst40dut3port4, abst40ate3port4},
		{abst40dut3port5, abst40ate3port5},
		{abst40dut3port6, abst40ate3port6},
		{abst40dut3port7, abst40ate3port7},
		{abst40dut3port8, abst40ate3port8},
		{abst40dut3port9, abst40ate3port9},
		{abst40dut3port10, abst40ate3port10},
		{abst40dut3port11, abst40ate3port11},
		{abst40dut3port12, abst40ate3port12},
		{abst40dut3port13, abst40ate3port13},
		{abst40dut3port14, abst40ate3port14},
		{abst40dut3port15, abst40ate3port15},
		{abst40dut3port16, abst40ate3port16},
		{abst40dut3port17, abst40ate3port17},
		{abst40dut3port18, abst40ate3port18},
		{abst40dut3port19, abst40ate3port19},
		{abst40dut3port20, abst40ate3port20},
		{abst40dut3port21, abst40ate3port21},
		{abst40dut3port22, abst40ate3port22},
		{abst40dut3port23, abst40ate3port23},
		{abst40dut3port24, abst40ate3port24},
		{abst40dut3port25, abst40ate3port25},
		{abst40dut3port26, abst40ate3port26},
		{abst40dut3port27, abst40ate3port27},
		{abst40dut3port28, abst40ate3port28},
		{abst40dut3port29, abst40ate3port29},
		{abst40dut3port30, abst40ate3port30},
		{abst40dut3port31, abst40ate3port31},
		{abst40dut3port32, abst40ate3port32},
		{abst40dut3port33, abst40ate3port33},
		{abst40dut3port34, abst40ate3port34},
		{abst40dut3port35, abst40ate3port35},
		{abst40dut3port36, abst40ate3port36},
		{abst40dut3port37, abst40ate3port37},
		{abst40dut3port38, abst40ate3port38},
		{abst40dut3port39, abst40ate3port39},
		{abst40dut3port40, abst40ate3port40},

		{abst40dut4port1, abst40ate4port1},
		{abst40dut4port2, abst40ate4port2},
		{abst40dut4port3, abst40ate4port3},
		{abst40dut4port4, abst40ate4port4},
		{abst40dut4port5, abst40ate4port5},
		{abst40dut4port6, abst40ate4port6},
		{abst40dut4port7, abst40ate4port7},
		{abst40dut4port8, abst40ate4port8},
		{abst40dut4port9, abst40ate4port9},
		{abst40dut4port10, abst40ate4port10},
		{abst40dut4port11, abst40ate4port11},
		{abst40dut4port12, abst40ate4port12},
		{abst40dut4port13, abst40ate4port13},
		{abst40dut4port14, abst40ate4port14},
		{abst40dut4port15, abst40ate4port15},
		{abst40dut4port16, abst40ate4port16},
		{abst40dut4port17, abst40ate4port17},
		{abst40dut4port18, abst40ate4port18},
		{abst40dut4port19, abst40ate4port19},
		{abst40dut4port20, abst40ate4port20},
		{abst40dut4port21, abst40ate4port21},
		{abst40dut4port22, abst40ate4port22},
		{abst40dut4port23, abst40ate4port23},
		{abst40dut4port24, abst40ate4port24},
		{abst40dut4port25, abst40ate4port25},
		{abst40dut4port26, abst40ate4port26},
		{abst40dut4port27, abst40ate4port27},
		{abst40dut4port28, abst40ate4port28},
		{abst40dut4port29, abst40ate4port29},
		{abst40dut4port30, abst40ate4port30},
		{abst40dut4port31, abst40ate4port31},
		{abst40dut4port32, abst40ate4port32},
		{abst40dut4port33, abst40ate4port33},
		{abst40dut4port34, abst40ate4port34},
		{abst40dut4port35, abst40ate4port35},
		{abst40dut4port36, abst40ate4port36},
		{abst40dut4port37, abst40ate4port37},
		{abst40dut4port38, abst40ate4port38},
		{abst40dut4port39, abst40ate4port39},
		{abst40dut4port40, abst40ate4port40}},
}

func TestSolveTestHybrid40(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraph40, superGraphTest40)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 8 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 320 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
