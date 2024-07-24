package query

import (
	"github.com/kirb-linux/kirb/helpers/net"
)

func SearchPkgs(search string) string {
	return net.Get("/pkgs/" + search)
}
