/* AUFGABE: Listen
 * Erreichbare Punkte: 10
 */
package aufgabe1

// AUFGABENSTELLUNG:
// Erwartet drei Zahlen m,n,max > 0 und liefert eine Liste mit allen
// gemeinsamen Vielfachen von m und n, die nicht größer als max sind.
func CommonMultiples(m, n, max int) []int {
	result := []int{}
	for x := m; x <= max; x += m {
		if x%n == 0 {
			result = append(result, x)
		}
	}
	return result
}
