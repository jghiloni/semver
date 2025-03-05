package commands

import (
	"errors"
	"fmt"
	"sort"

	"github.com/alecthomas/kong"
	"github.com/jghiloni/semver"
)

type ReleaseCommand struct {
	FailOnError bool `short:"f" default:"true" help:"If the latest version is not a prerelease, fail"`
}

func (cmd *ReleaseCommand) Run(k *kong.Context, versions semver.Versions) error {
	if len(versions) == 0 {
		if cmd.FailOnError {
			return errEmptyInputStream
		}

		fmt.Fprintln(k.Stderr, errEmptyInputStream)
		return nil
	}

	sort.Sort(sort.Reverse(versions))

	latest := versions[0]
	isPrerelease := len(latest.Prerelease()) > 0
	if !isPrerelease {
		msg := fmt.Sprintf("%q is not a prerelease and cannot be released", latest)
		if cmd.FailOnError {
			return errors.New(msg)
		}

		fmt.Fprintln(k.Stderr, msg)
		return nil
	}

	latest.SetPrelease("")
	latest.SetBuildMetadata("")

	fmt.Fprintln(k.Stdout, latest)
	return nil
}
