package refrigerator

import (
	frameHandler "INS_REFRIGERATOR/src/core/handler"
	refrigeratorProto "INS_REFRIGERATOR/src/core/pbs"
	"INS_REFRIGERATOR/src/features"
	"INS_REFRIGERATOR/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const RfPort string = "10004" //* Instance Refrigerator will use port 10002

type Refrigerator struct {
}

func Init() *Refrigerator {
	return &Refrigerator{}
}

func (lb *Refrigerator) Start() error {
	lis, err := net.Listen("tcp", ":"+RfPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	refrigeratorProto.RegisterRefrigeratorServer(grpcServer, &frameHandler.InstanceHandler{})
	//* Refrigerator will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE REFRIGERATOR IN RUNNING...")

	terminal.ClearTerminal()         //* Clears Terminal, starts initialization.
	features.PrintRefrigerator(true) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
