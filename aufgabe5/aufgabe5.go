/* AUFGABE: Structs
 * Erreichbare Punkte: 20
 */
package aufgabe5

// AUFGABENSTELLUNG:
// Reparieren Sie die folgende Methode.
// Es müssen alle Tests grün werden.

// Fügt ein Element zu dict hinzu,
// das das Wortpaar (de,en) enthält.
// Falls dict schon einen Eintrag für de enthält,
// soll dessen en ersetzt werden.
func (dict *Dictionary) SetEntry(de, en string) {
	dict.Entries = append(dict.Entries, Entry{de, en})
}
