package erkhttp_test

import (
	"errors"
	"testing"

	"github.com/JosiahWitt/erk"
	"github.com/JosiahWitt/erkhttp"
	"github.com/matryer/is"
)

func TestGetHTTPStatus(t *testing.T) {
	table := []struct {
		Name           string
		Err            error
		ExpectedStatus int
		ExpectedOK     bool
	}{
		{
			Name:           "with erk error with embedded status",
			Err:            erk.New(ErkItemNotFound{}, "item not found"),
			ExpectedStatus: 404,
			ExpectedOK:     true,
		},
		{
			Name:           "with erk error with no embedded status",
			Err:            erk.New(ErkUnknown{}, "I don't know what happened."),
			ExpectedStatus: 500,
			ExpectedOK:     false,
		},
		{
			Name:           "with non-erk error",
			Err:            errors.New("not an erk error"),
			ExpectedStatus: 500,
			ExpectedOK:     false,
		},
		{
			Name:           "with nil error",
			Err:            nil,
			ExpectedStatus: 500,
			ExpectedOK:     false,
		},
	}

	for _, entry := range table {
		entry := entry // Pin range variable

		t.Run(entry.Name, func(t *testing.T) {
			is := is.New(t)

			status, ok := erkhttp.GetHTTPStatus(entry.Err)
			is.Equal(status, entry.ExpectedStatus)
			is.Equal(ok, entry.ExpectedOK)
		})
	}
}

type ErkItemNotFound struct {
	erk.DefaultKind
	erkhttp.StatusNotFound
}

type ErkUnknown struct{ erk.DefaultKind }
