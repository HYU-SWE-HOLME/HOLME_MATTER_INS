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
	} else {
		curtainType += "Vertical"
	}
	curtainState += strconv.Itoa(degree)
	return "Curtain state: " + curtainType + " / " + curtainState + "% opened"
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - degree: int -> how much the curtain is opened

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

// PrintCurtain
/*
- Prints curtain.
- Reads curtain image and prints it.
- Parameter:
  - degree: int

- Return Value: void
	- It will do its task and exit the function.
*/
func PrintCurtain(degree int) {
	if degree < 50 { // closed(down)
		buf := readImageFile(0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.HiBlack(string(buf))
		color.HiBlack(returnFormattedMsg(false, false, false, degree))
	} else { //opened(up)
		buf := readImageFile(100)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.White(string(buf))
		color.White(returnFormattedMsg(false, false, false, degree))
	}
}
