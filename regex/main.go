package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	text := "FooBarBaz"
	matched, err := regexp.MatchString("Bar", text)
	fmt.Println(matched, err) // true nil
	matched, err = regexp.MatchString("^Foo", text)
	fmt.Println(matched, err) // true nil
	matched, err = regexp.MatchString("Baz$", text)
	fmt.Println(matched, err) // true nil
	matched, err = regexp.MatchString("bad[", text)
	fmt.Println(matched, err) // false error parsing regexp, missing closing ]

	// The file haiku.txct contains
	// Out of memory
	// We wish to hold the whole sky,
	// but we never will.

	bytes, err := ioutil.ReadFile("haiku.txt")
	if err != nil {
		log.Fatal(err)
	}
	matched, err = regexp.Match("whole sky", bytes)
	fmt.Println(matched, err) // true nil
}
