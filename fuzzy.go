package fuzzytime

import (
	"errors"
	"strings"
)

// Span represents the range [Begin,End)
type Span struct {
	Begin int
	End   int
}

// DefaultContext is a predefined context which bails out if timezones or dates are ambiguous
var DefaultContext Context = Context{
	DateResolver: func(a, b, c int) (Date, error) {
		return Date{}, errors.New("ambiguous date")
	},
	TZResolver: DefaultTZResolver(""),
}

// USContext is a prefefined context which opts for US timezones and mm/dd/yy dates
var USContext Context = Context{
	DateResolver: MDYResolver,
	TZResolver:   DefaultTZResolver("US"),
}

// WesternContext is a predefined context which opts for UK and US timezones
// and dd/mm/yy dates
var WesternContext Context = Context{
	DateResolver: DMYResolver,
	TZResolver:   DefaultTZResolver("GB,US"),
}

// Extract tries to parse a Date and Time from a string.
// Equivalent to DefaultContext.Extract()
func Extract(s string) DateTime { return DefaultContext.Extract(s) }

// Extract tries to parse a Time from a string.
// Equivalent to DefaultContext.ExtractTime()
func ExtractTime(s string) (Time, Span, error) { return DefaultContext.ExtractTime(s) }

// Extract tries to parse a Date from a string.
// Equivalent to DefaultContext.ExtractDate()
func ExtractDate(s string) (Date, Span, error) { return DefaultContext.ExtractDate(s) }

// Context provides helper functions to resolve ambiguous dates and timezones.
// For example, "CST" can mean China Standard Time, Central Standard Time in
// or Central Standard Time in Australia.
// Or, the date "5/2/10". It could Feburary 5th, 2010 or May 2nd 2010. Or even
// Feb 10th 2005, depending on country. Even "05/02/2010" is ambiguous.
type Context struct {
	// DateResolver is called when ambigous dates are encountered eg (10/11/12)
	// it should return a date or an error
	DateResolver func(a, b, c int) (Date, error)
	//TZResolver returns the offset in seconds from UTC of the named zone (eg "EST").
	TZResolver func(name string) (int, error)
}

// Extract tries to parse a Date and Time from a string
func (ctx *Context) Extract(s string) DateTime {
	ft, span, err := ctx.ExtractTime(s)
	if err != nil {
		return DateTime{}
	}
	if !ft.Empty() {
		// snip the matched time out of the string
		// (hack for nasty case where an hour can look like a 2-digit year)
		s = s[:span.Begin] + s[span.End:]
	}

	fd, _, _ := ctx.ExtractDate(s)

	return DateTime{fd, ft}
}

// DefaultTZResolver returns a TZResolver function which uses a list of country codes in
// preferredLocales to resolve ambigous timezones.
// For example, if you were expecting Bangladeshi times, then:
//     DefaultTZResolver("BD")
// would treat "BST" as Bangladesh Standard Time rather than British Summer Time
func DefaultTZResolver(preferredLocales string) func(name string) (int, error) {
	codes := strings.Split(strings.ToUpper(preferredLocales), ",")

	return func(name string) (int, error) {
		matches := FindTimeZone(name)
		if len(matches) == 1 {
			return TZToOffset(matches[0].Offset)
		} else if len(matches) > 1 {
			// try preferred locales in order of preference
			for _, cc := range codes {
				for _, tz := range matches {
					if strings.Contains(tz.Locale, cc) {
						return TZToOffset(tz.Offset)
					}
				}
			}
			return 0, errors.New("ambiguous timezone")
		} else {
			return 0, errors.New("unknown timezone")
		}
	}
}

// DMYResolver treats ambiguous dates as DD/MM/YY
func DMYResolver(a, b, c int) (Date, error) {
	c = ExtendYear(c)
	return *NewDate(c, b, a), nil
}

// MDYResolver treats ambiguous dates as MM/DD/YY
func MDYResolver(a, b, c int) (Date, error) {
	c = ExtendYear(c)
	return *NewDate(c, a, b), nil
}
