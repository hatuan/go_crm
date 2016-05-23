package models

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"errors"
	"database/sql"
)

type Organization struct {
	ID                string     `json:"id"`
	Code              string     `json:"code"`
	Name              string     `json:"name"`
	RecCreatedByID    string     `json:"rec_created_by_id"`
	RecCreatedByUser  User       `json:"rec_created_by_user"`
	RecCreated        *Timestamp `json:"rec_created"`
	RecModifiedByID   string     `json:"rec_modified_by_id"`
	RecModifiedByUser User       `json:"rec_modified_by_user"`
	RecModified       *Timestamp `json:"rec_modified"`
	Status            int8       `json:"status"`
	Version           int16      `json:"version"`
	ClientID          string     `json:"client_id"`
	Client            Client     `json:"client"`
}

var ErrRootOrganizationNotFound = errors.New("Error RootOrganization Not Found")
var ErrOrganizationNotFound = errors.New("Error Organization Not Found")

func (o Organization) GetRootOrganization() (Organization, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)

		return Organization{}, err
	}
	defer db.Close()

	rootOrganization := Organization{}

	err = db.Get(&rootOrganization, "SELECT * FROM organization WHERE client_id=$1 AND code='*'", o.ClientID)
	if err == sql.ErrNoRows {
		return rootOrganization, ErrRootOrganizationNotFound
	} else if err != nil {
		log.Error(err)
		return rootOrganization, err
	}
	return rootOrganization, nil
}

// Get returns the user that the given id corresponds to. If no user is found, an
// error is thrown.
func (o *Organization) Get(id string) error {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Get(&o, "SELECT * FROM organization WHERE id=$1::uuid", id)
	if err == sql.ErrNoRows {
		return ErrOrganizationNotFound
	} else if err != nil {
		return err
	}
	return nil
}