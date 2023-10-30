package hivemind

import (
	"errors"
	"net/http"
)

func InstanceRequestHandler(instanceRequest InstanceRequest, wr http.ResponseWriter) InstanceResponse {
	instanceType := instanceRequest.InstType
	instanceData := instanceRequest.Payload

	//* Define a map of handlers for each instance type.
	handlers := map[int]func(string) InstanceResponse{
		1: HandleLightBulb,
		2: HandleCurtain,
		3: HandleAircon,
		4: HandleRefrigerator,
		5: HandleWaterDispenser,
		6: HandleTelevision,
		7: HandleSoundbar,
		8: HandleMassageChair,
		9: HandleAiSpeaker,
	}

	//* Check if the handler for the specified instance type exists.
	handler, exists := handlers[instanceType]
	if !exists {
		//* Handle the case where the instance type is not recognized.
		return InstanceResponse{error: errors.New("unknown instance type")}
	}

	//* Call the appropriate handler and return its response.
	return handler(instanceData)
}
