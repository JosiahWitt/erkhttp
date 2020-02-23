// Package erkhttp allows embeding HTTP statuses in erk error kinds.
package erkhttp

import "github.com/JosiahWitt/erk"

// HTTPStatusable items can return an HTTP status.
type HTTPStatusable interface {
	HTTPStatus() int
}

// GetHTTPStatus from the kind of the provided error.
//
// If the status cannot be found, it defaults to 500 (Internal Server Error), and returns false.
func GetHTTPStatus(err error) (int, bool) {
	kind := erk.GetKind(err)
	if statusable, ok := kind.(HTTPStatusable); ok {
		return statusable.HTTPStatus(), true
	}

	return 500, false
}
