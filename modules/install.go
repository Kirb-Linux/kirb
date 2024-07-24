package modules

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/codeclysm/extract/v3"
	"github.com/fatih/color"
	"github.com/kirb-linux/kirb/globals"
	"github.com/kirb-linux/kirb/helpers"
	"github.com/kirb-linux/kirb/helpers/query"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Package struct {
	Name          string
	Filename      string
	Cloneurl      string
	Workdir       string
	Installscript string
	Checksum      string
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func download(url string, name string) error {
	err := DownloadFile(name, url)
	if err != nil {
		return err
	}
	color.Green("Downloaded %s from %s\n", name, url)
	return err
}

// Install the target package
func prep() Package {

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: install <package>")
		os.Exit(0)
	}

	fmt.Println("Looking for:", flag.Args()[1])

	var target Package

	Data := []byte(query.SearchPkgs(flag.Args()[1]))

	err := json.Unmarshal(Data, &target)

	if err != nil {
		fmt.Println(err)
	}

	return target
}

func Install() {
	pkgInfo := prep()

	fmt.Println("Installing", pkgInfo.Name, "...")

	err := os.Chdir("/tmp")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = download(pkgInfo.Cloneurl, pkgInfo.Filename)
	if err != nil {
		return
	}

	file, err := os.Open(filepath.Join("/tmp", pkgInfo.Filename))

	hash := helpers.Sha256(file)

	if hash != pkgInfo.Checksum {

		color.Red("Hash is mismatch (expected %s, got %s)", pkgInfo.Checksum, hash)
		os.Exit(1)
	}

	file, err = os.Open(filepath.Join("/tmp", pkgInfo.Filename))

	color.Green("Unarchiving files..")

	err = extract.Gz(context.TODO(), file, "/tmp", nil)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Chdir(filepath.Join(os.TempDir(), pkgInfo.Workdir))
	if err != nil {
		color.Red("Something went wrong. Contact the package author.")
		log.Fatal(err)
	}

	color.Green("Executing install script")

	cmd := exec.Command("sh", "-c", pkgInfo.Installscript)
	stdout, err := cmd.StdoutPipe()
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	if globals.Quiet == false {
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Println(msg)
		}
	}
	cmd.Wait()

	color.Green(pkgInfo.Name + " installed successfully! Cleaning up")

	err = os.RemoveAll(filepath.Join(os.TempDir(), pkgInfo.Workdir))
	if err != nil {
		log.Fatal(err)
	}
}
