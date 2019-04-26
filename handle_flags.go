package main

import (
	"flag"
)

var do_fetch bool

func init() {
	flag.BoolVar(&do_fetch, "fetch", true, "Fetch the latest share prices")
}

func handle_flags() (bool){
	flag.Parse()
	return do_fetch
}
