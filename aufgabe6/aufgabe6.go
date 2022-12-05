/* AUFGABE: Structs
 * Erreichbare Punkte: 20
 */
package aufgabe6

import (
	"fmt"
)

// AUFGABENSTELLUNG:
// Im Folgenden ist ein Datentyp für die Einträge eines Wörterbuchs gegeben.
// Ändern Sie diesen Datentyp so ab, dass er anstelle eines einzelnen englischen
// Worts eine Liste von Wörtern enthält.
// Passen Sie die Funktionen NewEntry(), Translations() und AddTranslation() an.

// Datentyp für Einträge eines Wörterbuchs.
type Entry struct {
	De string
	En string
}

// Liefert einen neuen Eintrag.
func NewEntry(de, en string) Entry {
	return Entry{de, en}
}

// Liefert eine String-Repräsentation eines Eintrags.
func (entry Entry) String() string {
	return fmt.Sprintf("%s : %v", entry.De, entry.Translations())
}

// Liefert einen String mit allen englischen Wörtern aus entry.
// Die einzelnen Wörter sollen mit Kommata getrennt sein.
func (entry Entry) Translations() string {
	return entry.En
}

// Fügt eine neue Übersetzung zu entry hinzu.
func (entry *Entry) AddTranslation(newEn string) {
	entry.En = newEn
}
