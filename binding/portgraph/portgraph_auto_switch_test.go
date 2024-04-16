// package portgraph

// import (
// 	"fmt"
// 	// "strings"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func TestSolveTest1(t *testing.T) {
// 	var (
// 		duts []*ConcreteNode
// 		ates []*ConcreteNode
// 		switches []*ConcreteNode
// 		switchPorts []*ConcretePort
// 		edges []*ConcreteEdge
// 		// abstractNodes []*AbstractNode
// 		// abstractEdges []*AbstractEdge
// 	)

// 	// Create 4 DUTs
// 	for i := 1; i <= 10000; i++ {
// 		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port1", i), Attrs: map[string]string{"speed": "S_400GB", "reserved": "no"}}
// 		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i), Ports: []*ConcretePort{dutPort}, Attrs: map[string]string{"vendor": "CISCO", "reserved": "no"}}
// 		duts = append(duts, dut)

// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)
// 		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 	}

// 	// Create 4 ATEs
// 	for i := 1; i <= 10000; i++ {
// 		atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port1", i), Attrs: map[string]string{"speed": "S_400GB", "reserved": "no"}}
// 		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i), Ports: []*ConcretePort{atePort}, Attrs: map[string]string{"vendor": "TGEN", "reserved": "no"}}
// 		ates = append(ates, ate)

// 		// Add edge from switch to ATE
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)
// 		edges = append(edges, &ConcreteEdge{Src: switchPort, Dst: atePort})
// 	}

// 	// Create one switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: switchPorts, // Now switchPorts has the correct data
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}
// 	switches = append(switches, switchNode)

// 	// Create the super graph
// 	superGraphTest := &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: append(append(duts, ates...), switches...),
// 		Edges: edges,
// 	}

// 	abstractNodes, abstractEdges := createAbstractNodesAndEdges()
// 	fmt.Println("Mohan")
// 	// for i := 1; i <= 100; i++ {
// 	// 	// Create abstract nodes
// 	// 	abstPortDesc := fmt.Sprintf("abst%d:port1", i)
// 	// 	abstPort := &AbstractEdge{
// 	// 		Desc:        abstPortDesc,
// 	// 		Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")},
// 	// 	}
// 	// 	// abstractPorts = append(abstractPorts, port)
// 	// 	abstNodeDesc := fmt.Sprintf("abst%d", i)
// 	// 	abstNode := &AbstractNode{
// 	// 		Desc:        abstNodeDesc,
// 	// 		Ports:       []*AbstractPort{abstPort},
// 	// 		Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "reserved": Equal("no")},
// 	// 	}
// 	// 	abstractNodes = append(abstractNodes, abstNode)
// 	// }
// 	// abstractEdges := []*AbstractEdge{}
// 	// for i := 1; i <= 5000; i++ {
// 	// 	// Create abstract nodes and ports for ATEs
// 	// 	abstEdgeDesc1 := fmt.Sprintf("abst%d:port1", i)
// 	// 	abstEdgeDesc2 := fmt.Sprintf("abst%d:port1", i+1) // Calculate port number for abst4
// 	// 	abstPort1 := &AbstractPort{
// 	// 		Desc:        abstEdgeDesc1,
// 	// 		Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")},
// 	// 	}
// 	// 	abstPort2 := &AbstractPort{
// 	// 		Desc:        abstEdgeDesc2,
// 	// 		Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")},
// 	// 	}
// 	// 	// Append pair of ports to abstractEdges
// 	// 	abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstPort1, Dst: abstPort2})
// 	// }

// 	tests := []struct {
// 		desc            string
// 		graph           *AbstractGraph
// 		wantNodes       map[*AbstractNode]*ConcreteNode
// 		wantPorts       map[*AbstractPort]*ConcretePort
// 		wantSolvedPorts []*AbstractPort
// 	}{{
// 		desc: "Four nodes, interconnected via Switch",
// 		graph: &AbstractGraph{
// 			Desc:  "Four nodes, interconnected via Switch",
// 			Nodes: abstractNodes,
// 			Edges: abstractEdges,
// 			// Nodes: []*AbstractNode{abst1, abst2, abst3, abst4},
// 			// Edges: []*AbstractEdge{{abst1port1, abst3port1}, {abst1port2, abst3port2}, {abst2port1, abst4port1}, {abst2port2, abst4port2}},

