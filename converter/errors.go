package converter

import (
	"fmt"
	"reflect"
)

func UnsupportedConversion(from, to reflect.Type) error {
	return fmt.Errorf("unsupported conversion from %v to %v", from.String(), to.String())
}
