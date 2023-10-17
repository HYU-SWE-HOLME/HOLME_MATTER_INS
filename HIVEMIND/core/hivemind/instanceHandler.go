package hivemind

import (
	"HIVEMIND/core/constants"
	instancePb "HIVEMIND/core/instances"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type InstanceHandler struct {
	instancePb.LightBulbClient
}

//* Instance 1. Light Bulb

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

	req := &instancePb.LightBulbFrame{Trigger: trigger, Degree: int64(degree), Color: color}

	_, err = client.HandleFrame(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while handling lightbulb request, %v", err)
		return false
	}

	return true
}

//* Instance 2: Curtain

func SendFrameDataToCurtain(
	isHorizontal bool, //* true: horizontal, false: vertical
	isCentreMode bool, //* true: centre, false: side
	isLeft bool, //* true: left -> right, false: top -> down
	degree int64, //* degree value
) bool {
	conn, err := grpc.Dial(constants.INS_CURTAIN, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect curtain instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing curtain instance")
			return
		}
	}(conn)

	client := instancePb.NewCurtainClient(conn)

	req := &instancePb.CurtainFrame{IsHorizontal: isHorizontal, IsCenterMode: isCentreMode, IsLeftOrTop: isLeft, Degree: degree}

	_, err = client.HandleFrame(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while handling curtain request, %v", err)
		return false
	}

	return true
}
