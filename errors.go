package semver

import "errors"

var (
	// ErrNotNumeric is returned when a given field is expected to be a number.
	// Numbers cannot be zero-padded in semver.
	ErrNotNumeric = errors.New("field must be numeric and not zero-padded")

	// ErrUnparseable is returned when parsing a semver that is not valid
	ErrUnparseable = errors.New("could not parse string as semantic version")
)
