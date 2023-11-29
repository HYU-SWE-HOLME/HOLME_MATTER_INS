package features

import (
	"INS_CURTAIN/src/terminal"
	"bytes"
	"os"
	"time"

	"github.com/fatih/color"
)

const linesToRead = 16

//returnFormattedMsg
/*
  - It formats the message and returns the proper message.
- Parameter:
  - isHorizontal: bool -> whether the curtain opens/closes horizontaly/verticaly
  - isCenterMode: bool -> if horizontal, center mode or side to side mode
  - isLeftOrTop: bool -> whether the curtain is opened from the top, down, left, right
  - degree: bool -> how much the curtain is opened
  - state: int -> what line is being printed
- Return Value: string
	- It will return a completely formatted message.
*/
func returnFormattedMsg(isHorizontal bool, isCenterMode bool, isLeftOrTop bool, degree int, state int) string {
	curtainType := ""
	curtainState := ""

	if degree > 50 {
		curtainState += " / closed"
	} else if state < 97 {
		curtainState += " / opening..."
	} else {
		curtainState += " / opened"
	}

	if isHorizontal {
		curtainType += "Horizontal"
		return "\nCurtain type: " + curtainType + curtainState
	} else {
		curtainType += "Vertical"
		return "\nCurtain type: " + curtainType + curtainState
	}
}

func extractLines(buf []byte, i, j int) []byte {
	lines := bytes.Split(buf, []byte("\n"))
	if j > len(lines) {
		j = len(lines)
	}
	if i > j {
		return nil
	}
	return bytes.Join(lines[i-1:j], []byte("\n"))
}

// readImageFile
/*
- It reads the image file, return the byte array available to print.
- Parameter:
  - degree: int -> how much the curtain is opened

- Return Value: []byte
	- It returns byte array which is printable.
*/
func readImageFile(degree int, startLine int) []byte {
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
		return extractLines(buf, startLine, startLine+linesToRead)
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
	if degree > 50 { // closed(down)
		buf := readImageFile(0, 0)
		terminal.ClearTerminal() //* Clear the terminal first.
		color.White(string(buf))
		color.White(returnFormattedMsg(false, false, false, degree, 0))
	} else { //opened(up)
		for i := 1; i <= 98; i += linesToRead {
			buf := readImageFile(100, i)
			terminal.ClearTerminal() //* Clear the terminal first.
			color.HiBlack(string(buf))
			color.HiBlack(returnFormattedMsg(false, false, false, degree, i))
			time.Sleep(1 * time.Second)
		}
	}
}
