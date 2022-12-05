package aufgabe5

import "fmt"

func ExampleDictionary_SetEntry() {
	d1 := Dictionary{[]Entry{}}

	d1.SetEntry("Himmel", "sky")
	fmt.Println(d1)
	fmt.Println()

	d1.SetEntry("Erde", "earth")
	fmt.Println(d1)
	fmt.Println()

	d1.SetEntry("Himmel", "heaven")
	fmt.Println(d1)
	fmt.Println()

	// Output:
	// 0 - Himmel : sky
	//
	// 0 - Himmel : sky
	// 1 - Erde : earth
	//
	// 0 - Himmel : heaven
	// 1 - Erde : earth
}
