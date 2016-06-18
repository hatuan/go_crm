package controllers

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"net/http"

	"github.com/gorilla/context"
)

func API_BusinessRelationTypes(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationTypes": models.BusinessRelationType{}}}, http.StatusInternalServerError)
		}

		client := new(models.Client)
		err := client.Get(requestUser.ClientID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, Data: map[string]interface{}{"Organizations": []models.Organization{}}, IsAuthenticated: true}, http.StatusInternalServerError)

		}
		organizations, err := client.GetOrganizations()
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, Data: map[string]interface{}{"Organizations": []models.Organization{}}, IsAuthenticated: true}, http.StatusInternalServerError)
		}
		JSONResponse(w, models.Response{ReturnStatus: true, Data: map[string]interface{}{"Organizations": organizations}, IsAuthenticated: true}, http.StatusOK)
	}
}
