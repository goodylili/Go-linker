package utils

import (
	"Golinker/config"
	"Golinker/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func FetchHandleRequestsRedirects(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var link models.LinkModel

	// handle error here
	config.TheDatabase.First(&link, params["id"])
	_, err := fmt.Fprintf(writer, "Redirecting you to %s", link.Link)
	if err != nil {
		return
	}
	http.Redirect(writer, request, link.Link, http.StatusSeeOther)
}
