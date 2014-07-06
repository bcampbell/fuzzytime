/*

Package fuzzytime helps with the parsing and representation of dates and times.


Sometimes, dates and times are ambiguous, and can't be parsed without
extra information (eg dd/mm/yy vs mm/dd/yy). The default behaviour is
for the Extract function to just return an error.
But this can be overriden by using a Context struct, which provides
functions to perform the decisions required in otherwise-ambiguous cases.



*/
package fuzzytime
