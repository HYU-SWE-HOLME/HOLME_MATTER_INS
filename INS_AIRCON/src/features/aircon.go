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
func returnFormattedMsg(trigger bool, mode string, airflowDirect bool, fanSpeed int,
	brightnessScreen int, objTemperature int, startWakeupTimer bool, startShutdwonTimer bool,
	stopWakeupTimer bool, stopShutdwonTimer bool, wakeupTime int, shutdownTime int) string {
	_airflowMode := ""
	if airflowDirect {
		_airflowMode = "direct"
	} else {
		_airflowMode = "indirect"
	}

	if trigger {
		if startWakeupTimer {
			return "Wake up timer set: " + strconv.Itoa(int(wakeupTime)) + ":00"
		} else if startShutdwonTimer {
			return "Shut down timer set: " + strconv.Itoa(int(shutdownTime)) + ":00"
		} else if stopWakeupTimer {
			return "Wake up timer stopped"
		} else if stopShutdwonTimer {
			return "Shut down timer stopped"
		} else { // No timer set
			return "Air conditioner state: ON\n" +
				"Air conditioner mode: " + mode +
				"\nAirflow: " + _airflowMode +
				"\nFan speed: " + strconv.Itoa(int(fanSpeed)) +
				"\nUser set temperature: " + strconv.Itoa(int(objTemperature)) +
				"\nScreen brightness: " + strconv.Itoa(int(brightnessScreen)) + "\n"
		}
	} else {
		return "Air conditioner state: OFF"
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
func PrintAirconOn(trigger bool, mode string, airflowDirect bool, fanSpeed int,
	brightnessScreen int, objTemperature int, startWakeupTimer bool, startShutdwonTimer bool,
	stopWakeupTimer bool, stopShutdwonTimer bool, wakeupTime int, shutdownTime int) {
	buf := readImageFile(true)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.HiBlack(string(buf))
	color.HiBlack(returnFormattedMsg(true, mode, airflowDirect, fanSpeed,
		brightnessScreen, objTemperature, startShutdwonTimer, startShutdwonTimer,
		stopWakeupTimer, stopShutdwonTimer, wakeupTime, shutdownTime))
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
func PrintAirconOff() {
	buf := readImageFile(false)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.White(string(buf))
	color.White(returnFormattedMsg(false, "", true, 0,
		0, 0, true, true,
		true, true, 0, 0))
}
