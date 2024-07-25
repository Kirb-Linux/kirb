package modules

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/kirb-linux/kirb/helpers/query"
	"os"
)

var pkg []byte

func parse() {
	pkg = []byte(query.SearchPkgs(flag.Args()[1]))

	var target Package

	err := json.Unmarshal(pkg, &target)

	if err != nil {
		panic(err)
	}

	fmt.Println("Found 1 package")

	color.Set(color.FgCyan)
	fmt.Printf("Package Name: %s\n", target.Name)
	fmt.Printf("Description: %s\n", target.Description)
	fmt.Printf("Checksum: %s\n", target.Checksum)
	fmt.Printf("Dependencies: %s\n", target.Dependencies)
	color.Unset()
}

func Search() {
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: search <package>")
		os.Exit(0)
	}

	fmt.Println("Looking for " + flag.Args()[1])

	parse()
}
