package xunsafe

import (
	"reflect"
	"time"
)

var (
	typeTime    = reflect.TypeOf(time.Time{})
	typeTimePtr = reflect.TypeOf(&time.Time{})
)
