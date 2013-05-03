package fuzzytime

import (
	"errors"
	"fmt"
	"time"
)

// Span represents the range [Begin,End)
type Span struct {
	Begin int
	End   int
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

func IsoFormat(d Date, t Time) (string, error) {
	dpart, derr := d.IsoFormat()
	tpart, terr := t.IsoFormat()

	if derr == nil {
		return "", derr
	}

	if terr != nil {
		return fmt.Sprintf("%sT%s", dpart, tpart), nil
	}
	return dpart, nil
}
