package fuzzytime

import (
//	"fmt"
	"regexp"
	"strconv"
	"strings"
	//	"time"
	"errors"
)

type FuzzyDate struct {
	Year      int
	Month     int
	Day       int
	DateBegin int
	DateEnd   int
}

type FuzzyTime struct {
	Hour      int
	Minute    int
	Second    int
	TZName    string
	TimeBegin int
	TimeEnd   int
}

// order is important(ish) - want to match as much of the string as we can
var dateCrackers = []*regexp.Regexp{
	//"Tuesday 16 December 2008"
	//"Tue 29 Jan 08"
	//"Monday, 22 October 2007"
	//"Tuesday, 21st January, 2003"
	regexp.MustCompile(`(?P<dayname>\w{3,})[.,\s]+(?P<day>\d{1,2})(?:st|nd|rd|th)?\s+(?P<month>\w{3,})[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "Friday    August    11, 2006"
	// "Tuesday October 14 2008"
	// "Thursday August 21 2008"
	// "Monday, May. 17, 2010"
	regexp.MustCompile(`(?P<dayname>\w{3,})[.,\s]+(?P<month>\w{3,})[.,\s]+(?P<day>\d{1,2})(?:st|nd|rd|th)?[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "9 Sep 2009", "09 Sep, 2009", "01 May 10"
	// "23rd November 2007", "22nd May 2008"
	regexp.MustCompile(`(?P<day>\d{1,2})(?:st|nd|rd|th)?\s+(?P<month>\w{3,})[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "Mar 3, 2007", "Jul 21, 08", "May 25 2010", "May 25th 2010", "February 10 2008"
	regexp.MustCompile(`(?P<month>\w{3,})[.,\s]+(?P<day>\d{1,2})(?:st|nd|rd|th)?[.,\s]+(?P<year>(\d{4})|(\d{2}))`),

	// "2010-04-02"
	regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{1,2})-(?P<day>\d{1,2})`),

	// "2007/03/18"
	regexp.MustCompile(`(?P<year>\d{4})/(?P<month>\d{1,2})/(?P<day>\d{1,2})`),
}

/*
    # "22/02/2008"
    # "22-02-2008"
    # "22.02.2008"
    r'(?P<day>\d{1,2})[/.-](?P<month>\d{1,2})[/.-](?P<year>\d{4})',
    # "09-Apr-2007", "09-Apr-07"
    r'(?P<day>\d{1,2})-(?P<month>\w{3,})-(?P<year>(\d{4})|(\d{2}))',


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
var tzPat string = `(?P<tz>Z|[A-Z]{2,10}|(([-+])(\d{2})((:?)(\d{2}))?))`
var ampmPat string = `(?:(?P<am>am)|(?P<pm>pm))`

var timeCrackers = []*regexp.Regexp{
	// "4:48PM GMT"
	regexp.MustCompile(`(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*` + ampmPat + `\s*` + tzPat),

	// "3:34PM"
	// "10:42 am"
	regexp.MustCompile(`(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*` + ampmPat),

	// "13:21:36 GMT"
	// "15:29 GMT"
	// "12:35:44+00:00"
	// "00.01 BST"
	regexp.MustCompile(`(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*` + tzPat),

	// "12.33"
	// "14:21"
	// TODO: BUG: this'll also pick up time from "30.25.2011"!
	regexp.MustCompile(`(?P<hour>\d{1,2})[:.](?P<min>\d{2})(?:[:.](?P<sec>\d{2}))?\s*`),

	// TODO: add support for microseconds?
}



// ExtractDate tries to parse a date from a string.
// It returns a FuzzyDate (and/or any error that might have occured)
func ExtractDate(s string) (FuzzyDate, error) {

	for _, pat := range dateCrackers {
		names := pat.SubexpNames()
		matchSpans := pat.FindStringSubmatchIndex(s)
		if matchSpans == nil {
			continue
		}

		var fd = FuzzyDate{}
		for i, name := range names {
			start, end := matchSpans[i*2], matchSpans[(i*2)+1]
			var sub string = ""
			if start >= 0 && end >= 0 {
				sub = strings.ToLower(s[start:end])
			}

			switch name {
			case "year":
				year, err := strconv.Atoi(sub)
				if err == nil {
					if year < 100 {
						year += 2000
					}
					fd.Year = year
				} else {
					break
				}
			case "month":
				month, err := strconv.Atoi(sub)
				if err == nil {
					// it was a number
					if month < 1 || month > 12 {
						break // month out of range
					}
					fd.Month = month
				} else {
					// try month name
					month, ok := monthLookup[sub]
					if !ok {
						break // nope.
					}
					fd.Month = month
				}
			case "cruftmonth":
				// special case to handle "Jan/Feb 2010"...
				// we'll make sure the first month is valid, then ignore it
				_, ok := monthLookup[sub]
				if !ok {
					break
				}
			case "day":
				day, err := strconv.Atoi(sub)
				if err != nil {
					break
				}
				if day < 1 || day > 31 {
					break
				}
				fd.Day = day
			}
			//			fmt.Printf("%d ('%s'): '%s' (%d-%d)\n", i, name, sub, start, end)
			//			fmt.Printf("%v\n",fd)
		}

		// got enough?
		if fd.Year != 0 && fd.Month != 0 && fd.Day != 0 {
			fd.DateBegin, fd.DateEnd = matchSpans[0], matchSpans[1]
			return fd, nil
		}
	}

	return FuzzyDate{}, errors.New("Date not found")
}

//return time.Date(fd.Year, fd.Month, fd.Day, 0, 0, 0, 0, time.UTC), nil

// ExtractTime tries to parse a time from a string.
// It returns a FuzzyTime (and/or any error that might have occured)
func ExtractTime(s string) (FuzzyTime, error) {
	for _, pat := range timeCrackers {
		names := pat.SubexpNames()
		matchSpans := pat.FindStringSubmatchIndex(s)
		if matchSpans == nil {
			continue
		}

		var hour, minute, second int = -1, -1, -1
		var am, pm bool = false, false
		var err error
		for i, name := range names {
			start, end := matchSpans[i*2], matchSpans[(i*2)+1]
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
			}

			//			fmt.Printf("%d ('%s'): '%s' (%d-%d)\n", i, name, sub, start, end)
			//			fmt.Printf("%v\n",fd)
		}

		// got enough?
		if hour >= 0 && minute >= 0 {
			// ok if seconds are missing - just assume zero
			if second == -1 {
				second = 0
			}
			if pm && hour >= 1 && hour <= 11 {
				hour += 12
			}
			if am && hour == 12 {
				hour -= 12
			}
			var ft = FuzzyTime{hour, minute, second, "", matchSpans[0], matchSpans[1]}
			return ft, nil
		}
	}

	return FuzzyTime{}, errors.New("Time not found")
}

