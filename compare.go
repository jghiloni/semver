package semver

import (
	"strings"
	"unicode"
)

// Compare compares two semantic versions based on the rules set out
// in the standard. If v is smaller than v2, Compare returns a value < 0.
// If v is larger than v2, Compare returns a value > 0. In both cases,
// the actual value returned is variable, but will conform to that signature.
// If the two versions are equivalent according to the standard, Compare returns 0.
func (v *Version) Compare(v2 *Version) int {
	majorDiff := compareNumericStrings(v.major, v2.major)
	if majorDiff != 0 {
		return majorDiff
	}

	minorDiff := compareNumericStrings(v.minor, v2.minor)
	if minorDiff != 0 {
		return minorDiff
	}

	patchDiff := compareNumericStrings(v.patch, v2.patch)
	if patchDiff != 0 {
		return patchDiff
	}

	vPre, v2Pre := v.Prerelease(), v2.Prerelease()
	// no prerelease data is always larger than some prerelease data if x.y.z are the same
	if len(vPre) == 0 && len(v2Pre) > 0 {
		return 1
	}

	if len(vPre) > 0 && len(v2Pre) == 0 {
		return -1
	}

	commonLen := min(len(vPre), len(v2Pre))
	for i := range commonLen {
		diff := comparePrereleaseStrings(vPre[i], v2Pre[i])
		if diff != 0 {
			return diff
		}
	}

	// if we're here, all the prerelease fields have been equal so far.
	// the version with more prerelease fields has higher precedence, or
	// is "larger". If they have the same length, that means they're equal
	return len(vPre) - len(v2Pre)
}

func compareNumericStrings(x, y string) int {
	if len(x) > len(y) {
		return 1
	}

	if len(x) < len(y) {
		return -1
	}

	// if we're here, x and y are the same length
	for i := range x {
		// time for shady byte math!
		bx := int(x[i] - '0')
		by := int(y[i] - '0')
		diff := bx - by
		if diff != 0 {
			return diff
		}
	}

	return 0
}

func comparePrereleaseStrings(x, y string) int {
	if strings.EqualFold(x, y) {
		return 0
	}

	if isStringNumeric(x) {
		if isStringNumeric(y) {
			return compareNumericStrings(x, y)
		}

		// numeric strings are always "less than" non-numeric
		return -1
	}

	if isStringNumeric(y) {
		// if we're here, x is non-numeric and y is numeric, so x is larger
		return 1
	}

	// they're both non-numeric strings
	return strings.Compare(strings.ToLower(x), strings.ToLower(y))
}

func isStringNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) > 1 && s[0] == '0' {
		return false
	}

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}
