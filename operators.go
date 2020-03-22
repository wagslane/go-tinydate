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
	t := td.ToTime()
	t = t.Add(d)
	newTD, _ := FromTime(t)
	return newTD
}

// AddDate adds a date
func (td TinyDate) AddDate(years int, months int, days int) TinyDate {
	t := td.ToTime()
	t = t.AddDate(years, months, days)
	newTD, _ := FromTime(t)
	return newTD
}

// Equal returns true if the dates are exactly the same
func (td TinyDate) Equal(tu TinyDate) bool {
	return td.year == tu.year && td.month == tu.month && td.day == tu.day
}

// Sub returns the duration (in days) between td and tu
func (td TinyDate) Sub(tu TinyDate) time.Duration {
	t := td.ToTime()
	u := tu.ToTime()
	return t.Sub(u)
}
