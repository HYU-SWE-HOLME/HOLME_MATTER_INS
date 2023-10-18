package handler

import (
	InstanceMassageChair "INS_MASSAGECHAIR/src/core/pbs"
	"INS_MASSAGECHAIR/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceMassageChair.MassageChairServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceMassageChair.MassageChairFrame) (*InstanceMassageChair.MassageChairRes, error) {
	trigger := frame.Trigger
	if trigger { // on
		features.PrintMassageChair(true)
	} else { // off
		features.PrintMassageChair(false)
	}
	return &InstanceMassageChair.MassageChairRes{Status: true, Error: ""}, nil
}
