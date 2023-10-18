/*
- Matter Instance: Water Dispenser
*/

package main

import (
	refrigerator "INS_WATERDISPENSER/src/core"
	"log"
)

func main() {
	//* Initiate WaterDispenser Instance
	wd := refrigerator.Init()
	err := wd.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 5 - Water Dispenser): %v", err)
	}
}
