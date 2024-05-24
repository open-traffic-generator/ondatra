// // 4 duts, 4 ates, 70 ports each and 3 switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	diffdut1port1  = &ConcretePort{Desc: "diffdut1:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port2  = &ConcretePort{Desc: "diffdut1:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port3  = &ConcretePort{Desc: "diffdut1:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port4  = &ConcretePort{Desc: "diffdut1:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port5  = &ConcretePort{Desc: "diffdut1:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port6  = &ConcretePort{Desc: "diffdut1:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port7  = &ConcretePort{Desc: "diffdut1:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port8  = &ConcretePort{Desc: "diffdut1:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port9  = &ConcretePort{Desc: "diffdut1:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port10 = &ConcretePort{Desc: "diffdut1:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut1port11 = &ConcretePort{Desc: "diffdut1:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port12 = &ConcretePort{Desc: "diffdut1:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port13 = &ConcretePort{Desc: "diffdut1:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port14 = &ConcretePort{Desc: "diffdut1:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port15 = &ConcretePort{Desc: "diffdut1:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port16 = &ConcretePort{Desc: "diffdut1:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port17 = &ConcretePort{Desc: "diffdut1:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port18 = &ConcretePort{Desc: "diffdut1:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port19 = &ConcretePort{Desc: "diffdut1:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port20 = &ConcretePort{Desc: "diffdut1:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut1port21 = &ConcretePort{Desc: "diffdut1:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port22 = &ConcretePort{Desc: "diffdut1:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port23 = &ConcretePort{Desc: "diffdut1:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port24 = &ConcretePort{Desc: "diffdut1:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port25 = &ConcretePort{Desc: "diffdut1:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port26 = &ConcretePort{Desc: "diffdut1:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port27 = &ConcretePort{Desc: "diffdut1:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port28 = &ConcretePort{Desc: "diffdut1:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port29 = &ConcretePort{Desc: "diffdut1:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port30 = &ConcretePort{Desc: "diffdut1:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port31 = &ConcretePort{Desc: "diffdut1:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port32 = &ConcretePort{Desc: "diffdut1:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port33 = &ConcretePort{Desc: "diffdut1:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port34 = &ConcretePort{Desc: "diffdut1:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port35 = &ConcretePort{Desc: "diffdut1:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port36 = &ConcretePort{Desc: "diffdut1:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port37 = &ConcretePort{Desc: "diffdut1:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port38 = &ConcretePort{Desc: "diffdut1:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port39 = &ConcretePort{Desc: "diffdut1:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port40 = &ConcretePort{Desc: "diffdut1:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port41 = &ConcretePort{Desc: "diffdut1:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port42 = &ConcretePort{Desc: "diffdut1:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port43 = &ConcretePort{Desc: "diffdut1:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port44 = &ConcretePort{Desc: "diffdut1:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port45 = &ConcretePort{Desc: "diffdut1:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port46 = &ConcretePort{Desc: "diffdut1:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port47 = &ConcretePort{Desc: "diffdut1:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port48 = &ConcretePort{Desc: "diffdut1:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port49 = &ConcretePort{Desc: "diffdut1:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port50 = &ConcretePort{Desc: "diffdut1:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port51 = &ConcretePort{Desc: "diffdut1:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port52 = &ConcretePort{Desc: "diffdut1:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port53 = &ConcretePort{Desc: "diffdut1:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port54 = &ConcretePort{Desc: "diffdut1:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port55 = &ConcretePort{Desc: "diffdut1:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port56 = &ConcretePort{Desc: "diffdut1:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port57 = &ConcretePort{Desc: "diffdut1:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port58 = &ConcretePort{Desc: "diffdut1:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port59 = &ConcretePort{Desc: "diffdut1:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port60 = &ConcretePort{Desc: "diffdut1:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port61 = &ConcretePort{Desc: "diffdut1:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port62 = &ConcretePort{Desc: "diffdut1:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port63 = &ConcretePort{Desc: "diffdut1:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port64 = &ConcretePort{Desc: "diffdut1:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port65 = &ConcretePort{Desc: "diffdut1:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port66 = &ConcretePort{Desc: "diffdut1:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port67 = &ConcretePort{Desc: "diffdut1:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port68 = &ConcretePort{Desc: "diffdut1:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port69 = &ConcretePort{Desc: "diffdut1:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1port70 = &ConcretePort{Desc: "diffdut1:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut1       = &ConcreteNode{Desc: "diffdut1", Ports: []*ConcretePort{diffdut1port1, diffdut1port2, diffdut1port3, diffdut1port4, diffdut1port5, diffdut1port6, diffdut1port7, diffdut1port8, diffdut1port9, diffdut1port10, diffdut1port11, diffdut1port12, diffdut1port13, diffdut1port14, diffdut1port15, diffdut1port16, diffdut1port17, diffdut1port18, diffdut1port19, diffdut1port20, diffdut1port21, diffdut1port22, diffdut1port23, diffdut1port24, diffdut1port25, diffdut1port26, diffdut1port27, diffdut1port28, diffdut1port29, diffdut1port30, diffdut1port31, diffdut1port32, diffdut1port33, diffdut1port34, diffdut1port35, diffdut1port36, diffdut1port37, diffdut1port38, diffdut1port39, diffdut1port40, diffdut1port41, diffdut1port42, diffdut1port43, diffdut1port44, diffdut1port45, diffdut1port46, diffdut1port47, diffdut1port48, diffdut1port49, diffdut1port50, diffdut1port51, diffdut1port52, diffdut1port53, diffdut1port54, diffdut1port55, diffdut1port56, diffdut1port57, diffdut1port58, diffdut1port59, diffdut1port60, diffdut1port61, diffdut1port62, diffdut1port63, diffdut1port64, diffdut1port65, diffdut1port66, diffdut1port67, diffdut1port68, diffdut1port69, diffdut1port70}, Attrs: map[string]string{"vendor": "CISCO", "role": "DUT", "model": "DUT"}}

	diffdut2port1  = &ConcretePort{Desc: "diffdut2:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port2  = &ConcretePort{Desc: "diffdut2:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port3  = &ConcretePort{Desc: "diffdut2:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port4  = &ConcretePort{Desc: "diffdut2:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port5  = &ConcretePort{Desc: "diffdut2:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port6  = &ConcretePort{Desc: "diffdut2:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port7  = &ConcretePort{Desc: "diffdut2:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port8  = &ConcretePort{Desc: "diffdut2:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port9  = &ConcretePort{Desc: "diffdut2:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port10 = &ConcretePort{Desc: "diffdut2:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut2port11 = &ConcretePort{Desc: "diffdut2:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port12 = &ConcretePort{Desc: "diffdut2:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port13 = &ConcretePort{Desc: "diffdut2:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port14 = &ConcretePort{Desc: "diffdut2:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port15 = &ConcretePort{Desc: "diffdut2:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port16 = &ConcretePort{Desc: "diffdut2:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port17 = &ConcretePort{Desc: "diffdut2:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port18 = &ConcretePort{Desc: "diffdut2:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port19 = &ConcretePort{Desc: "diffdut2:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port20 = &ConcretePort{Desc: "diffdut2:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut2port21 = &ConcretePort{Desc: "diffdut2:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port22 = &ConcretePort{Desc: "diffdut2:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port23 = &ConcretePort{Desc: "diffdut2:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port24 = &ConcretePort{Desc: "diffdut2:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port25 = &ConcretePort{Desc: "diffdut2:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port26 = &ConcretePort{Desc: "diffdut2:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port27 = &ConcretePort{Desc: "diffdut2:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port28 = &ConcretePort{Desc: "diffdut2:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port29 = &ConcretePort{Desc: "diffdut2:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port30 = &ConcretePort{Desc: "diffdut2:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port31 = &ConcretePort{Desc: "diffdut2:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port32 = &ConcretePort{Desc: "diffdut2:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port33 = &ConcretePort{Desc: "diffdut2:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port34 = &ConcretePort{Desc: "diffdut2:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port35 = &ConcretePort{Desc: "diffdut2:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port36 = &ConcretePort{Desc: "diffdut2:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port37 = &ConcretePort{Desc: "diffdut2:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port38 = &ConcretePort{Desc: "diffdut2:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port39 = &ConcretePort{Desc: "diffdut2:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port40 = &ConcretePort{Desc: "diffdut2:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port41 = &ConcretePort{Desc: "diffdut2:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port42 = &ConcretePort{Desc: "diffdut2:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port43 = &ConcretePort{Desc: "diffdut2:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port44 = &ConcretePort{Desc: "diffdut2:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port45 = &ConcretePort{Desc: "diffdut2:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port46 = &ConcretePort{Desc: "diffdut2:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port47 = &ConcretePort{Desc: "diffdut2:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port48 = &ConcretePort{Desc: "diffdut2:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port49 = &ConcretePort{Desc: "diffdut2:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port50 = &ConcretePort{Desc: "diffdut2:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port51 = &ConcretePort{Desc: "diffdut2:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port52 = &ConcretePort{Desc: "diffdut2:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port53 = &ConcretePort{Desc: "diffdut2:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port54 = &ConcretePort{Desc: "diffdut2:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port55 = &ConcretePort{Desc: "diffdut2:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port56 = &ConcretePort{Desc: "diffdut2:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port57 = &ConcretePort{Desc: "diffdut2:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port58 = &ConcretePort{Desc: "diffdut2:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port59 = &ConcretePort{Desc: "diffdut2:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port60 = &ConcretePort{Desc: "diffdut2:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port61 = &ConcretePort{Desc: "diffdut2:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port62 = &ConcretePort{Desc: "diffdut2:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port63 = &ConcretePort{Desc: "diffdut2:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port64 = &ConcretePort{Desc: "diffdut2:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port65 = &ConcretePort{Desc: "diffdut2:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port66 = &ConcretePort{Desc: "diffdut2:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port67 = &ConcretePort{Desc: "diffdut2:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port68 = &ConcretePort{Desc: "diffdut2:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port69 = &ConcretePort{Desc: "diffdut2:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2port70 = &ConcretePort{Desc: "diffdut2:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut2       = &ConcreteNode{Desc: "diffdut2", Ports: []*ConcretePort{diffdut2port1, diffdut2port2, diffdut2port3, diffdut2port4, diffdut2port5, diffdut2port6, diffdut2port7, diffdut2port8, diffdut2port9, diffdut2port10, diffdut2port11, diffdut2port12, diffdut2port13, diffdut2port14, diffdut2port15, diffdut2port16, diffdut2port17, diffdut2port18, diffdut2port19, diffdut2port20, diffdut2port21, diffdut2port22, diffdut2port23, diffdut2port24, diffdut2port25, diffdut2port26, diffdut2port27, diffdut2port28, diffdut2port29, diffdut2port30, diffdut2port31, diffdut2port32, diffdut2port33, diffdut2port34, diffdut2port35, diffdut2port36, diffdut2port37, diffdut2port38, diffdut2port39, diffdut2port40, diffdut2port41, diffdut2port42, diffdut2port43, diffdut2port44, diffdut2port45, diffdut2port46, diffdut2port47, diffdut2port48, diffdut2port49, diffdut2port50, diffdut2port51, diffdut2port52, diffdut2port53, diffdut2port54, diffdut2port55, diffdut2port56, diffdut2port57, diffdut2port58, diffdut2port59, diffdut2port60, diffdut2port61, diffdut2port62, diffdut2port63, diffdut2port64, diffdut2port65, diffdut2port66, diffdut2port67, diffdut2port68, diffdut2port69, diffdut2port70}, Attrs: map[string]string{"vendor": "CISCO", "role": "DUT", "model": "DUT"}}

	diffdut3port1  = &ConcretePort{Desc: "diffdut3:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port2  = &ConcretePort{Desc: "diffdut3:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port3  = &ConcretePort{Desc: "diffdut3:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port4  = &ConcretePort{Desc: "diffdut3:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port5  = &ConcretePort{Desc: "diffdut3:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port6  = &ConcretePort{Desc: "diffdut3:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port7  = &ConcretePort{Desc: "diffdut3:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port8  = &ConcretePort{Desc: "diffdut3:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port9  = &ConcretePort{Desc: "diffdut3:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port10 = &ConcretePort{Desc: "diffdut3:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut3port11 = &ConcretePort{Desc: "diffdut3:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port12 = &ConcretePort{Desc: "diffdut3:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port13 = &ConcretePort{Desc: "diffdut3:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port14 = &ConcretePort{Desc: "diffdut3:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port15 = &ConcretePort{Desc: "diffdut3:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port16 = &ConcretePort{Desc: "diffdut3:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port17 = &ConcretePort{Desc: "diffdut3:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port18 = &ConcretePort{Desc: "diffdut3:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port19 = &ConcretePort{Desc: "diffdut3:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port20 = &ConcretePort{Desc: "diffdut3:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut3port21 = &ConcretePort{Desc: "diffdut3:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port22 = &ConcretePort{Desc: "diffdut3:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port23 = &ConcretePort{Desc: "diffdut3:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port24 = &ConcretePort{Desc: "diffdut3:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port25 = &ConcretePort{Desc: "diffdut3:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port26 = &ConcretePort{Desc: "diffdut3:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port27 = &ConcretePort{Desc: "diffdut3:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port28 = &ConcretePort{Desc: "diffdut3:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port29 = &ConcretePort{Desc: "diffdut3:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port30 = &ConcretePort{Desc: "diffdut3:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port31 = &ConcretePort{Desc: "diffdut3:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port32 = &ConcretePort{Desc: "diffdut3:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port33 = &ConcretePort{Desc: "diffdut3:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port34 = &ConcretePort{Desc: "diffdut3:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port35 = &ConcretePort{Desc: "diffdut3:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port36 = &ConcretePort{Desc: "diffdut3:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port37 = &ConcretePort{Desc: "diffdut3:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port38 = &ConcretePort{Desc: "diffdut3:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port39 = &ConcretePort{Desc: "diffdut3:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port40 = &ConcretePort{Desc: "diffdut3:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port41 = &ConcretePort{Desc: "diffdut3:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port42 = &ConcretePort{Desc: "diffdut3:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port43 = &ConcretePort{Desc: "diffdut3:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port44 = &ConcretePort{Desc: "diffdut3:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port45 = &ConcretePort{Desc: "diffdut3:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port46 = &ConcretePort{Desc: "diffdut3:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port47 = &ConcretePort{Desc: "diffdut3:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port48 = &ConcretePort{Desc: "diffdut3:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port49 = &ConcretePort{Desc: "diffdut3:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port50 = &ConcretePort{Desc: "diffdut3:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port51 = &ConcretePort{Desc: "diffdut3:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port52 = &ConcretePort{Desc: "diffdut3:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port53 = &ConcretePort{Desc: "diffdut3:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port54 = &ConcretePort{Desc: "diffdut3:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port55 = &ConcretePort{Desc: "diffdut3:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port56 = &ConcretePort{Desc: "diffdut3:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port57 = &ConcretePort{Desc: "diffdut3:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port58 = &ConcretePort{Desc: "diffdut3:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port59 = &ConcretePort{Desc: "diffdut3:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port60 = &ConcretePort{Desc: "diffdut3:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port61 = &ConcretePort{Desc: "diffdut3:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port62 = &ConcretePort{Desc: "diffdut3:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port63 = &ConcretePort{Desc: "diffdut3:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port64 = &ConcretePort{Desc: "diffdut3:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port65 = &ConcretePort{Desc: "diffdut3:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port66 = &ConcretePort{Desc: "diffdut3:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port67 = &ConcretePort{Desc: "diffdut3:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port68 = &ConcretePort{Desc: "diffdut3:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port69 = &ConcretePort{Desc: "diffdut3:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3port70 = &ConcretePort{Desc: "diffdut3:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut3       = &ConcreteNode{Desc: "diffdut3", Ports: []*ConcretePort{diffdut3port1, diffdut3port2, diffdut3port3, diffdut3port4, diffdut3port5, diffdut3port6, diffdut3port7, diffdut3port8, diffdut3port9, diffdut3port10, diffdut3port11, diffdut3port12, diffdut3port13, diffdut3port14, diffdut3port15, diffdut3port16, diffdut3port17, diffdut3port18, diffdut3port19, diffdut3port20, diffdut3port21, diffdut3port22, diffdut3port23, diffdut3port24, diffdut3port25, diffdut3port26, diffdut3port27, diffdut3port28, diffdut3port29, diffdut3port30, diffdut3port31, diffdut3port32, diffdut3port33, diffdut3port34, diffdut3port35, diffdut3port36, diffdut3port37, diffdut3port38, diffdut3port39, diffdut3port40, diffdut3port41, diffdut3port42, diffdut3port43, diffdut3port44, diffdut3port45, diffdut3port46, diffdut3port47, diffdut3port48, diffdut3port49, diffdut3port50, diffdut3port51, diffdut3port52, diffdut3port53, diffdut3port54, diffdut3port55, diffdut3port56, diffdut3port57, diffdut3port58, diffdut3port59, diffdut3port60, diffdut3port61, diffdut3port62, diffdut3port63, diffdut3port64, diffdut3port65, diffdut3port66, diffdut3port67, diffdut3port68, diffdut3port69, diffdut3port70}, Attrs: map[string]string{"vendor": "CISCO", "role": "DUT", "model": "DUT"}}

	diffdut4port1  = &ConcretePort{Desc: "diffdut4:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port2  = &ConcretePort{Desc: "diffdut4:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port3  = &ConcretePort{Desc: "diffdut4:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port4  = &ConcretePort{Desc: "diffdut4:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port5  = &ConcretePort{Desc: "diffdut4:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port6  = &ConcretePort{Desc: "diffdut4:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port7  = &ConcretePort{Desc: "diffdut4:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port8  = &ConcretePort{Desc: "diffdut4:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port9  = &ConcretePort{Desc: "diffdut4:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port10 = &ConcretePort{Desc: "diffdut4:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffdut4port11 = &ConcretePort{Desc: "diffdut4:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port12 = &ConcretePort{Desc: "diffdut4:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port13 = &ConcretePort{Desc: "diffdut4:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port14 = &ConcretePort{Desc: "diffdut4:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port15 = &ConcretePort{Desc: "diffdut4:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port16 = &ConcretePort{Desc: "diffdut4:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port17 = &ConcretePort{Desc: "diffdut4:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port18 = &ConcretePort{Desc: "diffdut4:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port19 = &ConcretePort{Desc: "diffdut4:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port20 = &ConcretePort{Desc: "diffdut4:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffdut4port21 = &ConcretePort{Desc: "diffdut4:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port22 = &ConcretePort{Desc: "diffdut4:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port23 = &ConcretePort{Desc: "diffdut4:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port24 = &ConcretePort{Desc: "diffdut4:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port25 = &ConcretePort{Desc: "diffdut4:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port26 = &ConcretePort{Desc: "diffdut4:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port27 = &ConcretePort{Desc: "diffdut4:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port28 = &ConcretePort{Desc: "diffdut4:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port29 = &ConcretePort{Desc: "diffdut4:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port30 = &ConcretePort{Desc: "diffdut4:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port31 = &ConcretePort{Desc: "diffdut4:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port32 = &ConcretePort{Desc: "diffdut4:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port33 = &ConcretePort{Desc: "diffdut4:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port34 = &ConcretePort{Desc: "diffdut4:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port35 = &ConcretePort{Desc: "diffdut4:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port36 = &ConcretePort{Desc: "diffdut4:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port37 = &ConcretePort{Desc: "diffdut4:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port38 = &ConcretePort{Desc: "diffdut4:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port39 = &ConcretePort{Desc: "diffdut4:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port40 = &ConcretePort{Desc: "diffdut4:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port41 = &ConcretePort{Desc: "diffdut4:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port42 = &ConcretePort{Desc: "diffdut4:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port43 = &ConcretePort{Desc: "diffdut4:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port44 = &ConcretePort{Desc: "diffdut4:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port45 = &ConcretePort{Desc: "diffdut4:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port46 = &ConcretePort{Desc: "diffdut4:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port47 = &ConcretePort{Desc: "diffdut4:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port48 = &ConcretePort{Desc: "diffdut4:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port49 = &ConcretePort{Desc: "diffdut4:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port50 = &ConcretePort{Desc: "diffdut4:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port51 = &ConcretePort{Desc: "diffdut4:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port52 = &ConcretePort{Desc: "diffdut4:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port53 = &ConcretePort{Desc: "diffdut4:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port54 = &ConcretePort{Desc: "diffdut4:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port55 = &ConcretePort{Desc: "diffdut4:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port56 = &ConcretePort{Desc: "diffdut4:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port57 = &ConcretePort{Desc: "diffdut4:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port58 = &ConcretePort{Desc: "diffdut4:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port59 = &ConcretePort{Desc: "diffdut4:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port60 = &ConcretePort{Desc: "diffdut4:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port61 = &ConcretePort{Desc: "diffdut4:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port62 = &ConcretePort{Desc: "diffdut4:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port63 = &ConcretePort{Desc: "diffdut4:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port64 = &ConcretePort{Desc: "diffdut4:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port65 = &ConcretePort{Desc: "diffdut4:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port66 = &ConcretePort{Desc: "diffdut4:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port67 = &ConcretePort{Desc: "diffdut4:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port68 = &ConcretePort{Desc: "diffdut4:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port69 = &ConcretePort{Desc: "diffdut4:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4port70 = &ConcretePort{Desc: "diffdut4:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffdut4       = &ConcreteNode{Desc: "diffdut4", Ports: []*ConcretePort{diffdut4port1, diffdut4port2, diffdut4port3, diffdut4port4, diffdut4port5, diffdut4port6, diffdut4port7, diffdut4port8, diffdut4port9, diffdut4port10, diffdut4port11, diffdut4port12, diffdut4port13, diffdut4port14, diffdut4port15, diffdut4port16, diffdut4port17, diffdut4port18, diffdut4port19, diffdut4port20, diffdut4port21, diffdut4port22, diffdut4port23, diffdut4port24, diffdut4port25, diffdut4port26, diffdut4port27, diffdut4port28, diffdut4port29, diffdut4port30, diffdut4port31, diffdut4port32, diffdut4port33, diffdut4port34, diffdut4port35, diffdut4port36, diffdut4port37, diffdut4port38, diffdut4port39, diffdut4port40, diffdut4port41, diffdut4port42, diffdut4port43, diffdut4port44, diffdut4port45, diffdut4port46, diffdut4port47, diffdut4port48, diffdut4port49, diffdut4port50, diffdut4port51, diffdut4port52, diffdut4port53, diffdut4port54, diffdut4port55, diffdut4port56, diffdut4port57, diffdut4port58, diffdut4port59, diffdut4port60, diffdut4port61, diffdut4port62, diffdut4port63, diffdut4port64, diffdut4port65, diffdut4port66, diffdut4port67, diffdut4port68, diffdut4port69, diffdut4port70}, Attrs: map[string]string{"vendor": "CISCO", "role": "DUT", "model": "DUT"}}

	diffate1port1  = &ConcretePort{Desc: "diffate1:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port2  = &ConcretePort{Desc: "diffate1:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port3  = &ConcretePort{Desc: "diffate1:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port4  = &ConcretePort{Desc: "diffate1:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port5  = &ConcretePort{Desc: "diffate1:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port6  = &ConcretePort{Desc: "diffate1:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port7  = &ConcretePort{Desc: "diffate1:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port8  = &ConcretePort{Desc: "diffate1:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port9  = &ConcretePort{Desc: "diffate1:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port10 = &ConcretePort{Desc: "diffate1:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate1port11 = &ConcretePort{Desc: "diffate1:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port12 = &ConcretePort{Desc: "diffate1:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port13 = &ConcretePort{Desc: "diffate1:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port14 = &ConcretePort{Desc: "diffate1:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port15 = &ConcretePort{Desc: "diffate1:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port16 = &ConcretePort{Desc: "diffate1:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port17 = &ConcretePort{Desc: "diffate1:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port18 = &ConcretePort{Desc: "diffate1:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port19 = &ConcretePort{Desc: "diffate1:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port20 = &ConcretePort{Desc: "diffate1:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate1port21 = &ConcretePort{Desc: "diffate1:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port22 = &ConcretePort{Desc: "diffate1:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port23 = &ConcretePort{Desc: "diffate1:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port24 = &ConcretePort{Desc: "diffate1:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port25 = &ConcretePort{Desc: "diffate1:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port26 = &ConcretePort{Desc: "diffate1:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port27 = &ConcretePort{Desc: "diffate1:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port28 = &ConcretePort{Desc: "diffate1:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port29 = &ConcretePort{Desc: "diffate1:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port30 = &ConcretePort{Desc: "diffate1:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port31 = &ConcretePort{Desc: "diffate1:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port32 = &ConcretePort{Desc: "diffate1:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port33 = &ConcretePort{Desc: "diffate1:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port34 = &ConcretePort{Desc: "diffate1:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port35 = &ConcretePort{Desc: "diffate1:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port36 = &ConcretePort{Desc: "diffate1:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port37 = &ConcretePort{Desc: "diffate1:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port38 = &ConcretePort{Desc: "diffate1:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port39 = &ConcretePort{Desc: "diffate1:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port40 = &ConcretePort{Desc: "diffate1:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port41 = &ConcretePort{Desc: "diffate1:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port42 = &ConcretePort{Desc: "diffate1:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port43 = &ConcretePort{Desc: "diffate1:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port44 = &ConcretePort{Desc: "diffate1:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port45 = &ConcretePort{Desc: "diffate1:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port46 = &ConcretePort{Desc: "diffate1:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port47 = &ConcretePort{Desc: "diffate1:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port48 = &ConcretePort{Desc: "diffate1:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port49 = &ConcretePort{Desc: "diffate1:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port50 = &ConcretePort{Desc: "diffate1:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port51 = &ConcretePort{Desc: "diffate1:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port52 = &ConcretePort{Desc: "diffate1:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port53 = &ConcretePort{Desc: "diffate1:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port54 = &ConcretePort{Desc: "diffate1:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port55 = &ConcretePort{Desc: "diffate1:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port56 = &ConcretePort{Desc: "diffate1:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port57 = &ConcretePort{Desc: "diffate1:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port58 = &ConcretePort{Desc: "diffate1:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port59 = &ConcretePort{Desc: "diffate1:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port60 = &ConcretePort{Desc: "diffate1:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port61 = &ConcretePort{Desc: "diffate1:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port62 = &ConcretePort{Desc: "diffate1:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port63 = &ConcretePort{Desc: "diffate1:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port64 = &ConcretePort{Desc: "diffate1:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port65 = &ConcretePort{Desc: "diffate1:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port66 = &ConcretePort{Desc: "diffate1:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port67 = &ConcretePort{Desc: "diffate1:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port68 = &ConcretePort{Desc: "diffate1:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port69 = &ConcretePort{Desc: "diffate1:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1port70 = &ConcretePort{Desc: "diffate1:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate1       = &ConcreteNode{Desc: "diffate1", Ports: []*ConcretePort{diffate1port1, diffate1port2, diffate1port3, diffate1port4, diffate1port5, diffate1port6, diffate1port7, diffate1port8, diffate1port9, diffate1port10, diffate1port11, diffate1port12, diffate1port13, diffate1port14, diffate1port15, diffate1port16, diffate1port17, diffate1port18, diffate1port19, diffate1port20, diffate1port21, diffate1port22, diffate1port23, diffate1port24, diffate1port25, diffate1port26, diffate1port27, diffate1port28, diffate1port29, diffate1port30, diffate1port31, diffate1port32, diffate1port33, diffate1port34, diffate1port35, diffate1port36, diffate1port37, diffate1port38, diffate1port39, diffate1port40, diffate1port41, diffate1port42, diffate1port43, diffate1port44, diffate1port45, diffate1port46, diffate1port47, diffate1port48, diffate1port49, diffate1port50, diffate1port51, diffate1port52, diffate1port53, diffate1port54, diffate1port55, diffate1port56, diffate1port57, diffate1port58, diffate1port59, diffate1port60, diffate1port61, diffate1port62, diffate1port63, diffate1port64, diffate1port65, diffate1port66, diffate1port67, diffate1port68, diffate1port69, diffate1port70}, Attrs: map[string]string{"vendor": "TGEN", "role": "ATE", "model": "ATE"}}

	diffate2port1  = &ConcretePort{Desc: "diffate2:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port2  = &ConcretePort{Desc: "diffate2:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port3  = &ConcretePort{Desc: "diffate2:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port4  = &ConcretePort{Desc: "diffate2:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port5  = &ConcretePort{Desc: "diffate2:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port6  = &ConcretePort{Desc: "diffate2:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port7  = &ConcretePort{Desc: "diffate2:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port8  = &ConcretePort{Desc: "diffate2:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port9  = &ConcretePort{Desc: "diffate2:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port10 = &ConcretePort{Desc: "diffate2:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate2port11 = &ConcretePort{Desc: "diffate2:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port12 = &ConcretePort{Desc: "diffate2:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port13 = &ConcretePort{Desc: "diffate2:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port14 = &ConcretePort{Desc: "diffate2:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port15 = &ConcretePort{Desc: "diffate2:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port16 = &ConcretePort{Desc: "diffate2:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port17 = &ConcretePort{Desc: "diffate2:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port18 = &ConcretePort{Desc: "diffate2:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port19 = &ConcretePort{Desc: "diffate2:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port20 = &ConcretePort{Desc: "diffate2:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate2port21 = &ConcretePort{Desc: "diffate2:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port22 = &ConcretePort{Desc: "diffate2:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port23 = &ConcretePort{Desc: "diffate2:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port24 = &ConcretePort{Desc: "diffate2:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port25 = &ConcretePort{Desc: "diffate2:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port26 = &ConcretePort{Desc: "diffate2:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port27 = &ConcretePort{Desc: "diffate2:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port28 = &ConcretePort{Desc: "diffate2:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port29 = &ConcretePort{Desc: "diffate2:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port30 = &ConcretePort{Desc: "diffate2:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port31 = &ConcretePort{Desc: "diffate2:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port32 = &ConcretePort{Desc: "diffate2:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port33 = &ConcretePort{Desc: "diffate2:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port34 = &ConcretePort{Desc: "diffate2:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port35 = &ConcretePort{Desc: "diffate2:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port36 = &ConcretePort{Desc: "diffate2:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port37 = &ConcretePort{Desc: "diffate2:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port38 = &ConcretePort{Desc: "diffate2:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port39 = &ConcretePort{Desc: "diffate2:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port40 = &ConcretePort{Desc: "diffate2:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port41 = &ConcretePort{Desc: "diffate2:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port42 = &ConcretePort{Desc: "diffate2:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port43 = &ConcretePort{Desc: "diffate2:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port44 = &ConcretePort{Desc: "diffate2:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port45 = &ConcretePort{Desc: "diffate2:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port46 = &ConcretePort{Desc: "diffate2:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port47 = &ConcretePort{Desc: "diffate2:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port48 = &ConcretePort{Desc: "diffate2:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port49 = &ConcretePort{Desc: "diffate2:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port50 = &ConcretePort{Desc: "diffate2:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port51 = &ConcretePort{Desc: "diffate2:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port52 = &ConcretePort{Desc: "diffate2:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port53 = &ConcretePort{Desc: "diffate2:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port54 = &ConcretePort{Desc: "diffate2:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port55 = &ConcretePort{Desc: "diffate2:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port56 = &ConcretePort{Desc: "diffate2:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port57 = &ConcretePort{Desc: "diffate2:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port58 = &ConcretePort{Desc: "diffate2:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port59 = &ConcretePort{Desc: "diffate2:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port60 = &ConcretePort{Desc: "diffate2:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port61 = &ConcretePort{Desc: "diffate2:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port62 = &ConcretePort{Desc: "diffate2:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port63 = &ConcretePort{Desc: "diffate2:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port64 = &ConcretePort{Desc: "diffate2:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port65 = &ConcretePort{Desc: "diffate2:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port66 = &ConcretePort{Desc: "diffate2:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port67 = &ConcretePort{Desc: "diffate2:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port68 = &ConcretePort{Desc: "diffate2:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port69 = &ConcretePort{Desc: "diffate2:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2port70 = &ConcretePort{Desc: "diffate2:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate2       = &ConcreteNode{Desc: "diffate2", Ports: []*ConcretePort{diffate2port1, diffate2port2, diffate2port3, diffate2port4, diffate2port5, diffate2port6, diffate2port7, diffate2port8, diffate2port9, diffate2port10, diffate2port11, diffate2port12, diffate2port13, diffate2port14, diffate2port15, diffate2port16, diffate2port17, diffate2port18, diffate2port19, diffate2port20, diffate2port21, diffate2port22, diffate2port23, diffate2port24, diffate2port25, diffate2port26, diffate2port27, diffate2port28, diffate2port29, diffate2port30, diffate2port31, diffate2port32, diffate2port33, diffate2port34, diffate2port35, diffate2port36, diffate2port37, diffate2port38, diffate2port39, diffate2port40, diffate2port41, diffate2port42, diffate2port43, diffate2port44, diffate2port45, diffate2port46, diffate2port47, diffate2port48, diffate2port49, diffate2port50, diffate2port51, diffate2port52, diffate2port53, diffate2port54, diffate2port55, diffate2port56, diffate2port57, diffate2port58, diffate2port59, diffate2port60, diffate2port61, diffate2port62, diffate2port63, diffate2port64, diffate2port65, diffate2port66, diffate2port67, diffate2port68, diffate2port69, diffate2port70}, Attrs: map[string]string{"vendor": "TGEN", "role": "ATE", "model": "ATE"}}

	diffate3port1  = &ConcretePort{Desc: "diffate3:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port2  = &ConcretePort{Desc: "diffate3:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port3  = &ConcretePort{Desc: "diffate3:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port4  = &ConcretePort{Desc: "diffate3:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port5  = &ConcretePort{Desc: "diffate3:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port6  = &ConcretePort{Desc: "diffate3:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port7  = &ConcretePort{Desc: "diffate3:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port8  = &ConcretePort{Desc: "diffate3:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port9  = &ConcretePort{Desc: "diffate3:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port10 = &ConcretePort{Desc: "diffate3:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate3port11 = &ConcretePort{Desc: "diffate3:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port12 = &ConcretePort{Desc: "diffate3:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port13 = &ConcretePort{Desc: "diffate3:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port14 = &ConcretePort{Desc: "diffate3:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port15 = &ConcretePort{Desc: "diffate3:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port16 = &ConcretePort{Desc: "diffate3:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port17 = &ConcretePort{Desc: "diffate3:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port18 = &ConcretePort{Desc: "diffate3:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port19 = &ConcretePort{Desc: "diffate3:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port20 = &ConcretePort{Desc: "diffate3:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate3port21 = &ConcretePort{Desc: "diffate3:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port22 = &ConcretePort{Desc: "diffate3:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port23 = &ConcretePort{Desc: "diffate3:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port24 = &ConcretePort{Desc: "diffate3:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port25 = &ConcretePort{Desc: "diffate3:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port26 = &ConcretePort{Desc: "diffate3:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port27 = &ConcretePort{Desc: "diffate3:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port28 = &ConcretePort{Desc: "diffate3:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port29 = &ConcretePort{Desc: "diffate3:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port30 = &ConcretePort{Desc: "diffate3:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port31 = &ConcretePort{Desc: "diffate3:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port32 = &ConcretePort{Desc: "diffate3:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port33 = &ConcretePort{Desc: "diffate3:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port34 = &ConcretePort{Desc: "diffate3:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port35 = &ConcretePort{Desc: "diffate3:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port36 = &ConcretePort{Desc: "diffate3:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port37 = &ConcretePort{Desc: "diffate3:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port38 = &ConcretePort{Desc: "diffate3:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port39 = &ConcretePort{Desc: "diffate3:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port40 = &ConcretePort{Desc: "diffate3:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port41 = &ConcretePort{Desc: "diffate3:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port42 = &ConcretePort{Desc: "diffate3:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port43 = &ConcretePort{Desc: "diffate3:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port44 = &ConcretePort{Desc: "diffate3:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port45 = &ConcretePort{Desc: "diffate3:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port46 = &ConcretePort{Desc: "diffate3:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port47 = &ConcretePort{Desc: "diffate3:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port48 = &ConcretePort{Desc: "diffate3:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port49 = &ConcretePort{Desc: "diffate3:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port50 = &ConcretePort{Desc: "diffate3:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port51 = &ConcretePort{Desc: "diffate3:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port52 = &ConcretePort{Desc: "diffate3:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port53 = &ConcretePort{Desc: "diffate3:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port54 = &ConcretePort{Desc: "diffate3:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port55 = &ConcretePort{Desc: "diffate3:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port56 = &ConcretePort{Desc: "diffate3:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port57 = &ConcretePort{Desc: "diffate3:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port58 = &ConcretePort{Desc: "diffate3:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port59 = &ConcretePort{Desc: "diffate3:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port60 = &ConcretePort{Desc: "diffate3:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port61 = &ConcretePort{Desc: "diffate3:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port62 = &ConcretePort{Desc: "diffate3:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port63 = &ConcretePort{Desc: "diffate3:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port64 = &ConcretePort{Desc: "diffate3:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port65 = &ConcretePort{Desc: "diffate3:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port66 = &ConcretePort{Desc: "diffate3:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port67 = &ConcretePort{Desc: "diffate3:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port68 = &ConcretePort{Desc: "diffate3:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port69 = &ConcretePort{Desc: "diffate3:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3port70 = &ConcretePort{Desc: "diffate3:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate3       = &ConcreteNode{Desc: "diffate3", Ports: []*ConcretePort{diffate3port1, diffate3port2, diffate3port3, diffate3port4, diffate3port5, diffate3port6, diffate3port7, diffate3port8, diffate3port9, diffate3port10, diffate3port11, diffate3port12, diffate3port13, diffate3port14, diffate3port15, diffate3port16, diffate3port17, diffate3port18, diffate3port19, diffate3port20, diffate3port21, diffate3port22, diffate3port23, diffate3port24, diffate3port25, diffate3port26, diffate3port27, diffate3port28, diffate3port29, diffate3port30, diffate3port31, diffate3port32, diffate3port33, diffate3port34, diffate3port35, diffate3port36, diffate3port37, diffate3port38, diffate3port39, diffate3port40, diffate3port41, diffate3port42, diffate3port43, diffate3port44, diffate3port45, diffate3port46, diffate3port47, diffate3port48, diffate3port49, diffate3port50, diffate3port51, diffate3port52, diffate3port53, diffate3port54, diffate3port55, diffate3port56, diffate3port57, diffate3port58, diffate3port59, diffate3port60, diffate3port61, diffate3port62, diffate3port63, diffate3port64, diffate3port65, diffate3port66, diffate3port67, diffate3port68, diffate3port69, diffate3port70}, Attrs: map[string]string{"vendor": "TGEN", "role": "ATE", "model": "ATE"}}

	diffate4port1  = &ConcretePort{Desc: "diffate4:port1", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port2  = &ConcretePort{Desc: "diffate4:port2", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port3  = &ConcretePort{Desc: "diffate4:port3", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port4  = &ConcretePort{Desc: "diffate4:port4", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port5  = &ConcretePort{Desc: "diffate4:port5", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port6  = &ConcretePort{Desc: "diffate4:port6", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port7  = &ConcretePort{Desc: "diffate4:port7", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port8  = &ConcretePort{Desc: "diffate4:port8", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port9  = &ConcretePort{Desc: "diffate4:port9", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port10 = &ConcretePort{Desc: "diffate4:port10", Attrs: map[string]string{"speed": "S_100G", "media": "copper"}}
	diffate4port11 = &ConcretePort{Desc: "diffate4:port11", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port12 = &ConcretePort{Desc: "diffate4:port12", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port13 = &ConcretePort{Desc: "diffate4:port13", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port14 = &ConcretePort{Desc: "diffate4:port14", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port15 = &ConcretePort{Desc: "diffate4:port15", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port16 = &ConcretePort{Desc: "diffate4:port16", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port17 = &ConcretePort{Desc: "diffate4:port17", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port18 = &ConcretePort{Desc: "diffate4:port18", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port19 = &ConcretePort{Desc: "diffate4:port19", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port20 = &ConcretePort{Desc: "diffate4:port20", Attrs: map[string]string{"speed": "S_200G", "media": "copper"}}
	diffate4port21 = &ConcretePort{Desc: "diffate4:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port22 = &ConcretePort{Desc: "diffate4:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port23 = &ConcretePort{Desc: "diffate4:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port24 = &ConcretePort{Desc: "diffate4:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port25 = &ConcretePort{Desc: "diffate4:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port26 = &ConcretePort{Desc: "diffate4:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port27 = &ConcretePort{Desc: "diffate4:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port28 = &ConcretePort{Desc: "diffate4:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port29 = &ConcretePort{Desc: "diffate4:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port30 = &ConcretePort{Desc: "diffate4:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port31 = &ConcretePort{Desc: "diffate4:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port32 = &ConcretePort{Desc: "diffate4:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port33 = &ConcretePort{Desc: "diffate4:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port34 = &ConcretePort{Desc: "diffate4:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port35 = &ConcretePort{Desc: "diffate4:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port36 = &ConcretePort{Desc: "diffate4:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port37 = &ConcretePort{Desc: "diffate4:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port38 = &ConcretePort{Desc: "diffate4:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port39 = &ConcretePort{Desc: "diffate4:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port40 = &ConcretePort{Desc: "diffate4:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port41 = &ConcretePort{Desc: "diffate4:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port42 = &ConcretePort{Desc: "diffate4:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port43 = &ConcretePort{Desc: "diffate4:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port44 = &ConcretePort{Desc: "diffate4:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port45 = &ConcretePort{Desc: "diffate4:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port46 = &ConcretePort{Desc: "diffate4:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port47 = &ConcretePort{Desc: "diffate4:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port48 = &ConcretePort{Desc: "diffate4:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port49 = &ConcretePort{Desc: "diffate4:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port50 = &ConcretePort{Desc: "diffate4:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port51 = &ConcretePort{Desc: "diffate4:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port52 = &ConcretePort{Desc: "diffate4:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port53 = &ConcretePort{Desc: "diffate4:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port54 = &ConcretePort{Desc: "diffate4:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port55 = &ConcretePort{Desc: "diffate4:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port56 = &ConcretePort{Desc: "diffate4:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port57 = &ConcretePort{Desc: "diffate4:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port58 = &ConcretePort{Desc: "diffate4:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port59 = &ConcretePort{Desc: "diffate4:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port60 = &ConcretePort{Desc: "diffate4:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port61 = &ConcretePort{Desc: "diffate4:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port62 = &ConcretePort{Desc: "diffate4:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port63 = &ConcretePort{Desc: "diffate4:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port64 = &ConcretePort{Desc: "diffate4:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port65 = &ConcretePort{Desc: "diffate4:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port66 = &ConcretePort{Desc: "diffate4:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port67 = &ConcretePort{Desc: "diffate4:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port68 = &ConcretePort{Desc: "diffate4:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port69 = &ConcretePort{Desc: "diffate4:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4port70 = &ConcretePort{Desc: "diffate4:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffate4       = &ConcreteNode{Desc: "diffate4", Ports: []*ConcretePort{diffate4port1, diffate4port2, diffate4port3, diffate4port4, diffate4port5, diffate4port6, diffate4port7, diffate4port8, diffate4port9, diffate4port10, diffate4port11, diffate4port12, diffate4port13, diffate4port14, diffate4port15, diffate4port16, diffate4port17, diffate4port18, diffate4port19, diffate4port20, diffate4port21, diffate4port22, diffate4port23, diffate4port24, diffate4port25, diffate4port26, diffate4port27, diffate4port28, diffate4port29, diffate4port30, diffate4port31, diffate4port32, diffate4port33, diffate4port34, diffate4port35, diffate4port36, diffate4port37, diffate4port38, diffate4port39, diffate4port40, diffate4port41, diffate4port42, diffate4port43, diffate4port44, diffate4port45, diffate4port46, diffate4port47, diffate4port48, diffate4port49, diffate4port50, diffate4port51, diffate4port52, diffate4port53, diffate4port54, diffate4port55, diffate4port56, diffate4port57, diffate4port58, diffate4port59, diffate4port60, diffate4port61, diffate4port62, diffate4port63, diffate4port64, diffate4port65, diffate4port66, diffate4port67, diffate4port68, diffate4port69, diffate4port70}, Attrs: map[string]string{"vendor": "TGEN", "role": "ATE", "model": "ATE"}}

	diffswitch1port1   = &ConcretePort{Desc: "diffswitch1:port1", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port2   = &ConcretePort{Desc: "diffswitch1:port2", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port3   = &ConcretePort{Desc: "diffswitch1:port3", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port4   = &ConcretePort{Desc: "diffswitch1:port4", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port5   = &ConcretePort{Desc: "diffswitch1:port5", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port6   = &ConcretePort{Desc: "diffswitch1:port6", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port7   = &ConcretePort{Desc: "diffswitch1:port7", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port8   = &ConcretePort{Desc: "diffswitch1:port8", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port9   = &ConcretePort{Desc: "diffswitch1:port9", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port10  = &ConcretePort{Desc: "diffswitch1:port10", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port11  = &ConcretePort{Desc: "diffswitch1:port11", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port12  = &ConcretePort{Desc: "diffswitch1:port12", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port13  = &ConcretePort{Desc: "diffswitch1:port13", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port14  = &ConcretePort{Desc: "diffswitch1:port14", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port15  = &ConcretePort{Desc: "diffswitch1:port15", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port16  = &ConcretePort{Desc: "diffswitch1:port16", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port17  = &ConcretePort{Desc: "diffswitch1:port17", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port18  = &ConcretePort{Desc: "diffswitch1:port18", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port19  = &ConcretePort{Desc: "diffswitch1:port19", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port20  = &ConcretePort{Desc: "diffswitch1:port20", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port21  = &ConcretePort{Desc: "diffswitch1:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port22  = &ConcretePort{Desc: "diffswitch1:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port23  = &ConcretePort{Desc: "diffswitch1:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port24  = &ConcretePort{Desc: "diffswitch1:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port25  = &ConcretePort{Desc: "diffswitch1:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port26  = &ConcretePort{Desc: "diffswitch1:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port27  = &ConcretePort{Desc: "diffswitch1:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port28  = &ConcretePort{Desc: "diffswitch1:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port29  = &ConcretePort{Desc: "diffswitch1:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port30  = &ConcretePort{Desc: "diffswitch1:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port31  = &ConcretePort{Desc: "diffswitch1:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port32  = &ConcretePort{Desc: "diffswitch1:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port33  = &ConcretePort{Desc: "diffswitch1:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port34  = &ConcretePort{Desc: "diffswitch1:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port35  = &ConcretePort{Desc: "diffswitch1:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port36  = &ConcretePort{Desc: "diffswitch1:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port37  = &ConcretePort{Desc: "diffswitch1:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port38  = &ConcretePort{Desc: "diffswitch1:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port39  = &ConcretePort{Desc: "diffswitch1:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port40  = &ConcretePort{Desc: "diffswitch1:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port41  = &ConcretePort{Desc: "diffswitch1:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port42  = &ConcretePort{Desc: "diffswitch1:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port43  = &ConcretePort{Desc: "diffswitch1:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port44  = &ConcretePort{Desc: "diffswitch1:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port45  = &ConcretePort{Desc: "diffswitch1:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port46  = &ConcretePort{Desc: "diffswitch1:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port47  = &ConcretePort{Desc: "diffswitch1:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port48  = &ConcretePort{Desc: "diffswitch1:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port49  = &ConcretePort{Desc: "diffswitch1:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port50  = &ConcretePort{Desc: "diffswitch1:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port51  = &ConcretePort{Desc: "diffswitch1:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port52  = &ConcretePort{Desc: "diffswitch1:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port53  = &ConcretePort{Desc: "diffswitch1:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port54  = &ConcretePort{Desc: "diffswitch1:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port55  = &ConcretePort{Desc: "diffswitch1:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port56  = &ConcretePort{Desc: "diffswitch1:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port57  = &ConcretePort{Desc: "diffswitch1:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port58  = &ConcretePort{Desc: "diffswitch1:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port59  = &ConcretePort{Desc: "diffswitch1:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port60  = &ConcretePort{Desc: "diffswitch1:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port61  = &ConcretePort{Desc: "diffswitch1:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port62  = &ConcretePort{Desc: "diffswitch1:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port63  = &ConcretePort{Desc: "diffswitch1:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port64  = &ConcretePort{Desc: "diffswitch1:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port65  = &ConcretePort{Desc: "diffswitch1:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port66  = &ConcretePort{Desc: "diffswitch1:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port67  = &ConcretePort{Desc: "diffswitch1:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port68  = &ConcretePort{Desc: "diffswitch1:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port69  = &ConcretePort{Desc: "diffswitch1:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port70  = &ConcretePort{Desc: "diffswitch1:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port71  = &ConcretePort{Desc: "diffswitch1:port71", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port72  = &ConcretePort{Desc: "diffswitch1:port72", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port73  = &ConcretePort{Desc: "diffswitch1:port73", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port74  = &ConcretePort{Desc: "diffswitch1:port74", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port75  = &ConcretePort{Desc: "diffswitch1:port75", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port76  = &ConcretePort{Desc: "diffswitch1:port76", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port77  = &ConcretePort{Desc: "diffswitch1:port77", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port78  = &ConcretePort{Desc: "diffswitch1:port78", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port79  = &ConcretePort{Desc: "diffswitch1:port79", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port80  = &ConcretePort{Desc: "diffswitch1:port80", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port81  = &ConcretePort{Desc: "diffswitch1:port81", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port82  = &ConcretePort{Desc: "diffswitch1:port82", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port83  = &ConcretePort{Desc: "diffswitch1:port83", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port84  = &ConcretePort{Desc: "diffswitch1:port84", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port85  = &ConcretePort{Desc: "diffswitch1:port85", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port86  = &ConcretePort{Desc: "diffswitch1:port86", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port87  = &ConcretePort{Desc: "diffswitch1:port87", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port88  = &ConcretePort{Desc: "diffswitch1:port88", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port89  = &ConcretePort{Desc: "diffswitch1:port89", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port90  = &ConcretePort{Desc: "diffswitch1:port90", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port91  = &ConcretePort{Desc: "diffswitch1:port91", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port92  = &ConcretePort{Desc: "diffswitch1:port92", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port93  = &ConcretePort{Desc: "diffswitch1:port93", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port94  = &ConcretePort{Desc: "diffswitch1:port94", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port95  = &ConcretePort{Desc: "diffswitch1:port95", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port96  = &ConcretePort{Desc: "diffswitch1:port96", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port97  = &ConcretePort{Desc: "diffswitch1:port97", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port98  = &ConcretePort{Desc: "diffswitch1:port98", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port99  = &ConcretePort{Desc: "diffswitch1:port99", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port100 = &ConcretePort{Desc: "diffswitch1:port100", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port101 = &ConcretePort{Desc: "diffswitch1:port101", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port102 = &ConcretePort{Desc: "diffswitch1:port102", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port103 = &ConcretePort{Desc: "diffswitch1:port103", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port104 = &ConcretePort{Desc: "diffswitch1:port104", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port105 = &ConcretePort{Desc: "diffswitch1:port105", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port106 = &ConcretePort{Desc: "diffswitch1:port106", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port107 = &ConcretePort{Desc: "diffswitch1:port107", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port108 = &ConcretePort{Desc: "diffswitch1:port108", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port109 = &ConcretePort{Desc: "diffswitch1:port109", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port110 = &ConcretePort{Desc: "diffswitch1:port110", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port111 = &ConcretePort{Desc: "diffswitch1:port111", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port112 = &ConcretePort{Desc: "diffswitch1:port112", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port113 = &ConcretePort{Desc: "diffswitch1:port113", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port114 = &ConcretePort{Desc: "diffswitch1:port114", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port115 = &ConcretePort{Desc: "diffswitch1:port115", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port116 = &ConcretePort{Desc: "diffswitch1:port116", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port117 = &ConcretePort{Desc: "diffswitch1:port117", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port118 = &ConcretePort{Desc: "diffswitch1:port118", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port119 = &ConcretePort{Desc: "diffswitch1:port119", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port120 = &ConcretePort{Desc: "diffswitch1:port120", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port121 = &ConcretePort{Desc: "diffswitch1:port121", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port122 = &ConcretePort{Desc: "diffswitch1:port122", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port123 = &ConcretePort{Desc: "diffswitch1:port123", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port124 = &ConcretePort{Desc: "diffswitch1:port124", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port125 = &ConcretePort{Desc: "diffswitch1:port125", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port126 = &ConcretePort{Desc: "diffswitch1:port126", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port127 = &ConcretePort{Desc: "diffswitch1:port127", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port128 = &ConcretePort{Desc: "diffswitch1:port128", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port129 = &ConcretePort{Desc: "diffswitch1:port129", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port130 = &ConcretePort{Desc: "diffswitch1:port130", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port131 = &ConcretePort{Desc: "diffswitch1:port131", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port132 = &ConcretePort{Desc: "diffswitch1:port132", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port133 = &ConcretePort{Desc: "diffswitch1:port133", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port134 = &ConcretePort{Desc: "diffswitch1:port134", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port135 = &ConcretePort{Desc: "diffswitch1:port135", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port136 = &ConcretePort{Desc: "diffswitch1:port136", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port137 = &ConcretePort{Desc: "diffswitch1:port137", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port138 = &ConcretePort{Desc: "diffswitch1:port138", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port139 = &ConcretePort{Desc: "diffswitch1:port139", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port140 = &ConcretePort{Desc: "diffswitch1:port140", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch1port141 = &ConcretePort{Desc: "diffswitch1:port141", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}

	diffswitch1 = &ConcreteNode{Desc: "diffswitch1", Ports: []*ConcretePort{
		diffswitch1port1,
		diffswitch1port2,
		diffswitch1port3,
		diffswitch1port4,
		diffswitch1port5,
		diffswitch1port6,
		diffswitch1port7,
		diffswitch1port8,
		diffswitch1port9,
		diffswitch1port10,
		diffswitch1port11,
		diffswitch1port12,
		diffswitch1port13,
		diffswitch1port14,
		diffswitch1port15,
		diffswitch1port16,
		diffswitch1port17,
		diffswitch1port18,
		diffswitch1port19,
		diffswitch1port20,
		diffswitch1port21,
		diffswitch1port22,
		diffswitch1port23,
		diffswitch1port24,
		diffswitch1port25,
		diffswitch1port26,
		diffswitch1port27,
		diffswitch1port28,
		diffswitch1port29,
		diffswitch1port30,
		diffswitch1port31,
		diffswitch1port32,
		diffswitch1port33,
		diffswitch1port34,
		diffswitch1port35,
		diffswitch1port36,
		diffswitch1port37,
		diffswitch1port38,
		diffswitch1port39,
		diffswitch1port40,
		diffswitch1port41,
		diffswitch1port42,
		diffswitch1port43,
		diffswitch1port44,
		diffswitch1port45,
		diffswitch1port46,
		diffswitch1port47,
		diffswitch1port48,
		diffswitch1port49,
		diffswitch1port50,
		diffswitch1port51,
		diffswitch1port52,
		diffswitch1port53,
		diffswitch1port54,
		diffswitch1port55,
		diffswitch1port56,
		diffswitch1port57,
		diffswitch1port58,
		diffswitch1port59,
		diffswitch1port60,
		diffswitch1port61,
		diffswitch1port62,
		diffswitch1port63,
		diffswitch1port64,
		diffswitch1port65,
		diffswitch1port66,
		diffswitch1port67,
		diffswitch1port68,
		diffswitch1port69,
		diffswitch1port70,
		diffswitch1port71,
		diffswitch1port72,
		diffswitch1port73,
		diffswitch1port74,
		diffswitch1port75,
		diffswitch1port76,
		diffswitch1port77,
		diffswitch1port78,
		diffswitch1port79,
		diffswitch1port80,
		diffswitch1port81,
		diffswitch1port82,
		diffswitch1port83,
		diffswitch1port84,
		diffswitch1port85,
		diffswitch1port86,
		diffswitch1port87,
		diffswitch1port88,
		diffswitch1port89,
		diffswitch1port90,
		diffswitch1port91,
		diffswitch1port92,
		diffswitch1port93,
		diffswitch1port94,
		diffswitch1port95,
		diffswitch1port96,
		diffswitch1port97,
		diffswitch1port98,
		diffswitch1port99,
		diffswitch1port100,
		diffswitch1port101,
		diffswitch1port102,
		diffswitch1port103,
		diffswitch1port104,
		diffswitch1port105,
		diffswitch1port106,
		diffswitch1port107,
		diffswitch1port108,
		diffswitch1port109,
		diffswitch1port110,
		diffswitch1port111,
		diffswitch1port112,
		diffswitch1port113,
		diffswitch1port114,
		diffswitch1port115,
		diffswitch1port116,
		diffswitch1port117,
		diffswitch1port118,
		diffswitch1port119,
		diffswitch1port120,
		diffswitch1port121,
		diffswitch1port122,
		diffswitch1port123,
		diffswitch1port124,
		diffswitch1port125,
		diffswitch1port126,
		diffswitch1port127,
		diffswitch1port128,
		diffswitch1port129,
		diffswitch1port130,
		diffswitch1port131,
		diffswitch1port132,
		diffswitch1port133,
		diffswitch1port134,
		diffswitch1port135,
		diffswitch1port136,
		diffswitch1port137,
		diffswitch1port138,
		diffswitch1port139,
		diffswitch1port140,
		diffswitch1port141}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	diffswitch2port1   = &ConcretePort{Desc: "diffswitch2:port1", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port2   = &ConcretePort{Desc: "diffswitch2:port2", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port3   = &ConcretePort{Desc: "diffswitch2:port3", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port4   = &ConcretePort{Desc: "diffswitch2:port4", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port5   = &ConcretePort{Desc: "diffswitch2:port5", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port6   = &ConcretePort{Desc: "diffswitch2:port6", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port7   = &ConcretePort{Desc: "diffswitch2:port7", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port8   = &ConcretePort{Desc: "diffswitch2:port8", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port9   = &ConcretePort{Desc: "diffswitch2:port9", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port10  = &ConcretePort{Desc: "diffswitch2:port10", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port11  = &ConcretePort{Desc: "diffswitch2:port11", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port12  = &ConcretePort{Desc: "diffswitch2:port12", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port13  = &ConcretePort{Desc: "diffswitch2:port13", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port14  = &ConcretePort{Desc: "diffswitch2:port14", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port15  = &ConcretePort{Desc: "diffswitch2:port15", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port16  = &ConcretePort{Desc: "diffswitch2:port16", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port17  = &ConcretePort{Desc: "diffswitch2:port17", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port18  = &ConcretePort{Desc: "diffswitch2:port18", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port19  = &ConcretePort{Desc: "diffswitch2:port19", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port20  = &ConcretePort{Desc: "diffswitch2:port20", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port21  = &ConcretePort{Desc: "diffswitch2:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port22  = &ConcretePort{Desc: "diffswitch2:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port23  = &ConcretePort{Desc: "diffswitch2:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port24  = &ConcretePort{Desc: "diffswitch2:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port25  = &ConcretePort{Desc: "diffswitch2:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port26  = &ConcretePort{Desc: "diffswitch2:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port27  = &ConcretePort{Desc: "diffswitch2:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port28  = &ConcretePort{Desc: "diffswitch2:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port29  = &ConcretePort{Desc: "diffswitch2:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port30  = &ConcretePort{Desc: "diffswitch2:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port31  = &ConcretePort{Desc: "diffswitch2:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port32  = &ConcretePort{Desc: "diffswitch2:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port33  = &ConcretePort{Desc: "diffswitch2:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port34  = &ConcretePort{Desc: "diffswitch2:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port35  = &ConcretePort{Desc: "diffswitch2:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port36  = &ConcretePort{Desc: "diffswitch2:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port37  = &ConcretePort{Desc: "diffswitch2:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port38  = &ConcretePort{Desc: "diffswitch2:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port39  = &ConcretePort{Desc: "diffswitch2:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port40  = &ConcretePort{Desc: "diffswitch2:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port41  = &ConcretePort{Desc: "diffswitch2:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port42  = &ConcretePort{Desc: "diffswitch2:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port43  = &ConcretePort{Desc: "diffswitch2:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port44  = &ConcretePort{Desc: "diffswitch2:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port45  = &ConcretePort{Desc: "diffswitch2:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port46  = &ConcretePort{Desc: "diffswitch2:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port47  = &ConcretePort{Desc: "diffswitch2:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port48  = &ConcretePort{Desc: "diffswitch2:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port49  = &ConcretePort{Desc: "diffswitch2:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port50  = &ConcretePort{Desc: "diffswitch2:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port51  = &ConcretePort{Desc: "diffswitch2:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port52  = &ConcretePort{Desc: "diffswitch2:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port53  = &ConcretePort{Desc: "diffswitch2:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port54  = &ConcretePort{Desc: "diffswitch2:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port55  = &ConcretePort{Desc: "diffswitch2:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port56  = &ConcretePort{Desc: "diffswitch2:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port57  = &ConcretePort{Desc: "diffswitch2:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port58  = &ConcretePort{Desc: "diffswitch2:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port59  = &ConcretePort{Desc: "diffswitch2:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port60  = &ConcretePort{Desc: "diffswitch2:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port61  = &ConcretePort{Desc: "diffswitch2:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port62  = &ConcretePort{Desc: "diffswitch2:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port63  = &ConcretePort{Desc: "diffswitch2:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port64  = &ConcretePort{Desc: "diffswitch2:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port65  = &ConcretePort{Desc: "diffswitch2:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port66  = &ConcretePort{Desc: "diffswitch2:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port67  = &ConcretePort{Desc: "diffswitch2:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port68  = &ConcretePort{Desc: "diffswitch2:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port69  = &ConcretePort{Desc: "diffswitch2:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port70  = &ConcretePort{Desc: "diffswitch2:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port71  = &ConcretePort{Desc: "diffswitch2:port71", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port72  = &ConcretePort{Desc: "diffswitch2:port72", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port73  = &ConcretePort{Desc: "diffswitch2:port73", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port74  = &ConcretePort{Desc: "diffswitch2:port74", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port75  = &ConcretePort{Desc: "diffswitch2:port75", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port76  = &ConcretePort{Desc: "diffswitch2:port76", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port77  = &ConcretePort{Desc: "diffswitch2:port77", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port78  = &ConcretePort{Desc: "diffswitch2:port78", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port79  = &ConcretePort{Desc: "diffswitch2:port79", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port80  = &ConcretePort{Desc: "diffswitch2:port80", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port81  = &ConcretePort{Desc: "diffswitch2:port81", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port82  = &ConcretePort{Desc: "diffswitch2:port82", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port83  = &ConcretePort{Desc: "diffswitch2:port83", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port84  = &ConcretePort{Desc: "diffswitch2:port84", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port85  = &ConcretePort{Desc: "diffswitch2:port85", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port86  = &ConcretePort{Desc: "diffswitch2:port86", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port87  = &ConcretePort{Desc: "diffswitch2:port87", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port88  = &ConcretePort{Desc: "diffswitch2:port88", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port89  = &ConcretePort{Desc: "diffswitch2:port89", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port90  = &ConcretePort{Desc: "diffswitch2:port90", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port91  = &ConcretePort{Desc: "diffswitch2:port91", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port92  = &ConcretePort{Desc: "diffswitch2:port92", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port93  = &ConcretePort{Desc: "diffswitch2:port93", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port94  = &ConcretePort{Desc: "diffswitch2:port94", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port95  = &ConcretePort{Desc: "diffswitch2:port95", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port96  = &ConcretePort{Desc: "diffswitch2:port96", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port97  = &ConcretePort{Desc: "diffswitch2:port97", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port98  = &ConcretePort{Desc: "diffswitch2:port98", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port99  = &ConcretePort{Desc: "diffswitch2:port99", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port100 = &ConcretePort{Desc: "diffswitch2:port100", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port101 = &ConcretePort{Desc: "diffswitch2:port101", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port102 = &ConcretePort{Desc: "diffswitch2:port102", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port103 = &ConcretePort{Desc: "diffswitch2:port103", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port104 = &ConcretePort{Desc: "diffswitch2:port104", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port105 = &ConcretePort{Desc: "diffswitch2:port105", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port106 = &ConcretePort{Desc: "diffswitch2:port106", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port107 = &ConcretePort{Desc: "diffswitch2:port107", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port108 = &ConcretePort{Desc: "diffswitch2:port108", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port109 = &ConcretePort{Desc: "diffswitch2:port109", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port110 = &ConcretePort{Desc: "diffswitch2:port110", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port111 = &ConcretePort{Desc: "diffswitch2:port111", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port112 = &ConcretePort{Desc: "diffswitch2:port112", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port113 = &ConcretePort{Desc: "diffswitch2:port113", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port114 = &ConcretePort{Desc: "diffswitch2:port114", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port115 = &ConcretePort{Desc: "diffswitch2:port115", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port116 = &ConcretePort{Desc: "diffswitch2:port116", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port117 = &ConcretePort{Desc: "diffswitch2:port117", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port118 = &ConcretePort{Desc: "diffswitch2:port118", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port119 = &ConcretePort{Desc: "diffswitch2:port119", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port120 = &ConcretePort{Desc: "diffswitch2:port120", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port121 = &ConcretePort{Desc: "diffswitch2:port121", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port122 = &ConcretePort{Desc: "diffswitch2:port122", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port123 = &ConcretePort{Desc: "diffswitch2:port123", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port124 = &ConcretePort{Desc: "diffswitch2:port124", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port125 = &ConcretePort{Desc: "diffswitch2:port125", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port126 = &ConcretePort{Desc: "diffswitch2:port126", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port127 = &ConcretePort{Desc: "diffswitch2:port127", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port128 = &ConcretePort{Desc: "diffswitch2:port128", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port129 = &ConcretePort{Desc: "diffswitch2:port129", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port130 = &ConcretePort{Desc: "diffswitch2:port130", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port131 = &ConcretePort{Desc: "diffswitch2:port131", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port132 = &ConcretePort{Desc: "diffswitch2:port132", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port133 = &ConcretePort{Desc: "diffswitch2:port133", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port134 = &ConcretePort{Desc: "diffswitch2:port134", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port135 = &ConcretePort{Desc: "diffswitch2:port135", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port136 = &ConcretePort{Desc: "diffswitch2:port136", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port137 = &ConcretePort{Desc: "diffswitch2:port137", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port138 = &ConcretePort{Desc: "diffswitch2:port138", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port139 = &ConcretePort{Desc: "diffswitch2:port139", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port140 = &ConcretePort{Desc: "diffswitch2:port140", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port141 = &ConcretePort{Desc: "diffswitch2:port141", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch2port142 = &ConcretePort{Desc: "diffswitch2:port142", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}

	diffswitch2 = &ConcreteNode{Desc: "diffswitch2", Ports: []*ConcretePort{
		diffswitch2port1,
		diffswitch2port2,
		diffswitch2port3,
		diffswitch2port4,
		diffswitch2port5,
		diffswitch2port6,
		diffswitch2port7,
		diffswitch2port8,
		diffswitch2port9,
		diffswitch2port10,
		diffswitch2port11,
		diffswitch2port12,
		diffswitch2port13,
		diffswitch2port14,
		diffswitch2port15,
		diffswitch2port16,
		diffswitch2port17,
		diffswitch2port18,
		diffswitch2port19,
		diffswitch2port20,
		diffswitch2port21,
		diffswitch2port22,
		diffswitch2port23,
		diffswitch2port24,
		diffswitch2port25,
		diffswitch2port26,
		diffswitch2port27,
		diffswitch2port28,
		diffswitch2port29,
		diffswitch2port30,
		diffswitch2port31,
		diffswitch2port32,
		diffswitch2port33,
		diffswitch2port34,
		diffswitch2port35,
		diffswitch2port36,
		diffswitch2port37,
		diffswitch2port38,
		diffswitch2port39,
		diffswitch2port40,
		diffswitch2port41,
		diffswitch2port42,
		diffswitch2port43,
		diffswitch2port44,
		diffswitch2port45,
		diffswitch2port46,
		diffswitch2port47,
		diffswitch2port48,
		diffswitch2port49,
		diffswitch2port50,
		diffswitch2port51,
		diffswitch2port52,
		diffswitch2port53,
		diffswitch2port54,
		diffswitch2port55,
		diffswitch2port56,
		diffswitch2port57,
		diffswitch2port58,
		diffswitch2port59,
		diffswitch2port60,
		diffswitch2port61,
		diffswitch2port62,
		diffswitch2port63,
		diffswitch2port64,
		diffswitch2port65,
		diffswitch2port66,
		diffswitch2port67,
		diffswitch2port68,
		diffswitch2port69,
		diffswitch2port70,
		diffswitch2port71,
		diffswitch2port72,
		diffswitch2port73,
		diffswitch2port74,
		diffswitch2port75,
		diffswitch2port76,
		diffswitch2port77,
		diffswitch2port78,
		diffswitch2port79,
		diffswitch2port80,
		diffswitch2port81,
		diffswitch2port82,
		diffswitch2port83,
		diffswitch2port84,
		diffswitch2port85,
		diffswitch2port86,
		diffswitch2port87,
		diffswitch2port88,
		diffswitch2port89,
		diffswitch2port90,
		diffswitch2port91,
		diffswitch2port92,
		diffswitch2port93,
		diffswitch2port94,
		diffswitch2port95,
		diffswitch2port96,
		diffswitch2port97,
		diffswitch2port98,
		diffswitch2port99,
		diffswitch2port100,
		diffswitch2port101,
		diffswitch2port102,
		diffswitch2port103,
		diffswitch2port104,
		diffswitch2port105,
		diffswitch2port106,
		diffswitch2port107,
		diffswitch2port108,
		diffswitch2port109,
		diffswitch2port110,
		diffswitch2port111,
		diffswitch2port112,
		diffswitch2port113,
		diffswitch2port114,
		diffswitch2port115,
		diffswitch2port116,
		diffswitch2port117,
		diffswitch2port118,
		diffswitch2port119,
		diffswitch2port120,
		diffswitch2port121,
		diffswitch2port122,
		diffswitch2port123,
		diffswitch2port124,
		diffswitch2port125,
		diffswitch2port126,
		diffswitch2port127,
		diffswitch2port128,
		diffswitch2port129,
		diffswitch2port130,
		diffswitch2port131,
		diffswitch2port132,
		diffswitch2port133,
		diffswitch2port134,
		diffswitch2port135,
		diffswitch2port136,
		diffswitch2port137,
		diffswitch2port138,
		diffswitch2port139,
		diffswitch2port140,
		diffswitch2port141,
		diffswitch2port142}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	diffswitch3port1   = &ConcretePort{Desc: "diffswitch3:port1", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port2   = &ConcretePort{Desc: "diffswitch3:port2", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port3   = &ConcretePort{Desc: "diffswitch3:port3", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port4   = &ConcretePort{Desc: "diffswitch3:port4", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port5   = &ConcretePort{Desc: "diffswitch3:port5", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port6   = &ConcretePort{Desc: "diffswitch3:port6", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port7   = &ConcretePort{Desc: "diffswitch3:port7", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port8   = &ConcretePort{Desc: "diffswitch3:port8", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port9   = &ConcretePort{Desc: "diffswitch3:port9", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port10  = &ConcretePort{Desc: "diffswitch3:port10", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port11  = &ConcretePort{Desc: "diffswitch3:port11", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port12  = &ConcretePort{Desc: "diffswitch3:port12", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port13  = &ConcretePort{Desc: "diffswitch3:port13", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port14  = &ConcretePort{Desc: "diffswitch3:port14", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port15  = &ConcretePort{Desc: "diffswitch3:port15", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port16  = &ConcretePort{Desc: "diffswitch3:port16", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port17  = &ConcretePort{Desc: "diffswitch3:port17", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port18  = &ConcretePort{Desc: "diffswitch3:port18", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port19  = &ConcretePort{Desc: "diffswitch3:port19", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port20  = &ConcretePort{Desc: "diffswitch3:port20", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port21  = &ConcretePort{Desc: "diffswitch3:port21", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port22  = &ConcretePort{Desc: "diffswitch3:port22", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port23  = &ConcretePort{Desc: "diffswitch3:port23", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port24  = &ConcretePort{Desc: "diffswitch3:port24", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port25  = &ConcretePort{Desc: "diffswitch3:port25", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port26  = &ConcretePort{Desc: "diffswitch3:port26", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port27  = &ConcretePort{Desc: "diffswitch3:port27", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port28  = &ConcretePort{Desc: "diffswitch3:port28", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port29  = &ConcretePort{Desc: "diffswitch3:port29", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port30  = &ConcretePort{Desc: "diffswitch3:port30", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port31  = &ConcretePort{Desc: "diffswitch3:port31", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port32  = &ConcretePort{Desc: "diffswitch3:port32", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port33  = &ConcretePort{Desc: "diffswitch3:port33", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port34  = &ConcretePort{Desc: "diffswitch3:port34", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port35  = &ConcretePort{Desc: "diffswitch3:port35", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port36  = &ConcretePort{Desc: "diffswitch3:port36", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port37  = &ConcretePort{Desc: "diffswitch3:port37", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port38  = &ConcretePort{Desc: "diffswitch3:port38", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port39  = &ConcretePort{Desc: "diffswitch3:port39", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port40  = &ConcretePort{Desc: "diffswitch3:port40", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port41  = &ConcretePort{Desc: "diffswitch3:port41", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port42  = &ConcretePort{Desc: "diffswitch3:port42", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port43  = &ConcretePort{Desc: "diffswitch3:port43", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port44  = &ConcretePort{Desc: "diffswitch3:port44", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port45  = &ConcretePort{Desc: "diffswitch3:port45", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port46  = &ConcretePort{Desc: "diffswitch3:port46", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port47  = &ConcretePort{Desc: "diffswitch3:port47", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port48  = &ConcretePort{Desc: "diffswitch3:port48", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port49  = &ConcretePort{Desc: "diffswitch3:port49", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port50  = &ConcretePort{Desc: "diffswitch3:port50", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port51  = &ConcretePort{Desc: "diffswitch3:port51", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port52  = &ConcretePort{Desc: "diffswitch3:port52", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port53  = &ConcretePort{Desc: "diffswitch3:port53", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port54  = &ConcretePort{Desc: "diffswitch3:port54", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port55  = &ConcretePort{Desc: "diffswitch3:port55", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port56  = &ConcretePort{Desc: "diffswitch3:port56", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port57  = &ConcretePort{Desc: "diffswitch3:port57", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port58  = &ConcretePort{Desc: "diffswitch3:port58", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port59  = &ConcretePort{Desc: "diffswitch3:port59", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port60  = &ConcretePort{Desc: "diffswitch3:port60", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port61  = &ConcretePort{Desc: "diffswitch3:port61", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port62  = &ConcretePort{Desc: "diffswitch3:port62", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port63  = &ConcretePort{Desc: "diffswitch3:port63", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port64  = &ConcretePort{Desc: "diffswitch3:port64", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port65  = &ConcretePort{Desc: "diffswitch3:port65", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port66  = &ConcretePort{Desc: "diffswitch3:port66", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port67  = &ConcretePort{Desc: "diffswitch3:port67", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port68  = &ConcretePort{Desc: "diffswitch3:port68", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port69  = &ConcretePort{Desc: "diffswitch3:port69", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port70  = &ConcretePort{Desc: "diffswitch3:port70", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port71  = &ConcretePort{Desc: "diffswitch3:port71", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port72  = &ConcretePort{Desc: "diffswitch3:port72", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port73  = &ConcretePort{Desc: "diffswitch3:port73", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port74  = &ConcretePort{Desc: "diffswitch3:port74", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port75  = &ConcretePort{Desc: "diffswitch3:port75", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port76  = &ConcretePort{Desc: "diffswitch3:port76", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port77  = &ConcretePort{Desc: "diffswitch3:port77", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port78  = &ConcretePort{Desc: "diffswitch3:port78", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port79  = &ConcretePort{Desc: "diffswitch3:port79", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port80  = &ConcretePort{Desc: "diffswitch3:port80", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port81  = &ConcretePort{Desc: "diffswitch3:port81", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port82  = &ConcretePort{Desc: "diffswitch3:port82", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port83  = &ConcretePort{Desc: "diffswitch3:port83", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port84  = &ConcretePort{Desc: "diffswitch3:port84", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port85  = &ConcretePort{Desc: "diffswitch3:port85", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port86  = &ConcretePort{Desc: "diffswitch3:port86", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port87  = &ConcretePort{Desc: "diffswitch3:port87", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port88  = &ConcretePort{Desc: "diffswitch3:port88", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port89  = &ConcretePort{Desc: "diffswitch3:port89", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port90  = &ConcretePort{Desc: "diffswitch3:port90", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port91  = &ConcretePort{Desc: "diffswitch3:port91", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port92  = &ConcretePort{Desc: "diffswitch3:port92", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port93  = &ConcretePort{Desc: "diffswitch3:port93", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port94  = &ConcretePort{Desc: "diffswitch3:port94", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port95  = &ConcretePort{Desc: "diffswitch3:port95", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port96  = &ConcretePort{Desc: "diffswitch3:port96", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port97  = &ConcretePort{Desc: "diffswitch3:port97", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port98  = &ConcretePort{Desc: "diffswitch3:port98", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port99  = &ConcretePort{Desc: "diffswitch3:port99", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port100 = &ConcretePort{Desc: "diffswitch3:port100", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port101 = &ConcretePort{Desc: "diffswitch3:port101", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port102 = &ConcretePort{Desc: "diffswitch3:port102", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port103 = &ConcretePort{Desc: "diffswitch3:port103", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port104 = &ConcretePort{Desc: "diffswitch3:port104", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port105 = &ConcretePort{Desc: "diffswitch3:port105", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port106 = &ConcretePort{Desc: "diffswitch3:port106", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port107 = &ConcretePort{Desc: "diffswitch3:port107", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port108 = &ConcretePort{Desc: "diffswitch3:port108", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port109 = &ConcretePort{Desc: "diffswitch3:port109", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port110 = &ConcretePort{Desc: "diffswitch3:port110", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port111 = &ConcretePort{Desc: "diffswitch3:port111", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port112 = &ConcretePort{Desc: "diffswitch3:port112", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port113 = &ConcretePort{Desc: "diffswitch3:port113", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port114 = &ConcretePort{Desc: "diffswitch3:port114", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port115 = &ConcretePort{Desc: "diffswitch3:port115", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port116 = &ConcretePort{Desc: "diffswitch3:port116", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port117 = &ConcretePort{Desc: "diffswitch3:port117", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port118 = &ConcretePort{Desc: "diffswitch3:port118", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port119 = &ConcretePort{Desc: "diffswitch3:port119", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port120 = &ConcretePort{Desc: "diffswitch3:port120", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port121 = &ConcretePort{Desc: "diffswitch3:port121", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port122 = &ConcretePort{Desc: "diffswitch3:port122", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port123 = &ConcretePort{Desc: "diffswitch3:port123", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port124 = &ConcretePort{Desc: "diffswitch3:port124", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port125 = &ConcretePort{Desc: "diffswitch3:port125", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port126 = &ConcretePort{Desc: "diffswitch3:port126", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port127 = &ConcretePort{Desc: "diffswitch3:port127", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port128 = &ConcretePort{Desc: "diffswitch3:port128", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port129 = &ConcretePort{Desc: "diffswitch3:port129", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port130 = &ConcretePort{Desc: "diffswitch3:port130", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port131 = &ConcretePort{Desc: "diffswitch3:port131", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port132 = &ConcretePort{Desc: "diffswitch3:port132", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port133 = &ConcretePort{Desc: "diffswitch3:port133", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port134 = &ConcretePort{Desc: "diffswitch3:port134", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port135 = &ConcretePort{Desc: "diffswitch3:port135", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port136 = &ConcretePort{Desc: "diffswitch3:port136", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port137 = &ConcretePort{Desc: "diffswitch3:port137", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port138 = &ConcretePort{Desc: "diffswitch3:port138", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port139 = &ConcretePort{Desc: "diffswitch3:port139", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port140 = &ConcretePort{Desc: "diffswitch3:port140", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}
	diffswitch3port141 = &ConcretePort{Desc: "diffswitch3:port141", Attrs: map[string]string{"speed": "S_400G", "media": "copper"}}

	diffswitch3 = &ConcreteNode{Desc: "diffswitch3", Ports: []*ConcretePort{
		diffswitch3port1,
		diffswitch3port2,
		diffswitch3port3,
		diffswitch3port4,
		diffswitch3port5,
		diffswitch3port6,
		diffswitch3port7,
		diffswitch3port8,
		diffswitch3port9,
		diffswitch3port10,
		diffswitch3port11,
		diffswitch3port12,
		diffswitch3port13,
		diffswitch3port14,
		diffswitch3port15,
		diffswitch3port16,
		diffswitch3port17,
		diffswitch3port18,
		diffswitch3port19,
		diffswitch3port20,
		diffswitch3port21,
		diffswitch3port22,
		diffswitch3port23,
		diffswitch3port24,
		diffswitch3port25,
		diffswitch3port26,
		diffswitch3port27,
		diffswitch3port28,
		diffswitch3port29,
		diffswitch3port30,
		diffswitch3port31,
		diffswitch3port32,
		diffswitch3port33,
		diffswitch3port34,
		diffswitch3port35,
		diffswitch3port36,
		diffswitch3port37,
		diffswitch3port38,
		diffswitch3port39,
		diffswitch3port40,
		diffswitch3port41,
		diffswitch3port42,
		diffswitch3port43,
		diffswitch3port44,
		diffswitch3port45,
		diffswitch3port46,
		diffswitch3port47,
		diffswitch3port48,
		diffswitch3port49,
		diffswitch3port50,
		diffswitch3port51,
		diffswitch3port52,
		diffswitch3port53,
		diffswitch3port54,
		diffswitch3port55,
		diffswitch3port56,
		diffswitch3port57,
		diffswitch3port58,
		diffswitch3port59,
		diffswitch3port60,
		diffswitch3port61,
		diffswitch3port62,
		diffswitch3port63,
		diffswitch3port64,
		diffswitch3port65,
		diffswitch3port66,
		diffswitch3port67,
		diffswitch3port68,
		diffswitch3port69,
		diffswitch3port70,
		diffswitch3port71,
		diffswitch3port72,
		diffswitch3port73,
		diffswitch3port74,
		diffswitch3port75,
		diffswitch3port76,
		diffswitch3port77,
		diffswitch3port78,
		diffswitch3port79,
		diffswitch3port80,
		diffswitch3port81,
		diffswitch3port82,
		diffswitch3port83,
		diffswitch3port84,
		diffswitch3port85,
		diffswitch3port86,
		diffswitch3port87,
		diffswitch3port88,
		diffswitch3port89,
		diffswitch3port90,
		diffswitch3port91,
		diffswitch3port92,
		diffswitch3port93,
		diffswitch3port94,
		diffswitch3port95,
		diffswitch3port96,
		diffswitch3port97,
		diffswitch3port98,
		diffswitch3port99,
		diffswitch3port100,
		diffswitch3port101,
		diffswitch3port102,
		diffswitch3port103,
		diffswitch3port104,
		diffswitch3port105,
		diffswitch3port106,
		diffswitch3port107,
		diffswitch3port108,
		diffswitch3port109,
		diffswitch3port110,
		diffswitch3port111,
		diffswitch3port112,
		diffswitch3port113,
		diffswitch3port114,
		diffswitch3port115,
		diffswitch3port116,
		diffswitch3port117,
		diffswitch3port118,
		diffswitch3port119,
		diffswitch3port120,
		diffswitch3port121,
		diffswitch3port122,
		diffswitch3port123,
		diffswitch3port124,
		diffswitch3port125,
		diffswitch3port126,
		diffswitch3port127,
		diffswitch3port128,
		diffswitch3port129,
		diffswitch3port130,
		diffswitch3port131,
		diffswitch3port132,
		diffswitch3port133,
		diffswitch3port134,
		diffswitch3port135,
		diffswitch3port136,
		diffswitch3port137,
		diffswitch3port138,
		diffswitch3port139,
		diffswitch3port140,
		diffswitch3port141}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}
)

var superGraphTestDiff = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		diffdut1,
		diffate1,
		diffdut2,
		diffate2,
		diffdut3,
		diffate3,
		diffdut4,
		diffate4,
		diffswitch1, diffswitch2, diffswitch3},
	Edges: []*ConcreteEdge{
		{Src: diffdut1port1, Dst: diffate1port1},
		{Src: diffdut1port2, Dst: diffate1port2},
		{Src: diffdut1port3, Dst: diffate1port3},
		{Src: diffdut1port4, Dst: diffate1port4},
		{Src: diffdut1port5, Dst: diffate1port5},
		{Src: diffdut1port6, Dst: diffate1port6},
		{Src: diffdut1port7, Dst: diffate1port7},
		{Src: diffdut1port8, Dst: diffate1port8},
		{Src: diffdut1port9, Dst: diffate1port9},
		{Src: diffdut1port10, Dst: diffate1port10},
		{Src: diffdut1port11, Dst: diffate1port11},
		{Src: diffdut1port12, Dst: diffate1port12},
		{Src: diffdut1port13, Dst: diffate1port13},
		{Src: diffdut1port14, Dst: diffate1port14},
		{Src: diffdut1port15, Dst: diffate1port15},
		{Src: diffdut1port16, Dst: diffate1port16},
		{Src: diffdut1port17, Dst: diffate1port17},
		{Src: diffdut1port18, Dst: diffate1port18},
		{Src: diffdut1port19, Dst: diffate1port19},
		{Src: diffdut1port20, Dst: diffate1port20},
		{Src: diffdut1port21, Dst: diffate1port21},
		{Src: diffdut1port22, Dst: diffate1port22},
		{Src: diffdut1port23, Dst: diffate1port23},
		{Src: diffdut1port24, Dst: diffate1port24},
		{Src: diffdut1port25, Dst: diffate1port25},
		{Src: diffdut1port26, Dst: diffate1port26},
		{Src: diffdut1port27, Dst: diffate1port27},
		{Src: diffdut1port28, Dst: diffate1port28},
		{Src: diffdut1port29, Dst: diffate1port29},
		{Src: diffdut1port30, Dst: diffate1port30},
		{Src: diffdut1port31, Dst: diffate1port31},
		{Src: diffdut1port32, Dst: diffate1port32},
		{Src: diffdut1port33, Dst: diffate1port33},
		{Src: diffdut1port34, Dst: diffate1port34},
		{Src: diffdut1port35, Dst: diffate1port35},

		{Src: diffdut1port36, Dst: diffswitch1port1},
		{Src: diffdut1port37, Dst: diffswitch1port2},
		{Src: diffdut1port38, Dst: diffswitch1port3},
		{Src: diffdut1port39, Dst: diffswitch1port4},
		{Src: diffdut1port40, Dst: diffswitch1port5},
		{Src: diffdut1port41, Dst: diffswitch1port6},
		{Src: diffdut1port42, Dst: diffswitch1port7},
		{Src: diffdut1port43, Dst: diffswitch1port8},
		{Src: diffdut1port44, Dst: diffswitch1port9},
		{Src: diffdut1port45, Dst: diffswitch1port10},
		{Src: diffdut1port46, Dst: diffswitch1port11},
		{Src: diffdut1port47, Dst: diffswitch1port12},
		{Src: diffdut1port48, Dst: diffswitch1port13},
		{Src: diffdut1port49, Dst: diffswitch1port14},
		{Src: diffdut1port50, Dst: diffswitch1port15},
		{Src: diffdut1port51, Dst: diffswitch1port16},
		{Src: diffdut1port52, Dst: diffswitch1port17},
		{Src: diffdut1port53, Dst: diffswitch1port18},
		{Src: diffdut1port54, Dst: diffswitch1port19},
		{Src: diffdut1port55, Dst: diffswitch1port20},
		{Src: diffdut1port56, Dst: diffswitch1port21},
		{Src: diffdut1port57, Dst: diffswitch1port22},
		{Src: diffdut1port58, Dst: diffswitch1port23},
		{Src: diffdut1port59, Dst: diffswitch1port24},
		{Src: diffdut1port60, Dst: diffswitch1port25},
		{Src: diffdut1port61, Dst: diffswitch1port26},
		{Src: diffdut1port62, Dst: diffswitch1port27},
		{Src: diffdut1port63, Dst: diffswitch1port28},
		{Src: diffdut1port64, Dst: diffswitch1port29},
		{Src: diffdut1port65, Dst: diffswitch1port30},
		{Src: diffdut1port66, Dst: diffswitch1port31},
		{Src: diffdut1port67, Dst: diffswitch1port32},
		{Src: diffdut1port68, Dst: diffswitch1port33},
		{Src: diffdut1port69, Dst: diffswitch1port34},
		{Src: diffdut1port70, Dst: diffswitch1port35},

		{Src: diffswitch1port36, Dst: diffate1port36},
		{Src: diffswitch1port37, Dst: diffate1port37},
		{Src: diffswitch1port38, Dst: diffate1port38},
		{Src: diffswitch1port39, Dst: diffate1port39},
		{Src: diffswitch1port40, Dst: diffate1port40},
		{Src: diffswitch1port41, Dst: diffate1port41},
		{Src: diffswitch1port42, Dst: diffate1port42},
		{Src: diffswitch1port43, Dst: diffate1port43},
		{Src: diffswitch1port44, Dst: diffate1port44},
		{Src: diffswitch1port45, Dst: diffate1port45},
		{Src: diffswitch1port46, Dst: diffate1port46},
		{Src: diffswitch1port47, Dst: diffate1port47},
		{Src: diffswitch1port48, Dst: diffate1port48},
		{Src: diffswitch1port49, Dst: diffate1port49},
		{Src: diffswitch1port50, Dst: diffate1port50},
		{Src: diffswitch1port51, Dst: diffate1port51},
		{Src: diffswitch1port52, Dst: diffate1port52},
		{Src: diffswitch1port53, Dst: diffate1port53},
		{Src: diffswitch1port54, Dst: diffate1port54},
		{Src: diffswitch1port55, Dst: diffate1port55},
		{Src: diffswitch1port56, Dst: diffate1port56},
		{Src: diffswitch1port57, Dst: diffate1port57},
		{Src: diffswitch1port58, Dst: diffate1port58},
		{Src: diffswitch1port59, Dst: diffate1port59},
		{Src: diffswitch1port60, Dst: diffate1port60},
		{Src: diffswitch1port61, Dst: diffate1port61},
		{Src: diffswitch1port62, Dst: diffate1port62},
		{Src: diffswitch1port63, Dst: diffate1port63},
		{Src: diffswitch1port64, Dst: diffate1port64},
		{Src: diffswitch1port65, Dst: diffate1port65},
		{Src: diffswitch1port66, Dst: diffate1port66},
		{Src: diffswitch1port67, Dst: diffate1port67},
		{Src: diffswitch1port68, Dst: diffate1port68},
		{Src: diffswitch1port69, Dst: diffate1port69},
		{Src: diffswitch1port70, Dst: diffate1port70},
		{Src: diffswitch1port141, Dst: diffswitch2port141},

		{Src: diffdut2port1, Dst: diffate2port1},
		{Src: diffdut2port2, Dst: diffate2port2},
		{Src: diffdut2port3, Dst: diffate2port3},
		{Src: diffdut2port4, Dst: diffate2port4},
		{Src: diffdut2port5, Dst: diffate2port5},
		{Src: diffdut2port6, Dst: diffate2port6},
		{Src: diffdut2port7, Dst: diffate2port7},
		{Src: diffdut2port8, Dst: diffate2port8},
		{Src: diffdut2port9, Dst: diffate2port9},
		{Src: diffdut2port10, Dst: diffate2port10},
		{Src: diffdut2port11, Dst: diffate2port11},
		{Src: diffdut2port12, Dst: diffate2port12},
		{Src: diffdut2port13, Dst: diffate2port13},
		{Src: diffdut2port14, Dst: diffate2port14},
		{Src: diffdut2port15, Dst: diffate2port15},
		{Src: diffdut2port16, Dst: diffate2port16},
		{Src: diffdut2port17, Dst: diffate2port17},
		{Src: diffdut2port18, Dst: diffate2port18},
		{Src: diffdut2port19, Dst: diffate2port19},
		{Src: diffdut2port20, Dst: diffate2port20},
		{Src: diffdut2port21, Dst: diffate2port21},
		{Src: diffdut2port22, Dst: diffate2port22},
		{Src: diffdut2port23, Dst: diffate2port23},
		{Src: diffdut2port24, Dst: diffate2port24},
		{Src: diffdut2port25, Dst: diffate2port25},
		{Src: diffdut2port26, Dst: diffate2port26},
		{Src: diffdut2port27, Dst: diffate2port27},
		{Src: diffdut2port28, Dst: diffate2port28},
		{Src: diffdut2port29, Dst: diffate2port29},
		{Src: diffdut2port30, Dst: diffate2port30},
		{Src: diffdut2port31, Dst: diffate2port31},
		{Src: diffdut2port32, Dst: diffate2port32},
		{Src: diffdut2port33, Dst: diffate2port33},
		{Src: diffdut2port34, Dst: diffate2port34},
		{Src: diffdut2port35, Dst: diffate2port35},

		{Src: diffdut2port36, Dst: diffswitch1port71},
		{Src: diffdut2port37, Dst: diffswitch1port72},
		{Src: diffdut2port38, Dst: diffswitch1port73},
		{Src: diffdut2port39, Dst: diffswitch1port74},
		{Src: diffdut2port40, Dst: diffswitch1port75},
		{Src: diffdut2port41, Dst: diffswitch1port76},
		{Src: diffdut2port42, Dst: diffswitch1port77},
		{Src: diffdut2port43, Dst: diffswitch1port78},
		{Src: diffdut2port44, Dst: diffswitch1port79},
		{Src: diffdut2port45, Dst: diffswitch1port80},
		{Src: diffdut2port46, Dst: diffswitch1port81},
		{Src: diffdut2port47, Dst: diffswitch1port82},
		{Src: diffdut2port48, Dst: diffswitch1port83},
		{Src: diffdut2port49, Dst: diffswitch1port84},
		{Src: diffdut2port50, Dst: diffswitch1port85},
		{Src: diffdut2port51, Dst: diffswitch1port86},
		{Src: diffdut2port52, Dst: diffswitch1port87},
		{Src: diffdut2port53, Dst: diffswitch1port88},
		{Src: diffdut2port54, Dst: diffswitch1port89},
		{Src: diffdut2port55, Dst: diffswitch1port90},
		{Src: diffdut2port56, Dst: diffswitch1port91},
		{Src: diffdut2port57, Dst: diffswitch1port92},
		{Src: diffdut2port58, Dst: diffswitch1port93},
		{Src: diffdut2port59, Dst: diffswitch1port94},
		{Src: diffdut2port60, Dst: diffswitch1port95},
		{Src: diffdut2port61, Dst: diffswitch1port96},
		{Src: diffdut2port62, Dst: diffswitch1port97},
		{Src: diffdut2port63, Dst: diffswitch1port98},
		{Src: diffdut2port64, Dst: diffswitch1port99},
		{Src: diffdut2port65, Dst: diffswitch1port100},
		{Src: diffdut2port66, Dst: diffswitch1port101},
		{Src: diffdut2port67, Dst: diffswitch1port102},
		{Src: diffdut2port68, Dst: diffswitch1port103},
		{Src: diffdut2port69, Dst: diffswitch1port104},
		{Src: diffdut2port70, Dst: diffswitch1port105},

		{Src: diffswitch1port106, Dst: diffate2port36},
		{Src: diffswitch1port107, Dst: diffate2port37},
		{Src: diffswitch1port108, Dst: diffate2port38},
		{Src: diffswitch1port109, Dst: diffate2port39},
		{Src: diffswitch1port110, Dst: diffate2port40},
		{Src: diffswitch1port111, Dst: diffate2port41},
		{Src: diffswitch1port112, Dst: diffate2port42},
		{Src: diffswitch1port113, Dst: diffate2port43},
		{Src: diffswitch1port114, Dst: diffate2port44},
		{Src: diffswitch1port115, Dst: diffate2port45},
		{Src: diffswitch1port116, Dst: diffate2port46},
		{Src: diffswitch1port117, Dst: diffate2port47},
		{Src: diffswitch1port118, Dst: diffate2port48},
		{Src: diffswitch1port119, Dst: diffate2port49},
		{Src: diffswitch1port120, Dst: diffate2port50},
		{Src: diffswitch1port121, Dst: diffate2port51},
		{Src: diffswitch1port122, Dst: diffate2port52},
		{Src: diffswitch1port123, Dst: diffate2port53},
		{Src: diffswitch1port124, Dst: diffate2port54},
		{Src: diffswitch1port125, Dst: diffate2port55},
		{Src: diffswitch1port126, Dst: diffate2port56},
		{Src: diffswitch1port127, Dst: diffate2port57},
		{Src: diffswitch1port128, Dst: diffate2port58},
		{Src: diffswitch1port129, Dst: diffate2port59},
		{Src: diffswitch1port130, Dst: diffate2port60},
		{Src: diffswitch1port131, Dst: diffate2port61},
		{Src: diffswitch1port132, Dst: diffate2port62},
		{Src: diffswitch1port133, Dst: diffate2port63},
		{Src: diffswitch1port134, Dst: diffate2port64},
		{Src: diffswitch1port135, Dst: diffate2port65},
		{Src: diffswitch1port136, Dst: diffate2port66},
		{Src: diffswitch1port137, Dst: diffate2port67},
		{Src: diffswitch1port138, Dst: diffate2port68},
		{Src: diffswitch1port139, Dst: diffate2port69},
		{Src: diffswitch1port140, Dst: diffate2port70},

		{Src: diffdut3port1, Dst: diffswitch2port1},
		{Src: diffdut3port2, Dst: diffswitch2port2},
		{Src: diffdut3port3, Dst: diffswitch2port3},
		{Src: diffdut3port4, Dst: diffswitch2port4},
		{Src: diffdut3port5, Dst: diffswitch2port5},
		{Src: diffdut3port6, Dst: diffswitch2port6},
		{Src: diffdut3port7, Dst: diffswitch2port7},
		{Src: diffdut3port8, Dst: diffswitch2port8},
		{Src: diffdut3port9, Dst: diffswitch2port9},
		{Src: diffdut3port10, Dst: diffswitch2port10},
		{Src: diffdut3port11, Dst: diffswitch2port11},
		{Src: diffdut3port12, Dst: diffswitch2port12},
		{Src: diffdut3port13, Dst: diffswitch2port13},
		{Src: diffdut3port14, Dst: diffswitch2port14},
		{Src: diffdut3port15, Dst: diffswitch2port15},
		{Src: diffdut3port16, Dst: diffswitch2port16},
		{Src: diffdut3port17, Dst: diffswitch2port17},
		{Src: diffdut3port18, Dst: diffswitch2port18},
		{Src: diffdut3port19, Dst: diffswitch2port19},
		{Src: diffdut3port20, Dst: diffswitch2port20},
		{Src: diffdut3port21, Dst: diffswitch2port21},
		{Src: diffdut3port22, Dst: diffswitch2port22},
		{Src: diffdut3port23, Dst: diffswitch2port23},
		{Src: diffdut3port24, Dst: diffswitch2port24},
		{Src: diffdut3port25, Dst: diffswitch2port25},
		{Src: diffdut3port26, Dst: diffswitch2port26},
		{Src: diffdut3port27, Dst: diffswitch2port27},
		{Src: diffdut3port28, Dst: diffswitch2port28},
		{Src: diffdut3port29, Dst: diffswitch2port29},
		{Src: diffdut3port30, Dst: diffswitch2port30},
		{Src: diffdut3port31, Dst: diffswitch2port31},
		{Src: diffdut3port32, Dst: diffswitch2port32},
		{Src: diffdut3port33, Dst: diffswitch2port33},
		{Src: diffdut3port34, Dst: diffswitch2port34},
		{Src: diffdut3port35, Dst: diffswitch2port35},

		{Src: diffswitch2port41, Dst: diffate3port1},
		{Src: diffswitch2port42, Dst: diffate3port2},
		{Src: diffswitch2port43, Dst: diffate3port3},
		{Src: diffswitch2port44, Dst: diffate3port4},
		{Src: diffswitch2port45, Dst: diffate3port5},
		{Src: diffswitch2port46, Dst: diffate3port6},
		{Src: diffswitch2port47, Dst: diffate3port7},
		{Src: diffswitch2port48, Dst: diffate3port8},
		{Src: diffswitch2port49, Dst: diffate3port9},
		{Src: diffswitch2port50, Dst: diffate3port10},
		{Src: diffswitch2port51, Dst: diffate3port11},
		{Src: diffswitch2port52, Dst: diffate3port12},
		{Src: diffswitch2port53, Dst: diffate3port13},
		{Src: diffswitch2port54, Dst: diffate3port14},
		{Src: diffswitch2port55, Dst: diffate3port15},
		{Src: diffswitch2port56, Dst: diffate3port16},
		{Src: diffswitch2port57, Dst: diffate3port17},
		{Src: diffswitch2port58, Dst: diffate3port18},
		{Src: diffswitch2port59, Dst: diffate3port19},
		{Src: diffswitch2port60, Dst: diffate3port20},
		{Src: diffswitch2port61, Dst: diffate3port21},
		{Src: diffswitch2port62, Dst: diffate3port22},
		{Src: diffswitch2port63, Dst: diffate3port23},
		{Src: diffswitch2port64, Dst: diffate3port24},
		{Src: diffswitch2port65, Dst: diffate3port25},
		{Src: diffswitch2port66, Dst: diffate3port26},
		{Src: diffswitch2port67, Dst: diffate3port27},
		{Src: diffswitch2port68, Dst: diffate3port28},
		{Src: diffswitch2port69, Dst: diffate3port29},
		{Src: diffswitch2port70, Dst: diffate3port30},
		{Src: diffswitch2port71, Dst: diffate3port31},
		{Src: diffswitch2port72, Dst: diffate3port32},
		{Src: diffswitch2port73, Dst: diffate3port33},
		{Src: diffswitch2port74, Dst: diffate3port34},
		{Src: diffswitch2port75, Dst: diffate3port35},

		{Src: diffdut4port1, Dst: diffswitch2port76},
		{Src: diffdut4port2, Dst: diffswitch2port77},
		{Src: diffdut4port3, Dst: diffswitch2port78},
		{Src: diffdut4port4, Dst: diffswitch2port79},
		{Src: diffdut4port5, Dst: diffswitch2port80},
		{Src: diffdut4port6, Dst: diffswitch2port81},
		{Src: diffdut4port7, Dst: diffswitch2port82},
		{Src: diffdut4port8, Dst: diffswitch2port83},
		{Src: diffdut4port9, Dst: diffswitch2port84},
		{Src: diffdut4port10, Dst: diffswitch2port85},
		{Src: diffdut4port11, Dst: diffswitch2port86},
		{Src: diffdut4port12, Dst: diffswitch2port87},
		{Src: diffdut4port13, Dst: diffswitch2port88},
		{Src: diffdut4port14, Dst: diffswitch2port89},
		{Src: diffdut4port15, Dst: diffswitch2port90},
		{Src: diffdut4port16, Dst: diffswitch2port91},
		{Src: diffdut4port17, Dst: diffswitch2port92},
		{Src: diffdut4port18, Dst: diffswitch2port93},
		{Src: diffdut4port19, Dst: diffswitch2port94},
		{Src: diffdut4port20, Dst: diffswitch2port95},
		{Src: diffdut4port21, Dst: diffswitch2port96},
		{Src: diffdut4port22, Dst: diffswitch2port97},
		{Src: diffdut4port23, Dst: diffswitch2port98},
		{Src: diffdut4port24, Dst: diffswitch2port99},
		{Src: diffdut4port25, Dst: diffswitch2port100},
		{Src: diffdut4port26, Dst: diffswitch2port101},
		{Src: diffdut4port27, Dst: diffswitch2port102},
		{Src: diffdut4port28, Dst: diffswitch2port103},
		{Src: diffdut4port29, Dst: diffswitch2port104},
		{Src: diffdut4port30, Dst: diffswitch2port105},
		{Src: diffdut4port31, Dst: diffswitch2port106},
		{Src: diffdut4port32, Dst: diffswitch2port107},
		{Src: diffdut4port33, Dst: diffswitch2port108},
		{Src: diffdut4port34, Dst: diffswitch2port109},
		{Src: diffdut4port35, Dst: diffswitch2port110},

		{Src: diffswitch2port111, Dst: diffate4port1},
		{Src: diffswitch2port112, Dst: diffate4port2},
		{Src: diffswitch2port113, Dst: diffate4port3},
		{Src: diffswitch2port114, Dst: diffate4port4},
		{Src: diffswitch2port115, Dst: diffate4port5},
		{Src: diffswitch2port116, Dst: diffate4port6},
		{Src: diffswitch2port117, Dst: diffate4port7},
		{Src: diffswitch2port118, Dst: diffate4port8},
		{Src: diffswitch2port119, Dst: diffate4port9},
		{Src: diffswitch2port120, Dst: diffate4port10},
		{Src: diffswitch2port121, Dst: diffate4port11},
		{Src: diffswitch2port122, Dst: diffate4port12},
		{Src: diffswitch2port123, Dst: diffate4port13},
		{Src: diffswitch2port124, Dst: diffate4port14},
		{Src: diffswitch2port125, Dst: diffate4port15},
		{Src: diffswitch2port126, Dst: diffate4port16},
		{Src: diffswitch2port127, Dst: diffate4port17},
		{Src: diffswitch2port128, Dst: diffate4port18},
		{Src: diffswitch2port129, Dst: diffate4port19},
		{Src: diffswitch2port130, Dst: diffate4port20},
		{Src: diffswitch2port131, Dst: diffate4port21},
		{Src: diffswitch2port132, Dst: diffate4port22},
		{Src: diffswitch2port133, Dst: diffate4port23},
		{Src: diffswitch2port134, Dst: diffate4port24},
		{Src: diffswitch2port135, Dst: diffate4port25},
		{Src: diffswitch2port136, Dst: diffate4port26},
		{Src: diffswitch2port137, Dst: diffate4port27},
		{Src: diffswitch2port138, Dst: diffate4port28},
		{Src: diffswitch2port139, Dst: diffate4port29},
		{Src: diffswitch2port140, Dst: diffate4port30},
		{Src: diffswitch2port36, Dst: diffate4port31},
		{Src: diffswitch2port37, Dst: diffate4port32},
		{Src: diffswitch2port38, Dst: diffate4port33},
		{Src: diffswitch2port39, Dst: diffate4port34},
		{Src: diffswitch2port40, Dst: diffate4port35},
		{Src: diffswitch2port142, Dst: diffswitch3port141},

		{Src: diffdut3port36, Dst: diffswitch3port1},
		{Src: diffdut3port37, Dst: diffswitch3port2},
		{Src: diffdut3port38, Dst: diffswitch3port3},
		{Src: diffdut3port39, Dst: diffswitch3port4},
		{Src: diffdut3port40, Dst: diffswitch3port5},
		{Src: diffdut3port41, Dst: diffswitch3port6},
		{Src: diffdut3port42, Dst: diffswitch3port7},
		{Src: diffdut3port43, Dst: diffswitch3port8},
		{Src: diffdut3port44, Dst: diffswitch3port9},
		{Src: diffdut3port45, Dst: diffswitch3port10},
		{Src: diffdut3port46, Dst: diffswitch3port11},
		{Src: diffdut3port47, Dst: diffswitch3port12},
		{Src: diffdut3port48, Dst: diffswitch3port13},
		{Src: diffdut3port49, Dst: diffswitch3port14},
		{Src: diffdut3port50, Dst: diffswitch3port15},
		{Src: diffdut3port51, Dst: diffswitch3port16},
		{Src: diffdut3port52, Dst: diffswitch3port17},
		{Src: diffdut3port53, Dst: diffswitch3port18},
		{Src: diffdut3port54, Dst: diffswitch3port19},
		{Src: diffdut3port55, Dst: diffswitch3port20},
		{Src: diffdut3port56, Dst: diffswitch3port21},
		{Src: diffdut3port57, Dst: diffswitch3port22},
		{Src: diffdut3port58, Dst: diffswitch3port23},
		{Src: diffdut3port59, Dst: diffswitch3port24},
		{Src: diffdut3port60, Dst: diffswitch3port25},
		{Src: diffdut3port61, Dst: diffswitch3port26},
		{Src: diffdut3port62, Dst: diffswitch3port27},
		{Src: diffdut3port63, Dst: diffswitch3port28},
		{Src: diffdut3port64, Dst: diffswitch3port29},
		{Src: diffdut3port65, Dst: diffswitch3port30},
		{Src: diffdut3port66, Dst: diffswitch3port31},
		{Src: diffdut3port67, Dst: diffswitch3port32},
		{Src: diffdut3port68, Dst: diffswitch3port33},
		{Src: diffdut3port69, Dst: diffswitch3port34},
		{Src: diffdut3port70, Dst: diffswitch3port35},

		{Src: diffswitch3port36, Dst: diffate3port36},
		{Src: diffswitch3port37, Dst: diffate3port37},
		{Src: diffswitch3port38, Dst: diffate3port38},
		{Src: diffswitch3port39, Dst: diffate3port39},
		{Src: diffswitch3port40, Dst: diffate3port40},
		{Src: diffswitch3port41, Dst: diffate3port41},
		{Src: diffswitch3port42, Dst: diffate3port42},
		{Src: diffswitch3port43, Dst: diffate3port43},
		{Src: diffswitch3port44, Dst: diffate3port44},
		{Src: diffswitch3port45, Dst: diffate3port45},
		{Src: diffswitch3port46, Dst: diffate3port46},
		{Src: diffswitch3port47, Dst: diffate3port47},
		{Src: diffswitch3port48, Dst: diffate3port48},
		{Src: diffswitch3port49, Dst: diffate3port49},
		{Src: diffswitch3port50, Dst: diffate3port50},
		{Src: diffswitch3port51, Dst: diffate3port51},
		{Src: diffswitch3port52, Dst: diffate3port52},
		{Src: diffswitch3port53, Dst: diffate3port53},
		{Src: diffswitch3port54, Dst: diffate3port54},
		{Src: diffswitch3port55, Dst: diffate3port55},
		{Src: diffswitch3port56, Dst: diffate3port56},
		{Src: diffswitch3port57, Dst: diffate3port57},
		{Src: diffswitch3port58, Dst: diffate3port58},
		{Src: diffswitch3port59, Dst: diffate3port59},
		{Src: diffswitch3port60, Dst: diffate3port60},
		{Src: diffswitch3port61, Dst: diffate3port61},
		{Src: diffswitch3port62, Dst: diffate3port62},
		{Src: diffswitch3port63, Dst: diffate3port63},
		{Src: diffswitch3port64, Dst: diffate3port64},
		{Src: diffswitch3port65, Dst: diffate3port65},
		{Src: diffswitch3port66, Dst: diffate3port66},
		{Src: diffswitch3port67, Dst: diffate3port67},
		{Src: diffswitch3port68, Dst: diffate3port68},
		{Src: diffswitch3port69, Dst: diffate3port69},
		{Src: diffswitch3port70, Dst: diffate3port70},

		{Src: diffdut4port36, Dst: diffswitch3port71},
		{Src: diffdut4port37, Dst: diffswitch3port72},
		{Src: diffdut4port38, Dst: diffswitch3port73},
		{Src: diffdut4port39, Dst: diffswitch3port74},
		{Src: diffdut4port40, Dst: diffswitch3port75},
		{Src: diffdut4port41, Dst: diffswitch3port76},
		{Src: diffdut4port42, Dst: diffswitch3port77},
		{Src: diffdut4port43, Dst: diffswitch3port78},
		{Src: diffdut4port44, Dst: diffswitch3port79},
		{Src: diffdut4port45, Dst: diffswitch3port80},
		{Src: diffdut4port46, Dst: diffswitch3port81},
		{Src: diffdut4port47, Dst: diffswitch3port82},
		{Src: diffdut4port48, Dst: diffswitch3port83},
		{Src: diffdut4port49, Dst: diffswitch3port84},
		{Src: diffdut4port50, Dst: diffswitch3port85},
		{Src: diffdut4port51, Dst: diffswitch3port86},
		{Src: diffdut4port52, Dst: diffswitch3port87},
		{Src: diffdut4port53, Dst: diffswitch3port88},
		{Src: diffdut4port54, Dst: diffswitch3port89},
		{Src: diffdut4port55, Dst: diffswitch3port90},
		{Src: diffdut4port56, Dst: diffswitch3port91},
		{Src: diffdut4port57, Dst: diffswitch3port92},
		{Src: diffdut4port58, Dst: diffswitch3port93},
		{Src: diffdut4port59, Dst: diffswitch3port94},
		{Src: diffdut4port60, Dst: diffswitch3port95},
		{Src: diffdut4port61, Dst: diffswitch3port96},
		{Src: diffdut4port62, Dst: diffswitch3port97},
		{Src: diffdut4port63, Dst: diffswitch3port98},
		{Src: diffdut4port64, Dst: diffswitch3port99},
		{Src: diffdut4port65, Dst: diffswitch3port100},
		{Src: diffdut4port66, Dst: diffswitch3port101},
		{Src: diffdut4port67, Dst: diffswitch3port102},
		{Src: diffdut4port68, Dst: diffswitch3port103},
		{Src: diffdut4port69, Dst: diffswitch3port104},
		{Src: diffdut4port70, Dst: diffswitch3port105},

		{Src: diffswitch3port106, Dst: diffate4port36},
		{Src: diffswitch3port107, Dst: diffate4port37},
		{Src: diffswitch3port108, Dst: diffate4port38},
		{Src: diffswitch3port109, Dst: diffate4port39},
		{Src: diffswitch3port110, Dst: diffate4port40},
		{Src: diffswitch3port111, Dst: diffate4port41},
		{Src: diffswitch3port112, Dst: diffate4port42},
		{Src: diffswitch3port113, Dst: diffate4port43},
		{Src: diffswitch3port114, Dst: diffate4port44},
		{Src: diffswitch3port115, Dst: diffate4port45},
		{Src: diffswitch3port116, Dst: diffate4port46},
		{Src: diffswitch3port117, Dst: diffate4port47},
		{Src: diffswitch3port118, Dst: diffate4port48},
		{Src: diffswitch3port119, Dst: diffate4port49},
		{Src: diffswitch3port120, Dst: diffate4port50},
		{Src: diffswitch3port121, Dst: diffate4port51},
		{Src: diffswitch3port122, Dst: diffate4port52},
		{Src: diffswitch3port123, Dst: diffate4port53},
		{Src: diffswitch3port124, Dst: diffate4port54},
		{Src: diffswitch3port125, Dst: diffate4port55},
		{Src: diffswitch3port126, Dst: diffate4port56},
		{Src: diffswitch3port127, Dst: diffate4port57},
		{Src: diffswitch3port128, Dst: diffate4port58},
		{Src: diffswitch3port129, Dst: diffate4port59},
		{Src: diffswitch3port130, Dst: diffate4port60},
		{Src: diffswitch3port131, Dst: diffate4port61},
		{Src: diffswitch3port132, Dst: diffate4port62},
		{Src: diffswitch3port133, Dst: diffate4port63},
		{Src: diffswitch3port134, Dst: diffate4port64},
		{Src: diffswitch3port135, Dst: diffate4port65},
		{Src: diffswitch3port136, Dst: diffate4port66},
		{Src: diffswitch3port137, Dst: diffate4port67},
		{Src: diffswitch3port138, Dst: diffate4port68},
		{Src: diffswitch3port139, Dst: diffate4port69},
		{Src: diffswitch3port140, Dst: diffate4port70},
	},
}

var (
	absdiffdut1port1  = &AbstractPort{Desc: "absdiffdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port2  = &AbstractPort{Desc: "absdiffdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port3  = &AbstractPort{Desc: "absdiffdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port4  = &AbstractPort{Desc: "absdiffdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port5  = &AbstractPort{Desc: "absdiffdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port6  = &AbstractPort{Desc: "absdiffdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port7  = &AbstractPort{Desc: "absdiffdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port8  = &AbstractPort{Desc: "absdiffdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port9  = &AbstractPort{Desc: "absdiffdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port10 = &AbstractPort{Desc: "absdiffdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut1port11 = &AbstractPort{Desc: "absdiffdut1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port12 = &AbstractPort{Desc: "absdiffdut1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port13 = &AbstractPort{Desc: "absdiffdut1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port14 = &AbstractPort{Desc: "absdiffdut1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port15 = &AbstractPort{Desc: "absdiffdut1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port16 = &AbstractPort{Desc: "absdiffdut1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port17 = &AbstractPort{Desc: "absdiffdut1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port18 = &AbstractPort{Desc: "absdiffdut1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port19 = &AbstractPort{Desc: "absdiffdut1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port20 = &AbstractPort{Desc: "absdiffdut1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut1port21 = &AbstractPort{Desc: "absdiffdut1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port22 = &AbstractPort{Desc: "absdiffdut1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port23 = &AbstractPort{Desc: "absdiffdut1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port24 = &AbstractPort{Desc: "absdiffdut1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port25 = &AbstractPort{Desc: "absdiffdut1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port26 = &AbstractPort{Desc: "absdiffdut1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port27 = &AbstractPort{Desc: "absdiffdut1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port28 = &AbstractPort{Desc: "absdiffdut1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port29 = &AbstractPort{Desc: "absdiffdut1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port30 = &AbstractPort{Desc: "absdiffdut1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port31 = &AbstractPort{Desc: "absdiffdut1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port32 = &AbstractPort{Desc: "absdiffdut1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port33 = &AbstractPort{Desc: "absdiffdut1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port34 = &AbstractPort{Desc: "absdiffdut1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port35 = &AbstractPort{Desc: "absdiffdut1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port36 = &AbstractPort{Desc: "absdiffdut1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port37 = &AbstractPort{Desc: "absdiffdut1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port38 = &AbstractPort{Desc: "absdiffdut1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port39 = &AbstractPort{Desc: "absdiffdut1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port40 = &AbstractPort{Desc: "absdiffdut1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port41 = &AbstractPort{Desc: "absdiffdut1:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port42 = &AbstractPort{Desc: "absdiffdut1:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port43 = &AbstractPort{Desc: "absdiffdut1:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port44 = &AbstractPort{Desc: "absdiffdut1:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port45 = &AbstractPort{Desc: "absdiffdut1:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port46 = &AbstractPort{Desc: "absdiffdut1:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port47 = &AbstractPort{Desc: "absdiffdut1:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port48 = &AbstractPort{Desc: "absdiffdut1:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port49 = &AbstractPort{Desc: "absdiffdut1:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port50 = &AbstractPort{Desc: "absdiffdut1:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port51 = &AbstractPort{Desc: "absdiffdut1:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port52 = &AbstractPort{Desc: "absdiffdut1:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port53 = &AbstractPort{Desc: "absdiffdut1:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port54 = &AbstractPort{Desc: "absdiffdut1:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port55 = &AbstractPort{Desc: "absdiffdut1:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port56 = &AbstractPort{Desc: "absdiffdut1:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port57 = &AbstractPort{Desc: "absdiffdut1:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port58 = &AbstractPort{Desc: "absdiffdut1:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port59 = &AbstractPort{Desc: "absdiffdut1:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port60 = &AbstractPort{Desc: "absdiffdut1:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port61 = &AbstractPort{Desc: "absdiffdut1:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port62 = &AbstractPort{Desc: "absdiffdut1:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port63 = &AbstractPort{Desc: "absdiffdut1:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port64 = &AbstractPort{Desc: "absdiffdut1:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port65 = &AbstractPort{Desc: "absdiffdut1:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port66 = &AbstractPort{Desc: "absdiffdut1:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port67 = &AbstractPort{Desc: "absdiffdut1:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port68 = &AbstractPort{Desc: "absdiffdut1:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port69 = &AbstractPort{Desc: "absdiffdut1:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut1port70 = &AbstractPort{Desc: "absdiffdut1:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffdut1 = &AbstractNode{Desc: "absdiffdut1", Ports: []*AbstractPort{
		absdiffdut1port1,
		absdiffdut1port2,
		absdiffdut1port3,
		absdiffdut1port4,
		absdiffdut1port5,
		absdiffdut1port6,
		absdiffdut1port7,
		absdiffdut1port8,
		absdiffdut1port9,
		absdiffdut1port10,
		absdiffdut1port11,
		absdiffdut1port12,
		absdiffdut1port13,
		absdiffdut1port14,
		absdiffdut1port15,
		absdiffdut1port16,
		absdiffdut1port17,
		absdiffdut1port18,
		absdiffdut1port19,
		absdiffdut1port20,
		absdiffdut1port21,
		absdiffdut1port22,
		absdiffdut1port23,
		absdiffdut1port24,
		absdiffdut1port25,
		absdiffdut1port26,
		absdiffdut1port27,
		absdiffdut1port28,
		absdiffdut1port29,
		absdiffdut1port30,
		absdiffdut1port31,
		absdiffdut1port32,
		absdiffdut1port33,
		absdiffdut1port34,
		absdiffdut1port35,
		absdiffdut1port36,
		absdiffdut1port37,
		absdiffdut1port38,
		absdiffdut1port39,
		absdiffdut1port40,
		absdiffdut1port41,
		absdiffdut1port42,
		absdiffdut1port43,
		absdiffdut1port44,
		absdiffdut1port45,
		absdiffdut1port46,
		absdiffdut1port47,
		absdiffdut1port48,
		absdiffdut1port49,
		absdiffdut1port50,
		absdiffdut1port51,
		absdiffdut1port52,
		absdiffdut1port53,
		absdiffdut1port54,
		absdiffdut1port55,
		absdiffdut1port56,
		absdiffdut1port57,
		absdiffdut1port58,
		absdiffdut1port59,
		absdiffdut1port60,
		absdiffdut1port61,
		absdiffdut1port62,
		absdiffdut1port63,
		absdiffdut1port64,
		absdiffdut1port65,
		absdiffdut1port66,
		absdiffdut1port67,
		absdiffdut1port68,
		absdiffdut1port69,
		absdiffdut1port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "role": Equal("DUT"), "model": Equal("DUT")}}

	absdiffdut2port1  = &AbstractPort{Desc: "absdiffdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port2  = &AbstractPort{Desc: "absdiffdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port3  = &AbstractPort{Desc: "absdiffdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port4  = &AbstractPort{Desc: "absdiffdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port5  = &AbstractPort{Desc: "absdiffdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port6  = &AbstractPort{Desc: "absdiffdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port7  = &AbstractPort{Desc: "absdiffdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port8  = &AbstractPort{Desc: "absdiffdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port9  = &AbstractPort{Desc: "absdiffdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port10 = &AbstractPort{Desc: "absdiffdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut2port11 = &AbstractPort{Desc: "absdiffdut2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port12 = &AbstractPort{Desc: "absdiffdut2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port13 = &AbstractPort{Desc: "absdiffdut2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port14 = &AbstractPort{Desc: "absdiffdut2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port15 = &AbstractPort{Desc: "absdiffdut2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port16 = &AbstractPort{Desc: "absdiffdut2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port17 = &AbstractPort{Desc: "absdiffdut2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port18 = &AbstractPort{Desc: "absdiffdut2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port19 = &AbstractPort{Desc: "absdiffdut2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port20 = &AbstractPort{Desc: "absdiffdut2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut2port21 = &AbstractPort{Desc: "absdiffdut2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port22 = &AbstractPort{Desc: "absdiffdut2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port23 = &AbstractPort{Desc: "absdiffdut2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port24 = &AbstractPort{Desc: "absdiffdut2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port25 = &AbstractPort{Desc: "absdiffdut2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port26 = &AbstractPort{Desc: "absdiffdut2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port27 = &AbstractPort{Desc: "absdiffdut2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port28 = &AbstractPort{Desc: "absdiffdut2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port29 = &AbstractPort{Desc: "absdiffdut2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port30 = &AbstractPort{Desc: "absdiffdut2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port31 = &AbstractPort{Desc: "absdiffdut2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port32 = &AbstractPort{Desc: "absdiffdut2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port33 = &AbstractPort{Desc: "absdiffdut2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port34 = &AbstractPort{Desc: "absdiffdut2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port35 = &AbstractPort{Desc: "absdiffdut2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port36 = &AbstractPort{Desc: "absdiffdut2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port37 = &AbstractPort{Desc: "absdiffdut2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port38 = &AbstractPort{Desc: "absdiffdut2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port39 = &AbstractPort{Desc: "absdiffdut2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port40 = &AbstractPort{Desc: "absdiffdut2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port41 = &AbstractPort{Desc: "absdiffdut2:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port42 = &AbstractPort{Desc: "absdiffdut2:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port43 = &AbstractPort{Desc: "absdiffdut2:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port44 = &AbstractPort{Desc: "absdiffdut2:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port45 = &AbstractPort{Desc: "absdiffdut2:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port46 = &AbstractPort{Desc: "absdiffdut2:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port47 = &AbstractPort{Desc: "absdiffdut2:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port48 = &AbstractPort{Desc: "absdiffdut2:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port49 = &AbstractPort{Desc: "absdiffdut2:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port50 = &AbstractPort{Desc: "absdiffdut2:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port51 = &AbstractPort{Desc: "absdiffdut2:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port52 = &AbstractPort{Desc: "absdiffdut2:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port53 = &AbstractPort{Desc: "absdiffdut2:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port54 = &AbstractPort{Desc: "absdiffdut2:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port55 = &AbstractPort{Desc: "absdiffdut2:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port56 = &AbstractPort{Desc: "absdiffdut2:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port57 = &AbstractPort{Desc: "absdiffdut2:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port58 = &AbstractPort{Desc: "absdiffdut2:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port59 = &AbstractPort{Desc: "absdiffdut2:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port60 = &AbstractPort{Desc: "absdiffdut2:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port61 = &AbstractPort{Desc: "absdiffdut2:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port62 = &AbstractPort{Desc: "absdiffdut2:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port63 = &AbstractPort{Desc: "absdiffdut2:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port64 = &AbstractPort{Desc: "absdiffdut2:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port65 = &AbstractPort{Desc: "absdiffdut2:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port66 = &AbstractPort{Desc: "absdiffdut2:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port67 = &AbstractPort{Desc: "absdiffdut2:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port68 = &AbstractPort{Desc: "absdiffdut2:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port69 = &AbstractPort{Desc: "absdiffdut2:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut2port70 = &AbstractPort{Desc: "absdiffdut2:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffdut2 = &AbstractNode{Desc: "absdiffdut2", Ports: []*AbstractPort{
		absdiffdut2port1,
		absdiffdut2port2,
		absdiffdut2port3,
		absdiffdut2port4,
		absdiffdut2port5,
		absdiffdut2port6,
		absdiffdut2port7,
		absdiffdut2port8,
		absdiffdut2port9,
		absdiffdut2port10,
		absdiffdut2port11,
		absdiffdut2port12,
		absdiffdut2port13,
		absdiffdut2port14,
		absdiffdut2port15,
		absdiffdut2port16,
		absdiffdut2port17,
		absdiffdut2port18,
		absdiffdut2port19,
		absdiffdut2port20,
		absdiffdut2port21,
		absdiffdut2port22,
		absdiffdut2port23,
		absdiffdut2port24,
		absdiffdut2port25,
		absdiffdut2port26,
		absdiffdut2port27,
		absdiffdut2port28,
		absdiffdut2port29,
		absdiffdut2port30,
		absdiffdut2port31,
		absdiffdut2port32,
		absdiffdut2port33,
		absdiffdut2port34,
		absdiffdut2port35,
		absdiffdut2port36,
		absdiffdut2port37,
		absdiffdut2port38,
		absdiffdut2port39,
		absdiffdut2port40,
		absdiffdut2port41,
		absdiffdut2port42,
		absdiffdut2port43,
		absdiffdut2port44,
		absdiffdut2port45,
		absdiffdut2port46,
		absdiffdut2port47,
		absdiffdut2port48,
		absdiffdut2port49,
		absdiffdut2port50,
		absdiffdut2port51,
		absdiffdut2port52,
		absdiffdut2port53,
		absdiffdut2port54,
		absdiffdut2port55,
		absdiffdut2port56,
		absdiffdut2port57,
		absdiffdut2port58,
		absdiffdut2port59,
		absdiffdut2port60,
		absdiffdut2port61,
		absdiffdut2port62,
		absdiffdut2port63,
		absdiffdut2port64,
		absdiffdut2port65,
		absdiffdut2port66,
		absdiffdut2port67,
		absdiffdut2port68,
		absdiffdut2port69,
		absdiffdut2port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "role": Equal("DUT"), "model": Equal("DUT")}}

	absdiffdut3port1  = &AbstractPort{Desc: "absdiffdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port2  = &AbstractPort{Desc: "absdiffdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port3  = &AbstractPort{Desc: "absdiffdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port4  = &AbstractPort{Desc: "absdiffdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port5  = &AbstractPort{Desc: "absdiffdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port6  = &AbstractPort{Desc: "absdiffdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port7  = &AbstractPort{Desc: "absdiffdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port8  = &AbstractPort{Desc: "absdiffdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port9  = &AbstractPort{Desc: "absdiffdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port10 = &AbstractPort{Desc: "absdiffdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut3port11 = &AbstractPort{Desc: "absdiffdut3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port12 = &AbstractPort{Desc: "absdiffdut3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port13 = &AbstractPort{Desc: "absdiffdut3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port14 = &AbstractPort{Desc: "absdiffdut3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port15 = &AbstractPort{Desc: "absdiffdut3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port16 = &AbstractPort{Desc: "absdiffdut3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port17 = &AbstractPort{Desc: "absdiffdut3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port18 = &AbstractPort{Desc: "absdiffdut3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port19 = &AbstractPort{Desc: "absdiffdut3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port20 = &AbstractPort{Desc: "absdiffdut3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut3port21 = &AbstractPort{Desc: "absdiffdut3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port22 = &AbstractPort{Desc: "absdiffdut3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port23 = &AbstractPort{Desc: "absdiffdut3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port24 = &AbstractPort{Desc: "absdiffdut3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port25 = &AbstractPort{Desc: "absdiffdut3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port26 = &AbstractPort{Desc: "absdiffdut3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port27 = &AbstractPort{Desc: "absdiffdut3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port28 = &AbstractPort{Desc: "absdiffdut3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port29 = &AbstractPort{Desc: "absdiffdut3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port30 = &AbstractPort{Desc: "absdiffdut3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port31 = &AbstractPort{Desc: "absdiffdut3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port32 = &AbstractPort{Desc: "absdiffdut3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port33 = &AbstractPort{Desc: "absdiffdut3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port34 = &AbstractPort{Desc: "absdiffdut3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port35 = &AbstractPort{Desc: "absdiffdut3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port36 = &AbstractPort{Desc: "absdiffdut3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port37 = &AbstractPort{Desc: "absdiffdut3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port38 = &AbstractPort{Desc: "absdiffdut3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port39 = &AbstractPort{Desc: "absdiffdut3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port40 = &AbstractPort{Desc: "absdiffdut3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port41 = &AbstractPort{Desc: "absdiffdut3:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port42 = &AbstractPort{Desc: "absdiffdut3:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port43 = &AbstractPort{Desc: "absdiffdut3:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port44 = &AbstractPort{Desc: "absdiffdut3:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port45 = &AbstractPort{Desc: "absdiffdut3:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port46 = &AbstractPort{Desc: "absdiffdut3:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port47 = &AbstractPort{Desc: "absdiffdut3:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port48 = &AbstractPort{Desc: "absdiffdut3:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port49 = &AbstractPort{Desc: "absdiffdut3:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port50 = &AbstractPort{Desc: "absdiffdut3:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port51 = &AbstractPort{Desc: "absdiffdut3:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port52 = &AbstractPort{Desc: "absdiffdut3:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port53 = &AbstractPort{Desc: "absdiffdut3:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port54 = &AbstractPort{Desc: "absdiffdut3:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port55 = &AbstractPort{Desc: "absdiffdut3:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port56 = &AbstractPort{Desc: "absdiffdut3:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port57 = &AbstractPort{Desc: "absdiffdut3:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port58 = &AbstractPort{Desc: "absdiffdut3:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port59 = &AbstractPort{Desc: "absdiffdut3:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port60 = &AbstractPort{Desc: "absdiffdut3:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port61 = &AbstractPort{Desc: "absdiffdut3:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port62 = &AbstractPort{Desc: "absdiffdut3:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port63 = &AbstractPort{Desc: "absdiffdut3:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port64 = &AbstractPort{Desc: "absdiffdut3:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port65 = &AbstractPort{Desc: "absdiffdut3:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port66 = &AbstractPort{Desc: "absdiffdut3:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port67 = &AbstractPort{Desc: "absdiffdut3:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port68 = &AbstractPort{Desc: "absdiffdut3:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port69 = &AbstractPort{Desc: "absdiffdut3:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut3port70 = &AbstractPort{Desc: "absdiffdut3:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffdut3 = &AbstractNode{Desc: "absdiffdut3", Ports: []*AbstractPort{
		absdiffdut3port1,
		absdiffdut3port2,
		absdiffdut3port3,
		absdiffdut3port4,
		absdiffdut3port5,
		absdiffdut3port6,
		absdiffdut3port7,
		absdiffdut3port8,
		absdiffdut3port9,
		absdiffdut3port10,
		absdiffdut3port11,
		absdiffdut3port12,
		absdiffdut3port13,
		absdiffdut3port14,
		absdiffdut3port15,
		absdiffdut3port16,
		absdiffdut3port17,
		absdiffdut3port18,
		absdiffdut3port19,
		absdiffdut3port20,
		absdiffdut3port21,
		absdiffdut3port22,
		absdiffdut3port23,
		absdiffdut3port24,
		absdiffdut3port25,
		absdiffdut3port26,
		absdiffdut3port27,
		absdiffdut3port28,
		absdiffdut3port29,
		absdiffdut3port30,
		absdiffdut3port31,
		absdiffdut3port32,
		absdiffdut3port33,
		absdiffdut3port34,
		absdiffdut3port35,
		absdiffdut3port36,
		absdiffdut3port37,
		absdiffdut3port38,
		absdiffdut3port39,
		absdiffdut3port40,
		absdiffdut3port41,
		absdiffdut3port42,
		absdiffdut3port43,
		absdiffdut3port44,
		absdiffdut3port45,
		absdiffdut3port46,
		absdiffdut3port47,
		absdiffdut3port48,
		absdiffdut3port49,
		absdiffdut3port50,
		absdiffdut3port51,
		absdiffdut3port52,
		absdiffdut3port53,
		absdiffdut3port54,
		absdiffdut3port55,
		absdiffdut3port56,
		absdiffdut3port57,
		absdiffdut3port58,
		absdiffdut3port59,
		absdiffdut3port60,
		absdiffdut3port61,
		absdiffdut3port62,
		absdiffdut3port63,
		absdiffdut3port64,
		absdiffdut3port65,
		absdiffdut3port66,
		absdiffdut3port67,
		absdiffdut3port68,
		absdiffdut3port69,
		absdiffdut3port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "role": Equal("DUT"), "model": Equal("DUT")}}

	absdiffdut4port1  = &AbstractPort{Desc: "absdiffdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port2  = &AbstractPort{Desc: "absdiffdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port3  = &AbstractPort{Desc: "absdiffdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port4  = &AbstractPort{Desc: "absdiffdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port5  = &AbstractPort{Desc: "absdiffdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port6  = &AbstractPort{Desc: "absdiffdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port7  = &AbstractPort{Desc: "absdiffdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port8  = &AbstractPort{Desc: "absdiffdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port9  = &AbstractPort{Desc: "absdiffdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port10 = &AbstractPort{Desc: "absdiffdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffdut4port11 = &AbstractPort{Desc: "absdiffdut4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port12 = &AbstractPort{Desc: "absdiffdut4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port13 = &AbstractPort{Desc: "absdiffdut4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port14 = &AbstractPort{Desc: "absdiffdut4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port15 = &AbstractPort{Desc: "absdiffdut4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port16 = &AbstractPort{Desc: "absdiffdut4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port17 = &AbstractPort{Desc: "absdiffdut4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port18 = &AbstractPort{Desc: "absdiffdut4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port19 = &AbstractPort{Desc: "absdiffdut4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port20 = &AbstractPort{Desc: "absdiffdut4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffdut4port21 = &AbstractPort{Desc: "absdiffdut4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port22 = &AbstractPort{Desc: "absdiffdut4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port23 = &AbstractPort{Desc: "absdiffdut4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port24 = &AbstractPort{Desc: "absdiffdut4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port25 = &AbstractPort{Desc: "absdiffdut4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port26 = &AbstractPort{Desc: "absdiffdut4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port27 = &AbstractPort{Desc: "absdiffdut4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port28 = &AbstractPort{Desc: "absdiffdut4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port29 = &AbstractPort{Desc: "absdiffdut4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port30 = &AbstractPort{Desc: "absdiffdut4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port31 = &AbstractPort{Desc: "absdiffdut4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port32 = &AbstractPort{Desc: "absdiffdut4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port33 = &AbstractPort{Desc: "absdiffdut4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port34 = &AbstractPort{Desc: "absdiffdut4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port35 = &AbstractPort{Desc: "absdiffdut4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port36 = &AbstractPort{Desc: "absdiffdut4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port37 = &AbstractPort{Desc: "absdiffdut4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port38 = &AbstractPort{Desc: "absdiffdut4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port39 = &AbstractPort{Desc: "absdiffdut4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port40 = &AbstractPort{Desc: "absdiffdut4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port41 = &AbstractPort{Desc: "absdiffdut4:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port42 = &AbstractPort{Desc: "absdiffdut4:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port43 = &AbstractPort{Desc: "absdiffdut4:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port44 = &AbstractPort{Desc: "absdiffdut4:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port45 = &AbstractPort{Desc: "absdiffdut4:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port46 = &AbstractPort{Desc: "absdiffdut4:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port47 = &AbstractPort{Desc: "absdiffdut4:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port48 = &AbstractPort{Desc: "absdiffdut4:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port49 = &AbstractPort{Desc: "absdiffdut4:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port50 = &AbstractPort{Desc: "absdiffdut4:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port51 = &AbstractPort{Desc: "absdiffdut4:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port52 = &AbstractPort{Desc: "absdiffdut4:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port53 = &AbstractPort{Desc: "absdiffdut4:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port54 = &AbstractPort{Desc: "absdiffdut4:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port55 = &AbstractPort{Desc: "absdiffdut4:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port56 = &AbstractPort{Desc: "absdiffdut4:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port57 = &AbstractPort{Desc: "absdiffdut4:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port58 = &AbstractPort{Desc: "absdiffdut4:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port59 = &AbstractPort{Desc: "absdiffdut4:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port60 = &AbstractPort{Desc: "absdiffdut4:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port61 = &AbstractPort{Desc: "absdiffdut4:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port62 = &AbstractPort{Desc: "absdiffdut4:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port63 = &AbstractPort{Desc: "absdiffdut4:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port64 = &AbstractPort{Desc: "absdiffdut4:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port65 = &AbstractPort{Desc: "absdiffdut4:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port66 = &AbstractPort{Desc: "absdiffdut4:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port67 = &AbstractPort{Desc: "absdiffdut4:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port68 = &AbstractPort{Desc: "absdiffdut4:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port69 = &AbstractPort{Desc: "absdiffdut4:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffdut4port70 = &AbstractPort{Desc: "absdiffdut4:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffdut4 = &AbstractNode{Desc: "absdiffdut4", Ports: []*AbstractPort{
		absdiffdut4port1,
		absdiffdut4port2,
		absdiffdut4port3,
		absdiffdut4port4,
		absdiffdut4port5,
		absdiffdut4port6,
		absdiffdut4port7,
		absdiffdut4port8,
		absdiffdut4port9,
		absdiffdut4port10,
		absdiffdut4port11,
		absdiffdut4port12,
		absdiffdut4port13,
		absdiffdut4port14,
		absdiffdut4port15,
		absdiffdut4port16,
		absdiffdut4port17,
		absdiffdut4port18,
		absdiffdut4port19,
		absdiffdut4port20,
		absdiffdut4port21,
		absdiffdut4port22,
		absdiffdut4port23,
		absdiffdut4port24,
		absdiffdut4port25,
		absdiffdut4port26,
		absdiffdut4port27,
		absdiffdut4port28,
		absdiffdut4port29,
		absdiffdut4port30,
		absdiffdut4port31,
		absdiffdut4port32,
		absdiffdut4port33,
		absdiffdut4port34,
		absdiffdut4port35,
		absdiffdut4port36,
		absdiffdut4port37,
		absdiffdut4port38,
		absdiffdut4port39,
		absdiffdut4port40,
		absdiffdut4port41,
		absdiffdut4port42,
		absdiffdut4port43,
		absdiffdut4port44,
		absdiffdut4port45,
		absdiffdut4port46,
		absdiffdut4port47,
		absdiffdut4port48,
		absdiffdut4port49,
		absdiffdut4port50,
		absdiffdut4port51,
		absdiffdut4port52,
		absdiffdut4port53,
		absdiffdut4port54,
		absdiffdut4port55,
		absdiffdut4port56,
		absdiffdut4port57,
		absdiffdut4port58,
		absdiffdut4port59,
		absdiffdut4port60,
		absdiffdut4port61,
		absdiffdut4port62,
		absdiffdut4port63,
		absdiffdut4port64,
		absdiffdut4port65,
		absdiffdut4port66,
		absdiffdut4port67,
		absdiffdut4port68,
		absdiffdut4port69,
		absdiffdut4port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "role": Equal("DUT"), "model": Equal("DUT")}}

	absdiffate1port1  = &AbstractPort{Desc: "absdiffate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port2  = &AbstractPort{Desc: "absdiffate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port3  = &AbstractPort{Desc: "absdiffate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port4  = &AbstractPort{Desc: "absdiffate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port5  = &AbstractPort{Desc: "absdiffate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port6  = &AbstractPort{Desc: "absdiffate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port7  = &AbstractPort{Desc: "absdiffate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port8  = &AbstractPort{Desc: "absdiffate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port9  = &AbstractPort{Desc: "absdiffate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port10 = &AbstractPort{Desc: "absdiffate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate1port11 = &AbstractPort{Desc: "absdiffate1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port12 = &AbstractPort{Desc: "absdiffate1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port13 = &AbstractPort{Desc: "absdiffate1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port14 = &AbstractPort{Desc: "absdiffate1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port15 = &AbstractPort{Desc: "absdiffate1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port16 = &AbstractPort{Desc: "absdiffate1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port17 = &AbstractPort{Desc: "absdiffate1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port18 = &AbstractPort{Desc: "absdiffate1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port19 = &AbstractPort{Desc: "absdiffate1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port20 = &AbstractPort{Desc: "absdiffate1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate1port21 = &AbstractPort{Desc: "absdiffate1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port22 = &AbstractPort{Desc: "absdiffate1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port23 = &AbstractPort{Desc: "absdiffate1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port24 = &AbstractPort{Desc: "absdiffate1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port25 = &AbstractPort{Desc: "absdiffate1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port26 = &AbstractPort{Desc: "absdiffate1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port27 = &AbstractPort{Desc: "absdiffate1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port28 = &AbstractPort{Desc: "absdiffate1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port29 = &AbstractPort{Desc: "absdiffate1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port30 = &AbstractPort{Desc: "absdiffate1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port31 = &AbstractPort{Desc: "absdiffate1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port32 = &AbstractPort{Desc: "absdiffate1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port33 = &AbstractPort{Desc: "absdiffate1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port34 = &AbstractPort{Desc: "absdiffate1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port35 = &AbstractPort{Desc: "absdiffate1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port36 = &AbstractPort{Desc: "absdiffate1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port37 = &AbstractPort{Desc: "absdiffate1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port38 = &AbstractPort{Desc: "absdiffate1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port39 = &AbstractPort{Desc: "absdiffate1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port40 = &AbstractPort{Desc: "absdiffate1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port41 = &AbstractPort{Desc: "absdiffate1:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port42 = &AbstractPort{Desc: "absdiffate1:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port43 = &AbstractPort{Desc: "absdiffate1:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port44 = &AbstractPort{Desc: "absdiffate1:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port45 = &AbstractPort{Desc: "absdiffate1:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port46 = &AbstractPort{Desc: "absdiffate1:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port47 = &AbstractPort{Desc: "absdiffate1:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port48 = &AbstractPort{Desc: "absdiffate1:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port49 = &AbstractPort{Desc: "absdiffate1:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port50 = &AbstractPort{Desc: "absdiffate1:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port51 = &AbstractPort{Desc: "absdiffate1:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port52 = &AbstractPort{Desc: "absdiffate1:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port53 = &AbstractPort{Desc: "absdiffate1:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port54 = &AbstractPort{Desc: "absdiffate1:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port55 = &AbstractPort{Desc: "absdiffate1:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port56 = &AbstractPort{Desc: "absdiffate1:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port57 = &AbstractPort{Desc: "absdiffate1:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port58 = &AbstractPort{Desc: "absdiffate1:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port59 = &AbstractPort{Desc: "absdiffate1:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port60 = &AbstractPort{Desc: "absdiffate1:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port61 = &AbstractPort{Desc: "absdiffate1:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port62 = &AbstractPort{Desc: "absdiffate1:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port63 = &AbstractPort{Desc: "absdiffate1:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port64 = &AbstractPort{Desc: "absdiffate1:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port65 = &AbstractPort{Desc: "absdiffate1:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port66 = &AbstractPort{Desc: "absdiffate1:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port67 = &AbstractPort{Desc: "absdiffate1:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port68 = &AbstractPort{Desc: "absdiffate1:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port69 = &AbstractPort{Desc: "absdiffate1:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate1port70 = &AbstractPort{Desc: "absdiffate1:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffate1 = &AbstractNode{Desc: "absdiffate1", Ports: []*AbstractPort{
		absdiffate1port1,
		absdiffate1port2,
		absdiffate1port3,
		absdiffate1port4,
		absdiffate1port5,
		absdiffate1port6,
		absdiffate1port7,
		absdiffate1port8,
		absdiffate1port9,
		absdiffate1port10,
		absdiffate1port11,
		absdiffate1port12,
		absdiffate1port13,
		absdiffate1port14,
		absdiffate1port15,
		absdiffate1port16,
		absdiffate1port17,
		absdiffate1port18,
		absdiffate1port19,
		absdiffate1port20,
		absdiffate1port21,
		absdiffate1port22,
		absdiffate1port23,
		absdiffate1port24,
		absdiffate1port25,
		absdiffate1port26,
		absdiffate1port27,
		absdiffate1port28,
		absdiffate1port29,
		absdiffate1port30,
		absdiffate1port31,
		absdiffate1port32,
		absdiffate1port33,
		absdiffate1port34,
		absdiffate1port35,
		absdiffate1port36,
		absdiffate1port37,
		absdiffate1port38,
		absdiffate1port39,
		absdiffate1port40,
		absdiffate1port41,
		absdiffate1port42,
		absdiffate1port43,
		absdiffate1port44,
		absdiffate1port45,
		absdiffate1port46,
		absdiffate1port47,
		absdiffate1port48,
		absdiffate1port49,
		absdiffate1port50,
		absdiffate1port51,
		absdiffate1port52,
		absdiffate1port53,
		absdiffate1port54,
		absdiffate1port55,
		absdiffate1port56,
		absdiffate1port57,
		absdiffate1port58,
		absdiffate1port59,
		absdiffate1port60,
		absdiffate1port61,
		absdiffate1port62,
		absdiffate1port63,
		absdiffate1port64,
		absdiffate1port65,
		absdiffate1port66,
		absdiffate1port67,
		absdiffate1port68,
		absdiffate1port69,
		absdiffate1port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN"), "role": Equal("ATE"), "model": Equal("ATE")}}

	absdiffate2port1  = &AbstractPort{Desc: "absdiffate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port2  = &AbstractPort{Desc: "absdiffate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port3  = &AbstractPort{Desc: "absdiffate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port4  = &AbstractPort{Desc: "absdiffate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port5  = &AbstractPort{Desc: "absdiffate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port6  = &AbstractPort{Desc: "absdiffate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port7  = &AbstractPort{Desc: "absdiffate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port8  = &AbstractPort{Desc: "absdiffate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port9  = &AbstractPort{Desc: "absdiffate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port10 = &AbstractPort{Desc: "absdiffate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate2port11 = &AbstractPort{Desc: "absdiffate2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port12 = &AbstractPort{Desc: "absdiffate2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port13 = &AbstractPort{Desc: "absdiffate2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port14 = &AbstractPort{Desc: "absdiffate2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port15 = &AbstractPort{Desc: "absdiffate2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port16 = &AbstractPort{Desc: "absdiffate2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port17 = &AbstractPort{Desc: "absdiffate2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port18 = &AbstractPort{Desc: "absdiffate2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port19 = &AbstractPort{Desc: "absdiffate2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port20 = &AbstractPort{Desc: "absdiffate2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate2port21 = &AbstractPort{Desc: "absdiffate2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port22 = &AbstractPort{Desc: "absdiffate2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port23 = &AbstractPort{Desc: "absdiffate2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port24 = &AbstractPort{Desc: "absdiffate2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port25 = &AbstractPort{Desc: "absdiffate2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port26 = &AbstractPort{Desc: "absdiffate2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port27 = &AbstractPort{Desc: "absdiffate2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port28 = &AbstractPort{Desc: "absdiffate2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port29 = &AbstractPort{Desc: "absdiffate2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port30 = &AbstractPort{Desc: "absdiffate2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port31 = &AbstractPort{Desc: "absdiffate2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port32 = &AbstractPort{Desc: "absdiffate2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port33 = &AbstractPort{Desc: "absdiffate2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port34 = &AbstractPort{Desc: "absdiffate2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port35 = &AbstractPort{Desc: "absdiffate2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port36 = &AbstractPort{Desc: "absdiffate2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port37 = &AbstractPort{Desc: "absdiffate2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port38 = &AbstractPort{Desc: "absdiffate2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port39 = &AbstractPort{Desc: "absdiffate2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port40 = &AbstractPort{Desc: "absdiffate2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port41 = &AbstractPort{Desc: "absdiffate2:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port42 = &AbstractPort{Desc: "absdiffate2:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port43 = &AbstractPort{Desc: "absdiffate2:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port44 = &AbstractPort{Desc: "absdiffate2:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port45 = &AbstractPort{Desc: "absdiffate2:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port46 = &AbstractPort{Desc: "absdiffate2:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port47 = &AbstractPort{Desc: "absdiffate2:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port48 = &AbstractPort{Desc: "absdiffate2:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port49 = &AbstractPort{Desc: "absdiffate2:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port50 = &AbstractPort{Desc: "absdiffate2:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port51 = &AbstractPort{Desc: "absdiffate2:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port52 = &AbstractPort{Desc: "absdiffate2:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port53 = &AbstractPort{Desc: "absdiffate2:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port54 = &AbstractPort{Desc: "absdiffate2:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port55 = &AbstractPort{Desc: "absdiffate2:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port56 = &AbstractPort{Desc: "absdiffate2:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port57 = &AbstractPort{Desc: "absdiffate2:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port58 = &AbstractPort{Desc: "absdiffate2:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port59 = &AbstractPort{Desc: "absdiffate2:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port60 = &AbstractPort{Desc: "absdiffate2:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port61 = &AbstractPort{Desc: "absdiffate2:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port62 = &AbstractPort{Desc: "absdiffate2:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port63 = &AbstractPort{Desc: "absdiffate2:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port64 = &AbstractPort{Desc: "absdiffate2:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port65 = &AbstractPort{Desc: "absdiffate2:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port66 = &AbstractPort{Desc: "absdiffate2:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port67 = &AbstractPort{Desc: "absdiffate2:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port68 = &AbstractPort{Desc: "absdiffate2:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port69 = &AbstractPort{Desc: "absdiffate2:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate2port70 = &AbstractPort{Desc: "absdiffate2:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffate2 = &AbstractNode{Desc: "absdiffate2", Ports: []*AbstractPort{
		absdiffate2port1,
		absdiffate2port2,
		absdiffate2port3,
		absdiffate2port4,
		absdiffate2port5,
		absdiffate2port6,
		absdiffate2port7,
		absdiffate2port8,
		absdiffate2port9,
		absdiffate2port10,
		absdiffate2port11,
		absdiffate2port12,
		absdiffate2port13,
		absdiffate2port14,
		absdiffate2port15,
		absdiffate2port16,
		absdiffate2port17,
		absdiffate2port18,
		absdiffate2port19,
		absdiffate2port20,
		absdiffate2port21,
		absdiffate2port22,
		absdiffate2port23,
		absdiffate2port24,
		absdiffate2port25,
		absdiffate2port26,
		absdiffate2port27,
		absdiffate2port28,
		absdiffate2port29,
		absdiffate2port30,
		absdiffate2port31,
		absdiffate2port32,
		absdiffate2port33,
		absdiffate2port34,
		absdiffate2port35,
		absdiffate2port36,
		absdiffate2port37,
		absdiffate2port38,
		absdiffate2port39,
		absdiffate2port40,
		absdiffate2port41,
		absdiffate2port42,
		absdiffate2port43,
		absdiffate2port44,
		absdiffate2port45,
		absdiffate2port46,
		absdiffate2port47,
		absdiffate2port48,
		absdiffate2port49,
		absdiffate2port50,
		absdiffate2port51,
		absdiffate2port52,
		absdiffate2port53,
		absdiffate2port54,
		absdiffate2port55,
		absdiffate2port56,
		absdiffate2port57,
		absdiffate2port58,
		absdiffate2port59,
		absdiffate2port60,
		absdiffate2port61,
		absdiffate2port62,
		absdiffate2port63,
		absdiffate2port64,
		absdiffate2port65,
		absdiffate2port66,
		absdiffate2port67,
		absdiffate2port68,
		absdiffate2port69,
		absdiffate2port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN"), "role": Equal("ATE"), "model": Equal("ATE")}}

	absdiffate3port1  = &AbstractPort{Desc: "absdiffate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port2  = &AbstractPort{Desc: "absdiffate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port3  = &AbstractPort{Desc: "absdiffate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port4  = &AbstractPort{Desc: "absdiffate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port5  = &AbstractPort{Desc: "absdiffate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port6  = &AbstractPort{Desc: "absdiffate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port7  = &AbstractPort{Desc: "absdiffate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port8  = &AbstractPort{Desc: "absdiffate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port9  = &AbstractPort{Desc: "absdiffate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port10 = &AbstractPort{Desc: "absdiffate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate3port11 = &AbstractPort{Desc: "absdiffate3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port12 = &AbstractPort{Desc: "absdiffate3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port13 = &AbstractPort{Desc: "absdiffate3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port14 = &AbstractPort{Desc: "absdiffate3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port15 = &AbstractPort{Desc: "absdiffate3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port16 = &AbstractPort{Desc: "absdiffate3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port17 = &AbstractPort{Desc: "absdiffate3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port18 = &AbstractPort{Desc: "absdiffate3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port19 = &AbstractPort{Desc: "absdiffate3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port20 = &AbstractPort{Desc: "absdiffate3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate3port21 = &AbstractPort{Desc: "absdiffate3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port22 = &AbstractPort{Desc: "absdiffate3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port23 = &AbstractPort{Desc: "absdiffate3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port24 = &AbstractPort{Desc: "absdiffate3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port25 = &AbstractPort{Desc: "absdiffate3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port26 = &AbstractPort{Desc: "absdiffate3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port27 = &AbstractPort{Desc: "absdiffate3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port28 = &AbstractPort{Desc: "absdiffate3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port29 = &AbstractPort{Desc: "absdiffate3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port30 = &AbstractPort{Desc: "absdiffate3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port31 = &AbstractPort{Desc: "absdiffate3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port32 = &AbstractPort{Desc: "absdiffate3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port33 = &AbstractPort{Desc: "absdiffate3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port34 = &AbstractPort{Desc: "absdiffate3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port35 = &AbstractPort{Desc: "absdiffate3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port36 = &AbstractPort{Desc: "absdiffate3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port37 = &AbstractPort{Desc: "absdiffate3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port38 = &AbstractPort{Desc: "absdiffate3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port39 = &AbstractPort{Desc: "absdiffate3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port40 = &AbstractPort{Desc: "absdiffate3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port41 = &AbstractPort{Desc: "absdiffate3:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port42 = &AbstractPort{Desc: "absdiffate3:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port43 = &AbstractPort{Desc: "absdiffate3:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port44 = &AbstractPort{Desc: "absdiffate3:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port45 = &AbstractPort{Desc: "absdiffate3:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port46 = &AbstractPort{Desc: "absdiffate3:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port47 = &AbstractPort{Desc: "absdiffate3:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port48 = &AbstractPort{Desc: "absdiffate3:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port49 = &AbstractPort{Desc: "absdiffate3:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port50 = &AbstractPort{Desc: "absdiffate3:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port51 = &AbstractPort{Desc: "absdiffate3:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port52 = &AbstractPort{Desc: "absdiffate3:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port53 = &AbstractPort{Desc: "absdiffate3:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port54 = &AbstractPort{Desc: "absdiffate3:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port55 = &AbstractPort{Desc: "absdiffate3:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port56 = &AbstractPort{Desc: "absdiffate3:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port57 = &AbstractPort{Desc: "absdiffate3:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port58 = &AbstractPort{Desc: "absdiffate3:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port59 = &AbstractPort{Desc: "absdiffate3:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port60 = &AbstractPort{Desc: "absdiffate3:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port61 = &AbstractPort{Desc: "absdiffate3:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port62 = &AbstractPort{Desc: "absdiffate3:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port63 = &AbstractPort{Desc: "absdiffate3:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port64 = &AbstractPort{Desc: "absdiffate3:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port65 = &AbstractPort{Desc: "absdiffate3:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port66 = &AbstractPort{Desc: "absdiffate3:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port67 = &AbstractPort{Desc: "absdiffate3:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port68 = &AbstractPort{Desc: "absdiffate3:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port69 = &AbstractPort{Desc: "absdiffate3:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate3port70 = &AbstractPort{Desc: "absdiffate3:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffate3 = &AbstractNode{Desc: "absdiffate3", Ports: []*AbstractPort{
		absdiffate3port1,
		absdiffate3port2,
		absdiffate3port3,
		absdiffate3port4,
		absdiffate3port5,
		absdiffate3port6,
		absdiffate3port7,
		absdiffate3port8,
		absdiffate3port9,
		absdiffate3port10,
		absdiffate3port11,
		absdiffate3port12,
		absdiffate3port13,
		absdiffate3port14,
		absdiffate3port15,
		absdiffate3port16,
		absdiffate3port17,
		absdiffate3port18,
		absdiffate3port19,
		absdiffate3port20,
		absdiffate3port21,
		absdiffate3port22,
		absdiffate3port23,
		absdiffate3port24,
		absdiffate3port25,
		absdiffate3port26,
		absdiffate3port27,
		absdiffate3port28,
		absdiffate3port29,
		absdiffate3port30,
		absdiffate3port31,
		absdiffate3port32,
		absdiffate3port33,
		absdiffate3port34,
		absdiffate3port35,
		absdiffate3port36,
		absdiffate3port37,
		absdiffate3port38,
		absdiffate3port39,
		absdiffate3port40,
		absdiffate3port41,
		absdiffate3port42,
		absdiffate3port43,
		absdiffate3port44,
		absdiffate3port45,
		absdiffate3port46,
		absdiffate3port47,
		absdiffate3port48,
		absdiffate3port49,
		absdiffate3port50,
		absdiffate3port51,
		absdiffate3port52,
		absdiffate3port53,
		absdiffate3port54,
		absdiffate3port55,
		absdiffate3port56,
		absdiffate3port57,
		absdiffate3port58,
		absdiffate3port59,
		absdiffate3port60,
		absdiffate3port61,
		absdiffate3port62,
		absdiffate3port63,
		absdiffate3port64,
		absdiffate3port65,
		absdiffate3port66,
		absdiffate3port67,
		absdiffate3port68,
		absdiffate3port69,
		absdiffate3port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN"), "role": Equal("ATE"), "model": Equal("ATE")}}

	absdiffate4port1  = &AbstractPort{Desc: "absdiffate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port2  = &AbstractPort{Desc: "absdiffate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port3  = &AbstractPort{Desc: "absdiffate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port4  = &AbstractPort{Desc: "absdiffate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port5  = &AbstractPort{Desc: "absdiffate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port6  = &AbstractPort{Desc: "absdiffate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port7  = &AbstractPort{Desc: "absdiffate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port8  = &AbstractPort{Desc: "absdiffate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port9  = &AbstractPort{Desc: "absdiffate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port10 = &AbstractPort{Desc: "absdiffate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_100G"), "media": Equal("copper")}}
	absdiffate4port11 = &AbstractPort{Desc: "absdiffate4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port12 = &AbstractPort{Desc: "absdiffate4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port13 = &AbstractPort{Desc: "absdiffate4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port14 = &AbstractPort{Desc: "absdiffate4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port15 = &AbstractPort{Desc: "absdiffate4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port16 = &AbstractPort{Desc: "absdiffate4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port17 = &AbstractPort{Desc: "absdiffate4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port18 = &AbstractPort{Desc: "absdiffate4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port19 = &AbstractPort{Desc: "absdiffate4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port20 = &AbstractPort{Desc: "absdiffate4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_200G"), "media": Equal("copper")}}
	absdiffate4port21 = &AbstractPort{Desc: "absdiffate4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port22 = &AbstractPort{Desc: "absdiffate4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port23 = &AbstractPort{Desc: "absdiffate4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port24 = &AbstractPort{Desc: "absdiffate4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port25 = &AbstractPort{Desc: "absdiffate4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port26 = &AbstractPort{Desc: "absdiffate4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port27 = &AbstractPort{Desc: "absdiffate4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port28 = &AbstractPort{Desc: "absdiffate4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port29 = &AbstractPort{Desc: "absdiffate4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port30 = &AbstractPort{Desc: "absdiffate4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port31 = &AbstractPort{Desc: "absdiffate4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port32 = &AbstractPort{Desc: "absdiffate4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port33 = &AbstractPort{Desc: "absdiffate4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port34 = &AbstractPort{Desc: "absdiffate4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port35 = &AbstractPort{Desc: "absdiffate4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port36 = &AbstractPort{Desc: "absdiffate4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port37 = &AbstractPort{Desc: "absdiffate4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port38 = &AbstractPort{Desc: "absdiffate4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port39 = &AbstractPort{Desc: "absdiffate4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port40 = &AbstractPort{Desc: "absdiffate4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port41 = &AbstractPort{Desc: "absdiffate4:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port42 = &AbstractPort{Desc: "absdiffate4:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port43 = &AbstractPort{Desc: "absdiffate4:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port44 = &AbstractPort{Desc: "absdiffate4:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port45 = &AbstractPort{Desc: "absdiffate4:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port46 = &AbstractPort{Desc: "absdiffate4:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port47 = &AbstractPort{Desc: "absdiffate4:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port48 = &AbstractPort{Desc: "absdiffate4:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port49 = &AbstractPort{Desc: "absdiffate4:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port50 = &AbstractPort{Desc: "absdiffate4:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port51 = &AbstractPort{Desc: "absdiffate4:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port52 = &AbstractPort{Desc: "absdiffate4:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port53 = &AbstractPort{Desc: "absdiffate4:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port54 = &AbstractPort{Desc: "absdiffate4:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port55 = &AbstractPort{Desc: "absdiffate4:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port56 = &AbstractPort{Desc: "absdiffate4:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port57 = &AbstractPort{Desc: "absdiffate4:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port58 = &AbstractPort{Desc: "absdiffate4:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port59 = &AbstractPort{Desc: "absdiffate4:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port60 = &AbstractPort{Desc: "absdiffate4:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port61 = &AbstractPort{Desc: "absdiffate4:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port62 = &AbstractPort{Desc: "absdiffate4:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port63 = &AbstractPort{Desc: "absdiffate4:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port64 = &AbstractPort{Desc: "absdiffate4:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port65 = &AbstractPort{Desc: "absdiffate4:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port66 = &AbstractPort{Desc: "absdiffate4:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port67 = &AbstractPort{Desc: "absdiffate4:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port68 = &AbstractPort{Desc: "absdiffate4:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port69 = &AbstractPort{Desc: "absdiffate4:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}
	absdiffate4port70 = &AbstractPort{Desc: "absdiffate4:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G"), "media": Equal("copper")}}

	absdiffate4 = &AbstractNode{Desc: "absdiffate4", Ports: []*AbstractPort{
		absdiffate4port1,
		absdiffate4port2,
		absdiffate4port3,
		absdiffate4port4,
		absdiffate4port5,
		absdiffate4port6,
		absdiffate4port7,
		absdiffate4port8,
		absdiffate4port9,
		absdiffate4port10,
		absdiffate4port11,
		absdiffate4port12,
		absdiffate4port13,
		absdiffate4port14,
		absdiffate4port15,
		absdiffate4port16,
		absdiffate4port17,
		absdiffate4port18,
		absdiffate4port19,
		absdiffate4port20,
		absdiffate4port21,
		absdiffate4port22,
		absdiffate4port23,
		absdiffate4port24,
		absdiffate4port25,
		absdiffate4port26,
		absdiffate4port27,
		absdiffate4port28,
		absdiffate4port29,
		absdiffate4port30,
		absdiffate4port31,
		absdiffate4port32,
		absdiffate4port33,
		absdiffate4port34,
		absdiffate4port35,
		absdiffate4port36,
		absdiffate4port37,
		absdiffate4port38,
		absdiffate4port39,
		absdiffate4port40,
		absdiffate4port41,
		absdiffate4port42,
		absdiffate4port43,
		absdiffate4port44,
		absdiffate4port45,
		absdiffate4port46,
		absdiffate4port47,
		absdiffate4port48,
		absdiffate4port49,
		absdiffate4port50,
		absdiffate4port51,
		absdiffate4port52,
		absdiffate4port53,
		absdiffate4port54,
		absdiffate4port55,
		absdiffate4port56,
		absdiffate4port57,
		absdiffate4port58,
		absdiffate4port59,
		absdiffate4port60,
		absdiffate4port61,
		absdiffate4port62,
		absdiffate4port63,
		absdiffate4port64,
		absdiffate4port65,
		absdiffate4port66,
		absdiffate4port67,
		absdiffate4port68,
		absdiffate4port69,
		absdiffate4port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN"), "role": Equal("ATE"), "model": Equal("ATE")}}
)

var abstractGraphDiff = &AbstractGraph{
	Nodes: []*AbstractNode{absdiffdut1,
		absdiffate1,
		absdiffdut2,
		absdiffate2,
		absdiffdut3,
		absdiffate3,
		absdiffdut4,
		absdiffate4},
	Edges: []*AbstractEdge{
		{absdiffdut1port1, absdiffate1port1},
		{absdiffdut1port2, absdiffate1port2},
		{absdiffdut1port3, absdiffate1port3},
		{absdiffdut1port4, absdiffate1port4},
		{absdiffdut1port5, absdiffate1port5},
		{absdiffdut1port6, absdiffate1port6},
		{absdiffdut1port7, absdiffate1port7},
		{absdiffdut1port8, absdiffate1port8},
		{absdiffdut1port9, absdiffate1port9},
		{absdiffdut1port10, absdiffate1port10},
		{absdiffdut1port11, absdiffate1port11},
		{absdiffdut1port12, absdiffate1port12},
		{absdiffdut1port13, absdiffate1port13},
		{absdiffdut1port14, absdiffate1port14},
		{absdiffdut1port15, absdiffate1port15},
		{absdiffdut1port16, absdiffate1port16},
		{absdiffdut1port17, absdiffate1port17},
		{absdiffdut1port18, absdiffate1port18},
		{absdiffdut1port19, absdiffate1port19},
		{absdiffdut1port20, absdiffate1port20},
		{absdiffdut1port21, absdiffate1port21},
		{absdiffdut1port22, absdiffate1port22},
		{absdiffdut1port23, absdiffate1port23},
		{absdiffdut1port24, absdiffate1port24},
		{absdiffdut1port25, absdiffate1port25},
		{absdiffdut1port26, absdiffate1port26},
		{absdiffdut1port27, absdiffate1port27},
		{absdiffdut1port28, absdiffate1port28},
		{absdiffdut1port29, absdiffate1port29},
		{absdiffdut1port30, absdiffate1port30},
		{absdiffdut1port31, absdiffate1port31},
		{absdiffdut1port32, absdiffate1port32},
		{absdiffdut1port33, absdiffate1port33},
		{absdiffdut1port34, absdiffate1port34},
		{absdiffdut1port35, absdiffate1port35},
		{absdiffdut1port36, absdiffate1port36},
		{absdiffdut1port37, absdiffate1port37},
		{absdiffdut1port38, absdiffate1port38},
		{absdiffdut1port39, absdiffate1port39},
		{absdiffdut1port40, absdiffate1port40},
		{absdiffdut1port41, absdiffate1port41},
		{absdiffdut1port42, absdiffate1port42},
		{absdiffdut1port43, absdiffate1port43},
		{absdiffdut1port44, absdiffate1port44},
		{absdiffdut1port45, absdiffate1port45},
		{absdiffdut1port46, absdiffate1port46},
		{absdiffdut1port47, absdiffate1port47},
		{absdiffdut1port48, absdiffate1port48},
		{absdiffdut1port49, absdiffate1port49},
		{absdiffdut1port50, absdiffate1port50},
		{absdiffdut1port51, absdiffate1port51},
		{absdiffdut1port52, absdiffate1port52},
		{absdiffdut1port53, absdiffate1port53},
		{absdiffdut1port54, absdiffate1port54},
		{absdiffdut1port55, absdiffate1port55},
		{absdiffdut1port56, absdiffate1port56},
		{absdiffdut1port57, absdiffate1port57},
		{absdiffdut1port58, absdiffate1port58},
		{absdiffdut1port59, absdiffate1port59},
		{absdiffdut1port60, absdiffate1port60},
		{absdiffdut1port61, absdiffate1port61},
		{absdiffdut1port62, absdiffate1port62},
		{absdiffdut1port63, absdiffate1port63},
		{absdiffdut1port64, absdiffate1port64},
		{absdiffdut1port65, absdiffate1port65},
		{absdiffdut1port66, absdiffate1port66},
		{absdiffdut1port67, absdiffate1port67},
		{absdiffdut1port68, absdiffate1port68},
		{absdiffdut1port69, absdiffate1port69},
		{absdiffdut1port70, absdiffate1port70},

		{absdiffdut2port1, absdiffate2port1},
		{absdiffdut2port2, absdiffate2port2},
		{absdiffdut2port3, absdiffate2port3},
		{absdiffdut2port4, absdiffate2port4},
		{absdiffdut2port5, absdiffate2port5},
		{absdiffdut2port6, absdiffate2port6},
		{absdiffdut2port7, absdiffate2port7},
		{absdiffdut2port8, absdiffate2port8},
		{absdiffdut2port9, absdiffate2port9},
		{absdiffdut2port10, absdiffate2port10},
		{absdiffdut2port11, absdiffate2port11},
		{absdiffdut2port12, absdiffate2port12},
		{absdiffdut2port13, absdiffate2port13},
		{absdiffdut2port14, absdiffate2port14},
		{absdiffdut2port15, absdiffate2port15},
		{absdiffdut2port16, absdiffate2port16},
		{absdiffdut2port17, absdiffate2port17},
		{absdiffdut2port18, absdiffate2port18},
		{absdiffdut2port19, absdiffate2port19},
		{absdiffdut2port20, absdiffate2port20},
		{absdiffdut2port21, absdiffate2port21},
		{absdiffdut2port22, absdiffate2port22},
		{absdiffdut2port23, absdiffate2port23},
		{absdiffdut2port24, absdiffate2port24},
		{absdiffdut2port25, absdiffate2port25},
		{absdiffdut2port26, absdiffate2port26},
		{absdiffdut2port27, absdiffate2port27},
		{absdiffdut2port28, absdiffate2port28},
		{absdiffdut2port29, absdiffate2port29},
		{absdiffdut2port30, absdiffate2port30},
		{absdiffdut2port31, absdiffate2port31},
		{absdiffdut2port32, absdiffate2port32},
		{absdiffdut2port33, absdiffate2port33},
		{absdiffdut2port34, absdiffate2port34},
		{absdiffdut2port35, absdiffate2port35},
		{absdiffdut2port36, absdiffate2port36},
		{absdiffdut2port37, absdiffate2port37},
		{absdiffdut2port38, absdiffate2port38},
		{absdiffdut2port39, absdiffate2port39},
		{absdiffdut2port40, absdiffate2port40},
		{absdiffdut2port41, absdiffate2port41},
		{absdiffdut2port42, absdiffate2port42},
		{absdiffdut2port43, absdiffate2port43},
		{absdiffdut2port44, absdiffate2port44},
		{absdiffdut2port45, absdiffate2port45},
		{absdiffdut2port46, absdiffate2port46},
		{absdiffdut2port47, absdiffate2port47},
		{absdiffdut2port48, absdiffate2port48},
		{absdiffdut2port49, absdiffate2port49},
		{absdiffdut2port50, absdiffate2port50},
		{absdiffdut2port51, absdiffate2port51},
		{absdiffdut2port52, absdiffate2port52},
		{absdiffdut2port53, absdiffate2port53},
		{absdiffdut2port54, absdiffate2port54},
		{absdiffdut2port55, absdiffate2port55},
		{absdiffdut2port56, absdiffate2port56},
		{absdiffdut2port57, absdiffate2port57},
		{absdiffdut2port58, absdiffate2port58},
		{absdiffdut2port59, absdiffate2port59},
		{absdiffdut2port60, absdiffate2port60},
		{absdiffdut2port61, absdiffate2port61},
		{absdiffdut2port62, absdiffate2port62},
		{absdiffdut2port63, absdiffate2port63},
		{absdiffdut2port64, absdiffate2port64},
		{absdiffdut2port65, absdiffate2port65},
		{absdiffdut2port66, absdiffate2port66},
		{absdiffdut2port67, absdiffate2port67},
		{absdiffdut2port68, absdiffate2port68},
		{absdiffdut2port69, absdiffate2port69},
		{absdiffdut2port70, absdiffate2port70},

		{absdiffdut3port1, absdiffate3port1},
		{absdiffdut3port2, absdiffate3port2},
		{absdiffdut3port3, absdiffate3port3},
		{absdiffdut3port4, absdiffate3port4},
		{absdiffdut3port5, absdiffate3port5},
		{absdiffdut3port6, absdiffate3port6},
		{absdiffdut3port7, absdiffate3port7},
		{absdiffdut3port8, absdiffate3port8},
		{absdiffdut3port9, absdiffate3port9},
		{absdiffdut3port10, absdiffate3port10},
		{absdiffdut3port11, absdiffate3port11},
		{absdiffdut3port12, absdiffate3port12},
		{absdiffdut3port13, absdiffate3port13},
		{absdiffdut3port14, absdiffate3port14},
		{absdiffdut3port15, absdiffate3port15},
		{absdiffdut3port16, absdiffate3port16},
		{absdiffdut3port17, absdiffate3port17},
		{absdiffdut3port18, absdiffate3port18},
		{absdiffdut3port19, absdiffate3port19},
		{absdiffdut3port20, absdiffate3port20},
		{absdiffdut3port21, absdiffate3port21},
		{absdiffdut3port22, absdiffate3port22},
		{absdiffdut3port23, absdiffate3port23},
		{absdiffdut3port24, absdiffate3port24},
		{absdiffdut3port25, absdiffate3port25},
		{absdiffdut3port26, absdiffate3port26},
		{absdiffdut3port27, absdiffate3port27},
		{absdiffdut3port28, absdiffate3port28},
		{absdiffdut3port29, absdiffate3port29},
		{absdiffdut3port30, absdiffate3port30},
		{absdiffdut3port31, absdiffate3port31},
		{absdiffdut3port32, absdiffate3port32},
		{absdiffdut3port33, absdiffate3port33},
		{absdiffdut3port34, absdiffate3port34},
		{absdiffdut3port35, absdiffate3port35},
		{absdiffdut3port36, absdiffate3port36},
		{absdiffdut3port37, absdiffate3port37},
		{absdiffdut3port38, absdiffate3port38},
		{absdiffdut3port39, absdiffate3port39},
		{absdiffdut3port40, absdiffate3port40},
		{absdiffdut3port41, absdiffate3port41},
		{absdiffdut3port42, absdiffate3port42},
		{absdiffdut3port43, absdiffate3port43},
		{absdiffdut3port44, absdiffate3port44},
		{absdiffdut3port45, absdiffate3port45},
		{absdiffdut3port46, absdiffate3port46},
		{absdiffdut3port47, absdiffate3port47},
		{absdiffdut3port48, absdiffate3port48},
		{absdiffdut3port49, absdiffate3port49},
		{absdiffdut3port50, absdiffate3port50},
		{absdiffdut3port51, absdiffate3port51},
		{absdiffdut3port52, absdiffate3port52},
		{absdiffdut3port53, absdiffate3port53},
		{absdiffdut3port54, absdiffate3port54},
		{absdiffdut3port55, absdiffate3port55},
		{absdiffdut3port56, absdiffate3port56},
		{absdiffdut3port57, absdiffate3port57},
		{absdiffdut3port58, absdiffate3port58},
		{absdiffdut3port59, absdiffate3port59},
		{absdiffdut3port60, absdiffate3port60},
		{absdiffdut3port61, absdiffate3port61},
		{absdiffdut3port62, absdiffate3port62},
		{absdiffdut3port63, absdiffate3port63},
		{absdiffdut3port64, absdiffate3port64},
		{absdiffdut3port65, absdiffate3port65},
		{absdiffdut3port66, absdiffate3port66},
		{absdiffdut3port67, absdiffate3port67},
		{absdiffdut3port68, absdiffate3port68},
		{absdiffdut3port69, absdiffate3port69},
		{absdiffdut3port70, absdiffate3port70},

		{absdiffdut4port1, absdiffate4port1},
		{absdiffdut4port2, absdiffate4port2},
		{absdiffdut4port3, absdiffate4port3},
		{absdiffdut4port4, absdiffate4port4},
		{absdiffdut4port5, absdiffate4port5},
		{absdiffdut4port6, absdiffate4port6},
		{absdiffdut4port7, absdiffate4port7},
		{absdiffdut4port8, absdiffate4port8},
		{absdiffdut4port9, absdiffate4port9},
		{absdiffdut4port10, absdiffate4port10},
		{absdiffdut4port11, absdiffate4port11},
		{absdiffdut4port12, absdiffate4port12},
		{absdiffdut4port13, absdiffate4port13},
		{absdiffdut4port14, absdiffate4port14},
		{absdiffdut4port15, absdiffate4port15},
		{absdiffdut4port16, absdiffate4port16},
		{absdiffdut4port17, absdiffate4port17},
		{absdiffdut4port18, absdiffate4port18},
		{absdiffdut4port19, absdiffate4port19},
		{absdiffdut4port20, absdiffate4port20},
		{absdiffdut4port21, absdiffate4port21},
		{absdiffdut4port22, absdiffate4port22},
		{absdiffdut4port23, absdiffate4port23},
		{absdiffdut4port24, absdiffate4port24},
		{absdiffdut4port25, absdiffate4port25},
		{absdiffdut4port26, absdiffate4port26},
		{absdiffdut4port27, absdiffate4port27},
		{absdiffdut4port28, absdiffate4port28},
		{absdiffdut4port29, absdiffate4port29},
		{absdiffdut4port30, absdiffate4port30},
		{absdiffdut4port31, absdiffate4port31},
		{absdiffdut4port32, absdiffate4port32},
		{absdiffdut4port33, absdiffate4port33},
		{absdiffdut4port34, absdiffate4port34},
		{absdiffdut4port35, absdiffate4port35},
		{absdiffdut4port36, absdiffate4port36},
		{absdiffdut4port37, absdiffate4port37},
		{absdiffdut4port38, absdiffate4port38},
		{absdiffdut4port39, absdiffate4port39},
		{absdiffdut4port40, absdiffate4port40},
		{absdiffdut4port41, absdiffate4port41},
		{absdiffdut4port42, absdiffate4port42},
		{absdiffdut4port43, absdiffate4port43},
		{absdiffdut4port44, absdiffate4port44},
		{absdiffdut4port45, absdiffate4port45},
		{absdiffdut4port46, absdiffate4port46},
		{absdiffdut4port47, absdiffate4port47},
		{absdiffdut4port48, absdiffate4port48},
		{absdiffdut4port49, absdiffate4port49},
		{absdiffdut4port50, absdiffate4port50},
		{absdiffdut4port51, absdiffate4port51},
		{absdiffdut4port52, absdiffate4port52},
		{absdiffdut4port53, absdiffate4port53},
		{absdiffdut4port54, absdiffate4port54},
		{absdiffdut4port55, absdiffate4port55},
		{absdiffdut4port56, absdiffate4port56},
		{absdiffdut4port57, absdiffate4port57},
		{absdiffdut4port58, absdiffate4port58},
		{absdiffdut4port59, absdiffate4port59},
		{absdiffdut4port60, absdiffate4port60},
		{absdiffdut4port61, absdiffate4port61},
		{absdiffdut4port62, absdiffate4port62},
		{absdiffdut4port63, absdiffate4port63},
		{absdiffdut4port64, absdiffate4port64},
		{absdiffdut4port65, absdiffate4port65},
		{absdiffdut4port66, absdiffate4port66},
		{absdiffdut4port67, absdiffate4port67},
		{absdiffdut4port68, absdiffate4port68},
		{absdiffdut4port69, absdiffate4port69},
		{absdiffdut4port70, absdiffate4port70}},
}

func TestSolveTestHybrid70Diff(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraphDiff, superGraphTestDiff)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 8 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 560 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
