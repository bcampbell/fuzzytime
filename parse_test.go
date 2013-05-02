package fuzzytime

import (
	"testing"
	//	"time"
	"fmt"
)

type dtTest struct {
	input string
	date  Date
	time  Time
}

type datetimeTest struct {
	input  string
	expect string
}

var noddydttests = []dtTest{
	{"Tuesday 16 December 2008", *NewDate(2008, 12, 16), Time{}},
}

var noddydatetests = []datetimeTest{
	{"Tuesday 16 December 2008", "2008-12-16"},
	{"Friday    August    11, 2006", "2006-08-11"},
	{"23rd November 2007", "2007-11-23"},
	{"Jul 21, 08", "2008-07-21"},
	{"2010-07-01", "2010-07-01"},
	{"2010/07/01", "2010-07-01"},
	{"may 2nd 1928", "1928-05-02"},
}

var noddytimetests = []datetimeTest{
	{"4:48pm GMT", "16:48:00"},
	{"4:48pm", "16:48:00"},
}

var dttests = []datetimeTest{
	{"2010-04-02T12:35:44+00:00", "2010-04-02T12:35:44Z"}, //{iso8601, bbc blogs)
	{"2008-03-10 13:21:36 GMT", "2008-03-10T13:21:36Z"},   //{technorati api)

	{"9 Sep 2009 12.33", "2009-09-09T12:33:00Z"},                   //(heraldscotland blogs)
	{"May 25 2010 3:34PM", "2010-05-25T15:34:00Z"},                 //(thetimes.co.uk)
	{"Thursday August 21 2008 10:42 am", "2008-08-21T10:42:00Z"},   //(guardian blogs in their new cms)
	{"Tuesday October 14 2008 00.01 BST", "2008-10-13T23:01:00Z"},  //(Guardian blogs in their new cms)
	{"Tuesday 16 December 2008 16.23 GMT", "2008-12-16T16:23:00Z"}, //(Guardian blogs in their new cms)
	{"3:19pm on Tue 29 Jan 08", "2008-01-29T15:19:00Z"},            //(herald blogs)
	{"2007/03/18 10:59:02", "2007-03-18T10:59:02Z"},
	{"Mar 3, 2007 12:00 AM", "2007-03-03T00:00:00Z"},
	{"Jul 21, 08 10:00 AM", "2008-07-21T10:00:00Z"},          //(mirror blogs)
	{"09-Apr-2007 00:00", "2007-04-09T00:00:00Z"},            //(times, sundaytimes)
	{"4:48PM GMT 22/02/2008", "2008-02-22T16:48:00Z"},        //(telegraph html articles)
	{"09-Apr-07 00:00", "2007-04-09T00:00:00Z"},              //(scotsman)
	{"Friday    August    11, 2006", "2006-08-11T00:00:00Z"}, //(express, guardian/observer)
	{"26 May 2007, 02:10:36 BST", "2007-05-26T01:10:36Z"},    //(newsoftheworld)
	{"2:43pm BST 16/04/2007", "2007-04-16T13:43:00Z"},        //(telegraph, after munging)
	{"20:12pm 23rd November 2007", "2007-11-23T20:12:00Z"},   //(dailymail)
	//TODO        {"2:42 PM on 22nd May 2008", "2008-05-22T14:42:00Z"}, //(dailymail)
	{"February 10 2008 22:05", "2008-02-10T22:05:00Z"}, //(ft)
	//        {"22 Oct 2007, //(weird non-ascii characters) at(weird non-ascii characters)11:23", "2007-10-22T11:23:00Z"}, //(telegraph blogs OLD!)
	{"Feb 2, 2009 at 17:01:09", "2009-02-02T17:01:09Z"},              //(telegraph blogs)
	{"18 Oct 07, 04:50 PM", "2007-10-18T16:50:00Z"},                  //(BBC blogs)
	{"02 August 2007  1:21 PM", "2007-08-02T13:21:00Z"},              //(Daily Mail blogs)
	{"October 22, 2007  5:31 PM", "2007-10-22T17:31:00Z"},            //(old Guardian blogs, ft blogs)
	{"October 15, 2007", "2007-10-15T00:00:00Z"},                     //(Times blogs)
	{"February 12 2008", "2008-02-12T00:00:00Z"},                     //(Herald)
	{"Monday, 22 October 2007", "2007-10-22T00:00:00Z"},              //(Independent blogs, Sun (page date))
	{"22 October 2007", "2007-10-22T00:00:00Z"},                      //(Sky News blogs)
	{"11 Dec 2007", "2007-12-11T00:00:00Z"},                          //(Sun (article date))
	{"12 February 2008", "2008-02-12T00:00:00Z"},                     //(scotsman)
	{"03/09/2007", "2007-09-03T00:00:00Z"},                           //(Sky News blogs, mirror)
	{"Tuesday, 21 January, 2003, 15:29 GMT", "2003-01-21T15:29:00Z"}, //(historical bbcnews)
	{"2003/01/21 15:29:49", "2003-01-21T15:29:49Z"},                  //(historical bbcnews (meta tag))
	{"2010-07-01", "2010-07-01T00:00:00Z"},
	{"2010/07/01", "2010-07-01T00:00:00Z"},
	{"Feb 20th, 2000", "2000-02-20T00:00:00Z"},
	{"May 2008", "2008-05-01T00:00:00Z"},
	{"Monday, May. 17, 2010", "2010-05-17T00:00:00Z"},        // (time.com)
	{"Thu Aug 25 10:46:55 BST 2011", "2011-08-25T09:46:55Z"}, // (www.yorkshireeveningpost.co.uk)

	//
	{"September, 26th 2011 by Christo Hall", "2011-09-26T00:00:00Z"}, // (www.thenewwolf.co.uk)
	{"Monday 30 July 2012 08.38 BST", "2012-07-30T07:38:00Z"},        // (guardian.co.uk)
}

func TestDates(t *testing.T) {
	for _, test := range noddydatetests {
		fd, _, err := ExtractDate(test.input)
		if err != nil {
			panic(err)
		}

		got := fmt.Sprintf("%04d-%02d-%02d", fd.Year(), fd.Month(), fd.Day())

		if err != nil {
			panic(err)
		}

		if got != test.expect {
			t.Errorf("ExtractDate('%v') = '%v', want '%v'", test.input, got, test.expect)
		}
	}
}

func TestTimes(t *testing.T) {
	for _, test := range noddytimetests {
		ft, _, err := ExtractTime(test.input)
		if err != nil {
			panic(err)
		}

		got := fmt.Sprintf("%02d:%02d:%02d", ft.Hour(), ft.Minute(), ft.Second())

		if err != nil {
			panic(err)
		}

		if got != test.expect {
			t.Errorf("ExtractTime('%v') = '%v', want '%v'", test.input, got, test.expect)
		}
	}
}

func TestDateTimes(t *testing.T) {
	for _, test := range noddydttests {
		fd, _, err := ExtractDate(test.input)
		if err != nil {
			panic(err)
		}

		if !fd.Equals(&test.date) {
			t.Errorf("ExtractDate('%v') = '%v', want '%v'", test.input, fd, test.date)
		}
	}
}
