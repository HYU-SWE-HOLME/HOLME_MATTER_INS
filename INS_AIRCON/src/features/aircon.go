package features

import (
	"INS_AIRCON/src/terminal"
	"os"
	"strconv"

	"github.com/fatih/color"
)

//returnFormattedMsg
/*
  - It formats the message and returns the proper message.
- Parameter:
  - trigger: bool -> 0: off, 1: on
  - temperature: int32 -> user set temperature
  - degree: int32 -> air conditioner intensity

- Return Value: string
	- Returns formatted message.
*/
func returnFormattedMsg(trigger bool, temperature int, degree int) string {
	if !trigger {
		return "Air conditioner state: OFF"
	} else {
		return "Air conditioner state: ON\nUser set temperature: " + strconv.Itoa(int(temperature)) + "\nWind intesity: " + strconv.Itoa(int(degree))
	}
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - trigger -> 0: off, 1: on

- Return Value: []byte
	- It returns printable byte array.
*/
func readImageFile(trigger bool) []byte {
	if trigger {
		buf, err := os.ReadFile("src/raw/on.txt")
		if err != nil {
			panic(err)
		}
		return buf
	} else {
		buf, err := os.ReadFile("src/raw/off.txt")
		if err != nil {
			panic(err)
		}
		return buf
	}
}

// PrintAirconOn
/*
- Reads aircon(turned on) image and prints it.
- Parameter:
  	- temperature: int
	- degree: int
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintAirconOn(temperature int, degree int) {
	buf := readImageFile(true)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.White(string(buf))
	color.White(returnFormattedMsg(true, temperature, degree))
}

// PrintAirconOff
/*
- Reads aircon(turned off) image and prints it.
- Parameter:
  	- temperature: int
	- degree: int
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintAirconOff(temperature int, degree int) {
	buf := readImageFile(false)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.HiBlack(string(buf))
	color.HiBlack(returnFormattedMsg(false, temperature, degree))
}
