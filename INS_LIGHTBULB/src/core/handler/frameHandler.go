package handler

import (
	InstanceLightBulb "INS_LIGHTBULB/src/core/pbs"
	"INS_LIGHTBULB/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceLightBulb.LightBulbServer
}

// * Map color in flag value
var colorMap = map[string]uint8{
	"white":  0,
	"yellow": 1,
}

//HandleFrame
/*
- It will handle the incoming frame, parse the data and handle the light bulb based on the data.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceLightBulb.LightBulbFrame) (*InstanceLightBulb.LightBulbRes, error) {
	//* Handle Incoming Frame.
	trigger, degree, color := frame.Trigger, frame.Degree, frame.Color

	if !trigger {
		features.PrintLightDisabled() //* The light is off; print the default disabled light.
	} else {
		features.PrintLightEnabled(int(degree), colorMap[color]) //* The light in on; print the dedicated values.
	}

	return &InstanceLightBulb.LightBulbRes{
		Status: true,
		Error:  "",
	}, nil
}
