package modules

import (
	"encoding/json"
	"github.com/kirb-linux/kirb/helpers/query"
	"os"
)

var Packages []string

func GetDeps(pkg string) {
}

func CalculateDeps(pkg string) []string {
	// Calculate the dependencies, resolve everything
	// Start with the parent one

	currPkg := []byte(query.SearchPkgs(pkg))

	var target Package

	err := json.Unmarshal(currPkg, &target)

	if err != nil {
		os.Exit(0)
	}

	for _, dep := range target.Dependencies {
		Packages = append(Packages, query.GetInfo(dep).Name)
	}

	return Packages
}
