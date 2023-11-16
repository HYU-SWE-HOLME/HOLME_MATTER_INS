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
func returnFormattedMsg(trigger bool, askForReplacement bool, replacement bool) string {
	if !trigger {
		return "AI Speaker state: OFF"
	} else if askForReplacement {
		return "It looks like there's a soundbar in this hotel.\nShall I play the music through the soundbar?"
	} else if replacement {
		return "Sure, I'll play music through the soundbar for you."
	} else { // no replacement
		return "I'll play the music for you"
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
func readImageFile(trigger bool, askForReplacement bool, replacement bool) []byte {
	if trigger && !askForReplacement && !replacement {
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
func PrintAiSpeaker(trigger bool, askForReplacement bool, replacement bool) {
	buf := readImageFile(trigger, askForReplacement, replacement)
	terminal.ClearTerminal() //* Clear the terminal first.
	if !trigger {
		color.White(string(buf))
		color.White(returnFormattedMsg(trigger, askForReplacement, replacement))
	} else if askForReplacement || (!askForReplacement && !replacement) {
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(trigger, askForReplacement, replacement))
	} else {
		color.White(string(buf))
		color.White(returnFormattedMsg(trigger, askForReplacement, replacement))
	}
}
