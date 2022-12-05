/* AUFGABE: Schere-Stein-Papier
 * Erreichbare Punkte: 25
 */
package aufgabe8

import "fmt"

// AUFGABENSTELLUNG:
// Hier ist ein Gerüst für das Spiel "Schere-Stein-Papier" vorgegeben.
// Vervollständigen Sie das Spiel, d.h. implementieren Sie die leeren Funktionen.

// Spielt eine Spielrunde Schere-Stein-Papier.
func PlayRPS() {
	inputP1 := ChooseItem(1)
	inputP2 := ChooseItem(2)

	result := GetResult(inputP1, inputP2)
	PrintResult(result)
}

// Fragt den Spieler mit der gegebenen Nummer
// nach seiner Wahl und liefert sie als int.
func ChooseItem(player int) int {
	fmt.Println("1: Rock")
	fmt.Println("2: Paper")
	fmt.Println("3: Scissors")
	fmt.Println()
	fmt.Printf("Player %d, choose your item: ", player)
	var input int
	fmt.Scanln(&input)
	if input >= 1 && input <= 3 {
		return input
	}
	fmt.Printf("%d is an invalid choice, try again!\n\n", input)
	return ChooseItem(player)
}

// Bestimmt das Ergebnis des Spiels.
// 0: Unentschieden
// 1: Spieler 1 gewinnt.
// 2: Spieler 2 gewinnt.
func GetResult(input1, input2 int) int {
	if input1 == input2 {
		return 0
	}
	if input1 == 3 && input2 == 1 {
		return 2
	}
	if input2 == 3 && input1 == 1 {
		return 1
	}
	if input1 > input2 {
		return 1
	}
	return 2
}

// Gibt das Ergebnis auf der Konsole aus.
func PrintResult(result int) {
	switch result {
	case 0:
		fmt.Println("Draw.")
	case 1:
		fmt.Println("Player 1 wins.")
	case 2:
		fmt.Println("Player 2 wins.")
	}
}
