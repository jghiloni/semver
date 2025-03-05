package commands_test

import (
	"io"
	"strings"
	"testing"

	"github.com/jghiloni/go-commonutils/utils"
	"github.com/jghiloni/semver"
	"github.com/jghiloni/semver/cli/commands"
)

var testCLIVersion = utils.Must(semver.ParseStrict("0.0.0"))

func TestNext(t *testing.T) {
	params := [][]string{
		{"major", "11.0.0"},
		{"minor", "10.21.0"},
		{"patch", "10.20.31"},
		{"prerelease", "10.20.31-rc.1"},
	}

	helper := func(field string, expected string) {
		t.Helper()

		stdin := strings.NewReader(fakeStdIn)

		out := &strings.Builder{}

		cliArgs := []string{"next", field}
		err := commands.Execute(&commands.ExecuteArgs{
			CLIVersion: testCLIVersion,
			Stdout:     out,
			Stderr:     io.Discard,
			Stdin:      stdin,
			Args:       cliArgs,
		})

		if err != nil {
			t.Fatal(err)
		}

		output := strings.TrimSpace(out.String())
		if output != expected {
			t.Fatalf("%q should match %q", output, expected)
		}
	}

	for _, testCase := range params {
		helper(testCase[0], testCase[1])
	}
}
