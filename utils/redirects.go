package utils

import (
	"Golinker/config"
	"Golinker/models"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func FetchHandleRequestsRedirects(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var link models.LinkModel
	exist := config.TheDatabase.Where("Identifier=?", params["link"]).First(&link)
	if !errors.Is(exist.Error, gorm.ErrRecordNotFound) && exist.Error != nil {
		ReturnNotFound(writer)
	}

	// figure out how to log out a response
	http.Redirect(writer, request, link.Link, http.StatusSeeOther)
}
