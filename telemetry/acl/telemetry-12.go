package acl

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-acl/acl/state/counter-capability with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_CounterCapabilityPath) Lookup(t testing.TB) *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	goStruct := &oc.Acl{}
	md, ok := oc.Lookup(t, n, "Acl", goStruct, true, false)
	if ok {
		return convertAcl_CounterCapabilityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/state/counter-capability with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_CounterCapabilityPath) Get(t testing.TB) oc.E_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/state/counter-capability with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_CounterCapabilityPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertAcl_CounterCapabilityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/state/counter-capability with a ONCE subscription.
func (n *Acl_CounterCapabilityPathAny) Get(t testing.TB) []oc.E_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Acl_ACL_COUNTER_CAPABILITY
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_CounterCapabilityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Acl_CounterCapabilityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool) *oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher {
	t.Helper()
	w := &oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher{}
	gs := &oc.Acl{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl", gs, queryPath, true, false)
		return convertAcl_CounterCapabilityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_CounterCapabilityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool) *oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher {
	t.Helper()
	return watch_Acl_CounterCapabilityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_CounterCapabilityPath) Await(t testing.TB, timeout time.Duration, val oc.E_Acl_ACL_COUNTER_CAPABILITY) *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/state/counter-capability failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/state/counter-capability to the batch object.
func (n *Acl_CounterCapabilityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_CounterCapabilityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	c := &oc.CollectionE_Acl_ACL_COUNTER_CAPABILITY{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/state/counter-capability with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_CounterCapabilityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY) bool) *oc.E_Acl_ACL_COUNTER_CAPABILITYWatcher {
	t.Helper()
	return watch_Acl_CounterCapabilityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/state/counter-capability to the batch object.
func (n *Acl_CounterCapabilityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertAcl_CounterCapabilityPath extracts the value of the leaf CounterCapability from its parent oc.Acl
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY.
func convertAcl_CounterCapabilityPath(t testing.TB, md *genutil.Metadata, parent *oc.Acl) *oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY {
	t.Helper()
	qv := &oc.QualifiedE_Acl_ACL_COUNTER_CAPABILITY{
		Metadata: md,
	}
	val := parent.CounterCapability
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_InterfacePath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface {
	t.Helper()
	goStruct := &oc.Acl_Interface{}
	md, ok := oc.Lookup(t, n, "Acl_Interface", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_InterfacePath) Get(t testing.TB) *oc.Acl_Interface {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface with a ONCE subscription.
func (n *Acl_InterfacePathAny) Get(t testing.TB) []*oc.Acl_Interface {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface {
	t.Helper()
	c := &oc.CollectionAcl_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface)))
		return false
	})
	return c
}

func watch_Acl_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	w := &oc.Acl_InterfaceWatcher{}
	gs := &oc.Acl_Interface{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_Interface{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	return watch_Acl_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_InterfacePath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface) *oc.QualifiedAcl_Interface {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface to the batch object.
func (n *Acl_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface {
	t.Helper()
	c := &oc.CollectionAcl_Interface{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface) bool) *oc.Acl_InterfaceWatcher {
	t.Helper()
	return watch_Acl_InterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface to the batch object.
func (n *Acl_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSetPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSetPath) Get(t testing.TB) *oc.Acl_Interface_EgressAclSet {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSetPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_EgressAclSet
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a ONCE subscription.
func (n *Acl_Interface_EgressAclSetPathAny) Get(t testing.TB) []*oc.Acl_Interface_EgressAclSet {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_EgressAclSet
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSetPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface_EgressAclSet)))
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSetPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	w := &oc.Acl_Interface_EgressAclSetWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_Interface_EgressAclSet{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_EgressAclSet)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSetPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSetPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSetPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface_EgressAclSet) *oc.QualifiedAcl_Interface_EgressAclSet {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface_EgressAclSet) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set to the batch object.
func (n *Acl_Interface_EgressAclSetPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSetPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSetPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet) bool) *oc.Acl_Interface_EgressAclSetWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSetPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set to the batch object.
func (n *Acl_Interface_EgressAclSetPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Lookup(t testing.TB) *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	md, ok := oc.Lookup(t, n, "Acl_Interface_EgressAclSet_AclEntry", goStruct, false, false)
	if ok {
		return (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Get(t testing.TB) *oc.Acl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Lookup(t testing.TB) []*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Acl_Interface_EgressAclSet_AclEntry{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a ONCE subscription.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Get(t testing.TB) []*oc.Acl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Acl_Interface_EgressAclSet_AclEntry
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Acl_Interface_EgressAclSet_AclEntry)))
		return false
	})
	return c
}

func watch_Acl_Interface_EgressAclSet_AclEntryPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	w := &oc.Acl_Interface_EgressAclSet_AclEntryWatcher{}
	gs := &oc.Acl_Interface_EgressAclSet_AclEntry{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Acl_Interface_EgressAclSet_AclEntry", gs, queryPath, false, false)
		return (&oc.QualifiedAcl_Interface_EgressAclSet_AclEntry{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedAcl_Interface_EgressAclSet_AclEntry)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntryPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Await(t testing.TB, timeout time.Duration, val *oc.Acl_Interface_EgressAclSet_AclEntry) *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntryPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionAcl_Interface_EgressAclSet_AclEntry {
	t.Helper()
	c := &oc.CollectionAcl_Interface_EgressAclSet_AclEntry{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedAcl_Interface_EgressAclSet_AclEntry) bool) *oc.Acl_Interface_EgressAclSet_AclEntryWatcher {
	t.Helper()
	return watch_Acl_Interface_EgressAclSet_AclEntryPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-acl/acl/interfaces/interface/egress-acl-sets/egress-acl-set/acl-entries/acl-entry to the batch object.
func (n *Acl_Interface_EgressAclSet_AclEntryPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}