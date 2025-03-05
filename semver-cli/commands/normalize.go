package commands

import (
	"fmt"
	"sort"

	"github.com/alecthomas/kong"
	"github.com/jghiloni/semver"
)

type NormalizeCommand struct {
	SortAscending  bool `aliases:"sort" xor:"sort" help:"If set, sort versions on stdin from smallest/oldest/lowest precedence to largest/newest/highest precedence"`
	SortDescending bool `xor:"sort" help:"If set, sort versions on stdin from largest/newest/highest precedence to smallest/oldest/lowest precedence"`
}

func (n *NormalizeCommand) Run(k *kong.Context, versions semver.Versions) error {
	if n.SortAscending {
		sort.Sort(versions)
	}

	if n.SortDescending {
		sort.Sort(sort.Reverse(versions))
	}

	for _, v := range versions {
		fmt.Fprintln(k.Stdout, v)
	}

	return nil
}
