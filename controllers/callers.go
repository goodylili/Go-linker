package controllers

import (
	"Golinker/config"
	"Golinker/models"
	"Golinker/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func SetLink(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var link models.LinkModel // creates link type
	err := json.NewDecoder(reader.Body).Decode(&link)
	link.GenerateIdentifier()
	if err != nil {
		log.Fatalln("There was a problem with the decoding on SetLink")

	}

	if strings.HasPrefix(link.Link, "http") == true && utils.ValidateURL(link.Link) == true {
		config.TheDatabase.Create(&link)

	} else {
		link.Link = "https://" + link.Link
		if utils.ValidateURL(link.Link) == true {
			config.TheDatabase.Create(&link)
		} else {
			// error message here should be in JSON
			utils.ReturnFormatError(writer)
		}

	}

	err = json.NewEncoder(writer).Encode(link)
	if err != nil {
		log.Println("There was an error encoding the output JSON", err)
	}
}

func Ping(writer http.ResponseWriter, _ *http.Request) {

	utils.ReturnSuccess(writer)
}
