//Let's have some fun with functions.
//Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

func fibonacci() func() int {
	prev := 0
	curr := 1
	return func() int {
		prev, curr = curr, prev+curr
		return curr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}