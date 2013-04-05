package fuzzytime

import (
	//	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type FuzzyDate struct {
	Year        int
	Month       time.Month
	Day         int
	Hour        int
	Minute      int
	Second      int
	Microsecond int
	TZName      string
}

func ExtractTime(s string) (time.Time, error) {

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
					fd.Month = time.Month(month)
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
				fd.Day = day
			case "hour":
				hour, err := strconv.Atoi(sub)
				if err != nil {
					break
				}
				fd.Hour = hour
			case "minute":
				minute, err := strconv.Atoi(sub)
				if err != nil {
					break
				}
				fd.Minute = minute
			case "second":
				second, err := strconv.Atoi(sub)
				if err != nil {
				}
				fd.Second = second
			}
			//			fmt.Printf("%d ('%s'): '%s' (%d-%d)\n", i, name, sub, start, end)
			//			fmt.Printf("%v\n",fd)
		}

		// got enough?
		if fd.Year != 0 && fd.Month != 0 && fd.Day != 0 {
			return time.Date(fd.Year, fd.Month, fd.Day, 0, 0, 0, 0, time.UTC), nil
		}
	}

	return time.Time{}, nil
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
}

/*

    # "2010-04-02"
    r'(?P<year>\d{4})-(?P<month>\d{1,2})-(?P<day>\d{1,2})',
    # "2007/03/18"
    r'(?P<year>\d{4})/(?P<month>\d{1,2})/(?P<day>\d{1,2})',
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
