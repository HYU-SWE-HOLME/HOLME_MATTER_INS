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

func HandleAircon(frameData string) InstanceResponse {
	var aircon = instances.Aircon{}
	err := json.Unmarshal([]byte(frameData), &aircon)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToAircon(aircon.Trigger, aircon.Temperature, aircon.WindDegree); !res {
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

func HandleRefrigerator(frameData string) InstanceResponse {
	var refrigerator = instances.Refrigerator{}
	err := json.Unmarshal([]byte(frameData), &refrigerator)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToRefrigerator(refrigerator.Trigger); !res {
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

func HandleWaterDispenser(frameData string) InstanceResponse {
	var waterDispenser = instances.WaterDispenser{}
	err := json.Unmarshal([]byte(frameData), &waterDispenser)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToWaterDispenser(waterDispenser.TriggerReminder, waterDispenser.TriggerWater); !res {
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

func HandleTelevision(frameData string) InstanceResponse {
	var television = instances.Television{}
	err := json.Unmarshal([]byte(frameData), &television)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToTelevision(television.Trigger); !res {
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

func HandleSoundbar(frameData string) InstanceResponse {
	var soundbar = instances.Soundbar{}
	err := json.Unmarshal([]byte(frameData), &soundbar)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToSoundbar(soundbar.Trigger); !res {
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

func HandleMassageChair(frameData string) InstanceResponse {
	var massageChair = instances.MassageChair{}
	err := json.Unmarshal([]byte(frameData), &massageChair)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToMassageChair(massageChair.Trigger); !res {
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

func HandleAiSpeaker(frameData string) InstanceResponse {
	var aiSpeaker = instances.AiSpeaker{}
	err := json.Unmarshal([]byte(frameData), &aiSpeaker)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
			err,
		}
	}
	if res := SendFrameDataToAiSpeaker(aiSpeaker.Trigger, aiSpeaker.ReplacementMsg); !res {
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
