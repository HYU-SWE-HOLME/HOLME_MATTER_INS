package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func readBannerFile() string {
	buf, err := os.ReadFile("raw/banner.txt")

	if err != nil {
		panic("Error while printing the banner file!")
	}

	return string(buf)
}

func PrintBanner() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic("Error while printing the banner file!")
	}
	banner := readBannerFile()
	fmt.Println(banner)

	return
}
