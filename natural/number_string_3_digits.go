package natural

var x int
var y int
var z int

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	// TODO
	if n < 0 || n > 999 {
		return ""
	}
	if n == 0 {
		return "null"
	}
	x = n / 100
	y = (n % 100) / 10
	z = n % 10

	if y == 0 { //1 bis 10 ohne Und
		return DigitString11(z)
	}

	if y == 1 { //sonderfälle
		switch z {
		case 0:
			return DigitString100(x) + "zehn"
		case 1:
			return DigitString100(x) + "elf"
		case 2:
			return DigitString100(x) + "zwölf"
		case 3:
			return DigitString100(x) + "dreizehn"
		case 4:
			return DigitString100(x) + "vierzehn"
		case 5:
			return DigitString100(x) + "fünfzehn"
		case 6:
			return DigitString100(x) + "sechzehn"
		case 7:
			return DigitString100(x) + "siebzehn"
		case 8:
			return DigitString100(x) + "achtzehn"
		case 9:
			return DigitString100(x) + "neunzehn"
		}
	}

	return DigitString100(x) + DigitString1(z) + DigitString10(y)
}
