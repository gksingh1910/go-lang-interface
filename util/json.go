package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, statusCode int16, payload interface{}) {
	fmt.Printf("\n payload is %v", payload)

	msg, err := json.Marshal(payload)
	fmt.Printf("\n payload is %v", payload)
	if err != nil {
		fmt.Printf("\n Couldn't marshal as json")
	}
	w.WriteHeader(int(statusCode))
	w.Write(msg)

}

func RespondWithError(w http.ResponseWriter, statusCode int16, message string) {
	if statusCode > 500 {
		fmt.Errorf("Respond  with 5xx error code %v", message)
	}
	type errorMessage struct {
		errorMsg   string
		statusCode int16
	}
	payload := errorMessage{}

	payload.errorMsg = message
	payload.statusCode = statusCode
	RespondWithJson(w, statusCode, payload)
}
