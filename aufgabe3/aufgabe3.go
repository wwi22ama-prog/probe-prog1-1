/* AUFGABE: Rekursion
 * Erreichbare Punkte:
 */
package aufgabe3

// AUFGABENSTELLUNG:
// Schreiben Sie eine rekursive Funktion power2(), die einen int-Parameter x erwartet.
// Die Funktion soll die Potenz "2 hoch x" berechnen und zurückliefern.
func Power2(x int) float64 {
	if x < 0 {
		return 1 / Power2(-x)
	}
	if x == 0 {
		return 1
	}
	return 2 * Power2(x-1)
}
