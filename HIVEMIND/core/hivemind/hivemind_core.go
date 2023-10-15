package hivemind

import (
	"HIVEMIND/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT string = "10000"

type Hivemind struct {
	//* Methods will be defined as a receiver below.
}

func Init() *Hivemind {
	return &Hivemind{}
}

func (hb *Hivemind) Start() error { //* Hivemind will initiated.

	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("ERROR! Occured at initialization phase of HIVEMIND.")
		return err
	}

	grpServer := grpc.NewServer()
	utils.PrintBanner()
	log.Println("HIVEMIND is running... HOLME is now available for current environment!")

	if err := grpServer.Serve(listener); err != nil {
		log.Fatalf("ERROR! Occured at initialization phase of HIVEMIND.")
		return err
	}

	return nil
}
