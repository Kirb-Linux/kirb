package modules

import (
	"encoding/json"
	"fmt"
	"os"
)

type Package struct {
	Name          string
	Cloneurl      string
	Installscript string
}

// Install the target package
func prep() Package {

	fmt.Println(len(os.Args))

	if len(os.Args) < 3 {
		fmt.Println("Usage: install <package>")
		os.Exit(0)
	}

	fmt.Println("Looking for: ", os.Args[2])

	var target Package

	Data := []byte(`{
		"name": "neofetch",
		"cloneurl": "https://github.com/dylanaraps/neofetch/archive/refs/tags/7.1.0.tar.gz",
		"installscript": "make install"
	}`)

	err := json.Unmarshal(Data, &target)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("", target)

	return target
}

func Install() {
	pkgInfo := prep()

	fmt.Println("Installing ", pkgInfo.Name, "...")

}
