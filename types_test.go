package xunsafe

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLookupType(t *testing.T) {
	var testCases = []struct {
		description string
		name        string
		expect      reflect.Type
	}{
		{
			description: "std lib type",
			name:        "time/Time",
			expect:      reflect.TypeOf(time.Time{}),
		},
		{
			description: "package type",
			name:        "github.com/viant/xunsafe/Field",
			expect:      reflect.TypeOf(Field{}),
		},
		{
			description: "builtin type",
			name:        "string",
			expect:      reflect.TypeOf(""),
		},
		{
			description: "unknown type",
			name:        "github.com/viant/xunsafe/NoSuchType",
			expect:      nil,
		},
	}

	for _, testCase := range testCases {
		actual := LookupType(testCase.name)
		assert.Equal(t, testCase.expect, actual, testCase.description)
	}
}

func TestPackageTypes(t *testing.T) {
	types := PackageTypes("github.com/viant/xunsafe")
	assert.Contains(t, types, reflect.TypeOf(Field{}))
	assert.Contains(t, types, reflect.TypeOf(Type{}))
	assert.Empty(t, PackageTypes("github.com/viant/no-such-package"))
}

func TestPackageNames(t *testing.T) {
	LookupType("time/Time")
	names := PackageNames()
	assert.Contains(t, names, "time")
	assert.Contains(t, names, "github.com/viant/xunsafe")
}
