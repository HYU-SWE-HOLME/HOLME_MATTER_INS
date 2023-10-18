package handler

import (
	InstanceRefrigerator "INS_REFRIGERATOR/src/core/pbs"
	"INS_REFRIGERATOR/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceRefrigerator.RefrigeratorServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceRefrigerator.RefrigeratorFrame) (*InstanceRefrigerator.RefrigeratorRes, error) {
	trigger := frame.Trigger
	if trigger { // on
		features.PrintRefrigerator(true)
	} else { // off
		features.PrintRefrigerator(false)
	}
	return &InstanceRefrigerator.RefrigeratorRes{Status: true, Error: ""}, nil
}
