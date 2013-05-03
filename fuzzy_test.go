package fuzzytime

import (
	"testing"
)

type dtTest struct {
	input string
	date  Date
	time  Time
}

// TODO: add some more tests with numeric timezones

var dateTimeTests = []dtTest{
	{"2010-04-02T12:35:44+00:00", *NewDate(2010, 4, 2), *NewTime(12, 35, 44, "+00:00")}, //{iso8601, bbc blogs)
	{"2008-03-10 13:21:36 GMT", *NewDate(2008, 3, 10), *NewTime(13, 21, 36, "GMT")},     //{technorati api)

	{"9 Sep 2009 12.33", *NewDate(2009, 9, 9), *NewTime(12, 33, 0, "")},                        //(heraldscotland blogs)
	{"May 25 2010 3:34PM", *NewDate(2010, 5, 25), *NewTime(15, 34, 0, "")},                     //(thetimes.co.uk)
	{"Thursday August 21 2008 10:42 am", *NewDate(2008, 8, 21), *NewTime(10, 42, 0, "")},       //(guardian blogs in their new cms)
	{"Tuesday October 14 2008 00.01 BST", *NewDate(2008, 10, 14), *NewTime(0, 1, 0, "BST")},    //(Guardian blogs in their new cms)
	{"Tuesday 16 December 2008 16.23 GMT", *NewDate(2008, 12, 16), *NewTime(16, 23, 0, "GMT")}, //(Guardian blogs in their new cms)
	{"3:19pm on Tue 29 Jan 08", *NewDate(2008, 1, 29), *NewTime(15, 19, 0, "")},                //(herald blogs)
	{"2007/03/18 10:59:02", *NewDate(2007, 3, 18), *NewTime(10, 59, 2, "")},
	{"Mar 3, 2007 12:00 AM", *NewDate(2007, 3, 3), *NewTime(0, 0, 0, "")},
	{"Jul 21, 08 10:00 AM", *NewDate(2008, 7, 21), *NewTime(10, 0, 0, "")},           //(mirror blogs)
	{"09-Apr-2007 00:00", *NewDate(2007, 4, 9), *NewTime(0, 0, 0, "")},               //(times, sundaytimes)
	{"4:48PM GMT 22/02/2008", *NewDate(2008, 2, 22), *NewTime(16, 48, 0, "GMT")},     //(telegraph html articles)
	{"09-Apr-07 00:00", *NewDate(2007, 4, 9), *NewTime(0, 0, 0, "")},                 //(scotsman)
	{"Friday    August    11, 2006", *NewDate(2006, 8, 11), Time{}},                  //(express, guardian/observer)
	{"26 May 2007, 02:10:36 BST", *NewDate(2007, 5, 26), *NewTime(2, 10, 36, "BST")}, //(newsoftheworld)
	{"2:43pm BST 16/04/2007", *NewDate(2007, 4, 16), *NewTime(14, 43, 0, "BST")},     //(telegraph, after munging)
	{"20:12pm 23rd November 2007", *NewDate(2007, 11, 23), *NewTime(20, 12, 0, "")},  //(dailymail)
	//TODO        {"2:42 PM on 22nd May 2008", *NewDate(2008,5,22),*NewTime(14,42,0,"")}, //(dailymail)
	{"February 10 2008 22:05", *NewDate(2008, 2, 10), *NewTime(22, 5, 0, "")}, //(ft)
	//        {"22 Oct 2007, //(weird non-ascii characters) at(weird non-ascii characters)11:23", *NewDate(2007,10,22),*NewTime(11,23,0,"")}, //(telegraph blogs OLD!)
	{"Feb 2, 2009 at 17:01:09", *NewDate(2009, 2, 2), *NewTime(17, 1, 9, "")},                   //(telegraph blogs)
	{"18 Oct 07, 04:50 PM", *NewDate(2007, 10, 18), *NewTime(16, 50, 0, "")},                    //(BBC blogs)
	{"02 August 2007  1:21 PM", *NewDate(2007, 8, 2), *NewTime(13, 21, 0, "")},                  //(Daily Mail blogs)
	{"October 22, 2007  5:31 PM", *NewDate(2007, 10, 22), *NewTime(17, 31, 0, "")},              //(old Guardian blogs, ft blogs)
	{"October 15, 2007", *NewDate(2007, 10, 15), Time{}},                                        //(Times blogs)
	{"February 12 2008", *NewDate(2008, 2, 12), Time{}},                                         //(Herald)
	{"Monday, 22 October 2007", *NewDate(2007, 10, 22), Time{}},                                 //(Independent blogs, Sun (page date))
	{"22 October 2007", *NewDate(2007, 10, 22), Time{}},                                         //(Sky News blogs)
	{"11 Dec 2007", *NewDate(2007, 12, 11), Time{}},                                             //(Sun (article date))
	{"12 February 2008", *NewDate(2008, 2, 12), Time{}},                                         //(scotsman)
	{"03/09/2007", *NewDate(2007, 9, 3), Time{}},                                                //(Sky News blogs, mirror)
	{"Tuesday, 21 January, 2003, 15:29 GMT", *NewDate(2003, 1, 21), *NewTime(15, 29, 0, "GMT")}, //(historical bbcnews)
	{"2003/01/21 15:29:49", *NewDate(2003, 1, 21), *NewTime(15, 29, 49, "")},                    //(historical bbcnews (meta tag))
	{"2010-07-01", *NewDate(2010, 7, 1), Time{}},
	{"2010/07/01", *NewDate(2010, 7, 1), Time{}},
	{"Feb 20th, 2000", *NewDate(2000, 2, 20), Time{}},
	{"Monday, May. 17, 2010", *NewDate(2010, 5, 17), Time{}}, // (time.com)

	// TODO: this is a tricky one where hour can get picked up as year if not careful!
	{"Thu Aug 25 10:46:55 BST 2011", *NewDate(2011, 8, 25), *NewTime(10, 46, 55, "BST")}, // (www.yorkshireeveningpost.co.uk)

	//
	{"September, 26th 2011 by Christo Hall", *NewDate(2011, 9, 26), Time{}},             // (www.thenewwolf.co.uk)
	{"Monday 30 July 2012 08.38 BST", *NewDate(2012, 7, 30), *NewTime(8, 38, 0, "BST")}, // (guardian.co.uk)

	// some more obscure cases...
	{"May 2008", *NewDate(2008, 5, 1), Time{}},
}

func TestDateTimes(t *testing.T) {
	for _, test := range dateTimeTests {
		fd, ft := Extract(test.input)

		if !fd.Equals(&test.date) {
			t.Errorf("ExtractDate('%v') = '%v', want '%v'", test.input, fd.String(), test.date.String())
		}

		if !ft.Equals(&test.time) {
			t.Errorf("ExtractTime('%v') = '%v', want '%v'", test.input, ft.String(), test.time.String())
		}
	}

	// Tricky case that needs fixing:
	fd, _ := ExtractDate("Thu Aug 25 10:46:55 BST 2011")
	if fd.Year() != 2011 {
		t.Errorf("2-digit year issue still needs fixing.")
	}

}
