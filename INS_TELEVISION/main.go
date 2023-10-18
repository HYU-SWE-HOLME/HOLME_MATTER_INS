/*
- Matter Instance: Television
*/

package main

import (
	television "INS_TELEVISION/src/core"
	"log"
)

func main() {
	//* Initiate Refrigerator Instance
	tv := television.Init()
	err := tv.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 4 - Refrigerator): %v", err)
	}
}
