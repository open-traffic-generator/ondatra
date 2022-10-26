// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package solver creates solutions from devices and topologies.
package solver

import (
	"fmt"
	"strings"

	log "github.com/golang/glog"
	"github.com/pborman/uuid"
	"github.com/openconfig/ondatra/binding"

	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	// TODO: Cleanup after Type deprecation
	ateTypes = map[tpb.Node_Type]bool{
		tpb.Node_IXIA_TG: true,
	}

	ateVendors = map[tpb.Vendor]bool{
		tpb.Vendor_KEYSIGHT: true,
	}

	deviceTypes = map[tpb.Node_Type]opb.Device_Vendor{
		tpb.Node_ARISTA_CEOS: opb.Device_ARISTA,
		// TODO: when Ondatra supports the OS dimension, use it to
		// distinguish CSR from CXR and CEVO from VMX.
		tpb.Node_CISCO_CSR:    opb.Device_CISCO,
		tpb.Node_CISCO_CXR:    opb.Device_CISCO,
		tpb.Node_CISCO_E8000:  opb.Device_CISCO,
		tpb.Node_CISCO_XRD:    opb.Device_CISCO,
		tpb.Node_IXIA_TG:      opb.Device_IXIA,
		tpb.Node_JUNIPER_CEVO: opb.Device_JUNIPER,
		tpb.Node_JUNIPER_VMX:  opb.Device_JUNIPER,
		tpb.Node_NOKIA_SRL:    opb.Device_NOKIA,
		tpb.Node_LEMMING:      opb.Device_OPENCONFIG,
	}

	deviceVendors = map[tpb.Vendor]opb.Device_Vendor{
		tpb.Vendor_ARISTA:     opb.Device_ARISTA,
		tpb.Vendor_CISCO:      opb.Device_CISCO,
		tpb.Vendor_KEYSIGHT:   opb.Device_IXIA,
		tpb.Vendor_JUNIPER:    opb.Device_JUNIPER,
		tpb.Vendor_NOKIA:      opb.Device_NOKIA,
		tpb.Vendor_OPENCONFIG: opb.Device_OPENCONFIG,
	}
)

func filterTopology(topo *tpb.Topology) *tpb.Topology {
	t := &tpb.Topology{}
	for _, node := range topo.GetNodes() {
		_, vendorOK := deviceVendors[node.GetVendor()]
		_, typeOK := deviceTypes[node.GetType()]
		// Only include nodes with known type/vendor.
		if vendorOK || typeOK {
			t.Nodes = append(t.Nodes, node)
		} else {
			log.Infof("No known device vendor for node %q (vendor %v, type %v), ignoring node", node.GetName(), node.GetVendor(), node.GetType())
		}
	}
	for _, link := range topo.GetLinks() {
		foundA := false
		foundZ := false
		for _, node := range t.GetNodes() {
			if link.GetANode() == node.GetName() {
				foundA = true
			}
			if link.GetZNode() == node.GetName() {
				foundZ = true
			}
		}
		// Only include links with nodes passing filter.
		if foundA && foundZ {
			t.Links = append(t.Links, link)
		}
	}
	return t
}

