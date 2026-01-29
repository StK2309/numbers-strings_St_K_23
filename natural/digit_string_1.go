package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	// TODO
	var word string
	switch digit {
	case 1:
		word = "ein"
	case 2:
		word = "zwei"
	case 3:
		word = "drei"
	case 4:
		word = "vier"
	case 5:
		word = "fünf"
	case 6:
		word = "sechs"
	case 7:
		word = "sieben"
	case 8:
		word = "acht"
	case 9:
		word = "neun"
	default:
		return ""
	}
	return word + "und"
}

func DigitString11(digit int) string {
	digits := []string{"", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	if digit < 0 || digit > 9 {
		return ""
	}
	if digit == 0 {
		return ""
	}
	return digits[digit]
}
