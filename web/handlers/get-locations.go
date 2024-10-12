package handlers

import (
	"net/http"
)

func (handlers *Handlers) GetLocations(w http.ResponseWriter, r *http.Request) {
	// locations, err := handlers.locSvc.GetLocations(r.Context())
	// if err != nil {
	// 	utils.SendError(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// utils.SendData(w, locations)
}
