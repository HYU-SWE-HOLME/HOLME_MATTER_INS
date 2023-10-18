package handler

import (
	InstanceTelevision "INS_TELEVISION/src/core/pbs"
	"INS_TELEVISION/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceTelevision.TelevisionServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceTelevision.TelevisionFrame) (*InstanceTelevision.TelevisionRes, error) {
	trigger := frame.Trigger
	if trigger { // on
		features.PrintTelevision(true)
	} else { // off
		features.PrintTelevision(false)
	}
	return &InstanceTelevision.TelevisionRes{Status: true, Error: ""}, nil
}
