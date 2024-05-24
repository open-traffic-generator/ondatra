// package portgraph

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func generateConcreteGraphHybrid(dutCount int, ateCount int, atePortsCount int, dutPortsCount int) *ConcreteGraph {
// 	var (
// 		nodes []*ConcreteNode
// 		// switchPorts []*ConcretePort
// 		edges        []*ConcreteEdge
// 		switchPorts1 []*ConcretePort
// 		switchPorts2 []*ConcretePort
// 		switchPorts3 []*ConcretePort
// 	)
// 	switchCount1 := 1
// 	switchCount2 := 1
// 	switchCount3 := 1
// 	// switchPortsCreated := false  // Flag to track if switch ports are created
// 	// switchPortsCreated1 := false // Flag to track if switch ports are created
// 	for i := 1; i <= dutCount/2; i++ {
// 		var dutports []*ConcretePort
// 		var ateports []*ConcretePort
// 		// var switchPorts1 []*ConcretePort
// 		// var switchPorts2 []*ConcretePort

// 		for j := 1; j <= dutPortsCount/2; j++ {
// 			dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i, j), Attrs: map[string]string{"speed": "S_400G"}}
// 			dutports = append(dutports, dutPort)
// 			// Add edge from DUT to switch
// 			atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port%d", i, j), Attrs: map[string]string{"speed": "S_400G"}}
// 			ateports = append(ateports, atePort)
// 			edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: atePort})
// 		}
// 		for j := 1; j <= dutPortsCount/2; j++ {
// 			dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i, j+dutPortsCount/2), Attrs: map[string]string{"speed": "S_400G"}}
// 			dutports = append(dutports, dutPort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount1)}
// 			switchPorts1 = append(switchPorts1, switchPort)
// 			switchCount1++
// 			edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 		}
// 		for j := 1; j <= dutPortsCount/2; j++ {
// 			atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port%d", i, j+dutPortsCount/2), Attrs: map[string]string{"speed": "S_400G"}}
// 			ateports = append(ateports, atePort)
// 			// Add edge from ATE to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount1)}
// 			switchPorts1 = append(switchPorts1, switchPort)
// 			switchCount1++
// 			edges = append(edges, &ConcreteEdge{Src: atePort, Dst: switchPort})
// 		}
// 		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i), Ports: dutports, Attrs: map[string]string{"vendor": "CISCO", "reserved": "no"}}
// 		// nodes = append(nodes, dut)
// 		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i), Ports: ateports, Attrs: map[string]string{"vendor": "TGEN", "reserved": "no"}}
// 		// if !switchPortsCreated {
// 		// 	switchPort1 := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount1)}
// 		// 	switchPorts1 = append(switchPorts1, switchPort1)
// 		// 	switchCount1++
// 		// 	switchPort2 := &ConcretePort{Desc: fmt.Sprintf("switch2:p%d", switchCount2)}
// 		// 	switchPorts2 = append(switchPorts2, switchPort2)
// 		// 	switchCount2++
// 		// 	edges = append(edges, &ConcreteEdge{Src: switchPort1, Dst: switchPort2})
// 		// 	switchPortsCreated = true
// 		// }

// 		nodes = append(nodes, dut, ate)
// 	}
// 	for i := dutCount / 2; i < dutCount; i++ {

// 		var dutports []*ConcretePort
// 		var ateports []*ConcretePort
// 		// var switchPorts2 []*ConcretePort
// 		// var switchPorts3 []*ConcretePort
// 		for j := 1; j <= dutPortsCount/2; j++ {
// 			dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i+1, j), Attrs: map[string]string{"speed": "S_400G"}}
// 			dutports = append(dutports, dutPort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch2:p%d", switchCount2)}
// 			switchPorts2 = append(switchPorts2, switchPort)
// 			switchCount2++
// 			edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 		}
// 		for j := 1; j <= dutPortsCount/2; j++ {
// 			atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port%d", i, j+dutPortsCount/2), Attrs: map[string]string{"speed": "S_400G"}}
// 			ateports = append(dutports, atePort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch2:p%d", switchCount2)}
// 			switchPorts2 = append(switchPorts2, switchPort)
// 			switchCount2++
// 			edges = append(edges, &ConcreteEdge{Src: atePort, Dst: switchPort})
// 		}
// 		for j := dutPortsCount / 2; j <= dutPortsCount; j++ {
// 			dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i+1, j+1), Attrs: map[string]string{"speed": "S_400G"}}
// 			dutports = append(dutports, dutPort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch3:p%d", switchCount3)}
// 			switchPorts3 = append(switchPorts3, switchPort)
// 			switchCount3++
// 			edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 		}
// 		for j := dutPortsCount / 2; j <= dutPortsCount; j++ {
// 			atePort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i+1, j+1), Attrs: map[string]string{"speed": "S_400G"}}
// 			ateports = append(dutports, atePort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch3:p%d", switchCount3)}
// 			switchPorts3 = append(switchPorts3, switchPort)
// 			switchCount3++
// 			edges = append(edges, &ConcreteEdge{Src: atePort, Dst: switchPort})
// 		}

// 		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i+1), Ports: dutports, Attrs: map[string]string{"vendor": "CISCO", "reserved": "no"}}
// 		// nodes = append(nodes, dut)
// 		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i+1), Ports: ateports, Attrs: map[string]string{"vendor": "TGEN", "reserved": "no"}}
// 		// if !switchPortsCreated1 {
// 		// 	switchPort2 := &ConcretePort{Desc: fmt.Sprintf("switch2:p%d", switchCount2)}
// 		// 	switchPorts2 = append(switchPorts2, switchPort2)
// 		// 	switchCount2++
// 		// 	switchPort3 := &ConcretePort{Desc: fmt.Sprintf("switch3:p%d", switchCount3)}
// 		// 	switchPorts3 = append(switchPorts2, switchPort3)
// 		// 	switchCount3++
// 		// 	edges = append(edges, &ConcreteEdge{Src: switchPort2, Dst: switchPort3})
// 		// 	switchPortsCreated1 = true
// 		// }

// 		nodes = append(nodes, dut, ate)
// 	}
// 	switchPort1 := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount1)}
// 	switchPorts1 = append(switchPorts1, switchPort1)
// 	switchCount1++
// 	switchPort2 := &ConcretePort{Desc: fmt.Sprintf("switch2:p%d", switchCount2)}
// 	switchPorts2 = append(switchPorts2, switchPort2)
// 	switchCount2++
// 	edges = append(edges, &ConcreteEdge{Src: switchPort1, Dst: switchPort2})

// 	switchPort12 := &ConcretePort{Desc: fmt.Sprintf("switch2:p%d", switchCount2)}
// 	switchPorts2 = append(switchPorts2, switchPort12)
// 	switchPort3 := &ConcretePort{Desc: fmt.Sprintf("switch3:p%d", switchCount3)}
// 	switchPorts3 = append(switchPorts3, switchPort3)
// 	edges = append(edges, &ConcreteEdge{Src: switchPort12, Dst: switchPort3})

// 	// Create switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: append([]*ConcretePort{}, switchPorts1...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}

// 	switchNode1 := &ConcreteNode{
// 		Desc:  "switch2",
// 		Ports: append([]*ConcretePort{}, switchPorts2...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw2"},
// 	}
// 	switchNode2 := &ConcreteNode{
// 		Desc:  "switch3",
// 		Ports: append([]*ConcretePort{}, switchPorts3...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw3"},
// 	}
// 	nodes = append(nodes, switchNode, switchNode1, switchNode2)

// 	// Create one switch
// 	// switchNode := &ConcreteNode{
// 	// 	Desc:  "switch1",
// 	// 	Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
// 	// 	Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	// }
// 	// nodes = append(nodes, switchNode)
// 	return &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: nodes,
// 		Edges: edges,
// 	}

// }

// func createAbstractNodesAndEdgesHybrid(dutCount int, ateCount int, dutPortsCount int, atePortsCount int) *AbstractGraph {
// 	var abstractNodes []*AbstractNode
// 	var abstractDutPortsAll []*AbstractPort
// 	var abstractAtePortsAll []*AbstractPort
// 	// halfNodes := totalNodes
// 	// fullNodes := totalNodes * 2
// 	for i := 1; i <= dutCount; i++ {
// 		var abstractDutPorts []*AbstractPort
// 		// if i > 5 && i < 9 {
// 		// 	dutPortsCount = 6
// 		// } else if i > 9 {
// 		// 	dutPortsCount = 5
// 		// }
// 		for j := 1; j <= dutPortsCount; j++ {
// 			abstdutPortDesc := fmt.Sprintf("abstdut%d:port%d", i, j)
// 			abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 			abstractDutPorts = append(abstractDutPorts, abstdutPort)
// 			abstractDutPortsAll = append(abstractDutPortsAll, abstdutPort)
// 		}
// 		abstdutNode := &AbstractNode{Desc: fmt.Sprintf("abstdut%d", i), Ports: abstractDutPorts, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
// 		abstractNodes = append(abstractNodes, abstdutNode)
// 	}
// 	for i := 1; i <= ateCount; i++ {
// 		var abstractAtePorts []*AbstractPort

// 		for j := 1; j <= atePortsCount; j++ {
// 			abstatePortDesc := fmt.Sprintf("abstate%d:port%d", i, j)
// 			abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 			abstractAtePorts = append(abstractAtePorts, abstatePort)
// 			abstractAtePortsAll = append(abstractAtePortsAll, abstatePort)
// 		}
// 		// atePortsCount = 20
// 		abstateNode := &AbstractNode{Desc: fmt.Sprintf("abstate%d", i), Ports: abstractAtePorts, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
// 		abstractNodes = append(abstractNodes, abstateNode)
// 	}
// 	var abstractEdges []*AbstractEdge
// 	count := len(abstractDutPortsAll)
// 	for i := 0; i < count; i++ {
// 		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractDutPortsAll[i], Dst: abstractAtePortsAll[i]})
// 	}

// 	return &AbstractGraph{
// 		Desc:  "",
// 		Nodes: abstractNodes,
// 		Edges: abstractEdges,
// 	}
// }

// func TestSolveTestHybrid(t *testing.T) {
// 	totalNodes := 1
// 	ateCount := 4
// 	dutCount := 4
// 	atePortsCount := 10
// 	dutPortsCount := 10
// 	graph := createAbstractNodesAndEdgesHybrid(dutCount, ateCount, dutPortsCount, atePortsCount)
// 	conGraph := generateConcreteGraphHybrid(dutCount, ateCount, atePortsCount, dutPortsCount)
// 	startTime := time.Now()
// 	a, err := Solve(context.Background(), graph, conGraph)
// 	endTime := time.Now()
// 	if err != nil {
// 		t.Fatalf("Solve got error %v, want nil", err)
// 	}
// 	if len(a.Node2Node) != 8 {
// 		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
// 	}
// 	if len(a.Port2Port) != 141 {
// 		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
// 	}
// 	elapsed := endTime.Sub(startTime)
// 	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// }

// // 4 duts, 4 ates, 70 ports each and 3 switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

var (
	dut1port1  = &ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port2  = &ConcretePort{Desc: "dut1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port3  = &ConcretePort{Desc: "dut1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port4  = &ConcretePort{Desc: "dut1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port5  = &ConcretePort{Desc: "dut1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port6  = &ConcretePort{Desc: "dut1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port7  = &ConcretePort{Desc: "dut1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port8  = &ConcretePort{Desc: "dut1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port9  = &ConcretePort{Desc: "dut1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port10 = &ConcretePort{Desc: "dut1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port11 = &ConcretePort{Desc: "dut1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port12 = &ConcretePort{Desc: "dut1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port13 = &ConcretePort{Desc: "dut1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port14 = &ConcretePort{Desc: "dut1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port15 = &ConcretePort{Desc: "dut1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port16 = &ConcretePort{Desc: "dut1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port17 = &ConcretePort{Desc: "dut1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port18 = &ConcretePort{Desc: "dut1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port19 = &ConcretePort{Desc: "dut1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port20 = &ConcretePort{Desc: "dut1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port21 = &ConcretePort{Desc: "dut1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port22 = &ConcretePort{Desc: "dut1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port23 = &ConcretePort{Desc: "dut1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port24 = &ConcretePort{Desc: "dut1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port25 = &ConcretePort{Desc: "dut1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port26 = &ConcretePort{Desc: "dut1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port27 = &ConcretePort{Desc: "dut1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port28 = &ConcretePort{Desc: "dut1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port29 = &ConcretePort{Desc: "dut1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port30 = &ConcretePort{Desc: "dut1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port31 = &ConcretePort{Desc: "dut1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port32 = &ConcretePort{Desc: "dut1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port33 = &ConcretePort{Desc: "dut1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port34 = &ConcretePort{Desc: "dut1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port35 = &ConcretePort{Desc: "dut1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port36 = &ConcretePort{Desc: "dut1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port37 = &ConcretePort{Desc: "dut1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port38 = &ConcretePort{Desc: "dut1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port39 = &ConcretePort{Desc: "dut1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port40 = &ConcretePort{Desc: "dut1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port41 = &ConcretePort{Desc: "dut1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port42 = &ConcretePort{Desc: "dut1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port43 = &ConcretePort{Desc: "dut1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port44 = &ConcretePort{Desc: "dut1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port45 = &ConcretePort{Desc: "dut1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port46 = &ConcretePort{Desc: "dut1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port47 = &ConcretePort{Desc: "dut1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port48 = &ConcretePort{Desc: "dut1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port49 = &ConcretePort{Desc: "dut1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port50 = &ConcretePort{Desc: "dut1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port51 = &ConcretePort{Desc: "dut1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port52 = &ConcretePort{Desc: "dut1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port53 = &ConcretePort{Desc: "dut1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port54 = &ConcretePort{Desc: "dut1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port55 = &ConcretePort{Desc: "dut1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port56 = &ConcretePort{Desc: "dut1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port57 = &ConcretePort{Desc: "dut1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port58 = &ConcretePort{Desc: "dut1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port59 = &ConcretePort{Desc: "dut1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port60 = &ConcretePort{Desc: "dut1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port61 = &ConcretePort{Desc: "dut1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port62 = &ConcretePort{Desc: "dut1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port63 = &ConcretePort{Desc: "dut1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port64 = &ConcretePort{Desc: "dut1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port65 = &ConcretePort{Desc: "dut1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port66 = &ConcretePort{Desc: "dut1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port67 = &ConcretePort{Desc: "dut1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port68 = &ConcretePort{Desc: "dut1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port69 = &ConcretePort{Desc: "dut1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	dut1port70 = &ConcretePort{Desc: "dut1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	dut1       = &ConcreteNode{Desc: "dut1", Ports: []*ConcretePort{dut1port1, dut1port2, dut1port3, dut1port4, dut1port5, dut1port6, dut1port7, dut1port8, dut1port9, dut1port10, dut1port11, dut1port12, dut1port13, dut1port14, dut1port15, dut1port16, dut1port17, dut1port18, dut1port19, dut1port20, dut1port21, dut1port22, dut1port23, dut1port24, dut1port25, dut1port26, dut1port27, dut1port28, dut1port29, dut1port30, dut1port31, dut1port32, dut1port33, dut1port34, dut1port35, dut1port36, dut1port37, dut1port38, dut1port39, dut1port40, dut1port41, dut1port42, dut1port43, dut1port44, dut1port45, dut1port46, dut1port47, dut1port48, dut1port49, dut1port50, dut1port51, dut1port52, dut1port53, dut1port54, dut1port55, dut1port56, dut1port57, dut1port58, dut1port59, dut1port60, dut1port61, dut1port62, dut1port63, dut1port64, dut1port65, dut1port66, dut1port67, dut1port68, dut1port69, dut1port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	dut2port1  = &ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port2  = &ConcretePort{Desc: "dut2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port3  = &ConcretePort{Desc: "dut2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port4  = &ConcretePort{Desc: "dut2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port5  = &ConcretePort{Desc: "dut2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port6  = &ConcretePort{Desc: "dut2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port7  = &ConcretePort{Desc: "dut2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port8  = &ConcretePort{Desc: "dut2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port9  = &ConcretePort{Desc: "dut2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port10 = &ConcretePort{Desc: "dut2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port11 = &ConcretePort{Desc: "dut2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port12 = &ConcretePort{Desc: "dut2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port13 = &ConcretePort{Desc: "dut2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port14 = &ConcretePort{Desc: "dut2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port15 = &ConcretePort{Desc: "dut2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port16 = &ConcretePort{Desc: "dut2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port17 = &ConcretePort{Desc: "dut2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port18 = &ConcretePort{Desc: "dut2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port19 = &ConcretePort{Desc: "dut2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port20 = &ConcretePort{Desc: "dut2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port21 = &ConcretePort{Desc: "dut2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port22 = &ConcretePort{Desc: "dut2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port23 = &ConcretePort{Desc: "dut2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port24 = &ConcretePort{Desc: "dut2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port25 = &ConcretePort{Desc: "dut2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port26 = &ConcretePort{Desc: "dut2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port27 = &ConcretePort{Desc: "dut2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port28 = &ConcretePort{Desc: "dut2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port29 = &ConcretePort{Desc: "dut2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port30 = &ConcretePort{Desc: "dut2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port31 = &ConcretePort{Desc: "dut2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port32 = &ConcretePort{Desc: "dut2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port33 = &ConcretePort{Desc: "dut2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port34 = &ConcretePort{Desc: "dut2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port35 = &ConcretePort{Desc: "dut2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port36 = &ConcretePort{Desc: "dut2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port37 = &ConcretePort{Desc: "dut2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port38 = &ConcretePort{Desc: "dut2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port39 = &ConcretePort{Desc: "dut2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port40 = &ConcretePort{Desc: "dut2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port41 = &ConcretePort{Desc: "dut2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port42 = &ConcretePort{Desc: "dut2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port43 = &ConcretePort{Desc: "dut2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port44 = &ConcretePort{Desc: "dut2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port45 = &ConcretePort{Desc: "dut2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port46 = &ConcretePort{Desc: "dut2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port47 = &ConcretePort{Desc: "dut2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port48 = &ConcretePort{Desc: "dut2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port49 = &ConcretePort{Desc: "dut2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port50 = &ConcretePort{Desc: "dut2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port51 = &ConcretePort{Desc: "dut2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port52 = &ConcretePort{Desc: "dut2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port53 = &ConcretePort{Desc: "dut2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port54 = &ConcretePort{Desc: "dut2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port55 = &ConcretePort{Desc: "dut2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port56 = &ConcretePort{Desc: "dut2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port57 = &ConcretePort{Desc: "dut2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port58 = &ConcretePort{Desc: "dut2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port59 = &ConcretePort{Desc: "dut2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port60 = &ConcretePort{Desc: "dut2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port61 = &ConcretePort{Desc: "dut2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port62 = &ConcretePort{Desc: "dut2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port63 = &ConcretePort{Desc: "dut2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port64 = &ConcretePort{Desc: "dut2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port65 = &ConcretePort{Desc: "dut2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port66 = &ConcretePort{Desc: "dut2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port67 = &ConcretePort{Desc: "dut2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port68 = &ConcretePort{Desc: "dut2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port69 = &ConcretePort{Desc: "dut2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	dut2port70 = &ConcretePort{Desc: "dut2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	dut2       = &ConcreteNode{Desc: "dut2", Ports: []*ConcretePort{dut2port1, dut2port2, dut2port3, dut2port4, dut2port5, dut2port6, dut2port7, dut2port8, dut2port9, dut2port10, dut2port11, dut2port12, dut2port13, dut2port14, dut2port15, dut2port16, dut2port17, dut2port18, dut2port19, dut2port20, dut2port21, dut2port22, dut2port23, dut2port24, dut2port25, dut2port26, dut2port27, dut2port28, dut2port29, dut2port30, dut2port31, dut2port32, dut2port33, dut2port34, dut2port35, dut2port36, dut2port37, dut2port38, dut2port39, dut2port40, dut2port41, dut2port42, dut2port43, dut2port44, dut2port45, dut2port46, dut2port47, dut2port48, dut2port49, dut2port50, dut2port51, dut2port52, dut2port53, dut2port54, dut2port55, dut2port56, dut2port57, dut2port58, dut2port59, dut2port60, dut2port61, dut2port62, dut2port63, dut2port64, dut2port65, dut2port66, dut2port67, dut2port68, dut2port69, dut2port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	dut3port1  = &ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port2  = &ConcretePort{Desc: "dut3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port3  = &ConcretePort{Desc: "dut3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port4  = &ConcretePort{Desc: "dut3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port5  = &ConcretePort{Desc: "dut3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port6  = &ConcretePort{Desc: "dut3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port7  = &ConcretePort{Desc: "dut3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port8  = &ConcretePort{Desc: "dut3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port9  = &ConcretePort{Desc: "dut3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port10 = &ConcretePort{Desc: "dut3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port11 = &ConcretePort{Desc: "dut3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port12 = &ConcretePort{Desc: "dut3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port13 = &ConcretePort{Desc: "dut3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port14 = &ConcretePort{Desc: "dut3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port15 = &ConcretePort{Desc: "dut3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port16 = &ConcretePort{Desc: "dut3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port17 = &ConcretePort{Desc: "dut3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port18 = &ConcretePort{Desc: "dut3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port19 = &ConcretePort{Desc: "dut3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port20 = &ConcretePort{Desc: "dut3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port21 = &ConcretePort{Desc: "dut3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port22 = &ConcretePort{Desc: "dut3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port23 = &ConcretePort{Desc: "dut3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port24 = &ConcretePort{Desc: "dut3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port25 = &ConcretePort{Desc: "dut3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port26 = &ConcretePort{Desc: "dut3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port27 = &ConcretePort{Desc: "dut3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port28 = &ConcretePort{Desc: "dut3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port29 = &ConcretePort{Desc: "dut3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port30 = &ConcretePort{Desc: "dut3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port31 = &ConcretePort{Desc: "dut3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port32 = &ConcretePort{Desc: "dut3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port33 = &ConcretePort{Desc: "dut3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port34 = &ConcretePort{Desc: "dut3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port35 = &ConcretePort{Desc: "dut3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port36 = &ConcretePort{Desc: "dut3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port37 = &ConcretePort{Desc: "dut3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port38 = &ConcretePort{Desc: "dut3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port39 = &ConcretePort{Desc: "dut3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port40 = &ConcretePort{Desc: "dut3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port41 = &ConcretePort{Desc: "dut3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port42 = &ConcretePort{Desc: "dut3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port43 = &ConcretePort{Desc: "dut3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port44 = &ConcretePort{Desc: "dut3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port45 = &ConcretePort{Desc: "dut3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port46 = &ConcretePort{Desc: "dut3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port47 = &ConcretePort{Desc: "dut3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port48 = &ConcretePort{Desc: "dut3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port49 = &ConcretePort{Desc: "dut3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port50 = &ConcretePort{Desc: "dut3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port51 = &ConcretePort{Desc: "dut3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port52 = &ConcretePort{Desc: "dut3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port53 = &ConcretePort{Desc: "dut3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port54 = &ConcretePort{Desc: "dut3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port55 = &ConcretePort{Desc: "dut3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port56 = &ConcretePort{Desc: "dut3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port57 = &ConcretePort{Desc: "dut3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port58 = &ConcretePort{Desc: "dut3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port59 = &ConcretePort{Desc: "dut3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port60 = &ConcretePort{Desc: "dut3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port61 = &ConcretePort{Desc: "dut3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port62 = &ConcretePort{Desc: "dut3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port63 = &ConcretePort{Desc: "dut3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port64 = &ConcretePort{Desc: "dut3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port65 = &ConcretePort{Desc: "dut3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port66 = &ConcretePort{Desc: "dut3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port67 = &ConcretePort{Desc: "dut3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port68 = &ConcretePort{Desc: "dut3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port69 = &ConcretePort{Desc: "dut3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	dut3port70 = &ConcretePort{Desc: "dut3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	dut3       = &ConcreteNode{Desc: "dut3", Ports: []*ConcretePort{dut3port1, dut3port2, dut3port3, dut3port4, dut3port5, dut3port6, dut3port7, dut3port8, dut3port9, dut3port10, dut3port11, dut3port12, dut3port13, dut3port14, dut3port15, dut3port16, dut3port17, dut3port18, dut3port19, dut3port20, dut3port21, dut3port22, dut3port23, dut3port24, dut3port25, dut3port26, dut3port27, dut3port28, dut3port29, dut3port30, dut3port31, dut3port32, dut3port33, dut3port34, dut3port35, dut3port36, dut3port37, dut3port38, dut3port39, dut3port40, dut3port41, dut3port42, dut3port43, dut3port44, dut3port45, dut3port46, dut3port47, dut3port48, dut3port49, dut3port50, dut3port51, dut3port52, dut3port53, dut3port54, dut3port55, dut3port56, dut3port57, dut3port58, dut3port59, dut3port60, dut3port61, dut3port62, dut3port63, dut3port64, dut3port65, dut3port66, dut3port67, dut3port68, dut3port69, dut3port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	dut4port1  = &ConcretePort{Desc: "dut4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port2  = &ConcretePort{Desc: "dut4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port3  = &ConcretePort{Desc: "dut4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port4  = &ConcretePort{Desc: "dut4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port5  = &ConcretePort{Desc: "dut4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port6  = &ConcretePort{Desc: "dut4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port7  = &ConcretePort{Desc: "dut4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port8  = &ConcretePort{Desc: "dut4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port9  = &ConcretePort{Desc: "dut4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port10 = &ConcretePort{Desc: "dut4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port11 = &ConcretePort{Desc: "dut4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port12 = &ConcretePort{Desc: "dut4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port13 = &ConcretePort{Desc: "dut4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port14 = &ConcretePort{Desc: "dut4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port15 = &ConcretePort{Desc: "dut4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port16 = &ConcretePort{Desc: "dut4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port17 = &ConcretePort{Desc: "dut4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port18 = &ConcretePort{Desc: "dut4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port19 = &ConcretePort{Desc: "dut4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port20 = &ConcretePort{Desc: "dut4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port21 = &ConcretePort{Desc: "dut4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port22 = &ConcretePort{Desc: "dut4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port23 = &ConcretePort{Desc: "dut4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port24 = &ConcretePort{Desc: "dut4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port25 = &ConcretePort{Desc: "dut4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port26 = &ConcretePort{Desc: "dut4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port27 = &ConcretePort{Desc: "dut4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port28 = &ConcretePort{Desc: "dut4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port29 = &ConcretePort{Desc: "dut4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port30 = &ConcretePort{Desc: "dut4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port31 = &ConcretePort{Desc: "dut4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port32 = &ConcretePort{Desc: "dut4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port33 = &ConcretePort{Desc: "dut4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port34 = &ConcretePort{Desc: "dut4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port35 = &ConcretePort{Desc: "dut4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port36 = &ConcretePort{Desc: "dut4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port37 = &ConcretePort{Desc: "dut4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port38 = &ConcretePort{Desc: "dut4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port39 = &ConcretePort{Desc: "dut4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port40 = &ConcretePort{Desc: "dut4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port41 = &ConcretePort{Desc: "dut4:port41", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port42 = &ConcretePort{Desc: "dut4:port42", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port43 = &ConcretePort{Desc: "dut4:port43", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port44 = &ConcretePort{Desc: "dut4:port44", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port45 = &ConcretePort{Desc: "dut4:port45", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port46 = &ConcretePort{Desc: "dut4:port46", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port47 = &ConcretePort{Desc: "dut4:port47", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port48 = &ConcretePort{Desc: "dut4:port48", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port49 = &ConcretePort{Desc: "dut4:port49", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port50 = &ConcretePort{Desc: "dut4:port50", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port51 = &ConcretePort{Desc: "dut4:port51", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port52 = &ConcretePort{Desc: "dut4:port52", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port53 = &ConcretePort{Desc: "dut4:port53", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port54 = &ConcretePort{Desc: "dut4:port54", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port55 = &ConcretePort{Desc: "dut4:port55", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port56 = &ConcretePort{Desc: "dut4:port56", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port57 = &ConcretePort{Desc: "dut4:port57", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port58 = &ConcretePort{Desc: "dut4:port58", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port59 = &ConcretePort{Desc: "dut4:port59", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port60 = &ConcretePort{Desc: "dut4:port60", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port61 = &ConcretePort{Desc: "dut4:port61", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port62 = &ConcretePort{Desc: "dut4:port62", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port63 = &ConcretePort{Desc: "dut4:port63", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port64 = &ConcretePort{Desc: "dut4:port64", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port65 = &ConcretePort{Desc: "dut4:port65", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port66 = &ConcretePort{Desc: "dut4:port66", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port67 = &ConcretePort{Desc: "dut4:port67", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port68 = &ConcretePort{Desc: "dut4:port68", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port69 = &ConcretePort{Desc: "dut4:port69", Attrs: map[string]string{"speed": "S_400G"}}
	dut4port70 = &ConcretePort{Desc: "dut4:port70", Attrs: map[string]string{"speed": "S_400G"}}
	dut4       = &ConcreteNode{Desc: "dut4", Ports: []*ConcretePort{dut4port1, dut4port2, dut4port3, dut4port4, dut4port5, dut4port6, dut4port7, dut4port8, dut4port9, dut4port10, dut4port11, dut4port12, dut4port13, dut4port14, dut4port15, dut4port16, dut4port17, dut4port18, dut4port19, dut4port20, dut4port21, dut4port22, dut4port23, dut4port24, dut4port25, dut4port26, dut4port27, dut4port28, dut4port29, dut4port30, dut4port31, dut4port32, dut4port33, dut4port34, dut4port35, dut4port36, dut4port37, dut4port38, dut4port39, dut4port40, dut4port41, dut4port42, dut4port43, dut4port44, dut4port45, dut4port46, dut4port47, dut4port48, dut4port49, dut4port50, dut4port51, dut4port52, dut4port53, dut4port54, dut4port55, dut4port56, dut4port57, dut4port58, dut4port59, dut4port60, dut4port61, dut4port62, dut4port63, dut4port64, dut4port65, dut4port66, dut4port67, dut4port68, dut4port69, dut4port70}, Attrs: map[string]string{"vendor": "CISCO"}}

	ate1port1  = &ConcretePort{Desc: "ate1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port2  = &ConcretePort{Desc: "ate1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port3  = &ConcretePort{Desc: "ate1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port4  = &ConcretePort{Desc: "ate1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port5  = &ConcretePort{Desc: "ate1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port6  = &ConcretePort{Desc: "ate1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port7  = &ConcretePort{Desc: "ate1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port8  = &ConcretePort{Desc: "ate1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port9  = &ConcretePort{Desc: "ate1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port10 = &ConcretePort{Desc: "ate1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port11 = &ConcretePort{Desc: "ate1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port12 = &ConcretePort{Desc: "ate1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port13 = &ConcretePort{Desc: "ate1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port14 = &ConcretePort{Desc: "ate1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port15 = &ConcretePort{Desc: "ate1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port16 = &ConcretePort{Desc: "ate1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port17 = &ConcretePort{Desc: "ate1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port18 = &ConcretePort{Desc: "ate1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port19 = &ConcretePort{Desc: "ate1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port20 = &ConcretePort{Desc: "ate1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port21 = &ConcretePort{Desc: "ate1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port22 = &ConcretePort{Desc: "ate1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port23 = &ConcretePort{Desc: "ate1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port24 = &ConcretePort{Desc: "ate1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port25 = &ConcretePort{Desc: "ate1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port26 = &ConcretePort{Desc: "ate1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port27 = &ConcretePort{Desc: "ate1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port28 = &ConcretePort{Desc: "ate1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port29 = &ConcretePort{Desc: "ate1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port30 = &ConcretePort{Desc: "ate1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port31 = &ConcretePort{Desc: "ate1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port32 = &ConcretePort{Desc: "ate1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port33 = &ConcretePort{Desc: "ate1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port34 = &ConcretePort{Desc: "ate1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port35 = &ConcretePort{Desc: "ate1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port36 = &ConcretePort{Desc: "ate1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port37 = &ConcretePort{Desc: "ate1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port38 = &ConcretePort{Desc: "ate1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port39 = &ConcretePort{Desc: "ate1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port40 = &ConcretePort{Desc: "ate1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port41 = &ConcretePort{Desc: "ate1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port42 = &ConcretePort{Desc: "ate1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port43 = &ConcretePort{Desc: "ate1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port44 = &ConcretePort{Desc: "ate1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port45 = &ConcretePort{Desc: "ate1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port46 = &ConcretePort{Desc: "ate1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port47 = &ConcretePort{Desc: "ate1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port48 = &ConcretePort{Desc: "ate1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port49 = &ConcretePort{Desc: "ate1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port50 = &ConcretePort{Desc: "ate1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port51 = &ConcretePort{Desc: "ate1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port52 = &ConcretePort{Desc: "ate1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port53 = &ConcretePort{Desc: "ate1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port54 = &ConcretePort{Desc: "ate1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port55 = &ConcretePort{Desc: "ate1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port56 = &ConcretePort{Desc: "ate1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port57 = &ConcretePort{Desc: "ate1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port58 = &ConcretePort{Desc: "ate1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port59 = &ConcretePort{Desc: "ate1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port60 = &ConcretePort{Desc: "ate1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port61 = &ConcretePort{Desc: "ate1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port62 = &ConcretePort{Desc: "ate1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port63 = &ConcretePort{Desc: "ate1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port64 = &ConcretePort{Desc: "ate1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port65 = &ConcretePort{Desc: "ate1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port66 = &ConcretePort{Desc: "ate1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port67 = &ConcretePort{Desc: "ate1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port68 = &ConcretePort{Desc: "ate1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port69 = &ConcretePort{Desc: "ate1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	ate1port70 = &ConcretePort{Desc: "ate1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	ate1       = &ConcreteNode{Desc: "ate1", Ports: []*ConcretePort{ate1port1, ate1port2, ate1port3, ate1port4, ate1port5, ate1port6, ate1port7, ate1port8, ate1port9, ate1port10, ate1port11, ate1port12, ate1port13, ate1port14, ate1port15, ate1port16, ate1port17, ate1port18, ate1port19, ate1port20, ate1port21, ate1port22, ate1port23, ate1port24, ate1port25, ate1port26, ate1port27, ate1port28, ate1port29, ate1port30, ate1port31, ate1port32, ate1port33, ate1port34, ate1port35, ate1port36, ate1port37, ate1port38, ate1port39, ate1port40, ate1port41, ate1port42, ate1port43, ate1port44, ate1port45, ate1port46, ate1port47, ate1port48, ate1port49, ate1port50, ate1port51, ate1port52, ate1port53, ate1port54, ate1port55, ate1port56, ate1port57, ate1port58, ate1port59, ate1port60, ate1port61, ate1port62, ate1port63, ate1port64, ate1port65, ate1port66, ate1port67, ate1port68, ate1port69, ate1port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	ate2port1  = &ConcretePort{Desc: "ate2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port2  = &ConcretePort{Desc: "ate2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port3  = &ConcretePort{Desc: "ate2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port4  = &ConcretePort{Desc: "ate2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port5  = &ConcretePort{Desc: "ate2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port6  = &ConcretePort{Desc: "ate2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port7  = &ConcretePort{Desc: "ate2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port8  = &ConcretePort{Desc: "ate2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port9  = &ConcretePort{Desc: "ate2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port10 = &ConcretePort{Desc: "ate2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port11 = &ConcretePort{Desc: "ate2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port12 = &ConcretePort{Desc: "ate2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port13 = &ConcretePort{Desc: "ate2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port14 = &ConcretePort{Desc: "ate2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port15 = &ConcretePort{Desc: "ate2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port16 = &ConcretePort{Desc: "ate2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port17 = &ConcretePort{Desc: "ate2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port18 = &ConcretePort{Desc: "ate2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port19 = &ConcretePort{Desc: "ate2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port20 = &ConcretePort{Desc: "ate2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port21 = &ConcretePort{Desc: "ate2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port22 = &ConcretePort{Desc: "ate2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port23 = &ConcretePort{Desc: "ate2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port24 = &ConcretePort{Desc: "ate2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port25 = &ConcretePort{Desc: "ate2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port26 = &ConcretePort{Desc: "ate2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port27 = &ConcretePort{Desc: "ate2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port28 = &ConcretePort{Desc: "ate2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port29 = &ConcretePort{Desc: "ate2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port30 = &ConcretePort{Desc: "ate2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port31 = &ConcretePort{Desc: "ate2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port32 = &ConcretePort{Desc: "ate2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port33 = &ConcretePort{Desc: "ate2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port34 = &ConcretePort{Desc: "ate2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port35 = &ConcretePort{Desc: "ate2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port36 = &ConcretePort{Desc: "ate2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port37 = &ConcretePort{Desc: "ate2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port38 = &ConcretePort{Desc: "ate2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port39 = &ConcretePort{Desc: "ate2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port40 = &ConcretePort{Desc: "ate2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port41 = &ConcretePort{Desc: "ate2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port42 = &ConcretePort{Desc: "ate2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port43 = &ConcretePort{Desc: "ate2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port44 = &ConcretePort{Desc: "ate2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port45 = &ConcretePort{Desc: "ate2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port46 = &ConcretePort{Desc: "ate2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port47 = &ConcretePort{Desc: "ate2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port48 = &ConcretePort{Desc: "ate2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port49 = &ConcretePort{Desc: "ate2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port50 = &ConcretePort{Desc: "ate2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port51 = &ConcretePort{Desc: "ate2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port52 = &ConcretePort{Desc: "ate2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port53 = &ConcretePort{Desc: "ate2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port54 = &ConcretePort{Desc: "ate2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port55 = &ConcretePort{Desc: "ate2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port56 = &ConcretePort{Desc: "ate2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port57 = &ConcretePort{Desc: "ate2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port58 = &ConcretePort{Desc: "ate2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port59 = &ConcretePort{Desc: "ate2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port60 = &ConcretePort{Desc: "ate2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port61 = &ConcretePort{Desc: "ate2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port62 = &ConcretePort{Desc: "ate2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port63 = &ConcretePort{Desc: "ate2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port64 = &ConcretePort{Desc: "ate2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port65 = &ConcretePort{Desc: "ate2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port66 = &ConcretePort{Desc: "ate2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port67 = &ConcretePort{Desc: "ate2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port68 = &ConcretePort{Desc: "ate2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port69 = &ConcretePort{Desc: "ate2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	ate2port70 = &ConcretePort{Desc: "ate2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	ate2       = &ConcreteNode{Desc: "ate2", Ports: []*ConcretePort{ate2port1, ate2port2, ate2port3, ate2port4, ate2port5, ate2port6, ate2port7, ate2port8, ate2port9, ate2port10, ate2port11, ate2port12, ate2port13, ate2port14, ate2port15, ate2port16, ate2port17, ate2port18, ate2port19, ate2port20, ate2port21, ate2port22, ate2port23, ate2port24, ate2port25, ate2port26, ate2port27, ate2port28, ate2port29, ate2port30, ate2port31, ate2port32, ate2port33, ate2port34, ate2port35, ate2port36, ate2port37, ate2port38, ate2port39, ate2port40, ate2port41, ate2port42, ate2port43, ate2port44, ate2port45, ate2port46, ate2port47, ate2port48, ate2port49, ate2port50, ate2port51, ate2port52, ate2port53, ate2port54, ate2port55, ate2port56, ate2port57, ate2port58, ate2port59, ate2port60, ate2port61, ate2port62, ate2port63, ate2port64, ate2port65, ate2port66, ate2port67, ate2port68, ate2port69, ate2port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	ate3port1  = &ConcretePort{Desc: "ate3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port2  = &ConcretePort{Desc: "ate3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port3  = &ConcretePort{Desc: "ate3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port4  = &ConcretePort{Desc: "ate3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port5  = &ConcretePort{Desc: "ate3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port6  = &ConcretePort{Desc: "ate3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port7  = &ConcretePort{Desc: "ate3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port8  = &ConcretePort{Desc: "ate3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port9  = &ConcretePort{Desc: "ate3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port10 = &ConcretePort{Desc: "ate3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port11 = &ConcretePort{Desc: "ate3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port12 = &ConcretePort{Desc: "ate3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port13 = &ConcretePort{Desc: "ate3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port14 = &ConcretePort{Desc: "ate3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port15 = &ConcretePort{Desc: "ate3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port16 = &ConcretePort{Desc: "ate3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port17 = &ConcretePort{Desc: "ate3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port18 = &ConcretePort{Desc: "ate3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port19 = &ConcretePort{Desc: "ate3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port20 = &ConcretePort{Desc: "ate3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port21 = &ConcretePort{Desc: "ate3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port22 = &ConcretePort{Desc: "ate3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port23 = &ConcretePort{Desc: "ate3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port24 = &ConcretePort{Desc: "ate3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port25 = &ConcretePort{Desc: "ate3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port26 = &ConcretePort{Desc: "ate3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port27 = &ConcretePort{Desc: "ate3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port28 = &ConcretePort{Desc: "ate3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port29 = &ConcretePort{Desc: "ate3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port30 = &ConcretePort{Desc: "ate3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port31 = &ConcretePort{Desc: "ate3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port32 = &ConcretePort{Desc: "ate3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port33 = &ConcretePort{Desc: "ate3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port34 = &ConcretePort{Desc: "ate3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port35 = &ConcretePort{Desc: "ate3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port36 = &ConcretePort{Desc: "ate3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port37 = &ConcretePort{Desc: "ate3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port38 = &ConcretePort{Desc: "ate3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port39 = &ConcretePort{Desc: "ate3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port40 = &ConcretePort{Desc: "ate3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port41 = &ConcretePort{Desc: "ate3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port42 = &ConcretePort{Desc: "ate3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port43 = &ConcretePort{Desc: "ate3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port44 = &ConcretePort{Desc: "ate3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port45 = &ConcretePort{Desc: "ate3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port46 = &ConcretePort{Desc: "ate3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port47 = &ConcretePort{Desc: "ate3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port48 = &ConcretePort{Desc: "ate3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port49 = &ConcretePort{Desc: "ate3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port50 = &ConcretePort{Desc: "ate3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port51 = &ConcretePort{Desc: "ate3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port52 = &ConcretePort{Desc: "ate3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port53 = &ConcretePort{Desc: "ate3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port54 = &ConcretePort{Desc: "ate3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port55 = &ConcretePort{Desc: "ate3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port56 = &ConcretePort{Desc: "ate3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port57 = &ConcretePort{Desc: "ate3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port58 = &ConcretePort{Desc: "ate3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port59 = &ConcretePort{Desc: "ate3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port60 = &ConcretePort{Desc: "ate3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port61 = &ConcretePort{Desc: "ate3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port62 = &ConcretePort{Desc: "ate3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port63 = &ConcretePort{Desc: "ate3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port64 = &ConcretePort{Desc: "ate3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port65 = &ConcretePort{Desc: "ate3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port66 = &ConcretePort{Desc: "ate3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port67 = &ConcretePort{Desc: "ate3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port68 = &ConcretePort{Desc: "ate3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port69 = &ConcretePort{Desc: "ate3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	ate3port70 = &ConcretePort{Desc: "ate3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	ate3       = &ConcreteNode{Desc: "ate3", Ports: []*ConcretePort{ate3port1, ate3port2, ate3port3, ate3port4, ate3port5, ate3port6, ate3port7, ate3port8, ate3port9, ate3port10, ate3port11, ate3port12, ate3port13, ate3port14, ate3port15, ate3port16, ate3port17, ate3port18, ate3port19, ate3port20, ate3port21, ate3port22, ate3port23, ate3port24, ate3port25, ate3port26, ate3port27, ate3port28, ate3port29, ate3port30, ate3port31, ate3port32, ate3port33, ate3port34, ate3port35, ate3port36, ate3port37, ate3port38, ate3port39, ate3port40, ate3port41, ate3port42, ate3port43, ate3port44, ate3port45, ate3port46, ate3port47, ate3port48, ate3port49, ate3port50, ate3port51, ate3port52, ate3port53, ate3port54, ate3port55, ate3port56, ate3port57, ate3port58, ate3port59, ate3port60, ate3port61, ate3port62, ate3port63, ate3port64, ate3port65, ate3port66, ate3port67, ate3port68, ate3port69, ate3port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	ate4port1  = &ConcretePort{Desc: "ate4:port1", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port2  = &ConcretePort{Desc: "ate4:port2", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port3  = &ConcretePort{Desc: "ate4:port3", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port4  = &ConcretePort{Desc: "ate4:port4", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port5  = &ConcretePort{Desc: "ate4:port5", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port6  = &ConcretePort{Desc: "ate4:port6", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port7  = &ConcretePort{Desc: "ate4:port7", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port8  = &ConcretePort{Desc: "ate4:port8", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port9  = &ConcretePort{Desc: "ate4:port9", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port10 = &ConcretePort{Desc: "ate4:port10", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port11 = &ConcretePort{Desc: "ate4:port11", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port12 = &ConcretePort{Desc: "ate4:port12", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port13 = &ConcretePort{Desc: "ate4:port13", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port14 = &ConcretePort{Desc: "ate4:port14", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port15 = &ConcretePort{Desc: "ate4:port15", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port16 = &ConcretePort{Desc: "ate4:port16", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port17 = &ConcretePort{Desc: "ate4:port17", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port18 = &ConcretePort{Desc: "ate4:port18", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port19 = &ConcretePort{Desc: "ate4:port19", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port20 = &ConcretePort{Desc: "ate4:port20", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port21 = &ConcretePort{Desc: "ate4:port21", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port22 = &ConcretePort{Desc: "ate4:port22", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port23 = &ConcretePort{Desc: "ate4:port23", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port24 = &ConcretePort{Desc: "ate4:port24", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port25 = &ConcretePort{Desc: "ate4:port25", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port26 = &ConcretePort{Desc: "ate4:port26", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port27 = &ConcretePort{Desc: "ate4:port27", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port28 = &ConcretePort{Desc: "ate4:port28", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port29 = &ConcretePort{Desc: "ate4:port29", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port30 = &ConcretePort{Desc: "ate4:port30", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port31 = &ConcretePort{Desc: "ate4:port31", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port32 = &ConcretePort{Desc: "ate4:port32", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port33 = &ConcretePort{Desc: "ate4:port33", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port34 = &ConcretePort{Desc: "ate4:port34", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port35 = &ConcretePort{Desc: "ate4:port35", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port36 = &ConcretePort{Desc: "ate4:port36", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port37 = &ConcretePort{Desc: "ate4:port37", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port38 = &ConcretePort{Desc: "ate4:port38", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port39 = &ConcretePort{Desc: "ate4:port39", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port40 = &ConcretePort{Desc: "ate4:port40", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port41 = &ConcretePort{Desc: "ate4:port41", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port42 = &ConcretePort{Desc: "ate4:port42", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port43 = &ConcretePort{Desc: "ate4:port43", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port44 = &ConcretePort{Desc: "ate4:port44", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port45 = &ConcretePort{Desc: "ate4:port45", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port46 = &ConcretePort{Desc: "ate4:port46", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port47 = &ConcretePort{Desc: "ate4:port47", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port48 = &ConcretePort{Desc: "ate4:port48", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port49 = &ConcretePort{Desc: "ate4:port49", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port50 = &ConcretePort{Desc: "ate4:port50", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port51 = &ConcretePort{Desc: "ate4:port51", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port52 = &ConcretePort{Desc: "ate4:port52", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port53 = &ConcretePort{Desc: "ate4:port53", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port54 = &ConcretePort{Desc: "ate4:port54", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port55 = &ConcretePort{Desc: "ate4:port55", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port56 = &ConcretePort{Desc: "ate4:port56", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port57 = &ConcretePort{Desc: "ate4:port57", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port58 = &ConcretePort{Desc: "ate4:port58", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port59 = &ConcretePort{Desc: "ate4:port59", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port60 = &ConcretePort{Desc: "ate4:port60", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port61 = &ConcretePort{Desc: "ate4:port61", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port62 = &ConcretePort{Desc: "ate4:port62", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port63 = &ConcretePort{Desc: "ate4:port63", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port64 = &ConcretePort{Desc: "ate4:port64", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port65 = &ConcretePort{Desc: "ate4:port65", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port66 = &ConcretePort{Desc: "ate4:port66", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port67 = &ConcretePort{Desc: "ate4:port67", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port68 = &ConcretePort{Desc: "ate4:port68", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port69 = &ConcretePort{Desc: "ate4:port69", Attrs: map[string]string{"speed": "S_400G"}}
	ate4port70 = &ConcretePort{Desc: "ate4:port70", Attrs: map[string]string{"speed": "S_400G"}}
	ate4       = &ConcreteNode{Desc: "ate4", Ports: []*ConcretePort{ate4port1, ate4port2, ate4port3, ate4port4, ate4port5, ate4port6, ate4port7, ate4port8, ate4port9, ate4port10, ate4port11, ate4port12, ate4port13, ate4port14, ate4port15, ate4port16, ate4port17, ate4port18, ate4port19, ate4port20, ate4port21, ate4port22, ate4port23, ate4port24, ate4port25, ate4port26, ate4port27, ate4port28, ate4port29, ate4port30, ate4port31, ate4port32, ate4port33, ate4port34, ate4port35, ate4port36, ate4port37, ate4port38, ate4port39, ate4port40, ate4port41, ate4port42, ate4port43, ate4port44, ate4port45, ate4port46, ate4port47, ate4port48, ate4port49, ate4port50, ate4port51, ate4port52, ate4port53, ate4port54, ate4port55, ate4port56, ate4port57, ate4port58, ate4port59, ate4port60, ate4port61, ate4port62, ate4port63, ate4port64, ate4port65, ate4port66, ate4port67, ate4port68, ate4port69, ate4port70}, Attrs: map[string]string{"vendor": "TGEN"}}

	switch1port1   = &ConcretePort{Desc: "switch1:port1", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port2   = &ConcretePort{Desc: "switch1:port2", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port3   = &ConcretePort{Desc: "switch1:port3", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port4   = &ConcretePort{Desc: "switch1:port4", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port5   = &ConcretePort{Desc: "switch1:port5", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port6   = &ConcretePort{Desc: "switch1:port6", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port7   = &ConcretePort{Desc: "switch1:port7", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port8   = &ConcretePort{Desc: "switch1:port8", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port9   = &ConcretePort{Desc: "switch1:port9", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port10  = &ConcretePort{Desc: "switch1:port10", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port11  = &ConcretePort{Desc: "switch1:port11", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port12  = &ConcretePort{Desc: "switch1:port12", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port13  = &ConcretePort{Desc: "switch1:port13", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port14  = &ConcretePort{Desc: "switch1:port14", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port15  = &ConcretePort{Desc: "switch1:port15", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port16  = &ConcretePort{Desc: "switch1:port16", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port17  = &ConcretePort{Desc: "switch1:port17", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port18  = &ConcretePort{Desc: "switch1:port18", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port19  = &ConcretePort{Desc: "switch1:port19", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port20  = &ConcretePort{Desc: "switch1:port20", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port21  = &ConcretePort{Desc: "switch1:port21", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port22  = &ConcretePort{Desc: "switch1:port22", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port23  = &ConcretePort{Desc: "switch1:port23", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port24  = &ConcretePort{Desc: "switch1:port24", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port25  = &ConcretePort{Desc: "switch1:port25", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port26  = &ConcretePort{Desc: "switch1:port26", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port27  = &ConcretePort{Desc: "switch1:port27", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port28  = &ConcretePort{Desc: "switch1:port28", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port29  = &ConcretePort{Desc: "switch1:port29", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port30  = &ConcretePort{Desc: "switch1:port30", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port31  = &ConcretePort{Desc: "switch1:port31", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port32  = &ConcretePort{Desc: "switch1:port32", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port33  = &ConcretePort{Desc: "switch1:port33", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port34  = &ConcretePort{Desc: "switch1:port34", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port35  = &ConcretePort{Desc: "switch1:port35", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port36  = &ConcretePort{Desc: "switch1:port36", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port37  = &ConcretePort{Desc: "switch1:port37", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port38  = &ConcretePort{Desc: "switch1:port38", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port39  = &ConcretePort{Desc: "switch1:port39", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port40  = &ConcretePort{Desc: "switch1:port40", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port41  = &ConcretePort{Desc: "switch1:port41", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port42  = &ConcretePort{Desc: "switch1:port42", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port43  = &ConcretePort{Desc: "switch1:port43", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port44  = &ConcretePort{Desc: "switch1:port44", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port45  = &ConcretePort{Desc: "switch1:port45", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port46  = &ConcretePort{Desc: "switch1:port46", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port47  = &ConcretePort{Desc: "switch1:port47", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port48  = &ConcretePort{Desc: "switch1:port48", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port49  = &ConcretePort{Desc: "switch1:port49", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port50  = &ConcretePort{Desc: "switch1:port50", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port51  = &ConcretePort{Desc: "switch1:port51", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port52  = &ConcretePort{Desc: "switch1:port52", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port53  = &ConcretePort{Desc: "switch1:port53", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port54  = &ConcretePort{Desc: "switch1:port54", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port55  = &ConcretePort{Desc: "switch1:port55", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port56  = &ConcretePort{Desc: "switch1:port56", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port57  = &ConcretePort{Desc: "switch1:port57", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port58  = &ConcretePort{Desc: "switch1:port58", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port59  = &ConcretePort{Desc: "switch1:port59", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port60  = &ConcretePort{Desc: "switch1:port60", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port61  = &ConcretePort{Desc: "switch1:port61", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port62  = &ConcretePort{Desc: "switch1:port62", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port63  = &ConcretePort{Desc: "switch1:port63", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port64  = &ConcretePort{Desc: "switch1:port64", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port65  = &ConcretePort{Desc: "switch1:port65", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port66  = &ConcretePort{Desc: "switch1:port66", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port67  = &ConcretePort{Desc: "switch1:port67", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port68  = &ConcretePort{Desc: "switch1:port68", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port69  = &ConcretePort{Desc: "switch1:port69", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port70  = &ConcretePort{Desc: "switch1:port70", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port71  = &ConcretePort{Desc: "switch1:port71", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port72  = &ConcretePort{Desc: "switch1:port72", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port73  = &ConcretePort{Desc: "switch1:port73", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port74  = &ConcretePort{Desc: "switch1:port74", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port75  = &ConcretePort{Desc: "switch1:port75", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port76  = &ConcretePort{Desc: "switch1:port76", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port77  = &ConcretePort{Desc: "switch1:port77", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port78  = &ConcretePort{Desc: "switch1:port78", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port79  = &ConcretePort{Desc: "switch1:port79", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port80  = &ConcretePort{Desc: "switch1:port80", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port81  = &ConcretePort{Desc: "switch1:port81", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port82  = &ConcretePort{Desc: "switch1:port82", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port83  = &ConcretePort{Desc: "switch1:port83", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port84  = &ConcretePort{Desc: "switch1:port84", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port85  = &ConcretePort{Desc: "switch1:port85", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port86  = &ConcretePort{Desc: "switch1:port86", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port87  = &ConcretePort{Desc: "switch1:port87", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port88  = &ConcretePort{Desc: "switch1:port88", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port89  = &ConcretePort{Desc: "switch1:port89", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port90  = &ConcretePort{Desc: "switch1:port90", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port91  = &ConcretePort{Desc: "switch1:port91", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port92  = &ConcretePort{Desc: "switch1:port92", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port93  = &ConcretePort{Desc: "switch1:port93", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port94  = &ConcretePort{Desc: "switch1:port94", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port95  = &ConcretePort{Desc: "switch1:port95", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port96  = &ConcretePort{Desc: "switch1:port96", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port97  = &ConcretePort{Desc: "switch1:port97", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port98  = &ConcretePort{Desc: "switch1:port98", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port99  = &ConcretePort{Desc: "switch1:port99", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port100 = &ConcretePort{Desc: "switch1:port100", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port101 = &ConcretePort{Desc: "switch1:port101", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port102 = &ConcretePort{Desc: "switch1:port102", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port103 = &ConcretePort{Desc: "switch1:port103", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port104 = &ConcretePort{Desc: "switch1:port104", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port105 = &ConcretePort{Desc: "switch1:port105", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port106 = &ConcretePort{Desc: "switch1:port106", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port107 = &ConcretePort{Desc: "switch1:port107", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port108 = &ConcretePort{Desc: "switch1:port108", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port109 = &ConcretePort{Desc: "switch1:port109", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port110 = &ConcretePort{Desc: "switch1:port110", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port111 = &ConcretePort{Desc: "switch1:port111", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port112 = &ConcretePort{Desc: "switch1:port112", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port113 = &ConcretePort{Desc: "switch1:port113", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port114 = &ConcretePort{Desc: "switch1:port114", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port115 = &ConcretePort{Desc: "switch1:port115", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port116 = &ConcretePort{Desc: "switch1:port116", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port117 = &ConcretePort{Desc: "switch1:port117", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port118 = &ConcretePort{Desc: "switch1:port118", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port119 = &ConcretePort{Desc: "switch1:port119", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port120 = &ConcretePort{Desc: "switch1:port120", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port121 = &ConcretePort{Desc: "switch1:port121", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port122 = &ConcretePort{Desc: "switch1:port122", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port123 = &ConcretePort{Desc: "switch1:port123", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port124 = &ConcretePort{Desc: "switch1:port124", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port125 = &ConcretePort{Desc: "switch1:port125", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port126 = &ConcretePort{Desc: "switch1:port126", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port127 = &ConcretePort{Desc: "switch1:port127", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port128 = &ConcretePort{Desc: "switch1:port128", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port129 = &ConcretePort{Desc: "switch1:port129", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port130 = &ConcretePort{Desc: "switch1:port130", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port131 = &ConcretePort{Desc: "switch1:port131", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port132 = &ConcretePort{Desc: "switch1:port132", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port133 = &ConcretePort{Desc: "switch1:port133", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port134 = &ConcretePort{Desc: "switch1:port134", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port135 = &ConcretePort{Desc: "switch1:port135", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port136 = &ConcretePort{Desc: "switch1:port136", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port137 = &ConcretePort{Desc: "switch1:port137", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port138 = &ConcretePort{Desc: "switch1:port138", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port139 = &ConcretePort{Desc: "switch1:port139", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port140 = &ConcretePort{Desc: "switch1:port140", Attrs: map[string]string{"speed": "S_400G"}}
	switch1port141 = &ConcretePort{Desc: "switch1:port141", Attrs: map[string]string{"speed": "S_400G"}}

	switch1 = &ConcreteNode{Desc: "switch1", Ports: []*ConcretePort{
		switch1port1,
		switch1port2,
		switch1port3,
		switch1port4,
		switch1port5,
		switch1port6,
		switch1port7,
		switch1port8,
		switch1port9,
		switch1port10,
		switch1port11,
		switch1port12,
		switch1port13,
		switch1port14,
		switch1port15,
		switch1port16,
		switch1port17,
		switch1port18,
		switch1port19,
		switch1port20,
		switch1port21,
		switch1port22,
		switch1port23,
		switch1port24,
		switch1port25,
		switch1port26,
		switch1port27,
		switch1port28,
		switch1port29,
		switch1port30,
		switch1port31,
		switch1port32,
		switch1port33,
		switch1port34,
		switch1port35,
		switch1port36,
		switch1port37,
		switch1port38,
		switch1port39,
		switch1port40,
		switch1port41,
		switch1port42,
		switch1port43,
		switch1port44,
		switch1port45,
		switch1port46,
		switch1port47,
		switch1port48,
		switch1port49,
		switch1port50,
		switch1port51,
		switch1port52,
		switch1port53,
		switch1port54,
		switch1port55,
		switch1port56,
		switch1port57,
		switch1port58,
		switch1port59,
		switch1port60,
		switch1port61,
		switch1port62,
		switch1port63,
		switch1port64,
		switch1port65,
		switch1port66,
		switch1port67,
		switch1port68,
		switch1port69,
		switch1port70,
		switch1port71,
		switch1port72,
		switch1port73,
		switch1port74,
		switch1port75,
		switch1port76,
		switch1port77,
		switch1port78,
		switch1port79,
		switch1port80,
		switch1port81,
		switch1port82,
		switch1port83,
		switch1port84,
		switch1port85,
		switch1port86,
		switch1port87,
		switch1port88,
		switch1port89,
		switch1port90,
		switch1port91,
		switch1port92,
		switch1port93,
		switch1port94,
		switch1port95,
		switch1port96,
		switch1port97,
		switch1port98,
		switch1port99,
		switch1port100,
		switch1port101,
		switch1port102,
		switch1port103,
		switch1port104,
		switch1port105,
		switch1port106,
		switch1port107,
		switch1port108,
		switch1port109,
		switch1port110,
		switch1port111,
		switch1port112,
		switch1port113,
		switch1port114,
		switch1port115,
		switch1port116,
		switch1port117,
		switch1port118,
		switch1port119,
		switch1port120,
		switch1port121,
		switch1port122,
		switch1port123,
		switch1port124,
		switch1port125,
		switch1port126,
		switch1port127,
		switch1port128,
		switch1port129,
		switch1port130,
		switch1port131,
		switch1port132,
		switch1port133,
		switch1port134,
		switch1port135,
		switch1port136,
		switch1port137,
		switch1port138,
		switch1port139,
		switch1port140,
		switch1port141}, Attrs: map[string]string{"role": "L1S", "name": "sw1"}}

	switch2port1   = &ConcretePort{Desc: "switch2:port1", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port2   = &ConcretePort{Desc: "switch2:port2", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port3   = &ConcretePort{Desc: "switch2:port3", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port4   = &ConcretePort{Desc: "switch2:port4", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port5   = &ConcretePort{Desc: "switch2:port5", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port6   = &ConcretePort{Desc: "switch2:port6", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port7   = &ConcretePort{Desc: "switch2:port7", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port8   = &ConcretePort{Desc: "switch2:port8", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port9   = &ConcretePort{Desc: "switch2:port9", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port10  = &ConcretePort{Desc: "switch2:port10", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port11  = &ConcretePort{Desc: "switch2:port11", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port12  = &ConcretePort{Desc: "switch2:port12", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port13  = &ConcretePort{Desc: "switch2:port13", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port14  = &ConcretePort{Desc: "switch2:port14", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port15  = &ConcretePort{Desc: "switch2:port15", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port16  = &ConcretePort{Desc: "switch2:port16", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port17  = &ConcretePort{Desc: "switch2:port17", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port18  = &ConcretePort{Desc: "switch2:port18", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port19  = &ConcretePort{Desc: "switch2:port19", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port20  = &ConcretePort{Desc: "switch2:port20", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port21  = &ConcretePort{Desc: "switch2:port21", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port22  = &ConcretePort{Desc: "switch2:port22", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port23  = &ConcretePort{Desc: "switch2:port23", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port24  = &ConcretePort{Desc: "switch2:port24", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port25  = &ConcretePort{Desc: "switch2:port25", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port26  = &ConcretePort{Desc: "switch2:port26", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port27  = &ConcretePort{Desc: "switch2:port27", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port28  = &ConcretePort{Desc: "switch2:port28", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port29  = &ConcretePort{Desc: "switch2:port29", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port30  = &ConcretePort{Desc: "switch2:port30", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port31  = &ConcretePort{Desc: "switch2:port31", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port32  = &ConcretePort{Desc: "switch2:port32", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port33  = &ConcretePort{Desc: "switch2:port33", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port34  = &ConcretePort{Desc: "switch2:port34", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port35  = &ConcretePort{Desc: "switch2:port35", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port36  = &ConcretePort{Desc: "switch2:port36", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port37  = &ConcretePort{Desc: "switch2:port37", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port38  = &ConcretePort{Desc: "switch2:port38", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port39  = &ConcretePort{Desc: "switch2:port39", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port40  = &ConcretePort{Desc: "switch2:port40", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port41  = &ConcretePort{Desc: "switch2:port41", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port42  = &ConcretePort{Desc: "switch2:port42", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port43  = &ConcretePort{Desc: "switch2:port43", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port44  = &ConcretePort{Desc: "switch2:port44", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port45  = &ConcretePort{Desc: "switch2:port45", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port46  = &ConcretePort{Desc: "switch2:port46", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port47  = &ConcretePort{Desc: "switch2:port47", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port48  = &ConcretePort{Desc: "switch2:port48", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port49  = &ConcretePort{Desc: "switch2:port49", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port50  = &ConcretePort{Desc: "switch2:port50", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port51  = &ConcretePort{Desc: "switch2:port51", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port52  = &ConcretePort{Desc: "switch2:port52", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port53  = &ConcretePort{Desc: "switch2:port53", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port54  = &ConcretePort{Desc: "switch2:port54", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port55  = &ConcretePort{Desc: "switch2:port55", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port56  = &ConcretePort{Desc: "switch2:port56", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port57  = &ConcretePort{Desc: "switch2:port57", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port58  = &ConcretePort{Desc: "switch2:port58", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port59  = &ConcretePort{Desc: "switch2:port59", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port60  = &ConcretePort{Desc: "switch2:port60", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port61  = &ConcretePort{Desc: "switch2:port61", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port62  = &ConcretePort{Desc: "switch2:port62", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port63  = &ConcretePort{Desc: "switch2:port63", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port64  = &ConcretePort{Desc: "switch2:port64", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port65  = &ConcretePort{Desc: "switch2:port65", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port66  = &ConcretePort{Desc: "switch2:port66", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port67  = &ConcretePort{Desc: "switch2:port67", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port68  = &ConcretePort{Desc: "switch2:port68", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port69  = &ConcretePort{Desc: "switch2:port69", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port70  = &ConcretePort{Desc: "switch2:port70", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port71  = &ConcretePort{Desc: "switch2:port71", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port72  = &ConcretePort{Desc: "switch2:port72", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port73  = &ConcretePort{Desc: "switch2:port73", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port74  = &ConcretePort{Desc: "switch2:port74", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port75  = &ConcretePort{Desc: "switch2:port75", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port76  = &ConcretePort{Desc: "switch2:port76", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port77  = &ConcretePort{Desc: "switch2:port77", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port78  = &ConcretePort{Desc: "switch2:port78", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port79  = &ConcretePort{Desc: "switch2:port79", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port80  = &ConcretePort{Desc: "switch2:port80", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port81  = &ConcretePort{Desc: "switch2:port81", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port82  = &ConcretePort{Desc: "switch2:port82", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port83  = &ConcretePort{Desc: "switch2:port83", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port84  = &ConcretePort{Desc: "switch2:port84", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port85  = &ConcretePort{Desc: "switch2:port85", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port86  = &ConcretePort{Desc: "switch2:port86", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port87  = &ConcretePort{Desc: "switch2:port87", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port88  = &ConcretePort{Desc: "switch2:port88", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port89  = &ConcretePort{Desc: "switch2:port89", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port90  = &ConcretePort{Desc: "switch2:port90", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port91  = &ConcretePort{Desc: "switch2:port91", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port92  = &ConcretePort{Desc: "switch2:port92", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port93  = &ConcretePort{Desc: "switch2:port93", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port94  = &ConcretePort{Desc: "switch2:port94", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port95  = &ConcretePort{Desc: "switch2:port95", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port96  = &ConcretePort{Desc: "switch2:port96", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port97  = &ConcretePort{Desc: "switch2:port97", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port98  = &ConcretePort{Desc: "switch2:port98", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port99  = &ConcretePort{Desc: "switch2:port99", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port100 = &ConcretePort{Desc: "switch2:port100", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port101 = &ConcretePort{Desc: "switch2:port101", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port102 = &ConcretePort{Desc: "switch2:port102", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port103 = &ConcretePort{Desc: "switch2:port103", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port104 = &ConcretePort{Desc: "switch2:port104", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port105 = &ConcretePort{Desc: "switch2:port105", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port106 = &ConcretePort{Desc: "switch2:port106", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port107 = &ConcretePort{Desc: "switch2:port107", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port108 = &ConcretePort{Desc: "switch2:port108", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port109 = &ConcretePort{Desc: "switch2:port109", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port110 = &ConcretePort{Desc: "switch2:port110", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port111 = &ConcretePort{Desc: "switch2:port111", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port112 = &ConcretePort{Desc: "switch2:port112", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port113 = &ConcretePort{Desc: "switch2:port113", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port114 = &ConcretePort{Desc: "switch2:port114", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port115 = &ConcretePort{Desc: "switch2:port115", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port116 = &ConcretePort{Desc: "switch2:port116", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port117 = &ConcretePort{Desc: "switch2:port117", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port118 = &ConcretePort{Desc: "switch2:port118", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port119 = &ConcretePort{Desc: "switch2:port119", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port120 = &ConcretePort{Desc: "switch2:port120", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port121 = &ConcretePort{Desc: "switch2:port121", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port122 = &ConcretePort{Desc: "switch2:port122", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port123 = &ConcretePort{Desc: "switch2:port123", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port124 = &ConcretePort{Desc: "switch2:port124", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port125 = &ConcretePort{Desc: "switch2:port125", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port126 = &ConcretePort{Desc: "switch2:port126", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port127 = &ConcretePort{Desc: "switch2:port127", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port128 = &ConcretePort{Desc: "switch2:port128", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port129 = &ConcretePort{Desc: "switch2:port129", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port130 = &ConcretePort{Desc: "switch2:port130", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port131 = &ConcretePort{Desc: "switch2:port131", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port132 = &ConcretePort{Desc: "switch2:port132", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port133 = &ConcretePort{Desc: "switch2:port133", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port134 = &ConcretePort{Desc: "switch2:port134", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port135 = &ConcretePort{Desc: "switch2:port135", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port136 = &ConcretePort{Desc: "switch2:port136", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port137 = &ConcretePort{Desc: "switch2:port137", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port138 = &ConcretePort{Desc: "switch2:port138", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port139 = &ConcretePort{Desc: "switch2:port139", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port140 = &ConcretePort{Desc: "switch2:port140", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port141 = &ConcretePort{Desc: "switch2:port141", Attrs: map[string]string{"speed": "S_400G"}}
	switch2port142 = &ConcretePort{Desc: "switch2:port142", Attrs: map[string]string{"speed": "S_400G"}}

	switch2 = &ConcreteNode{Desc: "switch2", Ports: []*ConcretePort{
		switch2port1,
		switch2port2,
		switch2port3,
		switch2port4,
		switch2port5,
		switch2port6,
		switch2port7,
		switch2port8,
		switch2port9,
		switch2port10,
		switch2port11,
		switch2port12,
		switch2port13,
		switch2port14,
		switch2port15,
		switch2port16,
		switch2port17,
		switch2port18,
		switch2port19,
		switch2port20,
		switch2port21,
		switch2port22,
		switch2port23,
		switch2port24,
		switch2port25,
		switch2port26,
		switch2port27,
		switch2port28,
		switch2port29,
		switch2port30,
		switch2port31,
		switch2port32,
		switch2port33,
		switch2port34,
		switch2port35,
		switch2port36,
		switch2port37,
		switch2port38,
		switch2port39,
		switch2port40,
		switch2port41,
		switch2port42,
		switch2port43,
		switch2port44,
		switch2port45,
		switch2port46,
		switch2port47,
		switch2port48,
		switch2port49,
		switch2port50,
		switch2port51,
		switch2port52,
		switch2port53,
		switch2port54,
		switch2port55,
		switch2port56,
		switch2port57,
		switch2port58,
		switch2port59,
		switch2port60,
		switch2port61,
		switch2port62,
		switch2port63,
		switch2port64,
		switch2port65,
		switch2port66,
		switch2port67,
		switch2port68,
		switch2port69,
		switch2port70,
		switch2port71,
		switch2port72,
		switch2port73,
		switch2port74,
		switch2port75,
		switch2port76,
		switch2port77,
		switch2port78,
		switch2port79,
		switch2port80,
		switch2port81,
		switch2port82,
		switch2port83,
		switch2port84,
		switch2port85,
		switch2port86,
		switch2port87,
		switch2port88,
		switch2port89,
		switch2port90,
		switch2port91,
		switch2port92,
		switch2port93,
		switch2port94,
		switch2port95,
		switch2port96,
		switch2port97,
		switch2port98,
		switch2port99,
		switch2port100,
		switch2port101,
		switch2port102,
		switch2port103,
		switch2port104,
		switch2port105,
		switch2port106,
		switch2port107,
		switch2port108,
		switch2port109,
		switch2port110,
		switch2port111,
		switch2port112,
		switch2port113,
		switch2port114,
		switch2port115,
		switch2port116,
		switch2port117,
		switch2port118,
		switch2port119,
		switch2port120,
		switch2port121,
		switch2port122,
		switch2port123,
		switch2port124,
		switch2port125,
		switch2port126,
		switch2port127,
		switch2port128,
		switch2port129,
		switch2port130,
		switch2port131,
		switch2port132,
		switch2port133,
		switch2port134,
		switch2port135,
		switch2port136,
		switch2port137,
		switch2port138,
		switch2port139,
		switch2port140,
		switch2port141,
		switch2port142}, Attrs: map[string]string{"role": "L1S", "name": "sw2"}}

	switch3port1   = &ConcretePort{Desc: "switch3:port1", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port2   = &ConcretePort{Desc: "switch3:port2", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port3   = &ConcretePort{Desc: "switch3:port3", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port4   = &ConcretePort{Desc: "switch3:port4", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port5   = &ConcretePort{Desc: "switch3:port5", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port6   = &ConcretePort{Desc: "switch3:port6", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port7   = &ConcretePort{Desc: "switch3:port7", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port8   = &ConcretePort{Desc: "switch3:port8", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port9   = &ConcretePort{Desc: "switch3:port9", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port10  = &ConcretePort{Desc: "switch3:port10", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port11  = &ConcretePort{Desc: "switch3:port11", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port12  = &ConcretePort{Desc: "switch3:port12", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port13  = &ConcretePort{Desc: "switch3:port13", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port14  = &ConcretePort{Desc: "switch3:port14", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port15  = &ConcretePort{Desc: "switch3:port15", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port16  = &ConcretePort{Desc: "switch3:port16", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port17  = &ConcretePort{Desc: "switch3:port17", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port18  = &ConcretePort{Desc: "switch3:port18", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port19  = &ConcretePort{Desc: "switch3:port19", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port20  = &ConcretePort{Desc: "switch3:port20", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port21  = &ConcretePort{Desc: "switch3:port21", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port22  = &ConcretePort{Desc: "switch3:port22", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port23  = &ConcretePort{Desc: "switch3:port23", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port24  = &ConcretePort{Desc: "switch3:port24", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port25  = &ConcretePort{Desc: "switch3:port25", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port26  = &ConcretePort{Desc: "switch3:port26", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port27  = &ConcretePort{Desc: "switch3:port27", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port28  = &ConcretePort{Desc: "switch3:port28", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port29  = &ConcretePort{Desc: "switch3:port29", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port30  = &ConcretePort{Desc: "switch3:port30", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port31  = &ConcretePort{Desc: "switch3:port31", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port32  = &ConcretePort{Desc: "switch3:port32", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port33  = &ConcretePort{Desc: "switch3:port33", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port34  = &ConcretePort{Desc: "switch3:port34", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port35  = &ConcretePort{Desc: "switch3:port35", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port36  = &ConcretePort{Desc: "switch3:port36", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port37  = &ConcretePort{Desc: "switch3:port37", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port38  = &ConcretePort{Desc: "switch3:port38", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port39  = &ConcretePort{Desc: "switch3:port39", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port40  = &ConcretePort{Desc: "switch3:port40", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port41  = &ConcretePort{Desc: "switch3:port41", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port42  = &ConcretePort{Desc: "switch3:port42", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port43  = &ConcretePort{Desc: "switch3:port43", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port44  = &ConcretePort{Desc: "switch3:port44", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port45  = &ConcretePort{Desc: "switch3:port45", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port46  = &ConcretePort{Desc: "switch3:port46", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port47  = &ConcretePort{Desc: "switch3:port47", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port48  = &ConcretePort{Desc: "switch3:port48", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port49  = &ConcretePort{Desc: "switch3:port49", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port50  = &ConcretePort{Desc: "switch3:port50", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port51  = &ConcretePort{Desc: "switch3:port51", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port52  = &ConcretePort{Desc: "switch3:port52", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port53  = &ConcretePort{Desc: "switch3:port53", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port54  = &ConcretePort{Desc: "switch3:port54", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port55  = &ConcretePort{Desc: "switch3:port55", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port56  = &ConcretePort{Desc: "switch3:port56", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port57  = &ConcretePort{Desc: "switch3:port57", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port58  = &ConcretePort{Desc: "switch3:port58", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port59  = &ConcretePort{Desc: "switch3:port59", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port60  = &ConcretePort{Desc: "switch3:port60", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port61  = &ConcretePort{Desc: "switch3:port61", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port62  = &ConcretePort{Desc: "switch3:port62", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port63  = &ConcretePort{Desc: "switch3:port63", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port64  = &ConcretePort{Desc: "switch3:port64", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port65  = &ConcretePort{Desc: "switch3:port65", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port66  = &ConcretePort{Desc: "switch3:port66", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port67  = &ConcretePort{Desc: "switch3:port67", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port68  = &ConcretePort{Desc: "switch3:port68", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port69  = &ConcretePort{Desc: "switch3:port69", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port70  = &ConcretePort{Desc: "switch3:port70", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port71  = &ConcretePort{Desc: "switch3:port71", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port72  = &ConcretePort{Desc: "switch3:port72", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port73  = &ConcretePort{Desc: "switch3:port73", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port74  = &ConcretePort{Desc: "switch3:port74", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port75  = &ConcretePort{Desc: "switch3:port75", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port76  = &ConcretePort{Desc: "switch3:port76", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port77  = &ConcretePort{Desc: "switch3:port77", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port78  = &ConcretePort{Desc: "switch3:port78", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port79  = &ConcretePort{Desc: "switch3:port79", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port80  = &ConcretePort{Desc: "switch3:port80", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port81  = &ConcretePort{Desc: "switch3:port81", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port82  = &ConcretePort{Desc: "switch3:port82", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port83  = &ConcretePort{Desc: "switch3:port83", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port84  = &ConcretePort{Desc: "switch3:port84", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port85  = &ConcretePort{Desc: "switch3:port85", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port86  = &ConcretePort{Desc: "switch3:port86", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port87  = &ConcretePort{Desc: "switch3:port87", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port88  = &ConcretePort{Desc: "switch3:port88", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port89  = &ConcretePort{Desc: "switch3:port89", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port90  = &ConcretePort{Desc: "switch3:port90", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port91  = &ConcretePort{Desc: "switch3:port91", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port92  = &ConcretePort{Desc: "switch3:port92", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port93  = &ConcretePort{Desc: "switch3:port93", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port94  = &ConcretePort{Desc: "switch3:port94", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port95  = &ConcretePort{Desc: "switch3:port95", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port96  = &ConcretePort{Desc: "switch3:port96", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port97  = &ConcretePort{Desc: "switch3:port97", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port98  = &ConcretePort{Desc: "switch3:port98", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port99  = &ConcretePort{Desc: "switch3:port99", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port100 = &ConcretePort{Desc: "switch3:port100", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port101 = &ConcretePort{Desc: "switch3:port101", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port102 = &ConcretePort{Desc: "switch3:port102", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port103 = &ConcretePort{Desc: "switch3:port103", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port104 = &ConcretePort{Desc: "switch3:port104", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port105 = &ConcretePort{Desc: "switch3:port105", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port106 = &ConcretePort{Desc: "switch3:port106", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port107 = &ConcretePort{Desc: "switch3:port107", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port108 = &ConcretePort{Desc: "switch3:port108", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port109 = &ConcretePort{Desc: "switch3:port109", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port110 = &ConcretePort{Desc: "switch3:port110", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port111 = &ConcretePort{Desc: "switch3:port111", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port112 = &ConcretePort{Desc: "switch3:port112", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port113 = &ConcretePort{Desc: "switch3:port113", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port114 = &ConcretePort{Desc: "switch3:port114", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port115 = &ConcretePort{Desc: "switch3:port115", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port116 = &ConcretePort{Desc: "switch3:port116", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port117 = &ConcretePort{Desc: "switch3:port117", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port118 = &ConcretePort{Desc: "switch3:port118", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port119 = &ConcretePort{Desc: "switch3:port119", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port120 = &ConcretePort{Desc: "switch3:port120", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port121 = &ConcretePort{Desc: "switch3:port121", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port122 = &ConcretePort{Desc: "switch3:port122", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port123 = &ConcretePort{Desc: "switch3:port123", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port124 = &ConcretePort{Desc: "switch3:port124", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port125 = &ConcretePort{Desc: "switch3:port125", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port126 = &ConcretePort{Desc: "switch3:port126", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port127 = &ConcretePort{Desc: "switch3:port127", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port128 = &ConcretePort{Desc: "switch3:port128", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port129 = &ConcretePort{Desc: "switch3:port129", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port130 = &ConcretePort{Desc: "switch3:port130", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port131 = &ConcretePort{Desc: "switch3:port131", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port132 = &ConcretePort{Desc: "switch3:port132", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port133 = &ConcretePort{Desc: "switch3:port133", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port134 = &ConcretePort{Desc: "switch3:port134", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port135 = &ConcretePort{Desc: "switch3:port135", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port136 = &ConcretePort{Desc: "switch3:port136", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port137 = &ConcretePort{Desc: "switch3:port137", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port138 = &ConcretePort{Desc: "switch3:port138", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port139 = &ConcretePort{Desc: "switch3:port139", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port140 = &ConcretePort{Desc: "switch3:port140", Attrs: map[string]string{"speed": "S_400G"}}
	switch3port141 = &ConcretePort{Desc: "switch3:port141", Attrs: map[string]string{"speed": "S_400G"}}

	switch3 = &ConcreteNode{Desc: "switch3", Ports: []*ConcretePort{
		switch3port1,
		switch3port2,
		switch3port3,
		switch3port4,
		switch3port5,
		switch3port6,
		switch3port7,
		switch3port8,
		switch3port9,
		switch3port10,
		switch3port11,
		switch3port12,
		switch3port13,
		switch3port14,
		switch3port15,
		switch3port16,
		switch3port17,
		switch3port18,
		switch3port19,
		switch3port20,
		switch3port21,
		switch3port22,
		switch3port23,
		switch3port24,
		switch3port25,
		switch3port26,
		switch3port27,
		switch3port28,
		switch3port29,
		switch3port30,
		switch3port31,
		switch3port32,
		switch3port33,
		switch3port34,
		switch3port35,
		switch3port36,
		switch3port37,
		switch3port38,
		switch3port39,
		switch3port40,
		switch3port41,
		switch3port42,
		switch3port43,
		switch3port44,
		switch3port45,
		switch3port46,
		switch3port47,
		switch3port48,
		switch3port49,
		switch3port50,
		switch3port51,
		switch3port52,
		switch3port53,
		switch3port54,
		switch3port55,
		switch3port56,
		switch3port57,
		switch3port58,
		switch3port59,
		switch3port60,
		switch3port61,
		switch3port62,
		switch3port63,
		switch3port64,
		switch3port65,
		switch3port66,
		switch3port67,
		switch3port68,
		switch3port69,
		switch3port70,
		switch3port71,
		switch3port72,
		switch3port73,
		switch3port74,
		switch3port75,
		switch3port76,
		switch3port77,
		switch3port78,
		switch3port79,
		switch3port80,
		switch3port81,
		switch3port82,
		switch3port83,
		switch3port84,
		switch3port85,
		switch3port86,
		switch3port87,
		switch3port88,
		switch3port89,
		switch3port90,
		switch3port91,
		switch3port92,
		switch3port93,
		switch3port94,
		switch3port95,
		switch3port96,
		switch3port97,
		switch3port98,
		switch3port99,
		switch3port100,
		switch3port101,
		switch3port102,
		switch3port103,
		switch3port104,
		switch3port105,
		switch3port106,
		switch3port107,
		switch3port108,
		switch3port109,
		switch3port110,
		switch3port111,
		switch3port112,
		switch3port113,
		switch3port114,
		switch3port115,
		switch3port116,
		switch3port117,
		switch3port118,
		switch3port119,
		switch3port120,
		switch3port121,
		switch3port122,
		switch3port123,
		switch3port124,
		switch3port125,
		switch3port126,
		switch3port127,
		switch3port128,
		switch3port129,
		switch3port130,
		switch3port131,
		switch3port132,
		switch3port133,
		switch3port134,
		switch3port135,
		switch3port136,
		switch3port137,
		switch3port138,
		switch3port139,
		switch3port140,
		switch3port141}, Attrs: map[string]string{"role": "L1S", "name": "sw3"}}
)

var superGraphT = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		dut1,
		ate1,
		dut2,
		ate2,
		dut3,
		ate3,
		dut4,
		ate4,
		switch1, switch2, switch3},
	Edges: []*ConcreteEdge{
		{Src: dut1port1, Dst: ate1port1},
		{Src: dut1port2, Dst: ate1port2},
		{Src: dut1port3, Dst: ate1port3},
		{Src: dut1port4, Dst: ate1port4},
		{Src: dut1port5, Dst: ate1port5},
		{Src: dut1port6, Dst: ate1port6},
		{Src: dut1port7, Dst: ate1port7},
		{Src: dut1port8, Dst: ate1port8},
		{Src: dut1port9, Dst: ate1port9},
		{Src: dut1port10, Dst: ate1port10},
		{Src: dut1port11, Dst: ate1port11},
		{Src: dut1port12, Dst: ate1port12},
		{Src: dut1port13, Dst: ate1port13},
		{Src: dut1port14, Dst: ate1port14},
		{Src: dut1port15, Dst: ate1port15},
		{Src: dut1port16, Dst: ate1port16},
		{Src: dut1port17, Dst: ate1port17},
		{Src: dut1port18, Dst: ate1port18},
		{Src: dut1port19, Dst: ate1port19},
		{Src: dut1port20, Dst: ate1port20},
		{Src: dut1port21, Dst: ate1port21},
		{Src: dut1port22, Dst: ate1port22},
		{Src: dut1port23, Dst: ate1port23},
		{Src: dut1port24, Dst: ate1port24},
		{Src: dut1port25, Dst: ate1port25},
		{Src: dut1port26, Dst: ate1port26},
		{Src: dut1port27, Dst: ate1port27},
		{Src: dut1port28, Dst: ate1port28},
		{Src: dut1port29, Dst: ate1port29},
		{Src: dut1port30, Dst: ate1port30},
		{Src: dut1port31, Dst: ate1port31},
		{Src: dut1port32, Dst: ate1port32},
		{Src: dut1port33, Dst: ate1port33},
		{Src: dut1port34, Dst: ate1port34},
		{Src: dut1port35, Dst: ate1port35},

		{Src: dut1port36, Dst: switch1port1},
		{Src: dut1port37, Dst: switch1port2},
		{Src: dut1port38, Dst: switch1port3},
		{Src: dut1port39, Dst: switch1port4},
		{Src: dut1port40, Dst: switch1port5},
		{Src: dut1port41, Dst: switch1port6},
		{Src: dut1port42, Dst: switch1port7},
		{Src: dut1port43, Dst: switch1port8},
		{Src: dut1port44, Dst: switch1port9},
		{Src: dut1port45, Dst: switch1port10},
		{Src: dut1port46, Dst: switch1port11},
		{Src: dut1port47, Dst: switch1port12},
		{Src: dut1port48, Dst: switch1port13},
		{Src: dut1port49, Dst: switch1port14},
		{Src: dut1port50, Dst: switch1port15},
		{Src: dut1port51, Dst: switch1port16},
		{Src: dut1port52, Dst: switch1port17},
		{Src: dut1port53, Dst: switch1port18},
		{Src: dut1port54, Dst: switch1port19},
		{Src: dut1port55, Dst: switch1port20},
		{Src: dut1port56, Dst: switch1port21},
		{Src: dut1port57, Dst: switch1port22},
		{Src: dut1port58, Dst: switch1port23},
		{Src: dut1port59, Dst: switch1port24},
		{Src: dut1port60, Dst: switch1port25},
		{Src: dut1port61, Dst: switch1port26},
		{Src: dut1port62, Dst: switch1port27},
		{Src: dut1port63, Dst: switch1port28},
		{Src: dut1port64, Dst: switch1port29},
		{Src: dut1port65, Dst: switch1port30},
		{Src: dut1port66, Dst: switch1port31},
		{Src: dut1port67, Dst: switch1port32},
		{Src: dut1port68, Dst: switch1port33},
		{Src: dut1port69, Dst: switch1port34},
		{Src: dut1port70, Dst: switch1port35},

		{Src: switch1port36, Dst: ate1port36},
		{Src: switch1port37, Dst: ate1port37},
		{Src: switch1port38, Dst: ate1port38},
		{Src: switch1port39, Dst: ate1port39},
		{Src: switch1port40, Dst: ate1port40},
		{Src: switch1port41, Dst: ate1port41},
		{Src: switch1port42, Dst: ate1port42},
		{Src: switch1port43, Dst: ate1port43},
		{Src: switch1port44, Dst: ate1port44},
		{Src: switch1port45, Dst: ate1port45},
		{Src: switch1port46, Dst: ate1port46},
		{Src: switch1port47, Dst: ate1port47},
		{Src: switch1port48, Dst: ate1port48},
		{Src: switch1port49, Dst: ate1port49},
		{Src: switch1port50, Dst: ate1port50},
		{Src: switch1port51, Dst: ate1port51},
		{Src: switch1port52, Dst: ate1port52},
		{Src: switch1port53, Dst: ate1port53},
		{Src: switch1port54, Dst: ate1port54},
		{Src: switch1port55, Dst: ate1port55},
		{Src: switch1port56, Dst: ate1port56},
		{Src: switch1port57, Dst: ate1port57},
		{Src: switch1port58, Dst: ate1port58},
		{Src: switch1port59, Dst: ate1port59},
		{Src: switch1port60, Dst: ate1port60},
		{Src: switch1port61, Dst: ate1port61},
		{Src: switch1port62, Dst: ate1port62},
		{Src: switch1port63, Dst: ate1port63},
		{Src: switch1port64, Dst: ate1port64},
		{Src: switch1port65, Dst: ate1port65},
		{Src: switch1port66, Dst: ate1port66},
		{Src: switch1port67, Dst: ate1port67},
		{Src: switch1port68, Dst: ate1port68},
		{Src: switch1port69, Dst: ate1port69},
		{Src: switch1port70, Dst: ate1port70},
		{Src: switch1port141, Dst: switch2port141},

		{Src: dut2port1, Dst: ate2port1},
		{Src: dut2port2, Dst: ate2port2},
		{Src: dut2port3, Dst: ate2port3},
		{Src: dut2port4, Dst: ate2port4},
		{Src: dut2port5, Dst: ate2port5},
		{Src: dut2port6, Dst: ate2port6},
		{Src: dut2port7, Dst: ate2port7},
		{Src: dut2port8, Dst: ate2port8},
		{Src: dut2port9, Dst: ate2port9},
		{Src: dut2port10, Dst: ate2port10},
		{Src: dut2port11, Dst: ate2port11},
		{Src: dut2port12, Dst: ate2port12},
		{Src: dut2port13, Dst: ate2port13},
		{Src: dut2port14, Dst: ate2port14},
		{Src: dut2port15, Dst: ate2port15},
		{Src: dut2port16, Dst: ate2port16},
		{Src: dut2port17, Dst: ate2port17},
		{Src: dut2port18, Dst: ate2port18},
		{Src: dut2port19, Dst: ate2port19},
		{Src: dut2port20, Dst: ate2port20},
		{Src: dut2port21, Dst: ate2port21},
		{Src: dut2port22, Dst: ate2port22},
		{Src: dut2port23, Dst: ate2port23},
		{Src: dut2port24, Dst: ate2port24},
		{Src: dut2port25, Dst: ate2port25},
		{Src: dut2port26, Dst: ate2port26},
		{Src: dut2port27, Dst: ate2port27},
		{Src: dut2port28, Dst: ate2port28},
		{Src: dut2port29, Dst: ate2port29},
		{Src: dut2port30, Dst: ate2port30},
		{Src: dut2port31, Dst: ate2port31},
		{Src: dut2port32, Dst: ate2port32},
		{Src: dut2port33, Dst: ate2port33},
		{Src: dut2port34, Dst: ate2port34},
		{Src: dut2port35, Dst: ate2port35},

		{Src: dut2port36, Dst: switch1port71},
		{Src: dut2port37, Dst: switch1port72},
		{Src: dut2port38, Dst: switch1port73},
		{Src: dut2port39, Dst: switch1port74},
		{Src: dut2port40, Dst: switch1port75},
		{Src: dut2port41, Dst: switch1port76},
		{Src: dut2port42, Dst: switch1port77},
		{Src: dut2port43, Dst: switch1port78},
		{Src: dut2port44, Dst: switch1port79},
		{Src: dut2port45, Dst: switch1port80},
		{Src: dut2port46, Dst: switch1port81},
		{Src: dut2port47, Dst: switch1port82},
		{Src: dut2port48, Dst: switch1port83},
		{Src: dut2port49, Dst: switch1port84},
		{Src: dut2port50, Dst: switch1port85},
		{Src: dut2port51, Dst: switch1port86},
		{Src: dut2port52, Dst: switch1port87},
		{Src: dut2port53, Dst: switch1port88},
		{Src: dut2port54, Dst: switch1port89},
		{Src: dut2port55, Dst: switch1port90},
		{Src: dut2port56, Dst: switch1port91},
		{Src: dut2port57, Dst: switch1port92},
		{Src: dut2port58, Dst: switch1port93},
		{Src: dut2port59, Dst: switch1port94},
		{Src: dut2port60, Dst: switch1port95},
		{Src: dut2port61, Dst: switch1port96},
		{Src: dut2port62, Dst: switch1port97},
		{Src: dut2port63, Dst: switch1port98},
		{Src: dut2port64, Dst: switch1port99},
		{Src: dut2port65, Dst: switch1port100},
		{Src: dut2port66, Dst: switch1port101},
		{Src: dut2port67, Dst: switch1port102},
		{Src: dut2port68, Dst: switch1port103},
		{Src: dut2port69, Dst: switch1port104},
		{Src: dut2port70, Dst: switch1port105},

		{Src: switch1port106, Dst: ate2port36},
		{Src: switch1port107, Dst: ate2port37},
		{Src: switch1port108, Dst: ate2port38},
		{Src: switch1port109, Dst: ate2port39},
		{Src: switch1port110, Dst: ate2port40},
		{Src: switch1port111, Dst: ate2port41},
		{Src: switch1port112, Dst: ate2port42},
		{Src: switch1port113, Dst: ate2port43},
		{Src: switch1port114, Dst: ate2port44},
		{Src: switch1port115, Dst: ate2port45},
		{Src: switch1port116, Dst: ate2port46},
		{Src: switch1port117, Dst: ate2port47},
		{Src: switch1port118, Dst: ate2port48},
		{Src: switch1port119, Dst: ate2port49},
		{Src: switch1port120, Dst: ate2port50},
		{Src: switch1port121, Dst: ate2port51},
		{Src: switch1port122, Dst: ate2port52},
		{Src: switch1port123, Dst: ate2port53},
		{Src: switch1port124, Dst: ate2port54},
		{Src: switch1port125, Dst: ate2port55},
		{Src: switch1port126, Dst: ate2port56},
		{Src: switch1port127, Dst: ate2port57},
		{Src: switch1port128, Dst: ate2port58},
		{Src: switch1port129, Dst: ate2port59},
		{Src: switch1port130, Dst: ate2port60},
		{Src: switch1port131, Dst: ate2port61},
		{Src: switch1port132, Dst: ate2port62},
		{Src: switch1port133, Dst: ate2port63},
		{Src: switch1port134, Dst: ate2port64},
		{Src: switch1port135, Dst: ate2port65},
		{Src: switch1port136, Dst: ate2port66},
		{Src: switch1port137, Dst: ate2port67},
		{Src: switch1port138, Dst: ate2port68},
		{Src: switch1port139, Dst: ate2port69},
		{Src: switch1port140, Dst: ate2port70},

		{Src: dut3port1, Dst: switch2port1},
		{Src: dut3port2, Dst: switch2port2},
		{Src: dut3port3, Dst: switch2port3},
		{Src: dut3port4, Dst: switch2port4},
		{Src: dut3port5, Dst: switch2port5},
		{Src: dut3port6, Dst: switch2port6},
		{Src: dut3port7, Dst: switch2port7},
		{Src: dut3port8, Dst: switch2port8},
		{Src: dut3port9, Dst: switch2port9},
		{Src: dut3port10, Dst: switch2port10},
		{Src: dut3port11, Dst: switch2port11},
		{Src: dut3port12, Dst: switch2port12},
		{Src: dut3port13, Dst: switch2port13},
		{Src: dut3port14, Dst: switch2port14},
		{Src: dut3port15, Dst: switch2port15},
		{Src: dut3port16, Dst: switch2port16},
		{Src: dut3port17, Dst: switch2port17},
		{Src: dut3port18, Dst: switch2port18},
		{Src: dut3port19, Dst: switch2port19},
		{Src: dut3port20, Dst: switch2port20},
		{Src: dut3port21, Dst: switch2port21},
		{Src: dut3port22, Dst: switch2port22},
		{Src: dut3port23, Dst: switch2port23},
		{Src: dut3port24, Dst: switch2port24},
		{Src: dut3port25, Dst: switch2port25},
		{Src: dut3port26, Dst: switch2port26},
		{Src: dut3port27, Dst: switch2port27},
		{Src: dut3port28, Dst: switch2port28},
		{Src: dut3port29, Dst: switch2port29},
		{Src: dut3port30, Dst: switch2port30},
		{Src: dut3port31, Dst: switch2port31},
		{Src: dut3port32, Dst: switch2port32},
		{Src: dut3port33, Dst: switch2port33},
		{Src: dut3port34, Dst: switch2port34},
		{Src: dut3port35, Dst: switch2port35},

		{Src: switch2port41, Dst: ate3port1},
		{Src: switch2port42, Dst: ate3port2},
		{Src: switch2port43, Dst: ate3port3},
		{Src: switch2port44, Dst: ate3port4},
		{Src: switch2port45, Dst: ate3port5},
		{Src: switch2port46, Dst: ate3port6},
		{Src: switch2port47, Dst: ate3port7},
		{Src: switch2port48, Dst: ate3port8},
		{Src: switch2port49, Dst: ate3port9},
		{Src: switch2port50, Dst: ate3port10},
		{Src: switch2port51, Dst: ate3port11},
		{Src: switch2port52, Dst: ate3port12},
		{Src: switch2port53, Dst: ate3port13},
		{Src: switch2port54, Dst: ate3port14},
		{Src: switch2port55, Dst: ate3port15},
		{Src: switch2port56, Dst: ate3port16},
		{Src: switch2port57, Dst: ate3port17},
		{Src: switch2port58, Dst: ate3port18},
		{Src: switch2port59, Dst: ate3port19},
		{Src: switch2port60, Dst: ate3port20},
		{Src: switch2port61, Dst: ate3port21},
		{Src: switch2port62, Dst: ate3port22},
		{Src: switch2port63, Dst: ate3port23},
		{Src: switch2port64, Dst: ate3port24},
		{Src: switch2port65, Dst: ate3port25},
		{Src: switch2port66, Dst: ate3port26},
		{Src: switch2port67, Dst: ate3port27},
		{Src: switch2port68, Dst: ate3port28},
		{Src: switch2port69, Dst: ate3port29},
		{Src: switch2port70, Dst: ate3port30},
		{Src: switch2port71, Dst: ate3port31},
		{Src: switch2port72, Dst: ate3port32},
		{Src: switch2port73, Dst: ate3port33},
		{Src: switch2port74, Dst: ate3port34},
		{Src: switch2port75, Dst: ate3port35},

		{Src: dut4port1, Dst: switch2port76},
		{Src: dut4port2, Dst: switch2port77},
		{Src: dut4port3, Dst: switch2port78},
		{Src: dut4port4, Dst: switch2port79},
		{Src: dut4port5, Dst: switch2port80},
		{Src: dut4port6, Dst: switch2port81},
		{Src: dut4port7, Dst: switch2port82},
		{Src: dut4port8, Dst: switch2port83},
		{Src: dut4port9, Dst: switch2port84},
		{Src: dut4port10, Dst: switch2port85},
		{Src: dut4port11, Dst: switch2port86},
		{Src: dut4port12, Dst: switch2port87},
		{Src: dut4port13, Dst: switch2port88},
		{Src: dut4port14, Dst: switch2port89},
		{Src: dut4port15, Dst: switch2port90},
		{Src: dut4port16, Dst: switch2port91},
		{Src: dut4port17, Dst: switch2port92},
		{Src: dut4port18, Dst: switch2port93},
		{Src: dut4port19, Dst: switch2port94},
		{Src: dut4port20, Dst: switch2port95},
		{Src: dut4port21, Dst: switch2port96},
		{Src: dut4port22, Dst: switch2port97},
		{Src: dut4port23, Dst: switch2port98},
		{Src: dut4port24, Dst: switch2port99},
		{Src: dut4port25, Dst: switch2port100},
		{Src: dut4port26, Dst: switch2port101},
		{Src: dut4port27, Dst: switch2port102},
		{Src: dut4port28, Dst: switch2port103},
		{Src: dut4port29, Dst: switch2port104},
		{Src: dut4port30, Dst: switch2port105},
		{Src: dut4port31, Dst: switch2port106},
		{Src: dut4port32, Dst: switch2port107},
		{Src: dut4port33, Dst: switch2port108},
		{Src: dut4port34, Dst: switch2port109},
		{Src: dut4port35, Dst: switch2port110},

		{Src: switch2port111, Dst: ate4port1},
		{Src: switch2port112, Dst: ate4port2},
		{Src: switch2port113, Dst: ate4port3},
		{Src: switch2port114, Dst: ate4port4},
		{Src: switch2port115, Dst: ate4port5},
		{Src: switch2port116, Dst: ate4port6},
		{Src: switch2port117, Dst: ate4port7},
		{Src: switch2port118, Dst: ate4port8},
		{Src: switch2port119, Dst: ate4port9},
		{Src: switch2port120, Dst: ate4port10},
		{Src: switch2port121, Dst: ate4port11},
		{Src: switch2port122, Dst: ate4port12},
		{Src: switch2port123, Dst: ate4port13},
		{Src: switch2port124, Dst: ate4port14},
		{Src: switch2port125, Dst: ate4port15},
		{Src: switch2port126, Dst: ate4port16},
		{Src: switch2port127, Dst: ate4port17},
		{Src: switch2port128, Dst: ate4port18},
		{Src: switch2port129, Dst: ate4port19},
		{Src: switch2port130, Dst: ate4port20},
		{Src: switch2port131, Dst: ate4port21},
		{Src: switch2port132, Dst: ate4port22},
		{Src: switch2port133, Dst: ate4port23},
		{Src: switch2port134, Dst: ate4port24},
		{Src: switch2port135, Dst: ate4port25},
		{Src: switch2port136, Dst: ate4port26},
		{Src: switch2port137, Dst: ate4port27},
		{Src: switch2port138, Dst: ate4port28},
		{Src: switch2port139, Dst: ate4port29},
		{Src: switch2port140, Dst: ate4port30},
		{Src: switch2port36, Dst: ate4port31},
		{Src: switch2port37, Dst: ate4port32},
		{Src: switch2port38, Dst: ate4port33},
		{Src: switch2port39, Dst: ate4port34},
		{Src: switch2port40, Dst: ate4port35},
		{Src: switch2port142, Dst: switch3port141},

		{Src: dut3port36, Dst: switch3port1},
		{Src: dut3port37, Dst: switch3port2},
		{Src: dut3port38, Dst: switch3port3},
		{Src: dut3port39, Dst: switch3port4},
		{Src: dut3port40, Dst: switch3port5},
		{Src: dut3port41, Dst: switch3port6},
		{Src: dut3port42, Dst: switch3port7},
		{Src: dut3port43, Dst: switch3port8},
		{Src: dut3port44, Dst: switch3port9},
		{Src: dut3port45, Dst: switch3port10},
		{Src: dut3port46, Dst: switch3port11},
		{Src: dut3port47, Dst: switch3port12},
		{Src: dut3port48, Dst: switch3port13},
		{Src: dut3port49, Dst: switch3port14},
		{Src: dut3port50, Dst: switch3port15},
		{Src: dut3port51, Dst: switch3port16},
		{Src: dut3port52, Dst: switch3port17},
		{Src: dut3port53, Dst: switch3port18},
		{Src: dut3port54, Dst: switch3port19},
		{Src: dut3port55, Dst: switch3port20},
		{Src: dut3port56, Dst: switch3port21},
		{Src: dut3port57, Dst: switch3port22},
		{Src: dut3port58, Dst: switch3port23},
		{Src: dut3port59, Dst: switch3port24},
		{Src: dut3port60, Dst: switch3port25},
		{Src: dut3port61, Dst: switch3port26},
		{Src: dut3port62, Dst: switch3port27},
		{Src: dut3port63, Dst: switch3port28},
		{Src: dut3port64, Dst: switch3port29},
		{Src: dut3port65, Dst: switch3port30},
		{Src: dut3port66, Dst: switch3port31},
		{Src: dut3port67, Dst: switch3port32},
		{Src: dut3port68, Dst: switch3port33},
		{Src: dut3port69, Dst: switch3port34},
		{Src: dut3port70, Dst: switch3port35},

		{Src: switch3port36, Dst: ate3port36},
		{Src: switch3port37, Dst: ate3port37},
		{Src: switch3port38, Dst: ate3port38},
		{Src: switch3port39, Dst: ate3port39},
		{Src: switch3port40, Dst: ate3port40},
		{Src: switch3port41, Dst: ate3port41},
		{Src: switch3port42, Dst: ate3port42},
		{Src: switch3port43, Dst: ate3port43},
		{Src: switch3port44, Dst: ate3port44},
		{Src: switch3port45, Dst: ate3port45},
		{Src: switch3port46, Dst: ate3port46},
		{Src: switch3port47, Dst: ate3port47},
		{Src: switch3port48, Dst: ate3port48},
		{Src: switch3port49, Dst: ate3port49},
		{Src: switch3port50, Dst: ate3port50},
		{Src: switch3port51, Dst: ate3port51},
		{Src: switch3port52, Dst: ate3port52},
		{Src: switch3port53, Dst: ate3port53},
		{Src: switch3port54, Dst: ate3port54},
		{Src: switch3port55, Dst: ate3port55},
		{Src: switch3port56, Dst: ate3port56},
		{Src: switch3port57, Dst: ate3port57},
		{Src: switch3port58, Dst: ate3port58},
		{Src: switch3port59, Dst: ate3port59},
		{Src: switch3port60, Dst: ate3port60},
		{Src: switch3port61, Dst: ate3port61},
		{Src: switch3port62, Dst: ate3port62},
		{Src: switch3port63, Dst: ate3port63},
		{Src: switch3port64, Dst: ate3port64},
		{Src: switch3port65, Dst: ate3port65},
		{Src: switch3port66, Dst: ate3port66},
		{Src: switch3port67, Dst: ate3port67},
		{Src: switch3port68, Dst: ate3port68},
		{Src: switch3port69, Dst: ate3port69},
		{Src: switch3port70, Dst: ate3port70},

		{Src: dut4port36, Dst: switch3port71},
		{Src: dut4port37, Dst: switch3port72},
		{Src: dut4port38, Dst: switch3port73},
		{Src: dut4port39, Dst: switch3port74},
		{Src: dut4port40, Dst: switch3port75},
		{Src: dut4port41, Dst: switch3port76},
		{Src: dut4port42, Dst: switch3port77},
		{Src: dut4port43, Dst: switch3port78},
		{Src: dut4port44, Dst: switch3port79},
		{Src: dut4port45, Dst: switch3port80},
		{Src: dut4port46, Dst: switch3port81},
		{Src: dut4port47, Dst: switch3port82},
		{Src: dut4port48, Dst: switch3port83},
		{Src: dut4port49, Dst: switch3port84},
		{Src: dut4port50, Dst: switch3port85},
		{Src: dut4port51, Dst: switch3port86},
		{Src: dut4port52, Dst: switch3port87},
		{Src: dut4port53, Dst: switch3port88},
		{Src: dut4port54, Dst: switch3port89},
		{Src: dut4port55, Dst: switch3port90},
		{Src: dut4port56, Dst: switch3port91},
		{Src: dut4port57, Dst: switch3port92},
		{Src: dut4port58, Dst: switch3port93},
		{Src: dut4port59, Dst: switch3port94},
		{Src: dut4port60, Dst: switch3port95},
		{Src: dut4port61, Dst: switch3port96},
		{Src: dut4port62, Dst: switch3port97},
		{Src: dut4port63, Dst: switch3port98},
		{Src: dut4port64, Dst: switch3port99},
		{Src: dut4port65, Dst: switch3port100},
		{Src: dut4port66, Dst: switch3port101},
		{Src: dut4port67, Dst: switch3port102},
		{Src: dut4port68, Dst: switch3port103},
		{Src: dut4port69, Dst: switch3port104},
		{Src: dut4port70, Dst: switch3port105},

		{Src: switch3port106, Dst: ate4port36},
		{Src: switch3port107, Dst: ate4port37},
		{Src: switch3port108, Dst: ate4port38},
		{Src: switch3port109, Dst: ate4port39},
		{Src: switch3port110, Dst: ate4port40},
		{Src: switch3port111, Dst: ate4port41},
		{Src: switch3port112, Dst: ate4port42},
		{Src: switch3port113, Dst: ate4port43},
		{Src: switch3port114, Dst: ate4port44},
		{Src: switch3port115, Dst: ate4port45},
		{Src: switch3port116, Dst: ate4port46},
		{Src: switch3port117, Dst: ate4port47},
		{Src: switch3port118, Dst: ate4port48},
		{Src: switch3port119, Dst: ate4port49},
		{Src: switch3port120, Dst: ate4port50},
		{Src: switch3port121, Dst: ate4port51},
		{Src: switch3port122, Dst: ate4port52},
		{Src: switch3port123, Dst: ate4port53},
		{Src: switch3port124, Dst: ate4port54},
		{Src: switch3port125, Dst: ate4port55},
		{Src: switch3port126, Dst: ate4port56},
		{Src: switch3port127, Dst: ate4port57},
		{Src: switch3port128, Dst: ate4port58},
		{Src: switch3port129, Dst: ate4port59},
		{Src: switch3port130, Dst: ate4port60},
		{Src: switch3port131, Dst: ate4port61},
		{Src: switch3port132, Dst: ate4port62},
		{Src: switch3port133, Dst: ate4port63},
		{Src: switch3port134, Dst: ate4port64},
		{Src: switch3port135, Dst: ate4port65},
		{Src: switch3port136, Dst: ate4port66},
		{Src: switch3port137, Dst: ate4port67},
		{Src: switch3port138, Dst: ate4port68},
		{Src: switch3port139, Dst: ate4port69},
		{Src: switch3port140, Dst: ate4port70},
	},
}

var (
	absdut1port1  = &AbstractPort{Desc: "absdut1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port2  = &AbstractPort{Desc: "absdut1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port3  = &AbstractPort{Desc: "absdut1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port4  = &AbstractPort{Desc: "absdut1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port5  = &AbstractPort{Desc: "absdut1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port6  = &AbstractPort{Desc: "absdut1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port7  = &AbstractPort{Desc: "absdut1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port8  = &AbstractPort{Desc: "absdut1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port9  = &AbstractPort{Desc: "absdut1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port10 = &AbstractPort{Desc: "absdut1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port11 = &AbstractPort{Desc: "absdut1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port12 = &AbstractPort{Desc: "absdut1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port13 = &AbstractPort{Desc: "absdut1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port14 = &AbstractPort{Desc: "absdut1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port15 = &AbstractPort{Desc: "absdut1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port16 = &AbstractPort{Desc: "absdut1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port17 = &AbstractPort{Desc: "absdut1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port18 = &AbstractPort{Desc: "absdut1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port19 = &AbstractPort{Desc: "absdut1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port20 = &AbstractPort{Desc: "absdut1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port21 = &AbstractPort{Desc: "absdut1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port22 = &AbstractPort{Desc: "absdut1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port23 = &AbstractPort{Desc: "absdut1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port24 = &AbstractPort{Desc: "absdut1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port25 = &AbstractPort{Desc: "absdut1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port26 = &AbstractPort{Desc: "absdut1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port27 = &AbstractPort{Desc: "absdut1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port28 = &AbstractPort{Desc: "absdut1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port29 = &AbstractPort{Desc: "absdut1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port30 = &AbstractPort{Desc: "absdut1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port31 = &AbstractPort{Desc: "absdut1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port32 = &AbstractPort{Desc: "absdut1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port33 = &AbstractPort{Desc: "absdut1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port34 = &AbstractPort{Desc: "absdut1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port35 = &AbstractPort{Desc: "absdut1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port36 = &AbstractPort{Desc: "absdut1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port37 = &AbstractPort{Desc: "absdut1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port38 = &AbstractPort{Desc: "absdut1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port39 = &AbstractPort{Desc: "absdut1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port40 = &AbstractPort{Desc: "absdut1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port41 = &AbstractPort{Desc: "absdut1:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port42 = &AbstractPort{Desc: "absdut1:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port43 = &AbstractPort{Desc: "absdut1:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port44 = &AbstractPort{Desc: "absdut1:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port45 = &AbstractPort{Desc: "absdut1:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port46 = &AbstractPort{Desc: "absdut1:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port47 = &AbstractPort{Desc: "absdut1:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port48 = &AbstractPort{Desc: "absdut1:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port49 = &AbstractPort{Desc: "absdut1:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port50 = &AbstractPort{Desc: "absdut1:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port51 = &AbstractPort{Desc: "absdut1:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port52 = &AbstractPort{Desc: "absdut1:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port53 = &AbstractPort{Desc: "absdut1:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port54 = &AbstractPort{Desc: "absdut1:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port55 = &AbstractPort{Desc: "absdut1:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port56 = &AbstractPort{Desc: "absdut1:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port57 = &AbstractPort{Desc: "absdut1:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port58 = &AbstractPort{Desc: "absdut1:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port59 = &AbstractPort{Desc: "absdut1:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port60 = &AbstractPort{Desc: "absdut1:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port61 = &AbstractPort{Desc: "absdut1:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port62 = &AbstractPort{Desc: "absdut1:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port63 = &AbstractPort{Desc: "absdut1:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port64 = &AbstractPort{Desc: "absdut1:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port65 = &AbstractPort{Desc: "absdut1:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port66 = &AbstractPort{Desc: "absdut1:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port67 = &AbstractPort{Desc: "absdut1:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port68 = &AbstractPort{Desc: "absdut1:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port69 = &AbstractPort{Desc: "absdut1:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut1port70 = &AbstractPort{Desc: "absdut1:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absdut1 = &AbstractNode{Desc: "absdut1", Ports: []*AbstractPort{
		absdut1port1,
		absdut1port2,
		absdut1port3,
		absdut1port4,
		absdut1port5,
		absdut1port6,
		absdut1port7,
		absdut1port8,
		absdut1port9,
		absdut1port10,
		absdut1port11,
		absdut1port12,
		absdut1port13,
		absdut1port14,
		absdut1port15,
		absdut1port16,
		absdut1port17,
		absdut1port18,
		absdut1port19,
		absdut1port20,
		absdut1port21,
		absdut1port22,
		absdut1port23,
		absdut1port24,
		absdut1port25,
		absdut1port26,
		absdut1port27,
		absdut1port28,
		absdut1port29,
		absdut1port30,
		absdut1port31,
		absdut1port32,
		absdut1port33,
		absdut1port34,
		absdut1port35,
		absdut1port36,
		absdut1port37,
		absdut1port38,
		absdut1port39,
		absdut1port40,
		absdut1port41,
		absdut1port42,
		absdut1port43,
		absdut1port44,
		absdut1port45,
		absdut1port46,
		absdut1port47,
		absdut1port48,
		absdut1port49,
		absdut1port50,
		absdut1port51,
		absdut1port52,
		absdut1port53,
		absdut1port54,
		absdut1port55,
		absdut1port56,
		absdut1port57,
		absdut1port58,
		absdut1port59,
		absdut1port60,
		absdut1port61,
		absdut1port62,
		absdut1port63,
		absdut1port64,
		absdut1port65,
		absdut1port66,
		absdut1port67,
		absdut1port68,
		absdut1port69,
		absdut1port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	absdut2port1  = &AbstractPort{Desc: "absdut2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port2  = &AbstractPort{Desc: "absdut2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port3  = &AbstractPort{Desc: "absdut2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port4  = &AbstractPort{Desc: "absdut2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port5  = &AbstractPort{Desc: "absdut2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port6  = &AbstractPort{Desc: "absdut2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port7  = &AbstractPort{Desc: "absdut2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port8  = &AbstractPort{Desc: "absdut2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port9  = &AbstractPort{Desc: "absdut2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port10 = &AbstractPort{Desc: "absdut2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port11 = &AbstractPort{Desc: "absdut2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port12 = &AbstractPort{Desc: "absdut2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port13 = &AbstractPort{Desc: "absdut2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port14 = &AbstractPort{Desc: "absdut2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port15 = &AbstractPort{Desc: "absdut2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port16 = &AbstractPort{Desc: "absdut2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port17 = &AbstractPort{Desc: "absdut2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port18 = &AbstractPort{Desc: "absdut2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port19 = &AbstractPort{Desc: "absdut2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port20 = &AbstractPort{Desc: "absdut2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port21 = &AbstractPort{Desc: "absdut2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port22 = &AbstractPort{Desc: "absdut2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port23 = &AbstractPort{Desc: "absdut2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port24 = &AbstractPort{Desc: "absdut2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port25 = &AbstractPort{Desc: "absdut2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port26 = &AbstractPort{Desc: "absdut2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port27 = &AbstractPort{Desc: "absdut2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port28 = &AbstractPort{Desc: "absdut2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port29 = &AbstractPort{Desc: "absdut2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port30 = &AbstractPort{Desc: "absdut2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port31 = &AbstractPort{Desc: "absdut2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port32 = &AbstractPort{Desc: "absdut2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port33 = &AbstractPort{Desc: "absdut2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port34 = &AbstractPort{Desc: "absdut2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port35 = &AbstractPort{Desc: "absdut2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port36 = &AbstractPort{Desc: "absdut2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port37 = &AbstractPort{Desc: "absdut2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port38 = &AbstractPort{Desc: "absdut2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port39 = &AbstractPort{Desc: "absdut2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port40 = &AbstractPort{Desc: "absdut2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port41 = &AbstractPort{Desc: "absdut2:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port42 = &AbstractPort{Desc: "absdut2:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port43 = &AbstractPort{Desc: "absdut2:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port44 = &AbstractPort{Desc: "absdut2:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port45 = &AbstractPort{Desc: "absdut2:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port46 = &AbstractPort{Desc: "absdut2:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port47 = &AbstractPort{Desc: "absdut2:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port48 = &AbstractPort{Desc: "absdut2:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port49 = &AbstractPort{Desc: "absdut2:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port50 = &AbstractPort{Desc: "absdut2:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port51 = &AbstractPort{Desc: "absdut2:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port52 = &AbstractPort{Desc: "absdut2:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port53 = &AbstractPort{Desc: "absdut2:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port54 = &AbstractPort{Desc: "absdut2:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port55 = &AbstractPort{Desc: "absdut2:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port56 = &AbstractPort{Desc: "absdut2:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port57 = &AbstractPort{Desc: "absdut2:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port58 = &AbstractPort{Desc: "absdut2:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port59 = &AbstractPort{Desc: "absdut2:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port60 = &AbstractPort{Desc: "absdut2:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port61 = &AbstractPort{Desc: "absdut2:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port62 = &AbstractPort{Desc: "absdut2:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port63 = &AbstractPort{Desc: "absdut2:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port64 = &AbstractPort{Desc: "absdut2:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port65 = &AbstractPort{Desc: "absdut2:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port66 = &AbstractPort{Desc: "absdut2:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port67 = &AbstractPort{Desc: "absdut2:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port68 = &AbstractPort{Desc: "absdut2:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port69 = &AbstractPort{Desc: "absdut2:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut2port70 = &AbstractPort{Desc: "absdut2:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absdut2 = &AbstractNode{Desc: "absdut2", Ports: []*AbstractPort{
		absdut2port1,
		absdut2port2,
		absdut2port3,
		absdut2port4,
		absdut2port5,
		absdut2port6,
		absdut2port7,
		absdut2port8,
		absdut2port9,
		absdut2port10,
		absdut2port11,
		absdut2port12,
		absdut2port13,
		absdut2port14,
		absdut2port15,
		absdut2port16,
		absdut2port17,
		absdut2port18,
		absdut2port19,
		absdut2port20,
		absdut2port21,
		absdut2port22,
		absdut2port23,
		absdut2port24,
		absdut2port25,
		absdut2port26,
		absdut2port27,
		absdut2port28,
		absdut2port29,
		absdut2port30,
		absdut2port31,
		absdut2port32,
		absdut2port33,
		absdut2port34,
		absdut2port35,
		absdut2port36,
		absdut2port37,
		absdut2port38,
		absdut2port39,
		absdut2port40,
		absdut2port41,
		absdut2port42,
		absdut2port43,
		absdut2port44,
		absdut2port45,
		absdut2port46,
		absdut2port47,
		absdut2port48,
		absdut2port49,
		absdut2port50,
		absdut2port51,
		absdut2port52,
		absdut2port53,
		absdut2port54,
		absdut2port55,
		absdut2port56,
		absdut2port57,
		absdut2port58,
		absdut2port59,
		absdut2port60,
		absdut2port61,
		absdut2port62,
		absdut2port63,
		absdut2port64,
		absdut2port65,
		absdut2port66,
		absdut2port67,
		absdut2port68,
		absdut2port69,
		absdut2port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	absdut3port1  = &AbstractPort{Desc: "absdut3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port2  = &AbstractPort{Desc: "absdut3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port3  = &AbstractPort{Desc: "absdut3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port4  = &AbstractPort{Desc: "absdut3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port5  = &AbstractPort{Desc: "absdut3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port6  = &AbstractPort{Desc: "absdut3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port7  = &AbstractPort{Desc: "absdut3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port8  = &AbstractPort{Desc: "absdut3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port9  = &AbstractPort{Desc: "absdut3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port10 = &AbstractPort{Desc: "absdut3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port11 = &AbstractPort{Desc: "absdut3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port12 = &AbstractPort{Desc: "absdut3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port13 = &AbstractPort{Desc: "absdut3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port14 = &AbstractPort{Desc: "absdut3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port15 = &AbstractPort{Desc: "absdut3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port16 = &AbstractPort{Desc: "absdut3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port17 = &AbstractPort{Desc: "absdut3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port18 = &AbstractPort{Desc: "absdut3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port19 = &AbstractPort{Desc: "absdut3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port20 = &AbstractPort{Desc: "absdut3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port21 = &AbstractPort{Desc: "absdut3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port22 = &AbstractPort{Desc: "absdut3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port23 = &AbstractPort{Desc: "absdut3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port24 = &AbstractPort{Desc: "absdut3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port25 = &AbstractPort{Desc: "absdut3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port26 = &AbstractPort{Desc: "absdut3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port27 = &AbstractPort{Desc: "absdut3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port28 = &AbstractPort{Desc: "absdut3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port29 = &AbstractPort{Desc: "absdut3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port30 = &AbstractPort{Desc: "absdut3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port31 = &AbstractPort{Desc: "absdut3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port32 = &AbstractPort{Desc: "absdut3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port33 = &AbstractPort{Desc: "absdut3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port34 = &AbstractPort{Desc: "absdut3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port35 = &AbstractPort{Desc: "absdut3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port36 = &AbstractPort{Desc: "absdut3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port37 = &AbstractPort{Desc: "absdut3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port38 = &AbstractPort{Desc: "absdut3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port39 = &AbstractPort{Desc: "absdut3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port40 = &AbstractPort{Desc: "absdut3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port41 = &AbstractPort{Desc: "absdut3:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port42 = &AbstractPort{Desc: "absdut3:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port43 = &AbstractPort{Desc: "absdut3:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port44 = &AbstractPort{Desc: "absdut3:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port45 = &AbstractPort{Desc: "absdut3:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port46 = &AbstractPort{Desc: "absdut3:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port47 = &AbstractPort{Desc: "absdut3:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port48 = &AbstractPort{Desc: "absdut3:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port49 = &AbstractPort{Desc: "absdut3:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port50 = &AbstractPort{Desc: "absdut3:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port51 = &AbstractPort{Desc: "absdut3:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port52 = &AbstractPort{Desc: "absdut3:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port53 = &AbstractPort{Desc: "absdut3:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port54 = &AbstractPort{Desc: "absdut3:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port55 = &AbstractPort{Desc: "absdut3:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port56 = &AbstractPort{Desc: "absdut3:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port57 = &AbstractPort{Desc: "absdut3:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port58 = &AbstractPort{Desc: "absdut3:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port59 = &AbstractPort{Desc: "absdut3:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port60 = &AbstractPort{Desc: "absdut3:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port61 = &AbstractPort{Desc: "absdut3:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port62 = &AbstractPort{Desc: "absdut3:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port63 = &AbstractPort{Desc: "absdut3:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port64 = &AbstractPort{Desc: "absdut3:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port65 = &AbstractPort{Desc: "absdut3:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port66 = &AbstractPort{Desc: "absdut3:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port67 = &AbstractPort{Desc: "absdut3:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port68 = &AbstractPort{Desc: "absdut3:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port69 = &AbstractPort{Desc: "absdut3:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut3port70 = &AbstractPort{Desc: "absdut3:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absdut3 = &AbstractNode{Desc: "absdut3", Ports: []*AbstractPort{
		absdut3port1,
		absdut3port2,
		absdut3port3,
		absdut3port4,
		absdut3port5,
		absdut3port6,
		absdut3port7,
		absdut3port8,
		absdut3port9,
		absdut3port10,
		absdut3port11,
		absdut3port12,
		absdut3port13,
		absdut3port14,
		absdut3port15,
		absdut3port16,
		absdut3port17,
		absdut3port18,
		absdut3port19,
		absdut3port20,
		absdut3port21,
		absdut3port22,
		absdut3port23,
		absdut3port24,
		absdut3port25,
		absdut3port26,
		absdut3port27,
		absdut3port28,
		absdut3port29,
		absdut3port30,
		absdut3port31,
		absdut3port32,
		absdut3port33,
		absdut3port34,
		absdut3port35,
		absdut3port36,
		absdut3port37,
		absdut3port38,
		absdut3port39,
		absdut3port40,
		absdut3port41,
		absdut3port42,
		absdut3port43,
		absdut3port44,
		absdut3port45,
		absdut3port46,
		absdut3port47,
		absdut3port48,
		absdut3port49,
		absdut3port50,
		absdut3port51,
		absdut3port52,
		absdut3port53,
		absdut3port54,
		absdut3port55,
		absdut3port56,
		absdut3port57,
		absdut3port58,
		absdut3port59,
		absdut3port60,
		absdut3port61,
		absdut3port62,
		absdut3port63,
		absdut3port64,
		absdut3port65,
		absdut3port66,
		absdut3port67,
		absdut3port68,
		absdut3port69,
		absdut3port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	absdut4port1  = &AbstractPort{Desc: "absdut4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port2  = &AbstractPort{Desc: "absdut4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port3  = &AbstractPort{Desc: "absdut4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port4  = &AbstractPort{Desc: "absdut4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port5  = &AbstractPort{Desc: "absdut4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port6  = &AbstractPort{Desc: "absdut4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port7  = &AbstractPort{Desc: "absdut4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port8  = &AbstractPort{Desc: "absdut4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port9  = &AbstractPort{Desc: "absdut4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port10 = &AbstractPort{Desc: "absdut4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port11 = &AbstractPort{Desc: "absdut4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port12 = &AbstractPort{Desc: "absdut4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port13 = &AbstractPort{Desc: "absdut4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port14 = &AbstractPort{Desc: "absdut4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port15 = &AbstractPort{Desc: "absdut4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port16 = &AbstractPort{Desc: "absdut4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port17 = &AbstractPort{Desc: "absdut4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port18 = &AbstractPort{Desc: "absdut4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port19 = &AbstractPort{Desc: "absdut4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port20 = &AbstractPort{Desc: "absdut4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port21 = &AbstractPort{Desc: "absdut4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port22 = &AbstractPort{Desc: "absdut4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port23 = &AbstractPort{Desc: "absdut4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port24 = &AbstractPort{Desc: "absdut4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port25 = &AbstractPort{Desc: "absdut4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port26 = &AbstractPort{Desc: "absdut4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port27 = &AbstractPort{Desc: "absdut4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port28 = &AbstractPort{Desc: "absdut4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port29 = &AbstractPort{Desc: "absdut4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port30 = &AbstractPort{Desc: "absdut4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port31 = &AbstractPort{Desc: "absdut4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port32 = &AbstractPort{Desc: "absdut4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port33 = &AbstractPort{Desc: "absdut4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port34 = &AbstractPort{Desc: "absdut4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port35 = &AbstractPort{Desc: "absdut4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port36 = &AbstractPort{Desc: "absdut4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port37 = &AbstractPort{Desc: "absdut4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port38 = &AbstractPort{Desc: "absdut4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port39 = &AbstractPort{Desc: "absdut4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port40 = &AbstractPort{Desc: "absdut4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port41 = &AbstractPort{Desc: "absdut4:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port42 = &AbstractPort{Desc: "absdut4:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port43 = &AbstractPort{Desc: "absdut4:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port44 = &AbstractPort{Desc: "absdut4:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port45 = &AbstractPort{Desc: "absdut4:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port46 = &AbstractPort{Desc: "absdut4:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port47 = &AbstractPort{Desc: "absdut4:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port48 = &AbstractPort{Desc: "absdut4:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port49 = &AbstractPort{Desc: "absdut4:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port50 = &AbstractPort{Desc: "absdut4:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port51 = &AbstractPort{Desc: "absdut4:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port52 = &AbstractPort{Desc: "absdut4:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port53 = &AbstractPort{Desc: "absdut4:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port54 = &AbstractPort{Desc: "absdut4:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port55 = &AbstractPort{Desc: "absdut4:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port56 = &AbstractPort{Desc: "absdut4:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port57 = &AbstractPort{Desc: "absdut4:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port58 = &AbstractPort{Desc: "absdut4:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port59 = &AbstractPort{Desc: "absdut4:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port60 = &AbstractPort{Desc: "absdut4:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port61 = &AbstractPort{Desc: "absdut4:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port62 = &AbstractPort{Desc: "absdut4:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port63 = &AbstractPort{Desc: "absdut4:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port64 = &AbstractPort{Desc: "absdut4:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port65 = &AbstractPort{Desc: "absdut4:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port66 = &AbstractPort{Desc: "absdut4:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port67 = &AbstractPort{Desc: "absdut4:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port68 = &AbstractPort{Desc: "absdut4:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port69 = &AbstractPort{Desc: "absdut4:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absdut4port70 = &AbstractPort{Desc: "absdut4:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absdut4 = &AbstractNode{Desc: "absdut4", Ports: []*AbstractPort{
		absdut4port1,
		absdut4port2,
		absdut4port3,
		absdut4port4,
		absdut4port5,
		absdut4port6,
		absdut4port7,
		absdut4port8,
		absdut4port9,
		absdut4port10,
		absdut4port11,
		absdut4port12,
		absdut4port13,
		absdut4port14,
		absdut4port15,
		absdut4port16,
		absdut4port17,
		absdut4port18,
		absdut4port19,
		absdut4port20,
		absdut4port21,
		absdut4port22,
		absdut4port23,
		absdut4port24,
		absdut4port25,
		absdut4port26,
		absdut4port27,
		absdut4port28,
		absdut4port29,
		absdut4port30,
		absdut4port31,
		absdut4port32,
		absdut4port33,
		absdut4port34,
		absdut4port35,
		absdut4port36,
		absdut4port37,
		absdut4port38,
		absdut4port39,
		absdut4port40,
		absdut4port41,
		absdut4port42,
		absdut4port43,
		absdut4port44,
		absdut4port45,
		absdut4port46,
		absdut4port47,
		absdut4port48,
		absdut4port49,
		absdut4port50,
		absdut4port51,
		absdut4port52,
		absdut4port53,
		absdut4port54,
		absdut4port55,
		absdut4port56,
		absdut4port57,
		absdut4port58,
		absdut4port59,
		absdut4port60,
		absdut4port61,
		absdut4port62,
		absdut4port63,
		absdut4port64,
		absdut4port65,
		absdut4port66,
		absdut4port67,
		absdut4port68,
		absdut4port69,
		absdut4port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}

	absate1port1  = &AbstractPort{Desc: "absate1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port2  = &AbstractPort{Desc: "absate1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port3  = &AbstractPort{Desc: "absate1:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port4  = &AbstractPort{Desc: "absate1:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port5  = &AbstractPort{Desc: "absate1:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port6  = &AbstractPort{Desc: "absate1:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port7  = &AbstractPort{Desc: "absate1:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port8  = &AbstractPort{Desc: "absate1:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port9  = &AbstractPort{Desc: "absate1:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port10 = &AbstractPort{Desc: "absate1:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port11 = &AbstractPort{Desc: "absate1:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port12 = &AbstractPort{Desc: "absate1:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port13 = &AbstractPort{Desc: "absate1:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port14 = &AbstractPort{Desc: "absate1:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port15 = &AbstractPort{Desc: "absate1:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port16 = &AbstractPort{Desc: "absate1:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port17 = &AbstractPort{Desc: "absate1:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port18 = &AbstractPort{Desc: "absate1:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port19 = &AbstractPort{Desc: "absate1:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port20 = &AbstractPort{Desc: "absate1:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port21 = &AbstractPort{Desc: "absate1:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port22 = &AbstractPort{Desc: "absate1:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port23 = &AbstractPort{Desc: "absate1:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port24 = &AbstractPort{Desc: "absate1:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port25 = &AbstractPort{Desc: "absate1:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port26 = &AbstractPort{Desc: "absate1:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port27 = &AbstractPort{Desc: "absate1:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port28 = &AbstractPort{Desc: "absate1:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port29 = &AbstractPort{Desc: "absate1:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port30 = &AbstractPort{Desc: "absate1:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port31 = &AbstractPort{Desc: "absate1:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port32 = &AbstractPort{Desc: "absate1:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port33 = &AbstractPort{Desc: "absate1:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port34 = &AbstractPort{Desc: "absate1:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port35 = &AbstractPort{Desc: "absate1:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port36 = &AbstractPort{Desc: "absate1:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port37 = &AbstractPort{Desc: "absate1:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port38 = &AbstractPort{Desc: "absate1:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port39 = &AbstractPort{Desc: "absate1:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port40 = &AbstractPort{Desc: "absate1:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port41 = &AbstractPort{Desc: "absate1:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port42 = &AbstractPort{Desc: "absate1:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port43 = &AbstractPort{Desc: "absate1:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port44 = &AbstractPort{Desc: "absate1:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port45 = &AbstractPort{Desc: "absate1:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port46 = &AbstractPort{Desc: "absate1:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port47 = &AbstractPort{Desc: "absate1:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port48 = &AbstractPort{Desc: "absate1:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port49 = &AbstractPort{Desc: "absate1:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port50 = &AbstractPort{Desc: "absate1:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port51 = &AbstractPort{Desc: "absate1:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port52 = &AbstractPort{Desc: "absate1:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port53 = &AbstractPort{Desc: "absate1:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port54 = &AbstractPort{Desc: "absate1:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port55 = &AbstractPort{Desc: "absate1:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port56 = &AbstractPort{Desc: "absate1:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port57 = &AbstractPort{Desc: "absate1:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port58 = &AbstractPort{Desc: "absate1:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port59 = &AbstractPort{Desc: "absate1:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port60 = &AbstractPort{Desc: "absate1:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port61 = &AbstractPort{Desc: "absate1:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port62 = &AbstractPort{Desc: "absate1:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port63 = &AbstractPort{Desc: "absate1:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port64 = &AbstractPort{Desc: "absate1:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port65 = &AbstractPort{Desc: "absate1:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port66 = &AbstractPort{Desc: "absate1:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port67 = &AbstractPort{Desc: "absate1:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port68 = &AbstractPort{Desc: "absate1:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port69 = &AbstractPort{Desc: "absate1:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate1port70 = &AbstractPort{Desc: "absate1:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absate1 = &AbstractNode{Desc: "absate1", Ports: []*AbstractPort{
		absate1port1,
		absate1port2,
		absate1port3,
		absate1port4,
		absate1port5,
		absate1port6,
		absate1port7,
		absate1port8,
		absate1port9,
		absate1port10,
		absate1port11,
		absate1port12,
		absate1port13,
		absate1port14,
		absate1port15,
		absate1port16,
		absate1port17,
		absate1port18,
		absate1port19,
		absate1port20,
		absate1port21,
		absate1port22,
		absate1port23,
		absate1port24,
		absate1port25,
		absate1port26,
		absate1port27,
		absate1port28,
		absate1port29,
		absate1port30,
		absate1port31,
		absate1port32,
		absate1port33,
		absate1port34,
		absate1port35,
		absate1port36,
		absate1port37,
		absate1port38,
		absate1port39,
		absate1port40,
		absate1port41,
		absate1port42,
		absate1port43,
		absate1port44,
		absate1port45,
		absate1port46,
		absate1port47,
		absate1port48,
		absate1port49,
		absate1port50,
		absate1port51,
		absate1port52,
		absate1port53,
		absate1port54,
		absate1port55,
		absate1port56,
		absate1port57,
		absate1port58,
		absate1port59,
		absate1port60,
		absate1port61,
		absate1port62,
		absate1port63,
		absate1port64,
		absate1port65,
		absate1port66,
		absate1port67,
		absate1port68,
		absate1port69,
		absate1port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	absate2port1  = &AbstractPort{Desc: "absate2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port2  = &AbstractPort{Desc: "absate2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port3  = &AbstractPort{Desc: "absate2:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port4  = &AbstractPort{Desc: "absate2:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port5  = &AbstractPort{Desc: "absate2:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port6  = &AbstractPort{Desc: "absate2:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port7  = &AbstractPort{Desc: "absate2:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port8  = &AbstractPort{Desc: "absate2:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port9  = &AbstractPort{Desc: "absate2:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port10 = &AbstractPort{Desc: "absate2:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port11 = &AbstractPort{Desc: "absate2:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port12 = &AbstractPort{Desc: "absate2:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port13 = &AbstractPort{Desc: "absate2:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port14 = &AbstractPort{Desc: "absate2:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port15 = &AbstractPort{Desc: "absate2:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port16 = &AbstractPort{Desc: "absate2:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port17 = &AbstractPort{Desc: "absate2:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port18 = &AbstractPort{Desc: "absate2:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port19 = &AbstractPort{Desc: "absate2:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port20 = &AbstractPort{Desc: "absate2:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port21 = &AbstractPort{Desc: "absate2:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port22 = &AbstractPort{Desc: "absate2:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port23 = &AbstractPort{Desc: "absate2:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port24 = &AbstractPort{Desc: "absate2:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port25 = &AbstractPort{Desc: "absate2:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port26 = &AbstractPort{Desc: "absate2:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port27 = &AbstractPort{Desc: "absate2:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port28 = &AbstractPort{Desc: "absate2:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port29 = &AbstractPort{Desc: "absate2:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port30 = &AbstractPort{Desc: "absate2:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port31 = &AbstractPort{Desc: "absate2:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port32 = &AbstractPort{Desc: "absate2:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port33 = &AbstractPort{Desc: "absate2:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port34 = &AbstractPort{Desc: "absate2:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port35 = &AbstractPort{Desc: "absate2:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port36 = &AbstractPort{Desc: "absate2:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port37 = &AbstractPort{Desc: "absate2:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port38 = &AbstractPort{Desc: "absate2:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port39 = &AbstractPort{Desc: "absate2:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port40 = &AbstractPort{Desc: "absate2:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port41 = &AbstractPort{Desc: "absate2:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port42 = &AbstractPort{Desc: "absate2:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port43 = &AbstractPort{Desc: "absate2:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port44 = &AbstractPort{Desc: "absate2:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port45 = &AbstractPort{Desc: "absate2:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port46 = &AbstractPort{Desc: "absate2:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port47 = &AbstractPort{Desc: "absate2:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port48 = &AbstractPort{Desc: "absate2:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port49 = &AbstractPort{Desc: "absate2:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port50 = &AbstractPort{Desc: "absate2:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port51 = &AbstractPort{Desc: "absate2:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port52 = &AbstractPort{Desc: "absate2:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port53 = &AbstractPort{Desc: "absate2:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port54 = &AbstractPort{Desc: "absate2:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port55 = &AbstractPort{Desc: "absate2:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port56 = &AbstractPort{Desc: "absate2:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port57 = &AbstractPort{Desc: "absate2:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port58 = &AbstractPort{Desc: "absate2:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port59 = &AbstractPort{Desc: "absate2:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port60 = &AbstractPort{Desc: "absate2:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port61 = &AbstractPort{Desc: "absate2:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port62 = &AbstractPort{Desc: "absate2:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port63 = &AbstractPort{Desc: "absate2:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port64 = &AbstractPort{Desc: "absate2:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port65 = &AbstractPort{Desc: "absate2:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port66 = &AbstractPort{Desc: "absate2:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port67 = &AbstractPort{Desc: "absate2:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port68 = &AbstractPort{Desc: "absate2:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port69 = &AbstractPort{Desc: "absate2:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate2port70 = &AbstractPort{Desc: "absate2:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absate2 = &AbstractNode{Desc: "absate2", Ports: []*AbstractPort{
		absate2port1,
		absate2port2,
		absate2port3,
		absate2port4,
		absate2port5,
		absate2port6,
		absate2port7,
		absate2port8,
		absate2port9,
		absate2port10,
		absate2port11,
		absate2port12,
		absate2port13,
		absate2port14,
		absate2port15,
		absate2port16,
		absate2port17,
		absate2port18,
		absate2port19,
		absate2port20,
		absate2port21,
		absate2port22,
		absate2port23,
		absate2port24,
		absate2port25,
		absate2port26,
		absate2port27,
		absate2port28,
		absate2port29,
		absate2port30,
		absate2port31,
		absate2port32,
		absate2port33,
		absate2port34,
		absate2port35,
		absate2port36,
		absate2port37,
		absate2port38,
		absate2port39,
		absate2port40,
		absate2port41,
		absate2port42,
		absate2port43,
		absate2port44,
		absate2port45,
		absate2port46,
		absate2port47,
		absate2port48,
		absate2port49,
		absate2port50,
		absate2port51,
		absate2port52,
		absate2port53,
		absate2port54,
		absate2port55,
		absate2port56,
		absate2port57,
		absate2port58,
		absate2port59,
		absate2port60,
		absate2port61,
		absate2port62,
		absate2port63,
		absate2port64,
		absate2port65,
		absate2port66,
		absate2port67,
		absate2port68,
		absate2port69,
		absate2port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	absate3port1  = &AbstractPort{Desc: "absate3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port2  = &AbstractPort{Desc: "absate3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port3  = &AbstractPort{Desc: "absate3:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port4  = &AbstractPort{Desc: "absate3:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port5  = &AbstractPort{Desc: "absate3:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port6  = &AbstractPort{Desc: "absate3:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port7  = &AbstractPort{Desc: "absate3:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port8  = &AbstractPort{Desc: "absate3:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port9  = &AbstractPort{Desc: "absate3:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port10 = &AbstractPort{Desc: "absate3:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port11 = &AbstractPort{Desc: "absate3:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port12 = &AbstractPort{Desc: "absate3:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port13 = &AbstractPort{Desc: "absate3:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port14 = &AbstractPort{Desc: "absate3:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port15 = &AbstractPort{Desc: "absate3:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port16 = &AbstractPort{Desc: "absate3:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port17 = &AbstractPort{Desc: "absate3:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port18 = &AbstractPort{Desc: "absate3:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port19 = &AbstractPort{Desc: "absate3:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port20 = &AbstractPort{Desc: "absate3:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port21 = &AbstractPort{Desc: "absate3:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port22 = &AbstractPort{Desc: "absate3:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port23 = &AbstractPort{Desc: "absate3:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port24 = &AbstractPort{Desc: "absate3:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port25 = &AbstractPort{Desc: "absate3:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port26 = &AbstractPort{Desc: "absate3:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port27 = &AbstractPort{Desc: "absate3:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port28 = &AbstractPort{Desc: "absate3:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port29 = &AbstractPort{Desc: "absate3:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port30 = &AbstractPort{Desc: "absate3:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port31 = &AbstractPort{Desc: "absate3:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port32 = &AbstractPort{Desc: "absate3:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port33 = &AbstractPort{Desc: "absate3:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port34 = &AbstractPort{Desc: "absate3:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port35 = &AbstractPort{Desc: "absate3:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port36 = &AbstractPort{Desc: "absate3:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port37 = &AbstractPort{Desc: "absate3:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port38 = &AbstractPort{Desc: "absate3:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port39 = &AbstractPort{Desc: "absate3:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port40 = &AbstractPort{Desc: "absate3:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port41 = &AbstractPort{Desc: "absate3:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port42 = &AbstractPort{Desc: "absate3:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port43 = &AbstractPort{Desc: "absate3:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port44 = &AbstractPort{Desc: "absate3:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port45 = &AbstractPort{Desc: "absate3:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port46 = &AbstractPort{Desc: "absate3:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port47 = &AbstractPort{Desc: "absate3:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port48 = &AbstractPort{Desc: "absate3:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port49 = &AbstractPort{Desc: "absate3:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port50 = &AbstractPort{Desc: "absate3:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port51 = &AbstractPort{Desc: "absate3:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port52 = &AbstractPort{Desc: "absate3:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port53 = &AbstractPort{Desc: "absate3:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port54 = &AbstractPort{Desc: "absate3:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port55 = &AbstractPort{Desc: "absate3:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port56 = &AbstractPort{Desc: "absate3:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port57 = &AbstractPort{Desc: "absate3:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port58 = &AbstractPort{Desc: "absate3:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port59 = &AbstractPort{Desc: "absate3:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port60 = &AbstractPort{Desc: "absate3:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port61 = &AbstractPort{Desc: "absate3:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port62 = &AbstractPort{Desc: "absate3:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port63 = &AbstractPort{Desc: "absate3:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port64 = &AbstractPort{Desc: "absate3:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port65 = &AbstractPort{Desc: "absate3:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port66 = &AbstractPort{Desc: "absate3:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port67 = &AbstractPort{Desc: "absate3:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port68 = &AbstractPort{Desc: "absate3:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port69 = &AbstractPort{Desc: "absate3:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate3port70 = &AbstractPort{Desc: "absate3:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absate3 = &AbstractNode{Desc: "absate3", Ports: []*AbstractPort{
		absate3port1,
		absate3port2,
		absate3port3,
		absate3port4,
		absate3port5,
		absate3port6,
		absate3port7,
		absate3port8,
		absate3port9,
		absate3port10,
		absate3port11,
		absate3port12,
		absate3port13,
		absate3port14,
		absate3port15,
		absate3port16,
		absate3port17,
		absate3port18,
		absate3port19,
		absate3port20,
		absate3port21,
		absate3port22,
		absate3port23,
		absate3port24,
		absate3port25,
		absate3port26,
		absate3port27,
		absate3port28,
		absate3port29,
		absate3port30,
		absate3port31,
		absate3port32,
		absate3port33,
		absate3port34,
		absate3port35,
		absate3port36,
		absate3port37,
		absate3port38,
		absate3port39,
		absate3port40,
		absate3port41,
		absate3port42,
		absate3port43,
		absate3port44,
		absate3port45,
		absate3port46,
		absate3port47,
		absate3port48,
		absate3port49,
		absate3port50,
		absate3port51,
		absate3port52,
		absate3port53,
		absate3port54,
		absate3port55,
		absate3port56,
		absate3port57,
		absate3port58,
		absate3port59,
		absate3port60,
		absate3port61,
		absate3port62,
		absate3port63,
		absate3port64,
		absate3port65,
		absate3port66,
		absate3port67,
		absate3port68,
		absate3port69,
		absate3port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}

	absate4port1  = &AbstractPort{Desc: "absate4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port2  = &AbstractPort{Desc: "absate4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port3  = &AbstractPort{Desc: "absate4:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port4  = &AbstractPort{Desc: "absate4:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port5  = &AbstractPort{Desc: "absate4:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port6  = &AbstractPort{Desc: "absate4:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port7  = &AbstractPort{Desc: "absate4:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port8  = &AbstractPort{Desc: "absate4:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port9  = &AbstractPort{Desc: "absate4:port9", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port10 = &AbstractPort{Desc: "absate4:port10", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port11 = &AbstractPort{Desc: "absate4:port11", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port12 = &AbstractPort{Desc: "absate4:port12", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port13 = &AbstractPort{Desc: "absate4:port13", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port14 = &AbstractPort{Desc: "absate4:port14", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port15 = &AbstractPort{Desc: "absate4:port15", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port16 = &AbstractPort{Desc: "absate4:port16", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port17 = &AbstractPort{Desc: "absate4:port17", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port18 = &AbstractPort{Desc: "absate4:port18", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port19 = &AbstractPort{Desc: "absate4:port19", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port20 = &AbstractPort{Desc: "absate4:port20", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port21 = &AbstractPort{Desc: "absate4:port21", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port22 = &AbstractPort{Desc: "absate4:port22", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port23 = &AbstractPort{Desc: "absate4:port23", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port24 = &AbstractPort{Desc: "absate4:port24", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port25 = &AbstractPort{Desc: "absate4:port25", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port26 = &AbstractPort{Desc: "absate4:port26", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port27 = &AbstractPort{Desc: "absate4:port27", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port28 = &AbstractPort{Desc: "absate4:port28", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port29 = &AbstractPort{Desc: "absate4:port29", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port30 = &AbstractPort{Desc: "absate4:port30", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port31 = &AbstractPort{Desc: "absate4:port31", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port32 = &AbstractPort{Desc: "absate4:port32", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port33 = &AbstractPort{Desc: "absate4:port33", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port34 = &AbstractPort{Desc: "absate4:port34", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port35 = &AbstractPort{Desc: "absate4:port35", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port36 = &AbstractPort{Desc: "absate4:port36", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port37 = &AbstractPort{Desc: "absate4:port37", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port38 = &AbstractPort{Desc: "absate4:port38", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port39 = &AbstractPort{Desc: "absate4:port39", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port40 = &AbstractPort{Desc: "absate4:port40", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port41 = &AbstractPort{Desc: "absate4:port41", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port42 = &AbstractPort{Desc: "absate4:port42", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port43 = &AbstractPort{Desc: "absate4:port43", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port44 = &AbstractPort{Desc: "absate4:port44", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port45 = &AbstractPort{Desc: "absate4:port45", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port46 = &AbstractPort{Desc: "absate4:port46", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port47 = &AbstractPort{Desc: "absate4:port47", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port48 = &AbstractPort{Desc: "absate4:port48", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port49 = &AbstractPort{Desc: "absate4:port49", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port50 = &AbstractPort{Desc: "absate4:port50", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port51 = &AbstractPort{Desc: "absate4:port51", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port52 = &AbstractPort{Desc: "absate4:port52", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port53 = &AbstractPort{Desc: "absate4:port53", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port54 = &AbstractPort{Desc: "absate4:port54", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port55 = &AbstractPort{Desc: "absate4:port55", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port56 = &AbstractPort{Desc: "absate4:port56", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port57 = &AbstractPort{Desc: "absate4:port57", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port58 = &AbstractPort{Desc: "absate4:port58", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port59 = &AbstractPort{Desc: "absate4:port59", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port60 = &AbstractPort{Desc: "absate4:port60", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port61 = &AbstractPort{Desc: "absate4:port61", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port62 = &AbstractPort{Desc: "absate4:port62", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port63 = &AbstractPort{Desc: "absate4:port63", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port64 = &AbstractPort{Desc: "absate4:port64", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port65 = &AbstractPort{Desc: "absate4:port65", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port66 = &AbstractPort{Desc: "absate4:port66", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port67 = &AbstractPort{Desc: "absate4:port67", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port68 = &AbstractPort{Desc: "absate4:port68", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port69 = &AbstractPort{Desc: "absate4:port69", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
	absate4port70 = &AbstractPort{Desc: "absate4:port70", Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}

	absate4 = &AbstractNode{Desc: "absate4", Ports: []*AbstractPort{
		absate4port1,
		absate4port2,
		absate4port3,
		absate4port4,
		absate4port5,
		absate4port6,
		absate4port7,
		absate4port8,
		absate4port9,
		absate4port10,
		absate4port11,
		absate4port12,
		absate4port13,
		absate4port14,
		absate4port15,
		absate4port16,
		absate4port17,
		absate4port18,
		absate4port19,
		absate4port20,
		absate4port21,
		absate4port22,
		absate4port23,
		absate4port24,
		absate4port25,
		absate4port26,
		absate4port27,
		absate4port28,
		absate4port29,
		absate4port30,
		absate4port31,
		absate4port32,
		absate4port33,
		absate4port34,
		absate4port35,
		absate4port36,
		absate4port37,
		absate4port38,
		absate4port39,
		absate4port40,
		absate4port41,
		absate4port42,
		absate4port43,
		absate4port44,
		absate4port45,
		absate4port46,
		absate4port47,
		absate4port48,
		absate4port49,
		absate4port50,
		absate4port51,
		absate4port52,
		absate4port53,
		absate4port54,
		absate4port55,
		absate4port56,
		absate4port57,
		absate4port58,
		absate4port59,
		absate4port60,
		absate4port61,
		absate4port62,
		absate4port63,
		absate4port64,
		absate4port65,
		absate4port66,
		absate4port67,
		absate4port68,
		absate4port69,
		absate4port70}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
)

var abstractGraph = &AbstractGraph{
	Nodes: []*AbstractNode{absdut1,
		absate1,
		absdut2,
		absate2,
		absdut3,
		absate3,
		absdut4,
		absate4},
	Edges: []*AbstractEdge{
		{absdut1port1, absate1port1},
		{absdut1port2, absate1port2},
		{absdut1port3, absate1port3},
		{absdut1port4, absate1port4},
		{absdut1port5, absate1port5},
		{absdut1port6, absate1port6},
		{absdut1port7, absate1port7},
		{absdut1port8, absate1port8},
		{absdut1port9, absate1port9},
		{absdut1port10, absate1port10},
		{absdut1port11, absate1port11},
		{absdut1port12, absate1port12},
		{absdut1port13, absate1port13},
		{absdut1port14, absate1port14},
		{absdut1port15, absate1port15},
		{absdut1port16, absate1port16},
		{absdut1port17, absate1port17},
		{absdut1port18, absate1port18},
		{absdut1port19, absate1port19},
		{absdut1port20, absate1port20},
		{absdut1port21, absate1port21},
		{absdut1port22, absate1port22},
		{absdut1port23, absate1port23},
		{absdut1port24, absate1port24},
		{absdut1port25, absate1port25},
		{absdut1port26, absate1port26},
		{absdut1port27, absate1port27},
		{absdut1port28, absate1port28},
		{absdut1port29, absate1port29},
		{absdut1port30, absate1port30},
		{absdut1port31, absate1port31},
		{absdut1port32, absate1port32},
		{absdut1port33, absate1port33},
		{absdut1port34, absate1port34},
		{absdut1port35, absate1port35},
		{absdut1port36, absate1port36},
		{absdut1port37, absate1port37},
		{absdut1port38, absate1port38},
		{absdut1port39, absate1port39},
		{absdut1port40, absate1port40},
		{absdut1port41, absate1port41},
		{absdut1port42, absate1port42},
		{absdut1port43, absate1port43},
		{absdut1port44, absate1port44},
		{absdut1port45, absate1port45},
		{absdut1port46, absate1port46},
		{absdut1port47, absate1port47},
		{absdut1port48, absate1port48},
		{absdut1port49, absate1port49},
		{absdut1port50, absate1port50},
		{absdut1port51, absate1port51},
		{absdut1port52, absate1port52},
		{absdut1port53, absate1port53},
		{absdut1port54, absate1port54},
		{absdut1port55, absate1port55},
		{absdut1port56, absate1port56},
		{absdut1port57, absate1port57},
		{absdut1port58, absate1port58},
		{absdut1port59, absate1port59},
		{absdut1port60, absate1port60},
		{absdut1port61, absate1port61},
		{absdut1port62, absate1port62},
		{absdut1port63, absate1port63},
		{absdut1port64, absate1port64},
		{absdut1port65, absate1port65},
		{absdut1port66, absate1port66},
		{absdut1port67, absate1port67},
		{absdut1port68, absate1port68},
		{absdut1port69, absate1port69},
		{absdut1port70, absate1port70},

		{absdut2port1, absate2port1},
		{absdut2port2, absate2port2},
		{absdut2port3, absate2port3},
		{absdut2port4, absate2port4},
		{absdut2port5, absate2port5},
		{absdut2port6, absate2port6},
		{absdut2port7, absate2port7},
		{absdut2port8, absate2port8},
		{absdut2port9, absate2port9},
		{absdut2port10, absate2port10},
		{absdut2port11, absate2port11},
		{absdut2port12, absate2port12},
		{absdut2port13, absate2port13},
		{absdut2port14, absate2port14},
		{absdut2port15, absate2port15},
		{absdut2port16, absate2port16},
		{absdut2port17, absate2port17},
		{absdut2port18, absate2port18},
		{absdut2port19, absate2port19},
		{absdut2port20, absate2port20},
		{absdut2port21, absate2port21},
		{absdut2port22, absate2port22},
		{absdut2port23, absate2port23},
		{absdut2port24, absate2port24},
		{absdut2port25, absate2port25},
		{absdut2port26, absate2port26},
		{absdut2port27, absate2port27},
		{absdut2port28, absate2port28},
		{absdut2port29, absate2port29},
		{absdut2port30, absate2port30},
		{absdut2port31, absate2port31},
		{absdut2port32, absate2port32},
		{absdut2port33, absate2port33},
		{absdut2port34, absate2port34},
		{absdut2port35, absate2port35},
		{absdut2port36, absate2port36},
		{absdut2port37, absate2port37},
		{absdut2port38, absate2port38},
		{absdut2port39, absate2port39},
		{absdut2port40, absate2port40},
		{absdut2port41, absate2port41},
		{absdut2port42, absate2port42},
		{absdut2port43, absate2port43},
		{absdut2port44, absate2port44},
		{absdut2port45, absate2port45},
		{absdut2port46, absate2port46},
		{absdut2port47, absate2port47},
		{absdut2port48, absate2port48},
		{absdut2port49, absate2port49},
		{absdut2port50, absate2port50},
		{absdut2port51, absate2port51},
		{absdut2port52, absate2port52},
		{absdut2port53, absate2port53},
		{absdut2port54, absate2port54},
		{absdut2port55, absate2port55},
		{absdut2port56, absate2port56},
		{absdut2port57, absate2port57},
		{absdut2port58, absate2port58},
		{absdut2port59, absate2port59},
		{absdut2port60, absate2port60},
		{absdut2port61, absate2port61},
		{absdut2port62, absate2port62},
		{absdut2port63, absate2port63},
		{absdut2port64, absate2port64},
		{absdut2port65, absate2port65},
		{absdut2port66, absate2port66},
		{absdut2port67, absate2port67},
		{absdut2port68, absate2port68},
		{absdut2port69, absate2port69},
		{absdut2port70, absate2port70},

		{absdut3port1, absate3port1},
		{absdut3port2, absate3port2},
		{absdut3port3, absate3port3},
		{absdut3port4, absate3port4},
		{absdut3port5, absate3port5},
		{absdut3port6, absate3port6},
		{absdut3port7, absate3port7},
		{absdut3port8, absate3port8},
		{absdut3port9, absate3port9},
		{absdut3port10, absate3port10},
		{absdut3port11, absate3port11},
		{absdut3port12, absate3port12},
		{absdut3port13, absate3port13},
		{absdut3port14, absate3port14},
		{absdut3port15, absate3port15},
		{absdut3port16, absate3port16},
		{absdut3port17, absate3port17},
		{absdut3port18, absate3port18},
		{absdut3port19, absate3port19},
		{absdut3port20, absate3port20},
		{absdut3port21, absate3port21},
		{absdut3port22, absate3port22},
		{absdut3port23, absate3port23},
		{absdut3port24, absate3port24},
		{absdut3port25, absate3port25},
		{absdut3port26, absate3port26},
		{absdut3port27, absate3port27},
		{absdut3port28, absate3port28},
		{absdut3port29, absate3port29},
		{absdut3port30, absate3port30},
		{absdut3port31, absate3port31},
		{absdut3port32, absate3port32},
		{absdut3port33, absate3port33},
		{absdut3port34, absate3port34},
		{absdut3port35, absate3port35},
		{absdut3port36, absate3port36},
		{absdut3port37, absate3port37},
		{absdut3port38, absate3port38},
		{absdut3port39, absate3port39},
		{absdut3port40, absate3port40},
		{absdut3port41, absate3port41},
		{absdut3port42, absate3port42},
		{absdut3port43, absate3port43},
		{absdut3port44, absate3port44},
		{absdut3port45, absate3port45},
		{absdut3port46, absate3port46},
		{absdut3port47, absate3port47},
		{absdut3port48, absate3port48},
		{absdut3port49, absate3port49},
		{absdut3port50, absate3port50},
		{absdut3port51, absate3port51},
		{absdut3port52, absate3port52},
		{absdut3port53, absate3port53},
		{absdut3port54, absate3port54},
		{absdut3port55, absate3port55},
		{absdut3port56, absate3port56},
		{absdut3port57, absate3port57},
		{absdut3port58, absate3port58},
		{absdut3port59, absate3port59},
		{absdut3port60, absate3port60},
		{absdut3port61, absate3port61},
		{absdut3port62, absate3port62},
		{absdut3port63, absate3port63},
		{absdut3port64, absate3port64},
		{absdut3port65, absate3port65},
		{absdut3port66, absate3port66},
		{absdut3port67, absate3port67},
		{absdut3port68, absate3port68},
		{absdut3port69, absate3port69},
		{absdut3port70, absate3port70},

		{absdut4port1, absate4port1},
		{absdut4port2, absate4port2},
		{absdut4port3, absate4port3},
		{absdut4port4, absate4port4},
		{absdut4port5, absate4port5},
		{absdut4port6, absate4port6},
		{absdut4port7, absate4port7},
		{absdut4port8, absate4port8},
		{absdut4port9, absate4port9},
		{absdut4port10, absate4port10},
		{absdut4port11, absate4port11},
		{absdut4port12, absate4port12},
		{absdut4port13, absate4port13},
		{absdut4port14, absate4port14},
		{absdut4port15, absate4port15},
		{absdut4port16, absate4port16},
		{absdut4port17, absate4port17},
		{absdut4port18, absate4port18},
		{absdut4port19, absate4port19},
		{absdut4port20, absate4port20},
		{absdut4port21, absate4port21},
		{absdut4port22, absate4port22},
		{absdut4port23, absate4port23},
		{absdut4port24, absate4port24},
		{absdut4port25, absate4port25},
		{absdut4port26, absate4port26},
		{absdut4port27, absate4port27},
		{absdut4port28, absate4port28},
		{absdut4port29, absate4port29},
		{absdut4port30, absate4port30},
		{absdut4port31, absate4port31},
		{absdut4port32, absate4port32},
		{absdut4port33, absate4port33},
		{absdut4port34, absate4port34},
		{absdut4port35, absate4port35},
		{absdut4port36, absate4port36},
		{absdut4port37, absate4port37},
		{absdut4port38, absate4port38},
		{absdut4port39, absate4port39},
		{absdut4port40, absate4port40},
		{absdut4port41, absate4port41},
		{absdut4port42, absate4port42},
		{absdut4port43, absate4port43},
		{absdut4port44, absate4port44},
		{absdut4port45, absate4port45},
		{absdut4port46, absate4port46},
		{absdut4port47, absate4port47},
		{absdut4port48, absate4port48},
		{absdut4port49, absate4port49},
		{absdut4port50, absate4port50},
		{absdut4port51, absate4port51},
		{absdut4port52, absate4port52},
		{absdut4port53, absate4port53},
		{absdut4port54, absate4port54},
		{absdut4port55, absate4port55},
		{absdut4port56, absate4port56},
		{absdut4port57, absate4port57},
		{absdut4port58, absate4port58},
		{absdut4port59, absate4port59},
		{absdut4port60, absate4port60},
		{absdut4port61, absate4port61},
		{absdut4port62, absate4port62},
		{absdut4port63, absate4port63},
		{absdut4port64, absate4port64},
		{absdut4port65, absate4port65},
		{absdut4port66, absate4port66},
		{absdut4port67, absate4port67},
		{absdut4port68, absate4port68},
		{absdut4port69, absate4port69},
		{absdut4port70, absate4port70}},
}

func TestSolveTestHybrid70(t *testing.T) {
	totalNodes := 20
	startTime := time.Now()
	a, err := Solve(context.Background(), abstractGraph, superGraphT)
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
