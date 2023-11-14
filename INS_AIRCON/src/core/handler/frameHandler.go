package handler

import (
	InstanceAircon "INS_AIRCON/src/core/pbs"
	"INS_AIRCON/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceAircon.AirconServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceAircon.AirconFrame) (*InstanceAircon.AirconRes, error) {

	trigger, mode, airflowDirect, fanSpeed, brightnessScreen, objTemperature, startWakeupTimer, startShutdownTimer, stopWakeupTimer, stopShutdownTimer, wakeupTime, shutdownTime := frame.Trigger, frame.Mode, frame.AirflowDirect, frame.FanSpeed, frame.BrightnessScreen, frame.ObjTemperature, frame.StartWakeupTimer, frame.StartShutdownTimer, frame.StopWakeupTimer, frame.StopShutdownTimer, frame.WakeupTime, frame.ShutdownTime

	if trigger { // on
		features.PrintAirconOn(true, mode, airflowDirect, int(fanSpeed),
			int(brightnessScreen), int(objTemperature), startWakeupTimer, startShutdownTimer,
			stopWakeupTimer, stopShutdownTimer, int(wakeupTime), int(shutdownTime))
	} else { // off
		features.PrintAirconOff()
	}
	return &InstanceAircon.AirconRes{Status: true, Error: ""}, nil
}
