package lenconv

// делим значение длины на 3.281
func FtToMet(f Ft) Meter {
	return Meter(f / 3.281)
}

// умножаем значение длины на 3.281
func MetToFt(m Meter) Ft {
	return Ft(m * 3.281)
}
