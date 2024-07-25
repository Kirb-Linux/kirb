package helpers

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func YesNo() {
	fmt.Print("Proceed with action? (y/N): ")
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		os.Exit(0)
	}

	if strings.Contains(input, "n") || strings.Contains(input, "N") {
		color.Red("Exiting...")
		os.Exit(0)
	}

	if strings.Contains(input, "y") || strings.Contains(input, "Y") {
		return
	}

	YesNo()
}
