package commands

import (
	"errors"
	"fmt"
	"sort"

	"github.com/alecthomas/kong"
	"github.com/jghiloni/semver"
)

type NextCommand struct {
	Field string `arg:"" enum:"major,minor,patch,prerelease" default:"patch" help:"Take the newest version from stdin and bump it according to the arg. If 'prerelease', it bumps the patch version and appends rc.1"`
}

func (n *NextCommand) Run(k *kong.Context, versions semver.Versions) error {
	if len(versions) == 0 {
		return errors.New("need at least one version")
	}

	sort.Sort(sort.Reverse(versions))

	latest := versions[0]
	switch n.Field {
	case "major":
		if err := latest.BumpMajor(); err != nil {
			return err
		}
	case "minor":
		if err := latest.BumpMinor(); err != nil {
			return err
		}
	case "patch":
		if err := latest.BumpPatch(); err != nil {
			return err
		}
	case "prerelease":
		if err := latest.BumpPatch(); err != nil {
			return err
		}

		latest.SetPrelease("rc.1")
	default:
		return fmt.Errorf("unrecognized field %s", n.Field)
	}

	fmt.Fprintln(k.Stdout, latest)
	return nil
}
