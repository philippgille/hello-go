package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

// A "method" is a function with a receiver argument
func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type myFloat float64

// Methods can only have receivers that are in the same package
// Alias a built-in method to create a method for a built-in type (like "int")
func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers enable a method to modify the value the pointer points to
// Otherwise, the parameter would be a copy of the value
func (v *vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// The same logic can be implemented in a function, which requires the parameter to be passed with "&x" to pass the pointer
func scaleAsFunction(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ===============

// An interface is a type that defines a set of method signatures
type abser interface {
	abs() float64
}

type myFloat2 float64

// Implementation of method signature for type myFloat
func (f myFloat2) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex2 struct {
	X, Y float64
}

// Implementation of method signature for type *vertex
func (v *vertex2) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// When declaring a variable of an interface type, all types that implement the defined method signatures can be assigned
func myInterface() {
	var a abser
	f := myFloat2(-math.Sqrt2)
	v := vertex2{3, 4}

	a = f  // a MyFloat implements abser
	a = &v // a *Vertex implements abser

	// In the following line, v is a vertex (not *vertex)
	// and does NOT implement abser.
	//a = v

	fmt.Println(a.abs())
}

// ==============

type i interface {
	M()
}

type t struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t t) M() {
	fmt.Println(t.S)
}

// =============

type t2 struct {
	S string
}

// Implementations can be for pointers...
func (t *t2) M() {
	fmt.Println(t.S)
}

type f float64

//... as well as for values
func (f f) M() {
	fmt.Println(f)
}

func describe(i i) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func myInterfaceValues() {
	var i i // Interface value

	i = &t2{"Hello"} // make i a pointer, so i.M() can be called
	describe(i)      // (&{Hello}, *main.t2)
	i.M()

	i = f(math.Pi)
	describe(i) // (3.141592653589793, main.f)
	i.M()
}

// ===========

type i2 interface {
	M()
}

type t3 struct {
	S string
}

// It's common to write methods that gracefully handle nil receivers
// This prevents null pointer exceptions
func (t *t3) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func nilUnderlyingValues() {
	var i i2 // Interface type

	var t *t3   // no value assigned
	i = t       // nil underlying value
	describe(i) // (<nil>, *main.t3)
	i.M()

	i = &t3{"hello"}
	describe(i) // (&{hello}, *main.t3)
	i.M()
}

// ===========

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Empty interface
// Every type satisfies the empty interface
// Use it in methods that shall operate on all types
func myEmptyInterface() {
	var i interface{}
	describe2(i) // (<nil>, <nil>)

	i = 42
	describe2(i) // (42, int)

	i = "hello"
	describe2(i) // (hello, string)
}

// =========

// With type assertions you can test whether the value underlying an interface value is of a specific type
func typeAssertions() {
	var i interface{} = "hello" // Interface value

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	//f = i.(float64) // panic
	//fmt.Println(f)
}

// Type switches
func do(i interface{}) {
	switch v := i.(type) { // Note the keyword "type" here, although the syntax looks like type assertion
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.abs())

	f := myFloat(-math.Sqrt2) // Cast
	fmt.Println(f.abs())

	v.scale(10)
	fmt.Println(v.abs()) // Prints 50 instead of 5, because scale() modified the value

	v = vertex{3, 4}
	scaleAsFunction(&v, 10) // Must be called with &v, to pass the pointer
	fmt.Println(v.abs())

	// Pointer indirection
	// When calling a method with a pointer receiver, you can use v directly, because the Go compiler makes "(&v)" out of it
	v = vertex{3, 4}
	v.scale(10) // Go calls this as "(&v).scale(10)"

	// When a pointer is needed as an argument, you can get the pointer directly when declaring the variable
	p := &vertex{3, 4}
	scaleAsFunction(p, 10)
	fmt.Println(p.abs()) // Note: This is the same in the opposite direction:
	// abs() has a value receiver, but can be called with "p", which Go automatically calls as "(*p).abs()"

	// Choosing value or pointer receiver:
	// Value receivers make the method immutable
	// Pointer receivers enable the modification of a value
	// and can also be more efficient, because the value doesn't need to be copied
	// All methods on a given type should have either value or pointer receivers, but not a mixture of both

	var i i = t{"hello"}
	i.M()

	myInterfaceValues()

	nilUnderlyingValues()

	// Calling a method on a nil interface leads to runtime error
	var i2 i2
	describe(i2) // (<nil>, <nil>)
	// i.M() // panic: runtime error: invalid memory address or nil pointer dereference

	myEmptyInterface()

	typeAssertions()

	do(21)
	do("hello")
	do(true)
}
