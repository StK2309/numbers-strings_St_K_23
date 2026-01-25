package natural

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	// TODO
	var word string
	switch digit {
	case 1:
		word = "einhundert"
	case 2:
		word = "zweihundert"
	case 3:
		word = "dreihundert"
	case 4:
		word = "vierhundert"
	case 5:
		word = "fünfhundert"
	case 6:
		word = "sechshundert"
	case 7:
		word = "siebenhundert"
	case 8:
		word = "achthundert"
	case 9:
		word = "neunhundert"
	}
	return word
}
