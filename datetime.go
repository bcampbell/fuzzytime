package fuzzytime

//  DateTime represents a set of fields for date and time, any of which may
// be unset. The default initialisation (ie DateTime{}) produces an empty
// DateTime (that is, with all fields unset)
type DateTime struct {
	Date
	Time
}

// Equals returns true if dates and times match
func (dt *DateTime) Equals(other *DateTime) bool {
	return dt.Date.Equals(&other.Date) && dt.Time.Equals(&other.Time)
}

// String returns "YYYY-MM-DD hh:mm:ss tz" with question marks in place of
// any missing values (except for timezone, which will be blank if missing)
func (dt *DateTime) String() string {
	return dt.Date.String() + " " + dt.Time.String()
}

// Empty tests if datetime is blank (ie all fields unset)
func (dt *DateTime) Empty() bool {
	return dt.Time.Empty() && dt.Date.Empty()
}

// IsoFormat returns "YYYY-MM-DDTHH:MM:SS+TZ" if both date and time defined,
// or just "YYYY-MM-DD" if only date defined.
// returns an error if any part of date is missing, or if time is non-empty,
// but missing hours or minutes (seconds will be assumed to be zero if unset)
func (dt *DateTime) IsoFormat() (string, error) {
	d, derr := dt.Date.IsoFormat()
	if derr != nil {
		return "", derr
	}

	if dt.Time.Empty() {
		// just the date.
		return d, derr
	}

	t, terr := dt.Date.IsoFormat()
	if terr != nil {
		return "", terr
	}

	return d + "T" + t, nil
}
