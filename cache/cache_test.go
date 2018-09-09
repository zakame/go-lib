package cache

import (
	"reflect"
	"testing"
)

func TestCache(t *testing.T) {
	c := New(2)

	expectMiss := func(k string) {
		v, ok := c.Get(k)
		if ok {
			t.Fatalf("expected cache miss on key %q but hit value %v", k, v)
		}
	}

	expectHit := func(k string, ev interface{}) {
		v, ok := c.Get(k)
		if !ok {
			t.Fatalf("expected cache(%q)=%v; but missed", k, ev)
		}
		if !reflect.DeepEqual(v, ev) {
			t.Fatalf("expected cache(%q)=%v; but got %v", k, ev, v)
		}
	}

	expectMiss("1")
	c.Add("1", "one")
	expectHit("1", "one")

	c.Add("2", "two")
	expectHit("1", "one")
	expectHit("2", "two")

	c.Add("3", "three")
	expectHit("3", "three")
	expectHit("2", "two")
	expectMiss("1")
}