// Solve creates a new Reservation from a desired testbed and an available topology.
func Solve(tb *opb.Testbed, topo *tpb.Topology) (*binding.Reservation, error) {
	topo = filterTopology(topo)
	devs := append(append([]*opb.Device{}, tb.GetDuts()...), tb.GetAtes()...)
	if numDevs, numNodes := len(devs), len(topo.GetNodes()); numDevs > numNodes {
		return nil, fmt.Errorf("not enough nodes in KNE topology for specified testbed: "+
			" testbed has %d devices and topology only has %d nodes", numDevs, numNodes)
	}
	if numTBLinks, numTopoLinks := len(tb.GetLinks()), len(topo.GetLinks()); numTBLinks > numTopoLinks {
		return nil, fmt.Errorf("not enough links in KNE topology for specified testbed "+
			" testbed has %d links and topology only has %d links", numTBLinks, numTopoLinks)
	}
	s := &solver{
		testbed:    tb,
		topology:   topo,
		dev2Ports:  make(map[*opb.Device]map[string]*opb.Port),
		dev2Links:  make(map[*opb.Device]map[*opb.Device]*devPG),
		node2Intfs: make(map[*tpb.Node]map[string]*intf),
		node2Links: make(map[*tpb.Node]map[*tpb.Node]*nodePG),
	}
	id2Dev := make(map[string]*opb.Device)
	name2Node := make(map[string]*tpb.Node)

	// Cache various info in the solver about the testbed and topology.
	for _, dev := range devs {
		id2Dev[dev.GetId()] = dev
		ports := make(map[string]*opb.Port)
		for _, port := range dev.GetPorts() {
			ports[port.GetId()] = port
		}
		s.dev2Ports[dev] = ports
		s.dev2Links[dev] = make(map[*opb.Device]*devPG)
	}
	// Build the dev => links map for testbed
	// Currently we will support port groups for ATE devices only; so group name may be specified
	// only for one end of a link. The other end of these group links must map to a single device,
	// for all members of a group. Group specified on both ends of a link is unsupported for now.
	devicePGMap := make(map[string]*opb.Device)
	for _, link := range tb.GetLinks() {
		srcParts := strings.Split(link.GetA(), ":")
		dstParts := strings.Split(link.GetB(), ":")
		srcDev := id2Dev[srcParts[0]]
		dstDev := id2Dev[dstParts[0]]
		srcPort := s.dev2Ports[srcDev][srcParts[1]]
		dstPort := s.dev2Ports[dstDev][dstParts[1]]
		group := srcPort.GetGroup()
		peerDev := dstDev
		if group == "" {
			group = dstPort.GetGroup()
			peerDev = srcDev
		} else if dstPort.GetGroup() != "" {
			return nil, fmt.Errorf("unsupported configuration in testbed; group specified at both ends of link: %v", link)
		}
		// For valid port groups ensure DUT end maps to the same device for all members of the group
		if group != "" {
			if expPeerDev, ok := devicePGMap[group]; ok {
				if peerDev != expPeerDev {
					return nil, fmt.Errorf("inconsistent port group configuration found in testbed")
				}
			} else {
				devicePGMap[group] = peerDev
			}
		}
		dLink := &devLink{srcDev: srcDev, dstDev: dstDev, srcPort: srcPort, dstPort: dstPort, group: group}
		if _, ok := s.dev2Links[srcDev][dstDev]; !ok {
			pg := &devPG{groupMap: make(map[string][]*devLink)}
			s.dev2Links[srcDev][dstDev] = pg
			s.dev2Links[dstDev][srcDev] = pg
		}
		s.dev2Links[srcDev][dstDev].groupMap[group] = append(s.dev2Links[srcDev][dstDev].groupMap[group], dLink)
	}
	// For dev / node port group arrange the group names in descending order of membership size.
	// So for multiple port groups between two devices (and nodes), we assign interfaces for largest port group first
	// and go down that list; all un-assigned interfaces can be allocated to regular port.
	// pgSizes - map of port group name and membership size
	// returns list of sorted port group names in the order of membership size, and number of total links
	sortGroupBySize := func(pgSizes map[string]int) ([]string, int) {
		// Sort port group names based on number of members; skip for regular ports / intfs
		var sorted []string
		var numLinks int
		for grp, linkCount := range pgSizes {
			numLinks = numLinks + linkCount
			if grp == "" {
				continue
			}
			inserted := false
			for index, gName := range sorted {
				if pgSizes[gName] < linkCount {
					sorted = append(sorted[:index+1], sorted[index:]...)
					sorted[index] = grp
					inserted = true
					break
				}
			}
			if !inserted {
				sorted = append(sorted, grp)
			}
		}
		return sorted, numLinks
	}
	pDevs := make(map[*opb.Device]bool)
	for srcDev := range s.dev2Links {
		for dstDev, devPG := range s.dev2Links[srcDev] {
			if pDevs[dstDev] {
				pgSizes := make(map[string]int)
				for gName, devLinks := range devPG.groupMap {
					pgSizes[gName] = len(devLinks)
				}
				devPG.sortedNames, devPG.totalLinks = sortGroupBySize(pgSizes)
			}
		}
		pDevs[srcDev] = true
	}

	for _, node := range s.topology.GetNodes() {
		name2Node[node.GetName()] = node
		s.node2Intfs[node] = make(map[string]*intf)
		s.node2Links[node] = make(map[*tpb.Node]*nodePG)
	}
	addIntf := func(nodeName, intfName string) {
		node := name2Node[nodeName]
		i, ok := s.node2Intfs[node][intfName]
		if !ok {
			i = &intf{name: intfName}
			s.node2Intfs[node][intfName] = i
		}
		i.vendorName = node.Interfaces[intfName].GetName()
		if i.vendorName == "" {
			i.vendorName = intfName
		}
	}
	for _, link := range s.topology.GetLinks() {
		addIntf(link.GetANode(), link.GetAInt())
		addIntf(link.GetZNode(), link.GetZInt())
	}
	// Build the node => links map for topology
	pg2Node := make(map[string]*tpb.Node)
	for _, link := range topo.GetLinks() {
		srcNode := name2Node[link.GetANode()]
		dstNode := name2Node[link.GetZNode()]
		srcIntf := s.node2Intfs[srcNode][link.GetAInt()]
		dstIntf := s.node2Intfs[dstNode][link.GetZInt()]
		group := srcNode.Interfaces[link.GetAInt()].GetGroup()
		dstGroup := dstNode.Interfaces[link.GetZInt()].GetGroup()
		peerNode := dstNode
		if group == "" {
			group = dstGroup
			peerNode = srcNode
		} else if dstGroup != "" {
			// Group specified on both ends of a link is unsupported for now.
			return nil, fmt.Errorf("unsupported configuration; group specified at both ends of link: %v", link)
		}
		// For valid port groups ensure DUT end maps to the same node for all members of the group
		if group != "" {
			if expPeerNode, ok := pg2Node[group]; ok {
				if peerNode != expPeerNode {
					return nil, fmt.Errorf("inconsistent port group configuration found in topology")
				}
			} else {
				pg2Node[group] = peerNode
			}
		}
		nLink := &nodeLink{srcNode: srcNode, dstNode: dstNode, srcIntf: srcIntf, dstIntf: dstIntf, group: group}
		if _, ok := s.node2Links[srcNode][dstNode]; !ok {
			pg := &nodePG{groupMap: make(map[string][]*nodeLink)}
			s.node2Links[srcNode][dstNode] = pg
			s.node2Links[dstNode][srcNode] = pg
		}
		s.node2Links[srcNode][dstNode].groupMap[group] = append(s.node2Links[srcNode][dstNode].groupMap[group], nLink)
	}
	pNode := make(map[*tpb.Node]bool)
	for srcNode := range s.node2Links {
		for dstNode, nodePG := range s.node2Links[srcNode] {
			if pNode[dstNode] {
				pgMap := make(map[string]int)
				for gName, nodeLinks := range nodePG.groupMap {
					pgMap[gName] = len(nodeLinks)
				}
				nodePG.sortedNames, nodePG.totalLinks = sortGroupBySize(pgMap)
			}
		}
		pNode[srcNode] = true
	}

	a, err := s.solve()
	if err != nil {
		return nil, err
	}

	res := &binding.Reservation{
		ID:   uuid.New(),
		DUTs: make(map[string]binding.DUT),
		ATEs: make(map[string]binding.ATE),
	}

	for _, dut := range tb.GetDuts() {
		resDUT, err := a.resolveDUT(dut)
		if err != nil {
			return nil, err
		}
		res.DUTs[dut.GetId()] = resDUT
	}
	for _, ate := range tb.GetAtes() {
		resATE, err := a.resolveATE(ate)
		if err != nil {
			return nil, err
		}
		res.ATEs[ate.GetId()] = resATE
	}
	return res, nil
}

