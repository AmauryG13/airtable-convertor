package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/amauryg13/airtable-convertor-go/convertor"
)

var fnUsage = func() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("  %s\n        %s\n", "filename", "Path of the file to be converted")
	flag.PrintDefaults()
}

type uwc []string

func (u *uwc) String() string {
	return fmt.Sprint(*u)
}

func (u *uwc) Set(value string) error {
	if len(*u) > 0 {
		return errors.New("Unwanted chars (uwc) flag already set")
	}
	for _, c := range strings.Split(value, ",") {
		*u = append(*u, c)
	}
	return nil
}

var filepath string
var sep string
var eol string
var uwcFlag uwc

func init() {
	flag.StringVar(&sep, "sep", ";", "Default record separator")
	flag.StringVar(&eol, "eol", "\n", "Default end of line character")
	flag.Var(&uwcFlag, "uwc", "Additional (comma separated) unwanted chars to removed")
}

func main() {
	flag.Parse()

	if flag.NArg() == 1 {
		filepath = flag.Args()[0]
	} else {
		fnUsage()
		os.Exit(1)
	}

	removedChars := uwcFlag

	var options = make(map[string]string)
	options["eol"] = eol
	options["sep"] = sep

	c := convertor.New(filepath, removedChars, options)
	c.Run()
}
