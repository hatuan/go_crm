package models

import (
	"database/sql"
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type BusinessRelationType struct {
	ID                string     `db:"id"`
	Code              string     `db:"code"`
	Description       string     `db:"description"`
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

// ErrBusinessRelationTypeDescriptionNotSpecified indicates there was no name given by the user
var ErrBusinessRelationTypeDescriptionNotSpecified = errors.New("BusinessRelationType's description not specified")

// ErrBusinessRelationTypeCodeNotSpecified indicates there was no code given by the user
var ErrBusinessRelationTypeCodeNotSpecified = errors.New("BusinessRelationType's code not specified")

// ErrBusinessRelationTypeCodeDuplicate indicates there was duplicate of code given by the user
var ErrBusinessRelationTypeCodeDuplicate = errors.New("BusinessRelationType's code is duplicate")

// ErrBusinessRelationTypeFatal indicates there was fatal error
var ErrBusinessRelationTypeFatal = errors.New("BusinessRelationType has fatal error")

// ErrBusinessRelationTypeValidate indicates there was validate error
var ErrBusinessRelationTypeValidate = errors.New("BusinessRelationType has validate error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *BusinessRelationType) Validate() map[string]InterfaceArray {
	validationErrors := make(map[string]InterfaceArray)

	if c.Code == "" {
		validationErrors["Code"] = append(validationErrors["Code"], ErrBusinessRelationTypeCodeNotSpecified.Error())
	}
	if c.Description == "" {
		validationErrors["Description"] = append(validationErrors["Description"], ErrBusinessRelationTypeDescriptionNotSpecified.Error())
	}
	if c.Code != "" {
		db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
		if err != nil {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrBusinessRelationTypeFatal.Error())
		}
		defer db.Close()
		var otherID string
		ID := EmptyUUID
		if c.ID != "" {
			ID = c.ID
		}
		err = db.Get(&otherID, "SELECT id FROM business_relation_type WHERE code = $1 AND id != $2 AND client_id = $3", c.Code, ID, c.ClientID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrBusinessRelationTypeFatal.Error())
		}
		if otherID != "" && err != sql.ErrNoRows {
			validationErrors["Code"] = append(validationErrors["Code"], ErrBusinessRelationTypeCodeDuplicate.Error())
		}
	}
	return validationErrors
}

func GetBusinessRelationTypes(orgID string, searchCondition string, infiniteScrollingInformation InfiniteScrollingInformation) ([]BusinessRelationType, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT business_relation_type.*, user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, organization.name as organization" +
		" FROM business_relation_type " +
		" INNER JOIN \"user\" as user_created ON business_relation_type.rec_created_by = user_created.id " +
		" INNER JOIN \"user\" as user_modified ON business_relation_type.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON business_relation_type.organization_id = organization.id "

	sqlWhere := " WHERE business_relation_type.organization_id = $1"
	if len(searchCondition) > 0 {
		sqlWhere += fmt.Sprintf(" AND %s", searchCondition)
	}

	var sqlOrder string
	if len(infiniteScrollingInformation.SortDirection) == 0 || infiniteScrollingInformation.SortDirection == "ASC" {
		//if len(infiniteScrollingInformation.After) >= 0 && len(infiniteScrollingInformation.SortExpression) > 0 {
		///	sqlWhere += fmt.Sprintf(" AND %s > $2", "business_relation_type."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		//}
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s ASC", "business_relation_type."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	} else { //sort DESC
		//if len(infiniteScrollingInformation.After) >= 0 && len(infiniteScrollingInformation.SortDirection) > 0 {
		//	sqlWhere += fmt.Sprintf(" AND %s < $2", "business_relation_type."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		//}
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s DESC", "business_relation_type."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	}
	sqlLimit := ""
	if len(infiniteScrollingInformation.FetchSize) > 0 {
		sqlLimit += fmt.Sprintf(" LIMIT %s ", infiniteScrollingInformation.FetchSize)
	}
	sqlString += sqlWhere + sqlOrder + sqlLimit
	log.Debug(sqlString)

	businessRelationTypes := []BusinessRelationType{}
	err = db.Select(&businessRelationTypes, sqlString, orgID)

	if err != nil {
		log.Error(err)
		return businessRelationTypes, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return businessRelationTypes, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(businessRelationTypes)) + " records found"}}
}

func PostBusinessRelationType(businessRelationType BusinessRelationType) (BusinessRelationType, TransactionalInformation) {
	if validateErrs := businessRelationType.Validate(); len(validateErrs) != 0 {
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationTypeValidate.Error()}, ValidationErrors: validateErrs}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	if businessRelationType.ID == "" {
		businessRelationType.ID = uuid.NewV4().String()
		businessRelationType.Version = 1
		stmt, _ := db.PrepareNamed("INSERT INTO business_relation_type(id, code, description, rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:id, :code, :description, :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id)")
		_, err := stmt.Exec(businessRelationType)
		if err != nil {
			log.Error(err)
			return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}

	} else {
		stmt, _ := db.PrepareNamed("UPDATE business_relation_type SET " +
			"code = :code," +
			"description = :description," +
			"status = :status," +
			"version = :version + 1," +
			"rec_modified_by = :rec_modified_by, rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

		result, err := stmt.Exec(businessRelationType)
		if err != nil {
			log.Error(err)
			return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		changes, err := result.RowsAffected()
		if err != nil {
			log.Error(err)
			return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		if changes == 0 {
			return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationTypeNotFound.Error()}}
		}
	}
	businessRelationType, _ = GetBusinessRelationTypeByID(businessRelationType.ID)
	return businessRelationType, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

// GetBusinessRelationTypeByID returns the BusinessRelationType that the given id corresponds to. If no BusinessRelationType is found, an
// error is thrown.
func GetBusinessRelationTypeByID(id string) (BusinessRelationType, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
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
	if err != nil && err == sql.ErrNoRows {
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationTypeNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return businessRelationType, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

// GetBusinessRelationTypeByCode returns the BusinessRelationType that the given id corresponds to.
// If no BusinessRelationType is found, an error is thrown.
func GetBusinessRelationTypeByCode(code string, orgID string) (BusinessRelationType, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	org, _ := GetOrganizationByID(orgID)

	businessRelationType := BusinessRelationType{}
	err = db.Get(&businessRelationType, "SELECT business_relation_type.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM business_relation_type "+
		"		INNER JOIN \"user\" as user_created ON business_relation_type.rec_created_by = user_created.id "+
		"		INNER JOIN \"user\" as user_modified ON business_relation_type.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON business_relation_type.organization_id = organization.id "+
		"	WHERE business_relation_type.code=$1 and business_relation_type.client_id=$2", code, org.ClientID)

	if err != nil && err == sql.ErrNoRows {
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationTypeNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return BusinessRelationType{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return businessRelationType, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

func DeleteBusinessRelationTypeById(orgID string, ids []string) TransactionalInformation {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	query, args, err := sqlx.In("DELETE FROM business_relation_type "+
		" WHERE business_relation_type.id IN (?) and business_relation_type.organization_id=?", ids, orgID)

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	_, err = db.Exec(query, args...)

	if err != nil && err == sql.ErrNoRows {
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrBusinessRelationTypeNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}
