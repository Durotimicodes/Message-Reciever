package main

import (
	"encoding/json"
	"net/http"
)


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
		payLoad := jsonResponse{
			Error: false,
			Message: "Hit the broker",
		}


	//Marshal the payload data
	byt,_:=json.MarshalIndent(payLoad, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(byt)
}
