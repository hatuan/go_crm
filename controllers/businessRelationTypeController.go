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
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationTypes": []models.BusinessRelationType{}}}, http.StatusInternalServerError)
			return
		}

		businessRelationTypes, err := models.GetBusinessRelationTypes(user.OrganizationID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationTypes": []models.BusinessRelationType{}}}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: true, Data: map[string]interface{}{"BusinessRelationTypes": businessRelationTypes}, IsAuthenticated: true}, http.StatusOK)
	}
}
