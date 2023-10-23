/*
- Matter Instance: Soundbar
*/

package main

import (
	soundbar "INS_SOUNDBAR/src/core"
	"log"
)

func main() {
	//* Initiate soundbar Instance
	sb := soundbar.Init()
	err := sb.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 7 - Soundbar): %v", err)
	}
}
