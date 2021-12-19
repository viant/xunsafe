# xunsafe (faster golang reflection)

[![GoReportCard](https://goreportcard.com/badge/github.com/viant/xunsafe)](https://goreportcard.com/report/github.com/viant/xunsafe)
[![GoDoc](https://godoc.org/github.com/viant/xunsafe?status.svg)](https://godoc.org/github.com/viant/xunsafe)

This library is compatible with Go 1.17+

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

	fooAddr := unsafe.Pointer(foo)
	*(fooID.Addr(fooAddr).(*int)) = 201
	fmt.Printf("foo.ID: %v\n", foo.ID)//prints 201
}
```

###### Field Value

Field Value returns an interface{} wrapping actual field value


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
	fmt.Printf("foo.ID: %v\n", fooID.Value(fooAddr))//prints 101
}
```

For base golang type Field Addr and Value got optimized with casting unsafe address to actual corresponding type. 
For example for filed with int type, the casting come in form ```(*int)(unsafe.Pointer(structAddr + field.Offset))```
in other cases ```reflect.NewAt(field.Type, unsafe.Pointer(structAddr + field.Offset) ``` is used.

###### Unsafe struct casting registry

Given that reflect.NewAt is quite slow, you can register custom unsafe type casting bypassing reflect.NewAt all together

```go
    xunsafe.Register(reflect.TypeOf(time.Time{}), func(addr unsafe.Pointer) interface{} {
		return (*time.Time)(addr)
	})
    xunsafe.Register(reflect.TypeOf(&time.Time{}), func(addr unsafe.Pointer) interface{} {
		return (**time.Time)(addr)
	})

```


### Slice range

```go
        type T {ID int}
        aSlice := NewSlice(reflect.TypeOf([]T))
		var items = []T{
			{ID:1}, {ID:2}, {ID:3},
        }   
		aSlice.Range(unsafe.Pointer(&items), func(index int, addr unsafe.Pointer) bool {
			item := (*T)(addr)
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
		index := aSlice.Index(unsafe.Pointer(&items))
		for i :=0;i<len(items);i++ {
            item := (*T)(index(i))
            fmt.Printf("%+v\n", item)
        }
	
```

### Slice appender

```go
        type T {ID int}
        aSlice := NewSlice(reflect.TypeOf([]T))
		var items []T
		
		appender := aSlice.Appender(unsafe.Pointer(&items))
        appender.Append(unsafe.Pointer(&T{ID:1}),unsafe.Pointer(&T{ID:2}))
        appender.Append(unsafe.Pointer(&T{ID:3}),unsafe.Pointer(&T{ID:4}))
        fmt.Printf("%v\n", items)
		
```


### Benchmark

Accessor/Mutator benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/viant/xunsafe
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkField_Accessor_Native-16               15784005                69.36 ns/op            0 B/op          0 allocs/op
BenchmarkField_Accessor_Xunsafe-16               6739143               178.8 ns/op             0 B/op          0 allocs/op
BenchmarkField_Accessor_Value-16                 2807060               439.8 ns/op            56 B/op          3 allocs/op
BenchmarkField_Accessor_Reflect-16               1470356               826.7 ns/op            72 B/op          4 allocs/op
BenchmarkField_Accessor_PtrXunsafe-16            4269766               278.6 ns/op             0 B/op          0 allocs/op
BenchmarkField_Accessor_Reflect_Ptr-16           1902192               635.8 ns/op             0 B/op          0 allocs/op
BenchmarkField_Mutator_Native-16                25268916                46.56 ns/op            0 B/op          0 allocs/op
Benchmark_Mutator_Fast-16                        9487674               125.5 ns/op             0 B/op          0 allocs/op
Benchmark_Mutator_Fast_Ptr-16                    5344639               224.0 ns/op             0 B/op          0 allocs/op
BenchmarkField_Mutator_Reflect-16                2005026               604.2 ns/op            48 B/op          3 allocs/op
BenchmarkField_Mutator_Reflect_Ptr-16            2423137               483.6 ns/op             0 B/op          0 allocs/op
BenchmarkSlice_Index_Native-16                  14769567                76.22 ns/op
BenchmarkSlice_Index_Xunsafe-16                  4705651               243.9 ns/op
BenchmarkSlice_Index_Reflect-16                   643263              2001 ns/op
BenchmarkAppender_Append_Xunsafe-16                57892             19709 ns/op
BenchmarkAppender_Append_Relfect-16                18657             67098 ns/op
BenchmarkAppender_Append_Native-16                354895              3180 ns/op
PASS
ok      github.com/viant/xunsafe        26.395s
awitas@AWITAS-C02C42QCMD6R xunsafe % go teset -bench=.                              
go teset: unknown command
Run 'go help' for usage.
awitas@AWITAS-C02C42QCMD6R xunsafe % go test -bench=. 
goos: darwin
goarch: amd64
pkg: github.com/viant/xunsafe
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkField_Accessor_Native-16               933782421                1.215 ns/op           0 B/op          0 allocs/op
BenchmarkField_Accessor_Xunsafe-16              594855732                1.934 ns/op           0 B/op          0 allocs/op
BenchmarkField_Accessor_Value-16                16291875                67.54 ns/op           44 B/op          3 allocs/op
BenchmarkField_Accessor_Reflect-16               9490488               115.7 ns/op            56 B/op          4 allocs/op
BenchmarkField_Accessor_PtrXunsafe-16           100000000               10.27 ns/op            0 B/op          0 allocs/op
BenchmarkField_Accessor_Reflect_Ptr-16          23711229                47.98 ns/op            0 B/op          0 allocs/op
BenchmarkField_Mutator_Native-16                1000000000               0.9083 ns/op          0 B/op          0 allocs/op
Benchmark_Mutator_Xunsafe-16                    778190858                1.529 ns/op           0 B/op          0 allocs/op
Benchmark_Mutator_Xunsafe_Ptr-16                156936856                7.637 ns/op           0 B/op          0 allocs/op
BenchmarkField_Mutator_Reflect-16               12456297                93.22 ns/op           32 B/op          3 allocs/op
BenchmarkField_Mutator_Reflect_Ptr-16           32652355                36.07 ns/op            0 B/op          0 allocs/op
BenchmarkSlice_Index_Native-16                  205128967                5.789 ns/op           0 B/op          0 allocs/op
BenchmarkSlice_Index_Xunsafe-16                 157107307                7.303 ns/op           0 B/op          0 allocs/op
BenchmarkSlice_Index_Reflect-16                  5529320               201.7 ns/op            80 B/op         10 allocs/op
BenchmarkAppender_Append_Xunsafe-16               857829              1331 ns/op            2128 B/op         13 allocs/op
BenchmarkAppender_Append_Relfect-16               140394              8425 ns/op            4464 B/op        109 allocs/op
BenchmarkAppender_Append_Native-16               2400387               480.7 ns/op          2040 B/op          8 allocs/op
```
* **'Native'** suffix represent statically typed code
* **'Xunsafe'** suffix represent reflection implemented by this library
* **'Reflect'** suffix represent implementation with golang native reflect package


## License

The source code is made available under the terms of the Apache License, Version 2, as stated in the file `LICENSE`.

Individual files may be made available under their own specific license,
all compatible with Apache License, Version 2. Please see individual files for details.

## Credits and Acknowledgements

**Library Author:** Adrian Witas

