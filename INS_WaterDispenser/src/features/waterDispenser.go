package features

import (
	"INS_WATERDISPENSER/src/terminal"
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
func returnFormattedMsg(triggerReminder bool, triggerWater bool) string {
	if !triggerReminder {
		return "Water dispenser"
	} else if !triggerWater {
		return "Water dispenser: 200 mL dispensed"
	} else {
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
	if !triggerReminder {
		buf := readImageFile(false, 0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(triggerReminder, triggerWater))
	} else if triggerWater {
		for i := 0; i < 10; i++ {
			figNum := i % 3
			buf := readImageFile(true, figNum)
			terminal.ClearTerminal() //* Clear the terminal first.
			color.White(string(buf))
			color.White(returnFormattedMsg(triggerReminder, triggerWater))
			time.Sleep(200 * time.Millisecond)
		}
		buf := readImageFile(false, 0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.White(string(buf))
		color.White(returnFormattedMsg(true, false))
	} else {
		buf := readImageFile(false, 0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.White(string(buf))
		color.White(returnFormattedMsg(triggerReminder, triggerWater))
	}
}
