package semver

import (
	"regexp"
	"strings"
)

var (
	strictParseRE   = regexp.MustCompile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	tolerantParseRE = regexp.MustCompile(`^v?(?P<major>0|[1-9]\d*)(?:\.(?P<minor>0|[1-9]\d*)(?:\.(?P<patch>0|[1-9]\d*))?)?(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)

	strictGroupNames   = strictParseRE.SubexpNames()
	tolerantGroupNames = tolerantParseRE.SubexpNames()
)

// ParseStrict parses a string into a semantic version that strictly follows the
// standard, and validates it
func ParseStrict(versionString string) (*Version, error) {
	return commonParse(versionString, strictParseRE, strictGroupNames, "")
}

// ParseTolerant allows more leeway in parsing -- it allows an optional "v" at the
// beginning of the string, and it allows minor or patch to be missing, to be
// replaced with zeroes. HOWEVER, if patch is present, then minor must be present
// as well. It validates the version before returning
func ParseTolerant(versionString string) (*Version, error) {
	return commonParse(versionString, tolerantParseRE, tolerantGroupNames, "0")
}

func commonParse(versionString string, re *regexp.Regexp, groupNames []string, defaultFieldVal string) (*Version, error) {
	var (
		major string = defaultFieldVal
		minor string = defaultFieldVal
		patch string = defaultFieldVal

		pre, meta []string
	)

	versionString = strings.TrimSpace(versionString)
	if versionString == "" {
		return nil, nil
	}

	allMatches := re.FindStringSubmatch(versionString)
	if len(allMatches) == 0 {
		return nil, ErrUnparseable
	}

	for groupIndex, value := range allMatches {
		name := groupNames[groupIndex]
		switch name {
		case "major":
			if value != "" {
				major = value
			}
		case "minor":
			if value != "" {
				minor = value
			}
		case "patch":
			if value != "" {
				patch = value
			}
		case "prerelease":
			if value != "" {
				pre = strings.Split(value, ".")
			}
		case "buildmetadata":
			if value != "" {
				meta = strings.Split(value, ".")
			}
		}
	}

	v := &Version{
		major, minor, patch, pre, meta,
	}

	return v, nil
}
