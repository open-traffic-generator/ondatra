// // 6 duts, 6 ates, 70 ports each and 5 switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	comdut1port1  = &ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port2  = &ConcretePort{Desc: "dut1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port3  = &ConcretePort{Desc: "dut1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port4  = &ConcretePort{Desc: "dut1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port5  = &ConcretePort{Desc: "dut1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port6  = &ConcretePort{Desc: "dut1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port7  = &ConcretePort{Desc: "dut1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port8  = &ConcretePort{Desc: "dut1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port9  = &ConcretePort{Desc: "dut1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port10 = &ConcretePort{Desc: "dut1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port11 = &ConcretePort{Desc: "dut1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port12 = &ConcretePort{Desc: "dut1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port13 = &ConcretePort{Desc: "dut1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port14 = &ConcretePort{Desc: "dut1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port15 = &ConcretePort{Desc: "dut1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port16 = &ConcretePort{Desc: "dut1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port17 = &ConcretePort{Desc: "dut1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port18 = &ConcretePort{Desc: "dut1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port19 = &ConcretePort{Desc: "dut1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port20 = &ConcretePort{Desc: "dut1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port21 = &ConcretePort{Desc: "dut1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port22 = &ConcretePort{Desc: "dut1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port23 = &ConcretePort{Desc: "dut1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port24 = &ConcretePort{Desc: "dut1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port25 = &ConcretePort{Desc: "dut1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port26 = &ConcretePort{Desc: "dut1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port27 = &ConcretePort{Desc: "dut1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port28 = &ConcretePort{Desc: "dut1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port29 = &ConcretePort{Desc: "dut1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port30 = &ConcretePort{Desc: "dut1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port31 = &ConcretePort{Desc: "dut1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port32 = &ConcretePort{Desc: "dut1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port33 = &ConcretePort{Desc: "dut1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port34 = &ConcretePort{Desc: "dut1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port35 = &ConcretePort{Desc: "dut1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port36 = &ConcretePort{Desc: "dut1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port37 = &ConcretePort{Desc: "dut1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port38 = &ConcretePort{Desc: "dut1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port39 = &ConcretePort{Desc: "dut1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port40 = &ConcretePort{Desc: "dut1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port41 = &ConcretePort{Desc: "dut1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port42 = &ConcretePort{Desc: "dut1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port43 = &ConcretePort{Desc: "dut1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port44 = &ConcretePort{Desc: "dut1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port45 = &ConcretePort{Desc: "dut1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port46 = &ConcretePort{Desc: "dut1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port47 = &ConcretePort{Desc: "dut1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port48 = &ConcretePort{Desc: "dut1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port49 = &ConcretePort{Desc: "dut1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port50 = &ConcretePort{Desc: "dut1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port51 = &ConcretePort{Desc: "dut1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port52 = &ConcretePort{Desc: "dut1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port53 = &ConcretePort{Desc: "dut1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port54 = &ConcretePort{Desc: "dut1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port55 = &ConcretePort{Desc: "dut1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port56 = &ConcretePort{Desc: "dut1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port57 = &ConcretePort{Desc: "dut1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port58 = &ConcretePort{Desc: "dut1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port59 = &ConcretePort{Desc: "dut1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port60 = &ConcretePort{Desc: "dut1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port61 = &ConcretePort{Desc: "dut1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port62 = &ConcretePort{Desc: "dut1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port63 = &ConcretePort{Desc: "dut1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port64 = &ConcretePort{Desc: "dut1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port65 = &ConcretePort{Desc: "dut1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port66 = &ConcretePort{Desc: "dut1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port67 = &ConcretePort{Desc: "dut1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port68 = &ConcretePort{Desc: "dut1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port69 = &ConcretePort{Desc: "dut1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1port70 = &ConcretePort{Desc: "dut1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comdut1       = &ConcreteNode{Desc: "dut1", Ports: []*ConcretePort{comdut1port1, comdut1port2, comdut1port3, comdut1port4, comdut1port5, comdut1port6, comdut1port7, comdut1port8, comdut1port9, comdut1port10, comdut1port11, comdut1port12, comdut1port13, comdut1port14, comdut1port15, comdut1port16, comdut1port17, comdut1port18, comdut1port19, comdut1port20, comdut1port21, comdut1port22, comdut1port23, comdut1port24, comdut1port25, comdut1port26, comdut1port27, comdut1port28, comdut1port29, comdut1port30, comdut1port31, comdut1port32, comdut1port33, comdut1port34, comdut1port35, comdut1port36, comdut1port37, comdut1port38, comdut1port39, comdut1port40, comdut1port41, comdut1port42, comdut1port43, comdut1port44, comdut1port45, comdut1port46, comdut1port47, comdut1port48, comdut1port49, comdut1port50, comdut1port51, comdut1port52, comdut1port53, comdut1port54, comdut1port55, comdut1port56, comdut1port57, comdut1port58, comdut1port59, comdut1port60, comdut1port61, comdut1port62, comdut1port63, comdut1port64, comdut1port65, comdut1port66, comdut1port67, comdut1port68, comdut1port69, comdut1port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	comdut2port1  = &ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port2  = &ConcretePort{Desc: "dut2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port3  = &ConcretePort{Desc: "dut2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port4  = &ConcretePort{Desc: "dut2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port5  = &ConcretePort{Desc: "dut2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port6  = &ConcretePort{Desc: "dut2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port7  = &ConcretePort{Desc: "dut2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port8  = &ConcretePort{Desc: "dut2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port9  = &ConcretePort{Desc: "dut2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port10 = &ConcretePort{Desc: "dut2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port11 = &ConcretePort{Desc: "dut2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port12 = &ConcretePort{Desc: "dut2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port13 = &ConcretePort{Desc: "dut2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port14 = &ConcretePort{Desc: "dut2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port15 = &ConcretePort{Desc: "dut2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port16 = &ConcretePort{Desc: "dut2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port17 = &ConcretePort{Desc: "dut2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port18 = &ConcretePort{Desc: "dut2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port19 = &ConcretePort{Desc: "dut2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port20 = &ConcretePort{Desc: "dut2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port21 = &ConcretePort{Desc: "dut2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port22 = &ConcretePort{Desc: "dut2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port23 = &ConcretePort{Desc: "dut2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port24 = &ConcretePort{Desc: "dut2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port25 = &ConcretePort{Desc: "dut2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port26 = &ConcretePort{Desc: "dut2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port27 = &ConcretePort{Desc: "dut2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port28 = &ConcretePort{Desc: "dut2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port29 = &ConcretePort{Desc: "dut2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port30 = &ConcretePort{Desc: "dut2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port31 = &ConcretePort{Desc: "dut2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port32 = &ConcretePort{Desc: "dut2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port33 = &ConcretePort{Desc: "dut2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port34 = &ConcretePort{Desc: "dut2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port35 = &ConcretePort{Desc: "dut2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port36 = &ConcretePort{Desc: "dut2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port37 = &ConcretePort{Desc: "dut2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port38 = &ConcretePort{Desc: "dut2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port39 = &ConcretePort{Desc: "dut2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port40 = &ConcretePort{Desc: "dut2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port41 = &ConcretePort{Desc: "dut2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port42 = &ConcretePort{Desc: "dut2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port43 = &ConcretePort{Desc: "dut2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port44 = &ConcretePort{Desc: "dut2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port45 = &ConcretePort{Desc: "dut2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port46 = &ConcretePort{Desc: "dut2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port47 = &ConcretePort{Desc: "dut2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port48 = &ConcretePort{Desc: "dut2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port49 = &ConcretePort{Desc: "dut2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port50 = &ConcretePort{Desc: "dut2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port51 = &ConcretePort{Desc: "dut2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port52 = &ConcretePort{Desc: "dut2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port53 = &ConcretePort{Desc: "dut2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port54 = &ConcretePort{Desc: "dut2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port55 = &ConcretePort{Desc: "dut2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port56 = &ConcretePort{Desc: "dut2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port57 = &ConcretePort{Desc: "dut2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port58 = &ConcretePort{Desc: "dut2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port59 = &ConcretePort{Desc: "dut2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port60 = &ConcretePort{Desc: "dut2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port61 = &ConcretePort{Desc: "dut2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port62 = &ConcretePort{Desc: "dut2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port63 = &ConcretePort{Desc: "dut2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port64 = &ConcretePort{Desc: "dut2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port65 = &ConcretePort{Desc: "dut2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port66 = &ConcretePort{Desc: "dut2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port67 = &ConcretePort{Desc: "dut2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port68 = &ConcretePort{Desc: "dut2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port69 = &ConcretePort{Desc: "dut2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2port70 = &ConcretePort{Desc: "dut2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comdut2       = &ConcreteNode{Desc: "dut2", Ports: []*ConcretePort{comdut2port1, comdut2port2, comdut2port3, comdut2port4, comdut2port5, comdut2port6, comdut2port7, comdut2port8, comdut2port9, comdut2port10, comdut2port11, comdut2port12, comdut2port13, comdut2port14, comdut2port15, comdut2port16, comdut2port17, comdut2port18, comdut2port19, comdut2port20, comdut2port21, comdut2port22, comdut2port23, comdut2port24, comdut2port25, comdut2port26, comdut2port27, comdut2port28, comdut2port29, comdut2port30, comdut2port31, comdut2port32, comdut2port33, comdut2port34, comdut2port35, comdut2port36, comdut2port37, comdut2port38, comdut2port39, comdut2port40, comdut2port41, comdut2port42, comdut2port43, comdut2port44, comdut2port45, comdut2port46, comdut2port47, comdut2port48, comdut2port49, comdut2port50, comdut2port51, comdut2port52, comdut2port53, comdut2port54, comdut2port55, comdut2port56, comdut2port57, comdut2port58, comdut2port59, comdut2port60, comdut2port61, comdut2port62, comdut2port63, comdut2port64, comdut2port65, comdut2port66, comdut2port67, comdut2port68, comdut2port69, comdut2port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	comdut3port1  = &ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port2  = &ConcretePort{Desc: "dut3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port3  = &ConcretePort{Desc: "dut3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port4  = &ConcretePort{Desc: "dut3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port5  = &ConcretePort{Desc: "dut3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port6  = &ConcretePort{Desc: "dut3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port7  = &ConcretePort{Desc: "dut3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port8  = &ConcretePort{Desc: "dut3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port9  = &ConcretePort{Desc: "dut3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port10 = &ConcretePort{Desc: "dut3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port11 = &ConcretePort{Desc: "dut3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port12 = &ConcretePort{Desc: "dut3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port13 = &ConcretePort{Desc: "dut3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port14 = &ConcretePort{Desc: "dut3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port15 = &ConcretePort{Desc: "dut3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port16 = &ConcretePort{Desc: "dut3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port17 = &ConcretePort{Desc: "dut3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port18 = &ConcretePort{Desc: "dut3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port19 = &ConcretePort{Desc: "dut3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port20 = &ConcretePort{Desc: "dut3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port21 = &ConcretePort{Desc: "dut3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port22 = &ConcretePort{Desc: "dut3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port23 = &ConcretePort{Desc: "dut3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port24 = &ConcretePort{Desc: "dut3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port25 = &ConcretePort{Desc: "dut3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port26 = &ConcretePort{Desc: "dut3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port27 = &ConcretePort{Desc: "dut3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port28 = &ConcretePort{Desc: "dut3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port29 = &ConcretePort{Desc: "dut3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port30 = &ConcretePort{Desc: "dut3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port31 = &ConcretePort{Desc: "dut3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port32 = &ConcretePort{Desc: "dut3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port33 = &ConcretePort{Desc: "dut3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port34 = &ConcretePort{Desc: "dut3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port35 = &ConcretePort{Desc: "dut3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port36 = &ConcretePort{Desc: "dut3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port37 = &ConcretePort{Desc: "dut3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port38 = &ConcretePort{Desc: "dut3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port39 = &ConcretePort{Desc: "dut3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port40 = &ConcretePort{Desc: "dut3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port41 = &ConcretePort{Desc: "dut3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port42 = &ConcretePort{Desc: "dut3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port43 = &ConcretePort{Desc: "dut3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port44 = &ConcretePort{Desc: "dut3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port45 = &ConcretePort{Desc: "dut3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port46 = &ConcretePort{Desc: "dut3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port47 = &ConcretePort{Desc: "dut3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port48 = &ConcretePort{Desc: "dut3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port49 = &ConcretePort{Desc: "dut3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port50 = &ConcretePort{Desc: "dut3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port51 = &ConcretePort{Desc: "dut3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port52 = &ConcretePort{Desc: "dut3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port53 = &ConcretePort{Desc: "dut3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port54 = &ConcretePort{Desc: "dut3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port55 = &ConcretePort{Desc: "dut3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port56 = &ConcretePort{Desc: "dut3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port57 = &ConcretePort{Desc: "dut3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port58 = &ConcretePort{Desc: "dut3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port59 = &ConcretePort{Desc: "dut3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port60 = &ConcretePort{Desc: "dut3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port61 = &ConcretePort{Desc: "dut3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port62 = &ConcretePort{Desc: "dut3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port63 = &ConcretePort{Desc: "dut3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port64 = &ConcretePort{Desc: "dut3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port65 = &ConcretePort{Desc: "dut3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port66 = &ConcretePort{Desc: "dut3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port67 = &ConcretePort{Desc: "dut3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port68 = &ConcretePort{Desc: "dut3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port69 = &ConcretePort{Desc: "dut3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3port70 = &ConcretePort{Desc: "dut3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comdut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{comdut3port1, comdut3port2, comdut3port3, comdut3port4, comdut3port5, comdut3port6, comdut3port7, comdut3port8, comdut3port9, comdut3port10, comdut3port11, comdut3port12, comdut3port13, comdut3port14, comdut3port15, comdut3port16, comdut3port17, comdut3port18, comdut3port19, comdut3port20, comdut3port21, comdut3port22, comdut3port23, comdut3port24, comdut3port25, comdut3port26, comdut3port27, comdut3port28, comdut3port29, comdut3port30, comdut3port31, comdut3port32, comdut3port33, comdut3port34, comdut3port35, comdut3port36, comdut3port37, comdut3port38, comdut3port39, comdut3port40, comdut3port41, comdut3port42, comdut3port43, comdut3port44, comdut3port45, comdut3port46, comdut3port47, comdut3port48, comdut3port49, comdut3port50, comdut3port51, comdut3port52, comdut3port53, comdut3port54, comdut3port55, comdut3port56, comdut3port57, comdut3port58, comdut3port59, comdut3port60, comdut3port61, comdut3port62, comdut3port63, comdut3port64, comdut3port65, comdut3port66, comdut3port67, comdut3port68, comdut3port69, comdut3port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	comdut4port1  = &ConcretePort{Desc: "dut4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port2  = &ConcretePort{Desc: "dut4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port3  = &ConcretePort{Desc: "dut4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port4  = &ConcretePort{Desc: "dut4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port5  = &ConcretePort{Desc: "dut4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port6  = &ConcretePort{Desc: "dut4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port7  = &ConcretePort{Desc: "dut4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port8  = &ConcretePort{Desc: "dut4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port9  = &ConcretePort{Desc: "dut4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port10 = &ConcretePort{Desc: "dut4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port11 = &ConcretePort{Desc: "dut4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port12 = &ConcretePort{Desc: "dut4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port13 = &ConcretePort{Desc: "dut4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port14 = &ConcretePort{Desc: "dut4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port15 = &ConcretePort{Desc: "dut4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port16 = &ConcretePort{Desc: "dut4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port17 = &ConcretePort{Desc: "dut4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port18 = &ConcretePort{Desc: "dut4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port19 = &ConcretePort{Desc: "dut4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port20 = &ConcretePort{Desc: "dut4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port21 = &ConcretePort{Desc: "dut4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port22 = &ConcretePort{Desc: "dut4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port23 = &ConcretePort{Desc: "dut4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port24 = &ConcretePort{Desc: "dut4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port25 = &ConcretePort{Desc: "dut4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port26 = &ConcretePort{Desc: "dut4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port27 = &ConcretePort{Desc: "dut4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port28 = &ConcretePort{Desc: "dut4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port29 = &ConcretePort{Desc: "dut4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port30 = &ConcretePort{Desc: "dut4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port31 = &ConcretePort{Desc: "dut4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port32 = &ConcretePort{Desc: "dut4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port33 = &ConcretePort{Desc: "dut4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port34 = &ConcretePort{Desc: "dut4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port35 = &ConcretePort{Desc: "dut4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port36 = &ConcretePort{Desc: "dut4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port37 = &ConcretePort{Desc: "dut4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port38 = &ConcretePort{Desc: "dut4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port39 = &ConcretePort{Desc: "dut4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port40 = &ConcretePort{Desc: "dut4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port41 = &ConcretePort{Desc: "dut4:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port42 = &ConcretePort{Desc: "dut4:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port43 = &ConcretePort{Desc: "dut4:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port44 = &ConcretePort{Desc: "dut4:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port45 = &ConcretePort{Desc: "dut4:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port46 = &ConcretePort{Desc: "dut4:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port47 = &ConcretePort{Desc: "dut4:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port48 = &ConcretePort{Desc: "dut4:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port49 = &ConcretePort{Desc: "dut4:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port50 = &ConcretePort{Desc: "dut4:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port51 = &ConcretePort{Desc: "dut4:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port52 = &ConcretePort{Desc: "dut4:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port53 = &ConcretePort{Desc: "dut4:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port54 = &ConcretePort{Desc: "dut4:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port55 = &ConcretePort{Desc: "dut4:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port56 = &ConcretePort{Desc: "dut4:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port57 = &ConcretePort{Desc: "dut4:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port58 = &ConcretePort{Desc: "dut4:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port59 = &ConcretePort{Desc: "dut4:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port60 = &ConcretePort{Desc: "dut4:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port61 = &ConcretePort{Desc: "dut4:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port62 = &ConcretePort{Desc: "dut4:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port63 = &ConcretePort{Desc: "dut4:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port64 = &ConcretePort{Desc: "dut4:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port65 = &ConcretePort{Desc: "dut4:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port66 = &ConcretePort{Desc: "dut4:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port67 = &ConcretePort{Desc: "dut4:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port68 = &ConcretePort{Desc: "dut4:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port69 = &ConcretePort{Desc: "dut4:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4port70 = &ConcretePort{Desc: "dut4:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comdut4       = &ConcreteNode{Desc: "dut4", Ports: []*ConcretePort{comdut4port1, comdut4port2, comdut4port3, comdut4port4, comdut4port5, comdut4port6, comdut4port7, comdut4port8, comdut4port9, comdut4port10, comdut4port11, comdut4port12, comdut4port13, comdut4port14, comdut4port15, comdut4port16, comdut4port17, comdut4port18, comdut4port19, comdut4port20, comdut4port21, comdut4port22, comdut4port23, comdut4port24, comdut4port25, comdut4port26, comdut4port27, comdut4port28, comdut4port29, comdut4port30, comdut4port31, comdut4port32, comdut4port33, comdut4port34, comdut4port35, comdut4port36, comdut4port37, comdut4port38, comdut4port39, comdut4port40, comdut4port41, comdut4port42, comdut4port43, comdut4port44, comdut4port45, comdut4port46, comdut4port47, comdut4port48, comdut4port49, comdut4port50, comdut4port51, comdut4port52, comdut4port53, comdut4port54, comdut4port55, comdut4port56, comdut4port57, comdut4port58, comdut4port59, comdut4port60, comdut4port61, comdut4port62, comdut4port63, comdut4port64, comdut4port65, comdut4port66, comdut4port67, comdut4port68, comdut4port69, comdut4port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	comdut5port1  = &ConcretePort{Desc: "dut5:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port2  = &ConcretePort{Desc: "dut5:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port3  = &ConcretePort{Desc: "dut5:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port4  = &ConcretePort{Desc: "dut5:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port5  = &ConcretePort{Desc: "dut5:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port6  = &ConcretePort{Desc: "dut5:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port7  = &ConcretePort{Desc: "dut5:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port8  = &ConcretePort{Desc: "dut5:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port9  = &ConcretePort{Desc: "dut5:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port10 = &ConcretePort{Desc: "dut5:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port11 = &ConcretePort{Desc: "dut5:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port12 = &ConcretePort{Desc: "dut5:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port13 = &ConcretePort{Desc: "dut5:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port14 = &ConcretePort{Desc: "dut5:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port15 = &ConcretePort{Desc: "dut5:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port16 = &ConcretePort{Desc: "dut5:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port17 = &ConcretePort{Desc: "dut5:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port18 = &ConcretePort{Desc: "dut5:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port19 = &ConcretePort{Desc: "dut5:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port20 = &ConcretePort{Desc: "dut5:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port21 = &ConcretePort{Desc: "dut5:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port22 = &ConcretePort{Desc: "dut5:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port23 = &ConcretePort{Desc: "dut5:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port24 = &ConcretePort{Desc: "dut5:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port25 = &ConcretePort{Desc: "dut5:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port26 = &ConcretePort{Desc: "dut5:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port27 = &ConcretePort{Desc: "dut5:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port28 = &ConcretePort{Desc: "dut5:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port29 = &ConcretePort{Desc: "dut5:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port30 = &ConcretePort{Desc: "dut5:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port31 = &ConcretePort{Desc: "dut5:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port32 = &ConcretePort{Desc: "dut5:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port33 = &ConcretePort{Desc: "dut5:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port34 = &ConcretePort{Desc: "dut5:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port35 = &ConcretePort{Desc: "dut5:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port36 = &ConcretePort{Desc: "dut5:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port37 = &ConcretePort{Desc: "dut5:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port38 = &ConcretePort{Desc: "dut5:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port39 = &ConcretePort{Desc: "dut5:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port40 = &ConcretePort{Desc: "dut5:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port41 = &ConcretePort{Desc: "dut5:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port42 = &ConcretePort{Desc: "dut5:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port43 = &ConcretePort{Desc: "dut5:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port44 = &ConcretePort{Desc: "dut5:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port45 = &ConcretePort{Desc: "dut5:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port46 = &ConcretePort{Desc: "dut5:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port47 = &ConcretePort{Desc: "dut5:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port48 = &ConcretePort{Desc: "dut5:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port49 = &ConcretePort{Desc: "dut5:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port50 = &ConcretePort{Desc: "dut5:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port51 = &ConcretePort{Desc: "dut5:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port52 = &ConcretePort{Desc: "dut5:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port53 = &ConcretePort{Desc: "dut5:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port54 = &ConcretePort{Desc: "dut5:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port55 = &ConcretePort{Desc: "dut5:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port56 = &ConcretePort{Desc: "dut5:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port57 = &ConcretePort{Desc: "dut5:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port58 = &ConcretePort{Desc: "dut5:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port59 = &ConcretePort{Desc: "dut5:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port60 = &ConcretePort{Desc: "dut5:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port61 = &ConcretePort{Desc: "dut5:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port62 = &ConcretePort{Desc: "dut5:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port63 = &ConcretePort{Desc: "dut5:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port64 = &ConcretePort{Desc: "dut5:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port65 = &ConcretePort{Desc: "dut5:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port66 = &ConcretePort{Desc: "dut5:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port67 = &ConcretePort{Desc: "dut5:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port68 = &ConcretePort{Desc: "dut5:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port69 = &ConcretePort{Desc: "dut5:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5port70 = &ConcretePort{Desc: "dut5:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comdut5       = &ConcreteNode{Desc: "dut5", Ports: []*ConcretePort{comdut5port1, comdut5port2, comdut5port3, comdut5port4, comdut5port5, comdut5port6, comdut5port7, comdut5port8, comdut5port9, comdut5port10, comdut5port11, comdut5port12, comdut5port13, comdut5port14, comdut5port15, comdut5port16, comdut5port17, comdut5port18, comdut5port19, comdut5port20, comdut5port21, comdut5port22, comdut5port23, comdut5port24, comdut5port25, comdut5port26, comdut5port27, comdut5port28, comdut5port29, comdut5port30, comdut5port31, comdut5port32, comdut5port33, comdut5port34, comdut5port35, comdut5port36, comdut5port37, comdut5port38, comdut5port39, comdut5port40, comdut5port41, comdut5port42, comdut5port43, comdut5port44, comdut5port45, comdut5port46, comdut5port47, comdut5port48, comdut5port49, comdut5port50, comdut5port51, comdut5port52, comdut5port53, comdut5port54, comdut5port55, comdut5port56, comdut5port57, comdut5port58, comdut5port59, comdut5port60, comdut5port61, comdut5port62, comdut5port63, comdut5port64, comdut5port65, comdut5port66, comdut5port67, comdut5port68, comdut5port69, comdut5port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	comdut6port1  = &ConcretePort{Desc: "dut6:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port2  = &ConcretePort{Desc: "dut6:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port3  = &ConcretePort{Desc: "dut6:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port4  = &ConcretePort{Desc: "dut6:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port5  = &ConcretePort{Desc: "dut6:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port6  = &ConcretePort{Desc: "dut6:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port7  = &ConcretePort{Desc: "dut6:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port8  = &ConcretePort{Desc: "dut6:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port9  = &ConcretePort{Desc: "dut6:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port10 = &ConcretePort{Desc: "dut6:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port11 = &ConcretePort{Desc: "dut6:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port12 = &ConcretePort{Desc: "dut6:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port13 = &ConcretePort{Desc: "dut6:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port14 = &ConcretePort{Desc: "dut6:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port15 = &ConcretePort{Desc: "dut6:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port16 = &ConcretePort{Desc: "dut6:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port17 = &ConcretePort{Desc: "dut6:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port18 = &ConcretePort{Desc: "dut6:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port19 = &ConcretePort{Desc: "dut6:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port20 = &ConcretePort{Desc: "dut6:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port21 = &ConcretePort{Desc: "dut6:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port22 = &ConcretePort{Desc: "dut6:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port23 = &ConcretePort{Desc: "dut6:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port24 = &ConcretePort{Desc: "dut6:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port25 = &ConcretePort{Desc: "dut6:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port26 = &ConcretePort{Desc: "dut6:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port27 = &ConcretePort{Desc: "dut6:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port28 = &ConcretePort{Desc: "dut6:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port29 = &ConcretePort{Desc: "dut6:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port30 = &ConcretePort{Desc: "dut6:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port31 = &ConcretePort{Desc: "dut6:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port32 = &ConcretePort{Desc: "dut6:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port33 = &ConcretePort{Desc: "dut6:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port34 = &ConcretePort{Desc: "dut6:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port35 = &ConcretePort{Desc: "dut6:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port36 = &ConcretePort{Desc: "dut6:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port37 = &ConcretePort{Desc: "dut6:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port38 = &ConcretePort{Desc: "dut6:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port39 = &ConcretePort{Desc: "dut6:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port40 = &ConcretePort{Desc: "dut6:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port41 = &ConcretePort{Desc: "dut6:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port42 = &ConcretePort{Desc: "dut6:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port43 = &ConcretePort{Desc: "dut6:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port44 = &ConcretePort{Desc: "dut6:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port45 = &ConcretePort{Desc: "dut6:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port46 = &ConcretePort{Desc: "dut6:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port47 = &ConcretePort{Desc: "dut6:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port48 = &ConcretePort{Desc: "dut6:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port49 = &ConcretePort{Desc: "dut6:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port50 = &ConcretePort{Desc: "dut6:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port51 = &ConcretePort{Desc: "dut6:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port52 = &ConcretePort{Desc: "dut6:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port53 = &ConcretePort{Desc: "dut6:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port54 = &ConcretePort{Desc: "dut6:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port55 = &ConcretePort{Desc: "dut6:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port56 = &ConcretePort{Desc: "dut6:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port57 = &ConcretePort{Desc: "dut6:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port58 = &ConcretePort{Desc: "dut6:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port59 = &ConcretePort{Desc: "dut6:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port60 = &ConcretePort{Desc: "dut6:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port61 = &ConcretePort{Desc: "dut6:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port62 = &ConcretePort{Desc: "dut6:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port63 = &ConcretePort{Desc: "dut6:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port64 = &ConcretePort{Desc: "dut6:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port65 = &ConcretePort{Desc: "dut6:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port66 = &ConcretePort{Desc: "dut6:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port67 = &ConcretePort{Desc: "dut6:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port68 = &ConcretePort{Desc: "dut6:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port69 = &ConcretePort{Desc: "dut6:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6port70 = &ConcretePort{Desc: "dut6:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comdut6       = &ConcreteNode{Desc: "dut6", Ports: []*ConcretePort{comdut6port1, comdut6port2, comdut6port3, comdut6port4, comdut6port5, comdut6port6, comdut6port7, comdut6port8, comdut6port9, comdut6port10, comdut6port11, comdut6port12, comdut6port13, comdut6port14, comdut6port15, comdut6port16, comdut6port17, comdut6port18, comdut6port19, comdut6port20, comdut6port21, comdut6port22, comdut6port23, comdut6port24, comdut6port25, comdut6port26, comdut6port27, comdut6port28, comdut6port29, comdut6port30, comdut6port31, comdut6port32, comdut6port33, comdut6port34, comdut6port35, comdut6port36, comdut6port37, comdut6port38, comdut6port39, comdut6port40, comdut6port41, comdut6port42, comdut6port43, comdut6port44, comdut6port45, comdut6port46, comdut6port47, comdut6port48, comdut6port49, comdut6port50, comdut6port51, comdut6port52, comdut6port53, comdut6port54, comdut6port55, comdut6port56, comdut6port57, comdut6port58, comdut6port59, comdut6port60, comdut6port61, comdut6port62, comdut6port63, comdut6port64, comdut6port65, comdut6port66, comdut6port67, comdut6port68, comdut6port69, comdut6port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	comate1port1  = &ConcretePort{Desc: "ate1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port2  = &ConcretePort{Desc: "ate1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port3  = &ConcretePort{Desc: "ate1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port4  = &ConcretePort{Desc: "ate1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port5  = &ConcretePort{Desc: "ate1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port6  = &ConcretePort{Desc: "ate1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port7  = &ConcretePort{Desc: "ate1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port8  = &ConcretePort{Desc: "ate1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port9  = &ConcretePort{Desc: "ate1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port10 = &ConcretePort{Desc: "ate1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port11 = &ConcretePort{Desc: "ate1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port12 = &ConcretePort{Desc: "ate1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port13 = &ConcretePort{Desc: "ate1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port14 = &ConcretePort{Desc: "ate1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port15 = &ConcretePort{Desc: "ate1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port16 = &ConcretePort{Desc: "ate1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port17 = &ConcretePort{Desc: "ate1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port18 = &ConcretePort{Desc: "ate1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port19 = &ConcretePort{Desc: "ate1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port20 = &ConcretePort{Desc: "ate1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port21 = &ConcretePort{Desc: "ate1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port22 = &ConcretePort{Desc: "ate1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port23 = &ConcretePort{Desc: "ate1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port24 = &ConcretePort{Desc: "ate1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port25 = &ConcretePort{Desc: "ate1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port26 = &ConcretePort{Desc: "ate1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port27 = &ConcretePort{Desc: "ate1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port28 = &ConcretePort{Desc: "ate1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port29 = &ConcretePort{Desc: "ate1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port30 = &ConcretePort{Desc: "ate1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port31 = &ConcretePort{Desc: "ate1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port32 = &ConcretePort{Desc: "ate1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port33 = &ConcretePort{Desc: "ate1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port34 = &ConcretePort{Desc: "ate1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port35 = &ConcretePort{Desc: "ate1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port36 = &ConcretePort{Desc: "ate1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port37 = &ConcretePort{Desc: "ate1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port38 = &ConcretePort{Desc: "ate1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port39 = &ConcretePort{Desc: "ate1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port40 = &ConcretePort{Desc: "ate1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port41 = &ConcretePort{Desc: "ate1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port42 = &ConcretePort{Desc: "ate1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port43 = &ConcretePort{Desc: "ate1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port44 = &ConcretePort{Desc: "ate1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port45 = &ConcretePort{Desc: "ate1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port46 = &ConcretePort{Desc: "ate1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port47 = &ConcretePort{Desc: "ate1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port48 = &ConcretePort{Desc: "ate1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port49 = &ConcretePort{Desc: "ate1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port50 = &ConcretePort{Desc: "ate1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port51 = &ConcretePort{Desc: "ate1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port52 = &ConcretePort{Desc: "ate1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port53 = &ConcretePort{Desc: "ate1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port54 = &ConcretePort{Desc: "ate1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port55 = &ConcretePort{Desc: "ate1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port56 = &ConcretePort{Desc: "ate1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port57 = &ConcretePort{Desc: "ate1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port58 = &ConcretePort{Desc: "ate1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port59 = &ConcretePort{Desc: "ate1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port60 = &ConcretePort{Desc: "ate1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port61 = &ConcretePort{Desc: "ate1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port62 = &ConcretePort{Desc: "ate1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port63 = &ConcretePort{Desc: "ate1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port64 = &ConcretePort{Desc: "ate1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port65 = &ConcretePort{Desc: "ate1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port66 = &ConcretePort{Desc: "ate1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port67 = &ConcretePort{Desc: "ate1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port68 = &ConcretePort{Desc: "ate1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port69 = &ConcretePort{Desc: "ate1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comate1port70 = &ConcretePort{Desc: "ate1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comate1       = &ConcreteNode{Desc: "ate1", Ports: []*ConcretePort{comate1port1, comate1port2, comate1port3, comate1port4, comate1port5, comate1port6, comate1port7, comate1port8, comate1port9, comate1port10, comate1port11, comate1port12, comate1port13, comate1port14, comate1port15, comate1port16, comate1port17, comate1port18, comate1port19, comate1port20, comate1port21, comate1port22, comate1port23, comate1port24, comate1port25, comate1port26, comate1port27, comate1port28, comate1port29, comate1port30, comate1port31, comate1port32, comate1port33, comate1port34, comate1port35, comate1port36, comate1port37, comate1port38, comate1port39, comate1port40, comate1port41, comate1port42, comate1port43, comate1port44, comate1port45, comate1port46, comate1port47, comate1port48, comate1port49, comate1port50, comate1port51, comate1port52, comate1port53, comate1port54, comate1port55, comate1port56, comate1port57, comate1port58, comate1port59, comate1port60, comate1port61, comate1port62, comate1port63, comate1port64, comate1port65, comate1port66, comate1port67, comate1port68, comate1port69, comate1port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	comate2port1  = &ConcretePort{Desc: "ate2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port2  = &ConcretePort{Desc: "ate2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port3  = &ConcretePort{Desc: "ate2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port4  = &ConcretePort{Desc: "ate2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port5  = &ConcretePort{Desc: "ate2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port6  = &ConcretePort{Desc: "ate2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port7  = &ConcretePort{Desc: "ate2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port8  = &ConcretePort{Desc: "ate2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port9  = &ConcretePort{Desc: "ate2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port10 = &ConcretePort{Desc: "ate2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port11 = &ConcretePort{Desc: "ate2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port12 = &ConcretePort{Desc: "ate2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port13 = &ConcretePort{Desc: "ate2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port14 = &ConcretePort{Desc: "ate2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port15 = &ConcretePort{Desc: "ate2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port16 = &ConcretePort{Desc: "ate2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port17 = &ConcretePort{Desc: "ate2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port18 = &ConcretePort{Desc: "ate2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port19 = &ConcretePort{Desc: "ate2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port20 = &ConcretePort{Desc: "ate2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port21 = &ConcretePort{Desc: "ate2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port22 = &ConcretePort{Desc: "ate2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port23 = &ConcretePort{Desc: "ate2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port24 = &ConcretePort{Desc: "ate2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port25 = &ConcretePort{Desc: "ate2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port26 = &ConcretePort{Desc: "ate2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port27 = &ConcretePort{Desc: "ate2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port28 = &ConcretePort{Desc: "ate2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port29 = &ConcretePort{Desc: "ate2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port30 = &ConcretePort{Desc: "ate2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port31 = &ConcretePort{Desc: "ate2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port32 = &ConcretePort{Desc: "ate2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port33 = &ConcretePort{Desc: "ate2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port34 = &ConcretePort{Desc: "ate2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port35 = &ConcretePort{Desc: "ate2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port36 = &ConcretePort{Desc: "ate2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port37 = &ConcretePort{Desc: "ate2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port38 = &ConcretePort{Desc: "ate2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port39 = &ConcretePort{Desc: "ate2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port40 = &ConcretePort{Desc: "ate2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port41 = &ConcretePort{Desc: "ate2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port42 = &ConcretePort{Desc: "ate2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port43 = &ConcretePort{Desc: "ate2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port44 = &ConcretePort{Desc: "ate2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port45 = &ConcretePort{Desc: "ate2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port46 = &ConcretePort{Desc: "ate2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port47 = &ConcretePort{Desc: "ate2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port48 = &ConcretePort{Desc: "ate2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port49 = &ConcretePort{Desc: "ate2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port50 = &ConcretePort{Desc: "ate2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port51 = &ConcretePort{Desc: "ate2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port52 = &ConcretePort{Desc: "ate2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port53 = &ConcretePort{Desc: "ate2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port54 = &ConcretePort{Desc: "ate2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port55 = &ConcretePort{Desc: "ate2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port56 = &ConcretePort{Desc: "ate2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port57 = &ConcretePort{Desc: "ate2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port58 = &ConcretePort{Desc: "ate2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port59 = &ConcretePort{Desc: "ate2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port60 = &ConcretePort{Desc: "ate2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port61 = &ConcretePort{Desc: "ate2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port62 = &ConcretePort{Desc: "ate2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port63 = &ConcretePort{Desc: "ate2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port64 = &ConcretePort{Desc: "ate2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port65 = &ConcretePort{Desc: "ate2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port66 = &ConcretePort{Desc: "ate2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port67 = &ConcretePort{Desc: "ate2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port68 = &ConcretePort{Desc: "ate2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port69 = &ConcretePort{Desc: "ate2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comate2port70 = &ConcretePort{Desc: "ate2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comate2       = &ConcreteNode{Desc: "ate2", Ports: []*ConcretePort{comate2port1, comate2port2, comate2port3, comate2port4, comate2port5, comate2port6, comate2port7, comate2port8, comate2port9, comate2port10, comate2port11, comate2port12, comate2port13, comate2port14, comate2port15, comate2port16, comate2port17, comate2port18, comate2port19, comate2port20, comate2port21, comate2port22, comate2port23, comate2port24, comate2port25, comate2port26, comate2port27, comate2port28, comate2port29, comate2port30, comate2port31, comate2port32, comate2port33, comate2port34, comate2port35, comate2port36, comate2port37, comate2port38, comate2port39, comate2port40, comate2port41, comate2port42, comate2port43, comate2port44, comate2port45, comate2port46, comate2port47, comate2port48, comate2port49, comate2port50, comate2port51, comate2port52, comate2port53, comate2port54, comate2port55, comate2port56, comate2port57, comate2port58, comate2port59, comate2port60, comate2port61, comate2port62, comate2port63, comate2port64, comate2port65, comate2port66, comate2port67, comate2port68, comate2port69, comate2port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	comate3port1  = &ConcretePort{Desc: "ate3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port2  = &ConcretePort{Desc: "ate3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port3  = &ConcretePort{Desc: "ate3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port4  = &ConcretePort{Desc: "ate3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port5  = &ConcretePort{Desc: "ate3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port6  = &ConcretePort{Desc: "ate3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port7  = &ConcretePort{Desc: "ate3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port8  = &ConcretePort{Desc: "ate3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port9  = &ConcretePort{Desc: "ate3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port10 = &ConcretePort{Desc: "ate3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port11 = &ConcretePort{Desc: "ate3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port12 = &ConcretePort{Desc: "ate3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port13 = &ConcretePort{Desc: "ate3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port14 = &ConcretePort{Desc: "ate3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port15 = &ConcretePort{Desc: "ate3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port16 = &ConcretePort{Desc: "ate3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port17 = &ConcretePort{Desc: "ate3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port18 = &ConcretePort{Desc: "ate3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port19 = &ConcretePort{Desc: "ate3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port20 = &ConcretePort{Desc: "ate3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port21 = &ConcretePort{Desc: "ate3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port22 = &ConcretePort{Desc: "ate3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port23 = &ConcretePort{Desc: "ate3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port24 = &ConcretePort{Desc: "ate3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port25 = &ConcretePort{Desc: "ate3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port26 = &ConcretePort{Desc: "ate3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port27 = &ConcretePort{Desc: "ate3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port28 = &ConcretePort{Desc: "ate3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port29 = &ConcretePort{Desc: "ate3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port30 = &ConcretePort{Desc: "ate3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port31 = &ConcretePort{Desc: "ate3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port32 = &ConcretePort{Desc: "ate3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port33 = &ConcretePort{Desc: "ate3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port34 = &ConcretePort{Desc: "ate3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port35 = &ConcretePort{Desc: "ate3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port36 = &ConcretePort{Desc: "ate3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port37 = &ConcretePort{Desc: "ate3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port38 = &ConcretePort{Desc: "ate3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port39 = &ConcretePort{Desc: "ate3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port40 = &ConcretePort{Desc: "ate3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port41 = &ConcretePort{Desc: "ate3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port42 = &ConcretePort{Desc: "ate3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port43 = &ConcretePort{Desc: "ate3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port44 = &ConcretePort{Desc: "ate3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port45 = &ConcretePort{Desc: "ate3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port46 = &ConcretePort{Desc: "ate3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port47 = &ConcretePort{Desc: "ate3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port48 = &ConcretePort{Desc: "ate3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port49 = &ConcretePort{Desc: "ate3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port50 = &ConcretePort{Desc: "ate3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port51 = &ConcretePort{Desc: "ate3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port52 = &ConcretePort{Desc: "ate3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port53 = &ConcretePort{Desc: "ate3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port54 = &ConcretePort{Desc: "ate3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port55 = &ConcretePort{Desc: "ate3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port56 = &ConcretePort{Desc: "ate3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port57 = &ConcretePort{Desc: "ate3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port58 = &ConcretePort{Desc: "ate3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port59 = &ConcretePort{Desc: "ate3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port60 = &ConcretePort{Desc: "ate3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port61 = &ConcretePort{Desc: "ate3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port62 = &ConcretePort{Desc: "ate3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port63 = &ConcretePort{Desc: "ate3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port64 = &ConcretePort{Desc: "ate3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port65 = &ConcretePort{Desc: "ate3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port66 = &ConcretePort{Desc: "ate3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port67 = &ConcretePort{Desc: "ate3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port68 = &ConcretePort{Desc: "ate3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port69 = &ConcretePort{Desc: "ate3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comate3port70 = &ConcretePort{Desc: "ate3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{comate3port1, comate3port2, comate3port3, comate3port4, comate3port5, comate3port6, comate3port7, comate3port8, comate3port9, comate3port10, comate3port11, comate3port12, comate3port13, comate3port14, comate3port15, comate3port16, comate3port17, comate3port18, comate3port19, comate3port20, comate3port21, comate3port22, comate3port23, comate3port24, comate3port25, comate3port26, comate3port27, comate3port28, comate3port29, comate3port30, comate3port31, comate3port32, comate3port33, comate3port34, comate3port35, comate3port36, comate3port37, comate3port38, comate3port39, comate3port40, comate3port41, comate3port42, comate3port43, comate3port44, comate3port45, comate3port46, comate3port47, comate3port48, comate3port49, comate3port50, comate3port51, comate3port52, comate3port53, comate3port54, comate3port55, comate3port56, comate3port57, comate3port58, comate3port59, comate3port60, comate3port61, comate3port62, comate3port63, comate3port64, comate3port65, comate3port66, comate3port67, comate3port68, comate3port69, comate3port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	comate4port1  = &ConcretePort{Desc: "ate4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port2  = &ConcretePort{Desc: "ate4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port3  = &ConcretePort{Desc: "ate4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port4  = &ConcretePort{Desc: "ate4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port5  = &ConcretePort{Desc: "ate4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port6  = &ConcretePort{Desc: "ate4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port7  = &ConcretePort{Desc: "ate4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port8  = &ConcretePort{Desc: "ate4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port9  = &ConcretePort{Desc: "ate4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port10 = &ConcretePort{Desc: "ate4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port11 = &ConcretePort{Desc: "ate4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port12 = &ConcretePort{Desc: "ate4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port13 = &ConcretePort{Desc: "ate4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port14 = &ConcretePort{Desc: "ate4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port15 = &ConcretePort{Desc: "ate4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port16 = &ConcretePort{Desc: "ate4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port17 = &ConcretePort{Desc: "ate4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port18 = &ConcretePort{Desc: "ate4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port19 = &ConcretePort{Desc: "ate4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port20 = &ConcretePort{Desc: "ate4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port21 = &ConcretePort{Desc: "ate4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port22 = &ConcretePort{Desc: "ate4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port23 = &ConcretePort{Desc: "ate4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port24 = &ConcretePort{Desc: "ate4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port25 = &ConcretePort{Desc: "ate4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port26 = &ConcretePort{Desc: "ate4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port27 = &ConcretePort{Desc: "ate4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port28 = &ConcretePort{Desc: "ate4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port29 = &ConcretePort{Desc: "ate4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port30 = &ConcretePort{Desc: "ate4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port31 = &ConcretePort{Desc: "ate4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port32 = &ConcretePort{Desc: "ate4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port33 = &ConcretePort{Desc: "ate4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port34 = &ConcretePort{Desc: "ate4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port35 = &ConcretePort{Desc: "ate4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port36 = &ConcretePort{Desc: "ate4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port37 = &ConcretePort{Desc: "ate4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port38 = &ConcretePort{Desc: "ate4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port39 = &ConcretePort{Desc: "ate4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port40 = &ConcretePort{Desc: "ate4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port41 = &ConcretePort{Desc: "ate4:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port42 = &ConcretePort{Desc: "ate4:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port43 = &ConcretePort{Desc: "ate4:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port44 = &ConcretePort{Desc: "ate4:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port45 = &ConcretePort{Desc: "ate4:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port46 = &ConcretePort{Desc: "ate4:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port47 = &ConcretePort{Desc: "ate4:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port48 = &ConcretePort{Desc: "ate4:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port49 = &ConcretePort{Desc: "ate4:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port50 = &ConcretePort{Desc: "ate4:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port51 = &ConcretePort{Desc: "ate4:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port52 = &ConcretePort{Desc: "ate4:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port53 = &ConcretePort{Desc: "ate4:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port54 = &ConcretePort{Desc: "ate4:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port55 = &ConcretePort{Desc: "ate4:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port56 = &ConcretePort{Desc: "ate4:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port57 = &ConcretePort{Desc: "ate4:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port58 = &ConcretePort{Desc: "ate4:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port59 = &ConcretePort{Desc: "ate4:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port60 = &ConcretePort{Desc: "ate4:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port61 = &ConcretePort{Desc: "ate4:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port62 = &ConcretePort{Desc: "ate4:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port63 = &ConcretePort{Desc: "ate4:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port64 = &ConcretePort{Desc: "ate4:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port65 = &ConcretePort{Desc: "ate4:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port66 = &ConcretePort{Desc: "ate4:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port67 = &ConcretePort{Desc: "ate4:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port68 = &ConcretePort{Desc: "ate4:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port69 = &ConcretePort{Desc: "ate4:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comate4port70 = &ConcretePort{Desc: "ate4:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comate4       = &ConcreteNode{Desc: "ate4", Ports: []*ConcretePort{comate4port1, comate4port2, comate4port3, comate4port4, comate4port5, comate4port6, comate4port7, comate4port8, comate4port9, comate4port10, comate4port11, comate4port12, comate4port13, comate4port14, comate4port15, comate4port16, comate4port17, comate4port18, comate4port19, comate4port20, comate4port21, comate4port22, comate4port23, comate4port24, comate4port25, comate4port26, comate4port27, comate4port28, comate4port29, comate4port30, comate4port31, comate4port32, comate4port33, comate4port34, comate4port35, comate4port36, comate4port37, comate4port38, comate4port39, comate4port40, comate4port41, comate4port42, comate4port43, comate4port44, comate4port45, comate4port46, comate4port47, comate4port48, comate4port49, comate4port50, comate4port51, comate4port52, comate4port53, comate4port54, comate4port55, comate4port56, comate4port57, comate4port58, comate4port59, comate4port60, comate4port61, comate4port62, comate4port63, comate4port64, comate4port65, comate4port66, comate4port67, comate4port68, comate4port69, comate4port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	comate5port1  = &ConcretePort{Desc: "ate5:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port2  = &ConcretePort{Desc: "ate5:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port3  = &ConcretePort{Desc: "ate5:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port4  = &ConcretePort{Desc: "ate5:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port5  = &ConcretePort{Desc: "ate5:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port6  = &ConcretePort{Desc: "ate5:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port7  = &ConcretePort{Desc: "ate5:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port8  = &ConcretePort{Desc: "ate5:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port9  = &ConcretePort{Desc: "ate5:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port10 = &ConcretePort{Desc: "ate5:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port11 = &ConcretePort{Desc: "ate5:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port12 = &ConcretePort{Desc: "ate5:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port13 = &ConcretePort{Desc: "ate5:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port14 = &ConcretePort{Desc: "ate5:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port15 = &ConcretePort{Desc: "ate5:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port16 = &ConcretePort{Desc: "ate5:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port17 = &ConcretePort{Desc: "ate5:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port18 = &ConcretePort{Desc: "ate5:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port19 = &ConcretePort{Desc: "ate5:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port20 = &ConcretePort{Desc: "ate5:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port21 = &ConcretePort{Desc: "ate5:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port22 = &ConcretePort{Desc: "ate5:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port23 = &ConcretePort{Desc: "ate5:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port24 = &ConcretePort{Desc: "ate5:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port25 = &ConcretePort{Desc: "ate5:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port26 = &ConcretePort{Desc: "ate5:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port27 = &ConcretePort{Desc: "ate5:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port28 = &ConcretePort{Desc: "ate5:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port29 = &ConcretePort{Desc: "ate5:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port30 = &ConcretePort{Desc: "ate5:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port31 = &ConcretePort{Desc: "ate5:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port32 = &ConcretePort{Desc: "ate5:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port33 = &ConcretePort{Desc: "ate5:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port34 = &ConcretePort{Desc: "ate5:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port35 = &ConcretePort{Desc: "ate5:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port36 = &ConcretePort{Desc: "ate5:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port37 = &ConcretePort{Desc: "ate5:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port38 = &ConcretePort{Desc: "ate5:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port39 = &ConcretePort{Desc: "ate5:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port40 = &ConcretePort{Desc: "ate5:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port41 = &ConcretePort{Desc: "ate5:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port42 = &ConcretePort{Desc: "ate5:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port43 = &ConcretePort{Desc: "ate5:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port44 = &ConcretePort{Desc: "ate5:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port45 = &ConcretePort{Desc: "ate5:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port46 = &ConcretePort{Desc: "ate5:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port47 = &ConcretePort{Desc: "ate5:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port48 = &ConcretePort{Desc: "ate5:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port49 = &ConcretePort{Desc: "ate5:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port50 = &ConcretePort{Desc: "ate5:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port51 = &ConcretePort{Desc: "ate5:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port52 = &ConcretePort{Desc: "ate5:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port53 = &ConcretePort{Desc: "ate5:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port54 = &ConcretePort{Desc: "ate5:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port55 = &ConcretePort{Desc: "ate5:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port56 = &ConcretePort{Desc: "ate5:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port57 = &ConcretePort{Desc: "ate5:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port58 = &ConcretePort{Desc: "ate5:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port59 = &ConcretePort{Desc: "ate5:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port60 = &ConcretePort{Desc: "ate5:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port61 = &ConcretePort{Desc: "ate5:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port62 = &ConcretePort{Desc: "ate5:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port63 = &ConcretePort{Desc: "ate5:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port64 = &ConcretePort{Desc: "ate5:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port65 = &ConcretePort{Desc: "ate5:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port66 = &ConcretePort{Desc: "ate5:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port67 = &ConcretePort{Desc: "ate5:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port68 = &ConcretePort{Desc: "ate5:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port69 = &ConcretePort{Desc: "ate5:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comate5port70 = &ConcretePort{Desc: "ate5:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comate5       = &ConcreteNode{Desc: "ate5", Ports: []*ConcretePort{comate5port1, comate5port2, comate5port3, comate5port4, comate5port5, comate5port6, comate5port7, comate5port8, comate5port9, comate5port10, comate5port11, comate5port12, comate5port13, comate5port14, comate5port15, comate5port16, comate5port17, comate5port18, comate5port19, comate5port20, comate5port21, comate5port22, comate5port23, comate5port24, comate5port25, comate5port26, comate5port27, comate5port28, comate5port29, comate5port30, comate5port31, comate5port32, comate5port33, comate5port34, comate5port35, comate5port36, comate5port37, comate5port38, comate5port39, comate5port40, comate5port41, comate5port42, comate5port43, comate5port44, comate5port45, comate5port46, comate5port47, comate5port48, comate5port49, comate5port50, comate5port51, comate5port52, comate5port53, comate5port54, comate5port55, comate5port56, comate5port57, comate5port58, comate5port59, comate5port60, comate5port61, comate5port62, comate5port63, comate5port64, comate5port65, comate5port66, comate5port67, comate5port68, comate5port69, comate5port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	comate6port1  = &ConcretePort{Desc: "ate6:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port2  = &ConcretePort{Desc: "ate6:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port3  = &ConcretePort{Desc: "ate6:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port4  = &ConcretePort{Desc: "ate6:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port5  = &ConcretePort{Desc: "ate6:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port6  = &ConcretePort{Desc: "ate6:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port7  = &ConcretePort{Desc: "ate6:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port8  = &ConcretePort{Desc: "ate6:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port9  = &ConcretePort{Desc: "ate6:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port10 = &ConcretePort{Desc: "ate6:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port11 = &ConcretePort{Desc: "ate6:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port12 = &ConcretePort{Desc: "ate6:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port13 = &ConcretePort{Desc: "ate6:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port14 = &ConcretePort{Desc: "ate6:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port15 = &ConcretePort{Desc: "ate6:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port16 = &ConcretePort{Desc: "ate6:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port17 = &ConcretePort{Desc: "ate6:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port18 = &ConcretePort{Desc: "ate6:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port19 = &ConcretePort{Desc: "ate6:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port20 = &ConcretePort{Desc: "ate6:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port21 = &ConcretePort{Desc: "ate6:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port22 = &ConcretePort{Desc: "ate6:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port23 = &ConcretePort{Desc: "ate6:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port24 = &ConcretePort{Desc: "ate6:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port25 = &ConcretePort{Desc: "ate6:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port26 = &ConcretePort{Desc: "ate6:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port27 = &ConcretePort{Desc: "ate6:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port28 = &ConcretePort{Desc: "ate6:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port29 = &ConcretePort{Desc: "ate6:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port30 = &ConcretePort{Desc: "ate6:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port31 = &ConcretePort{Desc: "ate6:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port32 = &ConcretePort{Desc: "ate6:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port33 = &ConcretePort{Desc: "ate6:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port34 = &ConcretePort{Desc: "ate6:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port35 = &ConcretePort{Desc: "ate6:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port36 = &ConcretePort{Desc: "ate6:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port37 = &ConcretePort{Desc: "ate6:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port38 = &ConcretePort{Desc: "ate6:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port39 = &ConcretePort{Desc: "ate6:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port40 = &ConcretePort{Desc: "ate6:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port41 = &ConcretePort{Desc: "ate6:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port42 = &ConcretePort{Desc: "ate6:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port43 = &ConcretePort{Desc: "ate6:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port44 = &ConcretePort{Desc: "ate6:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port45 = &ConcretePort{Desc: "ate6:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port46 = &ConcretePort{Desc: "ate6:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port47 = &ConcretePort{Desc: "ate6:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port48 = &ConcretePort{Desc: "ate6:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port49 = &ConcretePort{Desc: "ate6:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port50 = &ConcretePort{Desc: "ate6:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port51 = &ConcretePort{Desc: "ate6:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port52 = &ConcretePort{Desc: "ate6:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port53 = &ConcretePort{Desc: "ate6:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port54 = &ConcretePort{Desc: "ate6:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port55 = &ConcretePort{Desc: "ate6:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port56 = &ConcretePort{Desc: "ate6:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port57 = &ConcretePort{Desc: "ate6:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port58 = &ConcretePort{Desc: "ate6:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port59 = &ConcretePort{Desc: "ate6:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port60 = &ConcretePort{Desc: "ate6:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port61 = &ConcretePort{Desc: "ate6:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port62 = &ConcretePort{Desc: "ate6:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port63 = &ConcretePort{Desc: "ate6:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port64 = &ConcretePort{Desc: "ate6:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port65 = &ConcretePort{Desc: "ate6:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port66 = &ConcretePort{Desc: "ate6:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port67 = &ConcretePort{Desc: "ate6:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port68 = &ConcretePort{Desc: "ate6:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port69 = &ConcretePort{Desc: "ate6:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comate6port70 = &ConcretePort{Desc: "ate6:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comate6       = &ConcreteNode{Desc: "ate6", Ports: []*ConcretePort{comate6port1, comate6port2, comate6port3, comate6port4, comate6port5, comate6port6, comate6port7, comate6port8, comate6port9, comate6port10, comate6port11, comate6port12, comate6port13, comate6port14, comate6port15, comate6port16, comate6port17, comate6port18, comate6port19, comate6port20, comate6port21, comate6port22, comate6port23, comate6port24, comate6port25, comate6port26, comate6port27, comate6port28, comate6port29, comate6port30, comate6port31, comate6port32, comate6port33, comate6port34, comate6port35, comate6port36, comate6port37, comate6port38, comate6port39, comate6port40, comate6port41, comate6port42, comate6port43, comate6port44, comate6port45, comate6port46, comate6port47, comate6port48, comate6port49, comate6port50, comate6port51, comate6port52, comate6port53, comate6port54, comate6port55, comate6port56, comate6port57, comate6port58, comate6port59, comate6port60, comate6port61, comate6port62, comate6port63, comate6port64, comate6port65, comate6port66, comate6port67, comate6port68, comate6port69, comate6port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	comswitch1port1   = &ConcretePort{Desc: "switch1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port2   = &ConcretePort{Desc: "switch1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port3   = &ConcretePort{Desc: "switch1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port4   = &ConcretePort{Desc: "switch1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port5   = &ConcretePort{Desc: "switch1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port6   = &ConcretePort{Desc: "switch1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port7   = &ConcretePort{Desc: "switch1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port8   = &ConcretePort{Desc: "switch1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port9   = &ConcretePort{Desc: "switch1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port10  = &ConcretePort{Desc: "switch1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port11  = &ConcretePort{Desc: "switch1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port12  = &ConcretePort{Desc: "switch1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port13  = &ConcretePort{Desc: "switch1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port14  = &ConcretePort{Desc: "switch1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port15  = &ConcretePort{Desc: "switch1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port16  = &ConcretePort{Desc: "switch1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port17  = &ConcretePort{Desc: "switch1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port18  = &ConcretePort{Desc: "switch1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port19  = &ConcretePort{Desc: "switch1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port20  = &ConcretePort{Desc: "switch1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port21  = &ConcretePort{Desc: "switch1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port22  = &ConcretePort{Desc: "switch1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port23  = &ConcretePort{Desc: "switch1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port24  = &ConcretePort{Desc: "switch1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port25  = &ConcretePort{Desc: "switch1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port26  = &ConcretePort{Desc: "switch1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port27  = &ConcretePort{Desc: "switch1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port28  = &ConcretePort{Desc: "switch1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port29  = &ConcretePort{Desc: "switch1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port30  = &ConcretePort{Desc: "switch1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port31  = &ConcretePort{Desc: "switch1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port32  = &ConcretePort{Desc: "switch1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port33  = &ConcretePort{Desc: "switch1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port34  = &ConcretePort{Desc: "switch1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port35  = &ConcretePort{Desc: "switch1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port36  = &ConcretePort{Desc: "switch1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port37  = &ConcretePort{Desc: "switch1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port38  = &ConcretePort{Desc: "switch1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port39  = &ConcretePort{Desc: "switch1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port40  = &ConcretePort{Desc: "switch1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port41  = &ConcretePort{Desc: "switch1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port42  = &ConcretePort{Desc: "switch1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port43  = &ConcretePort{Desc: "switch1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port44  = &ConcretePort{Desc: "switch1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port45  = &ConcretePort{Desc: "switch1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port46  = &ConcretePort{Desc: "switch1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port47  = &ConcretePort{Desc: "switch1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port48  = &ConcretePort{Desc: "switch1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port49  = &ConcretePort{Desc: "switch1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port50  = &ConcretePort{Desc: "switch1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port51  = &ConcretePort{Desc: "switch1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port52  = &ConcretePort{Desc: "switch1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port53  = &ConcretePort{Desc: "switch1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port54  = &ConcretePort{Desc: "switch1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port55  = &ConcretePort{Desc: "switch1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port56  = &ConcretePort{Desc: "switch1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port57  = &ConcretePort{Desc: "switch1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port58  = &ConcretePort{Desc: "switch1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port59  = &ConcretePort{Desc: "switch1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port60  = &ConcretePort{Desc: "switch1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port61  = &ConcretePort{Desc: "switch1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port62  = &ConcretePort{Desc: "switch1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port63  = &ConcretePort{Desc: "switch1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port64  = &ConcretePort{Desc: "switch1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port65  = &ConcretePort{Desc: "switch1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port66  = &ConcretePort{Desc: "switch1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port67  = &ConcretePort{Desc: "switch1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port68  = &ConcretePort{Desc: "switch1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port69  = &ConcretePort{Desc: "switch1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port70  = &ConcretePort{Desc: "switch1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port71  = &ConcretePort{Desc: "switch1:port71", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port72  = &ConcretePort{Desc: "switch1:port72", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port73  = &ConcretePort{Desc: "switch1:port73", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port74  = &ConcretePort{Desc: "switch1:port74", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port75  = &ConcretePort{Desc: "switch1:port75", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port76  = &ConcretePort{Desc: "switch1:port76", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port77  = &ConcretePort{Desc: "switch1:port77", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port78  = &ConcretePort{Desc: "switch1:port78", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port79  = &ConcretePort{Desc: "switch1:port79", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port80  = &ConcretePort{Desc: "switch1:port80", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port81  = &ConcretePort{Desc: "switch1:port81", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port82  = &ConcretePort{Desc: "switch1:port82", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port83  = &ConcretePort{Desc: "switch1:port83", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port84  = &ConcretePort{Desc: "switch1:port84", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port85  = &ConcretePort{Desc: "switch1:port85", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port86  = &ConcretePort{Desc: "switch1:port86", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port87  = &ConcretePort{Desc: "switch1:port87", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port88  = &ConcretePort{Desc: "switch1:port88", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port89  = &ConcretePort{Desc: "switch1:port89", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port90  = &ConcretePort{Desc: "switch1:port90", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port91  = &ConcretePort{Desc: "switch1:port91", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port92  = &ConcretePort{Desc: "switch1:port92", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port93  = &ConcretePort{Desc: "switch1:port93", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port94  = &ConcretePort{Desc: "switch1:port94", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port95  = &ConcretePort{Desc: "switch1:port95", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port96  = &ConcretePort{Desc: "switch1:port96", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port97  = &ConcretePort{Desc: "switch1:port97", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port98  = &ConcretePort{Desc: "switch1:port98", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port99  = &ConcretePort{Desc: "switch1:port99", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port100 = &ConcretePort{Desc: "switch1:port100", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port101 = &ConcretePort{Desc: "switch1:port101", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port102 = &ConcretePort{Desc: "switch1:port102", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port103 = &ConcretePort{Desc: "switch1:port103", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port104 = &ConcretePort{Desc: "switch1:port104", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port105 = &ConcretePort{Desc: "switch1:port105", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port106 = &ConcretePort{Desc: "switch1:port106", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port107 = &ConcretePort{Desc: "switch1:port107", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port108 = &ConcretePort{Desc: "switch1:port108", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port109 = &ConcretePort{Desc: "switch1:port109", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port110 = &ConcretePort{Desc: "switch1:port110", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port111 = &ConcretePort{Desc: "switch1:port111", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port112 = &ConcretePort{Desc: "switch1:port112", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port113 = &ConcretePort{Desc: "switch1:port113", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port114 = &ConcretePort{Desc: "switch1:port114", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port115 = &ConcretePort{Desc: "switch1:port115", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port116 = &ConcretePort{Desc: "switch1:port116", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port117 = &ConcretePort{Desc: "switch1:port117", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port118 = &ConcretePort{Desc: "switch1:port118", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port119 = &ConcretePort{Desc: "switch1:port119", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port120 = &ConcretePort{Desc: "switch1:port120", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port121 = &ConcretePort{Desc: "switch1:port121", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port122 = &ConcretePort{Desc: "switch1:port122", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port123 = &ConcretePort{Desc: "switch1:port123", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port124 = &ConcretePort{Desc: "switch1:port124", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port125 = &ConcretePort{Desc: "switch1:port125", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port126 = &ConcretePort{Desc: "switch1:port126", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port127 = &ConcretePort{Desc: "switch1:port127", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port128 = &ConcretePort{Desc: "switch1:port128", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port129 = &ConcretePort{Desc: "switch1:port129", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port130 = &ConcretePort{Desc: "switch1:port130", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port131 = &ConcretePort{Desc: "switch1:port131", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port132 = &ConcretePort{Desc: "switch1:port132", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port133 = &ConcretePort{Desc: "switch1:port133", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port134 = &ConcretePort{Desc: "switch1:port134", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port135 = &ConcretePort{Desc: "switch1:port135", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port136 = &ConcretePort{Desc: "switch1:port136", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port137 = &ConcretePort{Desc: "switch1:port137", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port138 = &ConcretePort{Desc: "switch1:port138", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port139 = &ConcretePort{Desc: "switch1:port139", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port140 = &ConcretePort{Desc: "switch1:port140", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch1port141 = &ConcretePort{Desc: "switch1:port141", Attrs: map[string]string{"speed": "S_400G"}}

	comswitch1 = &ConcreteNode{Desc: "switch1", Ports: []*ConcretePort{
		comswitch1port1,
		comswitch1port2,
		comswitch1port3,
		comswitch1port4,
		comswitch1port5,
		comswitch1port6,
		comswitch1port7,
		comswitch1port8,
		comswitch1port9,
		comswitch1port10,
		comswitch1port11,
		comswitch1port12,
		comswitch1port13,
		comswitch1port14,
		comswitch1port15,
		comswitch1port16,
		comswitch1port17,
		comswitch1port18,
		comswitch1port19,
		comswitch1port20,
		comswitch1port21,
		comswitch1port22,
		comswitch1port23,
		comswitch1port24,
		comswitch1port25,
		comswitch1port26,
		comswitch1port27,
		comswitch1port28,
		comswitch1port29,
		comswitch1port30,
		comswitch1port31,
		comswitch1port32,
		comswitch1port33,
		comswitch1port34,
		comswitch1port35,
		comswitch1port36,
		comswitch1port37,
		comswitch1port38,
		comswitch1port39,
		comswitch1port40,
		comswitch1port41,
		comswitch1port42,
		comswitch1port43,
		comswitch1port44,
		comswitch1port45,
		comswitch1port46,
		comswitch1port47,
		comswitch1port48,
		comswitch1port49,
		comswitch1port50,
		comswitch1port51,
		comswitch1port52,
		comswitch1port53,
		comswitch1port54,
		comswitch1port55,
		comswitch1port56,
		comswitch1port57,
		comswitch1port58,
		comswitch1port59,
		comswitch1port60,
		comswitch1port61,
		comswitch1port62,
		comswitch1port63,
		comswitch1port64,
		comswitch1port65,
		comswitch1port66,
		comswitch1port67,
		comswitch1port68,
		comswitch1port69,
		comswitch1port70,
		comswitch1port71,
		comswitch1port72,
		comswitch1port73,
		comswitch1port74,
		comswitch1port75,
		comswitch1port76,
		comswitch1port77,
		comswitch1port78,
		comswitch1port79,
		comswitch1port80,
		comswitch1port81,
		comswitch1port82,
		comswitch1port83,
		comswitch1port84,
		comswitch1port85,
		comswitch1port86,
		comswitch1port87,
		comswitch1port88,
		comswitch1port89,
		comswitch1port90,
		comswitch1port91,
		comswitch1port92,
		comswitch1port93,
		comswitch1port94,
		comswitch1port95,
		comswitch1port96,
		comswitch1port97,
		comswitch1port98,
		comswitch1port99,
		comswitch1port100,
		comswitch1port101,
		comswitch1port102,
		comswitch1port103,
		comswitch1port104,
		comswitch1port105,
		comswitch1port106,
		comswitch1port107,
		comswitch1port108,
		comswitch1port109,
		comswitch1port110,
		comswitch1port111,
		comswitch1port112,
		comswitch1port113,
		comswitch1port114,
		comswitch1port115,
		comswitch1port116,
		comswitch1port117,
		comswitch1port118,
		comswitch1port119,
		comswitch1port120,
		comswitch1port121,
		comswitch1port122,
		comswitch1port123,
		comswitch1port124,
		comswitch1port125,
		comswitch1port126,
		comswitch1port127,
		comswitch1port128,
		comswitch1port129,
		comswitch1port130,
		comswitch1port131,
		comswitch1port132,
		comswitch1port133,
		comswitch1port134,
		comswitch1port135,
		comswitch1port136,
		comswitch1port137,
		comswitch1port138,
		comswitch1port139,
		comswitch1port140,
		comswitch1port141}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	comswitch2port1   = &ConcretePort{Desc: "switch2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port2   = &ConcretePort{Desc: "switch2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port3   = &ConcretePort{Desc: "switch2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port4   = &ConcretePort{Desc: "switch2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port5   = &ConcretePort{Desc: "switch2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port6   = &ConcretePort{Desc: "switch2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port7   = &ConcretePort{Desc: "switch2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port8   = &ConcretePort{Desc: "switch2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port9   = &ConcretePort{Desc: "switch2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port10  = &ConcretePort{Desc: "switch2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port11  = &ConcretePort{Desc: "switch2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port12  = &ConcretePort{Desc: "switch2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port13  = &ConcretePort{Desc: "switch2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port14  = &ConcretePort{Desc: "switch2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port15  = &ConcretePort{Desc: "switch2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port16  = &ConcretePort{Desc: "switch2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port17  = &ConcretePort{Desc: "switch2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port18  = &ConcretePort{Desc: "switch2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port19  = &ConcretePort{Desc: "switch2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port20  = &ConcretePort{Desc: "switch2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port21  = &ConcretePort{Desc: "switch2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port22  = &ConcretePort{Desc: "switch2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port23  = &ConcretePort{Desc: "switch2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port24  = &ConcretePort{Desc: "switch2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port25  = &ConcretePort{Desc: "switch2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port26  = &ConcretePort{Desc: "switch2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port27  = &ConcretePort{Desc: "switch2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port28  = &ConcretePort{Desc: "switch2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port29  = &ConcretePort{Desc: "switch2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port30  = &ConcretePort{Desc: "switch2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port31  = &ConcretePort{Desc: "switch2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port32  = &ConcretePort{Desc: "switch2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port33  = &ConcretePort{Desc: "switch2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port34  = &ConcretePort{Desc: "switch2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port35  = &ConcretePort{Desc: "switch2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port36  = &ConcretePort{Desc: "switch2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port37  = &ConcretePort{Desc: "switch2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port38  = &ConcretePort{Desc: "switch2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port39  = &ConcretePort{Desc: "switch2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port40  = &ConcretePort{Desc: "switch2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port41  = &ConcretePort{Desc: "switch2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port42  = &ConcretePort{Desc: "switch2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port43  = &ConcretePort{Desc: "switch2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port44  = &ConcretePort{Desc: "switch2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port45  = &ConcretePort{Desc: "switch2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port46  = &ConcretePort{Desc: "switch2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port47  = &ConcretePort{Desc: "switch2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port48  = &ConcretePort{Desc: "switch2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port49  = &ConcretePort{Desc: "switch2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port50  = &ConcretePort{Desc: "switch2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port51  = &ConcretePort{Desc: "switch2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port52  = &ConcretePort{Desc: "switch2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port53  = &ConcretePort{Desc: "switch2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port54  = &ConcretePort{Desc: "switch2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port55  = &ConcretePort{Desc: "switch2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port56  = &ConcretePort{Desc: "switch2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port57  = &ConcretePort{Desc: "switch2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port58  = &ConcretePort{Desc: "switch2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port59  = &ConcretePort{Desc: "switch2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port60  = &ConcretePort{Desc: "switch2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port61  = &ConcretePort{Desc: "switch2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port62  = &ConcretePort{Desc: "switch2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port63  = &ConcretePort{Desc: "switch2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port64  = &ConcretePort{Desc: "switch2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port65  = &ConcretePort{Desc: "switch2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port66  = &ConcretePort{Desc: "switch2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port67  = &ConcretePort{Desc: "switch2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port68  = &ConcretePort{Desc: "switch2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port69  = &ConcretePort{Desc: "switch2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port70  = &ConcretePort{Desc: "switch2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port71  = &ConcretePort{Desc: "switch2:port71", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port72  = &ConcretePort{Desc: "switch2:port72", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port73  = &ConcretePort{Desc: "switch2:port73", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port74  = &ConcretePort{Desc: "switch2:port74", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port75  = &ConcretePort{Desc: "switch2:port75", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port76  = &ConcretePort{Desc: "switch2:port76", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port77  = &ConcretePort{Desc: "switch2:port77", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port78  = &ConcretePort{Desc: "switch2:port78", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port79  = &ConcretePort{Desc: "switch2:port79", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port80  = &ConcretePort{Desc: "switch2:port80", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port81  = &ConcretePort{Desc: "switch2:port81", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port82  = &ConcretePort{Desc: "switch2:port82", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port83  = &ConcretePort{Desc: "switch2:port83", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port84  = &ConcretePort{Desc: "switch2:port84", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port85  = &ConcretePort{Desc: "switch2:port85", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port86  = &ConcretePort{Desc: "switch2:port86", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port87  = &ConcretePort{Desc: "switch2:port87", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port88  = &ConcretePort{Desc: "switch2:port88", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port89  = &ConcretePort{Desc: "switch2:port89", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port90  = &ConcretePort{Desc: "switch2:port90", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port91  = &ConcretePort{Desc: "switch2:port91", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port92  = &ConcretePort{Desc: "switch2:port92", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port93  = &ConcretePort{Desc: "switch2:port93", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port94  = &ConcretePort{Desc: "switch2:port94", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port95  = &ConcretePort{Desc: "switch2:port95", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port96  = &ConcretePort{Desc: "switch2:port96", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port97  = &ConcretePort{Desc: "switch2:port97", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port98  = &ConcretePort{Desc: "switch2:port98", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port99  = &ConcretePort{Desc: "switch2:port99", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port100 = &ConcretePort{Desc: "switch2:port100", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port101 = &ConcretePort{Desc: "switch2:port101", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port102 = &ConcretePort{Desc: "switch2:port102", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port103 = &ConcretePort{Desc: "switch2:port103", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port104 = &ConcretePort{Desc: "switch2:port104", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port105 = &ConcretePort{Desc: "switch2:port105", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port106 = &ConcretePort{Desc: "switch2:port106", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port107 = &ConcretePort{Desc: "switch2:port107", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port108 = &ConcretePort{Desc: "switch2:port108", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port109 = &ConcretePort{Desc: "switch2:port109", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port110 = &ConcretePort{Desc: "switch2:port110", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port111 = &ConcretePort{Desc: "switch2:port111", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port112 = &ConcretePort{Desc: "switch2:port112", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port113 = &ConcretePort{Desc: "switch2:port113", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port114 = &ConcretePort{Desc: "switch2:port114", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port115 = &ConcretePort{Desc: "switch2:port115", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port116 = &ConcretePort{Desc: "switch2:port116", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port117 = &ConcretePort{Desc: "switch2:port117", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port118 = &ConcretePort{Desc: "switch2:port118", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port119 = &ConcretePort{Desc: "switch2:port119", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port120 = &ConcretePort{Desc: "switch2:port120", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port121 = &ConcretePort{Desc: "switch2:port121", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port122 = &ConcretePort{Desc: "switch2:port122", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port123 = &ConcretePort{Desc: "switch2:port123", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port124 = &ConcretePort{Desc: "switch2:port124", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port125 = &ConcretePort{Desc: "switch2:port125", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port126 = &ConcretePort{Desc: "switch2:port126", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port127 = &ConcretePort{Desc: "switch2:port127", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port128 = &ConcretePort{Desc: "switch2:port128", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port129 = &ConcretePort{Desc: "switch2:port129", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port130 = &ConcretePort{Desc: "switch2:port130", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port131 = &ConcretePort{Desc: "switch2:port131", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port132 = &ConcretePort{Desc: "switch2:port132", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port133 = &ConcretePort{Desc: "switch2:port133", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port134 = &ConcretePort{Desc: "switch2:port134", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port135 = &ConcretePort{Desc: "switch2:port135", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port136 = &ConcretePort{Desc: "switch2:port136", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port137 = &ConcretePort{Desc: "switch2:port137", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port138 = &ConcretePort{Desc: "switch2:port138", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port139 = &ConcretePort{Desc: "switch2:port139", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port140 = &ConcretePort{Desc: "switch2:port140", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port141 = &ConcretePort{Desc: "switch2:port141", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch2port142 = &ConcretePort{Desc: "switch2:port142", Attrs: map[string]string{"speed": "S_400G"}}

	comswitch2 = &ConcreteNode{Desc: "switch2", Ports: []*ConcretePort{
		comswitch2port1,
		comswitch2port2,
		comswitch2port3,
		comswitch2port4,
		comswitch2port5,
		comswitch2port6,
		comswitch2port7,
		comswitch2port8,
		comswitch2port9,
		comswitch2port10,
		comswitch2port11,
		comswitch2port12,
		comswitch2port13,
		comswitch2port14,
		comswitch2port15,
		comswitch2port16,
		comswitch2port17,
		comswitch2port18,
		comswitch2port19,
		comswitch2port20,
		comswitch2port21,
		comswitch2port22,
		comswitch2port23,
		comswitch2port24,
		comswitch2port25,
		comswitch2port26,
		comswitch2port27,
		comswitch2port28,
		comswitch2port29,
		comswitch2port30,
		comswitch2port31,
		comswitch2port32,
		comswitch2port33,
		comswitch2port34,
		comswitch2port35,
		comswitch2port36,
		comswitch2port37,
		comswitch2port38,
		comswitch2port39,
		comswitch2port40,
		comswitch2port41,
		comswitch2port42,
		comswitch2port43,
		comswitch2port44,
		comswitch2port45,
		comswitch2port46,
		comswitch2port47,
		comswitch2port48,
		comswitch2port49,
		comswitch2port50,
		comswitch2port51,
		comswitch2port52,
		comswitch2port53,
		comswitch2port54,
		comswitch2port55,
		comswitch2port56,
		comswitch2port57,
		comswitch2port58,
		comswitch2port59,
		comswitch2port60,
		comswitch2port61,
		comswitch2port62,
		comswitch2port63,
		comswitch2port64,
		comswitch2port65,
		comswitch2port66,
		comswitch2port67,
		comswitch2port68,
		comswitch2port69,
		comswitch2port70,
		comswitch2port71,
		comswitch2port72,
		comswitch2port73,
		comswitch2port74,
		comswitch2port75,
		comswitch2port76,
		comswitch2port77,
		comswitch2port78,
		comswitch2port79,
		comswitch2port80,
		comswitch2port81,
		comswitch2port82,
		comswitch2port83,
		comswitch2port84,
		comswitch2port85,
		comswitch2port86,
		comswitch2port87,
		comswitch2port88,
		comswitch2port89,
		comswitch2port90,
		comswitch2port91,
		comswitch2port92,
		comswitch2port93,
		comswitch2port94,
		comswitch2port95,
		comswitch2port96,
		comswitch2port97,
		comswitch2port98,
		comswitch2port99,
		comswitch2port100,
		comswitch2port101,
		comswitch2port102,
		comswitch2port103,
		comswitch2port104,
		comswitch2port105,
		comswitch2port106,
		comswitch2port107,
		comswitch2port108,
		comswitch2port109,
		comswitch2port110,
		comswitch2port111,
		comswitch2port112,
		comswitch2port113,
		comswitch2port114,
		comswitch2port115,
		comswitch2port116,
		comswitch2port117,
		comswitch2port118,
		comswitch2port119,
		comswitch2port120,
		comswitch2port121,
		comswitch2port122,
		comswitch2port123,
		comswitch2port124,
		comswitch2port125,
		comswitch2port126,
		comswitch2port127,
		comswitch2port128,
		comswitch2port129,
		comswitch2port130,
		comswitch2port131,
		comswitch2port132,
		comswitch2port133,
		comswitch2port134,
		comswitch2port135,
		comswitch2port136,
		comswitch2port137,
		comswitch2port138,
		comswitch2port139,
		comswitch2port140,
		comswitch2port141,
		comswitch2port142}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	comswitch3port1   = &ConcretePort{Desc: "switch3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port2   = &ConcretePort{Desc: "switch3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port3   = &ConcretePort{Desc: "switch3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port4   = &ConcretePort{Desc: "switch3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port5   = &ConcretePort{Desc: "switch3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port6   = &ConcretePort{Desc: "switch3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port7   = &ConcretePort{Desc: "switch3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port8   = &ConcretePort{Desc: "switch3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port9   = &ConcretePort{Desc: "switch3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port10  = &ConcretePort{Desc: "switch3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port11  = &ConcretePort{Desc: "switch3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port12  = &ConcretePort{Desc: "switch3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port13  = &ConcretePort{Desc: "switch3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port14  = &ConcretePort{Desc: "switch3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port15  = &ConcretePort{Desc: "switch3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port16  = &ConcretePort{Desc: "switch3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port17  = &ConcretePort{Desc: "switch3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port18  = &ConcretePort{Desc: "switch3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port19  = &ConcretePort{Desc: "switch3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port20  = &ConcretePort{Desc: "switch3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port21  = &ConcretePort{Desc: "switch3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port22  = &ConcretePort{Desc: "switch3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port23  = &ConcretePort{Desc: "switch3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port24  = &ConcretePort{Desc: "switch3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port25  = &ConcretePort{Desc: "switch3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port26  = &ConcretePort{Desc: "switch3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port27  = &ConcretePort{Desc: "switch3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port28  = &ConcretePort{Desc: "switch3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port29  = &ConcretePort{Desc: "switch3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port30  = &ConcretePort{Desc: "switch3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port31  = &ConcretePort{Desc: "switch3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port32  = &ConcretePort{Desc: "switch3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port33  = &ConcretePort{Desc: "switch3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port34  = &ConcretePort{Desc: "switch3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port35  = &ConcretePort{Desc: "switch3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port36  = &ConcretePort{Desc: "switch3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port37  = &ConcretePort{Desc: "switch3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port38  = &ConcretePort{Desc: "switch3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port39  = &ConcretePort{Desc: "switch3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port40  = &ConcretePort{Desc: "switch3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port41  = &ConcretePort{Desc: "switch3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port42  = &ConcretePort{Desc: "switch3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port43  = &ConcretePort{Desc: "switch3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port44  = &ConcretePort{Desc: "switch3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port45  = &ConcretePort{Desc: "switch3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port46  = &ConcretePort{Desc: "switch3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port47  = &ConcretePort{Desc: "switch3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port48  = &ConcretePort{Desc: "switch3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port49  = &ConcretePort{Desc: "switch3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port50  = &ConcretePort{Desc: "switch3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port51  = &ConcretePort{Desc: "switch3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port52  = &ConcretePort{Desc: "switch3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port53  = &ConcretePort{Desc: "switch3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port54  = &ConcretePort{Desc: "switch3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port55  = &ConcretePort{Desc: "switch3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port56  = &ConcretePort{Desc: "switch3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port57  = &ConcretePort{Desc: "switch3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port58  = &ConcretePort{Desc: "switch3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port59  = &ConcretePort{Desc: "switch3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port60  = &ConcretePort{Desc: "switch3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port61  = &ConcretePort{Desc: "switch3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port62  = &ConcretePort{Desc: "switch3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port63  = &ConcretePort{Desc: "switch3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port64  = &ConcretePort{Desc: "switch3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port65  = &ConcretePort{Desc: "switch3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port66  = &ConcretePort{Desc: "switch3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port67  = &ConcretePort{Desc: "switch3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port68  = &ConcretePort{Desc: "switch3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port69  = &ConcretePort{Desc: "switch3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port70  = &ConcretePort{Desc: "switch3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port71  = &ConcretePort{Desc: "switch3:port71", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port72  = &ConcretePort{Desc: "switch3:port72", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port73  = &ConcretePort{Desc: "switch3:port73", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port74  = &ConcretePort{Desc: "switch3:port74", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port75  = &ConcretePort{Desc: "switch3:port75", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port76  = &ConcretePort{Desc: "switch3:port76", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port77  = &ConcretePort{Desc: "switch3:port77", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port78  = &ConcretePort{Desc: "switch3:port78", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port79  = &ConcretePort{Desc: "switch3:port79", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port80  = &ConcretePort{Desc: "switch3:port80", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port81  = &ConcretePort{Desc: "switch3:port81", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port82  = &ConcretePort{Desc: "switch3:port82", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port83  = &ConcretePort{Desc: "switch3:port83", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port84  = &ConcretePort{Desc: "switch3:port84", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port85  = &ConcretePort{Desc: "switch3:port85", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port86  = &ConcretePort{Desc: "switch3:port86", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port87  = &ConcretePort{Desc: "switch3:port87", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port88  = &ConcretePort{Desc: "switch3:port88", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port89  = &ConcretePort{Desc: "switch3:port89", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port90  = &ConcretePort{Desc: "switch3:port90", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port91  = &ConcretePort{Desc: "switch3:port91", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port92  = &ConcretePort{Desc: "switch3:port92", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port93  = &ConcretePort{Desc: "switch3:port93", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port94  = &ConcretePort{Desc: "switch3:port94", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port95  = &ConcretePort{Desc: "switch3:port95", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port96  = &ConcretePort{Desc: "switch3:port96", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port97  = &ConcretePort{Desc: "switch3:port97", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port98  = &ConcretePort{Desc: "switch3:port98", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port99  = &ConcretePort{Desc: "switch3:port99", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port100 = &ConcretePort{Desc: "switch3:port100", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port101 = &ConcretePort{Desc: "switch3:port101", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port102 = &ConcretePort{Desc: "switch3:port102", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port103 = &ConcretePort{Desc: "switch3:port103", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port104 = &ConcretePort{Desc: "switch3:port104", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port105 = &ConcretePort{Desc: "switch3:port105", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port106 = &ConcretePort{Desc: "switch3:port106", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port107 = &ConcretePort{Desc: "switch3:port107", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port108 = &ConcretePort{Desc: "switch3:port108", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port109 = &ConcretePort{Desc: "switch3:port109", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port110 = &ConcretePort{Desc: "switch3:port110", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port111 = &ConcretePort{Desc: "switch3:port111", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port112 = &ConcretePort{Desc: "switch3:port112", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port113 = &ConcretePort{Desc: "switch3:port113", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port114 = &ConcretePort{Desc: "switch3:port114", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port115 = &ConcretePort{Desc: "switch3:port115", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port116 = &ConcretePort{Desc: "switch3:port116", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port117 = &ConcretePort{Desc: "switch3:port117", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port118 = &ConcretePort{Desc: "switch3:port118", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port119 = &ConcretePort{Desc: "switch3:port119", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port120 = &ConcretePort{Desc: "switch3:port120", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port121 = &ConcretePort{Desc: "switch3:port121", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port122 = &ConcretePort{Desc: "switch3:port122", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port123 = &ConcretePort{Desc: "switch3:port123", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port124 = &ConcretePort{Desc: "switch3:port124", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port125 = &ConcretePort{Desc: "switch3:port125", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port126 = &ConcretePort{Desc: "switch3:port126", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port127 = &ConcretePort{Desc: "switch3:port127", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port128 = &ConcretePort{Desc: "switch3:port128", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port129 = &ConcretePort{Desc: "switch3:port129", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port130 = &ConcretePort{Desc: "switch3:port130", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port131 = &ConcretePort{Desc: "switch3:port131", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port132 = &ConcretePort{Desc: "switch3:port132", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port133 = &ConcretePort{Desc: "switch3:port133", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port134 = &ConcretePort{Desc: "switch3:port134", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port135 = &ConcretePort{Desc: "switch3:port135", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port136 = &ConcretePort{Desc: "switch3:port136", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port137 = &ConcretePort{Desc: "switch3:port137", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port138 = &ConcretePort{Desc: "switch3:port138", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port139 = &ConcretePort{Desc: "switch3:port139", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port140 = &ConcretePort{Desc: "switch3:port140", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port141 = &ConcretePort{Desc: "switch3:port141", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch3port142 = &ConcretePort{Desc: "switch3:port142", Attrs: map[string]string{"speed": "S_400G"}}

	comswitch3 = &ConcreteNode{Desc: "switch3", Ports: []*ConcretePort{
		comswitch3port1,
		comswitch3port2,
		comswitch3port3,
		comswitch3port4,
		comswitch3port5,
		comswitch3port6,
		comswitch3port7,
		comswitch3port8,
		comswitch3port9,
		comswitch3port10,
		comswitch3port11,
		comswitch3port12,
		comswitch3port13,
		comswitch3port14,
		comswitch3port15,
		comswitch3port16,
		comswitch3port17,
		comswitch3port18,
		comswitch3port19,
		comswitch3port20,
		comswitch3port21,
		comswitch3port22,
		comswitch3port23,
		comswitch3port24,
		comswitch3port25,
		comswitch3port26,
		comswitch3port27,
		comswitch3port28,
		comswitch3port29,
		comswitch3port30,
		comswitch3port31,
		comswitch3port32,
		comswitch3port33,
		comswitch3port34,
		comswitch3port35,
		comswitch3port36,
		comswitch3port37,
		comswitch3port38,
		comswitch3port39,
		comswitch3port40,
		comswitch3port41,
		comswitch3port42,
		comswitch3port43,
		comswitch3port44,
		comswitch3port45,
		comswitch3port46,
		comswitch3port47,
		comswitch3port48,
		comswitch3port49,
		comswitch3port50,
		comswitch3port51,
		comswitch3port52,
		comswitch3port53,
		comswitch3port54,
		comswitch3port55,
		comswitch3port56,
		comswitch3port57,
		comswitch3port58,
		comswitch3port59,
		comswitch3port60,
		comswitch3port61,
		comswitch3port62,
		comswitch3port63,
		comswitch3port64,
		comswitch3port65,
		comswitch3port66,
		comswitch3port67,
		comswitch3port68,
		comswitch3port69,
		comswitch3port70,
		comswitch3port71,
		comswitch3port72,
		comswitch3port73,
		comswitch3port74,
		comswitch3port75,
		comswitch3port76,
		comswitch3port77,
		comswitch3port78,
		comswitch3port79,
		comswitch3port80,
		comswitch3port81,
		comswitch3port82,
		comswitch3port83,
		comswitch3port84,
		comswitch3port85,
		comswitch3port86,
		comswitch3port87,
		comswitch3port88,
		comswitch3port89,
		comswitch3port90,
		comswitch3port91,
		comswitch3port92,
		comswitch3port93,
		comswitch3port94,
		comswitch3port95,
		comswitch3port96,
		comswitch3port97,
		comswitch3port98,
		comswitch3port99,
		comswitch3port100,
		comswitch3port101,
		comswitch3port102,
		comswitch3port103,
		comswitch3port104,
		comswitch3port105,
		comswitch3port106,
		comswitch3port107,
		comswitch3port108,
		comswitch3port109,
		comswitch3port110,
		comswitch3port111,
		comswitch3port112,
		comswitch3port113,
		comswitch3port114,
		comswitch3port115,
		comswitch3port116,
		comswitch3port117,
		comswitch3port118,
		comswitch3port119,
		comswitch3port120,
		comswitch3port121,
		comswitch3port122,
		comswitch3port123,
		comswitch3port124,
		comswitch3port125,
		comswitch3port126,
		comswitch3port127,
		comswitch3port128,
		comswitch3port129,
		comswitch3port130,
		comswitch3port131,
		comswitch3port132,
		comswitch3port133,
		comswitch3port134,
		comswitch3port135,
		comswitch3port136,
		comswitch3port137,
		comswitch3port138,
		comswitch3port139,
		comswitch3port140,
		comswitch3port141, comswitch3port142}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}

	comswitch4port1   = &ConcretePort{Desc: "switch4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port2   = &ConcretePort{Desc: "switch4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port3   = &ConcretePort{Desc: "switch4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port4   = &ConcretePort{Desc: "switch4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port5   = &ConcretePort{Desc: "switch4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port6   = &ConcretePort{Desc: "switch4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port7   = &ConcretePort{Desc: "switch4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port8   = &ConcretePort{Desc: "switch4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port9   = &ConcretePort{Desc: "switch4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port10  = &ConcretePort{Desc: "switch4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port11  = &ConcretePort{Desc: "switch4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port12  = &ConcretePort{Desc: "switch4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port13  = &ConcretePort{Desc: "switch4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port14  = &ConcretePort{Desc: "switch4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port15  = &ConcretePort{Desc: "switch4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port16  = &ConcretePort{Desc: "switch4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port17  = &ConcretePort{Desc: "switch4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port18  = &ConcretePort{Desc: "switch4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port19  = &ConcretePort{Desc: "switch4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port20  = &ConcretePort{Desc: "switch4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port21  = &ConcretePort{Desc: "switch4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port22  = &ConcretePort{Desc: "switch4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port23  = &ConcretePort{Desc: "switch4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port24  = &ConcretePort{Desc: "switch4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port25  = &ConcretePort{Desc: "switch4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port26  = &ConcretePort{Desc: "switch4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port27  = &ConcretePort{Desc: "switch4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port28  = &ConcretePort{Desc: "switch4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port29  = &ConcretePort{Desc: "switch4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port30  = &ConcretePort{Desc: "switch4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port31  = &ConcretePort{Desc: "switch4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port32  = &ConcretePort{Desc: "switch4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port33  = &ConcretePort{Desc: "switch4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port34  = &ConcretePort{Desc: "switch4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port35  = &ConcretePort{Desc: "switch4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port36  = &ConcretePort{Desc: "switch4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port37  = &ConcretePort{Desc: "switch4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port38  = &ConcretePort{Desc: "switch4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port39  = &ConcretePort{Desc: "switch4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port40  = &ConcretePort{Desc: "switch4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port41  = &ConcretePort{Desc: "switch4:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port42  = &ConcretePort{Desc: "switch4:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port43  = &ConcretePort{Desc: "switch4:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port44  = &ConcretePort{Desc: "switch4:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port45  = &ConcretePort{Desc: "switch4:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port46  = &ConcretePort{Desc: "switch4:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port47  = &ConcretePort{Desc: "switch4:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port48  = &ConcretePort{Desc: "switch4:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port49  = &ConcretePort{Desc: "switch4:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port50  = &ConcretePort{Desc: "switch4:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port51  = &ConcretePort{Desc: "switch4:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port52  = &ConcretePort{Desc: "switch4:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port53  = &ConcretePort{Desc: "switch4:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port54  = &ConcretePort{Desc: "switch4:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port55  = &ConcretePort{Desc: "switch4:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port56  = &ConcretePort{Desc: "switch4:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port57  = &ConcretePort{Desc: "switch4:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port58  = &ConcretePort{Desc: "switch4:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port59  = &ConcretePort{Desc: "switch4:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port60  = &ConcretePort{Desc: "switch4:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port61  = &ConcretePort{Desc: "switch4:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port62  = &ConcretePort{Desc: "switch4:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port63  = &ConcretePort{Desc: "switch4:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port64  = &ConcretePort{Desc: "switch4:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port65  = &ConcretePort{Desc: "switch4:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port66  = &ConcretePort{Desc: "switch4:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port67  = &ConcretePort{Desc: "switch4:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port68  = &ConcretePort{Desc: "switch4:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port69  = &ConcretePort{Desc: "switch4:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port70  = &ConcretePort{Desc: "switch4:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port71  = &ConcretePort{Desc: "switch4:port71", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port72  = &ConcretePort{Desc: "switch4:port72", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port73  = &ConcretePort{Desc: "switch4:port73", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port74  = &ConcretePort{Desc: "switch4:port74", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port75  = &ConcretePort{Desc: "switch4:port75", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port76  = &ConcretePort{Desc: "switch4:port76", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port77  = &ConcretePort{Desc: "switch4:port77", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port78  = &ConcretePort{Desc: "switch4:port78", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port79  = &ConcretePort{Desc: "switch4:port79", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port80  = &ConcretePort{Desc: "switch4:port80", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port81  = &ConcretePort{Desc: "switch4:port81", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port82  = &ConcretePort{Desc: "switch4:port82", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port83  = &ConcretePort{Desc: "switch4:port83", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port84  = &ConcretePort{Desc: "switch4:port84", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port85  = &ConcretePort{Desc: "switch4:port85", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port86  = &ConcretePort{Desc: "switch4:port86", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port87  = &ConcretePort{Desc: "switch4:port87", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port88  = &ConcretePort{Desc: "switch4:port88", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port89  = &ConcretePort{Desc: "switch4:port89", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port90  = &ConcretePort{Desc: "switch4:port90", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port91  = &ConcretePort{Desc: "switch4:port91", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port92  = &ConcretePort{Desc: "switch4:port92", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port93  = &ConcretePort{Desc: "switch4:port93", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port94  = &ConcretePort{Desc: "switch4:port94", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port95  = &ConcretePort{Desc: "switch4:port95", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port96  = &ConcretePort{Desc: "switch4:port96", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port97  = &ConcretePort{Desc: "switch4:port97", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port98  = &ConcretePort{Desc: "switch4:port98", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port99  = &ConcretePort{Desc: "switch4:port99", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port100 = &ConcretePort{Desc: "switch4:port100", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port101 = &ConcretePort{Desc: "switch4:port101", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port102 = &ConcretePort{Desc: "switch4:port102", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port103 = &ConcretePort{Desc: "switch4:port103", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port104 = &ConcretePort{Desc: "switch4:port104", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port105 = &ConcretePort{Desc: "switch4:port105", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port106 = &ConcretePort{Desc: "switch4:port106", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port107 = &ConcretePort{Desc: "switch4:port107", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port108 = &ConcretePort{Desc: "switch4:port108", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port109 = &ConcretePort{Desc: "switch4:port109", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port110 = &ConcretePort{Desc: "switch4:port110", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port111 = &ConcretePort{Desc: "switch4:port111", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port112 = &ConcretePort{Desc: "switch4:port112", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port113 = &ConcretePort{Desc: "switch4:port113", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port114 = &ConcretePort{Desc: "switch4:port114", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port115 = &ConcretePort{Desc: "switch4:port115", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port116 = &ConcretePort{Desc: "switch4:port116", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port117 = &ConcretePort{Desc: "switch4:port117", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port118 = &ConcretePort{Desc: "switch4:port118", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port119 = &ConcretePort{Desc: "switch4:port119", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port120 = &ConcretePort{Desc: "switch4:port120", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port121 = &ConcretePort{Desc: "switch4:port121", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port122 = &ConcretePort{Desc: "switch4:port122", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port123 = &ConcretePort{Desc: "switch4:port123", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port124 = &ConcretePort{Desc: "switch4:port124", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port125 = &ConcretePort{Desc: "switch4:port125", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port126 = &ConcretePort{Desc: "switch4:port126", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port127 = &ConcretePort{Desc: "switch4:port127", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port128 = &ConcretePort{Desc: "switch4:port128", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port129 = &ConcretePort{Desc: "switch4:port129", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port130 = &ConcretePort{Desc: "switch4:port130", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port131 = &ConcretePort{Desc: "switch4:port131", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port132 = &ConcretePort{Desc: "switch4:port132", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port133 = &ConcretePort{Desc: "switch4:port133", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port134 = &ConcretePort{Desc: "switch4:port134", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port135 = &ConcretePort{Desc: "switch4:port135", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port136 = &ConcretePort{Desc: "switch4:port136", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port137 = &ConcretePort{Desc: "switch4:port137", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port138 = &ConcretePort{Desc: "switch4:port138", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port139 = &ConcretePort{Desc: "switch4:port139", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port140 = &ConcretePort{Desc: "switch4:port140", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port141 = &ConcretePort{Desc: "switch4:port141", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch4port142 = &ConcretePort{Desc: "switch4:port142", Attrs: map[string]string{"speed": "S_400G"}}

	comswitch4 = &ConcreteNode{Desc: "switch4", Ports: []*ConcretePort{
		comswitch4port1,
		comswitch4port2,
		comswitch4port3,
		comswitch4port4,
		comswitch4port5,
		comswitch4port6,
		comswitch4port7,
		comswitch4port8,
		comswitch4port9,
		comswitch4port10,
		comswitch4port11,
		comswitch4port12,
		comswitch4port13,
		comswitch4port14,
		comswitch4port15,
		comswitch4port16,
		comswitch4port17,
		comswitch4port18,
		comswitch4port19,
		comswitch4port20,
		comswitch4port21,
		comswitch4port22,
		comswitch4port23,
		comswitch4port24,
		comswitch4port25,
		comswitch4port26,
		comswitch4port27,
		comswitch4port28,
		comswitch4port29,
		comswitch4port30,
		comswitch4port31,
		comswitch4port32,
		comswitch4port33,
		comswitch4port34,
		comswitch4port35,
		comswitch4port36,
		comswitch4port37,
		comswitch4port38,
		comswitch4port39,
		comswitch4port40,
		comswitch4port41,
		comswitch4port42,
		comswitch4port43,
		comswitch4port44,
		comswitch4port45,
		comswitch4port46,
		comswitch4port47,
		comswitch4port48,
		comswitch4port49,
		comswitch4port50,
		comswitch4port51,
		comswitch4port52,
		comswitch4port53,
		comswitch4port54,
		comswitch4port55,
		comswitch4port56,
		comswitch4port57,
		comswitch4port58,
		comswitch4port59,
		comswitch4port60,
		comswitch4port61,
		comswitch4port62,
		comswitch4port63,
		comswitch4port64,
		comswitch4port65,
		comswitch4port66,
		comswitch4port67,
		comswitch4port68,
		comswitch4port69,
		comswitch4port70,
		comswitch4port71,
		comswitch4port72,
		comswitch4port73,
		comswitch4port74,
		comswitch4port75,
		comswitch4port76,
		comswitch4port77,
		comswitch4port78,
		comswitch4port79,
		comswitch4port80,
		comswitch4port81,
		comswitch4port82,
		comswitch4port83,
		comswitch4port84,
		comswitch4port85,
		comswitch4port86,
		comswitch4port87,
		comswitch4port88,
		comswitch4port89,
		comswitch4port90,
		comswitch4port91,
		comswitch4port92,
		comswitch4port93,
		comswitch4port94,
		comswitch4port95,
		comswitch4port96,
		comswitch4port97,
		comswitch4port98,
		comswitch4port99,
		comswitch4port100,
		comswitch4port101,
		comswitch4port102,
		comswitch4port103,
		comswitch4port104,
		comswitch4port105,
		comswitch4port106,
		comswitch4port107,
		comswitch4port108,
		comswitch4port109,
		comswitch4port110,
		comswitch4port111,
		comswitch4port112,
		comswitch4port113,
		comswitch4port114,
		comswitch4port115,
		comswitch4port116,
		comswitch4port117,
		comswitch4port118,
		comswitch4port119,
		comswitch4port120,
		comswitch4port121,
		comswitch4port122,
		comswitch4port123,
		comswitch4port124,
		comswitch4port125,
		comswitch4port126,
		comswitch4port127,
		comswitch4port128,
		comswitch4port129,
		comswitch4port130,
		comswitch4port131,
		comswitch4port132,
		comswitch4port133,
		comswitch4port134,
		comswitch4port135,
		comswitch4port136,
		comswitch4port137,
		comswitch4port138,
		comswitch4port139,
		comswitch4port140,
		comswitch4port141,
		comswitch4port142}, Attrs: map[string]string{"role": "L1S", "name": "sw4"}}

	comswitch5port1   = &ConcretePort{Desc: "switch5:port1", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port2   = &ConcretePort{Desc: "switch5:port2", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port3   = &ConcretePort{Desc: "switch5:port3", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port4   = &ConcretePort{Desc: "switch5:port4", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port5   = &ConcretePort{Desc: "switch5:port5", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port6   = &ConcretePort{Desc: "switch5:port6", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port7   = &ConcretePort{Desc: "switch5:port7", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port8   = &ConcretePort{Desc: "switch5:port8", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port9   = &ConcretePort{Desc: "switch5:port9", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port10  = &ConcretePort{Desc: "switch5:port10", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port11  = &ConcretePort{Desc: "switch5:port11", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port12  = &ConcretePort{Desc: "switch5:port12", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port13  = &ConcretePort{Desc: "switch5:port13", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port14  = &ConcretePort{Desc: "switch5:port14", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port15  = &ConcretePort{Desc: "switch5:port15", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port16  = &ConcretePort{Desc: "switch5:port16", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port17  = &ConcretePort{Desc: "switch5:port17", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port18  = &ConcretePort{Desc: "switch5:port18", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port19  = &ConcretePort{Desc: "switch5:port19", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port20  = &ConcretePort{Desc: "switch5:port20", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port21  = &ConcretePort{Desc: "switch5:port21", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port22  = &ConcretePort{Desc: "switch5:port22", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port23  = &ConcretePort{Desc: "switch5:port23", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port24  = &ConcretePort{Desc: "switch5:port24", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port25  = &ConcretePort{Desc: "switch5:port25", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port26  = &ConcretePort{Desc: "switch5:port26", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port27  = &ConcretePort{Desc: "switch5:port27", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port28  = &ConcretePort{Desc: "switch5:port28", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port29  = &ConcretePort{Desc: "switch5:port29", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port30  = &ConcretePort{Desc: "switch5:port30", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port31  = &ConcretePort{Desc: "switch5:port31", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port32  = &ConcretePort{Desc: "switch5:port32", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port33  = &ConcretePort{Desc: "switch5:port33", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port34  = &ConcretePort{Desc: "switch5:port34", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port35  = &ConcretePort{Desc: "switch5:port35", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port36  = &ConcretePort{Desc: "switch5:port36", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port37  = &ConcretePort{Desc: "switch5:port37", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port38  = &ConcretePort{Desc: "switch5:port38", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port39  = &ConcretePort{Desc: "switch5:port39", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port40  = &ConcretePort{Desc: "switch5:port40", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port41  = &ConcretePort{Desc: "switch5:port41", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port42  = &ConcretePort{Desc: "switch5:port42", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port43  = &ConcretePort{Desc: "switch5:port43", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port44  = &ConcretePort{Desc: "switch5:port44", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port45  = &ConcretePort{Desc: "switch5:port45", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port46  = &ConcretePort{Desc: "switch5:port46", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port47  = &ConcretePort{Desc: "switch5:port47", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port48  = &ConcretePort{Desc: "switch5:port48", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port49  = &ConcretePort{Desc: "switch5:port49", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port50  = &ConcretePort{Desc: "switch5:port50", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port51  = &ConcretePort{Desc: "switch5:port51", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port52  = &ConcretePort{Desc: "switch5:port52", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port53  = &ConcretePort{Desc: "switch5:port53", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port54  = &ConcretePort{Desc: "switch5:port54", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port55  = &ConcretePort{Desc: "switch5:port55", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port56  = &ConcretePort{Desc: "switch5:port56", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port57  = &ConcretePort{Desc: "switch5:port57", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port58  = &ConcretePort{Desc: "switch5:port58", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port59  = &ConcretePort{Desc: "switch5:port59", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port60  = &ConcretePort{Desc: "switch5:port60", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port61  = &ConcretePort{Desc: "switch5:port61", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port62  = &ConcretePort{Desc: "switch5:port62", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port63  = &ConcretePort{Desc: "switch5:port63", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port64  = &ConcretePort{Desc: "switch5:port64", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port65  = &ConcretePort{Desc: "switch5:port65", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port66  = &ConcretePort{Desc: "switch5:port66", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port67  = &ConcretePort{Desc: "switch5:port67", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port68  = &ConcretePort{Desc: "switch5:port68", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port69  = &ConcretePort{Desc: "switch5:port69", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port70  = &ConcretePort{Desc: "switch5:port70", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port71  = &ConcretePort{Desc: "switch5:port71", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port72  = &ConcretePort{Desc: "switch5:port72", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port73  = &ConcretePort{Desc: "switch5:port73", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port74  = &ConcretePort{Desc: "switch5:port74", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port75  = &ConcretePort{Desc: "switch5:port75", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port76  = &ConcretePort{Desc: "switch5:port76", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port77  = &ConcretePort{Desc: "switch5:port77", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port78  = &ConcretePort{Desc: "switch5:port78", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port79  = &ConcretePort{Desc: "switch5:port79", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port80  = &ConcretePort{Desc: "switch5:port80", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port81  = &ConcretePort{Desc: "switch5:port81", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port82  = &ConcretePort{Desc: "switch5:port82", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port83  = &ConcretePort{Desc: "switch5:port83", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port84  = &ConcretePort{Desc: "switch5:port84", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port85  = &ConcretePort{Desc: "switch5:port85", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port86  = &ConcretePort{Desc: "switch5:port86", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port87  = &ConcretePort{Desc: "switch5:port87", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port88  = &ConcretePort{Desc: "switch5:port88", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port89  = &ConcretePort{Desc: "switch5:port89", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port90  = &ConcretePort{Desc: "switch5:port90", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port91  = &ConcretePort{Desc: "switch5:port91", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port92  = &ConcretePort{Desc: "switch5:port92", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port93  = &ConcretePort{Desc: "switch5:port93", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port94  = &ConcretePort{Desc: "switch5:port94", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port95  = &ConcretePort{Desc: "switch5:port95", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port96  = &ConcretePort{Desc: "switch5:port96", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port97  = &ConcretePort{Desc: "switch5:port97", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port98  = &ConcretePort{Desc: "switch5:port98", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port99  = &ConcretePort{Desc: "switch5:port99", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port100 = &ConcretePort{Desc: "switch5:port100", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port101 = &ConcretePort{Desc: "switch5:port101", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port102 = &ConcretePort{Desc: "switch5:port102", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port103 = &ConcretePort{Desc: "switch5:port103", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port104 = &ConcretePort{Desc: "switch5:port104", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port105 = &ConcretePort{Desc: "switch5:port105", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port106 = &ConcretePort{Desc: "switch5:port106", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port107 = &ConcretePort{Desc: "switch5:port107", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port108 = &ConcretePort{Desc: "switch5:port108", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port109 = &ConcretePort{Desc: "switch5:port109", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port110 = &ConcretePort{Desc: "switch5:port110", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port111 = &ConcretePort{Desc: "switch5:port111", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port112 = &ConcretePort{Desc: "switch5:port112", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port113 = &ConcretePort{Desc: "switch5:port113", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port114 = &ConcretePort{Desc: "switch5:port114", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port115 = &ConcretePort{Desc: "switch5:port115", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port116 = &ConcretePort{Desc: "switch5:port116", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port117 = &ConcretePort{Desc: "switch5:port117", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port118 = &ConcretePort{Desc: "switch5:port118", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port119 = &ConcretePort{Desc: "switch5:port119", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port120 = &ConcretePort{Desc: "switch5:port120", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port121 = &ConcretePort{Desc: "switch5:port121", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port122 = &ConcretePort{Desc: "switch5:port122", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port123 = &ConcretePort{Desc: "switch5:port123", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port124 = &ConcretePort{Desc: "switch5:port124", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port125 = &ConcretePort{Desc: "switch5:port125", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port126 = &ConcretePort{Desc: "switch5:port126", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port127 = &ConcretePort{Desc: "switch5:port127", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port128 = &ConcretePort{Desc: "switch5:port128", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port129 = &ConcretePort{Desc: "switch5:port129", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port130 = &ConcretePort{Desc: "switch5:port130", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port131 = &ConcretePort{Desc: "switch5:port131", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port132 = &ConcretePort{Desc: "switch5:port132", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port133 = &ConcretePort{Desc: "switch5:port133", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port134 = &ConcretePort{Desc: "switch5:port134", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port135 = &ConcretePort{Desc: "switch5:port135", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port136 = &ConcretePort{Desc: "switch5:port136", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port137 = &ConcretePort{Desc: "switch5:port137", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port138 = &ConcretePort{Desc: "switch5:port138", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port139 = &ConcretePort{Desc: "switch5:port139", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port140 = &ConcretePort{Desc: "switch5:port140", Attrs: map[string]string{"speed": "S_400G"}}
	comswitch5port141 = &ConcretePort{Desc: "switch5:port141", Attrs: map[string]string{"speed": "S_400G"}}

	comswitch5 = &ConcreteNode{Desc: "switch5", Ports: []*ConcretePort{
		comswitch5port1,
		comswitch5port2,
		comswitch5port3,
		comswitch5port4,
		comswitch5port5,
		comswitch5port6,
		comswitch5port7,
		comswitch5port8,
		comswitch5port9,
		comswitch5port10,
		comswitch5port11,
		comswitch5port12,
		comswitch5port13,
		comswitch5port14,
		comswitch5port15,
		comswitch5port16,
		comswitch5port17,
		comswitch5port18,
		comswitch5port19,
		comswitch5port20,
		comswitch5port21,
		comswitch5port22,
		comswitch5port23,
		comswitch5port24,
		comswitch5port25,
		comswitch5port26,
		comswitch5port27,
		comswitch5port28,
		comswitch5port29,
		comswitch5port30,
		comswitch5port31,
		comswitch5port32,
		comswitch5port33,
		comswitch5port34,
		comswitch5port35,
		comswitch5port36,
		comswitch5port37,
		comswitch5port38,
		comswitch5port39,
		comswitch5port40,
		comswitch5port41,
		comswitch5port42,
		comswitch5port43,
		comswitch5port44,
		comswitch5port45,
		comswitch5port46,
		comswitch5port47,
		comswitch5port48,
		comswitch5port49,
		comswitch5port50,
		comswitch5port51,
		comswitch5port52,
		comswitch5port53,
		comswitch5port54,
		comswitch5port55,
		comswitch5port56,
		comswitch5port57,
		comswitch5port58,
		comswitch5port59,
		comswitch5port60,
		comswitch5port61,
		comswitch5port62,
		comswitch5port63,
		comswitch5port64,
		comswitch5port65,
		comswitch5port66,
		comswitch5port67,
		comswitch5port68,
		comswitch5port69,
		comswitch5port70,
		comswitch5port71,
		comswitch5port72,
		comswitch5port73,
		comswitch5port74,
		comswitch5port75,
		comswitch5port76,
		comswitch5port77,
		comswitch5port78,
		comswitch5port79,
		comswitch5port80,
		comswitch5port81,
		comswitch5port82,
		comswitch5port83,
		comswitch5port84,
		comswitch5port85,
		comswitch5port86,
		comswitch5port87,
		comswitch5port88,
		comswitch5port89,
		comswitch5port90,
		comswitch5port91,
		comswitch5port92,
		comswitch5port93,
		comswitch5port94,
		comswitch5port95,
		comswitch5port96,
		comswitch5port97,
		comswitch5port98,
		comswitch5port99,
		comswitch5port100,
		comswitch5port101,
		comswitch5port102,
		comswitch5port103,
		comswitch5port104,
		comswitch5port105,
		comswitch5port106,
		comswitch5port107,
		comswitch5port108,
		comswitch5port109,
		comswitch5port110,
		comswitch5port111,
		comswitch5port112,
		comswitch5port113,
		comswitch5port114,
		comswitch5port115,
		comswitch5port116,
		comswitch5port117,
		comswitch5port118,
		comswitch5port119,
		comswitch5port120,
		comswitch5port121,
		comswitch5port122,
		comswitch5port123,
		comswitch5port124,
		comswitch5port125,
		comswitch5port126,
		comswitch5port127,
		comswitch5port128,
		comswitch5port129,
		comswitch5port130,
		comswitch5port131,
		comswitch5port132,
		comswitch5port133,
		comswitch5port134,
		comswitch5port135,
		comswitch5port136,
		comswitch5port137,
		comswitch5port138,
		comswitch5port139,
		comswitch5port140,
		comswitch5port141}, Attrs: map[string]string{"role": "L1S", "name": "sw5"}}
)

var superGraphTestcom = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		comdut1,
		comate1,
		comdut2,
		comate2,
		comdut3,
		comate3,
		comdut4,
		comate4,
		comdut5,
		comate5,
		comdut6,
		comate6,
		comswitch1, comswitch2, comswitch3, comswitch4, comswitch5},
	Edges: []*ConcreteEdge{
		{Src: comdut1port1, Dst: comate1port1},
		{Src: comdut1port2, Dst: comate1port2},
		{Src: comdut1port3, Dst: comate1port3},
		{Src: comdut1port4, Dst: comate1port4},
		{Src: comdut1port5, Dst: comate1port5},
		{Src: comdut1port6, Dst: comate1port6},
		{Src: comdut1port7, Dst: comate1port7},
		{Src: comdut1port8, Dst: comate1port8},
		{Src: comdut1port9, Dst: comate1port9},
		{Src: comdut1port10, Dst: comate1port10},
		{Src: comdut1port11, Dst: comate1port11},
		{Src: comdut1port12, Dst: comate1port12},
		{Src: comdut1port13, Dst: comate1port13},
		{Src: comdut1port14, Dst: comate1port14},
		{Src: comdut1port15, Dst: comate1port15},
		{Src: comdut1port16, Dst: comate1port16},
		{Src: comdut1port17, Dst: comate1port17},
		{Src: comdut1port18, Dst: comate1port18},
		{Src: comdut1port19, Dst: comate1port19},
		{Src: comdut1port20, Dst: comate1port20},
		{Src: comdut1port21, Dst: comate1port21},
		{Src: comdut1port22, Dst: comate1port22},
		{Src: comdut1port23, Dst: comate1port23},
		{Src: comdut1port24, Dst: comate1port24},
		{Src: comdut1port25, Dst: comate1port25},
		{Src: comdut1port26, Dst: comate1port26},
		{Src: comdut1port27, Dst: comate1port27},
		{Src: comdut1port28, Dst: comate1port28},
		{Src: comdut1port29, Dst: comate1port29},
		{Src: comdut1port30, Dst: comate1port30},
		{Src: comdut1port31, Dst: comate1port31},
		{Src: comdut1port32, Dst: comate1port32},
		{Src: comdut1port33, Dst: comate1port33},
		{Src: comdut1port34, Dst: comate1port34},
		{Src: comdut1port35, Dst: comate1port35},

		{Src: comdut1port36, Dst: comswitch1port1},
		{Src: comdut1port37, Dst: comswitch1port2},
		{Src: comdut1port38, Dst: comswitch1port3},
		{Src: comdut1port39, Dst: comswitch1port4},
		{Src: comdut1port40, Dst: comswitch1port5},
		{Src: comdut1port41, Dst: comswitch1port6},
		{Src: comdut1port42, Dst: comswitch1port7},
		{Src: comdut1port43, Dst: comswitch1port8},
		{Src: comdut1port44, Dst: comswitch1port9},
		{Src: comdut1port45, Dst: comswitch1port10},
		{Src: comdut1port46, Dst: comswitch1port11},
		{Src: comdut1port47, Dst: comswitch1port12},
		{Src: comdut1port48, Dst: comswitch1port13},
		{Src: comdut1port49, Dst: comswitch1port14},
		{Src: comdut1port50, Dst: comswitch1port15},
		{Src: comdut1port51, Dst: comswitch1port16},
		{Src: comdut1port52, Dst: comswitch1port17},
		{Src: comdut1port53, Dst: comswitch1port18},
		{Src: comdut1port54, Dst: comswitch1port19},
		{Src: comdut1port55, Dst: comswitch1port20},
		{Src: comdut1port56, Dst: comswitch1port21},
		{Src: comdut1port57, Dst: comswitch1port22},
		{Src: comdut1port58, Dst: comswitch1port23},
		{Src: comdut1port59, Dst: comswitch1port24},
		{Src: comdut1port60, Dst: comswitch1port25},
		{Src: comdut1port61, Dst: comswitch1port26},
		{Src: comdut1port62, Dst: comswitch1port27},
		{Src: comdut1port63, Dst: comswitch1port28},
		{Src: comdut1port64, Dst: comswitch1port29},
		{Src: comdut1port65, Dst: comswitch1port30},
		{Src: comdut1port66, Dst: comswitch1port31},
		{Src: comdut1port67, Dst: comswitch1port32},
		{Src: comdut1port68, Dst: comswitch1port33},
		{Src: comdut1port69, Dst: comswitch1port34},
		{Src: comdut1port70, Dst: comswitch1port35},

		{Src: comswitch1port36, Dst: comate1port36},
		{Src: comswitch1port37, Dst: comate1port37},
		{Src: comswitch1port38, Dst: comate1port38},
		{Src: comswitch1port39, Dst: comate1port39},
		{Src: comswitch1port40, Dst: comate1port40},
		{Src: comswitch1port41, Dst: comate1port41},
		{Src: comswitch1port42, Dst: comate1port42},
		{Src: comswitch1port43, Dst: comate1port43},
		{Src: comswitch1port44, Dst: comate1port44},
		{Src: comswitch1port45, Dst: comate1port45},
		{Src: comswitch1port46, Dst: comate1port46},
		{Src: comswitch1port47, Dst: comate1port47},
		{Src: comswitch1port48, Dst: comate1port48},
		{Src: comswitch1port49, Dst: comate1port49},
		{Src: comswitch1port50, Dst: comate1port50},
		{Src: comswitch1port51, Dst: comate1port51},
		{Src: comswitch1port52, Dst: comate1port52},
		{Src: comswitch1port53, Dst: comate1port53},
		{Src: comswitch1port54, Dst: comate1port54},
		{Src: comswitch1port55, Dst: comate1port55},
		{Src: comswitch1port56, Dst: comate1port56},
		{Src: comswitch1port57, Dst: comate1port57},
		{Src: comswitch1port58, Dst: comate1port58},
		{Src: comswitch1port59, Dst: comate1port59},
		{Src: comswitch1port60, Dst: comate1port60},
		{Src: comswitch1port61, Dst: comate1port61},
		{Src: comswitch1port62, Dst: comate1port62},
		{Src: comswitch1port63, Dst: comate1port63},
		{Src: comswitch1port64, Dst: comate1port64},
		{Src: comswitch1port65, Dst: comate1port65},
		{Src: comswitch1port66, Dst: comate1port66},
		{Src: comswitch1port67, Dst: comate1port67},
		{Src: comswitch1port68, Dst: comate1port68},
		{Src: comswitch1port69, Dst: comate1port69},
		{Src: comswitch1port70, Dst: comate1port70},
		{Src: comswitch1port141, Dst: comswitch2port141},

		{Src: comdut2port1, Dst: comate2port1},
		{Src: comdut2port2, Dst: comate2port2},
		{Src: comdut2port3, Dst: comate2port3},
		{Src: comdut2port4, Dst: comate2port4},
		{Src: comdut2port5, Dst: comate2port5},
		{Src: comdut2port6, Dst: comate2port6},
		{Src: comdut2port7, Dst: comate2port7},
		{Src: comdut2port8, Dst: comate2port8},
		{Src: comdut2port9, Dst: comate2port9},
		{Src: comdut2port10, Dst: comate2port10},
		{Src: comdut2port11, Dst: comate2port11},
		{Src: comdut2port12, Dst: comate2port12},
		{Src: comdut2port13, Dst: comate2port13},
		{Src: comdut2port14, Dst: comate2port14},
		{Src: comdut2port15, Dst: comate2port15},
		{Src: comdut2port16, Dst: comate2port16},
		{Src: comdut2port17, Dst: comate2port17},
		{Src: comdut2port18, Dst: comate2port18},
		{Src: comdut2port19, Dst: comate2port19},
		{Src: comdut2port20, Dst: comate2port20},
		{Src: comdut2port21, Dst: comate2port21},
		{Src: comdut2port22, Dst: comate2port22},
		{Src: comdut2port23, Dst: comate2port23},
		{Src: comdut2port24, Dst: comate2port24},
		{Src: comdut2port25, Dst: comate2port25},
		{Src: comdut2port26, Dst: comate2port26},
		{Src: comdut2port27, Dst: comate2port27},
		{Src: comdut2port28, Dst: comate2port28},
		{Src: comdut2port29, Dst: comate2port29},
		{Src: comdut2port30, Dst: comate2port30},
		{Src: comdut2port31, Dst: comate2port31},
		{Src: comdut2port32, Dst: comate2port32},
		{Src: comdut2port33, Dst: comate2port33},
		{Src: comdut2port34, Dst: comate2port34},
		{Src: comdut2port35, Dst: comate2port35},

		{Src: comdut2port36, Dst: comswitch1port71},
		{Src: comdut2port37, Dst: comswitch1port72},
		{Src: comdut2port38, Dst: comswitch1port73},
		{Src: comdut2port39, Dst: comswitch1port74},
		{Src: comdut2port40, Dst: comswitch1port75},
		{Src: comdut2port41, Dst: comswitch1port76},
		{Src: comdut2port42, Dst: comswitch1port77},
		{Src: comdut2port43, Dst: comswitch1port78},
		{Src: comdut2port44, Dst: comswitch1port79},
		{Src: comdut2port45, Dst: comswitch1port80},
		{Src: comdut2port46, Dst: comswitch1port81},
		{Src: comdut2port47, Dst: comswitch1port82},
		{Src: comdut2port48, Dst: comswitch1port83},
		{Src: comdut2port49, Dst: comswitch1port84},
		{Src: comdut2port50, Dst: comswitch1port85},
		{Src: comdut2port51, Dst: comswitch1port86},
		{Src: comdut2port52, Dst: comswitch1port87},
		{Src: comdut2port53, Dst: comswitch1port88},
		{Src: comdut2port54, Dst: comswitch1port89},
		{Src: comdut2port55, Dst: comswitch1port90},
		{Src: comdut2port56, Dst: comswitch1port91},
		{Src: comdut2port57, Dst: comswitch1port92},
		{Src: comdut2port58, Dst: comswitch1port93},
		{Src: comdut2port59, Dst: comswitch1port94},
		{Src: comdut2port60, Dst: comswitch1port95},
		{Src: comdut2port61, Dst: comswitch1port96},
		{Src: comdut2port62, Dst: comswitch1port97},
		{Src: comdut2port63, Dst: comswitch1port98},
		{Src: comdut2port64, Dst: comswitch1port99},
		{Src: comdut2port65, Dst: comswitch1port100},
		{Src: comdut2port66, Dst: comswitch1port101},
		{Src: comdut2port67, Dst: comswitch1port102},
		{Src: comdut2port68, Dst: comswitch1port103},
		{Src: comdut2port69, Dst: comswitch1port104},
		{Src: comdut2port70, Dst: comswitch1port105},

		{Src: comswitch1port106, Dst: comate2port36},
		{Src: comswitch1port107, Dst: comate2port37},
		{Src: comswitch1port108, Dst: comate2port38},
		{Src: comswitch1port109, Dst: comate2port39},
		{Src: comswitch1port110, Dst: comate2port40},
		{Src: comswitch1port111, Dst: comate2port41},
		{Src: comswitch1port112, Dst: comate2port42},
		{Src: comswitch1port113, Dst: comate2port43},
		{Src: comswitch1port114, Dst: comate2port44},
		{Src: comswitch1port115, Dst: comate2port45},
		{Src: comswitch1port116, Dst: comate2port46},
		{Src: comswitch1port117, Dst: comate2port47},
		{Src: comswitch1port118, Dst: comate2port48},
		{Src: comswitch1port119, Dst: comate2port49},
		{Src: comswitch1port120, Dst: comate2port50},
		{Src: comswitch1port121, Dst: comate2port51},
		{Src: comswitch1port122, Dst: comate2port52},
		{Src: comswitch1port123, Dst: comate2port53},
		{Src: comswitch1port124, Dst: comate2port54},
		{Src: comswitch1port125, Dst: comate2port55},
		{Src: comswitch1port126, Dst: comate2port56},
		{Src: comswitch1port127, Dst: comate2port57},
		{Src: comswitch1port128, Dst: comate2port58},
		{Src: comswitch1port129, Dst: comate2port59},
		{Src: comswitch1port130, Dst: comate2port60},
		{Src: comswitch1port131, Dst: comate2port61},
		{Src: comswitch1port132, Dst: comate2port62},
		{Src: comswitch1port133, Dst: comate2port63},
		{Src: comswitch1port134, Dst: comate2port64},
		{Src: comswitch1port135, Dst: comate2port65},
		{Src: comswitch1port136, Dst: comate2port66},
		{Src: comswitch1port137, Dst: comate2port67},
		{Src: comswitch1port138, Dst: comate2port68},
		{Src: comswitch1port139, Dst: comate2port69},
		{Src: comswitch1port140, Dst: comate2port70},

		{Src: comdut3port1, Dst: comswitch2port1},
		{Src: comdut3port2, Dst: comswitch2port2},
		{Src: comdut3port3, Dst: comswitch2port3},
		{Src: comdut3port4, Dst: comswitch2port4},
		{Src: comdut3port5, Dst: comswitch2port5},
		{Src: comdut3port6, Dst: comswitch2port6},
		{Src: comdut3port7, Dst: comswitch2port7},
		{Src: comdut3port8, Dst: comswitch2port8},
		{Src: comdut3port9, Dst: comswitch2port9},
		{Src: comdut3port10, Dst: comswitch2port10},
		{Src: comdut3port11, Dst: comswitch2port11},
		{Src: comdut3port12, Dst: comswitch2port12},
		{Src: comdut3port13, Dst: comswitch2port13},
		{Src: comdut3port14, Dst: comswitch2port14},
		{Src: comdut3port15, Dst: comswitch2port15},
		{Src: comdut3port16, Dst: comswitch2port16},
		{Src: comdut3port17, Dst: comswitch2port17},
		{Src: comdut3port18, Dst: comswitch2port18},
		{Src: comdut3port19, Dst: comswitch2port19},
		{Src: comdut3port20, Dst: comswitch2port20},
		{Src: comdut3port21, Dst: comswitch2port21},
		{Src: comdut3port22, Dst: comswitch2port22},
		{Src: comdut3port23, Dst: comswitch2port23},
		{Src: comdut3port24, Dst: comswitch2port24},
		{Src: comdut3port25, Dst: comswitch2port25},
		{Src: comdut3port26, Dst: comswitch2port26},
		{Src: comdut3port27, Dst: comswitch2port27},
		{Src: comdut3port28, Dst: comswitch2port28},
		{Src: comdut3port29, Dst: comswitch2port29},
		{Src: comdut3port30, Dst: comswitch2port30},
		{Src: comdut3port31, Dst: comswitch2port31},
		{Src: comdut3port32, Dst: comswitch2port32},
		{Src: comdut3port33, Dst: comswitch2port33},
		{Src: comdut3port34, Dst: comswitch2port34},
		{Src: comdut3port35, Dst: comswitch2port35},

		{Src: comswitch2port41, Dst: comate3port1},
		{Src: comswitch2port42, Dst: comate3port2},
		{Src: comswitch2port43, Dst: comate3port3},
		{Src: comswitch2port44, Dst: comate3port4},
		{Src: comswitch2port45, Dst: comate3port5},
		{Src: comswitch2port46, Dst: comate3port6},
		{Src: comswitch2port47, Dst: comate3port7},
		{Src: comswitch2port48, Dst: comate3port8},
		{Src: comswitch2port49, Dst: comate3port9},
		{Src: comswitch2port50, Dst: comate3port10},
		{Src: comswitch2port51, Dst: comate3port11},
		{Src: comswitch2port52, Dst: comate3port12},
		{Src: comswitch2port53, Dst: comate3port13},
		{Src: comswitch2port54, Dst: comate3port14},
		{Src: comswitch2port55, Dst: comate3port15},
		{Src: comswitch2port56, Dst: comate3port16},
		{Src: comswitch2port57, Dst: comate3port17},
		{Src: comswitch2port58, Dst: comate3port18},
		{Src: comswitch2port59, Dst: comate3port19},
		{Src: comswitch2port60, Dst: comate3port20},
		{Src: comswitch2port61, Dst: comate3port21},
		{Src: comswitch2port62, Dst: comate3port22},
		{Src: comswitch2port63, Dst: comate3port23},
		{Src: comswitch2port64, Dst: comate3port24},
		{Src: comswitch2port65, Dst: comate3port25},
		{Src: comswitch2port66, Dst: comate3port26},
		{Src: comswitch2port67, Dst: comate3port27},
		{Src: comswitch2port68, Dst: comate3port28},
		{Src: comswitch2port69, Dst: comate3port29},
		{Src: comswitch2port70, Dst: comate3port30},
		{Src: comswitch2port71, Dst: comate3port31},
		{Src: comswitch2port72, Dst: comate3port32},
		{Src: comswitch2port73, Dst: comate3port33},
		{Src: comswitch2port74, Dst: comate3port34},
		{Src: comswitch2port75, Dst: comate3port35},

		{Src: comdut4port1, Dst: comswitch2port76},
		{Src: comdut4port2, Dst: comswitch2port77},
		{Src: comdut4port3, Dst: comswitch2port78},
		{Src: comdut4port4, Dst: comswitch2port79},
		{Src: comdut4port5, Dst: comswitch2port80},
		{Src: comdut4port6, Dst: comswitch2port81},
		{Src: comdut4port7, Dst: comswitch2port82},
		{Src: comdut4port8, Dst: comswitch2port83},
		{Src: comdut4port9, Dst: comswitch2port84},
		{Src: comdut4port10, Dst: comswitch2port85},
		{Src: comdut4port11, Dst: comswitch2port86},
		{Src: comdut4port12, Dst: comswitch2port87},
		{Src: comdut4port13, Dst: comswitch2port88},
		{Src: comdut4port14, Dst: comswitch2port89},
		{Src: comdut4port15, Dst: comswitch2port90},
		{Src: comdut4port16, Dst: comswitch2port91},
		{Src: comdut4port17, Dst: comswitch2port92},
		{Src: comdut4port18, Dst: comswitch2port93},
		{Src: comdut4port19, Dst: comswitch2port94},
		{Src: comdut4port20, Dst: comswitch2port95},
		{Src: comdut4port21, Dst: comswitch2port96},
		{Src: comdut4port22, Dst: comswitch2port97},
		{Src: comdut4port23, Dst: comswitch2port98},
		{Src: comdut4port24, Dst: comswitch2port99},
		{Src: comdut4port25, Dst: comswitch2port100},
		{Src: comdut4port26, Dst: comswitch2port101},
		{Src: comdut4port27, Dst: comswitch2port102},
		{Src: comdut4port28, Dst: comswitch2port103},
		{Src: comdut4port29, Dst: comswitch2port104},
		{Src: comdut4port30, Dst: comswitch2port105},
		{Src: comdut4port31, Dst: comswitch2port106},
		{Src: comdut4port32, Dst: comswitch2port107},
		{Src: comdut4port33, Dst: comswitch2port108},
		{Src: comdut4port34, Dst: comswitch2port109},
		{Src: comdut4port35, Dst: comswitch2port110},

		{Src: comswitch2port111, Dst: comate4port1},
		{Src: comswitch2port112, Dst: comate4port2},
		{Src: comswitch2port113, Dst: comate4port3},
		{Src: comswitch2port114, Dst: comate4port4},
		{Src: comswitch2port115, Dst: comate4port5},
		{Src: comswitch2port116, Dst: comate4port6},
		{Src: comswitch2port117, Dst: comate4port7},
		{Src: comswitch2port118, Dst: comate4port8},
		{Src: comswitch2port119, Dst: comate4port9},
		{Src: comswitch2port120, Dst: comate4port10},
		{Src: comswitch2port121, Dst: comate4port11},
		{Src: comswitch2port122, Dst: comate4port12},
		{Src: comswitch2port123, Dst: comate4port13},
		{Src: comswitch2port124, Dst: comate4port14},
		{Src: comswitch2port125, Dst: comate4port15},
		{Src: comswitch2port126, Dst: comate4port16},
		{Src: comswitch2port127, Dst: comate4port17},
		{Src: comswitch2port128, Dst: comate4port18},
		{Src: comswitch2port129, Dst: comate4port19},
		{Src: comswitch2port130, Dst: comate4port20},
		{Src: comswitch2port131, Dst: comate4port21},
		{Src: comswitch2port132, Dst: comate4port22},
		{Src: comswitch2port133, Dst: comate4port23},
		{Src: comswitch2port134, Dst: comate4port24},
		{Src: comswitch2port135, Dst: comate4port25},
		{Src: comswitch2port136, Dst: comate4port26},
		{Src: comswitch2port137, Dst: comate4port27},
		{Src: comswitch2port138, Dst: comate4port28},
		{Src: comswitch2port139, Dst: comate4port29},
		{Src: comswitch2port140, Dst: comate4port30},
		{Src: comswitch2port36, Dst: comate4port31},
		{Src: comswitch2port37, Dst: comate4port32},
		{Src: comswitch2port38, Dst: comate4port33},
		{Src: comswitch2port39, Dst: comate4port34},
		{Src: comswitch2port40, Dst: comate4port35},
		{Src: comswitch2port142, Dst: comswitch3port141},

		{Src: comdut3port36, Dst: comswitch3port1},
		{Src: comdut3port37, Dst: comswitch3port2},
		{Src: comdut3port38, Dst: comswitch3port3},
		{Src: comdut3port39, Dst: comswitch3port4},
		{Src: comdut3port40, Dst: comswitch3port5},
		{Src: comdut3port41, Dst: comswitch3port6},
		{Src: comdut3port42, Dst: comswitch3port7},
		{Src: comdut3port43, Dst: comswitch3port8},
		{Src: comdut3port44, Dst: comswitch3port9},
		{Src: comdut3port45, Dst: comswitch3port10},
		{Src: comdut3port46, Dst: comswitch3port11},
		{Src: comdut3port47, Dst: comswitch3port12},
		{Src: comdut3port48, Dst: comswitch3port13},
		{Src: comdut3port49, Dst: comswitch3port14},
		{Src: comdut3port50, Dst: comswitch3port15},
		{Src: comdut3port51, Dst: comswitch3port16},
		{Src: comdut3port52, Dst: comswitch3port17},
		{Src: comdut3port53, Dst: comswitch3port18},
		{Src: comdut3port54, Dst: comswitch3port19},
		{Src: comdut3port55, Dst: comswitch3port20},
		{Src: comdut3port56, Dst: comswitch3port21},
		{Src: comdut3port57, Dst: comswitch3port22},
		{Src: comdut3port58, Dst: comswitch3port23},
		{Src: comdut3port59, Dst: comswitch3port24},
		{Src: comdut3port60, Dst: comswitch3port25},
		{Src: comdut3port61, Dst: comswitch3port26},
		{Src: comdut3port62, Dst: comswitch3port27},
		{Src: comdut3port63, Dst: comswitch3port28},
		{Src: comdut3port64, Dst: comswitch3port29},
		{Src: comdut3port65, Dst: comswitch3port30},
		{Src: comdut3port66, Dst: comswitch3port31},
		{Src: comdut3port67, Dst: comswitch3port32},
		{Src: comdut3port68, Dst: comswitch3port33},
		{Src: comdut3port69, Dst: comswitch3port34},
		{Src: comdut3port70, Dst: comswitch3port35},

		{Src: comswitch3port36, Dst: comate3port36},
		{Src: comswitch3port37, Dst: comate3port37},
		{Src: comswitch3port38, Dst: comate3port38},
		{Src: comswitch3port39, Dst: comate3port39},
		{Src: comswitch3port40, Dst: comate3port40},
		{Src: comswitch3port41, Dst: comate3port41},
		{Src: comswitch3port42, Dst: comate3port42},
		{Src: comswitch3port43, Dst: comate3port43},
		{Src: comswitch3port44, Dst: comate3port44},
		{Src: comswitch3port45, Dst: comate3port45},
		{Src: comswitch3port46, Dst: comate3port46},
		{Src: comswitch3port47, Dst: comate3port47},
		{Src: comswitch3port48, Dst: comate3port48},
		{Src: comswitch3port49, Dst: comate3port49},
		{Src: comswitch3port50, Dst: comate3port50},
		{Src: comswitch3port51, Dst: comate3port51},
		{Src: comswitch3port52, Dst: comate3port52},
		{Src: comswitch3port53, Dst: comate3port53},
		{Src: comswitch3port54, Dst: comate3port54},
		{Src: comswitch3port55, Dst: comate3port55},
		{Src: comswitch3port56, Dst: comate3port56},
		{Src: comswitch3port57, Dst: comate3port57},
		{Src: comswitch3port58, Dst: comate3port58},
		{Src: comswitch3port59, Dst: comate3port59},
		{Src: comswitch3port60, Dst: comate3port60},
		{Src: comswitch3port61, Dst: comate3port61},
		{Src: comswitch3port62, Dst: comate3port62},
		{Src: comswitch3port63, Dst: comate3port63},
		{Src: comswitch3port64, Dst: comate3port64},
		{Src: comswitch3port65, Dst: comate3port65},
		{Src: comswitch3port66, Dst: comate3port66},
		{Src: comswitch3port67, Dst: comate3port67},
		{Src: comswitch3port68, Dst: comate3port68},
		{Src: comswitch3port69, Dst: comate3port69},
		{Src: comswitch3port70, Dst: comate3port70},

		{Src: comdut4port36, Dst: comswitch3port71},
		{Src: comdut4port37, Dst: comswitch3port72},
		{Src: comdut4port38, Dst: comswitch3port73},
		{Src: comdut4port39, Dst: comswitch3port74},
		{Src: comdut4port40, Dst: comswitch3port75},
		{Src: comdut4port41, Dst: comswitch3port76},
		{Src: comdut4port42, Dst: comswitch3port77},
		{Src: comdut4port43, Dst: comswitch3port78},
		{Src: comdut4port44, Dst: comswitch3port79},
		{Src: comdut4port45, Dst: comswitch3port80},
		{Src: comdut4port46, Dst: comswitch3port81},
		{Src: comdut4port47, Dst: comswitch3port82},
		{Src: comdut4port48, Dst: comswitch3port83},
		{Src: comdut4port49, Dst: comswitch3port84},
		{Src: comdut4port50, Dst: comswitch3port85},
		{Src: comdut4port51, Dst: comswitch3port86},
		{Src: comdut4port52, Dst: comswitch3port87},
		{Src: comdut4port53, Dst: comswitch3port88},
		{Src: comdut4port54, Dst: comswitch3port89},
		{Src: comdut4port55, Dst: comswitch3port90},
		{Src: comdut4port56, Dst: comswitch3port91},
		{Src: comdut4port57, Dst: comswitch3port92},
		{Src: comdut4port58, Dst: comswitch3port93},
		{Src: comdut4port59, Dst: comswitch3port94},
		{Src: comdut4port60, Dst: comswitch3port95},
		{Src: comdut4port61, Dst: comswitch3port96},
		{Src: comdut4port62, Dst: comswitch3port97},
		{Src: comdut4port63, Dst: comswitch3port98},
		{Src: comdut4port64, Dst: comswitch3port99},
		{Src: comdut4port65, Dst: comswitch3port100},
		{Src: comdut4port66, Dst: comswitch3port101},
		{Src: comdut4port67, Dst: comswitch3port102},
		{Src: comdut4port68, Dst: comswitch3port103},
		{Src: comdut4port69, Dst: comswitch3port104},
		{Src: comdut4port70, Dst: comswitch3port105},

		{Src: comswitch3port106, Dst: comate4port36},
		{Src: comswitch3port107, Dst: comate4port37},
		{Src: comswitch3port108, Dst: comate4port38},
		{Src: comswitch3port109, Dst: comate4port39},
		{Src: comswitch3port110, Dst: comate4port40},
		{Src: comswitch3port111, Dst: comate4port41},
		{Src: comswitch3port112, Dst: comate4port42},
		{Src: comswitch3port113, Dst: comate4port43},
		{Src: comswitch3port114, Dst: comate4port44},
		{Src: comswitch3port115, Dst: comate4port45},
		{Src: comswitch3port116, Dst: comate4port46},
		{Src: comswitch3port117, Dst: comate4port47},
		{Src: comswitch3port118, Dst: comate4port48},
		{Src: comswitch3port119, Dst: comate4port49},
		{Src: comswitch3port120, Dst: comate4port50},
		{Src: comswitch3port121, Dst: comate4port51},
		{Src: comswitch3port122, Dst: comate4port52},
		{Src: comswitch3port123, Dst: comate4port53},
		{Src: comswitch3port124, Dst: comate4port54},
		{Src: comswitch3port125, Dst: comate4port55},
		{Src: comswitch3port126, Dst: comate4port56},
		{Src: comswitch3port127, Dst: comate4port57},
		{Src: comswitch3port128, Dst: comate4port58},
		{Src: comswitch3port129, Dst: comate4port59},
		{Src: comswitch3port130, Dst: comate4port60},
		{Src: comswitch3port131, Dst: comate4port61},
		{Src: comswitch3port132, Dst: comate4port62},
		{Src: comswitch3port133, Dst: comate4port63},
		{Src: comswitch3port134, Dst: comate4port64},
		{Src: comswitch3port135, Dst: comate4port65},
		{Src: comswitch3port136, Dst: comate4port66},
		{Src: comswitch3port137, Dst: comate4port67},
		{Src: comswitch3port138, Dst: comate4port68},
		{Src: comswitch3port139, Dst: comate4port69},
		{Src: comswitch3port140, Dst: comate4port70},
		{Src: comswitch3port142, Dst: comswitch4port141},

		{Src: comdut5port1, Dst: comswitch4port1},
		{Src: comdut5port2, Dst: comswitch4port2},
		{Src: comdut5port3, Dst: comswitch4port3},
		{Src: comdut5port4, Dst: comswitch4port4},
		{Src: comdut5port5, Dst: comswitch4port5},
		{Src: comdut5port6, Dst: comswitch4port6},
		{Src: comdut5port7, Dst: comswitch4port7},
		{Src: comdut5port8, Dst: comswitch4port8},
		{Src: comdut5port9, Dst: comswitch4port9},
		{Src: comdut5port10, Dst: comswitch4port10},
		{Src: comdut5port11, Dst: comswitch4port11},
		{Src: comdut5port12, Dst: comswitch4port12},
		{Src: comdut5port13, Dst: comswitch4port13},
		{Src: comdut5port14, Dst: comswitch4port14},
		{Src: comdut5port15, Dst: comswitch4port15},
		{Src: comdut5port16, Dst: comswitch4port16},
		{Src: comdut5port17, Dst: comswitch4port17},
		{Src: comdut5port18, Dst: comswitch4port18},
		{Src: comdut5port19, Dst: comswitch4port19},
		{Src: comdut5port20, Dst: comswitch4port20},
		{Src: comdut5port21, Dst: comswitch4port21},
		{Src: comdut5port22, Dst: comswitch4port22},
		{Src: comdut5port23, Dst: comswitch4port23},
		{Src: comdut5port24, Dst: comswitch4port24},
		{Src: comdut5port25, Dst: comswitch4port25},
		{Src: comdut5port26, Dst: comswitch4port26},
		{Src: comdut5port27, Dst: comswitch4port27},
		{Src: comdut5port28, Dst: comswitch4port28},
		{Src: comdut5port29, Dst: comswitch4port29},
		{Src: comdut5port30, Dst: comswitch4port30},
		{Src: comdut5port31, Dst: comswitch4port31},
		{Src: comdut5port32, Dst: comswitch4port32},
		{Src: comdut5port33, Dst: comswitch4port33},
		{Src: comdut5port34, Dst: comswitch4port34},
		{Src: comdut5port35, Dst: comswitch4port35},

		{Src: comswitch4port36, Dst: comate5port1},
		{Src: comswitch4port37, Dst: comate5port2},
		{Src: comswitch4port38, Dst: comate5port3},
		{Src: comswitch4port39, Dst: comate5port4},
		{Src: comswitch4port40, Dst: comate5port5},
		{Src: comswitch4port41, Dst: comate5port6},
		{Src: comswitch4port42, Dst: comate5port7},
		{Src: comswitch4port43, Dst: comate5port8},
		{Src: comswitch4port44, Dst: comate5port9},
		{Src: comswitch4port45, Dst: comate5port10},
		{Src: comswitch4port46, Dst: comate5port11},
		{Src: comswitch4port47, Dst: comate5port12},
		{Src: comswitch4port48, Dst: comate5port13},
		{Src: comswitch4port49, Dst: comate5port14},
		{Src: comswitch4port50, Dst: comate5port15},
		{Src: comswitch4port51, Dst: comate5port16},
		{Src: comswitch4port52, Dst: comate5port17},
		{Src: comswitch4port53, Dst: comate5port18},
		{Src: comswitch4port54, Dst: comate5port19},
		{Src: comswitch4port55, Dst: comate5port20},
		{Src: comswitch4port56, Dst: comate5port21},
		{Src: comswitch4port57, Dst: comate5port22},
		{Src: comswitch4port58, Dst: comate5port23},
		{Src: comswitch4port59, Dst: comate5port24},
		{Src: comswitch4port60, Dst: comate5port25},
		{Src: comswitch4port61, Dst: comate5port26},
		{Src: comswitch4port62, Dst: comate5port27},
		{Src: comswitch4port63, Dst: comate5port28},
		{Src: comswitch4port64, Dst: comate5port29},
		{Src: comswitch4port65, Dst: comate5port30},
		{Src: comswitch4port66, Dst: comate5port31},
		{Src: comswitch4port67, Dst: comate5port32},
		{Src: comswitch4port68, Dst: comate5port33},
		{Src: comswitch4port69, Dst: comate5port34},
		{Src: comswitch4port70, Dst: comate5port35},

		{Src: comdut6port1, Dst: comswitch4port71},
		{Src: comdut6port2, Dst: comswitch4port72},
		{Src: comdut6port3, Dst: comswitch4port73},
		{Src: comdut6port4, Dst: comswitch4port74},
		{Src: comdut6port5, Dst: comswitch4port75},
		{Src: comdut6port6, Dst: comswitch4port76},
		{Src: comdut6port7, Dst: comswitch4port77},
		{Src: comdut6port8, Dst: comswitch4port78},
		{Src: comdut6port9, Dst: comswitch4port79},
		{Src: comdut6port10, Dst: comswitch4port80},
		{Src: comdut6port11, Dst: comswitch4port81},
		{Src: comdut6port12, Dst: comswitch4port82},
		{Src: comdut6port13, Dst: comswitch4port83},
		{Src: comdut6port14, Dst: comswitch4port84},
		{Src: comdut6port15, Dst: comswitch4port85},
		{Src: comdut6port16, Dst: comswitch4port86},
		{Src: comdut6port17, Dst: comswitch4port87},
		{Src: comdut6port18, Dst: comswitch4port88},
		{Src: comdut6port19, Dst: comswitch4port89},
		{Src: comdut6port20, Dst: comswitch4port90},

		{Src: comswitch4port91, Dst: comate6port1},
		{Src: comswitch4port92, Dst: comate6port2},
		{Src: comswitch4port93, Dst: comate6port3},
		{Src: comswitch4port94, Dst: comate6port4},
		{Src: comswitch4port95, Dst: comate6port5},
		{Src: comswitch4port96, Dst: comate6port6},
		{Src: comswitch4port97, Dst: comate6port7},
		{Src: comswitch4port98, Dst: comate6port8},
		{Src: comswitch4port99, Dst: comate6port9},
		{Src: comswitch4port100, Dst: comate6port10},
		{Src: comswitch4port101, Dst: comate6port11},
		{Src: comswitch4port102, Dst: comate6port12},
		{Src: comswitch4port103, Dst: comate6port13},
		{Src: comswitch4port104, Dst: comate6port14},
		{Src: comswitch4port105, Dst: comate6port15},
		{Src: comswitch4port106, Dst: comate6port16},
		{Src: comswitch4port107, Dst: comate6port17},
		{Src: comswitch4port108, Dst: comate6port18},
		{Src: comswitch4port109, Dst: comate6port19},
		{Src: comswitch4port110, Dst: comate6port20},
		{Src: comswitch4port111, Dst: comswitch5port141},

		{Src: comdut5port36, Dst: comswitch5port1},
		{Src: comdut5port37, Dst: comswitch5port2},
		{Src: comdut5port38, Dst: comswitch5port3},
		{Src: comdut5port39, Dst: comswitch5port4},
		{Src: comdut5port40, Dst: comswitch5port5},
		{Src: comdut5port41, Dst: comswitch5port6},
		{Src: comdut5port42, Dst: comswitch5port7},
		{Src: comdut5port43, Dst: comswitch5port8},
		{Src: comdut5port44, Dst: comswitch5port9},
		{Src: comdut5port45, Dst: comswitch5port10},
		{Src: comdut5port46, Dst: comswitch5port11},
		{Src: comdut5port47, Dst: comswitch5port12},
		{Src: comdut5port48, Dst: comswitch5port13},
		{Src: comdut5port49, Dst: comswitch5port14},
		{Src: comdut5port50, Dst: comswitch5port15},
		{Src: comdut5port51, Dst: comswitch5port16},
		{Src: comdut5port52, Dst: comswitch5port17},
		{Src: comdut5port53, Dst: comswitch5port18},
		{Src: comdut5port54, Dst: comswitch5port19},
		{Src: comdut5port55, Dst: comswitch5port20},
		{Src: comdut5port56, Dst: comswitch5port21},
		{Src: comdut5port57, Dst: comswitch5port22},
		{Src: comdut5port58, Dst: comswitch5port23},
		{Src: comdut5port59, Dst: comswitch5port24},
		{Src: comdut5port60, Dst: comswitch5port25},
		{Src: comdut5port61, Dst: comswitch5port26},
		{Src: comdut5port62, Dst: comswitch5port27},
		{Src: comdut5port63, Dst: comswitch5port28},
		{Src: comdut5port64, Dst: comswitch5port29},
		{Src: comdut5port65, Dst: comswitch5port30},
		{Src: comdut5port66, Dst: comswitch5port31},
		{Src: comdut5port67, Dst: comswitch5port32},
		{Src: comdut5port68, Dst: comswitch5port33},
		{Src: comdut5port69, Dst: comswitch5port34},
		{Src: comdut5port70, Dst: comswitch5port35},

		{Src: comswitch5port36, Dst: comate5port36},
		{Src: comswitch5port37, Dst: comate5port37},
		{Src: comswitch5port38, Dst: comate5port38},
		{Src: comswitch5port39, Dst: comate5port39},
		{Src: comswitch5port40, Dst: comate5port40},
		{Src: comswitch5port41, Dst: comate5port41},
		{Src: comswitch5port42, Dst: comate5port42},
		{Src: comswitch5port43, Dst: comate5port43},
		{Src: comswitch5port44, Dst: comate5port44},
		{Src: comswitch5port45, Dst: comate5port45},
		{Src: comswitch5port46, Dst: comate5port46},
		{Src: comswitch5port47, Dst: comate5port47},
		{Src: comswitch5port48, Dst: comate5port48},
		{Src: comswitch5port49, Dst: comate5port49},
		{Src: comswitch5port50, Dst: comate5port50},
		{Src: comswitch5port51, Dst: comate5port51},
		{Src: comswitch5port52, Dst: comate5port52},
		{Src: comswitch5port53, Dst: comate5port53},
		{Src: comswitch5port54, Dst: comate5port54},
		{Src: comswitch5port55, Dst: comate5port55},
		{Src: comswitch5port56, Dst: comate5port56},
		{Src: comswitch5port57, Dst: comate5port57},
		{Src: comswitch5port58, Dst: comate5port58},
		{Src: comswitch5port59, Dst: comate5port59},
		{Src: comswitch5port60, Dst: comate5port60},
		{Src: comswitch5port61, Dst: comate5port61},
		{Src: comswitch5port62, Dst: comate5port62},
		{Src: comswitch5port63, Dst: comate5port63},
		{Src: comswitch5port64, Dst: comate5port64},
		{Src: comswitch5port65, Dst: comate5port65},
		{Src: comswitch5port66, Dst: comate5port66},
		{Src: comswitch5port67, Dst: comate5port67},
		{Src: comswitch5port68, Dst: comate5port68},
		{Src: comswitch5port69, Dst: comate5port69},
		{Src: comswitch5port70, Dst: comate5port70},

		{Src: comdut6port21, Dst: comswitch5port71},
		{Src: comdut6port22, Dst: comswitch5port72},
		{Src: comdut6port23, Dst: comswitch5port73},
		{Src: comdut6port24, Dst: comswitch5port74},
		{Src: comdut6port25, Dst: comswitch5port75},
		{Src: comdut6port26, Dst: comswitch5port76},
		{Src: comdut6port27, Dst: comswitch5port77},
		{Src: comdut6port28, Dst: comswitch5port78},
		{Src: comdut6port29, Dst: comswitch5port79},
		{Src: comdut6port30, Dst: comswitch5port80},
		{Src: comdut6port31, Dst: comswitch5port81},
		{Src: comdut6port32, Dst: comswitch5port82},
		{Src: comdut6port33, Dst: comswitch5port83},
		{Src: comdut6port34, Dst: comswitch5port84},
		{Src: comdut6port35, Dst: comswitch5port85},
		{Src: comdut6port36, Dst: comswitch5port86},
		{Src: comdut6port37, Dst: comswitch5port87},
		{Src: comdut6port38, Dst: comswitch5port88},
		{Src: comdut6port39, Dst: comswitch5port89},
		{Src: comdut6port40, Dst: comswitch5port90},

		{Src: comswitch5port91, Dst: comate6port21},
		{Src: comswitch5port92, Dst: comate6port22},
		{Src: comswitch5port93, Dst: comate6port23},
		{Src: comswitch5port94, Dst: comate6port24},
		{Src: comswitch5port95, Dst: comate6port25},
		{Src: comswitch5port96, Dst: comate6port26},
		{Src: comswitch5port97, Dst: comate6port27},
		{Src: comswitch5port98, Dst: comate6port28},
		{Src: comswitch5port99, Dst: comate6port29},
		{Src: comswitch5port100, Dst: comate6port30},
		{Src: comswitch5port101, Dst: comate6port31},
		{Src: comswitch5port102, Dst: comate6port32},
		{Src: comswitch5port103, Dst: comate6port33},
		{Src: comswitch5port104, Dst: comate6port34},
		{Src: comswitch5port105, Dst: comate6port35},
		{Src: comswitch5port106, Dst: comate6port36},
		{Src: comswitch5port107, Dst: comate6port37},
		{Src: comswitch5port108, Dst: comate6port38},
		{Src: comswitch5port109, Dst: comate6port39},
		{Src: comswitch5port110, Dst: comate6port40},

		{Src: comdut6port41, Dst: comate6port41},
		{Src: comdut6port42, Dst: comate6port42},
		{Src: comdut6port43, Dst: comate6port43},
		{Src: comdut6port44, Dst: comate6port44},
		{Src: comdut6port45, Dst: comate6port45},
		{Src: comdut6port46, Dst: comate6port46},
		{Src: comdut6port47, Dst: comate6port47},
		{Src: comdut6port48, Dst: comate6port48},
		{Src: comdut6port49, Dst: comate6port49},
		{Src: comdut6port50, Dst: comate6port50},
		{Src: comdut6port51, Dst: comate6port51},
		{Src: comdut6port52, Dst: comate6port52},
		{Src: comdut6port53, Dst: comate6port53},
		{Src: comdut6port54, Dst: comate6port54},
		{Src: comdut6port55, Dst: comate6port55},
		{Src: comdut6port56, Dst: comate6port56},
		{Src: comdut6port57, Dst: comate6port57},
		{Src: comdut6port58, Dst: comate6port58},
		{Src: comdut6port59, Dst: comate6port59},
		{Src: comdut6port60, Dst: comate6port60},
		{Src: comdut6port61, Dst: comate6port61},
		{Src: comdut6port62, Dst: comate6port62},
		{Src: comdut6port63, Dst: comate6port63},
		{Src: comdut6port64, Dst: comate6port64},
		{Src: comdut6port65, Dst: comate6port65},
		{Src: comdut6port66, Dst: comate6port66},
		{Src: comdut6port67, Dst: comate6port67},
		{Src: comdut6port68, Dst: comate6port68},
		{Src: comdut6port69, Dst: comate6port69},
		{Src: comdut6port70, Dst: comate6port70},
	},
}

var (
	abscomdut1port1  = &AbstractPort{Desc: "absdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port2  = &AbstractPort{Desc: "absdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port3  = &AbstractPort{Desc: "absdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port4  = &AbstractPort{Desc: "absdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port5  = &AbstractPort{Desc: "absdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port6  = &AbstractPort{Desc: "absdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port7  = &AbstractPort{Desc: "absdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port8  = &AbstractPort{Desc: "absdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port9  = &AbstractPort{Desc: "absdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port10 = &AbstractPort{Desc: "absdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port11 = &AbstractPort{Desc: "absdut1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port12 = &AbstractPort{Desc: "absdut1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port13 = &AbstractPort{Desc: "absdut1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port14 = &AbstractPort{Desc: "absdut1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port15 = &AbstractPort{Desc: "absdut1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port16 = &AbstractPort{Desc: "absdut1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port17 = &AbstractPort{Desc: "absdut1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port18 = &AbstractPort{Desc: "absdut1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port19 = &AbstractPort{Desc: "absdut1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port20 = &AbstractPort{Desc: "absdut1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port21 = &AbstractPort{Desc: "absdut1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port22 = &AbstractPort{Desc: "absdut1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port23 = &AbstractPort{Desc: "absdut1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port24 = &AbstractPort{Desc: "absdut1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port25 = &AbstractPort{Desc: "absdut1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port26 = &AbstractPort{Desc: "absdut1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port27 = &AbstractPort{Desc: "absdut1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port28 = &AbstractPort{Desc: "absdut1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port29 = &AbstractPort{Desc: "absdut1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port30 = &AbstractPort{Desc: "absdut1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port31 = &AbstractPort{Desc: "absdut1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port32 = &AbstractPort{Desc: "absdut1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port33 = &AbstractPort{Desc: "absdut1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port34 = &AbstractPort{Desc: "absdut1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port35 = &AbstractPort{Desc: "absdut1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port36 = &AbstractPort{Desc: "absdut1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port37 = &AbstractPort{Desc: "absdut1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port38 = &AbstractPort{Desc: "absdut1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port39 = &AbstractPort{Desc: "absdut1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port40 = &AbstractPort{Desc: "absdut1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port41 = &AbstractPort{Desc: "absdut1:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port42 = &AbstractPort{Desc: "absdut1:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port43 = &AbstractPort{Desc: "absdut1:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port44 = &AbstractPort{Desc: "absdut1:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port45 = &AbstractPort{Desc: "absdut1:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port46 = &AbstractPort{Desc: "absdut1:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port47 = &AbstractPort{Desc: "absdut1:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port48 = &AbstractPort{Desc: "absdut1:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port49 = &AbstractPort{Desc: "absdut1:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port50 = &AbstractPort{Desc: "absdut1:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port51 = &AbstractPort{Desc: "absdut1:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port52 = &AbstractPort{Desc: "absdut1:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port53 = &AbstractPort{Desc: "absdut1:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port54 = &AbstractPort{Desc: "absdut1:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port55 = &AbstractPort{Desc: "absdut1:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port56 = &AbstractPort{Desc: "absdut1:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port57 = &AbstractPort{Desc: "absdut1:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port58 = &AbstractPort{Desc: "absdut1:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port59 = &AbstractPort{Desc: "absdut1:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port60 = &AbstractPort{Desc: "absdut1:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port61 = &AbstractPort{Desc: "absdut1:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port62 = &AbstractPort{Desc: "absdut1:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port63 = &AbstractPort{Desc: "absdut1:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port64 = &AbstractPort{Desc: "absdut1:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port65 = &AbstractPort{Desc: "absdut1:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port66 = &AbstractPort{Desc: "absdut1:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port67 = &AbstractPort{Desc: "absdut1:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port68 = &AbstractPort{Desc: "absdut1:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port69 = &AbstractPort{Desc: "absdut1:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut1port70 = &AbstractPort{Desc: "absdut1:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomdut1 = &AbstractNode{Desc: "absdut1", Ports: []*AbstractPort{
		abscomdut1port1,
		abscomdut1port2,
		abscomdut1port3,
		abscomdut1port4,
		abscomdut1port5,
		abscomdut1port6,
		abscomdut1port7,
		abscomdut1port8,
		abscomdut1port9,
		abscomdut1port10,
		abscomdut1port11,
		abscomdut1port12,
		abscomdut1port13,
		abscomdut1port14,
		abscomdut1port15,
		abscomdut1port16,
		abscomdut1port17,
		abscomdut1port18,
		abscomdut1port19,
		abscomdut1port20,
		abscomdut1port21,
		abscomdut1port22,
		abscomdut1port23,
		abscomdut1port24,
		abscomdut1port25,
		abscomdut1port26,
		abscomdut1port27,
		abscomdut1port28,
		abscomdut1port29,
		abscomdut1port30,
		abscomdut1port31,
		abscomdut1port32,
		abscomdut1port33,
		abscomdut1port34,
		abscomdut1port35,
		abscomdut1port36,
		abscomdut1port37,
		abscomdut1port38,
		abscomdut1port39,
		abscomdut1port40,
		abscomdut1port41,
		abscomdut1port42,
		abscomdut1port43,
		abscomdut1port44,
		abscomdut1port45,
		abscomdut1port46,
		abscomdut1port47,
		abscomdut1port48,
		abscomdut1port49,
		abscomdut1port50,
		abscomdut1port51,
		abscomdut1port52,
		abscomdut1port53,
		abscomdut1port54,
		abscomdut1port55,
		abscomdut1port56,
		abscomdut1port57,
		abscomdut1port58,
		abscomdut1port59,
		abscomdut1port60,
		abscomdut1port61,
		abscomdut1port62,
		abscomdut1port63,
		abscomdut1port64,
		abscomdut1port65,
		abscomdut1port66,
		abscomdut1port67,
		abscomdut1port68,
		abscomdut1port69,
		abscomdut1port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abscomdut2port1  = &AbstractPort{Desc: "absdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port2  = &AbstractPort{Desc: "absdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port3  = &AbstractPort{Desc: "absdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port4  = &AbstractPort{Desc: "absdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port5  = &AbstractPort{Desc: "absdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port6  = &AbstractPort{Desc: "absdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port7  = &AbstractPort{Desc: "absdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port8  = &AbstractPort{Desc: "absdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port9  = &AbstractPort{Desc: "absdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port10 = &AbstractPort{Desc: "absdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port11 = &AbstractPort{Desc: "absdut2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port12 = &AbstractPort{Desc: "absdut2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port13 = &AbstractPort{Desc: "absdut2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port14 = &AbstractPort{Desc: "absdut2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port15 = &AbstractPort{Desc: "absdut2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port16 = &AbstractPort{Desc: "absdut2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port17 = &AbstractPort{Desc: "absdut2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port18 = &AbstractPort{Desc: "absdut2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port19 = &AbstractPort{Desc: "absdut2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port20 = &AbstractPort{Desc: "absdut2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port21 = &AbstractPort{Desc: "absdut2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port22 = &AbstractPort{Desc: "absdut2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port23 = &AbstractPort{Desc: "absdut2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port24 = &AbstractPort{Desc: "absdut2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port25 = &AbstractPort{Desc: "absdut2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port26 = &AbstractPort{Desc: "absdut2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port27 = &AbstractPort{Desc: "absdut2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port28 = &AbstractPort{Desc: "absdut2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port29 = &AbstractPort{Desc: "absdut2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port30 = &AbstractPort{Desc: "absdut2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port31 = &AbstractPort{Desc: "absdut2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port32 = &AbstractPort{Desc: "absdut2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port33 = &AbstractPort{Desc: "absdut2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port34 = &AbstractPort{Desc: "absdut2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port35 = &AbstractPort{Desc: "absdut2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port36 = &AbstractPort{Desc: "absdut2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port37 = &AbstractPort{Desc: "absdut2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port38 = &AbstractPort{Desc: "absdut2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port39 = &AbstractPort{Desc: "absdut2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port40 = &AbstractPort{Desc: "absdut2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port41 = &AbstractPort{Desc: "absdut2:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port42 = &AbstractPort{Desc: "absdut2:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port43 = &AbstractPort{Desc: "absdut2:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port44 = &AbstractPort{Desc: "absdut2:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port45 = &AbstractPort{Desc: "absdut2:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port46 = &AbstractPort{Desc: "absdut2:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port47 = &AbstractPort{Desc: "absdut2:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port48 = &AbstractPort{Desc: "absdut2:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port49 = &AbstractPort{Desc: "absdut2:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port50 = &AbstractPort{Desc: "absdut2:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port51 = &AbstractPort{Desc: "absdut2:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port52 = &AbstractPort{Desc: "absdut2:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port53 = &AbstractPort{Desc: "absdut2:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port54 = &AbstractPort{Desc: "absdut2:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port55 = &AbstractPort{Desc: "absdut2:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port56 = &AbstractPort{Desc: "absdut2:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port57 = &AbstractPort{Desc: "absdut2:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port58 = &AbstractPort{Desc: "absdut2:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port59 = &AbstractPort{Desc: "absdut2:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port60 = &AbstractPort{Desc: "absdut2:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port61 = &AbstractPort{Desc: "absdut2:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port62 = &AbstractPort{Desc: "absdut2:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port63 = &AbstractPort{Desc: "absdut2:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port64 = &AbstractPort{Desc: "absdut2:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port65 = &AbstractPort{Desc: "absdut2:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port66 = &AbstractPort{Desc: "absdut2:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port67 = &AbstractPort{Desc: "absdut2:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port68 = &AbstractPort{Desc: "absdut2:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port69 = &AbstractPort{Desc: "absdut2:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut2port70 = &AbstractPort{Desc: "absdut2:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomdut2 = &AbstractNode{Desc: "absdut2", Ports: []*AbstractPort{
		abscomdut2port1,
		abscomdut2port2,
		abscomdut2port3,
		abscomdut2port4,
		abscomdut2port5,
		abscomdut2port6,
		abscomdut2port7,
		abscomdut2port8,
		abscomdut2port9,
		abscomdut2port10,
		abscomdut2port11,
		abscomdut2port12,
		abscomdut2port13,
		abscomdut2port14,
		abscomdut2port15,
		abscomdut2port16,
		abscomdut2port17,
		abscomdut2port18,
		abscomdut2port19,
		abscomdut2port20,
		abscomdut2port21,
		abscomdut2port22,
		abscomdut2port23,
		abscomdut2port24,
		abscomdut2port25,
		abscomdut2port26,
		abscomdut2port27,
		abscomdut2port28,
		abscomdut2port29,
		abscomdut2port30,
		abscomdut2port31,
		abscomdut2port32,
		abscomdut2port33,
		abscomdut2port34,
		abscomdut2port35,
		abscomdut2port36,
		abscomdut2port37,
		abscomdut2port38,
		abscomdut2port39,
		abscomdut2port40,
		abscomdut2port41,
		abscomdut2port42,
		abscomdut2port43,
		abscomdut2port44,
		abscomdut2port45,
		abscomdut2port46,
		abscomdut2port47,
		abscomdut2port48,
		abscomdut2port49,
		abscomdut2port50,
		abscomdut2port51,
		abscomdut2port52,
		abscomdut2port53,
		abscomdut2port54,
		abscomdut2port55,
		abscomdut2port56,
		abscomdut2port57,
		abscomdut2port58,
		abscomdut2port59,
		abscomdut2port60,
		abscomdut2port61,
		abscomdut2port62,
		abscomdut2port63,
		abscomdut2port64,
		abscomdut2port65,
		abscomdut2port66,
		abscomdut2port67,
		abscomdut2port68,
		abscomdut2port69,
		abscomdut2port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abscomdut3port1  = &AbstractPort{Desc: "absdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port2  = &AbstractPort{Desc: "absdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port3  = &AbstractPort{Desc: "absdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port4  = &AbstractPort{Desc: "absdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port5  = &AbstractPort{Desc: "absdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port6  = &AbstractPort{Desc: "absdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port7  = &AbstractPort{Desc: "absdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port8  = &AbstractPort{Desc: "absdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port9  = &AbstractPort{Desc: "absdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port10 = &AbstractPort{Desc: "absdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port11 = &AbstractPort{Desc: "absdut3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port12 = &AbstractPort{Desc: "absdut3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port13 = &AbstractPort{Desc: "absdut3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port14 = &AbstractPort{Desc: "absdut3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port15 = &AbstractPort{Desc: "absdut3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port16 = &AbstractPort{Desc: "absdut3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port17 = &AbstractPort{Desc: "absdut3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port18 = &AbstractPort{Desc: "absdut3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port19 = &AbstractPort{Desc: "absdut3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port20 = &AbstractPort{Desc: "absdut3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port21 = &AbstractPort{Desc: "absdut3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port22 = &AbstractPort{Desc: "absdut3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port23 = &AbstractPort{Desc: "absdut3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port24 = &AbstractPort{Desc: "absdut3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port25 = &AbstractPort{Desc: "absdut3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port26 = &AbstractPort{Desc: "absdut3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port27 = &AbstractPort{Desc: "absdut3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port28 = &AbstractPort{Desc: "absdut3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port29 = &AbstractPort{Desc: "absdut3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port30 = &AbstractPort{Desc: "absdut3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port31 = &AbstractPort{Desc: "absdut3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port32 = &AbstractPort{Desc: "absdut3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port33 = &AbstractPort{Desc: "absdut3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port34 = &AbstractPort{Desc: "absdut3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port35 = &AbstractPort{Desc: "absdut3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port36 = &AbstractPort{Desc: "absdut3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port37 = &AbstractPort{Desc: "absdut3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port38 = &AbstractPort{Desc: "absdut3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port39 = &AbstractPort{Desc: "absdut3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port40 = &AbstractPort{Desc: "absdut3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port41 = &AbstractPort{Desc: "absdut3:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port42 = &AbstractPort{Desc: "absdut3:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port43 = &AbstractPort{Desc: "absdut3:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port44 = &AbstractPort{Desc: "absdut3:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port45 = &AbstractPort{Desc: "absdut3:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port46 = &AbstractPort{Desc: "absdut3:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port47 = &AbstractPort{Desc: "absdut3:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port48 = &AbstractPort{Desc: "absdut3:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port49 = &AbstractPort{Desc: "absdut3:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port50 = &AbstractPort{Desc: "absdut3:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port51 = &AbstractPort{Desc: "absdut3:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port52 = &AbstractPort{Desc: "absdut3:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port53 = &AbstractPort{Desc: "absdut3:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port54 = &AbstractPort{Desc: "absdut3:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port55 = &AbstractPort{Desc: "absdut3:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port56 = &AbstractPort{Desc: "absdut3:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port57 = &AbstractPort{Desc: "absdut3:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port58 = &AbstractPort{Desc: "absdut3:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port59 = &AbstractPort{Desc: "absdut3:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port60 = &AbstractPort{Desc: "absdut3:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port61 = &AbstractPort{Desc: "absdut3:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port62 = &AbstractPort{Desc: "absdut3:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port63 = &AbstractPort{Desc: "absdut3:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port64 = &AbstractPort{Desc: "absdut3:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port65 = &AbstractPort{Desc: "absdut3:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port66 = &AbstractPort{Desc: "absdut3:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port67 = &AbstractPort{Desc: "absdut3:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port68 = &AbstractPort{Desc: "absdut3:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port69 = &AbstractPort{Desc: "absdut3:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut3port70 = &AbstractPort{Desc: "absdut3:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomdut3 = &AbstractNode{Desc: "absdut3", Ports: []*AbstractPort{
		abscomdut3port1,
		abscomdut3port2,
		abscomdut3port3,
		abscomdut3port4,
		abscomdut3port5,
		abscomdut3port6,
		abscomdut3port7,
		abscomdut3port8,
		abscomdut3port9,
		abscomdut3port10,
		abscomdut3port11,
		abscomdut3port12,
		abscomdut3port13,
		abscomdut3port14,
		abscomdut3port15,
		abscomdut3port16,
		abscomdut3port17,
		abscomdut3port18,
		abscomdut3port19,
		abscomdut3port20,
		abscomdut3port21,
		abscomdut3port22,
		abscomdut3port23,
		abscomdut3port24,
		abscomdut3port25,
		abscomdut3port26,
		abscomdut3port27,
		abscomdut3port28,
		abscomdut3port29,
		abscomdut3port30,
		abscomdut3port31,
		abscomdut3port32,
		abscomdut3port33,
		abscomdut3port34,
		abscomdut3port35,
		abscomdut3port36,
		abscomdut3port37,
		abscomdut3port38,
		abscomdut3port39,
		abscomdut3port40,
		abscomdut3port41,
		abscomdut3port42,
		abscomdut3port43,
		abscomdut3port44,
		abscomdut3port45,
		abscomdut3port46,
		abscomdut3port47,
		abscomdut3port48,
		abscomdut3port49,
		abscomdut3port50,
		abscomdut3port51,
		abscomdut3port52,
		abscomdut3port53,
		abscomdut3port54,
		abscomdut3port55,
		abscomdut3port56,
		abscomdut3port57,
		abscomdut3port58,
		abscomdut3port59,
		abscomdut3port60,
		abscomdut3port61,
		abscomdut3port62,
		abscomdut3port63,
		abscomdut3port64,
		abscomdut3port65,
		abscomdut3port66,
		abscomdut3port67,
		abscomdut3port68,
		abscomdut3port69,
		abscomdut3port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abscomdut4port1  = &AbstractPort{Desc: "absdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port2  = &AbstractPort{Desc: "absdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port3  = &AbstractPort{Desc: "absdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port4  = &AbstractPort{Desc: "absdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port5  = &AbstractPort{Desc: "absdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port6  = &AbstractPort{Desc: "absdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port7  = &AbstractPort{Desc: "absdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port8  = &AbstractPort{Desc: "absdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port9  = &AbstractPort{Desc: "absdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port10 = &AbstractPort{Desc: "absdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port11 = &AbstractPort{Desc: "absdut4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port12 = &AbstractPort{Desc: "absdut4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port13 = &AbstractPort{Desc: "absdut4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port14 = &AbstractPort{Desc: "absdut4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port15 = &AbstractPort{Desc: "absdut4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port16 = &AbstractPort{Desc: "absdut4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port17 = &AbstractPort{Desc: "absdut4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port18 = &AbstractPort{Desc: "absdut4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port19 = &AbstractPort{Desc: "absdut4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port20 = &AbstractPort{Desc: "absdut4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port21 = &AbstractPort{Desc: "absdut4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port22 = &AbstractPort{Desc: "absdut4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port23 = &AbstractPort{Desc: "absdut4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port24 = &AbstractPort{Desc: "absdut4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port25 = &AbstractPort{Desc: "absdut4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port26 = &AbstractPort{Desc: "absdut4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port27 = &AbstractPort{Desc: "absdut4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port28 = &AbstractPort{Desc: "absdut4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port29 = &AbstractPort{Desc: "absdut4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port30 = &AbstractPort{Desc: "absdut4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port31 = &AbstractPort{Desc: "absdut4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port32 = &AbstractPort{Desc: "absdut4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port33 = &AbstractPort{Desc: "absdut4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port34 = &AbstractPort{Desc: "absdut4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port35 = &AbstractPort{Desc: "absdut4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port36 = &AbstractPort{Desc: "absdut4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port37 = &AbstractPort{Desc: "absdut4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port38 = &AbstractPort{Desc: "absdut4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port39 = &AbstractPort{Desc: "absdut4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port40 = &AbstractPort{Desc: "absdut4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port41 = &AbstractPort{Desc: "absdut4:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port42 = &AbstractPort{Desc: "absdut4:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port43 = &AbstractPort{Desc: "absdut4:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port44 = &AbstractPort{Desc: "absdut4:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port45 = &AbstractPort{Desc: "absdut4:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port46 = &AbstractPort{Desc: "absdut4:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port47 = &AbstractPort{Desc: "absdut4:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port48 = &AbstractPort{Desc: "absdut4:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port49 = &AbstractPort{Desc: "absdut4:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port50 = &AbstractPort{Desc: "absdut4:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port51 = &AbstractPort{Desc: "absdut4:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port52 = &AbstractPort{Desc: "absdut4:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port53 = &AbstractPort{Desc: "absdut4:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port54 = &AbstractPort{Desc: "absdut4:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port55 = &AbstractPort{Desc: "absdut4:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port56 = &AbstractPort{Desc: "absdut4:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port57 = &AbstractPort{Desc: "absdut4:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port58 = &AbstractPort{Desc: "absdut4:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port59 = &AbstractPort{Desc: "absdut4:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port60 = &AbstractPort{Desc: "absdut4:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port61 = &AbstractPort{Desc: "absdut4:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port62 = &AbstractPort{Desc: "absdut4:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port63 = &AbstractPort{Desc: "absdut4:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port64 = &AbstractPort{Desc: "absdut4:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port65 = &AbstractPort{Desc: "absdut4:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port66 = &AbstractPort{Desc: "absdut4:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port67 = &AbstractPort{Desc: "absdut4:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port68 = &AbstractPort{Desc: "absdut4:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port69 = &AbstractPort{Desc: "absdut4:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut4port70 = &AbstractPort{Desc: "absdut4:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomdut4 = &AbstractNode{Desc: "absdut4", Ports: []*AbstractPort{
		abscomdut4port1,
		abscomdut4port2,
		abscomdut4port3,
		abscomdut4port4,
		abscomdut4port5,
		abscomdut4port6,
		abscomdut4port7,
		abscomdut4port8,
		abscomdut4port9,
		abscomdut4port10,
		abscomdut4port11,
		abscomdut4port12,
		abscomdut4port13,
		abscomdut4port14,
		abscomdut4port15,
		abscomdut4port16,
		abscomdut4port17,
		abscomdut4port18,
		abscomdut4port19,
		abscomdut4port20,
		abscomdut4port21,
		abscomdut4port22,
		abscomdut4port23,
		abscomdut4port24,
		abscomdut4port25,
		abscomdut4port26,
		abscomdut4port27,
		abscomdut4port28,
		abscomdut4port29,
		abscomdut4port30,
		abscomdut4port31,
		abscomdut4port32,
		abscomdut4port33,
		abscomdut4port34,
		abscomdut4port35,
		abscomdut4port36,
		abscomdut4port37,
		abscomdut4port38,
		abscomdut4port39,
		abscomdut4port40,
		abscomdut4port41,
		abscomdut4port42,
		abscomdut4port43,
		abscomdut4port44,
		abscomdut4port45,
		abscomdut4port46,
		abscomdut4port47,
		abscomdut4port48,
		abscomdut4port49,
		abscomdut4port50,
		abscomdut4port51,
		abscomdut4port52,
		abscomdut4port53,
		abscomdut4port54,
		abscomdut4port55,
		abscomdut4port56,
		abscomdut4port57,
		abscomdut4port58,
		abscomdut4port59,
		abscomdut4port60,
		abscomdut4port61,
		abscomdut4port62,
		abscomdut4port63,
		abscomdut4port64,
		abscomdut4port65,
		abscomdut4port66,
		abscomdut4port67,
		abscomdut4port68,
		abscomdut4port69,
		abscomdut4port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abscomdut5port1  = &AbstractPort{Desc: "absdut5:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port2  = &AbstractPort{Desc: "absdut5:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port3  = &AbstractPort{Desc: "absdut5:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port4  = &AbstractPort{Desc: "absdut5:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port5  = &AbstractPort{Desc: "absdut5:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port6  = &AbstractPort{Desc: "absdut5:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port7  = &AbstractPort{Desc: "absdut5:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port8  = &AbstractPort{Desc: "absdut5:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port9  = &AbstractPort{Desc: "absdut5:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port10 = &AbstractPort{Desc: "absdut5:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port11 = &AbstractPort{Desc: "absdut5:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port12 = &AbstractPort{Desc: "absdut5:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port13 = &AbstractPort{Desc: "absdut5:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port14 = &AbstractPort{Desc: "absdut5:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port15 = &AbstractPort{Desc: "absdut5:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port16 = &AbstractPort{Desc: "absdut5:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port17 = &AbstractPort{Desc: "absdut5:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port18 = &AbstractPort{Desc: "absdut5:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port19 = &AbstractPort{Desc: "absdut5:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port20 = &AbstractPort{Desc: "absdut5:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port21 = &AbstractPort{Desc: "absdut5:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port22 = &AbstractPort{Desc: "absdut5:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port23 = &AbstractPort{Desc: "absdut5:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port24 = &AbstractPort{Desc: "absdut5:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port25 = &AbstractPort{Desc: "absdut5:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port26 = &AbstractPort{Desc: "absdut5:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port27 = &AbstractPort{Desc: "absdut5:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port28 = &AbstractPort{Desc: "absdut5:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port29 = &AbstractPort{Desc: "absdut5:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port30 = &AbstractPort{Desc: "absdut5:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port31 = &AbstractPort{Desc: "absdut5:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port32 = &AbstractPort{Desc: "absdut5:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port33 = &AbstractPort{Desc: "absdut5:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port34 = &AbstractPort{Desc: "absdut5:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port35 = &AbstractPort{Desc: "absdut5:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port36 = &AbstractPort{Desc: "absdut5:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port37 = &AbstractPort{Desc: "absdut5:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port38 = &AbstractPort{Desc: "absdut5:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port39 = &AbstractPort{Desc: "absdut5:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port40 = &AbstractPort{Desc: "absdut5:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port41 = &AbstractPort{Desc: "absdut5:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port42 = &AbstractPort{Desc: "absdut5:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port43 = &AbstractPort{Desc: "absdut5:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port44 = &AbstractPort{Desc: "absdut5:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port45 = &AbstractPort{Desc: "absdut5:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port46 = &AbstractPort{Desc: "absdut5:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port47 = &AbstractPort{Desc: "absdut5:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port48 = &AbstractPort{Desc: "absdut5:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port49 = &AbstractPort{Desc: "absdut5:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port50 = &AbstractPort{Desc: "absdut5:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port51 = &AbstractPort{Desc: "absdut5:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port52 = &AbstractPort{Desc: "absdut5:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port53 = &AbstractPort{Desc: "absdut5:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port54 = &AbstractPort{Desc: "absdut5:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port55 = &AbstractPort{Desc: "absdut5:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port56 = &AbstractPort{Desc: "absdut5:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port57 = &AbstractPort{Desc: "absdut5:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port58 = &AbstractPort{Desc: "absdut5:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port59 = &AbstractPort{Desc: "absdut5:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port60 = &AbstractPort{Desc: "absdut5:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port61 = &AbstractPort{Desc: "absdut5:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port62 = &AbstractPort{Desc: "absdut5:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port63 = &AbstractPort{Desc: "absdut5:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port64 = &AbstractPort{Desc: "absdut5:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port65 = &AbstractPort{Desc: "absdut5:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port66 = &AbstractPort{Desc: "absdut5:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port67 = &AbstractPort{Desc: "absdut5:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port68 = &AbstractPort{Desc: "absdut5:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port69 = &AbstractPort{Desc: "absdut5:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut5port70 = &AbstractPort{Desc: "absdut5:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomdut5 = &AbstractNode{Desc: "absdut5", Ports: []*AbstractPort{
		abscomdut5port1,
		abscomdut5port2,
		abscomdut5port3,
		abscomdut5port4,
		abscomdut5port5,
		abscomdut5port6,
		abscomdut5port7,
		abscomdut5port8,
		abscomdut5port9,
		abscomdut5port10,
		abscomdut5port11,
		abscomdut5port12,
		abscomdut5port13,
		abscomdut5port14,
		abscomdut5port15,
		abscomdut5port16,
		abscomdut5port17,
		abscomdut5port18,
		abscomdut5port19,
		abscomdut5port20,
		abscomdut5port21,
		abscomdut5port22,
		abscomdut5port23,
		abscomdut5port24,
		abscomdut5port25,
		abscomdut5port26,
		abscomdut5port27,
		abscomdut5port28,
		abscomdut5port29,
		abscomdut5port30,
		abscomdut5port31,
		abscomdut5port32,
		abscomdut5port33,
		abscomdut5port34,
		abscomdut5port35,
		abscomdut5port36,
		abscomdut5port37,
		abscomdut5port38,
		abscomdut5port39,
		abscomdut5port40,
		abscomdut5port41,
		abscomdut5port42,
		abscomdut5port43,
		abscomdut5port44,
		abscomdut5port45,
		abscomdut5port46,
		abscomdut5port47,
		abscomdut5port48,
		abscomdut5port49,
		abscomdut5port50,
		abscomdut5port51,
		abscomdut5port52,
		abscomdut5port53,
		abscomdut5port54,
		abscomdut5port55,
		abscomdut5port56,
		abscomdut5port57,
		abscomdut5port58,
		abscomdut5port59,
		abscomdut5port60,
		abscomdut5port61,
		abscomdut5port62,
		abscomdut5port63,
		abscomdut5port64,
		abscomdut5port65,
		abscomdut5port66,
		abscomdut5port67,
		abscomdut5port68,
		abscomdut5port69,
		abscomdut5port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abscomdut6port1  = &AbstractPort{Desc: "absdut6:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port2  = &AbstractPort{Desc: "absdut6:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port3  = &AbstractPort{Desc: "absdut6:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port4  = &AbstractPort{Desc: "absdut6:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port5  = &AbstractPort{Desc: "absdut6:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port6  = &AbstractPort{Desc: "absdut6:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port7  = &AbstractPort{Desc: "absdut6:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port8  = &AbstractPort{Desc: "absdut6:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port9  = &AbstractPort{Desc: "absdut6:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port10 = &AbstractPort{Desc: "absdut6:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port11 = &AbstractPort{Desc: "absdut6:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port12 = &AbstractPort{Desc: "absdut6:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port13 = &AbstractPort{Desc: "absdut6:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port14 = &AbstractPort{Desc: "absdut6:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port15 = &AbstractPort{Desc: "absdut6:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port16 = &AbstractPort{Desc: "absdut6:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port17 = &AbstractPort{Desc: "absdut6:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port18 = &AbstractPort{Desc: "absdut6:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port19 = &AbstractPort{Desc: "absdut6:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port20 = &AbstractPort{Desc: "absdut6:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port21 = &AbstractPort{Desc: "absdut6:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port22 = &AbstractPort{Desc: "absdut6:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port23 = &AbstractPort{Desc: "absdut6:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port24 = &AbstractPort{Desc: "absdut6:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port25 = &AbstractPort{Desc: "absdut6:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port26 = &AbstractPort{Desc: "absdut6:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port27 = &AbstractPort{Desc: "absdut6:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port28 = &AbstractPort{Desc: "absdut6:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port29 = &AbstractPort{Desc: "absdut6:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port30 = &AbstractPort{Desc: "absdut6:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port31 = &AbstractPort{Desc: "absdut6:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port32 = &AbstractPort{Desc: "absdut6:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port33 = &AbstractPort{Desc: "absdut6:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port34 = &AbstractPort{Desc: "absdut6:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port35 = &AbstractPort{Desc: "absdut6:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port36 = &AbstractPort{Desc: "absdut6:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port37 = &AbstractPort{Desc: "absdut6:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port38 = &AbstractPort{Desc: "absdut6:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port39 = &AbstractPort{Desc: "absdut6:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port40 = &AbstractPort{Desc: "absdut6:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port41 = &AbstractPort{Desc: "absdut6:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port42 = &AbstractPort{Desc: "absdut6:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port43 = &AbstractPort{Desc: "absdut6:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port44 = &AbstractPort{Desc: "absdut6:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port45 = &AbstractPort{Desc: "absdut6:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port46 = &AbstractPort{Desc: "absdut6:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port47 = &AbstractPort{Desc: "absdut6:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port48 = &AbstractPort{Desc: "absdut6:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port49 = &AbstractPort{Desc: "absdut6:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port50 = &AbstractPort{Desc: "absdut6:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port51 = &AbstractPort{Desc: "absdut6:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port52 = &AbstractPort{Desc: "absdut6:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port53 = &AbstractPort{Desc: "absdut6:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port54 = &AbstractPort{Desc: "absdut6:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port55 = &AbstractPort{Desc: "absdut6:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port56 = &AbstractPort{Desc: "absdut6:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port57 = &AbstractPort{Desc: "absdut6:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port58 = &AbstractPort{Desc: "absdut6:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port59 = &AbstractPort{Desc: "absdut6:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port60 = &AbstractPort{Desc: "absdut6:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port61 = &AbstractPort{Desc: "absdut6:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port62 = &AbstractPort{Desc: "absdut6:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port63 = &AbstractPort{Desc: "absdut6:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port64 = &AbstractPort{Desc: "absdut6:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port65 = &AbstractPort{Desc: "absdut6:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port66 = &AbstractPort{Desc: "absdut6:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port67 = &AbstractPort{Desc: "absdut6:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port68 = &AbstractPort{Desc: "absdut6:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port69 = &AbstractPort{Desc: "absdut6:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomdut6port70 = &AbstractPort{Desc: "absdut6:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomdut6 = &AbstractNode{Desc: "absdut6", Ports: []*AbstractPort{
		abscomdut6port1,
		abscomdut6port2,
		abscomdut6port3,
		abscomdut6port4,
		abscomdut6port5,
		abscomdut6port6,
		abscomdut6port7,
		abscomdut6port8,
		abscomdut6port9,
		abscomdut6port10,
		abscomdut6port11,
		abscomdut6port12,
		abscomdut6port13,
		abscomdut6port14,
		abscomdut6port15,
		abscomdut6port16,
		abscomdut6port17,
		abscomdut6port18,
		abscomdut6port19,
		abscomdut6port20,
		abscomdut6port21,
		abscomdut6port22,
		abscomdut6port23,
		abscomdut6port24,
		abscomdut6port25,
		abscomdut6port26,
		abscomdut6port27,
		abscomdut6port28,
		abscomdut6port29,
		abscomdut6port30,
		abscomdut6port31,
		abscomdut6port32,
		abscomdut6port33,
		abscomdut6port34,
		abscomdut6port35,
		abscomdut6port36,
		abscomdut6port37,
		abscomdut6port38,
		abscomdut6port39,
		abscomdut6port40,
		abscomdut6port41,
		abscomdut6port42,
		abscomdut6port43,
		abscomdut6port44,
		abscomdut6port45,
		abscomdut6port46,
		abscomdut6port47,
		abscomdut6port48,
		abscomdut6port49,
		abscomdut6port50,
		abscomdut6port51,
		abscomdut6port52,
		abscomdut6port53,
		abscomdut6port54,
		abscomdut6port55,
		abscomdut6port56,
		abscomdut6port57,
		abscomdut6port58,
		abscomdut6port59,
		abscomdut6port60,
		abscomdut6port61,
		abscomdut6port62,
		abscomdut6port63,
		abscomdut6port64,
		abscomdut6port65,
		abscomdut6port66,
		abscomdut6port67,
		abscomdut6port68,
		abscomdut6port69,
		abscomdut6port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	abscomate1port1  = &AbstractPort{Desc: "absate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port2  = &AbstractPort{Desc: "absate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port3  = &AbstractPort{Desc: "absate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port4  = &AbstractPort{Desc: "absate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port5  = &AbstractPort{Desc: "absate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port6  = &AbstractPort{Desc: "absate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port7  = &AbstractPort{Desc: "absate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port8  = &AbstractPort{Desc: "absate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port9  = &AbstractPort{Desc: "absate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port10 = &AbstractPort{Desc: "absate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port11 = &AbstractPort{Desc: "absate1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port12 = &AbstractPort{Desc: "absate1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port13 = &AbstractPort{Desc: "absate1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port14 = &AbstractPort{Desc: "absate1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port15 = &AbstractPort{Desc: "absate1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port16 = &AbstractPort{Desc: "absate1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port17 = &AbstractPort{Desc: "absate1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port18 = &AbstractPort{Desc: "absate1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port19 = &AbstractPort{Desc: "absate1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port20 = &AbstractPort{Desc: "absate1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port21 = &AbstractPort{Desc: "absate1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port22 = &AbstractPort{Desc: "absate1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port23 = &AbstractPort{Desc: "absate1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port24 = &AbstractPort{Desc: "absate1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port25 = &AbstractPort{Desc: "absate1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port26 = &AbstractPort{Desc: "absate1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port27 = &AbstractPort{Desc: "absate1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port28 = &AbstractPort{Desc: "absate1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port29 = &AbstractPort{Desc: "absate1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port30 = &AbstractPort{Desc: "absate1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port31 = &AbstractPort{Desc: "absate1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port32 = &AbstractPort{Desc: "absate1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port33 = &AbstractPort{Desc: "absate1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port34 = &AbstractPort{Desc: "absate1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port35 = &AbstractPort{Desc: "absate1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port36 = &AbstractPort{Desc: "absate1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port37 = &AbstractPort{Desc: "absate1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port38 = &AbstractPort{Desc: "absate1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port39 = &AbstractPort{Desc: "absate1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port40 = &AbstractPort{Desc: "absate1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port41 = &AbstractPort{Desc: "absate1:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port42 = &AbstractPort{Desc: "absate1:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port43 = &AbstractPort{Desc: "absate1:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port44 = &AbstractPort{Desc: "absate1:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port45 = &AbstractPort{Desc: "absate1:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port46 = &AbstractPort{Desc: "absate1:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port47 = &AbstractPort{Desc: "absate1:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port48 = &AbstractPort{Desc: "absate1:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port49 = &AbstractPort{Desc: "absate1:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port50 = &AbstractPort{Desc: "absate1:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port51 = &AbstractPort{Desc: "absate1:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port52 = &AbstractPort{Desc: "absate1:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port53 = &AbstractPort{Desc: "absate1:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port54 = &AbstractPort{Desc: "absate1:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port55 = &AbstractPort{Desc: "absate1:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port56 = &AbstractPort{Desc: "absate1:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port57 = &AbstractPort{Desc: "absate1:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port58 = &AbstractPort{Desc: "absate1:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port59 = &AbstractPort{Desc: "absate1:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port60 = &AbstractPort{Desc: "absate1:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port61 = &AbstractPort{Desc: "absate1:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port62 = &AbstractPort{Desc: "absate1:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port63 = &AbstractPort{Desc: "absate1:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port64 = &AbstractPort{Desc: "absate1:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port65 = &AbstractPort{Desc: "absate1:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port66 = &AbstractPort{Desc: "absate1:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port67 = &AbstractPort{Desc: "absate1:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port68 = &AbstractPort{Desc: "absate1:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port69 = &AbstractPort{Desc: "absate1:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate1port70 = &AbstractPort{Desc: "absate1:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomate1 = &AbstractNode{Desc: "absate1", Ports: []*AbstractPort{
		abscomate1port1,
		abscomate1port2,
		abscomate1port3,
		abscomate1port4,
		abscomate1port5,
		abscomate1port6,
		abscomate1port7,
		abscomate1port8,
		abscomate1port9,
		abscomate1port10,
		abscomate1port11,
		abscomate1port12,
		abscomate1port13,
		abscomate1port14,
		abscomate1port15,
		abscomate1port16,
		abscomate1port17,
		abscomate1port18,
		abscomate1port19,
		abscomate1port20,
		abscomate1port21,
		abscomate1port22,
		abscomate1port23,
		abscomate1port24,
		abscomate1port25,
		abscomate1port26,
		abscomate1port27,
		abscomate1port28,
		abscomate1port29,
		abscomate1port30,
		abscomate1port31,
		abscomate1port32,
		abscomate1port33,
		abscomate1port34,
		abscomate1port35,
		abscomate1port36,
		abscomate1port37,
		abscomate1port38,
		abscomate1port39,
		abscomate1port40,
		abscomate1port41,
		abscomate1port42,
		abscomate1port43,
		abscomate1port44,
		abscomate1port45,
		abscomate1port46,
		abscomate1port47,
		abscomate1port48,
		abscomate1port49,
		abscomate1port50,
		abscomate1port51,
		abscomate1port52,
		abscomate1port53,
		abscomate1port54,
		abscomate1port55,
		abscomate1port56,
		abscomate1port57,
		abscomate1port58,
		abscomate1port59,
		abscomate1port60,
		abscomate1port61,
		abscomate1port62,
		abscomate1port63,
		abscomate1port64,
		abscomate1port65,
		abscomate1port66,
		abscomate1port67,
		abscomate1port68,
		abscomate1port69,
		abscomate1port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abscomate2port1  = &AbstractPort{Desc: "absate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port2  = &AbstractPort{Desc: "absate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port3  = &AbstractPort{Desc: "absate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port4  = &AbstractPort{Desc: "absate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port5  = &AbstractPort{Desc: "absate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port6  = &AbstractPort{Desc: "absate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port7  = &AbstractPort{Desc: "absate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port8  = &AbstractPort{Desc: "absate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port9  = &AbstractPort{Desc: "absate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port10 = &AbstractPort{Desc: "absate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port11 = &AbstractPort{Desc: "absate2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port12 = &AbstractPort{Desc: "absate2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port13 = &AbstractPort{Desc: "absate2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port14 = &AbstractPort{Desc: "absate2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port15 = &AbstractPort{Desc: "absate2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port16 = &AbstractPort{Desc: "absate2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port17 = &AbstractPort{Desc: "absate2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port18 = &AbstractPort{Desc: "absate2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port19 = &AbstractPort{Desc: "absate2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port20 = &AbstractPort{Desc: "absate2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port21 = &AbstractPort{Desc: "absate2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port22 = &AbstractPort{Desc: "absate2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port23 = &AbstractPort{Desc: "absate2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port24 = &AbstractPort{Desc: "absate2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port25 = &AbstractPort{Desc: "absate2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port26 = &AbstractPort{Desc: "absate2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port27 = &AbstractPort{Desc: "absate2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port28 = &AbstractPort{Desc: "absate2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port29 = &AbstractPort{Desc: "absate2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port30 = &AbstractPort{Desc: "absate2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port31 = &AbstractPort{Desc: "absate2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port32 = &AbstractPort{Desc: "absate2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port33 = &AbstractPort{Desc: "absate2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port34 = &AbstractPort{Desc: "absate2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port35 = &AbstractPort{Desc: "absate2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port36 = &AbstractPort{Desc: "absate2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port37 = &AbstractPort{Desc: "absate2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port38 = &AbstractPort{Desc: "absate2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port39 = &AbstractPort{Desc: "absate2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port40 = &AbstractPort{Desc: "absate2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port41 = &AbstractPort{Desc: "absate2:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port42 = &AbstractPort{Desc: "absate2:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port43 = &AbstractPort{Desc: "absate2:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port44 = &AbstractPort{Desc: "absate2:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port45 = &AbstractPort{Desc: "absate2:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port46 = &AbstractPort{Desc: "absate2:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port47 = &AbstractPort{Desc: "absate2:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port48 = &AbstractPort{Desc: "absate2:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port49 = &AbstractPort{Desc: "absate2:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port50 = &AbstractPort{Desc: "absate2:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port51 = &AbstractPort{Desc: "absate2:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port52 = &AbstractPort{Desc: "absate2:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port53 = &AbstractPort{Desc: "absate2:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port54 = &AbstractPort{Desc: "absate2:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port55 = &AbstractPort{Desc: "absate2:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port56 = &AbstractPort{Desc: "absate2:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port57 = &AbstractPort{Desc: "absate2:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port58 = &AbstractPort{Desc: "absate2:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port59 = &AbstractPort{Desc: "absate2:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port60 = &AbstractPort{Desc: "absate2:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port61 = &AbstractPort{Desc: "absate2:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port62 = &AbstractPort{Desc: "absate2:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port63 = &AbstractPort{Desc: "absate2:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port64 = &AbstractPort{Desc: "absate2:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port65 = &AbstractPort{Desc: "absate2:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port66 = &AbstractPort{Desc: "absate2:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port67 = &AbstractPort{Desc: "absate2:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port68 = &AbstractPort{Desc: "absate2:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port69 = &AbstractPort{Desc: "absate2:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate2port70 = &AbstractPort{Desc: "absate2:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomate2 = &AbstractNode{Desc: "absate2", Ports: []*AbstractPort{
		abscomate2port1,
		abscomate2port2,
		abscomate2port3,
		abscomate2port4,
		abscomate2port5,
		abscomate2port6,
		abscomate2port7,
		abscomate2port8,
		abscomate2port9,
		abscomate2port10,
		abscomate2port11,
		abscomate2port12,
		abscomate2port13,
		abscomate2port14,
		abscomate2port15,
		abscomate2port16,
		abscomate2port17,
		abscomate2port18,
		abscomate2port19,
		abscomate2port20,
		abscomate2port21,
		abscomate2port22,
		abscomate2port23,
		abscomate2port24,
		abscomate2port25,
		abscomate2port26,
		abscomate2port27,
		abscomate2port28,
		abscomate2port29,
		abscomate2port30,
		abscomate2port31,
		abscomate2port32,
		abscomate2port33,
		abscomate2port34,
		abscomate2port35,
		abscomate2port36,
		abscomate2port37,
		abscomate2port38,
		abscomate2port39,
		abscomate2port40,
		abscomate2port41,
		abscomate2port42,
		abscomate2port43,
		abscomate2port44,
		abscomate2port45,
		abscomate2port46,
		abscomate2port47,
		abscomate2port48,
		abscomate2port49,
		abscomate2port50,
		abscomate2port51,
		abscomate2port52,
		abscomate2port53,
		abscomate2port54,
		abscomate2port55,
		abscomate2port56,
		abscomate2port57,
		abscomate2port58,
		abscomate2port59,
		abscomate2port60,
		abscomate2port61,
		abscomate2port62,
		abscomate2port63,
		abscomate2port64,
		abscomate2port65,
		abscomate2port66,
		abscomate2port67,
		abscomate2port68,
		abscomate2port69,
		abscomate2port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abscomate3port1  = &AbstractPort{Desc: "absate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port2  = &AbstractPort{Desc: "absate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port3  = &AbstractPort{Desc: "absate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port4  = &AbstractPort{Desc: "absate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port5  = &AbstractPort{Desc: "absate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port6  = &AbstractPort{Desc: "absate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port7  = &AbstractPort{Desc: "absate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port8  = &AbstractPort{Desc: "absate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port9  = &AbstractPort{Desc: "absate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port10 = &AbstractPort{Desc: "absate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port11 = &AbstractPort{Desc: "absate3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port12 = &AbstractPort{Desc: "absate3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port13 = &AbstractPort{Desc: "absate3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port14 = &AbstractPort{Desc: "absate3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port15 = &AbstractPort{Desc: "absate3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port16 = &AbstractPort{Desc: "absate3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port17 = &AbstractPort{Desc: "absate3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port18 = &AbstractPort{Desc: "absate3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port19 = &AbstractPort{Desc: "absate3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port20 = &AbstractPort{Desc: "absate3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port21 = &AbstractPort{Desc: "absate3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port22 = &AbstractPort{Desc: "absate3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port23 = &AbstractPort{Desc: "absate3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port24 = &AbstractPort{Desc: "absate3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port25 = &AbstractPort{Desc: "absate3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port26 = &AbstractPort{Desc: "absate3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port27 = &AbstractPort{Desc: "absate3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port28 = &AbstractPort{Desc: "absate3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port29 = &AbstractPort{Desc: "absate3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port30 = &AbstractPort{Desc: "absate3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port31 = &AbstractPort{Desc: "absate3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port32 = &AbstractPort{Desc: "absate3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port33 = &AbstractPort{Desc: "absate3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port34 = &AbstractPort{Desc: "absate3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port35 = &AbstractPort{Desc: "absate3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port36 = &AbstractPort{Desc: "absate3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port37 = &AbstractPort{Desc: "absate3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port38 = &AbstractPort{Desc: "absate3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port39 = &AbstractPort{Desc: "absate3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port40 = &AbstractPort{Desc: "absate3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port41 = &AbstractPort{Desc: "absate3:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port42 = &AbstractPort{Desc: "absate3:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port43 = &AbstractPort{Desc: "absate3:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port44 = &AbstractPort{Desc: "absate3:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port45 = &AbstractPort{Desc: "absate3:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port46 = &AbstractPort{Desc: "absate3:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port47 = &AbstractPort{Desc: "absate3:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port48 = &AbstractPort{Desc: "absate3:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port49 = &AbstractPort{Desc: "absate3:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port50 = &AbstractPort{Desc: "absate3:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port51 = &AbstractPort{Desc: "absate3:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port52 = &AbstractPort{Desc: "absate3:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port53 = &AbstractPort{Desc: "absate3:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port54 = &AbstractPort{Desc: "absate3:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port55 = &AbstractPort{Desc: "absate3:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port56 = &AbstractPort{Desc: "absate3:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port57 = &AbstractPort{Desc: "absate3:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port58 = &AbstractPort{Desc: "absate3:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port59 = &AbstractPort{Desc: "absate3:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port60 = &AbstractPort{Desc: "absate3:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port61 = &AbstractPort{Desc: "absate3:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port62 = &AbstractPort{Desc: "absate3:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port63 = &AbstractPort{Desc: "absate3:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port64 = &AbstractPort{Desc: "absate3:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port65 = &AbstractPort{Desc: "absate3:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port66 = &AbstractPort{Desc: "absate3:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port67 = &AbstractPort{Desc: "absate3:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port68 = &AbstractPort{Desc: "absate3:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port69 = &AbstractPort{Desc: "absate3:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate3port70 = &AbstractPort{Desc: "absate3:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomate3 = &AbstractNode{Desc: "absate3", Ports: []*AbstractPort{
		abscomate3port1,
		abscomate3port2,
		abscomate3port3,
		abscomate3port4,
		abscomate3port5,
		abscomate3port6,
		abscomate3port7,
		abscomate3port8,
		abscomate3port9,
		abscomate3port10,
		abscomate3port11,
		abscomate3port12,
		abscomate3port13,
		abscomate3port14,
		abscomate3port15,
		abscomate3port16,
		abscomate3port17,
		abscomate3port18,
		abscomate3port19,
		abscomate3port20,
		abscomate3port21,
		abscomate3port22,
		abscomate3port23,
		abscomate3port24,
		abscomate3port25,
		abscomate3port26,
		abscomate3port27,
		abscomate3port28,
		abscomate3port29,
		abscomate3port30,
		abscomate3port31,
		abscomate3port32,
		abscomate3port33,
		abscomate3port34,
		abscomate3port35,
		abscomate3port36,
		abscomate3port37,
		abscomate3port38,
		abscomate3port39,
		abscomate3port40,
		abscomate3port41,
		abscomate3port42,
		abscomate3port43,
		abscomate3port44,
		abscomate3port45,
		abscomate3port46,
		abscomate3port47,
		abscomate3port48,
		abscomate3port49,
		abscomate3port50,
		abscomate3port51,
		abscomate3port52,
		abscomate3port53,
		abscomate3port54,
		abscomate3port55,
		abscomate3port56,
		abscomate3port57,
		abscomate3port58,
		abscomate3port59,
		abscomate3port60,
		abscomate3port61,
		abscomate3port62,
		abscomate3port63,
		abscomate3port64,
		abscomate3port65,
		abscomate3port66,
		abscomate3port67,
		abscomate3port68,
		abscomate3port69,
		abscomate3port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abscomate4port1  = &AbstractPort{Desc: "absate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port2  = &AbstractPort{Desc: "absate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port3  = &AbstractPort{Desc: "absate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port4  = &AbstractPort{Desc: "absate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port5  = &AbstractPort{Desc: "absate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port6  = &AbstractPort{Desc: "absate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port7  = &AbstractPort{Desc: "absate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port8  = &AbstractPort{Desc: "absate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port9  = &AbstractPort{Desc: "absate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port10 = &AbstractPort{Desc: "absate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port11 = &AbstractPort{Desc: "absate4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port12 = &AbstractPort{Desc: "absate4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port13 = &AbstractPort{Desc: "absate4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port14 = &AbstractPort{Desc: "absate4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port15 = &AbstractPort{Desc: "absate4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port16 = &AbstractPort{Desc: "absate4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port17 = &AbstractPort{Desc: "absate4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port18 = &AbstractPort{Desc: "absate4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port19 = &AbstractPort{Desc: "absate4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port20 = &AbstractPort{Desc: "absate4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port21 = &AbstractPort{Desc: "absate4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port22 = &AbstractPort{Desc: "absate4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port23 = &AbstractPort{Desc: "absate4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port24 = &AbstractPort{Desc: "absate4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port25 = &AbstractPort{Desc: "absate4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port26 = &AbstractPort{Desc: "absate4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port27 = &AbstractPort{Desc: "absate4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port28 = &AbstractPort{Desc: "absate4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port29 = &AbstractPort{Desc: "absate4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port30 = &AbstractPort{Desc: "absate4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port31 = &AbstractPort{Desc: "absate4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port32 = &AbstractPort{Desc: "absate4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port33 = &AbstractPort{Desc: "absate4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port34 = &AbstractPort{Desc: "absate4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port35 = &AbstractPort{Desc: "absate4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port36 = &AbstractPort{Desc: "absate4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port37 = &AbstractPort{Desc: "absate4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port38 = &AbstractPort{Desc: "absate4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port39 = &AbstractPort{Desc: "absate4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port40 = &AbstractPort{Desc: "absate4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port41 = &AbstractPort{Desc: "absate4:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port42 = &AbstractPort{Desc: "absate4:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port43 = &AbstractPort{Desc: "absate4:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port44 = &AbstractPort{Desc: "absate4:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port45 = &AbstractPort{Desc: "absate4:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port46 = &AbstractPort{Desc: "absate4:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port47 = &AbstractPort{Desc: "absate4:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port48 = &AbstractPort{Desc: "absate4:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port49 = &AbstractPort{Desc: "absate4:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port50 = &AbstractPort{Desc: "absate4:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port51 = &AbstractPort{Desc: "absate4:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port52 = &AbstractPort{Desc: "absate4:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port53 = &AbstractPort{Desc: "absate4:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port54 = &AbstractPort{Desc: "absate4:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port55 = &AbstractPort{Desc: "absate4:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port56 = &AbstractPort{Desc: "absate4:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port57 = &AbstractPort{Desc: "absate4:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port58 = &AbstractPort{Desc: "absate4:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port59 = &AbstractPort{Desc: "absate4:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port60 = &AbstractPort{Desc: "absate4:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port61 = &AbstractPort{Desc: "absate4:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port62 = &AbstractPort{Desc: "absate4:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port63 = &AbstractPort{Desc: "absate4:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port64 = &AbstractPort{Desc: "absate4:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port65 = &AbstractPort{Desc: "absate4:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port66 = &AbstractPort{Desc: "absate4:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port67 = &AbstractPort{Desc: "absate4:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port68 = &AbstractPort{Desc: "absate4:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port69 = &AbstractPort{Desc: "absate4:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate4port70 = &AbstractPort{Desc: "absate4:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomate4 = &AbstractNode{Desc: "absate4", Ports: []*AbstractPort{
		abscomate4port1,
		abscomate4port2,
		abscomate4port3,
		abscomate4port4,
		abscomate4port5,
		abscomate4port6,
		abscomate4port7,
		abscomate4port8,
		abscomate4port9,
		abscomate4port10,
		abscomate4port11,
		abscomate4port12,
		abscomate4port13,
		abscomate4port14,
		abscomate4port15,
		abscomate4port16,
		abscomate4port17,
		abscomate4port18,
		abscomate4port19,
		abscomate4port20,
		abscomate4port21,
		abscomate4port22,
		abscomate4port23,
		abscomate4port24,
		abscomate4port25,
		abscomate4port26,
		abscomate4port27,
		abscomate4port28,
		abscomate4port29,
		abscomate4port30,
		abscomate4port31,
		abscomate4port32,
		abscomate4port33,
		abscomate4port34,
		abscomate4port35,
		abscomate4port36,
		abscomate4port37,
		abscomate4port38,
		abscomate4port39,
		abscomate4port40,
		abscomate4port41,
		abscomate4port42,
		abscomate4port43,
		abscomate4port44,
		abscomate4port45,
		abscomate4port46,
		abscomate4port47,
		abscomate4port48,
		abscomate4port49,
		abscomate4port50,
		abscomate4port51,
		abscomate4port52,
		abscomate4port53,
		abscomate4port54,
		abscomate4port55,
		abscomate4port56,
		abscomate4port57,
		abscomate4port58,
		abscomate4port59,
		abscomate4port60,
		abscomate4port61,
		abscomate4port62,
		abscomate4port63,
		abscomate4port64,
		abscomate4port65,
		abscomate4port66,
		abscomate4port67,
		abscomate4port68,
		abscomate4port69,
		abscomate4port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abscomate5port1  = &AbstractPort{Desc: "absate5:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port2  = &AbstractPort{Desc: "absate5:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port3  = &AbstractPort{Desc: "absate5:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port4  = &AbstractPort{Desc: "absate5:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port5  = &AbstractPort{Desc: "absate5:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port6  = &AbstractPort{Desc: "absate5:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port7  = &AbstractPort{Desc: "absate5:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port8  = &AbstractPort{Desc: "absate5:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port9  = &AbstractPort{Desc: "absate5:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port10 = &AbstractPort{Desc: "absate5:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port11 = &AbstractPort{Desc: "absate5:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port12 = &AbstractPort{Desc: "absate5:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port13 = &AbstractPort{Desc: "absate5:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port14 = &AbstractPort{Desc: "absate5:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port15 = &AbstractPort{Desc: "absate5:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port16 = &AbstractPort{Desc: "absate5:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port17 = &AbstractPort{Desc: "absate5:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port18 = &AbstractPort{Desc: "absate5:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port19 = &AbstractPort{Desc: "absate5:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port20 = &AbstractPort{Desc: "absate5:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port21 = &AbstractPort{Desc: "absate5:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port22 = &AbstractPort{Desc: "absate5:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port23 = &AbstractPort{Desc: "absate5:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port24 = &AbstractPort{Desc: "absate5:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port25 = &AbstractPort{Desc: "absate5:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port26 = &AbstractPort{Desc: "absate5:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port27 = &AbstractPort{Desc: "absate5:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port28 = &AbstractPort{Desc: "absate5:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port29 = &AbstractPort{Desc: "absate5:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port30 = &AbstractPort{Desc: "absate5:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port31 = &AbstractPort{Desc: "absate5:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port32 = &AbstractPort{Desc: "absate5:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port33 = &AbstractPort{Desc: "absate5:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port34 = &AbstractPort{Desc: "absate5:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port35 = &AbstractPort{Desc: "absate5:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port36 = &AbstractPort{Desc: "absate5:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port37 = &AbstractPort{Desc: "absate5:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port38 = &AbstractPort{Desc: "absate5:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port39 = &AbstractPort{Desc: "absate5:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port40 = &AbstractPort{Desc: "absate5:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port41 = &AbstractPort{Desc: "absate5:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port42 = &AbstractPort{Desc: "absate5:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port43 = &AbstractPort{Desc: "absate5:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port44 = &AbstractPort{Desc: "absate5:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port45 = &AbstractPort{Desc: "absate5:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port46 = &AbstractPort{Desc: "absate5:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port47 = &AbstractPort{Desc: "absate5:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port48 = &AbstractPort{Desc: "absate5:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port49 = &AbstractPort{Desc: "absate5:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port50 = &AbstractPort{Desc: "absate5:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port51 = &AbstractPort{Desc: "absate5:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port52 = &AbstractPort{Desc: "absate5:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port53 = &AbstractPort{Desc: "absate5:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port54 = &AbstractPort{Desc: "absate5:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port55 = &AbstractPort{Desc: "absate5:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port56 = &AbstractPort{Desc: "absate5:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port57 = &AbstractPort{Desc: "absate5:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port58 = &AbstractPort{Desc: "absate5:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port59 = &AbstractPort{Desc: "absate5:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port60 = &AbstractPort{Desc: "absate5:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port61 = &AbstractPort{Desc: "absate5:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port62 = &AbstractPort{Desc: "absate5:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port63 = &AbstractPort{Desc: "absate5:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port64 = &AbstractPort{Desc: "absate5:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port65 = &AbstractPort{Desc: "absate5:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port66 = &AbstractPort{Desc: "absate5:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port67 = &AbstractPort{Desc: "absate5:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port68 = &AbstractPort{Desc: "absate5:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port69 = &AbstractPort{Desc: "absate5:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate5port70 = &AbstractPort{Desc: "absate5:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomate5 = &AbstractNode{Desc: "absate5", Ports: []*AbstractPort{
		abscomate5port1,
		abscomate5port2,
		abscomate5port3,
		abscomate5port4,
		abscomate5port5,
		abscomate5port6,
		abscomate5port7,
		abscomate5port8,
		abscomate5port9,
		abscomate5port10,
		abscomate5port11,
		abscomate5port12,
		abscomate5port13,
		abscomate5port14,
		abscomate5port15,
		abscomate5port16,
		abscomate5port17,
		abscomate5port18,
		abscomate5port19,
		abscomate5port20,
		abscomate5port21,
		abscomate5port22,
		abscomate5port23,
		abscomate5port24,
		abscomate5port25,
		abscomate5port26,
		abscomate5port27,
		abscomate5port28,
		abscomate5port29,
		abscomate5port30,
		abscomate5port31,
		abscomate5port32,
		abscomate5port33,
		abscomate5port34,
		abscomate5port35,
		abscomate5port36,
		abscomate5port37,
		abscomate5port38,
		abscomate5port39,
		abscomate5port40,
		abscomate5port41,
		abscomate5port42,
		abscomate5port43,
		abscomate5port44,
		abscomate5port45,
		abscomate5port46,
		abscomate5port47,
		abscomate5port48,
		abscomate5port49,
		abscomate5port50,
		abscomate5port51,
		abscomate5port52,
		abscomate5port53,
		abscomate5port54,
		abscomate5port55,
		abscomate5port56,
		abscomate5port57,
		abscomate5port58,
		abscomate5port59,
		abscomate5port60,
		abscomate5port61,
		abscomate5port62,
		abscomate5port63,
		abscomate5port64,
		abscomate5port65,
		abscomate5port66,
		abscomate5port67,
		abscomate5port68,
		abscomate5port69,
		abscomate5port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	abscomate6port1  = &AbstractPort{Desc: "absate6:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port2  = &AbstractPort{Desc: "absate6:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port3  = &AbstractPort{Desc: "absate6:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port4  = &AbstractPort{Desc: "absate6:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port5  = &AbstractPort{Desc: "absate6:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port6  = &AbstractPort{Desc: "absate6:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port7  = &AbstractPort{Desc: "absate6:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port8  = &AbstractPort{Desc: "absate6:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port9  = &AbstractPort{Desc: "absate6:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port10 = &AbstractPort{Desc: "absate6:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port11 = &AbstractPort{Desc: "absate6:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port12 = &AbstractPort{Desc: "absate6:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port13 = &AbstractPort{Desc: "absate6:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port14 = &AbstractPort{Desc: "absate6:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port15 = &AbstractPort{Desc: "absate6:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port16 = &AbstractPort{Desc: "absate6:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port17 = &AbstractPort{Desc: "absate6:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port18 = &AbstractPort{Desc: "absate6:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port19 = &AbstractPort{Desc: "absate6:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port20 = &AbstractPort{Desc: "absate6:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port21 = &AbstractPort{Desc: "absate6:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port22 = &AbstractPort{Desc: "absate6:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port23 = &AbstractPort{Desc: "absate6:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port24 = &AbstractPort{Desc: "absate6:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port25 = &AbstractPort{Desc: "absate6:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port26 = &AbstractPort{Desc: "absate6:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port27 = &AbstractPort{Desc: "absate6:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port28 = &AbstractPort{Desc: "absate6:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port29 = &AbstractPort{Desc: "absate6:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port30 = &AbstractPort{Desc: "absate6:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port31 = &AbstractPort{Desc: "absate6:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port32 = &AbstractPort{Desc: "absate6:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port33 = &AbstractPort{Desc: "absate6:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port34 = &AbstractPort{Desc: "absate6:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port35 = &AbstractPort{Desc: "absate6:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port36 = &AbstractPort{Desc: "absate6:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port37 = &AbstractPort{Desc: "absate6:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port38 = &AbstractPort{Desc: "absate6:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port39 = &AbstractPort{Desc: "absate6:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port40 = &AbstractPort{Desc: "absate6:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port41 = &AbstractPort{Desc: "absate6:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port42 = &AbstractPort{Desc: "absate6:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port43 = &AbstractPort{Desc: "absate6:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port44 = &AbstractPort{Desc: "absate6:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port45 = &AbstractPort{Desc: "absate6:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port46 = &AbstractPort{Desc: "absate6:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port47 = &AbstractPort{Desc: "absate6:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port48 = &AbstractPort{Desc: "absate6:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port49 = &AbstractPort{Desc: "absate6:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port50 = &AbstractPort{Desc: "absate6:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port51 = &AbstractPort{Desc: "absate6:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port52 = &AbstractPort{Desc: "absate6:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port53 = &AbstractPort{Desc: "absate6:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port54 = &AbstractPort{Desc: "absate6:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port55 = &AbstractPort{Desc: "absate6:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port56 = &AbstractPort{Desc: "absate6:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port57 = &AbstractPort{Desc: "absate6:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port58 = &AbstractPort{Desc: "absate6:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port59 = &AbstractPort{Desc: "absate6:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port60 = &AbstractPort{Desc: "absate6:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port61 = &AbstractPort{Desc: "absate6:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port62 = &AbstractPort{Desc: "absate6:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port63 = &AbstractPort{Desc: "absate6:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port64 = &AbstractPort{Desc: "absate6:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port65 = &AbstractPort{Desc: "absate6:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port66 = &AbstractPort{Desc: "absate6:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port67 = &AbstractPort{Desc: "absate6:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port68 = &AbstractPort{Desc: "absate6:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port69 = &AbstractPort{Desc: "absate6:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	abscomate6port70 = &AbstractPort{Desc: "absate6:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	abscomate6 = &AbstractNode{Desc: "absate6", Ports: []*AbstractPort{
		abscomate6port1,
		abscomate6port2,
		abscomate6port3,
		abscomate6port4,
		abscomate6port5,
		abscomate6port6,
		abscomate6port7,
		abscomate6port8,
		abscomate6port9,
		abscomate6port10,
		abscomate6port11,
		abscomate6port12,
		abscomate6port13,
		abscomate6port14,
		abscomate6port15,
		abscomate6port16,
		abscomate6port17,
		abscomate6port18,
		abscomate6port19,
		abscomate6port20,
		abscomate6port21,
		abscomate6port22,
		abscomate6port23,
		abscomate6port24,
		abscomate6port25,
		abscomate6port26,
		abscomate6port27,
		abscomate6port28,
		abscomate6port29,
		abscomate6port30,
		abscomate6port31,
		abscomate6port32,
		abscomate6port33,
		abscomate6port34,
		abscomate6port35,
		abscomate6port36,
		abscomate6port37,
		abscomate6port38,
		abscomate6port39,
		abscomate6port40,
		abscomate6port41,
		abscomate6port42,
		abscomate6port43,
		abscomate6port44,
		abscomate6port45,
		abscomate6port46,
		abscomate6port47,
		abscomate6port48,
		abscomate6port49,
		abscomate6port50,
		abscomate6port51,
		abscomate6port52,
		abscomate6port53,
		abscomate6port54,
		abscomate6port55,
		abscomate6port56,
		abscomate6port57,
		abscomate6port58,
		abscomate6port59,
		abscomate6port60,
		abscomate6port61,
		abscomate6port62,
		abscomate6port63,
		abscomate6port64,
		abscomate6port65,
		abscomate6port66,
		abscomate6port67,
		abscomate6port68,
		abscomate6port69,
		abscomate6port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
)

var abstractGraphcom = &AbstractGraph{
	Nodes: []*AbstractNode{abscomdut1,
		abscomate1,
		abscomdut2,
		abscomate2,
		abscomdut3,
		abscomate3,
		abscomdut4,
		abscomate4, abscomdut5,
		abscomate5,
		abscomdut6,
		abscomate6},
	Edges: []*AbstractEdge{
		{abscomdut1port1, abscomate1port1},
		{abscomdut1port2, abscomate1port2},
		{abscomdut1port3, abscomate1port3},
		{abscomdut1port4, abscomate1port4},
		{abscomdut1port5, abscomate1port5},
		{abscomdut1port6, abscomate1port6},
		{abscomdut1port7, abscomate1port7},
		{abscomdut1port8, abscomate1port8},
		{abscomdut1port9, abscomate1port9},
		{abscomdut1port10, abscomate1port10},
		{abscomdut1port11, abscomate1port11},
		{abscomdut1port12, abscomate1port12},
		{abscomdut1port13, abscomate1port13},
		{abscomdut1port14, abscomate1port14},
		{abscomdut1port15, abscomate1port15},
		{abscomdut1port16, abscomate1port16},
		{abscomdut1port17, abscomate1port17},
		{abscomdut1port18, abscomate1port18},
		{abscomdut1port19, abscomate1port19},
		{abscomdut1port20, abscomate1port20},
		{abscomdut1port21, abscomate1port21},
		{abscomdut1port22, abscomate1port22},
		{abscomdut1port23, abscomate1port23},
		{abscomdut1port24, abscomate1port24},
		{abscomdut1port25, abscomate1port25},
		{abscomdut1port26, abscomate1port26},
		{abscomdut1port27, abscomate1port27},
		{abscomdut1port28, abscomate1port28},
		{abscomdut1port29, abscomate1port29},
		{abscomdut1port30, abscomate1port30},
		{abscomdut1port31, abscomate1port31},
		{abscomdut1port32, abscomate1port32},
		{abscomdut1port33, abscomate1port33},
		{abscomdut1port34, abscomate1port34},
		{abscomdut1port35, abscomate1port35},
		{abscomdut1port36, abscomate1port36},
		{abscomdut1port37, abscomate1port37},
		{abscomdut1port38, abscomate1port38},
		{abscomdut1port39, abscomate1port39},
		{abscomdut1port40, abscomate1port40},
		{abscomdut1port41, abscomate1port41},
		{abscomdut1port42, abscomate1port42},
		{abscomdut1port43, abscomate1port43},
		{abscomdut1port44, abscomate1port44},
		{abscomdut1port45, abscomate1port45},
		{abscomdut1port46, abscomate1port46},
		{abscomdut1port47, abscomate1port47},
		{abscomdut1port48, abscomate1port48},
		{abscomdut1port49, abscomate1port49},
		{abscomdut1port50, abscomate1port50},
		{abscomdut1port51, abscomate1port51},
		{abscomdut1port52, abscomate1port52},
		{abscomdut1port53, abscomate1port53},
		{abscomdut1port54, abscomate1port54},
		{abscomdut1port55, abscomate1port55},
		{abscomdut1port56, abscomate1port56},
		{abscomdut1port57, abscomate1port57},
		{abscomdut1port58, abscomate1port58},
		{abscomdut1port59, abscomate1port59},
		{abscomdut1port60, abscomate1port60},
		{abscomdut1port61, abscomate1port61},
		{abscomdut1port62, abscomate1port62},
		{abscomdut1port63, abscomate1port63},
		{abscomdut1port64, abscomate1port64},
		{abscomdut1port65, abscomate1port65},
		{abscomdut1port66, abscomate1port66},
		{abscomdut1port67, abscomate1port67},
		{abscomdut1port68, abscomate1port68},
		{abscomdut1port69, abscomate1port69},
		{abscomdut1port70, abscomate1port70},

		{abscomdut2port1, abscomate2port1},
		{abscomdut2port2, abscomate2port2},
		{abscomdut2port3, abscomate2port3},
		{abscomdut2port4, abscomate2port4},
		{abscomdut2port5, abscomate2port5},
		{abscomdut2port6, abscomate2port6},
		{abscomdut2port7, abscomate2port7},
		{abscomdut2port8, abscomate2port8},
		{abscomdut2port9, abscomate2port9},
		{abscomdut2port10, abscomate2port10},
		{abscomdut2port11, abscomate2port11},
		{abscomdut2port12, abscomate2port12},
		{abscomdut2port13, abscomate2port13},
		{abscomdut2port14, abscomate2port14},
		{abscomdut2port15, abscomate2port15},
		{abscomdut2port16, abscomate2port16},
		{abscomdut2port17, abscomate2port17},
		{abscomdut2port18, abscomate2port18},
		{abscomdut2port19, abscomate2port19},
		{abscomdut2port20, abscomate2port20},
		{abscomdut2port21, abscomate2port21},
		{abscomdut2port22, abscomate2port22},
		{abscomdut2port23, abscomate2port23},
		{abscomdut2port24, abscomate2port24},
		{abscomdut2port25, abscomate2port25},
		{abscomdut2port26, abscomate2port26},
		{abscomdut2port27, abscomate2port27},
		{abscomdut2port28, abscomate2port28},
		{abscomdut2port29, abscomate2port29},
		{abscomdut2port30, abscomate2port30},
		{abscomdut2port31, abscomate2port31},
		{abscomdut2port32, abscomate2port32},
		{abscomdut2port33, abscomate2port33},
		{abscomdut2port34, abscomate2port34},
		{abscomdut2port35, abscomate2port35},
		{abscomdut2port36, abscomate2port36},
		{abscomdut2port37, abscomate2port37},
		{abscomdut2port38, abscomate2port38},
		{abscomdut2port39, abscomate2port39},
		{abscomdut2port40, abscomate2port40},
		{abscomdut2port41, abscomate2port41},
		{abscomdut2port42, abscomate2port42},
		{abscomdut2port43, abscomate2port43},
		{abscomdut2port44, abscomate2port44},
		{abscomdut2port45, abscomate2port45},
		{abscomdut2port46, abscomate2port46},
		{abscomdut2port47, abscomate2port47},
		{abscomdut2port48, abscomate2port48},
		{abscomdut2port49, abscomate2port49},
		{abscomdut2port50, abscomate2port50},
		{abscomdut2port51, abscomate2port51},
		{abscomdut2port52, abscomate2port52},
		{abscomdut2port53, abscomate2port53},
		{abscomdut2port54, abscomate2port54},
		{abscomdut2port55, abscomate2port55},
		{abscomdut2port56, abscomate2port56},
		{abscomdut2port57, abscomate2port57},
		{abscomdut2port58, abscomate2port58},
		{abscomdut2port59, abscomate2port59},
		{abscomdut2port60, abscomate2port60},
		{abscomdut2port61, abscomate2port61},
		{abscomdut2port62, abscomate2port62},
		{abscomdut2port63, abscomate2port63},
		{abscomdut2port64, abscomate2port64},
		{abscomdut2port65, abscomate2port65},
		{abscomdut2port66, abscomate2port66},
		{abscomdut2port67, abscomate2port67},
		{abscomdut2port68, abscomate2port68},
		{abscomdut2port69, abscomate2port69},
		{abscomdut2port70, abscomate2port70},

		{abscomdut3port1, abscomate3port1},
		{abscomdut3port2, abscomate3port2},
		{abscomdut3port3, abscomate3port3},
		{abscomdut3port4, abscomate3port4},
		{abscomdut3port5, abscomate3port5},
		{abscomdut3port6, abscomate3port6},
		{abscomdut3port7, abscomate3port7},
		{abscomdut3port8, abscomate3port8},
		{abscomdut3port9, abscomate3port9},
		{abscomdut3port10, abscomate3port10},
		{abscomdut3port11, abscomate3port11},
		{abscomdut3port12, abscomate3port12},
		{abscomdut3port13, abscomate3port13},
		{abscomdut3port14, abscomate3port14},
		{abscomdut3port15, abscomate3port15},
		{abscomdut3port16, abscomate3port16},
		{abscomdut3port17, abscomate3port17},
		{abscomdut3port18, abscomate3port18},
		{abscomdut3port19, abscomate3port19},
		{abscomdut3port20, abscomate3port20},
		{abscomdut3port21, abscomate3port21},
		{abscomdut3port22, abscomate3port22},
		{abscomdut3port23, abscomate3port23},
		{abscomdut3port24, abscomate3port24},
		{abscomdut3port25, abscomate3port25},
		{abscomdut3port26, abscomate3port26},
		{abscomdut3port27, abscomate3port27},
		{abscomdut3port28, abscomate3port28},
		{abscomdut3port29, abscomate3port29},
		{abscomdut3port30, abscomate3port30},
		{abscomdut3port31, abscomate3port31},
		{abscomdut3port32, abscomate3port32},
		{abscomdut3port33, abscomate3port33},
		{abscomdut3port34, abscomate3port34},
		{abscomdut3port35, abscomate3port35},
		{abscomdut3port36, abscomate3port36},
		{abscomdut3port37, abscomate3port37},
		{abscomdut3port38, abscomate3port38},
		{abscomdut3port39, abscomate3port39},
		{abscomdut3port40, abscomate3port40},
		{abscomdut3port41, abscomate3port41},
		{abscomdut3port42, abscomate3port42},
		{abscomdut3port43, abscomate3port43},
		{abscomdut3port44, abscomate3port44},
		{abscomdut3port45, abscomate3port45},
		{abscomdut3port46, abscomate3port46},
		{abscomdut3port47, abscomate3port47},
		{abscomdut3port48, abscomate3port48},
		{abscomdut3port49, abscomate3port49},
		{abscomdut3port50, abscomate3port50},
		{abscomdut3port51, abscomate3port51},
		{abscomdut3port52, abscomate3port52},
		{abscomdut3port53, abscomate3port53},
		{abscomdut3port54, abscomate3port54},
		{abscomdut3port55, abscomate3port55},
		{abscomdut3port56, abscomate3port56},
		{abscomdut3port57, abscomate3port57},
		{abscomdut3port58, abscomate3port58},
		{abscomdut3port59, abscomate3port59},
		{abscomdut3port60, abscomate3port60},
		{abscomdut3port61, abscomate3port61},
		{abscomdut3port62, abscomate3port62},
		{abscomdut3port63, abscomate3port63},
		{abscomdut3port64, abscomate3port64},
		{abscomdut3port65, abscomate3port65},
		{abscomdut3port66, abscomate3port66},
		{abscomdut3port67, abscomate3port67},
		{abscomdut3port68, abscomate3port68},
		{abscomdut3port69, abscomate3port69},
		{abscomdut3port70, abscomate3port70},

		{abscomdut4port1, abscomate4port1},
		{abscomdut4port2, abscomate4port2},
		{abscomdut4port3, abscomate4port3},
		{abscomdut4port4, abscomate4port4},
		{abscomdut4port5, abscomate4port5},
		{abscomdut4port6, abscomate4port6},
		{abscomdut4port7, abscomate4port7},
		{abscomdut4port8, abscomate4port8},
		{abscomdut4port9, abscomate4port9},
		{abscomdut4port10, abscomate4port10},
		{abscomdut4port11, abscomate4port11},
		{abscomdut4port12, abscomate4port12},
		{abscomdut4port13, abscomate4port13},
		{abscomdut4port14, abscomate4port14},
		{abscomdut4port15, abscomate4port15},
		{abscomdut4port16, abscomate4port16},
		{abscomdut4port17, abscomate4port17},
		{abscomdut4port18, abscomate4port18},
		{abscomdut4port19, abscomate4port19},
		{abscomdut4port20, abscomate4port20},
		{abscomdut4port21, abscomate4port21},
		{abscomdut4port22, abscomate4port22},
		{abscomdut4port23, abscomate4port23},
		{abscomdut4port24, abscomate4port24},
		{abscomdut4port25, abscomate4port25},
		{abscomdut4port26, abscomate4port26},
		{abscomdut4port27, abscomate4port27},
		{abscomdut4port28, abscomate4port28},
		{abscomdut4port29, abscomate4port29},
		{abscomdut4port30, abscomate4port30},
		{abscomdut4port31, abscomate4port31},
		{abscomdut4port32, abscomate4port32},
		{abscomdut4port33, abscomate4port33},
		{abscomdut4port34, abscomate4port34},
		{abscomdut4port35, abscomate4port35},
		{abscomdut4port36, abscomate4port36},
		{abscomdut4port37, abscomate4port37},
		{abscomdut4port38, abscomate4port38},
		{abscomdut4port39, abscomate4port39},
		{abscomdut4port40, abscomate4port40},
		{abscomdut4port41, abscomate4port41},
		{abscomdut4port42, abscomate4port42},
		{abscomdut4port43, abscomate4port43},
		{abscomdut4port44, abscomate4port44},
		{abscomdut4port45, abscomate4port45},
		{abscomdut4port46, abscomate4port46},
		{abscomdut4port47, abscomate4port47},
		{abscomdut4port48, abscomate4port48},
		{abscomdut4port49, abscomate4port49},
		{abscomdut4port50, abscomate4port50},
		{abscomdut4port51, abscomate4port51},
		{abscomdut4port52, abscomate4port52},
		{abscomdut4port53, abscomate4port53},
		{abscomdut4port54, abscomate4port54},
		{abscomdut4port55, abscomate4port55},
		{abscomdut4port56, abscomate4port56},
		{abscomdut4port57, abscomate4port57},
		{abscomdut4port58, abscomate4port58},
		{abscomdut4port59, abscomate4port59},
		{abscomdut4port60, abscomate4port60},
		{abscomdut4port61, abscomate4port61},
		{abscomdut4port62, abscomate4port62},
		{abscomdut4port63, abscomate4port63},
		{abscomdut4port64, abscomate4port64},
		{abscomdut4port65, abscomate4port65},
		{abscomdut4port66, abscomate4port66},
		{abscomdut4port67, abscomate4port67},
		{abscomdut4port68, abscomate4port68},
		{abscomdut4port69, abscomate4port69},
		{abscomdut4port70, abscomate4port70},

		{abscomdut5port1, abscomate5port1},
		{abscomdut5port2, abscomate5port2},
		{abscomdut5port3, abscomate5port3},
		{abscomdut5port4, abscomate5port4},
		{abscomdut5port5, abscomate5port5},
		{abscomdut5port6, abscomate5port6},
		{abscomdut5port7, abscomate5port7},
		{abscomdut5port8, abscomate5port8},
		{abscomdut5port9, abscomate5port9},
		{abscomdut5port10, abscomate5port10},
		{abscomdut5port11, abscomate5port11},
		{abscomdut5port12, abscomate5port12},
		{abscomdut5port13, abscomate5port13},
		{abscomdut5port14, abscomate5port14},
		{abscomdut5port15, abscomate5port15},
		{abscomdut5port16, abscomate5port16},
		{abscomdut5port17, abscomate5port17},
		{abscomdut5port18, abscomate5port18},
		{abscomdut5port19, abscomate5port19},
		{abscomdut5port20, abscomate5port20},
		{abscomdut5port21, abscomate5port21},
		{abscomdut5port22, abscomate5port22},
		{abscomdut5port23, abscomate5port23},
		{abscomdut5port24, abscomate5port24},
		{abscomdut5port25, abscomate5port25},
		{abscomdut5port26, abscomate5port26},
		{abscomdut5port27, abscomate5port27},
		{abscomdut5port28, abscomate5port28},
		{abscomdut5port29, abscomate5port29},
		{abscomdut5port30, abscomate5port30},
		{abscomdut5port31, abscomate5port31},
		{abscomdut5port32, abscomate5port32},
		{abscomdut5port33, abscomate5port33},
		{abscomdut5port34, abscomate5port34},
		{abscomdut5port35, abscomate5port35},
		{abscomdut5port36, abscomate5port36},
		{abscomdut5port37, abscomate5port37},
		{abscomdut5port38, abscomate5port38},
		{abscomdut5port39, abscomate5port39},
		{abscomdut5port40, abscomate5port40},
		{abscomdut5port41, abscomate5port41},
		{abscomdut5port42, abscomate5port42},
		{abscomdut5port43, abscomate5port43},
		{abscomdut5port44, abscomate5port44},
		{abscomdut5port45, abscomate5port45},
		{abscomdut5port46, abscomate5port46},
		{abscomdut5port47, abscomate5port47},
		{abscomdut5port48, abscomate5port48},
		{abscomdut5port49, abscomate5port49},
		{abscomdut5port50, abscomate5port50},
		{abscomdut5port51, abscomate5port51},
		{abscomdut5port52, abscomate5port52},
		{abscomdut5port53, abscomate5port53},
		{abscomdut5port54, abscomate5port54},
		{abscomdut5port55, abscomate5port55},
		{abscomdut5port56, abscomate5port56},
		{abscomdut5port57, abscomate5port57},
		{abscomdut5port58, abscomate5port58},
		{abscomdut5port59, abscomate5port59},
		{abscomdut5port60, abscomate5port60},
		{abscomdut5port61, abscomate5port61},
		{abscomdut5port62, abscomate5port62},
		{abscomdut5port63, abscomate5port63},
		{abscomdut5port64, abscomate5port64},
		{abscomdut5port65, abscomate5port65},
		{abscomdut5port66, abscomate5port66},
		{abscomdut5port67, abscomate5port67},
		{abscomdut5port68, abscomate5port68},
		{abscomdut5port69, abscomate5port69},
		{abscomdut5port70, abscomate5port70},

		{abscomdut6port1, abscomate6port1},
		{abscomdut6port2, abscomate6port2},
		{abscomdut6port3, abscomate6port3},
		{abscomdut6port4, abscomate6port4},
		{abscomdut6port5, abscomate6port5},
		{abscomdut6port6, abscomate6port6},
		{abscomdut6port7, abscomate6port7},
		{abscomdut6port8, abscomate6port8},
		{abscomdut6port9, abscomate6port9},
		{abscomdut6port10, abscomate6port10},
		{abscomdut6port11, abscomate6port11},
		{abscomdut6port12, abscomate6port12},
		{abscomdut6port13, abscomate6port13},
		{abscomdut6port14, abscomate6port14},
		{abscomdut6port15, abscomate6port15},
		{abscomdut6port16, abscomate6port16},
		{abscomdut6port17, abscomate6port17},
		{abscomdut6port18, abscomate6port18},
		{abscomdut6port19, abscomate6port19},
		{abscomdut6port20, abscomate6port20},
		{abscomdut6port21, abscomate6port21},
		{abscomdut6port22, abscomate6port22},
		{abscomdut6port23, abscomate6port23},
		{abscomdut6port24, abscomate6port24},
		{abscomdut6port25, abscomate6port25},
		{abscomdut6port26, abscomate6port26},
		{abscomdut6port27, abscomate6port27},
		{abscomdut6port28, abscomate6port28},
		{abscomdut6port29, abscomate6port29},
		{abscomdut6port30, abscomate6port30},
		{abscomdut6port31, abscomate6port31},
		{abscomdut6port32, abscomate6port32},
		{abscomdut6port33, abscomate6port33},
		{abscomdut6port34, abscomate6port34},
		{abscomdut6port35, abscomate6port35},
		{abscomdut6port36, abscomate6port36},
		{abscomdut6port37, abscomate6port37},
		{abscomdut6port38, abscomate6port38},
		{abscomdut6port39, abscomate6port39},
		{abscomdut6port40, abscomate6port40},
		{abscomdut6port41, abscomate6port41},
		{abscomdut6port42, abscomate6port42},
		{abscomdut6port43, abscomate6port43},
		{abscomdut6port44, abscomate6port44},
		{abscomdut6port45, abscomate6port45},
		{abscomdut6port46, abscomate6port46},
		{abscomdut6port47, abscomate6port47},
		{abscomdut6port48, abscomate6port48},
		{abscomdut6port49, abscomate6port49},
		{abscomdut6port50, abscomate6port50},
		{abscomdut6port51, abscomate6port51},
		{abscomdut6port52, abscomate6port52},
		{abscomdut6port53, abscomate6port53},
		{abscomdut6port54, abscomate6port54},
		{abscomdut6port55, abscomate6port55},
		{abscomdut6port56, abscomate6port56},
		{abscomdut6port57, abscomate6port57},
		{abscomdut6port58, abscomate6port58},
		{abscomdut6port59, abscomate6port59},
		{abscomdut6port60, abscomate6port60},
		{abscomdut6port61, abscomate6port61},
		{abscomdut6port62, abscomate6port62},
		{abscomdut6port63, abscomate6port63},
		{abscomdut6port64, abscomate6port64},
		{abscomdut6port65, abscomate6port65},
		{abscomdut6port66, abscomate6port66},
		{abscomdut6port67, abscomate6port67},
		{abscomdut6port68, abscomate6port68},
		{abscomdut6port69, abscomate6port69},
		{abscomdut6port70, abscomate6port70}},
}

func TestSolveTestHybridcom(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraphcom, superGraphTestcom)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 12 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 840 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
