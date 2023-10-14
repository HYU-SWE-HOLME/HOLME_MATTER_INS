package main

import (
	"HIVEMIND/handler"
	"HIVEMIND/utils"
	"log"
	"net/http"
)

const PORT string = "10000"

func main() {
	utils.PrintBanner()
	log.Println("HIVEMIND is now running on port 10000")
	err := http.ListenAndServe(":"+PORT, handler.RequestHandler())

	if err != nil {
		log.Fatalf("ERROR! Occured at initialization phase of HIVEMIND.")
	}
}
