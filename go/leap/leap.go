// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear determins if a given year is a leap year
func IsLeapYear(year int) bool {
	// on every year that is evenly divisible by 4
	// except every year that is evenly divisible by 100
	// unless the year is also evenly divisible by 400
	leapYear := false
	if year%400 == 0 {
		leapYear = true
	} else if year%100 == 0 {
		leapYear = false
	} else if year%4 == 0 {
		leapYear = true
	}
	return leapYear
}
