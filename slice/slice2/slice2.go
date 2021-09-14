package slice2

import "fmt"

func Slice2() {
	a := make([]int, 32)
	b := a[1:16]
	fmt.Printf("&a: %p\n", &a)
	fmt.Printf("&b: %p\n", &b)

	a = append(a, 1)
	a[2] = 42
	fmt.Printf("&a: %p\n", &a)
}