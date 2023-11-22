package hivemind

import (
	"HIVEMIND/core/constants"
	instancePb "HIVEMIND/core/instances"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
	//* TODO: Add response handler

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
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling curtain request, %v", err)
		return false
	}

	return true
}

//* Instance 3: Aircon

func SendFrameDataToAircon(
	trigger bool, // 0: off / 1: on
	mode string,
	airflowDirect bool,
	fanSpeed int,
	brightnessScreen int,
	objTemperature int, // user set temperature
	startWakeupTimer bool,
	startShutdownTimer bool,
	stopWakeupTimer bool,
	stopShutdownTimer bool,
	wakeupTime int,
	shutdownTime int,
) bool {
	conn, err := grpc.Dial(constants.INS_AIRCON, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect aircon instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing aircon instance")
			return
		}
	}(conn)

	client := instancePb.NewAirconClient(conn)

	req := &instancePb.AirconFrame{Trigger: trigger,
		Mode:               mode,
		AirflowDirect:      airflowDirect,
		FanSpeed:           int64(fanSpeed),
		BrightnessScreen:   int64(brightnessScreen),
		ObjTemperature:     int64(objTemperature), // user set temperature
		StartWakeupTimer:   startWakeupTimer,
		StartShutdownTimer: startShutdownTimer,
		StopWakeupTimer:    stopWakeupTimer,
		StopShutdownTimer:  stopShutdownTimer,
		WakeupTime:         int64(wakeupTime),
		ShutdownTime:       int64(shutdownTime),
	}

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling aircon request, %v", err)
		return false
	}

	return true
}

//* Instance 4: Refrigerator

func SendFrameDataToRefrigerator(
	trigger bool, // 0: off / 1: on
) bool {
	conn, err := grpc.Dial(constants.INS_REFRIGERATOR, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect refrigerator instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing refrigerator instance")
			return
		}
	}(conn)

	client := instancePb.NewRefrigeratorClient(conn)

	req := &instancePb.RefrigeratorFrame{Trigger: trigger}

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling refrigerator request, %v", err)
		return false
	}

	return true
}

//* Instance 5: Water Dispenser

func SendFrameDataToWaterDispenser(
	triggerReminder bool, // 0: before reminder / 1: after reminder
	triggerWater bool, // 0: before dispensing water / 1: after dispensing water
) bool {
	conn, err := grpc.Dial(constants.INS_WATERDISPENSER, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect water dispenser instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing refrigerator instance")
			return
		}
	}(conn)

	client := instancePb.NewWaterDispenserClient(conn)

	req := &instancePb.WaterDispenserFrame{TriggerReminder: triggerReminder, TriggerWater: triggerWater}

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling water dispenser request, %v", err)
		return false
	}

	return true
}

//* Instance 6: Television

func SendFrameDataToTelevision(
	trigger bool, // 0: music off / 1: music on
) bool {
	conn, err := grpc.Dial(constants.INS_TELEVISION, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect television  instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing television instance")
			return
		}
	}(conn)

	client := instancePb.NewTelevisionClient(conn)

	req := &instancePb.TelevisionFrame{Trigger: trigger}

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling television request, %v", err)
		return false
	}

	return true
}

//* Instance 7: Soundbar

func SendFrameDataToSoundbar(
	trigger bool, // 0: music off / 1: music on
) bool {
	conn, err := grpc.Dial(constants.INS_SOUNDBAR, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect soundbar instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing soundbar instance")
			return
		}
	}(conn)

	client := instancePb.NewSoundbarClient(conn)

	req := &instancePb.SoundbarFrame{Trigger: trigger}

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling soundbar request, %v", err)
		return false
	}

	return true
}

//* Instance 8: Massage Chair

func SendFrameDataToMassageChair(
	trigger bool, // 0: off / 1: on
) bool {
	conn, err := grpc.Dial(constants.INS_MASSAGECHAIR, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect massage chair instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing massage chair instance")
			return
		}
	}(conn)

	client := instancePb.NewMassageChairClient(conn)

	req := &instancePb.MassageChairFrame{Trigger: trigger}

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling massage chair request, %v", err)
		return false
	}

	return true
}

//* Instance 9: AI Speaker

func SendFrameDataToAiSpeaker(
	trigger bool, // 0: music off / 1: music on
	askForReplacement bool,
	replacement bool,
) bool {
	conn, err := grpc.Dial(constants.INS_AISPEAKER, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect AI speaker instance")
		return false
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Error while closing AI speaker instance")
			return
		}
	}(conn)

	client := instancePb.NewAiSpeakerClient(conn)

	req := &instancePb.AiSpeakerFrame{Trigger: trigger, AskForReplacement: askForReplacement, Replacement: replacement} //TODO

	_, err = client.HandleFrame(context.Background(), req)
	//* TODO: Add response handler

	if err != nil {
		log.Fatalf("Error while handling AI speaker request, %v", err)
		return false
	}

	return true
}
