package fuzzytime

import (
	"fmt"
)

func ExampleExtract() {

	inputs := []string{
		"Wed Apr 16 17:32:51 NZST 2014",
		"2010-02-01T13:14:43Z", // an iso 8601 form
		"no date or time info here",
		"Published on March 10th, 1999 by Brian Credability",
		"2:51pm",
	}

	for _, inp := range inputs {
		dt := Extract(inp)
		fmt.Println(dt.ISOFormat())
	}

	// Output:
	// 2014-04-16T17:32:51+12:00
	// 2010-02-01T13:14:43Z
	//
	// 1999-03-10
	// T14:51
}

func ExampleContext() {
	inputs := []string{
		"01/02/03",
		"12/23/99",
		"10:25CST",
	}
	// USA context:
	fmt.Println("in USA:")
	for _, inp := range inputs {
		dt := USContext.Extract(inp)
		fmt.Println(dt.ISOFormat())
	}

	// custom context for Australia:
	aussie := Context{
		DateResolver: DMYResolver,
		TZResolver:   DefaultTZResolver("AU"),
	}

	fmt.Println("in Australia:")
	for _, inp := range inputs {
		dt := aussie.Extract(inp)
		fmt.Println(dt.ISOFormat())
	}
	// Output:
	// in USA:
	// 2003-01-02
	// 1999-12-23
	// T10:25-06:00
	// in Australia:
	// 2003-02-01
	//
	// T10:25+09:30

}
