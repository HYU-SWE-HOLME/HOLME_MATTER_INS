package massageChair

import (
	frameHandler "INS_MASSAGECHAIR/src/core/handler"
	massageChairProto "INS_MASSAGECHAIR/src/core/pbs"
	"INS_MASSAGECHAIR/src/features"
	"INS_MASSAGECHAIR/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const MCPort string = "10008" //* Instance MassageChair will use port 10008

type MassageChair struct {
}

func Init() *MassageChair {
	return &MassageChair{}
}

func (lb *MassageChair) Start() error {
	lis, err := net.Listen("tcp", ":"+MCPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	massageChairProto.RegisterMassageChairServer(grpcServer, &frameHandler.InstanceHandler{})
	//* MassageChair will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE MASSAGECHAIR RUNNING...")

	terminal.ClearTerminal()          //* Clears Terminal, starts initialization.
	features.PrintMassageChair(false) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