// ServiceDUT is a DUT that contains a service map.
type ServiceDUT struct {
	*binding.AbstractDUT
	Services map[string]*tpb.Service
	Cert     *tpb.CertificateCfg
}

// Service returns the KNE service details for a given service name.
func (d *ServiceDUT) Service(service string) (*tpb.Service, error) {
	srv, ok := d.Services[service]
	if !ok {
		return nil, fmt.Errorf("service %q not found on DUT %q", service, d.Name())
	}
	return srv, nil
}

// ServiceATE is an ATE that contains a service map.
type ServiceATE struct {
	*binding.AbstractATE
	Services map[string]*tpb.Service
	Cert     *tpb.CertificateCfg
}

// Service returns the KNE service details for a given service name.
func (a *ServiceATE) Service(service string) (*tpb.Service, error) {
	srv, ok := a.Services[service]
	if !ok {
		return nil, fmt.Errorf("service %q not found on ATE %q", service, a.Name())
	}
	return srv, nil
}

type assign struct {
	dev2Node  map[*opb.Device]*tpb.Node
	port2Intf map[*opb.Port]*intf
}

type intf struct {
	name       string
	vendorName string
}

type devPG struct {
	groupMap    map[string][]*devLink
	sortedNames []string
	totalLinks  int
}

type nodePG struct {
	groupMap    map[string][]*nodeLink
	sortedNames []string
	totalLinks  int
}

