package semver_test

import (
	"testing"

	"github.com/jghiloni/semver"
)

func TestMorePrereleaseComparisons(t *testing.T) {
	root, _ := semver.ParseStrict("0.0.0")

	v1 := root.Clone()
	v2 := root.Clone()

	v1.SetPrelease("13234")
	v2.SetPrelease("2")

	cmp := v1.Compare(v2)
	if cmp <= 0 {
		t.Fatalf("%q has higher precedence than %q", v1, v2)
	}

	v1.SetPrelease("rc")
	cmp = v1.Compare(v2)
	if cmp <= 0 {
		t.Fatalf("%q has higher precedence than %q", v1, v2)
	}

	var n *semver.Version = nil
	ns := n.String()
	if ns != "" {
		t.Fatal("a nil Version should stringify to an empty string")
	}
}