// 		},
// 	}}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			startTime := time.Now()
// 			a, err := Solve(context.Background(), tc.graph, superGraphTest)
// 			endTime := time.Now()
// 			if err != nil {
// 				t.Fatalf("Solve got error %v, want nil", err)
// 			}
// 			fmt.Println("Length of Node: ", len(a.Node2Node))
// 			if len(a.Node2Node) != 10000 {
// 				t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 4)
// 			}
// 			if len(a.Port2Port) != 10000 {
// 				t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), (len(tc.wantPorts) + len(tc.wantSolvedPorts)))
// 			}
// 			elapsed := endTime.Sub(startTime)
// 			fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 			fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// 		})
// 	}

// }
// func createAbstractNodesAndEdges() ([]*AbstractNode, []*AbstractEdge) {
// 	var abstractNodes []*AbstractNode
// 	var abstractPorts []*AbstractPort
// 	for i := 1; i <= 10000; i++ {
// 		abstPortDesc := fmt.Sprintf("abst%d:port1", i)
// 		abstPort := &AbstractPort{Desc: abstPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")}}
// 		abstractPorts = append(abstractPorts, abstPort)
// 		abstNodeDesc := fmt.Sprintf("abst%d", i)
// 		abstNode := &AbstractNode{Desc: abstNodeDesc, Ports: []*AbstractPort{abstPort}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "reserved": Equal("no")}}
// 		abstractNodes = append(abstractNodes, abstNode)
// 	}
// 	for i := 10000; i <= 20000; i++ {
// 		abstPortDesc := fmt.Sprintf("abst%d:port1", i)
// 		abstPort := &AbstractPort{Desc: abstPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")}}
// 		abstractPorts = append(abstractPorts, abstPort)
// 		abstNodeDesc := fmt.Sprintf("abst%d", i)
// 		abstNode := &AbstractNode{Desc: abstNodeDesc, Ports: []*AbstractPort{abstPort}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN"), "reserved": Equal("no")}}
// 		abstractNodes = append(abstractNodes, abstNode)
// 	}

// 	var abstractEdges []*AbstractEdge
// 	// for i := 1; i <= 5000; i++ {
// 	// 	abstEdgeDesc1 := fmt.Sprintf("abst%d:port1", i)
// 	// 	abstEdgeDesc2 := fmt.Sprintf("abst%d:port1", i+1)
// 	// 	abstPort1 := &AbstractPort{Desc: abstEdgeDesc1, Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")}}
// 	// 	abstPort2 := &AbstractPort{Desc: abstEdgeDesc2, Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")}}
// 	// 	abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstPort1, Dst: abstPort2})
// 	// }
// 	return abstractNodes, abstractEdges
// }

// package portgraph

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func TestSolveTest1(t *testing.T) {
// 	var (
// 		duts        []*ConcreteNode
// 		ates        []*ConcreteNode
// 		switches    []*ConcreteNode
// 		switchPorts []*ConcretePort
// 		edges       []*ConcreteEdge
// 	)
// 	totalNodes := 100
// 	// Create 4 DUTs
// 	for i := 1; i <= totalNodes; i++ {
// 		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port1", i), Attrs: map[string]string{"speed": "S_400GB", "reserved": "no"}}
// 		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i), Ports: []*ConcretePort{dutPort}, Attrs: map[string]string{"vendor": "CISCO", "reserved": "no"}}
// 		duts = append(duts, dut)

// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)
// 		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 	}

// 	// Create 4 ATEs
// 	for i := 1; i <= totalNodes; i++ {
// 		atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port1", i), Attrs: map[string]string{"speed": "S_400GB", "reserved": "no"}}
// 		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i), Ports: []*ConcretePort{atePort}, Attrs: map[string]string{"vendor": "TGEN", "reserved": "no"}}
// 		ates = append(ates, ate)

// 		// Add edge from switch to ATE
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)
// 		edges = append(edges, &ConcreteEdge{Src: switchPort, Dst: atePort})
// 	}

