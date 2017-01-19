package controllers

import (
	"encoding/json"
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"strconv"

	"net/http"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func API_ProfileQuestionnaires(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		user, err := models.GetUser(*requestUser.ID)
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
		user, err := models.GetUser(*requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaire": []models.ProfileQuestionnaireHeader{}}}, http.StatusBadRequest)
			return
		}
		if profileQuestionnaire.ID == nil {
			profileQuestionnaire.RecCreatedByID = *user.ID
			profileQuestionnaire.RecModifiedByID = *user.ID
			profileQuestionnaire.RecCreated = &models.Timestamp{time.Now()}
			profileQuestionnaire.RecModified = &models.Timestamp{time.Now()}
			profileQuestionnaire.ClientID = user.ClientID
			profileQuestionnaire.OrganizationID = user.OrganizationID
		} else {
			profileQuestionnaire.RecModifiedByID = *user.ID
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
		user, err := models.GetUser(*requestUser.ID)
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
		ID, err := strconv.ParseInt(r.URL.Query().Get("ID"), 10, 64)
		if err != nil {
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

func API_ProfileQuestionnaireLines_By_HeaderId(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		params := mux.Vars(r)
		HeaderID, err := strconv.ParseInt(params["headerid"], 10, 64)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}
		profileQuestionnaireLines, tranInfo := models.GetProfileQuestionnaireLinesByHeaderID(HeaderID)
		if !tranInfo.ReturnStatus {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, Data: map[string]interface{}{"ProfileQuestionnaireLines": profileQuestionnaireLines}, IsAuthenticated: true}, http.StatusOK)
	case r.Method == "POST":
		params := mux.Vars(r)
		HeaderID, err := strconv.ParseInt(params["headerid"], 10, 64)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}

		profileQuestionnaireLines := make([]models.ProfileQuestionnaireLine, 0)
		err = json.NewDecoder(r.Body).Decode(&profileQuestionnaireLines)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}
		user, err := models.GetUser(*requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}

		//range uses a[i] as its second value for arrays/slices, which effectively means that the value is copied, making the original value untouchable.
		for index, _ := range profileQuestionnaireLines {
			if profileQuestionnaireLines[index].ID == nil {
				profileQuestionnaireLines[index].Version = 1
				profileQuestionnaireLines[index].RecCreatedByID = *user.ID
				profileQuestionnaireLines[index].RecModifiedByID = *user.ID
				profileQuestionnaireLines[index].RecCreated = &models.Timestamp{time.Now()}
				profileQuestionnaireLines[index].RecModified = &models.Timestamp{time.Now()}
				profileQuestionnaireLines[index].ClientID = user.ClientID
				profileQuestionnaireLines[index].OrganizationID = user.OrganizationID
			} else {
				profileQuestionnaireLines[index].RecModifiedByID = *user.ID
				profileQuestionnaireLines[index].RecModified = &models.Timestamp{time.Now()}
			}
		}

		profileQuestionnaireLines, tranInfor := models.PostProfileQuestionnaireLines(HeaderID, profileQuestionnaireLines)

		if tranInfor.ReturnStatus == false && len(tranInfor.ValidationErrors) > 0 {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, ValidationErrors: tranInfor.ValidationErrors, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		} else if tranInfor.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}

		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": profileQuestionnaireLines}}, http.StatusOK)
	}
}
