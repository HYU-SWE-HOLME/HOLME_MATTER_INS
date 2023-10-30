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
			var resp RESTResponse
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

			err = sendResponse(resp, wr)
			if err != nil {
				ErrorLog(fmt.Sprintf("Error while responding to the server - %v", err))
			}

		}
	})

	mux.Handle("/request", reqHandler)
	return mux
}