// 	// Create one switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}
// 	switches = append(switches, switchNode)

// 	// Create the super graph
// 	superGraphTest1 := &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: append(append(duts, ates...), switches...),
// 		Edges: edges,
// 	}

// 	abstractNodes, abstractEdges := createAbstractNodesAndEdges(totalNodes)
// 	tests := []struct {
// 		desc            string
// 		graph           *AbstractGraph
// 		wantNodes       map[*AbstractNode]*ConcreteNode
// 		wantPorts       map[*ConcretePort]*AbstractPort
// 		wantSolvedPorts []*AbstractPort
// 	}{{
// 		desc: "Four nodes, interconnected via Switch",
// 		graph: &AbstractGraph{
// 			Desc:  "Four nodes, interconnected via Switch",
// 			Nodes: abstractNodes,
// 			Edges: abstractEdges,
// 		},
// 	}}
// 	for _, tc := range tests {
// 		t.Run(tc.desc, func(t *testing.T) {
// 			startTime := time.Now()
// 			a, err := Solve(context.Background(), tc.graph, superGraphTest1)
// 			endTime := time.Now()
// 			if err != nil {
// 				t.Fatalf("Solve got error %v, want nil", err)
// 			}
// 			fmt.Println("Length of Node: ", len(a.Node2Node))
// 			if len(a.Node2Node) != totalNodes*2 {
// 				t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 4)
// 			}
// 			if len(a.Port2Port) != totalNodes*2 {
// 				t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), (len(tc.wantPorts) + len(tc.wantSolvedPorts)))
// 			}
// 			elapsed := endTime.Sub(startTime)
// 			fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 			fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// 		})
// 	}

// }
// func createAbstractNodesAndEdges(totalNodes int) ([]*AbstractNode, []*AbstractEdge) {
// 	var abstractNodes []*AbstractNode
// 	var abstractPorts []*AbstractPort
// 	halfNodes := totalNodes
// 	fullNodes := totalNodes * 2
// 	for i := 1; i <= halfNodes; i++ {
// 		abstPortDesc := fmt.Sprintf("abst%d:port1", i)
// 		abstPort := &AbstractPort{Desc: abstPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")}}
// 		abstractPorts = append(abstractPorts, abstPort)
// 		abstNodeDesc := fmt.Sprintf("abst%d", i)
// 		abstNode := &AbstractNode{Desc: abstNodeDesc, Ports: []*AbstractPort{abstPort}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO"), "reserved": Equal("no")}}
// 		abstractNodes = append(abstractNodes, abstNode)
// 	}
// 	for i := halfNodes + 1; i <= fullNodes; i++ {
// 		abstPortDesc := fmt.Sprintf("abst%d:port1", i)
// 		abstPort := &AbstractPort{Desc: abstPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400GB"), "reserved": Equal("no")}}
// 		abstractPorts = append(abstractPorts, abstPort)
// 		abstNodeDesc := fmt.Sprintf("abst%d", i)
// 		abstNode := &AbstractNode{Desc: abstNodeDesc, Ports: []*AbstractPort{abstPort}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN"), "reserved": Equal("no")}}
// 		abstractNodes = append(abstractNodes, abstNode)
// 	}

// 	var abstractEdges []*AbstractEdge
// 	count := len(abstractPorts) / 2
// 	for i := 0; i < count; i++ {
// 		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractPorts[i], Dst: abstractPorts[i+count]})
// 	}
// 	return abstractNodes, abstractEdges
// }

