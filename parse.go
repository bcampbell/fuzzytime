package fuzzytime

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func (d *Date) Empty() bool {
	if d.HasYear() || d.HasMonth() || d.HasDay() {
		return false
	}
	return true
}

// NewDate creates a Date with all fields set
func NewDate(y, m, d int) *Date {
	return &Date{y, m, d}
}

const (
	hourFlag   int = 0x01
	minuteFlag int = 0x02
	secondFlag int = 0x04
	tzFlag     int = 0x08
)

// Time represents a set of time fields, any of which may be unset.
// The default initialisation (ie Time{}) produces a Time with all fields unset.
type Time struct {
	set    int // flags to show which fields are set
	hour   int
	minute int
	second int
	tzName string // TODO: still mulling timezones over
}

// Hour returns the hour (result undefined if field unset)
func (t *Time) Hour() int { return t.hour }

// Minute returns the minute (result undefined if field unset)
func (t *Time) Minute() int { return t.minute }

// Second returns the second (result undefined if field unset)
func (t *Time) Second() int { return t.second }

func (t *Time) SetHour(hour int)     { t.hour = hour }
func (t *Time) SetMinute(minute int) { t.minute = minute }
func (t *Time) SetSecond(second int) { t.second = second }
func (t *Time) HasHour() bool        { return (t.set & hourFlag) != 0 }
func (t *Time) HasMinute() bool      { return (t.set & minuteFlag) != 0 }
func (t *Time) HasSecond() bool      { return (t.set & secondFlag) != 0 }
func (t *Time) HasTZ() bool          { return (t.set & tzFlag) != 0 }
func NewTime(h, m, s int, tz string) *Time {
	return &Time{hourFlag | minuteFlag | secondFlag | tzFlag, h, m, s, tz}
}

// Equals returns true if times match
func (t *Time) Equals(other *Time) bool {
	if t.set != other.set {
		return false
	}
	if t.HasHour() && t.hour != other.hour {
		return false
	}
	if t.HasMinute() && t.minute != other.minute {
		return false
	}
	if t.HasSecond() && t.second != other.second {
		return false
	}
	if t.HasTZ() && t.tzName != other.tzName {
		return false
	}
	return true
}

func (t *Time) Empty() bool {
	return t.set == 0
}

// Span represents the range [Begin,End)
type Span struct {
	Begin int
	End   int
}

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
    # dd-mm-yy
    r'(?P<day>\d{1,2})-(?P<month>\d{1,2})-(?P<year>\d{2})',
    # dd/mm/yy
    r'(?P<day>\d{1,2})/(?P<month>\d{1,2})/(?P<year>\d{2})',
    # dd.mm.yy
    r'(?P<day>\d{1,2})[.](?P<month>\d{1,2})[.](?P<year>\d{2})',

    # TODO:
    # mm/dd/yy
    # dd.mm.yy
    # etc...
    # YYYYMMDD

    # TODO:
    # year/month only

    # "May/June 2011" (common for publications) - just use second month
    r'(?P<cruftmonth>\w{3,})/(?P<month>\w{3,})\s+(?P<year>\d{4})',

    # "May 2011"
    r'(?P<month>\w{3,})\s+(?P<year>\d{4})',
]

*/

// "BST" ,"+02:00", "+02"
var tzPat string = `(?i)(?P<tz>Z|[A-Z]{2,10}|(([-+])(\d{2})((:?)(\d{2}))?))`
var ampmPat string = `(?i)(?:(?P<am>am)|(?P<pm>pm))`

var timeCrackers = []*regexp.Regexp{
	// "4:48PM GMT"
	regexp.MustCompile(`(?i)(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*` + ampmPat + `\s*` + tzPat),

	// "3:34PM"
	// "10:42 am"
	regexp.MustCompile(`(?i)(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*` + ampmPat),

	// "13:21:36 GMT"
	// "15:29 GMT"
	// "12:35:44+00:00"
	// "00.01 BST"
	regexp.MustCompile(`(?i)(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*` + tzPat),

	// "12.33"
	// "14:21"
	// TODO: BUG: this'll also pick up time from "30.25.2011"!
	regexp.MustCompile(`(?i)(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*`),

	// TODO: add support for microseconds?
}

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

//return time.Date(fd.Year, fd.Month, fd.Day, 0, 0, 0, 0, time.UTC), nil

// ExtractTime tries to parse a time from a string.
// It returns a Time and a Span indicating which part of string matched
func ExtractTime(s string) (Time, Span) {
	for _, pat := range timeCrackers {
		names := pat.SubexpNames()
		matchSpans := pat.FindStringSubmatchIndex(s)
		if matchSpans == nil {
			continue
		}

		var hour, minute, second int = -1, -1, -1
		var am, pm bool = false, false
		var tzName string = ""
		var err error
		for i, name := range names {
			start, end := matchSpans[i*2], matchSpans[(i*2)+1]
			if start == end {
				continue
			}
			var sub string = ""
			if start >= 0 && end >= 0 {
				sub = strings.ToLower(s[start:end])
			}

			switch name {
			case "hour":
				hour, err = strconv.Atoi(sub)
				if err != nil {
					break
				}
			case "min":
				minute, err = strconv.Atoi(sub)
				if err != nil {
					break
				}
			case "sec":
				second, err = strconv.Atoi(sub)
				if err != nil {
					break
				}
			case "am":
				am = true
			case "pm":
				pm = true
			case "tz":
				tzName = strings.ToUpper(sub)
			}

		}

		// got enough?
		if hour >= 0 && minute >= 0 {
			// ok if seconds are missing - just assume zero
			if second == -1 {
				second = 0
			}
			if pm && (hour >= 1) && (hour <= 11) {
				hour += 12
			}
			if am && (hour == 12) {
				hour -= 12
			}
			var ft = *NewTime(hour, minute, second, tzName)
			var span = Span{matchSpans[0], matchSpans[1]}
			return ft, span
		}
	}

	// nothing. Just return an empty time and span
	return Time{}, Span{}
}

// TODO: return sensible span(s)!
func Extract(s string) (Date, Time) {
	ft, span := ExtractTime(s)
	if !ft.Empty() {
		// snip the matched time out of the string
		// (hack for nasty case where an hour can look like a 2-digit year)
		s = s[:span.Begin] + s[span.End:]
	}

	fd, _ := ExtractDate(s)

	return fd, ft
}

func Parse(s string) (time.Time, error) {
	ft, span := ExtractTime(s)
	if !ft.Empty() {
		// snip the matched time out of the string
		s = s[:span.Begin] + s[span.End:]
	}

	fd, _ := ExtractDate(s)

	if !fd.Empty() {
		if !ft.Empty() {
			return time.Date(fd.Year(), time.Month(fd.Month()), fd.Day(), ft.Hour(), ft.Minute(), ft.Second(), 0, time.UTC), nil
		} else {
			// ok if time missing
			return time.Date(fd.Year(), time.Month(fd.Month()), fd.Day(), 0, 0, 0, 0, time.UTC), nil
		}
	}
	return time.Time{}, errors.New("no date found")
}
