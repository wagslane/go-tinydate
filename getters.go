package tinydate

import "time"

// Year returns the year as the readable numerical value e.g. 2020
func (td TinyDate) Year() int {
	return int(td.year)
}

// Month returns the month of the date 1-12
func (td TinyDate) Month() time.Month {
	return time.Month(td.month + 1)
}

// Day retruns the day of the month as an integer
func (td TinyDate) Day() int {
	return int(td.day + 1)
}

// String returns the month formatted as a readable string
func (td TinyDate) String() string {
	return td.Format(iso8601Date)
}

// IsZero returns true if the date represents the first day,
// January 01, 0000
func (td TinyDate) IsZero() bool {
	return td.year == 0 && td.month == 0 && td.day == 0
}

// Unix returns the date as a Unix timestamp
// where the precision is only to the day. Hours, minutes, and seconds
// are just padded zeros
func (td TinyDate) Unix() int64 {
	t := td.ToTime()
	return t.Unix()
}

// UnixNano returns the date as a Unix timestamp in nanoseconds
// where the precision is only to the day. Hours, minutes, seconds and nanoseconds
// are just padded zeros
func (td TinyDate) UnixNano() int64 {
	t := td.ToTime()
	return t.UnixNano()
}

// Weekday returns the day of the week
func (td TinyDate) Weekday() time.Weekday {
	t := td.ToTime()
	return t.Weekday()
}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years, and [1,366] in leap years
func (td TinyDate) YearDay() int {
	t := td.ToTime()
	return t.YearDay()
}

// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (td TinyDate) ISOWeek() (year, week int) {
	t := td.ToTime()
	return t.ISOWeek()
}

// Date returns the readable numerical values of the date
func (td TinyDate) Date() (year int, month time.Month, day int) {
	return int(td.year), time.Month(td.Month()), int(td.day)
}
