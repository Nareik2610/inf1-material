package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion FilterVowels.
 * ERREICHBARE PUNKTE: 10
 */

// FilterVowels erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Vokale entfernt werden.
// Anmerkung: Es müssen nur Kleinbuchstaben beachtet werden.
func FilterVowels(s string) string {
	result := ""
	// TODO
	//aeiou
	if len(s) == 0 {
		return result
	}
	if s[0] != 'a' && s[0] != 'e' && s[0] != 'i' && s[0] != 'o' && s[0] != 'u' {
		return string(s[0]) + FilterVowels(s[1:])

	}

	return FilterVowels(s[1:])
}
