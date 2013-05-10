package fuzzytime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

// IsoFormat returns "YYYY-MM-DD" (or error if fields missing)
func (d *Date) IsoFormat() (string, error) {
	// require full date
	if !(d.HasYear() && d.HasMonth() && d.HasDay()) {
		return "", errors.New("date information missing")
	}

	return fmt.Sprintf("%04d-%02d-%02d", d.Year(), d.Month(), d.Day()), nil
}

// NewDate creates a Date with all fields set
func NewDate(y, m, d int) *Date {
	return &Date{y, m, d}
}

// dateCracker is a set of regexps for various date formats
// order is important(ish) - want to match as much of the string as we can
var dateCrackers = []*regexp.Regexp{
	//"Tuesday 16 December 2008"
	//"Tue 29 Jan 08"
	//"Monday, 22 October 2007"
	//"Tuesday, 21st January, 2003"
	regexp.MustCompile(`(?i)(?P<dayname>\w{3,})[.,\s]+(?P<day>\d{1,2})(?:st|nd|rd|th)?\s+(?P<month>\w{3,})[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "Friday    August    11, 2006"
	// "Tuesday October 14 2008"
	// "Thursday August 21 2008"
	// "Monday, May. 17, 2010"
	regexp.MustCompile(`(?i)(?P<dayname>\w{3,})[.,\s]+(?P<month>\w{3,})[.,\s]+(?P<day>\d{1,2})(?:st|nd|rd|th)?[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "9 Sep 2009", "09 Sep, 2009", "01 May 10"
	// "23rd November 2007", "22nd May 2008"
	regexp.MustCompile(`(?i)(?P<day>\d{1,2})(?:st|nd|rd|th)?\s+(?P<month>\w{3,})[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "Mar 3, 2007", "Jul 21, 08", "May 25 2010", "May 25th 2010", "February 10 2008"
	regexp.MustCompile(`(?i)(?P<month>\w{3,})[.,\s]+(?P<day>\d{1,2})(?:st|nd|rd|th)?[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "2010-04-02"
	regexp.MustCompile(`(?i)(?P<year>\d{4})-(?P<month>\d{1,2})-(?P<day>\d{1,2})`),

	// "2007/03/18"
	regexp.MustCompile(`(?i)(?P<year>\d{4})/(?P<month>\d{1,2})/(?P<day>\d{1,2})`),

	// "22/02/2008"
	// "22-02-2008"
	// "22.02.2008"
	regexp.MustCompile(`(?i)(?P<day>\d{1,2})[/.-](?P<month>\d{1,2})[/.-](?P<year>\d{4})`),
	// "09-Apr-2007", "09-Apr-07"
	regexp.MustCompile(`(?i)(?P<day>\d{1,2})-(?P<month>\w{3,})-(?P<year>(\d{4})|(\d{2}))`),
}

/*
    Ambiguous formats...
    # dd-mm-yy
    r'(?P<day>\d{1,2})-(?P<month>\d{1,2})-(?P<year>\d{2})',
    # dd/mm/yy
    r'(?P<day>\d{1,2})/(?P<month>\d{1,2})/(?P<year>\d{2})',
    # dd.mm.yy
    r'(?P<day>\d{1,2})[.](?P<month>\d{1,2})[.](?P<year>\d{2})',

    also others, eg:  japan uses yy/mm/dd

    # TODO:
    # year/month only

    # "May/June 2011" (common for publications) - just use second month
    r'(?P<cruftmonth>\w{3,})/(?P<month>\w{3,})\s+(?P<year>\d{4})',

    # "May 2011"
    r'(?P<month>\w{3,})\s+(?P<year>\d{4})',
]

*/

// ExtractDate tries to parse a date from a string.
// It returns a Date and Span indicating which part of string matched.
func ExtractDate(s string) (fd Date, span Span) {

	for _, pat := range dateCrackers {
		names := pat.SubexpNames()
		matchSpans := pat.FindStringSubmatchIndex(s)
		if matchSpans == nil {
			continue
		}

		for i, name := range names {
			start, end := matchSpans[i*2], matchSpans[(i*2)+1]
			var sub string = ""
			if start >= 0 && end >= 0 {
				sub = strings.ToLower(s[start:end])
			}

			switch name {
			case "year":
				year, e := strconv.Atoi(sub)
				if e == nil {
					if year < 100 {
						year += 2000
					}
					fd.SetYear(year)
				} else {
					break
				}
			case "month":
				month, e := strconv.Atoi(sub)
				if e == nil {
					// it was a number
					if month < 1 || month > 12 {
						break // month out of range
					}
					fd.SetMonth(month)
				} else {
					// try month name
					month, ok := monthLookup[sub]
					if !ok {
						break // nope.
					}
					fd.SetMonth(month)
				}
			case "cruftmonth":
				// special case to handle "Jan/Feb 2010"...
				// we'll make sure the first month is valid, then ignore it
				_, ok := monthLookup[sub]
				if !ok {
					break
				}
			case "day":
				day, e := strconv.Atoi(sub)
				if e != nil {
					break
				}
				if day < 1 || day > 31 {
					break
				}
				fd.SetDay(day)
			}
		}

		// got enough?
		if fd.HasYear() && fd.HasMonth() && fd.HasDay() {
			span.Begin, span.End = matchSpans[0], matchSpans[1]
			return
		}
	}

	// nothing. Just return an empty date and span
	fd = Date{}
	span = Span{}
	return
}
