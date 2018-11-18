package optional

import (
	"fmt"
)

func ExampleAbsent() {
	opt := Absent()
	fmt.Println(opt)

	val := opt.Or(1)
	fmt.Println("val:", val)

	opt.Substitute(Of(10)).
		IfPresent(func(i interface{}) {
			fmt.Printf("I am a consumer and will consume the value[%v].\n", i)
		})

	// Example output:
	// &{}
	// val:1
	// I am a consumer and will consume the value[10].
}

func ExampleOf() {
	opt := Of(1)
	fmt.Println(opt)

	val := opt.Or(10)
	fmt.Println("val:", val)

	sum := opt.Transform(
		func(i interface{}) interface{} {
			return i.(int) + i.(int)
		}).Get()
	fmt.Println("sum:", sum)

	// Example output:
	// &{1}
	// val:10
	// sum:2
}

func ExampleOfNilable() {
	opt1 := OfNilable(1)
	fmt.Println("opt1: ", opt1)
	opt2 := OfNilable(nil)
	fmt.Println("opt2: ", opt2)

	consumer := func(i interface{}) {
		fmt.Printf("I am a consumer and will consume the value[%v].\n", i)
	}
	worker := func() {
		fmt.Println("I am a worker.")
	}

	opt1.IfPresentOrElse(consumer, worker)
	opt2.IfPresentOrElse(consumer, worker)

	// Example output:
	// opt1:  &{1}
	// opt2:  &{}
	// I am a consumer and will consume the value[1].
	// I am a worker.
}

func ExampleOptional_IsPresent() {
	opt1 := OfNilable(1)
	fmt.Println("opt1.IsPresent(): ", opt1.IsPresent())
	opt2 := OfNilable(nil)
	fmt.Println("opt2.IsPresent(): ", opt2.IsPresent())

	// Example output:
	// opt1.IsPresent():  true
	// opt2.IsPresent():  false
}

func ExampleOptional_IsAbsent() {
	opt1 := OfNilable(1)
	fmt.Println("opt1.IsAbsent(): ", opt1.IsAbsent())
	opt2 := OfNilable(nil)
	fmt.Println("opt2.IsAbsent(): ", opt2.IsAbsent())

	// Example output:
	// opt1.IsAbsent():  false
	// opt2.IsAbsent():  true
}

func ExampleOptional_Get() {
	opt1 := OfNilable(1)
	fmt.Println("opt1.Get(): ", opt1.Get())
	opt2 := OfNilable(nil)
	fmt.Println("opt2.Get(): ", opt2.Get())

	// Example output:
	// opt1.Get():  1
	// panic: Optional.get() cannot be called on an absent value: illegal status
}

func ExampleOptional_Or() {
	opt1 := OfNilable(1)
	fmt.Println("opt1.Or(0): ", opt1.Or(0))
	opt2 := OfNilable(nil)
	fmt.Println("opt2.Or(0): ", opt2.Or(0))

	// Example output:
	// opt1.Or(0):  1
	// opt2.Or(0):  0
}

func ExampleOptional_OrNil() {
	opt1 := OfNilable(1)
	fmt.Println("opt1.OrNil(): ", opt1.OrNil())
	opt2 := OfNilable(nil)
	fmt.Println("opt2.OrNil(): ", opt2.OrNil())

	// Example output:
	// opt1.OrNil():  1
	// opt2.OrNil():  <nil>
}

func ExampleOptional_OrSupply() {
	supplier := func() interface{} {
		return 0
	}
	opt1 := OfNilable(1)
	fmt.Println("opt1.OrSupply(supplier): ", opt1.OrSupply(supplier))
	opt2 := OfNilable(nil)
	fmt.Println("opt2.OrSupply(supplier): ", opt2.OrSupply(supplier))

	// Example output:
	// opt1.OrSupply(supplier):  1
	// opt2.OrSupply(supplier):  0
}

func ExampleOptional_Substitute() {
	subsitutation := Of(0)
	opt1 := OfNilable(1)
	fmt.Println("opt1.Substitute(subsitutation): ", opt1.Substitute(subsitutation))
	opt2 := OfNilable(nil)
	fmt.Println("opt2.Substitute(subsitutation): ", opt2.Substitute(subsitutation))

	// Example output:
	// opt1.Substitute(subsitutation):  &{1}
	// opt2.Substitute(subsitutation):  &{0}
}

func ExampleOptional_Transform() {
	transformation := func(i interface{}) interface{} {
		x := i.(int)
		return x + x
	}

	opt1 := OfNilable(1)
	fmt.Println("opt1.Transform(transformation): ", opt1.Transform(transformation))
	opt2 := OfNilable(nil)
	fmt.Println("opt2.Transform(transformation): ", opt2.Transform(transformation))

	// Example output:
	// opt1.Transform(transformation):  &{2}
	// opt2.Transform(transformation):  &{}
}

func ExampleOptional_IfPresent() {
	consumer := func(i interface{}) {
		fmt.Println(i)
	}

	opt1 := OfNilable(1)
	fmt.Print("opt1.IfPresent(transformation): ")
	opt1.IfPresent(consumer)
	opt2 := OfNilable(nil)
	fmt.Print("opt2.IfPresent(transformation): ")
	opt2.IfPresent(consumer)

	// Example output:
	// opt1.IfPresent(transformation): 1
	// opt2.IfPresent(transformation):
}

func ExampleOptional_IfPresentOrElse() {
	consumer := func(i interface{}) {
		fmt.Println(i)
	}
	worker := func() { fmt.Println("worker") }

	opt1 := OfNilable(1)
	fmt.Print("opt1.IfPresentOrElse(transformation): ")
	opt1.IfPresentOrElse(consumer, worker)
	opt2 := OfNilable(nil)
	fmt.Print("opt2.IfPresentOrElse(transformation): ")
	opt2.IfPresentOrElse(consumer, worker)

	// Example output:
	// opt1.IfPresentOrElse(transformation): 1
	// opt2.IfPresentOrElse(transformation): worker
}
