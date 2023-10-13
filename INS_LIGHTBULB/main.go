/*
- Matter Instance: Light bulb
*/

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	lightBulb "INS_LIGHTBULB/src/core"
	frameHandler "INS_LIGHTBULB/src/core/handler"
	"INS_LIGHTBULB/src/features"
	"INS_LIGHTBULB/src/terminal"
)

const LbPort string = "10001" //* Instance Light Bulb will use port 10001

func main() {
	lis, err := net.Listen("tcp", ":"+LbPort)

	if err != nil {
		log.Fatalf("Failed to initate features: %v", err)
	}

	grpcServer := grpc.NewServer()
	lightBulb.RegisterLightBulbServer(grpcServer, &frameHandler.InstanceHandler{})
	//* Light bulb will use grpcServer, initialize it with frame handler.
	log.Printf("Starting Instance...")
	log.Printf("💡💡INSTANCE LIGHTBULB IN RUNNING...💡💡")

	terminal.ClearTerminal()      //* Clears Terminal, starts initialization.
	features.PrintLightDisabled() //* Prints the default images as a initial state.
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate features: %v", err)
	}
}
