package soundbar

import (
	frameHandler "INS_SOUNDBAR/src/core/handler"
	soundbarProto "INS_SOUNDBAR/src/core/pbs"
	"INS_SOUNDBAR/src/features"
	"INS_SOUNDBAR/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const SBPort string = "10007" //* Instance Soundbar will use port 10007

type Soundbar struct {
}

func Init() *Soundbar {
	return &Soundbar{}
}

func (lb *Soundbar) Start() error {
	lis, err := net.Listen("tcp", ":"+SBPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	soundbarProto.RegisterSoundbarServer(grpcServer, &frameHandler.InstanceHandler{})
	//* Soundbar will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE SOUNDBAR RUNNING...")

	terminal.ClearTerminal()      //* Clears Terminal, starts initialization.
	features.PrintSoundbar(false) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