type devLink struct {
	srcDev, dstDev   *opb.Device
	srcPort, dstPort *opb.Port
	group            string
}

type nodeLink struct {
	srcNode, dstNode *tpb.Node
	srcIntf, dstIntf *intf
	group            string
}

func (a *assign) String() string {
	var sb strings.Builder
	for dut, node := range a.dev2Node {
		sb.WriteString(fmt.Sprintf("%s->%s\n", dut.GetId(), node.GetName()))
		for _, p := range dut.GetPorts() {
			sb.WriteString(fmt.Sprintf("  %s->%s\n", p.GetId(), a.port2Intf[p]))
		}
	}
	return sb.String()
}

func (a *assign) resolveDUT(dev *opb.Device) (*ServiceDUT, error) {
	dims, srvs, cert, err := a.resolveDevice(dev)
	if err != nil {
		return nil, err
	}
	return &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{dims},
		Services:    srvs,
		Cert:        cert,
	}, nil
}

func (a *assign) resolveATE(dev *opb.Device) (*ServiceATE, error) {
	dims, srvs, cert, err := a.resolveDevice(dev)
	if err != nil {
		return nil, err
	}
	return &ServiceATE{
		AbstractATE: &binding.AbstractATE{dims},
		Services:    srvs,
		Cert:        cert,
	}, nil
}

func (a *assign) resolveDevice(dev *opb.Device) (*binding.Dims, map[string]*tpb.Service, *tpb.CertificateCfg, error) {
	node, ok := a.dev2Node[dev]
	if !ok {
		return nil, nil, nil, fmt.Errorf("node %q not resolved", dev.GetId())
	}
	// TODO: Cleanup after Type deprecation
	vendor, ok := deviceVendors[node.GetVendor()]
	if !ok {
		vendor, ok = deviceTypes[node.GetType()]
		if !ok {
			return nil, nil, nil, fmt.Errorf("no known device vendor for node %q (vendor %v, type %v)", node.GetName(), node.GetVendor(), node.GetType())
		}
	}
	dims := &binding.Dims{
		Name:            node.GetName(),
		Vendor:          vendor,
		HardwareModel:   node.GetModel(),
		SoftwareVersion: node.GetOs(),
		Ports:           make(map[string]*binding.Port),
	}
	for _, p := range dev.GetPorts() {
		dims.Ports[p.GetId()] = &binding.Port{Name: a.port2Intf[p].vendorName}
	}
	sm := map[string]*tpb.Service{}
	for _, s := range node.GetServices() {
		sm[s.GetName()] = s
	}
	return dims, sm, node.GetConfig().GetCert(), nil
}

type solver struct {
	testbed    *opb.Testbed
	topology   *tpb.Topology
	dev2Ports  map[*opb.Device]map[string]*opb.Port
	dev2Links  map[*opb.Device]map[*opb.Device]*devPG
	node2Intfs map[*tpb.Node]map[string]*intf
	node2Links map[*tpb.Node]map[*tpb.Node]*nodePG
}

func (s *solver) solve() (*assign, error) {
	// Find all the matching device->node assignments, and
	// for each of those, all the port->intf assignments.
	dev2Nodes := make(map[*opb.Device][]*tpb.Node)
	for _, dut := range s.testbed.GetDuts() {
		nodeList, err := s.nodeMatches(dut, false)
		if err != nil {
			return nil, err
		}
		dev2Nodes[dut] = nodeList
	}
	for _, ate := range s.testbed.GetAtes() {
		nodeList, err := s.nodeMatches(ate, true)
		if err != nil {
			return nil, err
		}
		dev2Nodes[ate] = nodeList
	}

	// Generate possible testbed device -> topology node combinations.
	dev2NodeChan := s.genDevCombos(dev2Nodes)
	for dev2Node := range dev2NodeChan {
		link2link := s.buildLink2Links(dev2Node)
		link2LinkChan := s.genLinkCombos(dev2Node, link2link)
		for dLink2nLink := range link2LinkChan {
			if a := s.newAssign(dev2Node, dLink2nLink); a != nil {
				// TODO: When solver is rewritten, signal the gorouting
				// channel to exit early here and avoid leaving the goroutine hanging.
				// Not disastrous but ideally the goroutines would terminate here.
				return a, nil
			}
		}
	}
	return nil, fmt.Errorf("no KNE topology matches the testbed")
}

