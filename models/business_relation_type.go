package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"

	"github.com/jmoiron/sqlx"
)

type BusinessRelationType struct {
	ID                string     `db:"id"`
	Code              string     `db:"code"`
	Name              string     `db:"name"`
	RecCreatedByID    string     `db:"rec_created_by"`
	RecCreatedByUser  string     `db:"rec_created_by_user"`
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

func GetBusinessRelationTypes(orgID string) ([]BusinessRelationType, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return []BusinessRelationType{}, err
	}
	defer db.Close()

	businessRelationTypes := []BusinessRelationType{}
	err = db.Select(&businessRelationTypes, "SELECT business_relation_type.*, user_created.name as rec_created_by_user, "+
		" user_modified.name as rec_modified_by_user, organization.name as organization"+
		" FROM business_relation_type "+
		" INNER JOIN \"user\" as user_created ON business_relation_type.rec_created_by = user_created.id "+
		" INNER JOIN \"user\" as user_modified ON business_relation_type.rec_modified_by = user_modified.id "+
		" INNER JOIN organization as organization ON business_relation_type.organization_id = organization.id "+
		" WHERE business_relation_type.organization_id = $1", orgID)

	if err != nil {
		log.Error(err)
		return businessRelationTypes, err
	}

	return businessRelationTypes, nil
}
