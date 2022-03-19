package internal

import "errors"

var UNKNOWN_PACKAGE = errors.New("unknown package")
var RATE_LIMIT = errors.New("rate limit exceeded")