// // Single switch test
// package portgraph

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func generateConcreteGraph(atePortsCount int, dutPortsCount int) *ConcreteGraph {
// 	var (
// 		nodes       []*ConcreteNode
// 		dutports    []*ConcretePort
// 		ateports    []*ConcretePort
// 		switchPorts []*ConcretePort
// 		edges       []*ConcreteEdge
// 	)

// 	for i := 1; i <= dutPortsCount; i++ {
// 		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
// 		dutports = append(dutports, dutPort)
// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)

// 		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 	}
// 	dut := &ConcreteNode{Desc: "dut1", Ports: dutports, Attrs: map[string]string{"vendor": "CISCO"}}
// 	// nodes = append(nodes, dut)
// 	for i := 1; i <= atePortsCount; i++ {
// 		atePort := &ConcretePort{Desc: fmt.Sprintf("ate1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
// 		ateports = append(ateports, atePort)
// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i+atePortsCount)}
// 		switchPorts = append(switchPorts, switchPort)

// 		edges = append(edges, &ConcreteEdge{Src: switchPort, Dst: atePort})
// 	}
// 	ate := &ConcreteNode{Desc: "ate1", Ports: ateports, Attrs: map[string]string{"vendor": "TGEN"}}

// 	// Create one switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}
// 	nodes = append(nodes, dut, ate, switchNode)
// 	return &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: nodes,
// 		Edges: edges,
// 	}

// }

// func createAbstractNodesAndEdges(dutPortsCount int, atePortsCount int) *AbstractGraph {
// 	var abstractNodes []*AbstractNode
// 	var abstractDutPorts []*AbstractPort
// 	var abstractAtePorts []*AbstractPort
// 	// halfNodes := totalNodes
// 	// fullNodes := totalNodes * 2
// 	for i := 1; i <= dutPortsCount; i++ {
// 		abstdutPortDesc := fmt.Sprintf("abstdut1:port%d", i)

// 		abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 		abstractDutPorts = append(abstractDutPorts, abstdutPort)
// 	}
// 	abstdutNode := &AbstractNode{Desc: "abstdut1", Ports: abstractDutPorts, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
// 	abstractNodes = append(abstractNodes, abstdutNode)
// 	for i := 1; i <= atePortsCount; i++ {
// 		abstatePortDesc := fmt.Sprintf("abstate1:port%d", i)
// 		abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 		abstractAtePorts = append(abstractAtePorts, abstatePort)
// 	}
// 	abstateNode := &AbstractNode{Desc: "abstate1", Ports: abstractAtePorts, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
// 	abstractNodes = append(abstractNodes, abstateNode)
// 	var abstractEdges []*AbstractEdge
// 	count := len(abstractDutPorts)
// 	for i := 0; i < count; i++ {
// 		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractDutPorts[i], Dst: abstractAtePorts[i]})
// 	}

// 	return &AbstractGraph{
// 		Desc:  "",
// 		Nodes: abstractNodes,
// 		Edges: abstractEdges,
// 	}
// }

// func TestSolveTest1(t *testing.T) {
// 	totalNodes := 1
// 	atePortsCount := 70
// 	dutPortsCount := 70
// 	graph := createAbstractNodesAndEdges(dutPortsCount, atePortsCount)
// 	conGraph := generateConcreteGraph(atePortsCount, dutPortsCount)
// 	startTime := time.Now()
// 	a, err := Solve(context.Background(), graph, conGraph)
// 	endTime := time.Now()
// 	if err != nil {
// 		t.Fatalf("Solve got error %v, want nil", err)
// 	}
// 	if len(a.Node2Node) != 2 {
// 		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
// 	}
// 	if len(a.Port2Port) != 2*dutPortsCount {
// 		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
// 	}
// 	elapsed := endTime.Sub(startTime)
// 	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// }

// // Single switch with duts 70 prts and ate 58 ports
// package portgraph

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func generateConcreteGraph(atePortsCount int, dutPortsCount int) *ConcreteGraph {
// 	var (
// 		nodes       []*ConcreteNode
// 		dutports    []*ConcretePort
// 		ateports    []*ConcretePort
// 		switchPorts []*ConcretePort
// 		edges       []*ConcreteEdge
// 	)

// 	for i := 1; i <= dutPortsCount; i++ {
// 		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
// 		dutports = append(dutports, dutPort)
// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)

// 		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 	}
// 	dut := &ConcreteNode{Desc: "dut1", Ports: dutports, Attrs: map[string]string{"vendor": "CISCO"}}
// 	// nodes = append(nodes, dut)
// 	for i := 1; i <= atePortsCount; i++ {
// 		atePort := &ConcretePort{Desc: fmt.Sprintf("ate1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
// 		ateports = append(ateports, atePort)
// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i+atePortsCount)}
// 		switchPorts = append(switchPorts, switchPort)

// 		edges = append(edges, &ConcreteEdge{Src: switchPort, Dst: atePort})
// 	}
// 	ate := &ConcreteNode{Desc: "ate1", Ports: ateports, Attrs: map[string]string{"vendor": "TGEN"}}

// 	// Create one switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}
// 	nodes = append(nodes, dut, ate, switchNode)
// 	return &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: nodes,
// 		Edges: edges,
// 	}

// }

// func createAbstractNodesAndEdges(dutPortsCount int, atePortsCount int) *AbstractGraph {
// 	var abstractNodes []*AbstractNode
// 	var abstractDutPorts []*AbstractPort
// 	var abstractAtePorts []*AbstractPort
// 	// halfNodes := totalNodes
// 	// fullNodes := totalNodes * 2
// 	for i := 1; i <= atePortsCount; i++ {
// 		abstdutPortDesc := fmt.Sprintf("abstdut1:port%d", i)

// 		abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 		abstractDutPorts = append(abstractDutPorts, abstdutPort)
// 	}
// 	abstdutNode := &AbstractNode{Desc: "abstdut1", Ports: abstractDutPorts, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
// 	abstractNodes = append(abstractNodes, abstdutNode)
// 	for i := 1; i <= atePortsCount; i++ {
// 		abstatePortDesc := fmt.Sprintf("abstate1:port%d", i)
// 		abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 		abstractAtePorts = append(abstractAtePorts, abstatePort)
// 	}
// 	abstateNode := &AbstractNode{Desc: "abstate1", Ports: abstractAtePorts, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
// 	abstractNodes = append(abstractNodes, abstateNode)
// 	var abstractEdges []*AbstractEdge
// 	count := len(abstractDutPorts)
// 	for i := 0; i < count; i++ {
// 		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractDutPorts[i], Dst: abstractAtePorts[i]})
// 	}

// 	return &AbstractGraph{
// 		Desc:  "",
// 		Nodes: abstractNodes,
// 		Edges: abstractEdges,
// 	}
// }

// func TestSolveTest1(t *testing.T) {
// 	totalNodes := 1
// 	atePortsCount := 58
// 	dutPortsCount := 70
// 	graph := createAbstractNodesAndEdges(dutPortsCount, atePortsCount)
// 	conGraph := generateConcreteGraph(atePortsCount, dutPortsCount)
// 	startTime := time.Now()
// 	a, err := Solve(context.Background(), graph, conGraph)
// 	endTime := time.Now()
// 	if err != nil {
// 		t.Fatalf("Solve got error %v, want nil", err)
// 	}
// 	if len(a.Node2Node) != 2 {
// 		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
// 	}
// 	if len(a.Port2Port) != 2*atePortsCount {
// 		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
// 	}
// 	elapsed := endTime.Sub(startTime)
// 	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// }

// // Single switch with duts 58 prts and ate 70 ports

// package portgraph

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func generateConcreteGraph(atePortsCount int, dutPortsCount int) *ConcreteGraph {
// 	var (
// 		nodes       []*ConcreteNode
// 		dutports    []*ConcretePort
// 		ateports    []*ConcretePort
// 		switchPorts []*ConcretePort
// 		edges       []*ConcreteEdge
// 	)

// 	for i := 1; i <= dutPortsCount; i++ {
// 		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
// 		dutports = append(dutports, dutPort)
// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
// 		switchPorts = append(switchPorts, switchPort)

// 		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 	}
// 	dut := &ConcreteNode{Desc: "dut1", Ports: dutports, Attrs: map[string]string{"vendor": "CISCO"}}
// 	// nodes = append(nodes, dut)
// 	for i := 1; i <= atePortsCount; i++ {
// 		atePort := &ConcretePort{Desc: fmt.Sprintf("ate1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
// 		ateports = append(ateports, atePort)
// 		// Add edge from DUT to switch
// 		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i+atePortsCount)}
// 		switchPorts = append(switchPorts, switchPort)

// 		edges = append(edges, &ConcreteEdge{Src: switchPort, Dst: atePort})
// 	}
// 	ate := &ConcreteNode{Desc: "ate1", Ports: ateports, Attrs: map[string]string{"vendor": "TGEN"}}

// 	// Create one switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}
// 	nodes = append(nodes, dut, ate, switchNode)
// 	return &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: nodes,
// 		Edges: edges,
// 	}

// }

// func createAbstractNodesAndEdges(dutPortsCount int, atePortsCount int) *AbstractGraph {
// 	var abstractNodes []*AbstractNode
// 	var abstractDutPorts []*AbstractPort
// 	var abstractAtePorts []*AbstractPort
// 	// halfNodes := totalNodes
// 	// fullNodes := totalNodes * 2
// 	for i := 1; i <= dutPortsCount; i++ {
// 		abstdutPortDesc := fmt.Sprintf("abstdut1:port%d", i)

// 		abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 		abstractDutPorts = append(abstractDutPorts, abstdutPort)
// 	}
// 	abstdutNode := &AbstractNode{Desc: "abstdut1", Ports: abstractDutPorts, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
// 	abstractNodes = append(abstractNodes, abstdutNode)
// 	for i := 1; i <= dutPortsCount; i++ {
// 		abstatePortDesc := fmt.Sprintf("abstate1:port%d", i)
// 		abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 		abstractAtePorts = append(abstractAtePorts, abstatePort)
// 	}
// 	abstateNode := &AbstractNode{Desc: "abstate1", Ports: abstractAtePorts, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
// 	abstractNodes = append(abstractNodes, abstateNode)
// 	var abstractEdges []*AbstractEdge
// 	count := len(abstractDutPorts)
// 	for i := 0; i < count; i++ {
// 		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractDutPorts[i], Dst: abstractAtePorts[i]})
// 	}

// 	return &AbstractGraph{
// 		Desc:  "",
// 		Nodes: abstractNodes,
// 		Edges: abstractEdges,
// 	}
// }

// func TestSolveTest1(t *testing.T) {
// 	totalNodes := 1
// 	atePortsCount := 70
// 	dutPortsCount := 58
// 	graph := createAbstractNodesAndEdges(dutPortsCount, atePortsCount)
// 	conGraph := generateConcreteGraph(atePortsCount, dutPortsCount)
// 	startTime := time.Now()
// 	a, err := Solve(context.Background(), graph, conGraph)
// 	endTime := time.Now()
// 	if err != nil {
// 		t.Fatalf("Solve got error %v, want nil", err)
// 	}
// 	if len(a.Node2Node) != 2 {
// 		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
// 	}
// 	if len(a.Port2Port) != 2*dutPortsCount {
// 		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
// 	}
// 	elapsed := endTime.Sub(startTime)
// 	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// }

// // // 10 duts, 10 ates and on switch
// package portgraph

// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	"golang.org/x/net/context"
// )

// func generateConcreteGraph(dutCount int, ateCount int, atePortsCount int, dutPortsCount int) *ConcreteGraph {
// 	var (
// 		nodes       []*ConcreteNode
// 		switchPorts []*ConcretePort
// 		edges       []*ConcreteEdge
// 	)
// 	switchCount := 1
// 	for i := 1; i <= dutCount; i++ {
// 		var dutports []*ConcretePort
// 		if i > 5 && i < 9 {
// 			dutPortsCount = 6
// 		} else if i > 9 {
// 			dutPortsCount = 5
// 		}
// 		for j := 1; j <= dutPortsCount; j++ {
// 			dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i, j), Attrs: map[string]string{"speed": "S_400G"}}
// 			dutports = append(dutports, dutPort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount)}
// 			switchPorts = append(switchPorts, switchPort)
// 			switchCount++
// 			edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
// 		}
// 		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i), Ports: dutports, Attrs: map[string]string{"vendor": "CISCO", "reserved": "no"}}
// 		nodes = append(nodes, dut)
// 	}
// 	for i := 1; i <= ateCount; i++ {
// 		var ateports []*ConcretePort
// 		if i > 5 && i < 9 {
// 			atePortsCount = 6
// 		} else if i > 9 {
// 			atePortsCount = 5
// 		}
// 		for j := 1; j <= atePortsCount; j++ {
// 			atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port%d", i, j), Attrs: map[string]string{"speed": "S_400G"}}
// 			ateports = append(ateports, atePort)
// 			// Add edge from DUT to switch
// 			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount)}
// 			switchPorts = append(switchPorts, switchPort)
// 			switchCount++
// 			edges = append(edges, &ConcreteEdge{Src: atePort, Dst: switchPort})
// 		}
// 		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i), Ports: ateports, Attrs: map[string]string{"vendor": "TGEN", "reserved": "no"}}
// 		nodes = append(nodes, ate)
// 	}

// 	// Create one switch
// 	switchNode := &ConcreteNode{
// 		Desc:  "switch1",
// 		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
// 		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
// 	}
// 	nodes = append(nodes, switchNode)
// 	return &ConcreteGraph{
// 		Desc:  "super",
// 		Nodes: nodes,
// 		Edges: edges,
// 	}

// }

// func createAbstractNodesAndEdges(dutCount int, ateCount int, dutPortsCount int, atePortsCount int) *AbstractGraph {
// 	var abstractNodes []*AbstractNode
// 	var abstractDutPortsAll []*AbstractPort
// 	var abstractAtePortsAll []*AbstractPort
// 	// halfNodes := totalNodes
// 	// fullNodes := totalNodes * 2
// 	for i := 1; i <= dutCount; i++ {
// 		var abstractDutPorts []*AbstractPort
// 		if i > 5 && i < 9 {
// 			dutPortsCount = 6
// 		} else if i > 9 {
// 			dutPortsCount = 5
// 		}
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
// 		if i > 5 && i < 9 {
// 			atePortsCount = 6
// 		} else if i > 9 {
// 			atePortsCount = 5
// 		}
// 		for j := 1; j <= atePortsCount; j++ {
// 			abstatePortDesc := fmt.Sprintf("abstate%d:port%d", i, j)
// 			abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
// 			abstractAtePorts = append(abstractAtePorts, abstatePort)
// 			abstractAtePortsAll = append(abstractAtePortsAll, abstatePort)
// 		}
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

// func TestSolveTest1(t *testing.T) {
// 	totalNodes := 1
// 	ateCount := 10
// 	dutCount := 10
// 	atePortsCount := 7
// 	dutPortsCount := 7
// 	graph := createAbstractNodesAndEdges(dutCount, ateCount, dutPortsCount, atePortsCount)
// 	conGraph := generateConcreteGraph(dutCount, ateCount, atePortsCount, dutPortsCount)
// 	startTime := time.Now()
// 	a, err := Solve(context.Background(), graph, conGraph)
// 	endTime := time.Now()
// 	if err != nil {
// 		t.Fatalf("Solve got error %v, want nil", err)
// 	}
// 	if len(a.Node2Node) != 20 {
// 		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
// 	}
// 	if len(a.Port2Port) != 128 {
// 		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
// 	}
// 	elapsed := endTime.Sub(startTime)
// 	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
// 	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
// }


// // 10 duts, 3 ates and on switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func generateConcreteGraph(dutCount int, ateCount int, atePortsCount int, dutPortsCount int) *ConcreteGraph {
	var (
		nodes       []*ConcreteNode
		switchPorts []*ConcretePort
		edges       []*ConcreteEdge
	)
	switchCount := 1
	for i := 1; i <= dutCount; i++ {
		var dutports []*ConcretePort
		if i > 5 && i < 9 {
			dutPortsCount = 6
		} else if i > 9 {
			dutPortsCount = 5
		}
		for j := 1; j <= dutPortsCount; j++ {
			dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port%d", i, j), Attrs: map[string]string{"speed": "S_400G"}}
			dutports = append(dutports, dutPort)
			// Add edge from DUT to switch
			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount)}
			switchPorts = append(switchPorts, switchPort)
			switchCount++
			edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
		}
		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i), Ports: dutports, Attrs: map[string]string{"vendor": "CISCO", "reserved": "no"}}
		nodes = append(nodes, dut)
	}
	for i := 1; i <= ateCount; i++ {
		var ateports []*ConcretePort		
		for j := 1; j <= atePortsCount; j++ {
			atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port%d", i, j), Attrs: map[string]string{"speed": "S_400G"}}
			ateports = append(ateports, atePort)
			// Add edge from DUT to switch
			switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", switchCount)}
			switchPorts = append(switchPorts, switchPort)
			switchCount++
			edges = append(edges, &ConcreteEdge{Src: atePort, Dst: switchPort})
		}
		atePortsCount = 20
		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i), Ports: ateports, Attrs: map[string]string{"vendor": "TGEN", "reserved": "no"}}
		nodes = append(nodes, ate)
	}

	// Create one switch
	switchNode := &ConcreteNode{
		Desc:  "switch1",
		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
	}
	nodes = append(nodes, switchNode)
	return &ConcreteGraph{
		Desc:  "super",
		Nodes: nodes,
		Edges: edges,
	}

}

