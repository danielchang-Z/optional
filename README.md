# Optional

Optional is a wrapper for representing 'optional' (or 'nullable') objects who may not (yet) contain a valid value.

[https://godoc.org/github.com/danielzhangy/optional](https://godoc.org/github.com/danielzhangy/optional)

`go get -u github.com/danielzhangy/optional`

The purpose of the package is to provide a type-level solution for representing optional values instead of using nil in code. Instead traditional nil value or reference  handling in Go is rough akin to

```golang
if obj != nil {
    // handle the non-nil object
} else {
    // handle the nil object
}
```
, optional package allows programmer to handle nil value or reference like
```golang
optional.Of("some object").
    IfPresent(func(i interface{}) {fmt.Println(i)})
```
.


# Usage
Using `optional.Absent()` to create an empty Optional, 

```golang
opt := Absent()
fmt.Println(opt)

// Example output:
// &{}
```

, or using `optional.Of()` to create an optional with given non-nil value,

```golang
opt := Of(1)
fmt.Println(opt)

// Example output:
// &{1}
```

. If you not sure that the value if nil or not, you can use `optional.OfNilable()` to create, it will generate a present or absent instance depended on the given value is non-nil or not.

```golang
opt1 := OfNilable(1)
fmt.Println("opt1: ", opt1)
opt2 := OfNilable(nil)
fmt.Println("opt2: ", opt2)

// Example output:
// opt1:  &{1}
// opt2:  &{}
```

[Read the package documentation for more information.](https://godoc.org/github.com/danielzhangy/optional)

# Contributing
We welcome pull requests, bug fixes and issue reports. With that said, the bar for adding new symbols to this package is intentionally set high.

Before proposing a change, please discuss your change by raising an issue.

# License
MIT License