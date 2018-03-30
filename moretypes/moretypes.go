package main

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

// "var p *int" declares a pointer to an int value
// "&x" generates a pointer to operand x.
// "*p" dereferences pointer p
func myPointer() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// A struct is a collection of fields.
type vertex struct {
	X int
	Y int
}

// Struct literals
var (
	v2 = vertex{1, 2}  // has type Vertex
	v3 = vertex{X: 1}  // Y:0 is implicit
	v4 = vertex{}      // X:0 and Y:0
	p2 = &vertex{1, 2} // has type *Vertex
)

// Arrays
// Arrays can't be resized
func myArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// Slices
// Slices select a "half-open" range, so it includes the first, but excludes the last element
// Working with slices is more common than working with arrays
func mySlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

// Modifying a slice modifies the elements of the underlying array
func modifySlice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// Slice literals
// Creating a slice from scratch automatically creates the underlying array
func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// Slice defaults - the low and high bounds can be omitted
func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// Slice length and capacity
func sliceLengthAndCapacity() {
	// Lambda
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

// Make slice
func makeSlice() {
	// Lambda
	printSlice := func(s string, x []int) {
		fmt.Printf("%s len=%d cap=%d %v\n",
			s, len(x), cap(x), x)
	}

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func sliceOfSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Appending to slice
// If the capacity is reached, a new underlying array is created
func appendToSlice() {
	// Lambda
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

// The range form of the for loop iterates over a slice or map.
// Range returns two values: The index and the element at that index
func myRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// One of the two return values can be omitted
func myRange2() {
	pow := make([]int, 10) // len 10, cap implicitly set to len, leading to [0 0 0 ...]
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// Pic is an exercise function
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dx)
	for x := range result {
		result[x] = make([]uint8, dy)
		for y := range result[x] {
			color := uint8((x + y) / 2)
			result[x][y] = color
		}
	}
	return result
}

type vertex2 struct {
	Lat, Long float64
}

// Maps can be created with make()
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
func myMap() {
	// Optional declaration:
	// var m map[string]Vertex
	m := make(map[string]vertex2)
	m["Bell Labs"] = vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

// "Map literals"
// Maps can be initialized during declaration, requiring key and value
func mapLiterals() {
	var m = map[string]vertex2{
		"Bell Labs": vertex2{
			40.68433, -74.39967,
		},
		"Google": vertex2{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.
func mapLiterals2() {
	var m = map[string]vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

// Mutating maps
func mutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

// WordCount is an exercise
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Split(s, " ")
	for _, word := range words {
		result[word] = result[word] + 1
	}
	return result
}

// Functions are values, too
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// Closures are functions with access to variables outside of their scope
// A closure keeps its state across calls
// Each closure has it's own state
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func myClosure() {
	pos, neg := adder(), adder() // Two closures
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // Each with its own state
			neg(-2*i), // =
		)
	}
}

// Exercise
// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	n := 0
	m := 1
	return func() int {
		result := n + m
		n = m
		m = result
		return result
	}
}
func fibTester() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	fmt.Printf("Hello world!\n")

	myPointer()

	fmt.Println(vertex{1, 2})

	// Accessing fields via dot
	v := vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs
	v = vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v2, p2, v3, v4)

	myArray()

	mySlice()

	modifySlice()

	sliceLiterals()

	sliceDefaults()

	sliceLengthAndCapacity()

	// The "zero value" of a slice is nil
	// A nil slice has a length and capacity of 0 and no underlying array
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	makeSlice()

	sliceOfSlice()

	appendToSlice()

	myRange()

	myRange2()

	pic.Show(Pic)

	myMap()

	mapLiterals()

	mapLiterals2()

	mutatingMaps()

	wc.Test(WordCount)

	functionValues()

	myClosure()

	fibTester()
}
