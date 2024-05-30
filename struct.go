package xunsafe

import (
	"reflect"
)

type (

	//Struct represents a struct
	Struct struct {
		Fields []Field
	}

	//Matcher represents a field matcher
	Matcher struct {
		keyFn func(string) string
		index map[string]*Field
	}
)

// Matcher creates a filed matched for supplied key Fn
func (s *Struct) Matcher(keyFn func(string) string) *Matcher {
	var matcher = Matcher{
		index: make(map[string]*Field, len(s.Fields)),
	}
	for i := range s.Fields {
		field := &s.Fields[i]
		matcher.index[keyFn(field.Name)] = field
	}
	return &matcher
}

func (s *Struct) MatchByType(target reflect.Type) *Field {
	for i := range s.Fields {
		field := &s.Fields[i]
		fType := field.Type
		if fType.Kind() == reflect.Ptr {
			fType = fType.Elem()
		}
		if fType == reflect.TypeOf(target) {
			return field
		}
	}
	return nil
}

// NewStruct creates a unsafe struct wrapper
func NewStruct(sType reflect.Type) *Struct {
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}
	result := &Struct{
		Fields: make([]Field, sType.NumField()),
	}
	for i := 0; i < sType.NumField(); i++ {
		result.Fields[i] = *NewField(sType.Field(i))
	}
	return result
}

// Match matches field with name
func (s *Matcher) Match(name string) *Field {
	return s.index[s.keyFn(name)]
}
