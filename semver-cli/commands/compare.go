package commands

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/jghiloni/semver"
)

type CompareCommand struct {
	Verbose   bool   `short:"v" help:"If set, it will print the input version as well as the comparison result"`
	Candidate string `arg:"" help:"The version against which to compare each input version. If the input version is smaller, it will return -1. If it is larger, it will return 1. If they are equivalent, it will return 0"`
}

func (c *CompareCommand) Run(k *kong.Context, versions semver.Versions) error {
	if len(versions) == 0 {
		return errEmptyInputStream
	}

	candidate, err := semver.ParseTolerant(c.Candidate)
	if err != nil {
		return err
	}

	for _, v := range versions {
		if c.Verbose {
			_, _ = fmt.Fprintf(k.Stdout, "%s: ", v)
		}
		_, _ = fmt.Fprintln(k.Stdout, v.Compare(candidate))
	}

	return nil
}
