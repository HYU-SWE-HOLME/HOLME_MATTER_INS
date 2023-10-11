package features

import (
	"github.com/fatih/color"
	"os"
	"strconv"
)

func returnPrintingMsg(state bool, degree int, color string) string {
	if !state {
		return "Light state: OFF"
	} else {
		return "Light state: ON\nCurrent Light Degree: " + strconv.Itoa(degree) + "%\nCurrent Light Color: " + color
	}
}

func PrintDefaultLight() {
	buf, err := os.ReadFile("src/raw/default.txt")
	if err != nil {
		panic(err)
	}
	color.HiBlack(string(buf))
	color.HiBlack(returnPrintingMsg(false, 0, ""))
}
