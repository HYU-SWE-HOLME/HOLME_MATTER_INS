package hivemind

import (
	"HIVEMIND/instances"
	"encoding/json"
)

func HandleLightBulb(frameData string) {
	var lightBulb = instances.LightBulb{}
	err := json.Unmarshal([]byte(frameData), &lightBulb)
	if err != nil {
		//* ERROR!!!!
		panic("Error for handling light bulb payload.")
	}
	SendFrameDataToLightBulb(lightBulb.Trigger, lightBulb.Degree, lightBulb.Color)
}
