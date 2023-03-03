//go:build !debug

package xunsafe

func (f *Field) MustBeAssignable(y interface{}) {}
