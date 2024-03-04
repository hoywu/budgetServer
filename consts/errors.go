package consts

import "errors"

var (
	ErrNotAllowed = errors.New("Rejected by policy")
	ErrPassword   = errors.New("Wrong password")
)
