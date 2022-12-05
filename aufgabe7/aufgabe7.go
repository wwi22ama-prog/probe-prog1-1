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
	return 1 + Div(x-y, y)
}