// buildLink2Links builds a map of dev to node links, for a device => node selection
// All port group links are matched (size in descending order) between device-peer pairs and node-peer pairs
// Device peer group links can be mapped to any corresponding node peer group links as long as the node group size
// is atleast as large as the device group size.
// d2n - is the device => node map
func (s *solver) buildLink2Links(d2n map[*opb.Device]*tpb.Node) map[*devLink][]*nodeLink {
	link2link := make(map[*devLink][]*nodeLink)
	assignedDevs := make(map[*opb.Device]bool)
	for dev, node := range d2n {
		for devPeer, devPG := range s.dev2Links[dev] {
			if _, ok := assignedDevs[devPeer]; ok {
				nodePeer := d2n[devPeer]
				nodePG := s.node2Links[node][nodePeer]
				allNodeLinks := []*nodeLink{}
				devGPIndex := 0
				for _, nodeGroupName := range nodePG.sortedNames {
					if len(devPG.sortedNames) > devGPIndex {
						devGroupName := devPG.sortedNames[devGPIndex]
						devPGSize := len(devPG.groupMap[devGroupName])
						nodePGSize := len(nodePG.groupMap[nodeGroupName])
						if devPGSize > nodePGSize {
							for _, link := range devPG.groupMap[devGroupName] {
								link2link[link] = allNodeLinks
							}
							devGPIndex++
						}
					}
					allNodeLinks = append(allNodeLinks, nodePG.groupMap[nodeGroupName]...)
				}
				// Finally add to link2link for all remaining devPG links
				for ; devGPIndex < len(devPG.sortedNames); devGPIndex++ {
					devGroupName := devPG.sortedNames[devGPIndex]
					for _, link := range devPG.groupMap[devGroupName] {
						link2link[link] = allNodeLinks
					}
				}

				if _, ok = devPG.groupMap[""]; ok {
					// Append all regular links
					allNodeLinks = append(allNodeLinks, nodePG.groupMap[""]...)
					for _, link := range devPG.groupMap[""] {
						link2link[link] = allNodeLinks
					}
				}
			}
		}
		assignedDevs[dev] = true
	}
	return link2link
}

func (s *solver) nodeMatches(dev *opb.Device, isATE bool) ([]*tpb.Node, error) {
	var nodeList []*tpb.Node
	for _, node := range s.topology.GetNodes() {
		// TODO: Cleanup after Type deprecation
		if ateVendors[node.GetVendor()] != isATE && ateTypes[node.GetType()] != isATE {
			continue
		}
		match := s.devMatch(dev, node)
		if match {
			log.V(1).Infof("Found match: testbed device %q -> KNE topology node %q", dev.GetId(), node.GetName())
			nodeList = append(nodeList, node)
		}
	}
	if len(nodeList) == 0 {
		return nil, fmt.Errorf("no node in KNE topology to match testbed device %q", dev.GetId())
	}
	return nodeList, nil
}

func (s *solver) devMatch(dev *opb.Device, node *tpb.Node) bool {
	if dev.GetHardwareModel() != "" && dev.GetHardwareModel() != hardwareModel(node) {
		return false
	}
	if dev.GetSoftwareVersion() != "" && dev.GetSoftwareVersion() != softwareVersion(node) {
		return false
	}
	vendor, ok := deviceVendors[node.GetVendor()]
	if !ok {
		vendor = deviceTypes[node.GetType()]
	}
	if v := dev.GetVendor(); v != opb.Device_VENDOR_UNSPECIFIED && v != vendor {
		return false
	}
	log.V(1).Infof("Found node match: %q", dev.GetId())
	intfs := s.node2Intfs[node]
	log.V(1).Infof("Interfaces: %v", intfs)
	// If the device needs more ports than the node, this node cannot match.
	return len(dev.GetPorts()) <= len(intfs)
}

