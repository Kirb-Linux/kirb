package net

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/kirb-linux/kirb/helpers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var mirror string

var target helpers.Config

func getSettings() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dat, err := os.ReadFile(filepath.Join(homedir, ".config", "kirb", "config.json"))
	if err != nil {
		color.Red("Something went wrong when reading the settings file.")
		log.Fatal(err)
	}

	err = json.Unmarshal(dat, &target)

	if err != nil {
		color.Red("An error occured while parsing the settings file.")
		log.Fatal(err)
	}

	mirror = target.Mirror

}

func Get(endpoint string) string {
	getSettings()

	resp, err := http.Get(mirror + endpoint)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
