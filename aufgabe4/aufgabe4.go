/* AUFGABE: Listen
 * Erreichbare Punkte:
 */
package aufgabe4

// AUFGABENSTELLUNG:
// Liefert eine Liste mit allen Stellen in list,
// an denen x vorkommt.
func FindAll(list []int, x int) []int {
	result := []int{}
	for pos, el := range list {
		if el == x {
			result = append(result, pos)
		}
	}
	return result
}
