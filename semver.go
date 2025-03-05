package semver

import (
	"fmt"
	"strings"
)

// Version represents a semantic version compliant with the v2.0.0 standard specified
// at https://semver.org
type Version struct {
	major string
	minor string
	patch string
	pre   []string
	meta  []string
}

func (v *Version) Clone() *Version {
	newpre := make([]string, len(v.pre))
	newmeta := make([]string, len(v.meta))

	copy(newpre, v.pre)
	copy(newmeta, v.meta)

	return &Version{
		major: v.major,
		minor: v.minor,
		patch: v.patch,
		pre:   newpre,
		meta:  newmeta,
	}
}

func (v *Version) Prerelease() []string {
	return v.pre
}

func (v *Version) BuildMetadata() []string {
	return v.meta
}

func (v *Version) String() string {
	if v == nil {
		return ""
	}

	w := &strings.Builder{}
	fmt.Fprintf(w, "%s.%s.%s", v.major, v.minor, v.patch)

	if len(v.pre) > 0 {
		fmt.Fprintf(w, "-%s", strings.Join(v.pre, "."))
	}

	if len(v.meta) > 0 {
		fmt.Fprintf(w, "+%s", strings.Join(v.meta, "."))
	}

	return w.String()
}
