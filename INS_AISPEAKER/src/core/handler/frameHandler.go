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
	trigger, askForReplacement, replacement := frame.Trigger, frame.AskForReplacement, frame.Replacement
	features.PrintAiSpeaker(trigger, askForReplacement, replacement)
	return &InstanceAiSpeaker.AiSpeakerRes{Status: true, Error: ""}, nil
}
