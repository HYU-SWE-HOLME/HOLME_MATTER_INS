package handler

import (
	lightBulb "INS_LIGHTBULB/src/core"
	"fmt"
	"golang.org/x/net/context"
)

type InstanceHandler struct {
	lightBulb.LightBulbServer
}

func (handler *InstanceHandler) HandleFrame(ctx context.Context, frame *lightBulb.LightBulbFrame) (*lightBulb.LightBulbRes, error) {
	//* Handle Incoming Frame.
	trigger, degree, color := frame.Trigger, frame.Degree, frame.Color
	fmt.Println(trigger)
	fmt.Println(degree)
	fmt.Println(color)

	return &lightBulb.LightBulbRes{}, nil
}
