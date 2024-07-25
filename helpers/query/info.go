package query

import (
	"encoding/json"
)

var pkg []byte

type Package struct {
	Name          string
	Filename      string
	Cloneurl      string
	Workdir       string
	Installscript string
	Checksum      string
	Description   string
	Dependencies  []string
}

func GetInfo(pkgname string) Package {
	pkg = []byte(SearchPkgs(pkgname))

	var target Package

	err := json.Unmarshal(pkg, &target)

	if err != nil {
		panic(err)
	}

	return Package(target)
}
