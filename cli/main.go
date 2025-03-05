package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/jghiloni/go-commonutils/utils"
	"github.com/jghiloni/semver"
	"github.com/jghiloni/semver/cli/commands"
)

var (
	version  string = "0.0.0"
	buildRef string = "local"
)

type cmdLineArgs struct {
	Normalize             *commands.NormalizeCommand `cmd:"" default:"1" help:"Parse the versions on stdin in a tolerant way and output them as compliant semver strings"`
	Next                  *commands.NextCommand      `cmd:"" help:"Take the latest version on stdin and increase it based on which field is chosen"`
	Version               kong.VersionFlag           `short:"V" help:"Show the version of this CLI and exit"`
	IgnoreInvalidVersions bool                       `short:"i" negatable:"" default:"true" help:"If set, discard any text on stdin that does not parse as a semver string"`
	Silent                bool                       `short:"q" help:"If set, don't output error messages"`
}

func main() {
	v, err := semver.ParseTolerant(version)
	if err != nil {
		v, _ = semver.ParseStrict("0.0.0+invalid-version")
	} else {
		v.SetBuildMetadata(buildRef)
	}

	vars := kong.Vars{
		"version": v.String(),
	}

	var args cmdLineArgs
	ctx := kong.Parse(&args, vars)

	if args.Silent {
		ctx.Stderr = io.Discard
	}

	versions := getVersions(ctx, args)
	ctx.FatalIfErrorf(ctx.Run(versions))
}

func getVersions(k *kong.Context, args cmdLineArgs) semver.Versions {
	inBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(k.Stderr, err)
		os.Exit(1)
	}

	versionStrings := strings.Fields(string(inBytes))
	var versions semver.Versions = utils.Map(versionStrings, func(vs string) *semver.Version {
		v, err := semver.ParseTolerant(vs)
		if err != nil {
			fmt.Fprintln(k.Stderr, err)
			if !args.IgnoreInvalidVersions {
				os.Exit(1)
			}

			return nil
		}
		return v
	})

	versions = utils.Filter(versions, func(v *semver.Version) bool {
		return v != nil
	})

	return versions
}
