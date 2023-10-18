package handler

import (
	InstanceWaterDispenser "INS_WATERDISPENSER/src/core/pbs"
	"INS_WATERDISPENSER/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceWaterDispenser.WaterDispenserServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceWaterDispenser.WaterDispenserFrame) (*InstanceWaterDispenser.WaterDispenserRes, error) {
	triggerReminder, triggerWater := frame.TriggerReminder, frame.TriggerWater
	if !triggerReminder { // before medication reminder
		features.PrintWaterDispenser(false, false)
	} else if !triggerWater { // after medication reminder, before user says "give me water"
		features.PrintWaterDispenser(true, false)
	} else { // after medication reminder, after user says "give me water"
		features.PrintWaterDispenser(true, true)
	}
	return &InstanceWaterDispenser.WaterDispenserRes{Status: true, Error: ""}, nil
}
