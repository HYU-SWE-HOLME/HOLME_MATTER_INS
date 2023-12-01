package features

import (
	"INS_TELEVISION/src/terminal"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

//returnFormattedMsg
/*
  - It formats the message and returns the proper message.
- Parameter:
  - trigger: bool -> 0: off, 1: on

- Return Value: string
	- Returns formatted message.
*/
func returnFormattedMsg(trigger bool) string {
	if trigger {
		return "Television: Turned ON"
	} else {
		return "Television: Turned OFF"
	}
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - trigger -> 0: music off, 1: music on

- Return Value: []byte
	- It returns printable byte array.
*/
func readImageFile(trigger bool, figNum int) []byte {
	if trigger {
		buf, err := os.ReadFile("src/raw/on_" + strconv.Itoa(figNum) + ".txt")
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

// PrintTelevision
/*
- Reads television image and prints it.
- Parameter:
	- trigger: bool -> 0: music off, 1: musinc on
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintTelevision(trigger bool) {
	if trigger {
		// i := 0
		for i := 0; i < 10; i++ {
			figNum := i % 3
			buf := readImageFile(true, figNum)
			terminal.ClearTerminal() //* Clear the terminal first.
			color.White(string(buf))
			color.White(returnFormattedMsg(true))
			time.Sleep(100 * time.Millisecond)
		}
		buf := readImageFile(true, 3)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.White(string(buf))
		color.White(returnFormattedMsg(true))
		time.Sleep(100 * time.Millisecond)
	} else {
		buf := readImageFile(false, 0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(false))
	}
}
