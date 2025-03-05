package semver

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

var bi1 = big.NewInt(1)

// BumpMajor attempts to increment the major field of the version by 1. If it's
// successful, it resets the minor and patch fields to 0 and unsets the prerelease
// and build metadata fields. It only errors if major is not parsable by
// (*math/big.Int).SetString, which should not occur.
func (v *Version) BumpMajor() error {
	var err error
	v.major, err = v.bumpNumericString(v.major)
	if err != nil {
		return fmt.Errorf("BumpMajor: %w", err)
	}

	v.minor = "0"
	v.patch = "0"
	v.pre = nil
	v.meta = nil

	return nil
}

// BumpMinor attempts to increment the minor field of the version by 1. If it's
// successful, it resets the patch field to 0 and unsets the prerelease
// and build metadata fields. It only errors if minor is not parsable by
// (*math/big.Int).SetString, which should not occur.
func (v *Version) BumpMinor() error {
	var err error
	v.minor, err = v.bumpNumericString(v.minor)
	if err != nil {
		return fmt.Errorf("BumpMinor: %w", err)
	}

	v.patch = "0"
	v.pre = nil
	v.meta = nil

	return nil
}

// BumpPatch attempts to increment the patch field of the version by 1. If it's
// successful, it unsets the prerelease and build metadata fields. It only
// errors if major is not parsable by (*math/big.Int).SetString, which should
// not occur.
func (v *Version) BumpPatch() error {
	var err error
	v.patch, err = v.bumpNumericString(v.patch)
	if err != nil {
		return fmt.Errorf("BumpPatch: %w", err)
	}

	v.pre = nil
	v.meta = nil

	return nil
}

// BumpPrerelease inspects the version's prerelease information, and if the last
// field is a number, increment it. If it is not, append ".1" to the end of the
// prerelease information. It returns an error if prerelease is empty.
func (v *Version) BumpPrerelease() error {
	if len(v.pre) == 0 {
		return errors.New("can't bump empty prerelease")
	}

	lastPre := v.pre[len(v.pre)-1]
	doAppend := true
	if isStringNumeric(lastPre) {
		bumped, err := v.bumpNumericString(lastPre)
		if err == nil {
			doAppend = false
			v.pre[len(v.pre)-1] = bumped
		}
	}

	if doAppend {
		v.pre = append(v.pre, "1")
	}
	v.meta = nil
	return nil
}

// SetPrerelease sets the prerelease information for the version. If the passed
// string is empty or contains only whitespace, it unsets the prerelease info.
// If it has data, it must be a dot-separated string, where each substring
// contains only letters, numbers, and dashes. If a substring is fully numeric,
// it must not be zero padded (if it starts with 0, it must be exactly 0).
func (v *Version) SetPrelease(pre string) error {
	pre = strings.TrimSpace(pre)
	if pre == "" {
		v.pre = nil
		return nil
	}

	v.pre = strings.Split(pre, ".")

	n, err := ParseStrict(v.String())
	if err != nil {
		return err
	}

	//lint:ignore SA4006 updates receiver
	v = n.Clone()

	return nil
}

// SetBuildMetadata sets the build metadata information for the version. If the passed
// string is empty or contains only whitespace, it unsets the prerelease info.
// If it has data, it must be a dot-separated string, where each substring
// contains only letters, numbers, and dashes.
func (v *Version) SetBuildMetadata(meta string) error {
	meta = strings.TrimSpace(meta)
	if meta == "" {
		v.meta = nil
		return nil
	}

	v.meta = strings.Split(meta, ".")

	n, err := ParseStrict(v.String())
	if err != nil {
		return err
	}

	//lint:ignore SA4006 updates receiver
	v = n.Clone()

	return nil
}

func (v *Version) bumpNumericString(s string) (string, error) {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return "", ErrNotNumeric
	}

	n = n.Add(n, bi1)

	return n.String(), nil
}
