package handlers

import (
	"fmt"
	"net/http"
	"post-service/web/utils"
)

func (handlers *Handlers) GetLocation(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	locationTitle := queryParams.Get("title")

	if locationTitle == "" {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("required title"))
		return
	}

	locations, err := handlers.locSvc.GetLocation(r.Context(), locationTitle)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, locations)
}
