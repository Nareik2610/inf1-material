package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	// TODO
	pos1 := -1
	pos2 := -1
	for pos, s := range list {
		if s == first {
			pos1 = pos

		}
		if s == last {
			pos2 = pos
		}
	}
	if pos2 <= pos1 {
		return []string{}
	}

	return append(list[:pos1], list[pos2+1:]...)
}
