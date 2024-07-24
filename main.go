package main

import (
	"flag"
	"fmt"
	"github.com/kirb-linux/kirb/globals"
	"github.com/kirb-linux/kirb/modules"
	"os"
)

var mirror string

func main() {

	modules.Prerun()

	var quietPtr = flag.Bool("quiet", false, "suppress output")

	flag.Parse()

	globals.Quiet = *quietPtr

	if len(flag.Args()) == 0 {
		fmt.Println("No argument specified.")
		os.Exit(0)
	}

	if flag.Args()[0] == "install" || flag.Args()[0] == "i" {
		modules.Install()
	}

	if flag.Args()[0] == "search" || flag.Args()[0] == "s" {
		modules.Search()
	}
}
