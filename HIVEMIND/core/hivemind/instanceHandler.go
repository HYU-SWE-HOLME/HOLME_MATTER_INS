package hivemind

import (
	"HIVEMIND/core/constants"
	instancePb "HIVEMIND/core/instances"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type InstanceHandler struct {
	instancePb.LightBulbClient
}

func SendFrameDataToLightBulb(
	trigger bool,
	degree int,
	color string,
) bool {
	conn, err := grpc.Dial(constants.INS_LIGHTBULB, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect lightbulb instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing lightbulb instance")
			return
		}
	}(conn)

	client := instancePb.NewLightBulbClient(conn)

	//* test
	req := &instancePb.LightBulbFrame{Trigger: trigger, Degree: int64(degree), Color: color}

	res, err := client.HandleFrame(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while handling lightbulb request, %v", err)
		return false
	}
	fmt.Println(res)

	return true
}
