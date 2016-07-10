package controllers

import (
	"encoding/json"
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"errors"
	"net/http"
	"time"

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
	case r.Method == "POST":
		businessRelationType := models.BusinessRelationType{}
		err := json.NewDecoder(r.Body).Decode(&businessRelationType)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": models.BusinessRelationType{}}}, http.StatusInternalServerError)
			return
		}
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationTypes": []models.BusinessRelationType{}}}, http.StatusInternalServerError)
			return
		}
		if businessRelationType.ID == "" {
			businessRelationType.RecCreatedByID, businessRelationType.RecModifiedByID = user.ID, user.ID
			businessRelationType.RecCreated, businessRelationType.RecModified = &models.Timestamp{time.Now()}, &models.Timestamp{time.Now()}
			businessRelationType.ClientID = user.ClientID
			businessRelationType.OrganizationID = user.OrganizationID
		} else {
			businessRelationType.RecModifiedByID = user.ID
			businessRelationType.RecModified = &models.Timestamp{time.Now()}
		}

		businessRelationType, err = models.PostBusinessRelationType(businessRelationType)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": businessRelationType}}, http.StatusInternalServerError)
		}
		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": businessRelationType}}, http.StatusOK)
	}
}

// ErrIDParameterNotFound is thrown when do not found ID in request
var ErrIDParameterNotFound = errors.New("ID Parameter Not Found")

func API_BusinessRelationType_Id(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "GET":
		ID := r.URL.Query().Get("ID")
		if ID == "" {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": models.BusinessRelationType{}}}, http.StatusInternalServerError)
			return
		}
		businessRelationType, err := models.GetBusinessRelationTypeByID(ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": models.BusinessRelationType{}}}, http.StatusInternalServerError)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: true, Data: map[string]interface{}{"BusinessRelationType": businessRelationType}, IsAuthenticated: true}, http.StatusOK)

	case r.Method == "DELETE":
	}
}
