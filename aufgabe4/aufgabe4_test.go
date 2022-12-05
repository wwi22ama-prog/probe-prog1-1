package aufgabe4

import "fmt"

func ExampleFindAll() {
	l1 := []int{1, 4, 17, 2, 1, 5, 10, 5, 2}
	fmt.Println(FindAll(l1, 1))
	fmt.Println(FindAll(l1, 4))
	fmt.Println(FindAll(l1, 5))
	fmt.Println(FindAll(l1, 42))

	// Output:
	// [0 4]
	// [1]
	// [5 7]
	// []
}
