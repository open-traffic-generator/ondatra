// // 10 duts, 3 ates and on switch
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func generateConcreteGraphAtes(dutCount int, ateCount int, atePortsCount int, dutPortsCount int) *ConcreteGraph {
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

func createAbstractNodesAndEdgesAtes(dutCount int, ateCount int, dutPortsCount int, atePortsCount int) *AbstractGraph {
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

func TestSolveTestAtes(t *testing.T) {
	totalNodes := 1
	ateCount := 3
	dutCount := 10
	atePortsCount := 24
	dutPortsCount := 7
	graph := createAbstractNodesAndEdgesAtes(dutCount, ateCount, dutPortsCount, atePortsCount)
	conGraph := generateConcreteGraphAtes(dutCount, ateCount, atePortsCount, dutPortsCount)
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
