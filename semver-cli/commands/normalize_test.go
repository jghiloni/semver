package commands_test

import (
	"io"
	"strings"
	"testing"

	"github.com/jghiloni/semver/semver-cli/commands"
)

func TestNormalize(t *testing.T) {
	testCases := [][]string{
		{normalized},
		{sortedAsc, "normalize", "--sort-ascending"},
		{sortedDesc, "normalize", "--sort-descending"},
	}

	helper := func(expectedOutput string, commandLineArgs []string) {
		t.Helper()

		stdin := strings.NewReader(fakeStdIn)
		stdout := &strings.Builder{}
		stderr := io.Discard

		err := commands.Execute(&commands.ExecuteArgs{
			CLIVersion: testCLIVersion,
			Stdout:     stdout,
			Stderr:     stderr,
			Stdin:      stdin,
			Args:       commandLineArgs,
		})

		if err != nil {
			t.Fatal(err)
		}

		actualOutput := strings.TrimSpace(stdout.String())
		expectedOutput = strings.TrimSpace(expectedOutput)
		if expectedOutput != actualOutput {
			t.Fatalf("%q is supposed to equal %q", actualOutput, expectedOutput)
		}
	}

	for _, testCase := range testCases {
		cliArgs := []string{}
		if len(testCase) > 1 {
			cliArgs = testCase[1:]
		}

		helper(testCase[0], cliArgs)
	}
}
