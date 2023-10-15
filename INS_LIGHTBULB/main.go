/*
- Matter Instance: Light bulb
*/

package main

import (
	lightBulb "INS_LIGHTBULB/src/core"
	"log"
)

func main() {
	//* Initiate LightBulb Instance
	lb := lightBulb.Init()
	err := lb.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 1 - light bulb): %v", err)
	}
}
