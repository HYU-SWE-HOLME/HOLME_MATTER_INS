package handler

import (
	InstanceCurtain "INS_CURTAIN/src/core/pbs"
	"INS_CURTAIN/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceCurtain.CurtainServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceCurtain.CurtainFrame) (*InstanceCurtain.CurtainRes, error) {
	degree := frame.Degree

	if degree < 50 {
		features.PrintClosedCurtain()

	} else {
		features.PrintOpenedCurtain()
	}

	return &InstanceCurtain.CurtainRes{}, nil
}
