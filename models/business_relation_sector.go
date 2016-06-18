package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"

	"github.com/jmoiron/sqlx"
)

type BusinessRelationSector struct {
	ID                string       `json:"id"`
	Code              string       `json:"code"`
	Name              string       `json:"name"`
	RecCreatedByID    string       `json:"rec_created_by_id"`
	RecCreatedByUser  User         `json:"rec_created_by_user" db:"-"`
	RecCreated        *Timestamp   `json:"rec_created"`
	RecModifiedByID   string       `json:"rec_modified_by_id"`
	RecModifiedByUser User         `json:"rec_modified_by_user" db:"-"`
	RecModified       *Timestamp   `json:"rec_modified"`
	Status            int8         `json:"status"`
	Version           int16        `json:"version"`
	ClientID          string       `json:"client_id"`
	Client            Client       `json:"client" db:"-"`
	OrganizationID    string       `json:"organization_id"`
	Organization      Organization `json:"organization" db:"-"`
}

func GetBusinessRelationSectors(orgID string) ([]BusinessRelationSector, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return []BusinessRelationSector{}, err
	}
	defer db.Close()

	businessRelationSectors := []BusinessRelationSector{}
	err = db.Select(&businessRelationSectors, "SELECT * FROM business_relation_sector WHERE organization_id = $1", orgID)
	if err != nil {
		log.Error(err)
		return businessRelationSectors, err
	}

	return businessRelationSectors, nil
}
