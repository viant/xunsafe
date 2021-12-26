# xunsafe (faster golang reflection)

[![GoReportCard](https://goreportcard.com/badge/github.com/viant/xunsafe)](https://goreportcard.com/report/github.com/viant/xunsafe)
[![GoDoc](https://godoc.org/github.com/viant/xunsafe?status.svg)](https://godoc.org/github.com/viant/xunsafe)

This library is compatible with Go 1.17+

Please refer to [`CHANGELOG.md`](CHANGELOG.md) if you encounter breaking changes.

- [Motivation](#motivation)
- [Introduction](#introduction)
- [Usage](#usage)
- [Bugs](#bugs)
- [Benchmark](#benchmark)
- [Contribution](#contributing-to-xunsafe)
- [License](#license)

## Motivation

In order to solve a problem that can work with any golang type, one has no choice but to use reflection.
Native golang reflection comes with hefty performance price, on benchmarking simple getter/setter case 
to manipulate struct dynamically I've seen around 100 time worst performance comparing to 
statically typed code. 
I believe that, native reflect package could be easily implemented way better to provide optimized performance.
This library comes with reflection implementation that greatly improved performance, that is  between 25 to 50x time faster than native golang reflection. 
What that means that extra overhead of using reflection is only around 1.5 to four times comparing to statically typed code.

## Introduction

In order to achieve better performance, this library uses unsafe.Pointer along with StructField.Offset to effectively access/modify struct fields.
On top of that most of implemented methods, inline giving substantial performance boost which is x40 times as opposed to the same not inlined version.

## Usage

######  Accessor/Mutator

```go
func Example_FastReflection() {
    
    type Foo struct {
        ID int
        Name string
    }
    fooType := reflect.TypeOf(Foo{})

    fooID := xunsafe.FieldByName(fooType, "ID")
    fooName := xunsafe.FieldByName(fooType, "Name")

    var foos = make([]Foo, 100)
    for i := range foos {
        fooAddr := unsafe.Pointer(&foos[i])
        fooID.SetInt(fooAddr, i)
        fooName.SetString(fooAddr, fmt.Sprintf("name %d", i))
    }

    for i := range foos {
        fooAddr := unsafe.Pointer(&foos[i])
        fmt.Printf("[%v] ID: %v, Name: %v\n", i, fooID.Int(fooAddr), fooName.String(fooAddr))
    }
}
```


###### Field Address

Field Addr returns an interface{} wrapping actual field pointer

```go
func ExampleAddr() {
	type Foo struct {
		ID int
		Name string
	}
	fooType := reflect.TypeOf(Foo{})
	fooID := xunsafe.FieldByName(fooType, "ID")
	foo := &Foo{ID: 101, Name: "name 101"}

	fooPtr := unsafe.Pointer(foo)
	*(fooID.Addr(fooPtr).(*int)) = 201
	fmt.Printf("foo.ID: %v\n", foo.ID)//prints 201
}
```

###### Field Value

Field Interface returns an interface{} wrapping actual field value

```go
func ExampleAddr() {
	type Foo struct {
		ID int
		Name string
	}
	fooType := reflect.TypeOf(Foo{})
	fooID := xunsafe.FieldByName(fooType, "ID")
	foo := &Foo{ID: 101, Name: "name 101"}

    fooAddr := unsafe.Pointer(foo)
	fmt.Printf("foo.ID: %v\n", fooID.Interface(fooAddr))//prints 101
}
```

For base golang type Field Addr and Value got optimized with casting unsafe address to actual corresponding type. 
For example for filed with int type, the casting come in form ```(*int)(unsafe.Pointer(structAddr + field.Offset))```


### Slice range

```go
        type T {ID int}
        aSlice := NewSlice(reflect.TypeOf([]T))
		var items = []T{
			{ID:1}, {ID:2}, {ID:3},
        }   
		aSlice.Range(unsafe.Pointer(&items), func(index int, item interface{}) bool {
			actual := item.(*T)
			fmt.Printf("%+v\n", item)
			return true //to continue
		})
```


### Slice index

```go
        type T {ID int}
        aSlice := NewSlice(reflect.TypeOf([]T))
		var items = []T{
			{ID:1}, {ID:2}, {ID:3},
        }   
		slicePtr := unsafe.Pointer(&items)
		for i :=0;i<len(items);i++ {
            item := aSlice.Index(slicePtr, i).(T)
            fmt.Printf("%+v\n", item)
        }
	
```

### Slice appender

```go
        type T {ID int}
        aSlice := NewSlice(reflect.TypeOf([]T))
		var items []T
		
		appender := aSlice.Appender(unsafe.Pointer(&items))
        appender.Append(T{ID:1},{ID:2})
        appender.Append(T{ID:3},T{ID:4})
        fmt.Printf("%v\n", items)
		
```


### Arbitrary type Ref/Deref/Pointer

Defined Type implements Pointer and Deref arbitrary type.

```go
        type T 
    	aType := xunsafe.NewType(reflect.TypeOf(T))
		var t T
		ref := aType.Ref(t) //return *T
		deref := aType.Deref(ref) //return T
		ptr := aType.Pointer(t) //return unsafe.Pointer
```

## Bugs

This package operates on unsafe.Pointer and also uses some redefined private reflect package types like rtype, emptyInterface.
While directed type should work well, some generic method like Field.Interface, Field.Set, may not support all data types.
User of the package should ensure the code is fully tested and run test with -race and  -gcflags=all=-d=checkptr flags


### Benchmark

Accessor/Mutator benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/viant/xunsafe
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkField_Accessor_Native-16                       886540330                1.209 ns/op           0 B/op          0 allocs/op
BenchmarkField_Accessor_Direct_Xunsafe-16               606187651                1.967 ns/op           0 B/op          0 allocs/op
BenchmarkField_Accessor_Interface_Xunsafe-16            256453082                4.520 ns/op           0 B/op          0 allocs/op
BenchmarkField_Accessor_Interface_Reflect-16            10056830               118.1 ns/op            56 B/op          4 allocs/op
BenchmarkField_Accessor_Addr_Xunsafe-16                 168350235                7.225 ns/op           0 B/op          0 allocs/op
BenchmarkField_Accessor_Addr_Reflect-16                 20753077                49.50 ns/op            0 B/op          0 allocs/op
BenchmarkField_Mutator_Native-16                        1000000000               0.9133 ns/op          0 B/op          0 allocs/op
Benchmark_Mutator_Direct_Xunsafe-16                     773207817                1.513 ns/op           0 B/op          0 allocs/op
Benchmark_Mutator_Set_Xunsafe-16                        458994487                2.750 ns/op           0 B/op          0 allocs/op
Benchmark_Mutator_Xunsafe_Ptr-16                        183253660                6.550 ns/op           0 B/op          0 allocs/op
BenchmarkField_Mutator_Reflect-16                       10741909                94.74 ns/op           32 B/op          3 allocs/op
BenchmarkField_Mutator_Addr_Reflect-16                  31762166                36.48 ns/op            0 B/op          0 allocs/op
BenchmarkSlice_Index_Native-16                          309795711                3.766 ns/op           0 B/op          0 allocs/op
BenchmarkSlice_Index_Xunsafe-16                         100000000               10.53 ns/op            0 B/op          0 allocs/op
BenchmarkSlice_Index_Reflect-16                          5836300               206.1 ns/op            80 B/op         10 allocs/op
BenchmarkAppender_Append_Xunsafe-16                       933180              1142 ns/op            2000 B/op         11 allocs/op
BenchmarkAppender_Append_Reflect-16                       130723              8838 ns/op            4464 B/op        109 allocs/op
BenchmarkAppender_Append_Native-16                       2436530               475.1 ns/op          2040 B/op          8 allocs/op
```
* **'Native'** suffix represent statically typed code
* **'Xunsafe'** suffix represent reflection implemented by this library
* **'Reflect'** suffix represent implementation with golang native reflect package


## Contributing to xunsafe

XUnsafe is an open source project and contributors are welcome!


It would be nice if at some point native golang library exposes ability to create ref/deref/pointer for an actual value behind the interface
without that can be inlined to remove dependency/exposure of the private reflect pacakge types from this library.

```go
    i := 101
    var v interface = i
	vPtr := xxxx.Pointer(v)
	vPtrInterface := xxxx.Ref(v)
	cloned := xxxx.Deref(vPtrInterface)
```


## License

The source code is made available under the terms of the Apache License, Version 2, as stated in the file `LICENSE`.

Individual files may be made available under their own specific license,
all compatible with Apache License, Version 2. Please see individual files for details.

## Credits and Acknowledgements

**Library Author:** Adrian Witas

