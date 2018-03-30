package main

import (
	"fmt"
	"math" // for Sqrt(int)
	"math/cmplx"
	"math/rand" // for Intn(int)
)

// Using import
func myRand(n int) int {
	return rand.Intn(n)
}

// Using other import
func mySqrt(n int) float64 {
	return math.Sqrt(float64(n)) // float64(n) casts n
}

// MyExport is an exported function because it starts with a capital letter
func MyExport() {
	fmt.Printf("hello export")
}

// Arguments
func add(a int, b int) int {
	return a + b
}

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "naked" return
}

// Variables with package scope
var java, csharp bool

// Variables with initializers take the type of the initializers
var i, j = 1, 2

// Variables can be declared in a block
var (
	ToBe          = false // Omitted type "bool"
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i) // Omitted type "complex128"
)

// Pi is a constant
// Constants can be character, string, boolean, or numeric values.
const Pi = 3.14

// Numeric constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Printf("hello, world\n")

	fmt.Println(myRand(10))

	fmt.Println(mySqrt(7))
	fmt.Println(add(3, 4))

	var a, b = swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(23))

	java = true // Note: the below introduced ":=" can be used as well, because "java" is already declared as var in the package scope
	fmt.Println(java, csharp)

	fmt.Println(i, j)

	// Inside of functions, "var" is not necessary and ":=" (short assignment statement) can be used
	k := 3
	fmt.Println(k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables without initial value are given their "zero" value
	var i int
	var f float64
	var b2 bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b2, s)

	// Type conversions
	var x, y int = 3, 4
	var f2 = math.Sqrt(float64(x*x + y*y))
	var z = uint(f2)
	fmt.Println(x, y, z)

	// Type inference
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

	fmt.Println("Happy", Pi, "Day")
}
