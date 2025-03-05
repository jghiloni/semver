package semver_test

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/jghiloni/semver"
)

func TestBumps(t *testing.T) {
	tests, err := os.Open("testdata/bumped.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer tests.Close()

	reader := csv.NewReader(tests)

	rows, err := reader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	// drop the header row
	rows = rows[1:]
	for _, row := range rows {
		test := row[0]
		major := row[1]
		minor := row[2]
		patch := row[3]
		pre := row[4]

		v, err := semver.ParseStrict(test)
		if err != nil {
			t.Fatalf("%q: %v", test, err)
		}

		maj := v.Clone()
		if err := maj.BumpMajor(); err != nil {
			t.Fatalf("%q: %v", maj, err)
		}

		if maj.String() != major {
			t.Fatalf("%q should equal %q", maj, major)
		}

		min := v.Clone()
		if err := min.BumpMinor(); err != nil {
			t.Fatalf("%q: %v", min, err)
		}

		if min.String() != minor {
			t.Fatalf("%q should equal %q", min, minor)
		}

		p := v.Clone()
		if err := p.BumpPatch(); err != nil {
			t.Fatalf("%q: %v", p, err)
		}

		if p.String() != patch {
			t.Fatalf("%q should equal %q", p, patch)
		}

		if pre != "" {
			preV := v.Clone()
			if err := preV.BumpPrerelease(); err != nil {
				t.Fatalf("%q should equal %q", preV, pre)
			}
		}
	}
}

func TestSets(t *testing.T) {
	v, err := semver.ParseStrict("0.0.0")
	if err != nil {
		t.Fatal(err)
	}

	v1 := v.Clone()
	v1.SetPrelease("foo.bar.baz")

	if v1.String() != "0.0.0-foo.bar.baz" {
		t.Fatalf("%q should equal 0.0.0-foo.bar.baz", v1)
	}

	if len(v.Prerelease()) != 0 {
		t.Fatalf("%s should have no prerelease data", v)
	}

	expectedPre := []string{"foo", "bar", "baz"}
	if len(v1.Prerelease()) != 3 {
		t.Fatal("prerelease should have 3 items")
	}

	for i := range 3 {
		if v1.Prerelease()[i] != expectedPre[i] {
			t.Fatalf("%q should equal %q", v1.Prerelease()[i], expectedPre[i])
		}
	}

	v2 := v.Clone()
	v2.SetBuildMetadata("build.0asdf")
	if v2.String() != "0.0.0+build.0asdf" {
		t.Fatalf("%q should equal 0.0.0+build.0asdf", v2)
	}

	expectedMeta := []string{"build", "0asdf"}
	if len(v2.BuildMetadata()) != 2 {
		t.Fatal("buildmetadata should have 2 items")
	}

	for i := range 2 {
		if v2.BuildMetadata()[i] != expectedMeta[i] {
			t.Fatalf("%q should equal %q", v2.BuildMetadata()[i], expectedMeta[i])
		}
	}
}

func TestBumpNegativeCases(t *testing.T) {
	v, err := semver.ParseStrict("0.0.0")
	if err != nil {
		t.Fatal(err)
	}

	err = v.Clone().BumpPrerelease()
	if err == nil {
		t.Fatal("BumpPrerelease on an empty prerelease should have errored")
	}

	err = v.Clone().SetPrelease("%")
	if err == nil {
		t.Fatal("SetPrerelease with an invalid character should have errored")
	}

	err = v.Clone().SetBuildMetadata(".")
	if err == nil {
		t.Fatal("SetBuildMetadata with an invalid character should have errored")
	}
}
