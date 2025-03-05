package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/jghiloni/go-commonutils/utils"
	"github.com/jghiloni/semver"
)

type cmdLineArgs struct {
	Normalize             *NormalizeCommand `cmd:"" default:"1" help:"Parse the versions on stdin in a tolerant way and output them as compliant semver strings"`
	Next                  *NextCommand      `cmd:"" help:"Take the latest version on stdin and increase it based on which field is chosen"`
	Version               kong.VersionFlag  `short:"V" help:"Show the version of this CLI and exit"`
	IgnoreInvalidVersions bool              `short:"i" negatable:"" default:"true" help:"If set, discard any text on stdin that does not parse as a semver string"`
	Silent                bool              `short:"q" help:"If set, don't output error messages"`
}

type ExecuteArgs struct {
	CLIVersion *semver.Version
	Stdout     io.Writer
	Stderr     io.Writer
	Stdin      io.Reader
	Args       []string
}

func Execute(cfg *ExecuteArgs) error {
	vars := kong.Vars{
		"version": cfg.CLIVersion.String(),
	}

	var args cmdLineArgs
	parser := kong.Must(&args, vars)
	ctx, err := parser.Parse(cfg.Args)
	if err != nil {
		return err
	}

	ctx.Stdout = cfg.Stdout
	ctx.Stderr = cfg.Stderr

	if args.Silent {
		ctx.Stderr = io.Discard
	}

	versions := getVersions(ctx, args, cfg.Stdin)
	return ctx.Run(versions)
}

func getVersions(k *kong.Context, args cmdLineArgs, stdin io.Reader) semver.Versions {
	inBytes, err := io.ReadAll(stdin)
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
