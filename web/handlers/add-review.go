package handlers

import "net/http"
type review struct{
	locationId int `json:"locationId"`
	
}
func(handlers *Handlers)AddReview(w http.ResponseWriter,r *http.Request){

}