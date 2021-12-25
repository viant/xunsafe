package xunsafe

import (
	"reflect"
	"strings"
)

type (

	//Struct represents a struct
	Struct struct {
		Field []Field
	}

	//Matcher represents a field matcher
	Matcher struct {
		fuzzy bool
		index map[string]*Field
	}
)

func (s *Struct) Matcher(fuzzy bool) *Matcher {
	var matcher = Matcher{
		fuzzy: fuzzy,
		index: make(map[string]*Field, len(s.Field)),
	}
	for i := range s.Field {
		field := &s.Field[i]
		matcher.index[matcher.key(field.Name)] = field
	}
	return &matcher
}

//NewStruct creates a unsafe struct wrapper
func NewStruct(sType reflect.Type) *Struct {
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}
	result := &Struct{
		Field: make([]Field, sType.NumField()),
	}
	for i := 0; i < sType.NumField(); i++ {
		result.Field[i] = *NewField(sType.Field(i))
	}
	return result
}

//Match matches field with name
func (s *Matcher) Match(name string) *Field {
	return s.index[s.key(name)]
}

func (s *Matcher) key(name string) string {
	if !s.fuzzy {
		return name
	}
	return strings.Replace(strings.ToLower(name), "_", "", strings.Count(name, "_"))
}
