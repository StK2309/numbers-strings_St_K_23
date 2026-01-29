package natural

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {
	// TODO
	if n < 0 || n > 999999 {
		return ""
	}
	if n == 0 {
		return "null"
	}
	thousands := n / 1000
	hundreds := n % 1000

	if thousands == 0 {
		return NumberString3Digits(hundreds)
	}
	if thousands == 1 {
		if hundreds == 0 {
			return "eintausend"
		}
		return "eintausend" + NumberString3Digits(hundreds)
	}
	if hundreds == 0 {
		return NumberString3Digits(thousands) + "tausend"
	}
	return NumberString3Digits(thousands) + "tausend" + NumberString3Digits(hundreds)
}
