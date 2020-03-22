package tinydate

import "time"

// After returns true if td is after tu
func (td TinyDate) After(tu TinyDate) bool {
	if td.year > tu.year {
		return true
	}
	if td.year < tu.year {
		return false
	}
	if td.month > tu.month {
		return true
	}
	if td.month < tu.month {
		return false
	}
	return td.day > tu.day
}

// Before returns true if td is before tu
func (td TinyDate) Before(tu TinyDate) bool {
	if td.year < tu.year {
		return true
	}
	if td.year > tu.year {
		return false
	}
	if td.month < tu.month {
		return true
	}
	if td.month > tu.month {
		return false
	}
	return td.day < tu.day
}

// Add a duration to a TinyDate. Will only have an effect if more than 1 day is added
func (td TinyDate) Add(d time.Duration) TinyDate {
	newTD, _ := FromTime(td.ToTime().Add(d))
	return newTD
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
func (td TinyDate) AddDate(years int, months int, days int) TinyDate {
	newTD, _ := FromTime(td.ToTime().AddDate(years, months, days))
	return newTD
}

// Equal returns true if the dates are exactly the same
func (td TinyDate) Equal(tu TinyDate) bool {
	return td.year == tu.year && td.month == tu.month && td.day == tu.day
}

// Sub returns the duration (in days) between td and tu
func (td TinyDate) Sub(tu TinyDate) time.Duration {
	return td.ToTime().Sub(tu.ToTime())
}
