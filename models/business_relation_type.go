package models

import (
	"database/sql"
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"errors"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
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

// ErrBusinessRelationTypeNotFound indicates there was no BusinessRelationType
var ErrBusinessRelationTypeNotFound = errors.New("BusinessRelationType not found")

// ErrBusinessRelationTypeNameNotSpecified indicates there was no name given by the user
var ErrBusinessRelationTypeNameNotSpecified = errors.New("BusinessRelationType's name not specified")

// ErrBusinessRelationTypeCodeNotSpecified indicates there was no code given by the user
var ErrBusinessRelationTypeCodeNotSpecified = errors.New("BusinessRelationType's code not specified")

// ErrBusinessRelationTypeCodeDuplicate indicates there was duplicate of code given by the user
var ErrBusinessRelationTypeCodeDuplicate = errors.New("BusinessRelationType's code is duplicate")

// ErrBusinessRelationTypeFatal indicates there was fatal error
var ErrBusinessRelationTypeFatal = errors.New("BusinessRelationType has fatal error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *BusinessRelationType) Validate() error {
	switch {
	case c.Code == "":
		return ErrBusinessRelationTypeCodeNotSpecified
	case c.Name == "":
		return ErrBusinessRelationTypeNameNotSpecified
	case c.Code != "":
		db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
		if err != nil {
			log.Fatal(err)
			return ErrBusinessRelationTypeFatal
		}
		defer db.Close()
		var id string
		err = db.Get(&id, "SELECT ID FROM business_relation_type WHERE code = $1 AND organization_id = $2", c.Code, c.OrganizationID)
		if err != nil && err != sql.ErrNoRows {
			log.Fatal(err)
			return ErrBusinessRelationTypeFatal
		}
		if id != c.ID && err != sql.ErrNoRows {
			return ErrBusinessRelationTypeCodeDuplicate
		}
	}
	return nil
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
		" WHERE business_relation_type.organization_id = $1 ORDER BY business_relation_type.code", orgID)

	if err != nil {
		log.Error(err)
		return businessRelationTypes, err
	}

	return businessRelationTypes, nil
}

func PostBusinessRelationType(businessRelationType BusinessRelationType) (BusinessRelationType, error) {
	if err := businessRelationType.Validate(); err != nil {
		return BusinessRelationType{}, err
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
		return BusinessRelationType{}, err
	}
	defer db.Close()

	if businessRelationType.ID == "" {
		businessRelationType.ID = uuid.NewV1().String()
		businessRelationType.Version = 1
		stmt, _ := db.PrepareNamed("INSERT INTO business_relation_type(id, code, name, rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:id, :code, :name, :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id)")
		_, err := stmt.Exec(businessRelationType)
		if err != nil {
			log.Error(err)
			return BusinessRelationType{}, err
		}

	} else {
		stmt, _ := db.PrepareNamed("UPDATE business_relation_type SET " +
			"code = :code," +
			"name = :name," +
			"status = :status," +
			"version = :version + 1," +
			"rec_modified_by = :rec_modified_by, rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

		_, err := stmt.Exec(businessRelationType)
		if err != nil {
			log.Error(err)
			return BusinessRelationType{}, err
		}
	}
	businessRelationType, _ = GetBusinessRelationTypeByID(businessRelationType.ID)
	return businessRelationType, nil
}

// GetBusinessRelationTypeByID returns the BusinessRelationType that the given id corresponds to. If no BusinessRelationType is found, an
// error is thrown.
func GetBusinessRelationTypeByID(id string) (BusinessRelationType, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	businessRelationType := BusinessRelationType{}
	err = db.Get(&businessRelationType, "SELECT business_relation_type.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM business_relation_type "+
		"		INNER JOIN \"user\" as user_created ON business_relation_type.rec_created_by = user_created.id "+
		"		INNER JOIN \"user\" as user_modified ON business_relation_type.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON business_relation_type.organization_id = organization.id "+
		"	WHERE business_relation_type.id=$1", id)
	if err != nil && err == ErrBusinessRelationTypeNotFound {
		return BusinessRelationType{}, ErrBusinessRelationTypeNotFound
	} else if err != nil {
		return BusinessRelationType{}, err
	}
	return businessRelationType, nil
}

// GetBusinessRelationTypeByCode returns the BusinessRelationType that the given id corresponds to. If no BusinessRelationType is found, an
// error is thrown.
func GetBusinessRelationTypeByCode(code string, orgID string) (BusinessRelationType, error) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	businessRelationType := BusinessRelationType{}
	err = db.Get(&businessRelationType, "SELECT business_relation_type.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM business_relation_type "+
		"		INNER JOIN \"user\" as user_created ON business_relation_type.rec_created_by = user_created.id "+
		"		INNER JOIN \"user\" as user_modified ON business_relation_type.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON business_relation_type.organization_id = organization.id "+
		"	WHERE business_relation_type.code=$1 and business_relation_type.organization_id=$2", code, orgID)

	if err != nil && err == sql.ErrNoRows {
		return BusinessRelationType{}, ErrBusinessRelationTypeNotFound
	} else if err != nil {
		return BusinessRelationType{}, err
	}
	return businessRelationType, nil
}
