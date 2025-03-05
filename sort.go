package semver

import "sort"

type Versions []*Version

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Versions) Less(i, j int) bool {
	return v[i].Compare(v[j]) < 0
}

var _ sort.Interface = Versions{}
