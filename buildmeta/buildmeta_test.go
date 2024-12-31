package buildmeta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	testFn     func(t assert.TestingT, expected interface{}, actual interface{}, args ...interface{}) bool
	testValues struct {
		version       string
		release       string
		builder       string
		buildDateTime string
	}
)

// по условию получаем функцию для проверки значений
func getAssertFn(inputArg, wantValue string) (testFn, string) {
	if inputArg == wantValue {
		return assert.Equal, "values should be equal"
	}
	return assert.NotEqual, "values should not be equal"
}

func TestBuildMeta(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		description string
		inputValues testValues
		want        testValues
	}{
		{
			"OK",
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
		},
		{
			"invalid version",
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
			testValues{"v0.0.2", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
		},
		{
			"invalid release",
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
			testValues{"v0.0.1", "release_2006-01-01", "jdoe", "2006-01-02 15:04:05"},
		},
		{
			"invalid builder",
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
			testValues{"v0.0.1", "release_2006-01-02", "johndoe", "2006-01-02 15:04:05"},
		},
		{
			"invalid time",
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:05"},
			testValues{"v0.0.1", "release_2006-01-02", "jdoe", "2006-01-02 15:04:06"},
		},
	}

	for _, testCase := range testCases {
		desc := testCase.description
		want := testCase.want
		input := testCase.inputValues
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			var msg string
			var assertFn testFn
			meta := NewBuildMeta(input.version, input.release, input.builder, input.buildDateTime)

			assertFn, msg = getAssertFn(input.version, want.version)
			assertFn(t, want.version, meta.GetVersion(), msg)

			assertFn, msg = getAssertFn(input.release, want.release)
			assertFn(t, want.release, meta.GetRelease(), msg)

			assertFn, msg = getAssertFn(input.builder, want.builder)
			assertFn(t, want.builder, meta.GetBuilder(), msg)

			assertFn, msg = getAssertFn(input.buildDateTime, want.buildDateTime)
			assertFn(t, want.buildDateTime, meta.GetBuildTime(), msg)
		})
	}
}
