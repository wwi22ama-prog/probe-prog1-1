/* AUFGABE: Rekursion
 * Erreichbare Punkte:
 */
package aufgabe7

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Funktion.
// Es müssen alle Tests grün werden.
// Zusatzanforderung: Die Funktion muss rekursiv bleiben.

// Berechnet den ganzzahligen Quotienten x / y.
// D.h. die Funktion ignoriert Nachkommastellen bzw. den Rest.
func Div(x, y int) int {
	if x < 0 {
		return -Div(-x, y)
	}
	if y < 0 {
		return -Div(x, -y)
	}
	if x < y {
		return 0
	}
	return 1 + Div(x-y, y)
}
