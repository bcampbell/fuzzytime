package fuzzytime

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
func (t *Time) TZ() string  { return t.tzName }

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
func (t *Time) String() string {
	var hour, minute, second, tz = "????", "??", "??", ""
	if t.HasHour() {
		hour = fmt.Sprintf("%02d", t.Hour())
	}
	if t.HasMinute() {
		minute = fmt.Sprintf("%02d", t.Minute())
	}
	if t.HasSecond() {
		second = fmt.Sprintf("%02d", t.Second())
	}
	if t.HasTZ() {
		tz = " " + t.TZ()
	}
	return hour + ":" + minute + ":" + second + tz
}

func (t *Time) Empty() bool {
	return t.set == 0
}

// match one of:
//  named timezone (eg, BST, NZDT etc)
//  Z
//  +hh:mm, +hhmm, or +hh
//  -hh:mm, -hhmm, or -hh
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
				if _, ok := tzTable[strings.ToUpper(sub)]; ok {
					tzName = strings.ToUpper(sub)
				} else {
					// doesn't look like a timezone after all
					break
				}
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
