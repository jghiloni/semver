package commands_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/jghiloni/go-fp/fslices"
	"github.com/jghiloni/semver/semver-cli/commands"
)

func TestCompare(t *testing.T) {
	expectedOutput := []string{"-1", "1", "1", "-1", "0", "-1", "-1", "-1", "-1", "-1", "-1", "-1", "-1", "-1", "1", "1", "1", "1"}

	stdin := strings.NewReader(fakeStdIn)
	stdout := new(strings.Builder)

	args := []string{"compare", "v1.1.2"}

	err := commands.Execute(&commands.ExecuteArgs{
		CLIVersion: testCLIVersion,
		Stdout:     stdout,
		Stderr:     io.Discard,
		Stdin:      stdin,
		Args:       args,
	})
	if err != nil {
		t.Fatal(err)
	}

	actualOutput := strings.Fields(strings.TrimSpace(stdout.String()))
	if !reflect.DeepEqual(actualOutput, expectedOutput) {
		t.Fatalf("expected comparison results %v, got %v", expectedOutput, actualOutput)
	}
}

func TestCompareVerbose(t *testing.T) {
	expectedOutput := `0.0.4: -1
  1.2.3: 1
  10.20.30: 1
  1.1.2-prerelease+meta: -1
  1.1.2+meta: 0
  1.0.0-alpha: -1
  1.0.0-beta: -1
  1.0.0-alpha.beta: -1
  1.0.0-alpha.beta.1: -1
  1.0.0-alpha.1: -1
  1.0.0-alpha0.valid: -1
  1.0.0-alpha.0valid: -1
  1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay: -1
  1.0.0-rc.1+build.1: -1
  2.0.0-rc.1+build.123: 1
  1.2.3-beta: 1
  10.2.3-DEV-SNAPSHOT: 1
  1.2.3-SNAPSHOT-123: 1`

	stdin := strings.NewReader(fakeStdIn)
	stdout := new(strings.Builder)

	args := []string{"compare", "v1.1.2", "-v"}

	err := commands.Execute(&commands.ExecuteArgs{
		CLIVersion: testCLIVersion,
		Stdout:     stdout,
		Stderr:     io.Discard,
		Stdin:      stdin,
		Args:       args,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedLines := fslices.Map(strings.Split(expectedOutput, "\n"), strings.TrimSpace)
	actualLines := fslices.Map(strings.Split(strings.TrimSpace(stdout.String()), "\n"), strings.TrimSpace)

	if len(expectedLines) != len(actualLines) {
		t.Fatalf("expected %d, got %d", len(expectedLines), len(actualLines))
	}

	for i := range expectedLines {
		if expectedLines[i] != actualLines[i] {
			t.Errorf("expected %s, got %s", expectedLines[i], actualLines[i])
		}
	}
}
