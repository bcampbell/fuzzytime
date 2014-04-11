package fuzzytime

import (
	"testing"
)

func TestDateTimes(t *testing.T) {
	// TODO: add some more tests with numeric timezones
	// TODO: use DateTime.String()-style strings for expected results
	testData := []struct {
		in       string
		expected string
	}{

		{"2010-04-02T12:35:44+00:00", "2010-04-02T12:35:44Z"},       // iso 8601
		{"2008-03-10 13:21:36 GMT", "2008-03-10T13:21:36Z"},         //
		{"9 Sep 2009 12.33", "2009-09-09T12:33"},                    //(heraldscotland blogs)
		{"May 25 2010 3:34PM", "2010-05-25T15:34"},                  //(thetimes.co.uk)
		{"Thursday August 21 2008 10:42 am", "2008-08-21T10:42"},    //(guardian blogs in their new cms)
		{"Tuesday 16 December 2008 16.23 GMT", "2008-12-16T16:23Z"}, //(Guardian blogs in their new cms)
		{"3:19pm on Tue 29 Jan 08", "2008-01-29T15:19"},             //(herald blogs)
		{"2007/03/18 10:59:02", "2007-03-18T10:59:02"},

		{"Mar 3, 2007 12:00 AM", "2007-03-03T00:00"},
		{"Jul 21, 08 10:00 AM", "2008-07-21T10:00"},                   //(mirror blogs)
		{"09-Apr-2007 00:00", "2007-04-09T00:00"},                     //(times, sundaytimes)
		{"4:48PM GMT 22/02/2008", "2008-02-22T16:48Z"},                //(telegraph html articles)
		{"09-Apr-07 00:00", "2007-04-09T00:00"},                       //(scotsman)
		{"Friday    August    11, 2006", "2006-08-11"},                //(express, guardian/observer)
		{"20:12pm 23rd November 2007", "2007-11-23T20:12"},            //(dailymail)
		{"2:42 PM on 22nd May 2008", "2008-05-22T14:42"},              //(dailymail)
		{"February 10 2008 22:05", "2008-02-10T22:05"},                //(ft)
		{"Feb 2, 2009 at 17:01:09", "2009-02-02T17:01:09"},            //(telegraph blogs)
		{"18 Oct 07, 04:50 PM", "2007-10-18T16:50"},                   //(BBC blogs)
		{"02 August 2007  1:21 PM", "2007-08-02T13:21"},               //(Daily Mail blogs)
		{"October 22, 2007  5:31 PM", "2007-10-22T17:31"},             //(old Guardian blogs, ft blogs)
		{"October 15, 2007", "2007-10-15"},                            //(Times blogs)
		{"February 12 2008", "2008-02-12"},                            //(Herald)
		{"Monday, 22 October 2007", "2007-10-22"},                     //(Independent blogs, Sun (page date))
		{"22 October 2007", "2007-10-22"},                             //(Sky News blogs)
		{"11 Dec 2007", "2007-12-11"},                                 //(Sun (article date))
		{"12 February 2008", "2008-02-12"},                            //(scotsman)
		{"Tuesday, 21 January, 2003, 15:29 GMT", "2003-01-21T15:29Z"}, //(historical bbcnews)
		{"2003/01/21 15:29:49", "2003-01-21T15:29:49"},                //(historical bbcnews (meta tag))
		{"2010-07-01", "2010-07-01"},
		{"2010/07/01", "2010-07-01"},
		{"Feb 20th, 2000", "2000-02-20"},
		{"Monday, May. 17, 2010", "2010-05-17"}, // (time.com)

		{"APRIL 10, 2014", "2014-04-10"}, // nytimes.com
		{"30.12.2011", "2011-12-30"},

		{"10 ABR 2014 - 20:36 CET", "2014-04-10T20:36+01:00"},      // elpais.com
		{"9:11 p.m. EDT April 10, 2014", "2014-04-10T21:11-04:00"}, // usatoday.com
		// with trailing text
		{"September, 26th 2011 by Christo Hall", "2011-09-26"}, // (www.thenewwolf.co.uk)

		// some more obscure cases...
		{"May 2008", "2008-05"},

		// BST is ambiguous
		//{"Tuesday October 14 2008 00.01 BST", "2008-10-14T00:01+01:00"}, //(Guardian blogs in their new cms)
		//{"26 May 2007, 02:10:36 BST", "2007-05-26T02:10:36+01:00"},                      //(newsoftheworld)
		//{"2:43pm BST 16/04/2007", "2007-04-16T14:43+01:00"},         //(telegraph, after munging)
		//{"Monday 30 July 2012 08.38 BST", *"2012-7-30T8:38:0+01:00")}, // (guardian.co.uk)

		// NOTE: this is a tricky one where hour can get picked up as year if not careful!
		{"Thu Aug 25 10:46:55 GMT 2011", "2011-08-25T10:46:55Z"}, // (www.yorkshireeveningpost.co.uk)

		// Other possible formats to support:
		// http://en.wikipedia.org/wiki/Date_and_time_notation_in_the_United_States#Date-time_group
		//{"091630Z JUL 11", "2011-07T09:16:30Z"

		// Ones that should fail

		// time or date?
		{"10.12", ""},
		// ambiguous (at least with the default date resolver)
		{"03/09/2007", ""}, //(Sky News blogs, mirror)
		{"03/09/12", ""},

		// ambiguous format, but with values that provide enough info
		// {"25/11/2004","2004-11-25"}
	}

	for _, dat := range testData {
		dt := Extract(dat.in)

		got := dt.ISOFormat()

		if got != dat.expected {
			t.Errorf("Extract(%s): expected %s, but got %s", dat.in, dat.expected, got)
		}
	}

}

// Test timezone parsing
func TestParseTimeZone(t *testing.T) {
	/*
		testData := []struct {
			in       string
			expected int
		}{
			{"Z", 0},
			{"+0100", 1 * 60 * 60},
			{"-0430", -(4*60*60 + 30*60)},
			{"NZDT", 13 * 60 * 60},
		}

		// TODO
	*/
}

func TestTZToOffset(t *testing.T) {
	testData := []struct {
		in       string
		expected int
	}{
		{"Z", 0},
		{"+00:00", 0},
		{"-00:00", 0},
		{"+0000", 0},
		{"+1000", 10 * 60 * 60},
		{"-01:35", -(1*60*60 + 35*60)},
	}
	for _, dat := range testData {
		got, err := TZToOffset(dat.in)
		if err != nil {
			t.Errorf("TZToOffset(%s) error: %s", dat.in, err)
		} else if got != dat.expected {
			t.Errorf("TZToOffset(%s): expected '%d' but got '%d'", dat.in, dat.expected, got)
		}
	}
}

func TestOffsetToTZ(t *testing.T) {

	testData := []struct {
		in       int
		expected string
	}{
		{0, "Z"},
		{30 * 60, "+00:30"},
		{-45 * 60, "-00:45"},
		{(10 * 60 * 60) + (0 * 60), "+10:00"},
		{-((10 * 60 * 60) + (15 * 60)), "-10:15"},
	}
	for _, dat := range testData {
		got := OffsetToTZ(dat.in)
		if got != dat.expected {
			t.Errorf("OffsetToTZ(%d): expected '%s' but got '%s'", dat.in, dat.expected, got)
		}
	}
}
