// Single switch test
package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func generateConcreteGraphSwitch(atePortsCount int, dutPortsCount int) *ConcreteGraph {
	var (
		nodes       []*ConcreteNode
		dutports    []*ConcretePort
		ateports    []*ConcretePort
		switchPorts []*ConcretePort
		edges       []*ConcreteEdge
	)

	for i := 1; i <= dutPortsCount; i++ {
		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
		dutports = append(dutports, dutPort)
		// Add edge from DUT to switch
		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i)}
		switchPorts = append(switchPorts, switchPort)

		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: switchPort})
	}
	dut := &ConcreteNode{Desc: "dut1", Ports: dutports, Attrs: map[string]string{"vendor": "CISCO"}}
	// nodes = append(nodes, dut)
	for i := 1; i <= atePortsCount; i++ {
		atePort := &ConcretePort{Desc: fmt.Sprintf("ate1:port%d", i), Attrs: map[string]string{"speed": "S_400G"}}
		ateports = append(ateports, atePort)
		// Add edge from DUT to switch
		switchPort := &ConcretePort{Desc: fmt.Sprintf("switch1:p%d", i+atePortsCount)}
		switchPorts = append(switchPorts, switchPort)

		edges = append(edges, &ConcreteEdge{Src: switchPort, Dst: atePort})
	}
	ate := &ConcreteNode{Desc: "ate1", Ports: ateports, Attrs: map[string]string{"vendor": "TGEN"}}

	// Create one switch
	switchNode := &ConcreteNode{
		Desc:  "switch1",
		Ports: append([]*ConcretePort{}, switchPorts...), // Copy switchPorts to avoid modification
		Attrs: map[string]string{"role": "L1S", "name": "sw1"},
	}
	nodes = append(nodes, dut, ate, switchNode)
	return &ConcreteGraph{
		Desc:  "super",
		Nodes: nodes,
		Edges: edges,
	}

}

func createAbstractNodesAndEdgesSwitch(dutPortsCount int, atePortsCount int) *AbstractGraph {
	var abstractNodes []*AbstractNode
	var abstractDutPorts []*AbstractPort
	var abstractAtePorts []*AbstractPort
	// halfNodes := totalNodes
	// fullNodes := totalNodes * 2
	for i := 1; i <= dutPortsCount; i++ {
		abstdutPortDesc := fmt.Sprintf("abstdut1:port%d", i)

		abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
		abstractDutPorts = append(abstractDutPorts, abstdutPort)
	}
	abstdutNode := &AbstractNode{Desc: "abstdut1", Ports: abstractDutPorts, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
	abstractNodes = append(abstractNodes, abstdutNode)
	for i := 1; i <= atePortsCount; i++ {
		abstatePortDesc := fmt.Sprintf("abstate1:port%d", i)
		abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
		abstractAtePorts = append(abstractAtePorts, abstatePort)
	}
	abstateNode := &AbstractNode{Desc: "abstate1", Ports: abstractAtePorts, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
	abstractNodes = append(abstractNodes, abstateNode)
	var abstractEdges []*AbstractEdge
	count := len(abstractDutPorts)
	for i := 0; i < count; i++ {
		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractDutPorts[i], Dst: abstractAtePorts[i]})
	}

	return &AbstractGraph{
		Desc:  "",
		Nodes: abstractNodes,
		Edges: abstractEdges,
	}
}

func TestSolveTestSwitch(t *testing.T) {
	totalNodes := 1
	atePortsCount := 70
	dutPortsCount := 70
	graph := createAbstractNodesAndEdgesSwitch(dutPortsCount, atePortsCount)
	conGraph := generateConcreteGraphSwitch(atePortsCount, dutPortsCount)
	startTime := time.Now()
	a, err := Solve(context.Background(), graph, conGraph)
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 2 {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 2*dutPortsCount {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
