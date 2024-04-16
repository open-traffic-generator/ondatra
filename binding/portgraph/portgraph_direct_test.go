package portgraph

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func generateConcreteGraphDirect(totalNodes int) *ConcreteGraph {
	var (
		nodes []*ConcreteNode
		edges []*ConcreteEdge
	)
	// Create 4 DUTs
	for i := 1; i <= totalNodes; i++ {
		dutPort := &ConcretePort{Desc: fmt.Sprintf("dut%d:port1", i), Attrs: map[string]string{"speed": "S_400G"}}
		dutPort1 := &ConcretePort{Desc: fmt.Sprintf("dut%d:port1", i), Attrs: map[string]string{"speed": "S_400G"}}
		dut := &ConcreteNode{Desc: fmt.Sprintf("dut%d", i), Ports: []*ConcretePort{dutPort, dutPort1}, Attrs: map[string]string{"vendor": "CISCO"}}
		nodes = append(nodes, dut)

		atePort := &ConcretePort{Desc: fmt.Sprintf("ate%d:port1", i), Attrs: map[string]string{"speed": "S_400G"}}
		atePort1 := &ConcretePort{Desc: fmt.Sprintf("ate%d:port1", i), Attrs: map[string]string{"speed": "S_400G"}}
		ate := &ConcreteNode{Desc: fmt.Sprintf("ate%d", i), Ports: []*ConcretePort{atePort, atePort1}, Attrs: map[string]string{"vendor": "TGEN"}}
		nodes = append(nodes, ate)

		edges = append(edges, &ConcreteEdge{Src: dutPort, Dst: atePort}, &ConcreteEdge{Src: dutPort1, Dst: atePort1})
	}

	return &ConcreteGraph{
		Desc:  "super",
		Nodes: nodes,
		Edges: edges,
	}

}

func createAbstractNodesAndEdgesDirect(totalNodes int) *AbstractGraph {
	var abstractNodes []*AbstractNode
	var abstractPorts []*AbstractPort
	halfNodes := totalNodes
	fullNodes := totalNodes * 2
	for i := 1; i <= halfNodes; i++ {
		abstdutPortDesc := fmt.Sprintf("abstdut%d:port1", i)
		abstdutPortDesc1 := fmt.Sprintf("abst%d:port2", i+halfNodes)
		abstdutPort := &AbstractPort{Desc: abstdutPortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
		abstdutPort1 := &AbstractPort{Desc: abstdutPortDesc1, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
		abstractPorts = append(abstractPorts, abstdutPort, abstdutPort1)
		abstdutNodeDesc := fmt.Sprintf("abst%d", i)
		abstdutNode := &AbstractNode{Desc: abstdutNodeDesc, Ports: []*AbstractPort{abstdutPort}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
		abstractNodes = append(abstractNodes, abstdutNode)
	}
	for i := halfNodes + 1; i <= fullNodes; i++ {
		abstatePortDesc := fmt.Sprintf("abstate%d:port1", i+fullNodes)
		abstatePortDesc1 := fmt.Sprintf("abst%d:port1", i+fullNodes+halfNodes)
		abstatePort := &AbstractPort{Desc: abstatePortDesc, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
		abstatePort1 := &AbstractPort{Desc: abstatePortDesc1, Constraints: map[string]PortConstraint{"speed": Equal("S_400G")}}
		abstractPorts = append(abstractPorts, abstatePort, abstatePort1)
		abstateNodeDesc := fmt.Sprintf("abstate%d", i)
		abstateNode := &AbstractNode{Desc: abstateNodeDesc, Ports: []*AbstractPort{abstatePort}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
		abstractNodes = append(abstractNodes, abstateNode)
	}

	var abstractEdges []*AbstractEdge
	count := len(abstractPorts) / 2
	for i := 0; i < count; i++ {
		abstractEdges = append(abstractEdges, &AbstractEdge{Src: abstractPorts[i], Dst: abstractPorts[i+count]})
	}

	return &AbstractGraph{
		Desc:  "",
		Nodes: abstractNodes,
		Edges: abstractEdges,
	}
}

func TestSolveTestDirect(t *testing.T) {
	totalNodes := 2000
	startTime := time.Now()
	a, err := Solve(context.Background(), createAbstractNodesAndEdgesDirect(totalNodes), generateConcreteGraphDirect(totalNodes))
	endTime := time.Now()
	if err != nil {
		t.Fatalf("Solve got error %v, want nil", err)
	}
	if len(a.Node2Node) != 2*totalNodes {
		t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), 2*totalNodes)
	}
	if len(a.Port2Port) != 2*totalNodes {
		t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), 2*totalNodes)
	}
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v seconds\n", elapsed.Seconds())
	fmt.Printf("Execution time: %v milliseconds\n", elapsed.Milliseconds())
}
