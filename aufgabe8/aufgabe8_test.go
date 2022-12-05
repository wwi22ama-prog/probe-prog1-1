package aufgabe8

import "fmt"

func ExampleGetResult() {
	fmt.Println(GetResult(1, 1)) // Unentschieden
	fmt.Println(GetResult(1, 2)) // p1: Stein, p2: Papier -> p2
	fmt.Println(GetResult(1, 3)) // p1: Stein, p2: Schere -> p1
	fmt.Println(GetResult(2, 1)) // p1: Papier, p2: Stein -> p1
	fmt.Println(GetResult(2, 2)) // Unentschieden
	fmt.Println(GetResult(2, 3)) // p1: Papier, p2: Schere -> p2
	fmt.Println(GetResult(3, 1)) // p1: Schere, p2: Stein -> p2
	fmt.Println(GetResult(3, 2)) // p1: Schere, p2: Papier -> p1
	fmt.Println(GetResult(3, 3)) // Unentschieden

	// Output:
	// 0
	// 2
	// 1
	// 1
	// 0
	// 2
	// 2
	// 1
	// 0
}
