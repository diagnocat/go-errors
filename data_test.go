package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapD(t *testing.T) {
	a := assert.New(t)
	commonErr := New("common error")
	someData := map[string]int{"id": 5}

	testCases := []struct {
		name string

		err error

		expData  interface{}
		expCause error
	}{
		{
			name: "has data",

			err: WrapD(commonErr, someData),

			expData:  someData,
			expCause: commonErr,
		},
		{
			name: "has no data",

			err: commonErr,

			expData:  nil,
			expCause: commonErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a.Equal(tc.expData, Data(tc.err))
			a.Equal(tc.expCause, Cause(tc.err))
		})
	}
}

func TestWrapMD(t *testing.T) {
	a := assert.New(t)
	commonErr := New("common error")
	someData := map[string]int{"id": 5}

	err := WrapMD(commonErr, "some message", someData)

	a.Equal(someData, Data(err))
	a.Equal(commonErr, Cause(err))
}
