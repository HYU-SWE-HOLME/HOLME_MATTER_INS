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
		}
	}
	if res := SendFrameDataToLightBulb(lightBulb.Trigger, lightBulb.Degree, lightBulb.Color); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleCurtain(frameData string) InstanceResponse {
	var curtain = instances.Curtain{}
	err := json.Unmarshal([]byte(frameData), &curtain)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToCurtain(curtain.IsHorizontal, curtain.IsCenterMode, curtain.IsLeftOrTop, int64(curtain.Degree)); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleAircon(frameData string) InstanceResponse {
	var aircon = instances.Aircon{}
	err := json.Unmarshal([]byte(frameData), &aircon)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToAircon(
		aircon.Trigger,
		aircon.Mode,
		aircon.AirflowDirect,
		aircon.FanSpeed,
		aircon.BrightnessScreen,
		aircon.ObjTemperature,
		aircon.StartWakeupTimer,
		aircon.StartShutdownTimer,
		aircon.StopWakeupTimer,
		aircon.StopShutdownTimer,
		aircon.WakeupTime,
		aircon.ShutdownTime,
	); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleRefrigerator(frameData string) InstanceResponse {
	var refrigerator = instances.Refrigerator{}
	err := json.Unmarshal([]byte(frameData), &refrigerator)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToRefrigerator(refrigerator.Trigger); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleWaterDispenser(frameData string) InstanceResponse {
	var waterDispenser = instances.WaterDispenser{}
	err := json.Unmarshal([]byte(frameData), &waterDispenser)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToWaterDispenser(waterDispenser.TriggerReminder, waterDispenser.TriggerWater); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleTelevision(frameData string) InstanceResponse {
	var television = instances.Television{}
	err := json.Unmarshal([]byte(frameData), &television)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToTelevision(television.Trigger); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleSoundbar(frameData string) InstanceResponse {
	var soundbar = instances.Soundbar{}
	err := json.Unmarshal([]byte(frameData), &soundbar)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToSoundbar(soundbar.Trigger); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleMassageChair(frameData string) InstanceResponse {
	var massageChair = instances.MassageChair{}
	err := json.Unmarshal([]byte(frameData), &massageChair)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToMassageChair(massageChair.Trigger); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}

func HandleAiSpeaker(frameData string) InstanceResponse {
	var aiSpeaker = instances.AiSpeaker{}
	err := json.Unmarshal([]byte(frameData), &aiSpeaker)
	if err != nil {
		//* ERROR!!!!
		return InstanceResponse{
			ERROR,
		}
	}
	if res := SendFrameDataToAiSpeaker(aiSpeaker.Trigger, aiSpeaker.AskForReplacement, aiSpeaker.Replacement); !res {
		//* Substitution Required.
		return InstanceResponse{
			NO_DEVICE,
		}
	}

	//* IoT Setting Applied.
	return InstanceResponse{
		SUCCESSFUL,
	}
}
