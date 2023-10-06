package main

import (
	inspb "Test/config"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT string = "9000"

type insServer struct {
	inspb.ConfigReqRespServer
}

func (serv *insServer) TestResponse(ctx context.Context, req *inspb.ConfigRequest) (resp *inspb.ConfigResponse, error error) {

	return &inspb.ConfigResponse{
		Payload: "Response Successful!",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+PORT)

	if err != nil { //* Error!
		log.Fatalf("ERROR! Failed to listen server: %v", err)
	}

	grpcServer := grpc.NewServer()
	inspb.RegisterConfigReqRespServer(grpcServer, &insServer{}) //* gRPC Register.
	log.Printf("💫Instnace is now working! Available at port number %s 💫", PORT)

	if grpcErr := grpcServer.Serve(lis); grpcErr != nil {
		log.Fatalf("ERROR! Failed to serve instance: %v", grpcErr)
	}
}
