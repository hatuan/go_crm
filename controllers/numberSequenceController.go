package controllers

/*
import (
	"encoding/json"
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"strconv"

	"net/http"
	"strings"
	"time"

	"github.com/gorilla/context"
)

func API_BusinessRelationType_Id(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "GET":
		ID, err := strconv.ParseInt(r.URL.Query().Get("ID"), 10, 64)
		if err != nil {
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{ErrIDParameterNotFound.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": models.BusinessRelationType{}}}, http.StatusBadRequest)
			return
		}
		businessRelationType, tranInfo := models.GetBusinessRelationTypeByID(ID)
		if !tranInfo.ReturnStatus {
			JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, IsAuthenticated: true, Data: map[string]interface{}{"BusinessRelationType": models.BusinessRelationType{}}}, http.StatusBadRequest)
			return
		}
		JSONResponse(w, models.Response{ReturnStatus: tranInfo.ReturnStatus, ReturnMessage: tranInfo.ReturnMessage, Data: map[string]interface{}{"BusinessRelationType": businessRelationType}, IsAuthenticated: true}, http.StatusOK)
	}
}
*/
