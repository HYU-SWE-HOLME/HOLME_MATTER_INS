/*
- Matter Instance: Light bulb
*/

package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const LbPort string = "10001"

func main() {
	lis, err := net.Listen("tcp", ":"+LbPort)

	if err != nil {
		log.Fatalf("Failed to initate instance: %v", err)
	}

	grpcServer := grpc.NewServer()
	log.Printf("Starting Instance...")
	log.Printf("💡💡INSTANCE LIGHTBULB IN RUNNING...💡💡")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initate instance: %v", err)
	}
}
