package main

import (
	"log-service/data"
	"net/http"
)

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read the json
	var requestPayload jsonPayload

	_ = app.readJSON(w, r, &requestPayload)

	entry := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(entry)
	if err != nil {
		app.errorJSON(w, err)
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
