package features

import (
	"INS_AISPEAKER/src/terminal"
	"os"

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
func returnFormattedMsg(trigger bool) string {
	if !trigger {
		return "AI Speaker state: OFF"
	} else {
		return "AI Speaker state: ON"
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

// PrintAiSpeaker
/*
- Reads aiSpeaker(turned on) image and prints it.
- Parameter:
  	- trigger: bool
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintAiSpeaker(trigger bool) {
	buf := readImageFile(trigger)
	terminal.ClearTerminal() //* Clear the terminal first.

	if trigger {
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(trigger))
	} else {
		color.White(string(buf))
		color.White(returnFormattedMsg(trigger))
	}
}
