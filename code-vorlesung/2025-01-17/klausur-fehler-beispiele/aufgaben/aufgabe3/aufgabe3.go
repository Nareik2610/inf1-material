package aufgabe3

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ElementDiffs.
 * ERREICHBARE PUNKTE: 10
 * RANDBEDINGUNG: Die Funktion muss rekursiv sein!
 */

// ElementDiffs erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils die Differenz der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementDiffs(l1, l2 []int) []int {
	// TODO

	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	return append([]int{l1[0] - l2[0]}, ElementDiffs(l1[1:], l2[1:])...)
}

// BEWERTUNG: 10 Punkte
