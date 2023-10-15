package hivemind

import (
	"encoding/json"
	"net/http"
)

type InstanceFrame struct {
	InstType int    `json:"instance"`
	Payload  string `json:"payload"`
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

			//* Based on instance type, unmarshal into new json code,
			//* then send the gRPC request to the dedicated instance.
			switch instanceType {
			case 1: //* Instance Light Bulb
				{
					HandleLightBulb(frameData)
				}

			}
		}
	})

	mux.Handle("/request", reqHandler)
	return mux
}
