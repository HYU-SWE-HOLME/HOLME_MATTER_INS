package features

import (
	"INS_CURTAIN/src/terminal"
	"os"
	"strconv"

	"github.com/fatih/color"
)

//returnFormattedMsg
/*
  - It formats the message and returns the proper message.
- Parameter:
  - isHorizontal: bool -> whether the curtain opens/closes horizontaly/verticaly
  - isCenterMode: bool -> if horizontal, center mode or side to side mode
  - isLeftOrTop: bool -> whether the curtain is opened from the top, down, left, right
  - degree: bool -> how much the curtain is opened

- Return Value: string
	- It will return a completely formatted message.
*/
func returnFormattedMsg(isHorizontal bool, isCenterMode bool, isLeftOrTop bool, degree int) string {
	curtainType := ""
	curtainState := ""

	if isHorizontal {
		curtainType += "Horizontal"
		if isCenterMode {
			curtainType += " / Center Mode"
		} else {
			curtainType += " / Side to Side Mode"
		}
		if isLeftOrTop {
			curtainState += "left"
		} else {
			curtainState += "right"
		}
	} else {
		curtainType += "Vertical"
		curtainState += "Right"
		if isLeftOrTop {
			curtainState += "top"
		} else {
			curtainState += "down"
		}
	}
	curtainState += " / " + strconv.Itoa(degree)
	if degree < 50 {
		return "Curtain type: " + curtainType + "\n" + "State of the curtain: closed from the " + curtainState + "%"
	} else {
		return "Curtain type: " + curtainType + "\n" + "State of the curtain: opened from the " + curtainState + "%"
	}
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - degree: bool -> how much the curtain is opened

- Return Value: []byte
	- It returns byte array which is printable.
*/
func readImageFile(degree int) []byte {
	if degree < 50 {
		buf, err := os.ReadFile("src/raw/closed.txt")
		if err != nil {
			panic(err)
		}
		return buf
	} else {
		buf, err := os.ReadFile("src/raw/opened.txt")
		if err != nil {
			panic(err)
		}
		return buf
	}
}

// PrintClosedCurtain
/*
- Prints closed curtain.
- Reads closed curtain image and prints it.
- Parameter:
  - none

- Return Value: void
	- It will do its task and exit the function.
*/
func PrintClosedCurtain() {
	buf := readImageFile(0)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.White(string(buf))
	color.White(returnFormattedMsg(false, false, false, 0))
}

// PrintOpenedCurtain
/*
- Prints opened curtain.
- Reads opned curtain image and prints it.
- Parameter:
  - none

- Return Value: void
	- It will do its task and exit the function.
*/
func PrintOpenedCurtain() {
	buf := readImageFile(100)
	terminal.ClearTerminal() //* Clear the terminal first.
	color.HiBlack(string(buf))
	color.HiBlack(returnFormattedMsg(false, false, false, 0))
}
