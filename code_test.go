package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWrapC(t *testing.T) {
	a := assert.New(t)
	commonErr := New("common error")

	testCases := []struct {
		name string

		err error

		expCode  int
		expCause error
	}{
		{
			name: "has code",

			err: WrapC(commonErr, http.StatusBadRequest),

			expCode:  http.StatusBadRequest,
			expCause: commonErr,
		},
		{
			name: "has no code",

			err: commonErr,

			expCode:  0,
			expCause: commonErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a.Equal(tc.expCode, Code(tc.err))
			a.Equal(tc.expCause, Cause(tc.err))
		})
	}
}

func TestWrapMC(t *testing.T) {
	a := assert.New(t)
	commonErr := New("common error")

	err := WrapMC(commonErr, "some message", http.StatusBadRequest)

	a.Equal(http.StatusBadRequest, Code(err))
	a.Equal(commonErr, Cause(err))
}
