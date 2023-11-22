package hivemind

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var instanceName = map[int]string{
	1: "Light Bulb",
	2: "Curtain",
	3: "AC",
	4: "Refrigerator",
	5: "Water Dispenser",
	6: "Television",
	7: "Sound Bar",
	8: "Massage Chair",
	9: "AI Speaker",
}

func sendSyncResponse(body RESTSyncResponse, wr http.ResponseWriter) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated) //* Indicate Response Created, regardless if the request was successful.
	err := json.NewEncoder(wr).Encode(body)
	if err != nil {
		return err
	}

	return nil
}

func sendPingResponse(body RESTPingResponse, wr http.ResponseWriter) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusCreated) //* Indicate Response Created, regardless if the request was successful.
	err := json.NewEncoder(wr).Encode(body)
	if err != nil {
		return err
	}

	return nil
}

func SyncRequestHandler() http.HandlerFunc {
	reqHandler := http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:

			var instanceFrame InstanceFrame
			var resp RESTSyncResponse
			err := json.NewDecoder(req.Body).Decode(&instanceFrame)

			if err != nil {
				panic("Error! Invalid frame received.")
			}

			requests := instanceFrame.Request

			for _, request := range requests {
				res := InstanceRequestHandler(request, wr)

				switch res.Ok {
				case SUCCESSFUL:
					{
						SuccessfulLog(fmt.Sprintf("Successfully Applied %s Settings", instanceName[request.InstType]))
						resp.Result = append(resp.Result, res)
					}
				case NO_DEVICE:
					{
						//WarningLog("There")
					}

				}
			}

			err = sendSyncResponse(resp, wr)
			if err != nil {
				ErrorLog(fmt.Sprintf("Error while responding to the server - %v", err))
			}

		}
	})

	return reqHandler
}

func PingRequestHandler() http.HandlerFunc {
	reqHandler := http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:

			var pingRequest PingRequest
			var resp RESTPingResponse
			err := json.NewDecoder(req.Body).Decode(&pingRequest)

			// DEMO: Assume there is all instances exists.
			numInstance := int(AI_SPEAKER) + 1
			instStatus := make([]PingResponse, numInstance)

			for i := range instStatus {
				instStatus[i] = PingResponse{
					InstanceId: i + 1,
					IsExist:    true, // All connected in demo.
				}
			}

			resp.Result = instStatus

			err = sendPingResponse(resp, wr)
			if err != nil {
				ErrorLog(fmt.Sprintf("Error while responding to the server - %v", err))
			}

		}
	})

	return reqHandler
}

func RequestHandler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/request", SyncRequestHandler())
	mux.Handle("/ping", PingRequestHandler())
	return mux
}
