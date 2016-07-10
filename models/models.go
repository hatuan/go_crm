package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"fmt"

	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Flash struct {
	Type    string
	Message string
}

type Response struct {
	ReturnStatus     bool
	ReturnMessage    []string
	ValidationErrors map[string]interface{}
	TotalPages       int
	TotalRows        int
	PageSize         int
	IsAuthenticated  bool
	Data             map[string]interface{}
}

type Token struct {
	TransactionalInformation
	Token string `json:"token"`
}

type LoginDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type TransactionalInformation struct {
	ReturnStatus     bool
	ReturnMessage    []string
	ValidationErrors map[string]interface{}
	TotalPages       int
	TotalRows        int
	PageSize         int
	IsAuthenticated  bool
}

type ApplicationMenuDTO struct {
	MenuID                 string `json:"menu_id"`
	Description            string `json:"description"`
	Route                  string `json:"route"`
	Module                 string `json:"module"`
	MenuOrder              int    `json:"menu_order"`
	RequiresAuthentication bool   `json:"requires_authentication"`
}

//ApplicationModelDTO user for return from  controllers.InitializeApplication
type ApplicationModelDTO struct {
	TransactionalInformation
	MenuItems []ApplicationMenuDTO `json:"menu_items"`
}

func CheckUnique(table, ID, code, orgID string) (bool, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	strSQL := fmt.Sprintf("SELECT id FROM %s WHERE code = $1 AND id <> $2 AND  organization_id = $3", table)
	log.Info(strSQL)

	var otherID string
	err = db.Get(&otherID, strSQL, code, ID, orgID)

	if err != nil && err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		log.Fatal(err)
		return false, err
	}
	return false, nil
}
