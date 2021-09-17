package slice4

import (
	"fmt"
	"reflect"
)

func Slice4() {
	v1 := [2]int{1, 2}
	v2 := [2]int{1, 2}
	fmt.Println("v1 == v2: ", reflect.DeepEqual(v1, v2))

	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	m2 := map[string]int{"one": 1, "three": 3, "two": 2}
	fmt.Println("m1 == m2: ", reflect.DeepEqual(m1, m2))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2: ", reflect.DeepEqual(s1, s2))
}