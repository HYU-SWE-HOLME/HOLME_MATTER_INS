package handler

import (
	InstanceSoundbar "INS_SOUNDBAR/src/core/pbs"
	"INS_SOUNDBAR/src/features"
	"context"
)

type InstanceHandler struct {
	InstanceSoundbar.SoundbarServer
}

//HandleFrame
/*
- Handles the incoming frame, parse the data, and handles the instance.
*/
func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *InstanceSoundbar.SoundbarFrame) (*InstanceSoundbar.SoundbarRes, error) {
	trigger := frame.Trigger
	if trigger { // on
		features.PrintSoundbar(true)
	} else { // off
		features.PrintSoundbar(false)
	}
	return &InstanceSoundbar.SoundbarRes{Status: true, Error: ""}, nil
}
