package controllers

import (
	"encoding/json"
	"net/http"
	"erpvietnam/crm/models"
	"fmt"
	"erpvietnam/crm/log"
)

//InitializeApplication run init menu ... before login
func InitializeApplication(w http.ResponseWriter, r *http.Request) {
	applicationModel := new (models.ApplicationModelDTO)

	//TODO: Implement InitializeApplication here

	JSONResponse(w, applicationModel, http.StatusOK)
}

// JSONResponse attempts to set the status code, c, and marshal the given interface, d, into a response that
// is written to the given ResponseWriter.
func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}