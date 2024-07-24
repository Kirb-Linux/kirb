package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/codeclysm/extract/v3"
	"github.com/fatih/color"
	"github.com/kirb-linux/kirb/helpers"
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
	Sha256        string
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
		"installscript": "make install",
		"filename": "7.1.0.tar.gz",
		"workdir": "neofetch-7.1.0",
		"sha256": "58a95e6b714e41efc804eca389a223309169b2def35e57fa934482a6b47c27e7"
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

	if hash != pkgInfo.Sha256 {

		color.Red("Hash is mismatch (expected %s, got %s)", pkgInfo.Sha256, hash)
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

	cmd := exec.Command("sh", "-c", pkgInfo.Installscript)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pkgInfo.Name, "installed successfully!")
}
