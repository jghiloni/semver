package semver_test

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/jghiloni/semver"
)

func TestParseStrict(t *testing.T) {
	testfile, err := os.Open("testdata/strict.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer testfile.Close()

	scanner := bufio.NewScanner(testfile)
	for scanner.Scan() {
		testVersion := scanner.Text()
		v, err := semver.ParseStrict(testVersion)
		if err != nil {
			t.Fatalf("%q: %v", testVersion, err)
		}

		if v.String() != testVersion {
			t.Fatalf("expected %q to equal %q", testVersion, v)
		}
	}
}

func TestParseTolerant(t *testing.T) {
	combined := &bytes.Buffer{}

	strictBytes, err := os.ReadFile("testdata/strict.txt")
	if err != nil {
		t.Fatal(err)
	}

	tolerantBytes, err := os.ReadFile("testdata/tolerant.txt")
	if err != nil {
		t.Fatal(err)
	}

	combined.Write(strictBytes)
	combined.WriteByte('\n')
	combined.Write(tolerantBytes)

	scanner := bufio.NewScanner(combined)
	for scanner.Scan() {
		testVersion := scanner.Text()
		_, err := semver.ParseTolerant(testVersion)
		if err != nil {
			t.Fatalf("%q: %v", testVersion, err)
		}
	}
}

func TestParseInvalid(t *testing.T) {
	testfile, err := os.Open("testdata/invalid.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer testfile.Close()

	scanner := bufio.NewScanner(testfile)
	for scanner.Scan() {
		testVersion := scanner.Text()
		_, err := semver.ParseTolerant(testVersion)
		if err == nil {
			t.Fatalf("%q: expected error missing", testVersion)
		}
	}
}
