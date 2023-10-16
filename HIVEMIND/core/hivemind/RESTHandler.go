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

func sendResponse(body RESTResponse, wr http.ResponseWriter) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated) //* Indicate Response Created, regardless if the request was successful.
	err := json.NewEncoder(wr).Encode(body)
	if err != nil {
		return err
	}

	return nil
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
					res := HandleLightBulb(frameData)

					switch res.Ok {
					case SUCCESSFUL:
						{
							SuccessfulLog("Successfully Applied Light Settings")

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
				}

			}
		}
	})

	mux.Handle("/request", reqHandler)
	return mux
}