func (s *solver) portMatch(port *opb.Port, intf *intf) bool {
	// TODO: When solver is rewritten, support more generic port matching.
	if port.GetSpeed() != opb.Port_SPEED_UNSPECIFIED {
		return false
	}
	return true
}

func hardwareModel(node *tpb.Node) string {
	return tpb.Node_Type_name[int32(node.GetType())]
}

func softwareVersion(node *tpb.Node) string {
	return hardwareModel(node)
}

// genLinkCombos yields port link->node link mappings
func (s *solver) genLinkCombos(d2n map[*opb.Device]*tpb.Node, m map[*devLink][]*nodeLink) <-chan map[*devLink]*nodeLink {
	keys := s.arrangeLinkKeys(m)
	ch := make(chan map[*devLink]*nodeLink)
	go func() {
		defer close(ch)
		s.genLinkRecurse(d2n, m, keys, make(map[*devLink]*nodeLink), make(map[string]string), make(map[*nodeLink]bool), ch)
	}()
	return ch
}

func (s *solver) genLinkRecurse(
	d2n map[*opb.Device]*tpb.Node,
	m map[*devLink][]*nodeLink,
	keys []*devLink,
	res map[*devLink]*nodeLink,
	grpRes map[string]string,
	used map[*nodeLink]bool,
	ch chan<- map[*devLink]*nodeLink) {
	if len(keys) == 0 {
		copy := make(map[*devLink]*nodeLink)
		for k, v := range res {
			copy[k] = v
		}
		ch <- copy
		return
	}
	first := keys[0]
	for _, i := range m[first] {
		if !used[i] {
			grpResUpdated := false
			// Ensure that for groups, all member links for a group get mapped to node links
			// that are co-located in a single group
			if first.group != "" {
				if gName, ok := grpRes[first.group]; ok {
					// Some previous member of device port group got mapped to a different node intf group
					if gName != i.group {
						continue
					}
				} else {
					// This is to keep track of which iteration is adding entry to grpRes map,
					// so as to delete it when res entry is deleted for that iteration
					grpResUpdated = true
				}
			}
			if !s.shouldLinkRecurse(d2n, first, i) {
				continue
			}
			res[first] = i
			grpRes[first.group] = i.group
			used[i] = true
			s.genLinkRecurse(d2n, m, keys[1:], res, grpRes, used, ch)
			delete(used, i)
			delete(res, first)
			if grpResUpdated {
				// Remove only if this link was first in the group to add the map entry
				delete(grpRes, first.group)
			}
		}
	}
}

// arrangeLinkKeys creates a list of links (between a device pair) from testbed.
// All regular links (with no group) are added at the end.
// m - is the map of all testbed links and corresponding potential topology links for mapping
// returns list of key links in the desirable processing order, with all regular links added at the end
func (s *solver) arrangeLinkKeys(m map[*devLink][]*nodeLink) []*devLink {
	// All regular link (port) allocation should happen last
	var allRlinks, orderedKeys []*devLink
	for k := range m {
		if k.group == "" {
			allRlinks = append(allRlinks, k)
		} else {
			orderedKeys = append(orderedKeys, k)
		}
	}
	return append(orderedKeys, allRlinks...)
}

// shouldLinkRecurse validates whether a topology link can be assigned for a testbed link by validating
// ports assigned match requirements. Other than port specific match, device peer group links have to "map"
// to node peer group links of the selected group.
// d2n - device to node mapped for this solution set.
// dLink - is the testbed link.
// nLink - is potential mapped node link.
// returns true if the ports match; false otherwise
func (s *solver) shouldLinkRecurse(d2n map[*opb.Device]*tpb.Node, dLink *devLink, nLink *nodeLink) bool {
	if d2n[dLink.srcDev] == nLink.srcNode {
		return s.portMatch(dLink.srcPort, nLink.srcIntf) && s.portMatch(dLink.dstPort, nLink.dstIntf)
	}
	return s.portMatch(dLink.srcPort, nLink.dstIntf) && s.portMatch(dLink.dstPort, nLink.srcIntf)
}

