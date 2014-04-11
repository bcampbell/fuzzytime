package fuzzytime

import (
	"errors"
)

// Span represents the range [Begin,End)
type Span struct {
	Begin int
	End   int
}

var DefaultContext Context = Context{
	DateResolver: func(a, b, c int) (Date, error) {
		return Date{}, errors.New("ambiguous date")
	},
	TZResolver: func(name string) (int, error) {
		// conservative resolver - reject ambiguous names
		matches := FindTimeZone(name)
		if len(matches) == 1 {
			return TZToOffset(matches[0].Offset)
		} else if len(matches) > 1 {
			return 0, errors.New("ambiguous timezone")
		} else {
			return 0, errors.New("unknown timezone")
		}
	},
}

//TODO:
// EuroCentricContext
// USCentricContext

func Extract(s string) DateTime         { return DefaultContext.Extract(s) }
func ExtractTime(s string) (Time, Span) { return DefaultContext.ExtractTime(s) }
func ExtractDate(s string) (Date, Span) { return DefaultContext.ExtractDate(s) }

// Context to help resolve ambiguous dates and timezones
type Context struct {
	// DateResolver is called when ambigous dates are encountered eg (10/11/12)
	// it should return a date or an error
	DateResolver func(a, b, c int) (Date, error)
	//TZResolver returns the offset in seconds from UTC of the named zone (eg "EST").
	TZResolver func(name string) (int, error)
}

func (ctx *Context) Extract(s string) DateTime {
	ft, span := ctx.ExtractTime(s)
	if !ft.Empty() {
		// snip the matched time out of the string
		// (hack for nasty case where an hour can look like a 2-digit year)
		s = s[:span.Begin] + s[span.End:]
	}

	fd, _ := ExtractDate(s)

	return DateTime{fd, ft}
}

// DefaultResolveTimeZone returns the offset of the named timezone.
// If the name is unknown or if there are multiple zones of that
// name then an error is returned.
// Note that some very common timezone names (eg EST, BST) are ambiguous
// This is the function used by DefaultContext
