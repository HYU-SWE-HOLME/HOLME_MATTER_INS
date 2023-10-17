/*
- Matter Instance: Refrigerator
*/

package main

import (
	refrigerator "INS_REFRIGERATOR/src/core"
	"log"
)

func main() {
	//* Initiate Refrigerator Instance
	rf := refrigerator.Init()
	err := rf.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 4 - Refrigerator): %v", err)
	}
}
