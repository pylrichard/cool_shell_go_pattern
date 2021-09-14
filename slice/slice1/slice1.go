package slice1

import "fmt"

func Slice1() {
	foo := make([]int, 6)
	foo[3] = 42
	foo[4] = 100
	fmt.Printf("foo: %v\n", foo)

	bar := foo[1:4]
	bar[1] = 99
	fmt.Printf("bar: %v\n", bar)
}