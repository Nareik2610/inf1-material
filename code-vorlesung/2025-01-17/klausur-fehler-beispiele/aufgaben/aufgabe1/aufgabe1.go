package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion EvenPrefix.
 * ERREICHBARE PUNKTE: 10
 */

// EvenPrefix erwartet eine Liste von Zahlen und liefert
// die längste Anfangs-Folge, die nur gerade Zahlen enthält.
// D.h. eine Liste mit allen geraden Zahlen vor der ersten ungeraden.
func EvenPrefix(list []int) []int {
	result := []int{}
	// TODO
	if len(list) == 0 {
		return result
	}
	if list[0]%2 != 0 {
		return result
	}

	return append(list[:1], EvenPrefix(list[1:])...)
}
