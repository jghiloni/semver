package semver_test

import (
	"bufio"
	"os"
	"sort"
	"testing"

	"github.com/jghiloni/semver"
)

func TestVersionsSorting(t *testing.T) {
	toSort := loadVersions(t, "testdata/strict.txt")
	presorted := loadVersions(t, "testdata/sorted-strict.txt")

	sort.Sort(toSort)
	if len(toSort) != len(presorted) {
		t.Fatal("sorted and presorted are not the same length")
	}

	for i := range toSort {
		if toSort[i].Compare(presorted[i]) != 0 {
			t.Fatalf("%q and %q are different", toSort[i], presorted[i])
		}
	}
}

func loadVersions(t *testing.T, file string) semver.Versions {
	t.Helper()

	fp, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	defer fp.Close()

	versions := make(semver.Versions, 0, 31)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		vs := scanner.Text()
		v, err := semver.ParseStrict(vs)
		if err != nil {
			t.Fatalf("%q: %v", vs, err)
		}

		versions = append(versions, v)
	}

	return versions
}
