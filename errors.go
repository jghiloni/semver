package semver

import "errors"

var (
	ErrMajorEmpty    = errors.New("major is empty")
	ErrMinorEmpty    = errors.New("minor is empty")
	ErrPatchEmpty    = errors.New("patch is empty")
	ErrNotNumeric    = errors.New("field must be numeric and not zero-padded")
	ErrInvalidString = errors.New("strings must only consist of the characters [0-9A-Za-z-]")
	ErrUnparseable   = errors.New("could not parse string as semantic version")
)