// genDutCombos yields device -> node mappings
func (s *solver) genDevCombos(m map[*opb.Device][]*tpb.Node) <-chan map[*opb.Device]*tpb.Node {
	var keys []*opb.Device
	for dev := range m {
		keys = append(keys, dev)
	}
	ch := make(chan map[*opb.Device]*tpb.Node)
	go func() {
		defer close(ch)
		s.genDevRecurse(m, keys, make(map[*opb.Device]*tpb.Node), make(map[*tpb.Node]bool), ch)
	}()
	return ch
}

func (s *solver) genDevRecurse(
	m map[*opb.Device][]*tpb.Node,
	keys []*opb.Device,
	res map[*opb.Device]*tpb.Node,
	used map[*tpb.Node]bool,
	ch chan<- map[*opb.Device]*tpb.Node) {
	if len(keys) == 0 {
		copy := make(map[*opb.Device]*tpb.Node)
		for k, v := range res {
			copy[k] = v
		}
		ch <- copy
		return
	}
	first := keys[0]
	for _, i := range m[first] {
		if !used[i] {
			if !s.shouldDevRecurse(first, i, res) {
				continue
			}
			res[first] = i
			used[i] = true
			s.genDevRecurse(m, keys[1:], res, used, ch)
			delete(used, i)
			delete(res, first)
		}
	}
}

// shouldDevRecurse validates whether a topology node can be assigned for a testbed device by validating
// required number of links are present between node peers.
// dev - is the testbed device.
// node - is potential mapped node.
// res - is working resultset based on past selections.
// returns true if the devices match; false otherwise
func (s *solver) shouldDevRecurse(dev *opb.Device, node *tpb.Node, res map[*opb.Device]*tpb.Node) bool {
	for devPeer, devPG := range s.dev2Links[dev] {
		if nodePeer, ok := res[devPeer]; ok {
			devRLinks := 0
			if _, ok = devPG.groupMap[""]; ok {
				devRLinks = len(devPG.groupMap[""])
			}
			if _, ok = s.node2Links[node][nodePeer]; !ok {
				// There is no link between current 'node' and previously selected peer node as required by testbed
				return false
			}
			nodePG := s.node2Links[node][nodePeer]
			nodeRLinks := nodePG.totalLinks
			if len(devPG.sortedNames) > len(nodePG.sortedNames) {
				// The required number of port groups between device and peer is more than those between node and peer
				return false
			}
			for index, dGroupName := range devPG.sortedNames {
				nGroupName := nodePG.sortedNames[index]
				dGroupSize := len(devPG.groupMap[dGroupName])
				nGroupSize := len(nodePG.groupMap[nGroupName])
				if dGroupSize > nGroupSize {
					// The port group size between device and peer is more than that between node and peer
					return false
				}
				// Balance links can be used as regular links
				nodeRLinks = nodeRLinks - dGroupSize
			}
			if devRLinks > nodeRLinks {
				// The required regular links between device and peer is more than what is available between node and peer
				return false
			}
		}
	}
	return true
}

func (s *solver) newAssign(dev2Node map[*opb.Device]*tpb.Node, link2link map[*devLink]*nodeLink) *assign {
	a := &assign{
		dev2Node:  make(map[*opb.Device]*tpb.Node),
		port2Intf: make(map[*opb.Port]*intf),
	}
	assignedIntfs := make(map[*intf]bool)
	// From the mapped link => link, build up the port => intf map
	for dLink, nLink := range link2link {
		if dev2Node[dLink.srcDev] == nLink.srcNode {
			a.port2Intf[dLink.srcPort] = nLink.srcIntf
			a.port2Intf[dLink.dstPort] = nLink.dstIntf
		} else {
			a.port2Intf[dLink.srcPort] = nLink.dstIntf
			a.port2Intf[dLink.dstPort] = nLink.srcIntf
		}
		assignedIntfs[nLink.srcIntf] = true
		assignedIntfs[nLink.dstIntf] = true
	}
	for d, n := range dev2Node {
		a.dev2Node[d] = n
		// Assign remaining ports not connected to any device
		for _, port := range s.dev2Ports[d] {
			if _, ok := a.port2Intf[port]; !ok {
				for _, intf := range s.node2Intfs[n] {
					if _, ok = assignedIntfs[intf]; !ok {
						a.port2Intf[port] = intf
						assignedIntfs[intf] = true
						break
					}
				}
			}
		}
	}
	return a
}
