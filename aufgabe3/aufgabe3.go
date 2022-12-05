/* AUFGABE: Rekursion
 * Erreichbare Punkte:
 */
package aufgabe3

// AUFGABENSTELLUNG:
// Liefert die Potenz "2 hoch x".
// Zusatzanforderung: Die Funktion muss rekursiv sein.
func Power2(x int) float64 {
	if x < 0 {
		return 1 / Power2(-x)
	}
	if x == 0 {
		return 1
	}
	return 2 * Power2(x-1)
}
