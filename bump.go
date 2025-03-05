package semver

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

var bi1 = big.NewInt(1)

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

func (v *Version) BumpPrerelease() error {
	if len(v.pre) == 0 {
		return errors.New("can't bump empty prerelease")
	}

	lastPre := v.pre[len(v.pre)-1]
	if isStringNumeric(lastPre) {
		bumped, err := v.bumpNumericString(lastPre)
		if err != nil {
			return fmt.Errorf("BumpPrerelease: %w", err)
		}
		v.pre[len(v.pre)-1] = bumped
		return nil
	}

	v.pre = append(v.pre, "1")
	v.meta = nil
	return nil
}

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
