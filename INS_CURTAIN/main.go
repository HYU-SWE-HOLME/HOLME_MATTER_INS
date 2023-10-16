/*
- Matter Instance: Curtain
*/

package main

import (
	curtain "INS_CURTAIN/src/core"
	"log"
)

func main() {
	//* Initiate LightBulb Instance
	lb := curtain.Init()
	err := lb.Start()
	if err != nil {
		log.Fatalf("FATAL ERROR (Instance 2 - curtain ): %v", err)
	}
}
