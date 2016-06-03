package controllers

import (
	"encoding/json"
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

//API_Users_Id Get & Update User
func API_Users_Id(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		log.Error(err.Error())
		JSONResponse(w, models.User{}, http.StatusInternalServerError)
		return
	}
	u, err := models.GetUser(id.String())
	if err != nil {
		log.Error(err.Error())
		JSONResponse(w, u, http.StatusInternalServerError)
		return
	}

	switch {
	case r.Method == "GET":
		JSONResponse(w, u, http.StatusOK)

	}
}

//API_User_Preference Get & Set User's Preference
func API_User_Preference(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := ctx.Get(r, "user").(models.User)

	switch {
	case r.Method == "GET":
		preference, err := requestUser.GetPreference()
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"Preference": models.PreferenceDTO{}}}, http.StatusInternalServerError)
		}
		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"Preference": preference}}, http.StatusOK)
	case r.Method == "POST": //update preference to user
		preference := new(models.PreferenceDTO)
		err := json.NewDecoder(r.Body).Decode(&preference)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"Preference": models.PreferenceDTO{}}}, http.StatusInternalServerError)
		}
		err = requestUser.SetPreference(*preference)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, models.Response{ReturnStatus: false, ReturnMessage: []string{err.Error()}, IsAuthenticated: true, Data: map[string]interface{}{"Preference": preference}}, http.StatusInternalServerError)
		}
		JSONResponse(w, models.Response{ReturnStatus: true, IsAuthenticated: true, Data: map[string]interface{}{"Preference": preference}}, http.StatusOK)
	}

}
