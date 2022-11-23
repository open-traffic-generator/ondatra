package otg

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ygot/ygot"
)

// QualifiedBgpPeer is a *BgpPeer with a corresponding timestamp.
type QualifiedBgpPeer struct {
	*genutil.Metadata
	val     *BgpPeer // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer sample, erroring out if not present.
func (q *QualifiedBgpPeer) Val(t testing.TB) *BgpPeer {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer sample.
func (q *QualifiedBgpPeer) SetVal(v *BgpPeer) *QualifiedBgpPeer {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer is a telemetry Collection whose Await method returns a slice of *BgpPeer samples.
type CollectionBgpPeer struct {
	W    *BgpPeerWatcher
	Data []*QualifiedBgpPeer
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer) Await(t testing.TB) []*QualifiedBgpPeer {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeerWatcher observes a stream of *BgpPeer samples.
type BgpPeerWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeerWatcher) Await(t testing.TB) (*QualifiedBgpPeer, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_Counters is a *BgpPeer_Counters with a corresponding timestamp.
type QualifiedBgpPeer_Counters struct {
	*genutil.Metadata
	val     *BgpPeer_Counters // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_Counters sample, erroring out if not present.
func (q *QualifiedBgpPeer_Counters) Val(t testing.TB) *BgpPeer_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_Counters sample.
func (q *QualifiedBgpPeer_Counters) SetVal(v *BgpPeer_Counters) *QualifiedBgpPeer_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_Counters is a telemetry Collection whose Await method returns a slice of *BgpPeer_Counters samples.
type CollectionBgpPeer_Counters struct {
	W    *BgpPeer_CountersWatcher
	Data []*QualifiedBgpPeer_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_Counters) Await(t testing.TB) []*QualifiedBgpPeer_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_CountersWatcher observes a stream of *BgpPeer_Counters samples.
type BgpPeer_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_CountersWatcher) Await(t testing.TB) (*QualifiedBgpPeer_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_UnicastIpv4Prefix is a *BgpPeer_UnicastIpv4Prefix with a corresponding timestamp.
type QualifiedBgpPeer_UnicastIpv4Prefix struct {
	*genutil.Metadata
	val     *BgpPeer_UnicastIpv4Prefix // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_UnicastIpv4Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_UnicastIpv4Prefix sample, erroring out if not present.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix) Val(t testing.TB) *BgpPeer_UnicastIpv4Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_UnicastIpv4Prefix sample.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix) SetVal(v *BgpPeer_UnicastIpv4Prefix) *QualifiedBgpPeer_UnicastIpv4Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_UnicastIpv4Prefix is a telemetry Collection whose Await method returns a slice of *BgpPeer_UnicastIpv4Prefix samples.
type CollectionBgpPeer_UnicastIpv4Prefix struct {
	W    *BgpPeer_UnicastIpv4PrefixWatcher
	Data []*QualifiedBgpPeer_UnicastIpv4Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_UnicastIpv4Prefix) Await(t testing.TB) []*QualifiedBgpPeer_UnicastIpv4Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_UnicastIpv4PrefixWatcher observes a stream of *BgpPeer_UnicastIpv4Prefix samples.
type BgpPeer_UnicastIpv4PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_UnicastIpv4Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_UnicastIpv4PrefixWatcher) Await(t testing.TB) (*QualifiedBgpPeer_UnicastIpv4Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_UnicastIpv4Prefix_AsPath is a *BgpPeer_UnicastIpv4Prefix_AsPath with a corresponding timestamp.
type QualifiedBgpPeer_UnicastIpv4Prefix_AsPath struct {
	*genutil.Metadata
	val     *BgpPeer_UnicastIpv4Prefix_AsPath // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_UnicastIpv4Prefix_AsPath sample, erroring out if not present.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) Val(t testing.TB) *BgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_UnicastIpv4Prefix_AsPath sample.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) SetVal(v *BgpPeer_UnicastIpv4Prefix_AsPath) *QualifiedBgpPeer_UnicastIpv4Prefix_AsPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix_AsPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_UnicastIpv4Prefix_AsPath is a telemetry Collection whose Await method returns a slice of *BgpPeer_UnicastIpv4Prefix_AsPath samples.
type CollectionBgpPeer_UnicastIpv4Prefix_AsPath struct {
	W    *BgpPeer_UnicastIpv4Prefix_AsPathWatcher
	Data []*QualifiedBgpPeer_UnicastIpv4Prefix_AsPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_UnicastIpv4Prefix_AsPath) Await(t testing.TB) []*QualifiedBgpPeer_UnicastIpv4Prefix_AsPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_UnicastIpv4Prefix_AsPathWatcher observes a stream of *BgpPeer_UnicastIpv4Prefix_AsPath samples.
type BgpPeer_UnicastIpv4Prefix_AsPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_UnicastIpv4Prefix_AsPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_UnicastIpv4Prefix_AsPathWatcher) Await(t testing.TB) (*QualifiedBgpPeer_UnicastIpv4Prefix_AsPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_UnicastIpv4Prefix_Community is a *BgpPeer_UnicastIpv4Prefix_Community with a corresponding timestamp.
type QualifiedBgpPeer_UnicastIpv4Prefix_Community struct {
	*genutil.Metadata
	val     *BgpPeer_UnicastIpv4Prefix_Community // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_UnicastIpv4Prefix_Community) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_UnicastIpv4Prefix_Community sample, erroring out if not present.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix_Community) Val(t testing.TB) *BgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_UnicastIpv4Prefix_Community sample.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix_Community) SetVal(v *BgpPeer_UnicastIpv4Prefix_Community) *QualifiedBgpPeer_UnicastIpv4Prefix_Community {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_UnicastIpv4Prefix_Community) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_UnicastIpv4Prefix_Community is a telemetry Collection whose Await method returns a slice of *BgpPeer_UnicastIpv4Prefix_Community samples.
type CollectionBgpPeer_UnicastIpv4Prefix_Community struct {
	W    *BgpPeer_UnicastIpv4Prefix_CommunityWatcher
	Data []*QualifiedBgpPeer_UnicastIpv4Prefix_Community
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_UnicastIpv4Prefix_Community) Await(t testing.TB) []*QualifiedBgpPeer_UnicastIpv4Prefix_Community {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_UnicastIpv4Prefix_CommunityWatcher observes a stream of *BgpPeer_UnicastIpv4Prefix_Community samples.
type BgpPeer_UnicastIpv4Prefix_CommunityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_UnicastIpv4Prefix_Community
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_UnicastIpv4Prefix_CommunityWatcher) Await(t testing.TB) (*QualifiedBgpPeer_UnicastIpv4Prefix_Community, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_UnicastIpv6Prefix is a *BgpPeer_UnicastIpv6Prefix with a corresponding timestamp.
type QualifiedBgpPeer_UnicastIpv6Prefix struct {
	*genutil.Metadata
	val     *BgpPeer_UnicastIpv6Prefix // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_UnicastIpv6Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_UnicastIpv6Prefix sample, erroring out if not present.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix) Val(t testing.TB) *BgpPeer_UnicastIpv6Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_UnicastIpv6Prefix sample.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix) SetVal(v *BgpPeer_UnicastIpv6Prefix) *QualifiedBgpPeer_UnicastIpv6Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_UnicastIpv6Prefix is a telemetry Collection whose Await method returns a slice of *BgpPeer_UnicastIpv6Prefix samples.
type CollectionBgpPeer_UnicastIpv6Prefix struct {
	W    *BgpPeer_UnicastIpv6PrefixWatcher
	Data []*QualifiedBgpPeer_UnicastIpv6Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_UnicastIpv6Prefix) Await(t testing.TB) []*QualifiedBgpPeer_UnicastIpv6Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_UnicastIpv6PrefixWatcher observes a stream of *BgpPeer_UnicastIpv6Prefix samples.
type BgpPeer_UnicastIpv6PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_UnicastIpv6Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_UnicastIpv6PrefixWatcher) Await(t testing.TB) (*QualifiedBgpPeer_UnicastIpv6Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_UnicastIpv6Prefix_AsPath is a *BgpPeer_UnicastIpv6Prefix_AsPath with a corresponding timestamp.
type QualifiedBgpPeer_UnicastIpv6Prefix_AsPath struct {
	*genutil.Metadata
	val     *BgpPeer_UnicastIpv6Prefix_AsPath // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_UnicastIpv6Prefix_AsPath) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_UnicastIpv6Prefix_AsPath sample, erroring out if not present.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix_AsPath) Val(t testing.TB) *BgpPeer_UnicastIpv6Prefix_AsPath {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_UnicastIpv6Prefix_AsPath sample.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix_AsPath) SetVal(v *BgpPeer_UnicastIpv6Prefix_AsPath) *QualifiedBgpPeer_UnicastIpv6Prefix_AsPath {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix_AsPath) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_UnicastIpv6Prefix_AsPath is a telemetry Collection whose Await method returns a slice of *BgpPeer_UnicastIpv6Prefix_AsPath samples.
type CollectionBgpPeer_UnicastIpv6Prefix_AsPath struct {
	W    *BgpPeer_UnicastIpv6Prefix_AsPathWatcher
	Data []*QualifiedBgpPeer_UnicastIpv6Prefix_AsPath
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_UnicastIpv6Prefix_AsPath) Await(t testing.TB) []*QualifiedBgpPeer_UnicastIpv6Prefix_AsPath {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_UnicastIpv6Prefix_AsPathWatcher observes a stream of *BgpPeer_UnicastIpv6Prefix_AsPath samples.
type BgpPeer_UnicastIpv6Prefix_AsPathWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_UnicastIpv6Prefix_AsPath
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_UnicastIpv6Prefix_AsPathWatcher) Await(t testing.TB) (*QualifiedBgpPeer_UnicastIpv6Prefix_AsPath, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedBgpPeer_UnicastIpv6Prefix_Community is a *BgpPeer_UnicastIpv6Prefix_Community with a corresponding timestamp.
type QualifiedBgpPeer_UnicastIpv6Prefix_Community struct {
	*genutil.Metadata
	val     *BgpPeer_UnicastIpv6Prefix_Community // val is the sample value.
	present bool
}

func (q *QualifiedBgpPeer_UnicastIpv6Prefix_Community) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *BgpPeer_UnicastIpv6Prefix_Community sample, erroring out if not present.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix_Community) Val(t testing.TB) *BgpPeer_UnicastIpv6Prefix_Community {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *BgpPeer_UnicastIpv6Prefix_Community sample.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix_Community) SetVal(v *BgpPeer_UnicastIpv6Prefix_Community) *QualifiedBgpPeer_UnicastIpv6Prefix_Community {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedBgpPeer_UnicastIpv6Prefix_Community) IsPresent() bool {
	return q != nil && q.present
}

// CollectionBgpPeer_UnicastIpv6Prefix_Community is a telemetry Collection whose Await method returns a slice of *BgpPeer_UnicastIpv6Prefix_Community samples.
type CollectionBgpPeer_UnicastIpv6Prefix_Community struct {
	W    *BgpPeer_UnicastIpv6Prefix_CommunityWatcher
	Data []*QualifiedBgpPeer_UnicastIpv6Prefix_Community
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionBgpPeer_UnicastIpv6Prefix_Community) Await(t testing.TB) []*QualifiedBgpPeer_UnicastIpv6Prefix_Community {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// BgpPeer_UnicastIpv6Prefix_CommunityWatcher observes a stream of *BgpPeer_UnicastIpv6Prefix_Community samples.
type BgpPeer_UnicastIpv6Prefix_CommunityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedBgpPeer_UnicastIpv6Prefix_Community
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *BgpPeer_UnicastIpv6Prefix_CommunityWatcher) Await(t testing.TB) (*QualifiedBgpPeer_UnicastIpv6Prefix_Community, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedDevice is a *Device with a corresponding timestamp.
type QualifiedDevice struct {
	*genutil.Metadata
	val     *Device // val is the sample value.
	present bool
}

func (q *QualifiedDevice) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Device sample, erroring out if not present.
func (q *QualifiedDevice) Val(t testing.TB) *Device {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Device sample.
func (q *QualifiedDevice) SetVal(v *Device) *QualifiedDevice {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedDevice) IsPresent() bool {
	return q != nil && q.present
}

// CollectionDevice is a telemetry Collection whose Await method returns a slice of *Device samples.
type CollectionDevice struct {
	W    *DeviceWatcher
	Data []*QualifiedDevice
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionDevice) Await(t testing.TB) []*QualifiedDevice {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// DeviceWatcher observes a stream of *Device samples.
type DeviceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedDevice
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *DeviceWatcher) Await(t testing.TB) (*QualifiedDevice, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow is a *Flow with a corresponding timestamp.
type QualifiedFlow struct {
	*genutil.Metadata
	val     *Flow // val is the sample value.
	present bool
}

func (q *QualifiedFlow) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow sample, erroring out if not present.
func (q *QualifiedFlow) Val(t testing.TB) *Flow {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow sample.
func (q *QualifiedFlow) SetVal(v *Flow) *QualifiedFlow {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow is a telemetry Collection whose Await method returns a slice of *Flow samples.
type CollectionFlow struct {
	W    *FlowWatcher
	Data []*QualifiedFlow
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow) Await(t testing.TB) []*QualifiedFlow {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// FlowWatcher observes a stream of *Flow samples.
type FlowWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *FlowWatcher) Await(t testing.TB) (*QualifiedFlow, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedFlow_Counters is a *Flow_Counters with a corresponding timestamp.
type QualifiedFlow_Counters struct {
	*genutil.Metadata
	val     *Flow_Counters // val is the sample value.
	present bool
}

func (q *QualifiedFlow_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Flow_Counters sample, erroring out if not present.
func (q *QualifiedFlow_Counters) Val(t testing.TB) *Flow_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Flow_Counters sample.
func (q *QualifiedFlow_Counters) SetVal(v *Flow_Counters) *QualifiedFlow_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedFlow_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionFlow_Counters is a telemetry Collection whose Await method returns a slice of *Flow_Counters samples.
type CollectionFlow_Counters struct {
	W    *Flow_CountersWatcher
	Data []*QualifiedFlow_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionFlow_Counters) Await(t testing.TB) []*QualifiedFlow_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Flow_CountersWatcher observes a stream of *Flow_Counters samples.
type Flow_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedFlow_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Flow_CountersWatcher) Await(t testing.TB) (*QualifiedFlow_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface is a *Interface with a corresponding timestamp.
type QualifiedInterface struct {
	*genutil.Metadata
	val     *Interface // val is the sample value.
	present bool
}

func (q *QualifiedInterface) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface sample, erroring out if not present.
func (q *QualifiedInterface) Val(t testing.TB) *Interface {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface sample.
func (q *QualifiedInterface) SetVal(v *Interface) *QualifiedInterface {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface is a telemetry Collection whose Await method returns a slice of *Interface samples.
type CollectionInterface struct {
	W    *InterfaceWatcher
	Data []*QualifiedInterface
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface) Await(t testing.TB) []*QualifiedInterface {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// InterfaceWatcher observes a stream of *Interface samples.
type InterfaceWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *InterfaceWatcher) Await(t testing.TB) (*QualifiedInterface, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ipv4Neighbor is a *Interface_Ipv4Neighbor with a corresponding timestamp.
type QualifiedInterface_Ipv4Neighbor struct {
	*genutil.Metadata
	val     *Interface_Ipv4Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ipv4Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ipv4Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_Ipv4Neighbor) Val(t testing.TB) *Interface_Ipv4Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ipv4Neighbor sample.
func (q *QualifiedInterface_Ipv4Neighbor) SetVal(v *Interface_Ipv4Neighbor) *QualifiedInterface_Ipv4Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ipv4Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ipv4Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_Ipv4Neighbor samples.
type CollectionInterface_Ipv4Neighbor struct {
	W    *Interface_Ipv4NeighborWatcher
	Data []*QualifiedInterface_Ipv4Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ipv4Neighbor) Await(t testing.TB) []*QualifiedInterface_Ipv4Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ipv4NeighborWatcher observes a stream of *Interface_Ipv4Neighbor samples.
type Interface_Ipv4NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ipv4Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ipv4NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_Ipv4Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedInterface_Ipv6Neighbor is a *Interface_Ipv6Neighbor with a corresponding timestamp.
type QualifiedInterface_Ipv6Neighbor struct {
	*genutil.Metadata
	val     *Interface_Ipv6Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedInterface_Ipv6Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *Interface_Ipv6Neighbor sample, erroring out if not present.
func (q *QualifiedInterface_Ipv6Neighbor) Val(t testing.TB) *Interface_Ipv6Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *Interface_Ipv6Neighbor sample.
func (q *QualifiedInterface_Ipv6Neighbor) SetVal(v *Interface_Ipv6Neighbor) *QualifiedInterface_Ipv6Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedInterface_Ipv6Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionInterface_Ipv6Neighbor is a telemetry Collection whose Await method returns a slice of *Interface_Ipv6Neighbor samples.
type CollectionInterface_Ipv6Neighbor struct {
	W    *Interface_Ipv6NeighborWatcher
	Data []*QualifiedInterface_Ipv6Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionInterface_Ipv6Neighbor) Await(t testing.TB) []*QualifiedInterface_Ipv6Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// Interface_Ipv6NeighborWatcher observes a stream of *Interface_Ipv6Neighbor samples.
type Interface_Ipv6NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedInterface_Ipv6Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *Interface_Ipv6NeighborWatcher) Await(t testing.TB) (*QualifiedInterface_Ipv6Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter is a *IsisRouter with a corresponding timestamp.
type QualifiedIsisRouter struct {
	*genutil.Metadata
	val     *IsisRouter // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter sample, erroring out if not present.
func (q *QualifiedIsisRouter) Val(t testing.TB) *IsisRouter {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter sample.
func (q *QualifiedIsisRouter) SetVal(v *IsisRouter) *QualifiedIsisRouter {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter is a telemetry Collection whose Await method returns a slice of *IsisRouter samples.
type CollectionIsisRouter struct {
	W    *IsisRouterWatcher
	Data []*QualifiedIsisRouter
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter) Await(t testing.TB) []*QualifiedIsisRouter {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouterWatcher observes a stream of *IsisRouter samples.
type IsisRouterWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouterWatcher) Await(t testing.TB) (*QualifiedIsisRouter, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_Counters is a *IsisRouter_Counters with a corresponding timestamp.
type QualifiedIsisRouter_Counters struct {
	*genutil.Metadata
	val     *IsisRouter_Counters // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_Counters) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_Counters sample, erroring out if not present.
func (q *QualifiedIsisRouter_Counters) Val(t testing.TB) *IsisRouter_Counters {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_Counters sample.
func (q *QualifiedIsisRouter_Counters) SetVal(v *IsisRouter_Counters) *QualifiedIsisRouter_Counters {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_Counters) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_Counters is a telemetry Collection whose Await method returns a slice of *IsisRouter_Counters samples.
type CollectionIsisRouter_Counters struct {
	W    *IsisRouter_CountersWatcher
	Data []*QualifiedIsisRouter_Counters
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_Counters) Await(t testing.TB) []*QualifiedIsisRouter_Counters {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_CountersWatcher observes a stream of *IsisRouter_Counters samples.
type IsisRouter_CountersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_Counters
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_CountersWatcher) Await(t testing.TB) (*QualifiedIsisRouter_Counters, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_Counters_Level1 is a *IsisRouter_Counters_Level1 with a corresponding timestamp.
type QualifiedIsisRouter_Counters_Level1 struct {
	*genutil.Metadata
	val     *IsisRouter_Counters_Level1 // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_Counters_Level1) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_Counters_Level1 sample, erroring out if not present.
func (q *QualifiedIsisRouter_Counters_Level1) Val(t testing.TB) *IsisRouter_Counters_Level1 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_Counters_Level1 sample.
func (q *QualifiedIsisRouter_Counters_Level1) SetVal(v *IsisRouter_Counters_Level1) *QualifiedIsisRouter_Counters_Level1 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_Counters_Level1) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_Counters_Level1 is a telemetry Collection whose Await method returns a slice of *IsisRouter_Counters_Level1 samples.
type CollectionIsisRouter_Counters_Level1 struct {
	W    *IsisRouter_Counters_Level1Watcher
	Data []*QualifiedIsisRouter_Counters_Level1
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_Counters_Level1) Await(t testing.TB) []*QualifiedIsisRouter_Counters_Level1 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_Counters_Level1Watcher observes a stream of *IsisRouter_Counters_Level1 samples.
type IsisRouter_Counters_Level1Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_Counters_Level1
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_Counters_Level1Watcher) Await(t testing.TB) (*QualifiedIsisRouter_Counters_Level1, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_Counters_Level2 is a *IsisRouter_Counters_Level2 with a corresponding timestamp.
type QualifiedIsisRouter_Counters_Level2 struct {
	*genutil.Metadata
	val     *IsisRouter_Counters_Level2 // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_Counters_Level2) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_Counters_Level2 sample, erroring out if not present.
func (q *QualifiedIsisRouter_Counters_Level2) Val(t testing.TB) *IsisRouter_Counters_Level2 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_Counters_Level2 sample.
func (q *QualifiedIsisRouter_Counters_Level2) SetVal(v *IsisRouter_Counters_Level2) *QualifiedIsisRouter_Counters_Level2 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_Counters_Level2) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_Counters_Level2 is a telemetry Collection whose Await method returns a slice of *IsisRouter_Counters_Level2 samples.
type CollectionIsisRouter_Counters_Level2 struct {
	W    *IsisRouter_Counters_Level2Watcher
	Data []*QualifiedIsisRouter_Counters_Level2
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_Counters_Level2) Await(t testing.TB) []*QualifiedIsisRouter_Counters_Level2 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_Counters_Level2Watcher observes a stream of *IsisRouter_Counters_Level2 samples.
type IsisRouter_Counters_Level2Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_Counters_Level2
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_Counters_Level2Watcher) Await(t testing.TB) (*QualifiedIsisRouter_Counters_Level2, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase is a *IsisRouter_LinkStateDatabase with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase) Val(t testing.TB) *IsisRouter_LinkStateDatabase {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase sample.
func (q *QualifiedIsisRouter_LinkStateDatabase) SetVal(v *IsisRouter_LinkStateDatabase) *QualifiedIsisRouter_LinkStateDatabase {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase samples.
type CollectionIsisRouter_LinkStateDatabase struct {
	W    *IsisRouter_LinkStateDatabaseWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabaseWatcher observes a stream of *IsisRouter_LinkStateDatabase samples.
type IsisRouter_LinkStateDatabaseWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabaseWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps is a *IsisRouter_LinkStateDatabase_Lsps with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps) SetVal(v *IsisRouter_LinkStateDatabase_Lsps) *QualifiedIsisRouter_LinkStateDatabase_Lsps {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps struct {
	W    *IsisRouter_LinkStateDatabase_LspsWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_LspsWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps samples.
type IsisRouter_LinkStateDatabase_LspsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_LspsWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs samples.
type IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_TlvsWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4ReachabilityWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_PrefixWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributesWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIpv4Reachability_Prefix_PrefixAttributes, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachabilityWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_NeighborWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_ExtendedIsReachability_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames is a *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames with a corresponding timestamp.
type QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames struct {
	*genutil.Metadata
	val     *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames // val is the sample value.
	present bool
}

func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames sample, erroring out if not present.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) Val(t testing.TB) *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames sample.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) SetVal(v *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) IsPresent() bool {
	return q != nil && q.present
}

// CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames is a telemetry Collection whose Await method returns a slice of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames samples.
type CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames struct {
	W    *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher
	Data []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames) Await(t testing.TB) []*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher observes a stream of *IsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames samples.
type IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *IsisRouter_LinkStateDatabase_Lsps_Tlvs_HostnamesWatcher) Await(t testing.TB) (*QualifiedIsisRouter_LinkStateDatabase_Lsps_Tlvs_Hostnames, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
