package aircon

import (
	frameHandler "INS_AIRCON/src/core/handler"
	airconProto "INS_AIRCON/src/core/pbs"
	"INS_AIRCON/src/features"
	"INS_AIRCON/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const AcPort string = "10003" //* Instance Aircon will use port 10003

type Aircon struct {
}

func Init() *Aircon {
	return &Aircon{}
}

func (lb *Aircon) Start() error {
	lis, err := net.Listen("tcp", ":"+AcPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	airconProto.RegisterAirconServer(grpcServer, &frameHandler.InstanceHandler{})
	//* Aircon will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE AIRCON IN RUNNING...")

	terminal.ClearTerminal()       //* Clears Terminal, starts initialization.
	features.PrintAirconOff(-1, 0) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
