package natural

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	// TODO
	var word string
	switch digit {
	case 1:
		word = "zwanzig"
	case 2:
		word = "dreißig"
	case 3:
		word = "vierzig"
	case 4:
		word = "fünfzig"
	case 5:
		word = "sechzig"
	case 6:
		word = "siebzig"
	case 7:
		word = "achtzig"
	case 8:
		word = "neunzig"
	default:
		return ""

	}
	return word
}
