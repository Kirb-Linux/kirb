package main

import (
	"fmt"
	"github.com/kirb-linux/kirb/modules"
	"os"
)

func main() {

	modules.Prerun()

	if len(os.Args) == 1 {
		fmt.Println("No argument specified.")
		os.Exit(0)
	}

	if os.Args[1] == "install" || os.Args[1] == "i" {
		modules.Install()
	}
}
