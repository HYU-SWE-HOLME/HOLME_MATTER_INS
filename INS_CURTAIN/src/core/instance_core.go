package curtain

import (
	frameHandler "INS_CURTAIN/src/core/handler"
	curtainProto "INS_CURTAIN/src/core/pbs"
	"INS_CURTAIN/src/features"
	"INS_CURTAIN/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const CtPort string = "10002" //* Instance Curtain will use port 10002

type Curtain struct {
}

func Init() *Curtain {
	return &Curtain{}
}

func (lb *Curtain) Start() error {
	lis, err := net.Listen("tcp", ":"+CtPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	curtainProto.RegisterCurtainServer(grpcServer, &frameHandler.InstanceHandler{})
	//* Curtain will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE CURTAIN RUNNING...")

	terminal.ClearTerminal()   //* Clears Terminal, starts initialization.
	features.PrintCurtain(100) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
