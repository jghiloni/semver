package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jghiloni/semver"
	"github.com/jghiloni/semver/semver-cli/commands"
)

var (
	version  string = "0.0.0"
	buildRef string = "local"
)

func main() {
	log.SetFlags(0)

	cliVersion, err := semver.ParseTolerant(fmt.Sprintf("%s+%s", version, buildRef))
	if err != nil {
		log.Fatal(err)
	}

	if err = commands.Execute(&commands.ExecuteArgs{
		CLIVersion: cliVersion,
		Stdout:     os.Stdout,
		Stderr:     os.Stderr,
		Stdin:      os.Stdin,
		Args:       os.Args[1:],
	}); err != nil {
		log.Fatal(err)
	}
}
