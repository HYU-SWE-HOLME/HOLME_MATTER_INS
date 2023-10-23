package aiSpeaker

import (
	frameHandler "INS_AISPEAKER/src/core/handler"
	aiSpeakerProto "INS_AISPEAKER/src/core/pbs"
	"INS_AISPEAKER/src/features"
	"INS_AISPEAKER/src/terminal"
	"log"
	"net"

	"google.golang.org/grpc"
)

const AcPort string = "10009" //* Instance AiSpeaker will use port 10009

type AiSpeaker struct {
}

func Init() *AiSpeaker {
	return &AiSpeaker{}
}

func (lb *AiSpeaker) Start() error {
	lis, err := net.Listen("tcp", ":"+AcPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()
	aiSpeakerProto.RegisterAiSpeakerServer(grpcServer, &frameHandler.InstanceHandler{})
	//* AiSpeaker will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("INSTANCE AISPEAKER RUNNING...")

	terminal.ClearTerminal()       //* Clears Terminal, starts initialization.
	features.PrintAiSpeaker(false) //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
		return err
	}

	return nil
}
