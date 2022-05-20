package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ReturnSuccess edit the bytes messages
func ReturnSuccess(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	var buffer bytes.Buffer
	buffer.WriteString(`{Response: "success", Message: "Your Request was successful"}`)
	err := json.NewEncoder(writer).Encode(buffer.String())
	if err != nil {
		return
	}

}

func ReturnFormatError(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	var buffer bytes.Buffer
	buffer.WriteString(`{Response: "Error", Message: "Your request isn't properly formatted'"}`)
	err := json.NewEncoder(writer).Encode(buffer.String())
	if err != nil {
		return
	}

}

func ReturnNotFound(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	var buffer bytes.Buffer
	buffer.WriteString(`{Response: "Error", Message: "You must provide a valid URL"}`)
	err := json.NewEncoder(writer).Encode(buffer.String())
	if err != nil {
		return
	}
}
