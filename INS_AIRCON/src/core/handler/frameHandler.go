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
	trigger, temperature, degree := frame.Trigger, frame.Temperature, frame.WindDegree
	if trigger { // on
		features.PrintAirconOn(int(temperature), int(degree))
	} else { // off
		features.PrintAirconOff(int(temperature), int(degree))
	}
	return &InstanceAircon.AirconRes{Status: true, Error: ""}, nil
}
