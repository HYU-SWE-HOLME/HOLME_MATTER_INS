/*
- Matter Instance: AiSpeaker
*/

package main

import (
	aiSpeaker "INS_AISPEAKER/src/core"
	"log"
)

func main() {
	//* Initiate AI Speaker Instance
	as := aiSpeaker.Init()
	err := as.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 9 - AI Speaker): %v", err)
	}
}
