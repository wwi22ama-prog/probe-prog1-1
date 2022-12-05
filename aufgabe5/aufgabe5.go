/* AUFGABE: Structs
 * Erreichbare Punkte:
 */
package aufgabe5

// AUFGABENSTELLUNG:
// Fügt ein Element zu dict hinzu,
// das das Wortpaar (de,en) enthält.
// Falls dict schon einen Eintrag für de enthält,
// soll dessen en ersetzt werden.
func (dict *Dictionary) SetEntry(de, en string) {
	for i := range dict.Entries {
		if dict.Entries[i].De == de {
			dict.Entries[i].En = en
			return
		}
	}
	dict.Entries = append(dict.Entries, Entry{de, en})
}
