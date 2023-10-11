package terminal

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
execTerminal(): Core Terminal Function - execute terminal functions while features running.
- Parameter:
  - name: string, a command for execution.
  - args: []string, a argument for the command.

- Return Value: Boolean
  - Ture if the execution is successful.
  - False if something gone wrong while execution.
*/
func execTerminal(name string, args []string) bool {
	var cmd *exec.Cmd //* Init cmd features

	if args == nil {
		//* No arguments
		cmd = exec.Command(name)
	} else {
		//* Argument exists.
		concatArgs := strings.Join(args[1:], " ") //* Concatenated Arguments
		cmd = exec.Command(name, concatArgs)
	}

	cmd.Stdout = os.Stdout //* Fix CMD Execution as standard output.

	if err := cmd.Run(); err != nil { //* Run Command
		//* Error occurred
		log.Fatalf("Failed to execute terminal: %v", err)
		return false
	}

	return true
}

func ClearTerminal() {
	if ret := execTerminal("clear", nil); !ret {
		log.Fatalf("Error on ClearTerminal()")
	}

	return
}
