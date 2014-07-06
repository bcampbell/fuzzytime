/*

Package fuzzytime helps with the parsing and representation of dates and times.

Fuzzytime defines types (Date, Time, DateTime) which have optional fields.
So you can represent a date with a year and month, but no day.

Sometimes dates and times are ambiguous and can't be parsed without
extra information (eg "dd/mm/yy" vs "mm/dd/yy"). The default behaviour when
such a data is encountered is for Extract() function to just return an error.
This can be overriden by using a Context struct, which provides
functions to perform the decisions required in otherwise-ambiguous cases.

Timezones are stored as an offset from UTC (in seconds).

*/
package fuzzytime
