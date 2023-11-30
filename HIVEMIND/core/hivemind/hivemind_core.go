package hivemind

import (
	"HIVEMIND/utils"
	"log"
	"net/http"
)

const PORT string = "10000"

type Hivemind struct {
	//* Methods will be defined as a receiver below.
}

func Init() *Hivemind {
	return &Hivemind{}
}

func (hb *Hivemind) Start() error { //* Hivemind will initiated.

	utils.PrintBanner()
	log.Println("HOLME HUB is running... HOLME is now available for current environment!")

	err := http.ListenAndServe(":"+PORT, RequestHandler())
	if err != nil {
		return err
	}

	return nil
}
