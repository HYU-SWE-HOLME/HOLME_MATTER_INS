package hivemind

import (
	"HIVEMIND/instances"
	"encoding/json"
)

func HandleLightBulb(frameData string) InstanceResponse {
	var lightBulb = instances.LightBulb{}
	err := json.Unmarshal([]byte(frameData), &lightBulb)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToLightBulb(lightBulb.Trigger, lightBulb.Degree, lightBulb.Color); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
			nil,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
		nil,
	}
}

func HandleCurtain(frameData string) InstanceResponse {
	var curtain = instances.Curtain{}
	err := json.Unmarshal([]byte(frameData), &curtain)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToCurtain(curtain.IsHorizontal, curtain.IsCenterMode, curtain.IsLeftOrTop, int64(curtain.Degree)); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
			nil,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
		nil,
	}
}
