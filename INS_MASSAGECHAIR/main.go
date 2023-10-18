/*
- Matter Instance: MassageChair
*/

package main

import (
	massageChair "INS_MASSAGECHAIR/src/core"
	"log"
)

func main() {
	//* Initiate massageChair Instance
	sb := massageChair.Init()
	err := sb.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 4 - Refrigerator): %v", err)
	}
}
