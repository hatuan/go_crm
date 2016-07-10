package controllers

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/models"
	"net/http"
)

func API_Check_Unique(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch {
	case r.Method == "POST":
		type CheckUniqueDTO struct {
			UserID string
			Table  string
			ID     string
			Code   string
		}
		//Print r.Body
		//body, err := ioutil.ReadAll(r.Body)
		//if err != nil {
		//	panic(err)
		//}
		//log.Info(string(body))

		err := r.ParseForm()
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, err.Error(), http.StatusOK)
			return
		}
		userID := r.Form.Get("UserID")
		code := r.Form.Get("Code")
		table := r.Form.Get("Table")
		recID := r.Form.Get("RecID")

		user, err := models.GetUser(userID)

		valid, err := models.CheckUnique(table, recID, code, user.OrganizationID)
		if err != nil {
			log.Error(err.Error())
			JSONResponse(w, err.Error(), http.StatusOK)
			return
		}

		if valid {
			JSONResponse(w, "true", http.StatusOK)
		} else {
			JSONResponse(w, nil, http.StatusOK)
		}
	}
}
