package features

import (
	"INS_WATERDISPENSER/src/terminal"
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
func returnFormattedMsg(triggerReminder bool, triggerWater bool) string {
	if !triggerReminder { // before medication reminder
		return "Water dispenser"
	} else if !triggerWater { // after medication reminder, before users says "give me wagter"
		return "Water dispenser: Take your medicine!"
	} else { // after medication reminder, after users says "give me wagter"
		return "Water dispenser: Dispensing water..."
	}
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - trigger -> 0: not dispensing water, 1: dispensing water

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

// PrintWaterDispenser
/*
- Reads waterDispenser image and prints it.
- Parameter:
	- triggerReminder: bool -> 0: before medication reminder, 1: after/during medication reminder
	- triggerWater: bool 0-> 0: water not dispensed 1: water dispensed
- Return Value: void
	- It will do its task and exit the function.
*/
func PrintWaterDispenser(triggerReminder bool, triggerWater bool) {
	if !triggerReminder { // before medication reminder
		buf := readImageFile(false)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.White(string(buf))
		color.White(returnFormattedMsg(false, false))
	} else if !triggerWater { // after medication reminder, before users says "give me wagter"
		buf := readImageFile(false)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(true, false))
	} else { // after medication reminder, after users says "give me wagter"
		buf := readImageFile(true)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(true, true))
	}
}
