package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hatuan/go_crm/log"
	"github.com/hatuan/go_crm/settings"

	"github.com/jmoiron/sqlx"
)

type BusinessRelationSector struct {
	ID                *int64     `db:"id" json:",string"`
	Code              string     `db:"code"`
	Description       string     `db:"description"`
	RecCreatedByID    int64      `db:"rec_created_by" json:",string"`
	RecCreatedByUser  string     `db:"rec_created_by_user"`
	RecCreated        *Timestamp `db:"rec_created_at"`
	RecModifiedByID   int64      `db:"rec_modified_by" json:",string"`
	RecModifiedByUser string     `db:"rec_modified_by_user"`
	RecModified       *Timestamp `db:"rec_modified_at"`
	Status            int8       `db:"status"`
	Version           int16      `db:"version"`
	ClientID          int64      `db:"client_id" json:",string"`
	OrganizationID    int64      `db:"organization_id" json:",string"`
	Organization      string     `db:"organization"`
}

// ErrBusinessRelationSectorNotFound indicates there was no BusinessRelationSector
var ErrBusinessRelationSectorNotFound = errors.New("BusinessRelationSector not found")

// ErrBusinessRelationSectorDescriptionNotSpecified indicates there was no name given by the user
var ErrBusinessRelationSectorDescriptionNotSpecified = errors.New("BusinessRelationSector's description not specified")

// ErrBusinessRelationSectorCodeNotSpecified indicates there was no code given by the user
var ErrBusinessRelationSectorCodeNotSpecified = errors.New("BusinessRelationSector's code not specified")

// ErrBusinessRelationSectorCodeDuplicate indicates there was duplicate of code given by the user
var ErrBusinessRelationSectorCodeDuplicate = errors.New("BusinessRelationSector's code is duplicate")

// ErrBusinessRelationSectorFatal indicates there was fatal error
var ErrBusinessRelationSectorFatal = errors.New("BusinessRelationSector has fatal error")

// ErrBusinessRelationSectorValidate indicates there was validate error
var ErrBusinessRelationSectorValidate = errors.New("BusinessRelationSector has validate error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *BusinessRelationSector) Validate() map[string]InterfaceArray {
	validationErrors := make(map[string]InterfaceArray)

	if c.Code == "" {
		validationErrors["Code"] = append(validationErrors["Code"], ErrBusinessRelationSectorCodeNotSpecified.Error())
	}
	if c.Description == "" {
		validationErrors["Description"] = append(validationErrors["Description"], ErrBusinessRelationSectorDescriptionNotSpecified.Error())
	}
	if c.Code != "" {
		db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
		if err != nil {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrBusinessRelationSectorFatal.Error())
		}
		defer db.Close()
		var otherID string
		ID := int64(0)
		if c.ID != nil {
			ID = *c.ID
		}
		err = db.Get(&otherID, "SELECT id FROM business_relation_sector WHERE code = $1 AND id != $2 AND client_id = $3", c.Code, ID, c.ClientID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrBusinessRelationSectorFatal.Error())
		}
		if otherID != "" && err != sql.ErrNoRows {
			validationErrors["Code"] = append(validationErrors["Code"], ErrBusinessRelationSectorCodeDuplicate.Error())
		}
	}

	return validationErrors
}

func GetBusinessRelationSectors(orgID int64, searchCondition string, infiniteScrollingInformation InfiniteScrollingInformation) ([]BusinessRelationSector, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT business_relation_sector.*, user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, organization.name as organization" +
		" FROM business_relation_sector " +
		" INNER JOIN user_profile as user_created ON business_relation_sector.rec_created_by = user_created.id " +
		" INNER JOIN user_profile as user_modified ON business_relation_sector.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON business_relation_sector.organization_id = organization.id "

	sqlWhere := " WHERE business_relation_sector.organization_id = $1"
	if len(searchCondition) > 0 {
		sqlWhere += fmt.Sprintf(" AND %s", searchCondition)
	}

	var sqlOrder string
	if len(infiniteScrollingInformation.SortDirection) == 0 || infiniteScrollingInformation.SortDirection == "ASC" {
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s ASC", "business_relation_sector."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	} else { //sort DESC
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s DESC", "business_relation_sector."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	}

	sqlLimit := ""
	if len(infiniteScrollingInformation.FetchSize) > 0 {
		sqlLimit += fmt.Sprintf(" LIMIT %s ", infiniteScrollingInformation.FetchSize)
	}
	sqlString += sqlWhere + sqlOrder + sqlLimit
	log.Debug(sqlString)

	businessRelationSectors := []BusinessRelationSector{}
	err = db.Select(&businessRelationSectors, sqlString, orgID)

	if err != nil {
		log.Error(err)
		return businessRelationSectors, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return businessRelationSectors, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(businessRelationSectors)) + " records found"}}
}

