# xunsafe (faster golang reflection)

[![GoReportCard](https://goreportcard.com/badge/github.com/viant/xunsafe)](https://goreportcard.com/report/github.com/viant/xunsafe)
[![GoDoc](https://godoc.org/github.com/viant/xunsafe?status.svg)](https://godoc.org/github.com/viant/xunsafe)

This library is compatible with Go 1.13+

Please refer to [`CHANGELOG.md`](CHANGELOG.md) if you encounter breaking changes.

- [Motivation](#motivation)
- [Introduction](#introduction)
- [Usage](#usage)
- [Benchmark](#benchmark)
- [Contribution](#contributing-to-xunsafe)
- [License](#license)

## Motivation

In order to solve a problem that can work with any golang type, one has no choice but to use reflection.
Native golang reflection comes with hefty performance price, on benchmarking simple getter/setter case 
to manipulate struct dynamically I've seen around 100 time worst performance comparing to 
statically typed code. 
I believe that, native reflect package could be easily implemented way better to provide optimized performance.
This library comes with reflection implementation that greatly improved performance, that is 50x time faster than native golang reflection. 
What that means that  extra overhead of using reflection is only around twice comparing to statically typed code.

## Introduction

In order to achieve better performance, this library uses unsafe.Pointer along with StructField.Offset to effectively access/modify struct fields.


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
        fooAddr := xunsafe.Addr(&foos[i])
        fooID.SetInt(fooAddr, i)
        fooName.SetString(fooAddr, fmt.Sprintf("name %d", i))
    }

    for i := range foos {
        fooAddr := xunsafe.Addr(&foos[i])
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

	fooAddr := xunsafe.Addr(foo)
	*(fooID.Addr(fooAddr).(*int)) = 201
	fmt.Printf("foo.ID: %v\n", foo.ID)//prints 201
}
```

###### Field Value

Field Addr returns an interface{} wrapping actual field value


```go
func ExampleAddr() {
	type Foo struct {
		ID int
		Name string
	}
	fooType := reflect.TypeOf(Foo{})
	fooID := xunsafe.FieldByName(fooType, "ID")
	foo := &Foo{ID: 101, Name: "name 101"}

	fooAddr := xunsafe.Addr(foo)
	fmt.Printf("foo.ID: %v\n", fooID.Value(fooAddr))//prints 101
}
```

For base golang type Field Addr and Value got optimized with casting unsafe address to actual corresponding type. 
For example for filed with int type, the casting come in form ```(*int)(unsafe.Pointer(structAddr + field.Offset))```
in other cases ```reflect.NewAt(field.Type, unsafe.Pointer(structAddr + field.Offset) ``` is used.

###### Unsafe struct casting registry

Given that reflect.NewAt is quite slow, you can register custom unsafe type casting bypassing reflect.NewAt all together

```go
    xunsafe.Register(reflect.TypeOf(time.Time{}), func(addr uintptr) interface{} {
		return (*time.Time)(unsafe.Pointer(addr))
	})
    xunsafe.Register(reflect.TypeOf(&time.Time{}), func(addr uintptr) interface{} {
		return (**time.Time)(unsafe.Pointer(addr))
	})

```

### Benchmark

Accessor/Mutator benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/viant/xunsafe
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkField_Accessor_Native-16         	890931754	          1.235 ns/op	       0 B/op	       0 allocs/op
BenchmarkField_Accessor_Fast-16           	558196452	          2.217 ns/op	       0 B/op	       0 allocs/op
BenchmarkField_Accessor_Reflect-16        	 8435084	          134.9 ns/op	      56 B/op	       4 allocs/op
BenchmarkField_Accessor_PtrFast-16        	94770246	          12.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkField_Accessor_Reflect_Ptr-16    	17914276	          68.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkField_Mutator_Native-16          	1000000000	         0.9271 ns/op	       0 B/op	       0 allocs/op
Benchmark_Mutator_Fast-16                 	879740266	          1.333 ns/op	       0 B/op	       0 allocs/op
Benchmark_Mutator_Fast_Ptr-16             	100000000	          10.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkField_Mutator_Reflect-16         	11812810	          99.02 ns/op	      32 B/op	       3 allocs/op
BenchmarkField_Mutator_Reflect_Ptr-16     	22566876	          54.51 ns/op	       0 B/op	       0 allocs/op
```
'Fast' suffix represent reflection implemented by this library.


## License

The source code is made available under the terms of the Apache License, Version 2, as stated in the file `LICENSE`.

Individual files may be made available under their own specific license,
all compatible with Apache License, Version 2. Please see individual files for details.

## Credits and Acknowledgements

**Library Author:** Adrian Witas

