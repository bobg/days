package days

// Delta computes the number of days between two dates, each given as
// Y, M, D. The result is positive iff the first date is earlier than
// the second, zero if they're the same date, and negative otherwise.
func Delta(y1, m1, d1, y2, m2, d2 int) int {
	if y1 > y2 {
		return -Delta(y2, m2, d2, y1, m1, d1)
	}
	if y1 == y2 && m1 > m2 {
		return -Delta(y2, m2, d2, y1, m1, d1)
	}
	if y1 == y2 && m1 == m2 && d1 > d2 {
		return -Delta(y2, m2, d2, y1, m1, d1)
	}
	// now date1 <= date2
	if y1 == y2 {
		return DeltaInYear(y1, m1, d1, m2, d2)
	}
	result := DeltaInYear(y1, m1, d1, 12, 31) + 1 // days to end of year 1
	result += DeltaInYear(y2, 1, 1, m2, d2)       // days from start of year 2
	deltaYears := y2 - y1
	if deltaYears > 1 {
		result += 365 * (deltaYears - 1)
		result += CountLeapYears(y1, y2)
	}
	return result
}

// DeltaInYear computes the number of days between two dates in the
// same year. The result is positive iff the first date is earlier
// than the second, zero if they're the same date, and negative
// otherwise.
func DeltaInYear(year, m1, d1, m2, d2 int) int {
	if m1 > m2 {
		return -DeltaInYear(year, m2, d2, m1, d1)
	}
	if m1 == m2 {
		return d2 - d1
	}
	result := MonthLen(year, m1) - d1 + 1 // days to end of month 1
	result += d2 - 1                      // days from start of month 2
	for m := m1 + 1; m < m2; m++ {
		result += MonthLen(year, m)
	}
	return result
}

// MonthLen reports the number of days in a given month.
func MonthLen(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	}
	if year%400 == 0 {
		return 29
	}
	if year%100 == 0 {
		return 28
	}
	if year%4 == 0 {
		return 29
	}
	return 28
}

// CountLeapYears counts the number of leap years between year y1 and year y2, inclusive.
func CountLeapYears(y1, y2 int) int {
	if y1 > y2 {
		return CountLeapYears(y2, y1)
	}
	y1--
	result := y2/4 - y1/4
	result -= y2/100 - y1/100
	result += y2/400 - y1/400
	return result
}
