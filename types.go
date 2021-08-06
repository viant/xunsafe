package xunsafe

import (
	"reflect"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})
var timeTypePtr = reflect.TypeOf(&time.Time{})
