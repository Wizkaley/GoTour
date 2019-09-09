package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib1 := 0
	fib2 := 1
	var fib3 int
	return func() int{
		
		fib3 = fib1 + fib2
		fib1 = fib2
		fib2 = fib3
		//fib = sum
		return fib1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
