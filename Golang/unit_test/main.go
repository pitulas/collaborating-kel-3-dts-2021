package unit_test

// jika nilai <= 40, maka nilai C
// jika nilai > 40 dan <= 70, maka nilai B
// jika nilai > 70, maka nilai A

func status(n int) string {
	if n <= 40 {
		return "C"
	} else if n > 40 && n <= 70 {
		return "B"
	} else {
		return "A"
	}
}
