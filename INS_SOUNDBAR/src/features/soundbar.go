package features

import (
	"INS_SOUNDBAR/src/terminal"
	"os"
	"strconv"
	"time"

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
		return "Soundbar: Music ON"
	} else {
		return "Soundbar: Music OFF"
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

// PrintSoundbar
/*
- Reads soundbar image and prints it.
- Parameter:
	- trigger: bool -> 0: music off, 1: musinc on
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintSoundbar(trigger bool) {
	if trigger {
		for i := 0; i < 99; i++ {
			figNum := i % 3
			buf := readImageFile(true, figNum)
			terminal.ClearTerminal() //* Clear the terminal first.
			color.White(string(buf))
			color.White(returnFormattedMsg(true))
			time.Sleep(200 * time.Millisecond)
		}
	} else {
		buf := readImageFile(false, 0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(false))
	}
}
