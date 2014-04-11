package fuzzytime

import (
	"fmt"
)

// A Date represents a year/month/day set where any of the three may be
// unset.
// default initialisation (ie Date{}) is a valid Date, with no fields set.
type Date struct {
	year, month, day int // internally, we'll say 0=undefined
}

// Year returns the year (result undefined if field unset)
func (d *Date) Year() int { return d.year }

// Month returns the month (result undefined if field unset)
func (d *Date) Month() int { return d.month }

// Day returns the day (result undefined if field unset)
func (d *Date) Day() int { return d.day }

func (d *Date) SetYear(year int)   { d.year = year }
func (d *Date) SetMonth(month int) { d.month = month }
func (d *Date) SetDay(day int)     { d.day = day }

func (d *Date) HasYear() bool  { return d.year != 0 }
func (d *Date) HasMonth() bool { return d.month != 0 }
func (d *Date) HasDay() bool   { return d.day != 0 }

// Equals returns true if dates match
func (d *Date) Equals(other *Date) bool {
	// TODO: should check if fields are set before comparing
	if d.year == other.year && d.month == other.month && d.day == other.day {
		return true
	}
	return false
}

func (d *Date) Conflicts(other *Date) bool {
	if d.HasYear() && other.HasYear() && d.Year() != other.Year() {
		return true
	}
	if d.HasMonth() && other.HasMonth() && d.Month() != other.Month() {
		return true
	}
	if d.HasDay() && other.HasDay() && d.Day() != other.Day() {
		return true
	}
	return false
}

// Empty tests if date is blank (ie all fields unset)
func (d *Date) Empty() bool {
	if d.HasYear() || d.HasMonth() || d.HasDay() {
		return false
	}
	return true
}

// String returns "YYYY-MM-DD" with question marks in place of
// any missing values
func (d *Date) String() string {
	var year, month, day = "????", "??", "??"
	if d.HasYear() {
		year = fmt.Sprintf("%04d", d.Year())
	}
	if d.HasMonth() {
		month = fmt.Sprintf("%02d", d.Month())
	}
	if d.HasDay() {
		day = fmt.Sprintf("%02d", d.Day())
	}

	return year + "-" + month + "-" + day
}

// ISOFormat returns "YYYY-MM-DD", "YYYY-MM" or "YYYY" depending on which
// fields are available (or "" if year is missing).
func (d *Date) ISOFormat() string {
	if d.HasYear() {
		if d.HasMonth() {
			if d.HasDay() {
				return fmt.Sprintf("%04d-%02d-%02d", d.Year(), d.Month(), d.Day())
			} else {

				return fmt.Sprintf("%04d-%02d", d.Year(), d.Month())
			}
		} else {
			return fmt.Sprintf("%04d", d.Year())
		}
	}
	return ""
}

// NewDate creates a Date with all fields set
func NewDate(y, m, d int) *Date {
	return &Date{y, m, d}
}
