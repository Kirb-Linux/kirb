package modules

import (
	"fmt"
	"github.com/kirb-linux/kirb/demo"
	"os"
)

var (
	conffile *os.FileInfo
	err      error
)

func Prerun() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Oops, something went wrong!", err)
		os.Exit(1)
	}
	_, err = os.Stat(homedir + "/.config/kirb/config.json")

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Kirb config file doesn't exist, creating a new one.")
			err = os.MkdirAll(homedir+"/.config/kirb", os.ModePerm)
			if err != nil {
				fmt.Println("Oops, something went wrong!", err)
				os.Exit(1)
			}

			f, err := os.Create(homedir + "/.config/kirb/config.json")

			err = os.WriteFile(homedir+"/.config/kirb/config.json", demo.DefConfig(), 0644)

			if err != nil {
				fmt.Println("Oops, something went wrong!", err)
				os.Exit(1)
			}
			defer f.Close()
			fmt.Println("Created Kirb config file. ", f.Name())
		}
	}

}
