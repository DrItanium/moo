package main

import (
	"flag"
	_ "github.com/DrItanium/moo"
)

func main() {
	flag.Parse()
	if flag.NArg() <= 1 {
		flag.Usage()
	} else {

	}
}
