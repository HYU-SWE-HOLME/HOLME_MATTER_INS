package handler

import (
	instances "HIVEMIND/instances"
	"encoding/json"
	"fmt"
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
					var lightBulb = instances.LightBulb{}
					err := json.Unmarshal([]byte(frameData), &lightBulb)
					if err != nil {
						//* ERROR!!!!
						panic("Error for handling light bulb payload.")
					}
					fmt.Println(lightBulb)
				}

			}
		}
	})

	mux.Handle("/request", reqHandler)
	return mux
}
