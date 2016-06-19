package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"

	"github.com/jmoiron/sqlx"
)

type BusinessRelationSector struct {
	ID                string     `db:"id"`
	Code              string     `db:"code"`
	Name              string     `db:"name"`
	RecCreatedByID    string     `db:"rec_created_by"`
	RecCreatedByName  string     `db:"rec_created_by_user"`
	RecCreated        *Timestamp `db:"rec_created_at"`
	RecModifiedByID   string     `db:"rec_modified_by"`
	RecModifiedByUser string     `db:"rec_modified_by_user"`
	RecModified       *Timestamp `db:"rec_modified_at"`
	Status            int8       `db:"status"`
	Version           int16      `db:"version"`
	ClientID          string     `db:"client_id"`
	OrganizationID    string     `db:"organization_id"`
	Organization      string     `db:"organization"`
}

func GetBusinessRelationSectors(orgID string) ([]BusinessRelationSector, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return []BusinessRelationSector{}, err
	}
	defer db.Close()

	businessRelationSectors := []BusinessRelationSector{}
	err = db.Select(&businessRelationSectors, "SELECT business_relation_sector.*, user_created.name as rec_created_by_user, "+
		" user_modified.name as rec_modified_by_user, organization.name as organization"+
		" FROM business_relation_sector "+
		" INNER JOIN \"user\" as user_created ON business_relation_sector.rec_created_by = user_created.id "+
		" INNER JOIN \"user\" as user_modified ON business_relation_sector.rec_modified_by = user_modified.id "+
		" INNER JOIN organization as organization ON business_relation_sector.organization_id = organization.id "+
		" WHERE business_relation_sector.organization_id = $1", orgID)

	if err != nil {
		log.Error(err)
		return businessRelationSectors, err
	}

	return businessRelationSectors, nil
}
