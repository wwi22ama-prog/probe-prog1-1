package aufgabe7

import "fmt"

func ExampleDiv() {

	fmt.Println(Div(3, 2))
	fmt.Println(Div(2, 3))
	fmt.Println(Div(20, 2))
	fmt.Println(Div(-3, 2))
	fmt.Println(Div(3, -2))
	fmt.Println(Div(-3, -2))

	// Output:
	// 1
	// 0
	// 10
	// -1
	// -1
	// 1
}
