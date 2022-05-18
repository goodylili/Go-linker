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

// SetLink manages the POST requests and migrates links into the database
func SetLink(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var link models.LinkModel // creates link type
	err := json.NewDecoder(reader.Body).Decode(&link)
	link.GenerateIdentifier()
	if err != nil {
		log.Fatalln("There was a problem with the decoding on SetLink")

	}

	if strings.HasPrefix(link.Link, "http") == true && utils.ValidateURL(link.Link) == true {
		config.TheDatabase.Create(&link) // migrates the data into the database
		// outputs the data to console, this should be removed
	} else {
		link.Link = "https://" + link.Link
		if utils.ValidateURL(link.Link) == true {
			config.TheDatabase.Create(&link)
		} else {
			// error message here should be in JSON
			utils.ReturnFormatError(writer)
		}

	}

	// return successful here
}

func Ping(writer http.ResponseWriter, _ *http.Request) {
	// reformat this properly
	utils.ReturnSuccess(writer)
}
