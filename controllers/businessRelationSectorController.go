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

func API_BusinessRelationSectors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := context.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSectors": []models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}

		infiniteScrollingInformation := models.InfiniteScrollingInformation{
			After:          r.URL.Query().Get("After"),
			FetchSize:      r.URL.Query().Get("FetchSize"),
			SortDirection:  r.URL.Query().Get("SortDirection"),
			SortExpression: r.URL.Query().Get("SortExpression")}

		businessRelationSectors, tranInfor := models.GetBusinessRelationSectors(user.OrganizationID, r.URL.Query().Get("Search"), infiniteScrollingInformation)
		if tranInfor.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSectors": []models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: true, TotalRows: len(businessRelationSectors), Data: map[string]interface{}{"BusinessRelationSectors": businessRelationSectors}, IsAuthenticated: true}, http.StatusOK)
	case r.Method == "POST":
		businessRelationSector := models.BusinessRelationSector{}
		err := json.NewDecoder(r.Body).Decode(&businessRelationSector)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": []models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}
		if businessRelationSector.ID == "" {
			businessRelationSector.RecCreatedByID = user.ID
			businessRelationSector.RecModifiedByID = user.ID
			businessRelationSector.RecCreated = &models.Timestamp{time.Now()}
			businessRelationSector.RecModified = &models.Timestamp{time.Now()}
			businessRelationSector.ClientID = user.ClientID
			businessRelationSector.OrganizationID = user.OrganizationID
		} else {
			businessRelationSector.RecModifiedByID = user.ID
			businessRelationSector.RecModified = &models.Timestamp{time.Now()}
		}

		businessRelationSector, tranInfor := models.PostBusinessRelationSector(businessRelationSector)
		if tranInfor.ReturnStatus == false && len(tranInfor.ValidationErrors) > 0 {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, ValidationErrors: tranInfor.ValidationErrors, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": businessRelationSector}}, http.StatusBadRequest)
			return
		} else if tranInfor.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfor.ReturnStatus, ReturnMessage: tranInfor.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": businessRelationSector}}, http.StatusBadRequest)
			return
		}

		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": businessRelationSector}}, http.StatusOK)

	case r.Method == "DELETE":
		user, err := models.GetUser(requestUser.ID)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSectors": []models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}
		ids := strings.Split(r.URL.Query().Get("ID"), ",")
		tranInfo := models.DeleteBusinessRelationSectorById(user.OrganizationID, ids)
		if tranInfo.ReturnStatus == false {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true}, http.StatusOK)
	}
}

func API_BusinessRelationSector_Id(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "GET":
		ID := r.URL.Query().Get("ID")
		if ID == "" {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}
		BusinessRelationSector, tranInfo := models.GetBusinessRelationSectorByID(ID)
		if !tranInfo.ReturnStatus {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationSector": models.BusinessRelationSector{}}}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, Data: map[string]interface{}{"BusinessRelationSector": BusinessRelationSector}, IsAuthenticated: true}, http.StatusOK)
	}
}
