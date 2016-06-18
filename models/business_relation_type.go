package models

import (
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"

	"github.com/jmoiron/sqlx"
)

type BusinessRelationType struct {
	ID                string       `db:"id"`
	Code              string       `db:"code"`
	Name              string       `db:"name"`
	RecCreatedByID    string       `db:"rec_created_by_id"`
	RecCreatedByUser  User         `db:"rec_created_by_user" db:"-"`
	RecCreated        *Timestamp   `db:"rec_created_at"`
	RecModifiedByID   string       `db:"rec_modified_by_id"`
	RecModifiedByUser User         `db:"rec_modified_by_user" db:"-"`
	RecModified       *Timestamp   `db:"rec_modified_at"`
	Status            int8         `db:"status"`
	Version           int16        `db:"version"`
	ClientID          string       `db:"client_id"`
	OrganizationID    string       `db:"organization_id"`
	Organization      Organization `db:"-"`
}

func GetBusinessRelationTypes(orgID string) ([]BusinessRelationType, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return []BusinessRelationType{}, err
	}
	defer db.Close()

	businessRelationTypes := []BusinessRelationType{}
	err = db.Select(&businessRelationTypes, "SELECT business_relation_type.*, rec_created_by_user.name, rec_modified_by_user.name, organization.name "+
		" FROM business_relation_type "+
		" INNER JOIN user as rec_created_by_user ON business_relation_type.rec_created_by_id = rec_created_by_user.id "+
		" INNER JOIN user as rec_modified_by_user ON business_relation_type.rec_modified_by_id = rec_modified_by_user.id "+
		" INNER JOIN organization as organization ON business_relation_type.organization_id = organization.id "+
		" WHERE organization_id = $1", orgID)
	if err != nil {
		log.Error(err)
		return businessRelationTypes, err
	}

	return businessRelationTypes, nil
}