func PostBusinessRelationSector(businessRelationSector BusinessRelationSector) (BusinessRelationSector, TransactionalInformation) {
	if validateErrs := businessRelationSector.Validate(); len(validateErrs) != 0 {
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationSectorValidate.Error()}, ValidationErrors: validateErrs}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	if businessRelationSector.ID == nil {
		businessRelationSector.Version = 1
		stmt, _ := db.PrepareNamed("INSERT INTO business_relation_sector(code, description, rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:code, :description, :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id) RETURNING id")
		var id int64
		err := stmt.Get(&id, businessRelationSector)
		if err != nil {
			log.Error(err)
			return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		} else {
			businessRelationSector.ID = &id
		}
	} else {
		stmt, _ := db.PrepareNamed("UPDATE business_relation_sector SET " +
			"code = :code," +
			"description = :description," +
			"status = :status," +
			"version = :version + 1," +
			"rec_modified_by = :rec_modified_by, rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

		result, err := stmt.Exec(businessRelationSector)
		if err != nil {
			log.Error(err)
			return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		changes, err := result.RowsAffected()
		if err != nil {
			log.Error(err)
			return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		if changes == 0 {
			return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationSectorNotFound.Error()}}
		}
	}

	businessRelationSector, _ = GetBusinessRelationSectorByID(*businessRelationSector.ID)
	return businessRelationSector, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

// GetBusinessRelationSectorByID returns the BusinessRelationSector that the given id corresponds to. If no BusinessRelationSector is found, an
// error is thrown.
func GetBusinessRelationSectorByID(id int64) (BusinessRelationSector, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	businessRelationSector := BusinessRelationSector{}
	err = db.Get(&businessRelationSector, "SELECT business_relation_sector.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM business_relation_sector "+
		"		INNER JOIN user_profile as user_created ON business_relation_sector.rec_created_by = user_created.id "+
		"		INNER JOIN user_profile as user_modified ON business_relation_sector.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON business_relation_sector.organization_id = organization.id "+
		"	WHERE business_relation_sector.id=$1", id)

	if err != nil && err == sql.ErrNoRows {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationSectorNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return businessRelationSector, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

// GetBusinessRelationSectorByCode returns the BusinessRelationSector that the given id corresponds to.
// If no BusinessRelationSector is found, an error is thrown.
func GetBusinessRelationSectorByCode(code string, orgID int64) (BusinessRelationSector, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	org, _ := GetOrganizationByID(orgID)

	businessRelationSector := BusinessRelationSector{}
	err = db.Get(&businessRelationSector, "SELECT business_relation_sector.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM business_relation_sector "+
		"		INNER JOIN user_profile as user_created ON business_relation_sector.rec_created_by = user_created.id "+
		"		INNER JOIN user_profile as user_modified ON business_relation_sector.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON business_relation_sector.organization_id = organization.id "+
		"	WHERE business_relation_sector.code=$1 and business_relation_sector.client_id=$2", code, org.ClientID)

	if err != nil && err == sql.ErrNoRows {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationSectorNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return BusinessRelationSector{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return businessRelationSector, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

func DeleteBusinessRelationSectorById(orgID int64, ids []string) TransactionalInformation {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	query, args, err := sqlx.In("DELETE FROM business_relation_sector "+
		" WHERE business_relation_sector.id IN (?) and business_relation_sector.organization_id=?", ids, orgID)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err = db.Exec(query, args...)
	if err != nil && err == sql.ErrNoRows {
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationSectorNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}
