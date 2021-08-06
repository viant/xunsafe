package xunsafe

import "unsafe"

//Getter represents a func returning field value pointer, it takes holder address
type Getter func(structAdd unsafe.Pointer) interface{}
