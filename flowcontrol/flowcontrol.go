package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// Basic for loop
func myFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// For loop without "init" and "post" statements
func myFor2() {
	sum := 1
	for sum < 1000 { // no init and post
		sum += sum
	}
	fmt.Println(sum)
}

// Infinite loops
func myFor3() {
	// for {
	// }
}

// If statements
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If statement with "short statement"
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// If-else statement, with "short statement"
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // v can be used here
	}
	// can't use v here, though
	return lim
}

// MySqrt is an exercise
func MySqrt(x float64) float64 {
	z := 1.0
	old := 0.0
	for z != old {
		old = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

// Defer statements get executed after a function returns
func myDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// When "stacking" defers, they get executed in LIFO order
func myStackedDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	fmt.Printf("Hello world!\n")

	myFor()

	myFor2()

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)

	fmt.Println(MySqrt(10000))

	// Switch statement
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch evaluation order

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch without condition as alternative to long if-else statements
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	myDefer()

	myStackedDefer()
}
