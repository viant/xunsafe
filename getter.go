package xunsafe

//Getter represents a func returning field value pointer, it takes holder address
type Getter func(structPtr uintptr) interface{}
