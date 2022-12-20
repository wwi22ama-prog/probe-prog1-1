package aufgabe6

import (
	"fmt"
	"strings"
)

func ExampleEntry_Translations() {
	e1 := NewEntry("Himmel", "sky")
	fmt.Println(e1)

	e1.AddTranslation("heaven")
	fmt.Println(e1)

	e2 := NewEntry("Fenster", "")
	e2.AddTranslation("window")
	fmt.Println(e2)

	// Output:
	// Himmel : sky
	// Himmel : sky,heaven
	// Fenster : window
}

func Example_beispiel_join() {
	// Beispiel für strings.Join():
	stringList := []string{"Hallo", "du", "schöne", "Welt"}

	s3 := strings.Join(stringList, " ")
	s4 := strings.Join(stringList, ":")
	s5 := strings.Join(stringList, "\n")
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	// Output:
	// Hallo du schöne Welt
	// Hallo:du:schöne:Welt
	// Hallo
	// du
	// schöne
	// Welt
}
