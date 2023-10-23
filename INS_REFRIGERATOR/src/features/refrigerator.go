package features

import (
	"INS_REFRIGERATOR/src/terminal"
	"os"

	"github.com/fatih/color"
)

//returnFormattedMsg
/*
  - It formats the message and returns the proper message.
- Parameter:
  - trigger(sleep mode): bool -> 0: off, 1: on

- Return Value: string
	- Returns formatted message.
*/
func returnFormattedMsg(trigger bool) string {
	if trigger {
		return "Refrigerator screen: ON"
	} else {
		return "Refrigerator screen: OFF"
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
func readImageFile() []byte {
	buf, err := os.ReadFile("src/raw/default.txt")
	if err != nil {
		panic(err)
	}
	return buf
}

// PrintRefrigerator
/*
- Reads refrigerator image and prints it.
- Parameter:
	- trigger: bool -> 0: sleep mode off, 1: sleep mode on
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintRefrigerator(trigger bool) {
	buf := readImageFile()
	terminal.ClearTerminal() //* Clear the terminal first.
	if trigger {
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(true))
	} else {
		color.White(string(buf))
		color.White(returnFormattedMsg(false))
	}
}
