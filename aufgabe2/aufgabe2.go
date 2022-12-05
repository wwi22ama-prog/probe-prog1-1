/* AUFGABE: Listen
 * Erreichbare Punkte:
 */
package aufgabe2

/* AUFGABENSTELLUNG: */
// Erwartet eine int-Slice list.
// Liefert eine int-Slice, die an Stelle n die Summe
// der Elemente aus list bis zu Stelle n enthält.
func ArraySums(list []int) []int {
	sum := 0
	result := []int{}
	for _, el := range list {
		sum += el
		result = append(result, sum)
	}
	return result
}
