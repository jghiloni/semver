package commands_test

import (
	"io"
	"strings"
	"testing"

	"github.com/jghiloni/semver/semver-cli/commands"
)

type expectedReleaseResult struct {
	expectedError  bool
	expectedResult string
}

func TestRelease(t *testing.T) {
	expectedResults := []expectedReleaseResult{
		{true, ""},
		{false, "10.2.3"},
		{false, "2.0.0"},
		{true, ""},
		{false, "1.2.3"},
		{false, "1.2.3"},
		{true, ""},
		{false, "1.1.2"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{false, "1.0.0"},
		{true, ""},
	}

	actualVersions := strings.Fields(sortedDesc)

	for i := range actualVersions {
		stdin := strings.NewReader(actualVersions[i])
		stdout := &strings.Builder{}
		stderr := io.Discard
		args := []string{"release"}

		err := commands.Execute(&commands.ExecuteArgs{
			CLIVersion: testCLIVersion,
			Stdout:     stdout,
			Stderr:     stderr,
			Stdin:      stdin,
			Args:       args,
		})

		hasErr := err != nil
		if !hasErr && expectedResults[i].expectedError {
			t.Fatal("expected error did not occur")
		}

		if hasErr && !expectedResults[i].expectedError {
			t.Fatal("unexpected error occurred")
		}

		if hasErr && expectedResults[i].expectedError {
			continue
		}

		actual := strings.TrimSpace(stdout.String())
		expected := expectedResults[i].expectedResult

		if actual != expected {
			t.Fatalf("%q should equal %q", actual, expected)
		}
	}
}
