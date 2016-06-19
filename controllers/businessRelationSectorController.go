package controllers

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"net/http"

	"github.com/gorilla/context"
)

func API_BusinessRelationSectors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSectors": []models.BusinessRelationSector{}}}, http.StatusInternalServerError)
			return
		}

		businessRelationSectors, err := models.GetBusinessRelationSectors(user.OrganizationID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSectors": []models.BusinessRelationSector{}}}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: true, Data: map[string]interface{}{"BusinessRelationSectors": businessRelationSectors}, IsAuthenticated: true}, http.StatusOK)
	}
}