func createAbstractNodesAndEdges(dutCount int, ateCount int, dutPortsCount int, atePortsCount int) *AbstractGraph {
	var abstractNodes []*AbstractNode
	var abstractDutPortsAll []*AbstractPort
	var abstractAtePortsAll []*AbstractPort
	// halfNodes := totalNodes
	// fullNodes := totalNodes * 2
	for i := 1; i <= dutCount; i++ {
		var abstractDutPorts []*AbstractPort
		if i > 5 && i < 9 {
			dutPortsCount = 6
		} else if i > 9 {
			dutPortsCount = 5
		}
		for j := 1; j <= dutPortsCount; j++ {
			abstdutPortDesc := fmt.Sprintf("abstdut%d:port%d", i, j)
			abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
			abstractDutPorts = append(abstractDutPorts, abstdutPort)
			abstractDutPortsAll = append(abstractDutPortsAll, abstdutPort)
		}
		abstdutNode := &AbstractNode{Desc: fmt.Sprintf("abstdut%d", i), Ports: abstractDutPorts, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
		abstractNodes = append(abstractNodes, abstdutNode)
	}
	for i := 1; i <= ateCount; i++ {
		var abstractAtePorts []*AbstractPort
		
		for j := 1; j <= atePortsCount; j++ {
			abstatePortDesc := fmt.Sprintf("abstate%d:port%d", i, j)
			abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
			abstractAtePorts = append(abstractAtePorts, abstatePort)
			abstractAtePortsAll = append(abstractAtePortsAll, abstatePort)
		}
		atePortsCount = 20
		abstateNode := &AbstractNode{Desc: fmt.Sprintf("abstate%d", i), Ports: abstractAtePorts, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
		abstractNodes = append(abstractNodes, abstateNode)
	}
	var abstractEdges []*AbstractEdge
	count := len(abstractDutPortsAll)
	for i := 0; i < count; i++ {
		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractDutPortsAll[i], Dst: abstractAtePortsAll[i]})
	}

	return &AbstractGraph{
		Desc:  "",
		Nodes: abstractNodes,
		Edges: abstractEdges,
	}
}

func TestSolveTest1(t *testing.T) {
	totalNodes := 1
	ateCount := 3
	dutCount := 10
	atePortsCount := 24
	dutPortsCount := 7
	graph := createAbstractNodesAndEdges(dutCount, ateCount, dutPortsCount, atePortsCount)
	conGraph := generateConcreteGraph(dutCount, ateCount, atePortsCount, dutPortsCount)
	startTime := time.Now()
	a, err := Solve(context.Background(), graph, conGraph)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 13 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 128 {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
