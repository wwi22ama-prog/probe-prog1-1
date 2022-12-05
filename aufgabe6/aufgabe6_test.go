package aufgabe6

import "fmt"

func ExampleEntry_Translations() {
	e1 := NewEntry("Himmel", "sky")
	fmt.Println(e1)

	e1.AddTranslation("heaven")
	fmt.Println(e1)

	// Output:
	// Himmel : sky
	// Himmel : sky,heaven
}
