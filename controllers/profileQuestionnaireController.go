package controllers

import (
	"encoding/json"
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"

	"net/http"
	"strings"
	"time"

	"github.com/gorilla/context"
)

func API_ProfileQuestionnaires(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaires": []models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}

		infiniteScrollingInformation := models.InfiniteScrollingInformation{
			After:          r.URL.Query().Get("After"),
			FetchSize:      r.URL.Query().Get("FetchSize"),
			SortDirection:  r.URL.Query().Get("SortDirection"),
			SortExpression: r.URL.Query().Get("SortExpression")}

		profileQuestionnaires, tranInfor := models.GetProfileQuestionnaireHeaders(user.OrganizationID, r.URL.Query().Get("Search"), infiniteScrollingInformation)
		if tranInfor.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaires": []models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: true, TotalRows: len(profileQuestionnaires), Data: map[string]interface{}{"ProfileQuestionnaires": profileQuestionnaires}, IsAuthenticated: true}, http.StatusOK)

	case r.Method == "POST":
		profileQuestionnaire := models.ProfileQuestionnaireHeader{}
		err := json.NewDecoder(r.Body).Decode(&profileQuestionnaire)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": []models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}
		if profileQuestionnaire.ID == "" {
			profileQuestionnaire.RecCreatedByID = user.ID
			profileQuestionnaire.RecModifiedByID = user.ID
			profileQuestionnaire.RecCreated = &models.Timestamp{time.Now()}
			profileQuestionnaire.RecModified = &models.Timestamp{time.Now()}
			profileQuestionnaire.ClientID = user.ClientID
			profileQuestionnaire.OrganizationID = user.OrganizationID
		} else {
			profileQuestionnaire.RecModifiedByID = user.ID
			profileQuestionnaire.RecModified = &models.Timestamp{time.Now()}
		}

		profileQuestionnaire, tranInfor := models.PostProfileQuestionnaireHeader(profileQuestionnaire)
		if tranInfor.ReturnStatus == false && len(tranInfor.ValidationErrors) > 0 {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, ValidationErrors: tranInfor.ValidationErrors, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": profileQuestionnaire}}, http.StatusBadRequest)
			return
		} else if tranInfor.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": profileQuestionnaire}}, http.StatusBadRequest)
			return
		}

		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": profileQuestionnaire}}, http.StatusOK)

	case r.Method == "DELETE":
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true}, http.StatusBadRequest)
			return
		}
		ids := strings.Split(r.URL.Query().Get("ID"), ",")
		tranInfo := models.DeleteProfileQuestionnaireHeaderById(user.OrganizationID, ids)
		if tranInfo.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true}, http.StatusOK)
	}
}

func API_ProfileQuestionnaire_Id(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "GET":
		ID := r.URL.Query().Get("ID")
		if ID == "" {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}
		profileQuestionnaire, tranInfo := models.GetProfileQuestionnaireHeaderByID(ID)
		if !tranInfo.ReturnStatus {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, Data: map[string]interface{}{"ProfileQuestionnaire": profileQuestionnaire}, IsAuthenticated: true}, http.StatusOK)
	}
}
