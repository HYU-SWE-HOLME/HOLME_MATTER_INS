package waterDispenser

import (
	frameHandler "INS_WATERDISPENSER/src/core/handler"
	waterDispenserProto "INS_WATERDISPENSER/src/core/pbs"
	"INS_WATERDISPENSER/src/features"
	"INS_WATERDISPENSER/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const wdPort string = "10005" //* Instance WaterDispenser will use port 10005

type WaterDispenser struct {
}

func Init() *WaterDispenser {
	return &WaterDispenser{}
}

func (lb *WaterDispenser) Start() error {
	lis, err := net.Listen("tcp", ":"+wdPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	waterDispenserProto.RegisterWaterDispenserServer(grpcServer, &frameHandler.InstanceHandler{})
	//* WaterDispenser will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE WATERDISPENSER IN RUNNING...")

	terminal.ClearTerminal()                   //* Clears Terminal, starts initialization.
	features.PrintWaterDispenser(false, false) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
