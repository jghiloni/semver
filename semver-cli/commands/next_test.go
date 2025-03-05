package commands_test

import (
	"io"
	"strings"
	"testing"

	"github.com/jghiloni/go-commonutils/utils"
	"github.com/jghiloni/semver"
	"github.com/jghiloni/semver/semver-cli/commands"
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

func TestNextWithExistingPrerelease(t *testing.T) {
	toTest := []string{
		"v1.0.0-rc",
		"2.1.0-rc.2",
		"1.6.4-beta.alpha.3",
		"5.0.0-rc.alpha0+meta",
		"1",
	}

	expectedResults := []string{
		"1.0.0-rc.1",
		"2.1.0-rc.3",
		"1.6.4-beta.alpha.4",
		"5.0.0-rc.alpha0.1",
		"1.0.1-rc.1",
	}

	helper := func(underTest string, expected string) {
		t.Helper()

		stdin := strings.NewReader(underTest)
		stdout := &strings.Builder{}
		args := []string{"next", "prerelease"}

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

		actual := strings.TrimSpace(stdout.String())
		if actual != expected {
			t.Fatalf("%q should equal %q", actual, expected)
		}
	}

	for i := range toTest {
		helper(toTest[i], expectedResults[i])
	}
}
