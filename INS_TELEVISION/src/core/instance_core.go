package television

import (
	frameHandler "INS_TELEVISION/src/core/handler"
	televisionProto "INS_TELEVISION/src/core/pbs"
	"INS_TELEVISION/src/features"
	"INS_TELEVISION/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const TvPort string = "10006" //* Instance Television will use port 10006

type Television struct {
}

func Init() *Television {
	return &Television{}
}

func (lb *Television) Start() error {
	lis, err := net.Listen("tcp", ":"+TvPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	televisionProto.RegisterTelevisionServer(grpcServer, &frameHandler.InstanceHandler{})
	//* Television will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE TELEVISION RUNNING...")

	terminal.ClearTerminal()        //* Clears Terminal, starts initialization.
	features.PrintTelevision(false) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
