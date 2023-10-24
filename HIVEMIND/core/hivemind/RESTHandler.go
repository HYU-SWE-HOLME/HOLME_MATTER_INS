package hivemind

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InstanceFrame struct {
	InstType int    `json:"instance"`
	User     string `json:"user"`
	Payload  string `json:"payload"`
}

type HandleFrameFunction func(string) InstanceResponse

func sendResponse(body RESTResponse, wr http.ResponseWriter) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated) //* Indicate Response Created, regardless if the request was successful.
	err := json.NewEncoder(wr).Encode(body)
	if err != nil {
		return err
	}

	return nil
}

func handleFrame(handler HandleFrameFunction, frameData string, instanceName string, wr http.ResponseWriter) InstanceResponse {
	res := handler(frameData)

	switch res.Ok {
	case SUCCESSFUL:
		{
			SuccessfulLog(fmt.Sprintf("Successfully Applied %s Settings", instanceName))

			resp := RESTResponse{SUCCESSFUL}

			err := sendResponse(resp, wr)
			if err != nil {
				ErrorLog(fmt.Sprintf("Error while responding to the server - %v", err))
			}
		}
	case NO_DEVICE:
		{
			//WarningLog("There")
		}

	}

	return res
}

func RequestHandler() *http.ServeMux {
	mux := http.NewServeMux()

	reqHandler := http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			var instanceFrame InstanceFrame
			err := json.NewDecoder(req.Body).Decode(&instanceFrame)

			if err != nil {
				panic("Error! Invalid frame received.")
			}
			instanceType := instanceFrame.InstType
			frameData := instanceFrame.Payload

			//TODO - Serialize the request in array
			//TODO - Add Info log in here.

			//* Based on instance type, unmarshal into new json code,
			//* then send the gRPC request to the dedicated instance.
			switch instanceType {
			case 1: //* Instance Light Bulb
				{
					_ = handleFrame(HandleLightBulb, frameData, "Light Bulb", wr)
					//* TODO: handling response required
				}
			case 2: //* Instance Curtain
				{
					_ = handleFrame(HandleCurtain, frameData, "Curtain", wr)
					//* TODO: handling response required
				}
			case 3: // Instance Aircon
				{
					_ = handleFrame(HandleAircon, frameData, "Aircon", wr)
					//* TODO: handling response required
				}
			case 4: // Instance Refrigerator
				{
					_ = handleFrame(HandleRefrigerator, frameData, "Refrigerator", wr)
					//* TODO: handling response required
				}
			case 5: // Instance Water Dispenser
				{
					_ = handleFrame(HandleWaterDispenser, frameData, "WaterDispenser", wr)
					//* TODO: handling response required
				}
			case 6: // Instance Television
				{
					_ = handleFrame(HandleTelevision, frameData, "Television", wr)
					//* TODO: handling response required
				}
			case 7: // Instance Soundbar
				{
					_ = handleFrame(HandleSoundbar, frameData, "Soundbar", wr)
					//* TODO: handling response required
				}
			case 8: // Instance Massage Chair
				{
					_ = handleFrame(HandleMassageChair, frameData, "MassageChair", wr)
					//* TODO: handling response required
				}
			case 9: // Instance AI Speaker
				{
					_ = handleFrame(HandleAiSpeaker, frameData, "AiSpeaker", wr)
					//* TODO: handling response required
				}
			}
		}
	})

	mux.Handle("/request", reqHandler)
	return mux
}
