package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/hatuan/go_crm/log"
	"github.com/hatuan/go_crm/models"

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

func API_ProfileQuestionnaire_By_HeaderId(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		params := mux.Vars(r)
		headerID, err := strconv.ParseInt(params["headerid"], 10, 64)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true}, http.StatusBadRequest)
			return
		}

		profileQuestionnaire, tranInfo := models.GetProfileQuestionnaireHeaderByID(headerID)
		if !tranInfo.ReturnStatus {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true}, http.StatusBadRequest)
			return
		}

		profileQuestionnaireLines, tranInfo := models.GetProfileQuestionnaireLinesByHeaderID(headerID)

		if !tranInfo.ReturnStatus {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true}, http.StatusBadRequest)
			return
		}

		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, Data: map[string]interface{}{"ProfileQuestionnaire": profileQuestionnaire, "ProfileQuestionnaireLines": profileQuestionnaireLines}, IsAuthenticated: true}, http.StatusOK)
	case r.Method == "POST":
		params := mux.Vars(r)
		headerID, err := strconv.ParseInt(params["headerid"], 10, 64)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}

		var postData struct {
			ProfileQuestionnaireLines []models.ProfileQuestionnaireLine
		}

		err = json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}

		profileQuestionnaireLines := postData.ProfileQuestionnaireLines

		user, err := models.GetUser(*requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"ProfileQuestionnaireLines": []models.ProfileQuestionnaireLine{}}}, http.StatusBadRequest)
			return
		}

		//range uses a[i] as its second value for arrays/slices, which effectively means that the value is copied, making the original value untouchable.
		for index := range profileQuestionnaireLines {
			if profileQuestionnaireLines[index].ID == nil {
				id, err := models.IDGenerator()
				if err != nil {
					log.Error(err.Error())
					JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true}, http.StatusBadRequest)
					return
				}
				profileQuestionnaireLines[index].ID = &id
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

			for ratingIndex := range profileQuestionnaireLines[index].Ratings {
				if profileQuestionnaireLines[index].Ratings[ratingIndex].ID == nil {
					id, err := models.IDGenerator()
					if err != nil {
						log.Error(err.Error())
						JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true}, http.StatusBadRequest)
						return
					}
					profileQuestionnaireLines[index].Ratings[ratingIndex].ID = &id
					profileQuestionnaireLines[index].Ratings[ratingIndex].Version = 1
					profileQuestionnaireLines[index].Ratings[ratingIndex].RecCreatedByID = *user.ID
					profileQuestionnaireLines[index].Ratings[ratingIndex].RecModifiedByID = *user.ID
					profileQuestionnaireLines[index].Ratings[ratingIndex].RecCreated = &models.Timestamp{time.Now()}
					profileQuestionnaireLines[index].Ratings[ratingIndex].RecModified = &models.Timestamp{time.Now()}
					profileQuestionnaireLines[index].Ratings[ratingIndex].ClientID = user.ClientID
					profileQuestionnaireLines[index].Ratings[ratingIndex].OrganizationID = user.OrganizationID
				} else {
					profileQuestionnaireLines[index].Ratings[ratingIndex].RecModifiedByID = *user.ID
					profileQuestionnaireLines[index].Ratings[ratingIndex].RecModified = &models.Timestamp{time.Now()}
				}
			}
		}

		profileQuestionnaireLines, tranInfor := models.PostProfileQuestionnaireLines(headerID, profileQuestionnaireLines)

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

func API_ProfileQuestionnaireLine_Id_Ratings(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "POST":
		var postData struct {
			ProfileQuestionnaireLine models.ProfileQuestionnaireLine
		}

		err := json.NewDecoder(r.Body).Decode(&postData)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"Ratings": []models.Rating{}}}, http.StatusBadRequest)
			return
		}

		profileQuestionnaireLine := postData.ProfileQuestionnaireLine

		user, err := models.GetUser(*requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"Ratings": []models.Rating{}}}, http.StatusBadRequest)
			return
		}

		ratings := profileQuestionnaireLine.Ratings
		//range uses a[i] as its second value for arrays/slices, which effectively means that the value is copied, making the original value untouchable.
		for index := range ratings {
			if ratings[index].ID == nil {
				id, err := models.IDGenerator()
				if err != nil {
					log.Error(err.Error())
					JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true}, http.StatusBadRequest)
					return
				}
				ratings[index].ID = &id
				ratings[index].Version = 1
				ratings[index].RecCreatedByID = *user.ID
				ratings[index].RecModifiedByID = *user.ID
				ratings[index].RecCreated = &models.Timestamp{time.Now()}
				ratings[index].RecModified = &models.Timestamp{time.Now()}
				ratings[index].ClientID = user.ClientID
				ratings[index].OrganizationID = user.OrganizationID
			} else {
				ratings[index].RecModifiedByID = *user.ID
				ratings[index].RecModified = &models.Timestamp{time.Now()}
			}
		}

		ratings, tranInfor := models.PostRatingsWithLineID(profileQuestionnaireLine.ProfileQuestionnaireHeaderID, *profileQuestionnaireLine.ID, ratings)

		if tranInfor.ReturnStatus == false && len(tranInfor.ValidationErrors) > 0 {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, ValidationErrors: tranInfor.ValidationErrors, IsAuthenticated: true, Data: map[string]interface{}{"Ratings": []models.Rating{}}}, http.StatusBadRequest)
			return
		} else if tranInfor.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"Ratings": []models.Rating{}}}, http.StatusBadRequest)
			return
		}

		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"Ratings": ratings}}, http.StatusOK)
	}
}
