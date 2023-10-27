package handler

import (
	InstanceAiSpeaker "INS_AISPEAKER/src/core/pbs"
	"INS_AISPEAKER/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceAiSpeaker.AiSpeakerServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceAiSpeaker.AiSpeakerFrame) (*InstanceAiSpeaker.AiSpeakerRes, error) {
	trigger, replacementMsg := frame.Trigger, frame.ReplacementMsg
	if trigger { // on
		features.PrintAiSpeaker(true, replacementMsg)
	} else { // off
		features.PrintAiSpeaker(false, replacementMsg)
	}
	return &InstanceAiSpeaker.AiSpeakerRes{Status: true, Error: ""}, nil
}
