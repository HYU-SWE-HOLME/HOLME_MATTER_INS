/*
- Matter Instance: Aircon
*/

package main

import (
	aircon "INS_AIRCON/src/core"
	"log"
)

func main() {
	//* Initiate aircon Instance
	ac := aircon.Init()
	err := ac.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 3 - air conditioner ): %v", err)
	}
}
