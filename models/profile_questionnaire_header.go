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

type ProfileQuestionnaireHeader struct {
	ID                        *int64                     `db:"id" json:",string"`
	Code                      string                     `db:"code"`
	Description               string                     `db:"description"`
	Priority                  int8                       `db:"priority"`
	ContactType               int8                       `db:"contact_type"`
	BusinessRelationTypeID    *int64                     `db:"business_relation_type_id"  json:",string"`
	BusinessRelationTypeCode  string                     `db:"business_relation_type_code"`
	ProfileQuestionnaireLines []ProfileQuestionnaireLine `db:"-"`
	RecCreatedByID            int64                      `db:"rec_created_by" json:",string"`
	RecCreatedByUser          string                     `db:"rec_created_by_user"`
	RecCreated                *Timestamp                 `db:"rec_created_at"`
	RecModifiedByID           int64                      `db:"rec_modified_by" json:",string"`
	RecModifiedByUser         string                     `db:"rec_modified_by_user"`
	RecModified               *Timestamp                 `db:"rec_modified_at"`
	Status                    int8                       `db:"status"`
	Version                   int16                      `db:"version"`
	ClientID                  int64                      `db:"client_id" json:",string"`
	OrganizationID            int64                      `db:"organization_id" json:",string"`
	Organization              string                     `db:"organization"`
}

// ErrProfileQuestionnaireHeaderNotFound indicates there was no ProfileQuestionnaireHeader
var ErrProfileQuestionnaireHeaderNotFound = errors.New("ProfileQuestionnaireHeader not found")

// ErrProfileQuestionnaireHeaderDescriptionNotSpecified indicates there was no description given by the user
var ErrProfileQuestionnaireHeaderDescriptionNotSpecified = errors.New("ProfileQuestionnaireHeader's description not specified")

// ErrProfileQuestionnaireHeaderCodeNotSpecified indicates there was no code given by the user
var ErrProfileQuestionnaireHeaderCodeNotSpecified = errors.New("ProfileQuestionnaireHeader's code not specified")

// ErrProfileQuestionnaireHeaderCodeDuplicate indicates there was duplicate of code given by the user
var ErrProfileQuestionnaireHeaderCodeDuplicate = errors.New("ProfileQuestionnaireHeader's code is duplicate")

// ErrProfileQuestionnaireHeaderFatal indicates there was fatal error
var ErrProfileQuestionnaireHeaderFatal = errors.New("ProfileQuestionnaireHeader has fatal error")

// ErrProfileQuestionnaireHeaderValidate indicates there was validate error
var ErrProfileQuestionnaireHeaderValidate = errors.New("ProfileQuestionnaireHeader has validate error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *ProfileQuestionnaireHeader) Validate() map[string]InterfaceArray {
	validationErrors := make(map[string]InterfaceArray)

	if c.Code == "" {
		validationErrors["Code"] = append(validationErrors["Code"], ErrProfileQuestionnaireHeaderCodeNotSpecified.Error())
	}
	if c.Description == "" {
		validationErrors["Description"] = append(validationErrors["Description"], ErrProfileQuestionnaireHeaderDescriptionNotSpecified.Error())
	}
	if c.Code != "" {
		db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
		if err != nil {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrProfileQuestionnaireHeaderFatal.Error())
		}
		defer db.Close()
		var otherID string
		ID := int64(0)
		if c.ID != nil {
			ID = *c.ID
		}
		err = db.Get(&otherID, "SELECT id FROM profile_questionnaire_header WHERE code = $1 AND id != $2 AND client_id = $3", c.Code, ID, c.ClientID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrProfileQuestionnaireHeaderFatal.Error())
		}
		if otherID != "" && err != sql.ErrNoRows {
			validationErrors["Code"] = append(validationErrors["Code"], ErrProfileQuestionnaireHeaderCodeDuplicate.Error())
		}
	}

	return validationErrors
}

func (c *ProfileQuestionnaireHeader) getDetails() []string {

	profileQuestionnaireLines, transactionInformation := GetProfileQuestionnaireLinesByHeaderID(*c.ID)

	if !transactionInformation.ReturnStatus {
		log.Error(transactionInformation.ReturnMessage)
		return transactionInformation.ReturnMessage
	}

	c.ProfileQuestionnaireLines = profileQuestionnaireLines
	return nil
}

func GetProfileQuestionnaireHeaders(orgID int64, searchCondition string, infiniteScrollingInformation InfiniteScrollingInformation) ([]ProfileQuestionnaireHeader, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT profile_questionnaire_header.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization, " +
		" COALESCE(business_relation_type.code, '') as business_relation_type_code" +
		" FROM profile_questionnaire_header " +
		" INNER JOIN user_profile as user_created ON profile_questionnaire_header.rec_created_by = user_created.id " +
		" INNER JOIN user_profile as user_modified ON profile_questionnaire_header.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON profile_questionnaire_header.organization_id = organization.id " +
		" LEFT JOIN business_relation_type as business_relation_type ON profile_questionnaire_header.business_relation_type_id = business_relation_type.id "

	sqlWhere := " WHERE profile_questionnaire_header.organization_id = $1"
	if len(searchCondition) > 0 {
		sqlWhere += fmt.Sprintf(" AND %s", searchCondition)
	}

	var sqlOrder string
	if len(infiniteScrollingInformation.SortDirection) == 0 || infiniteScrollingInformation.SortDirection == "ASC" {
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s ASC", "profile_questionnaire_header."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	} else { //sort DESC
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s DESC", "profile_questionnaire_header."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	}

	sqlLimit := ""
	if len(infiniteScrollingInformation.FetchSize) > 0 {
		sqlLimit += fmt.Sprintf(" LIMIT %s ", infiniteScrollingInformation.FetchSize)
	}
	sqlString += sqlWhere + sqlOrder + sqlLimit
	log.Debug(sqlString)

	ProfileQuestionnaireHeaders := []ProfileQuestionnaireHeader{}
	err = db.Select(&ProfileQuestionnaireHeaders, sqlString, orgID)

	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireHeaders, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return ProfileQuestionnaireHeaders, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(ProfileQuestionnaireHeaders)) + " records found"}}
}

func PostProfileQuestionnaireHeader(profileQuestionnaireHeader ProfileQuestionnaireHeader) (ProfileQuestionnaireHeader, TransactionalInformation) {
	if validateErrs := profileQuestionnaireHeader.Validate(); len(validateErrs) != 0 {
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireHeaderValidate.Error()}, ValidationErrors: validateErrs}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	if profileQuestionnaireHeader.ID == nil {
		profileQuestionnaireHeader.Version = 1
		stmt, err := db.PrepareNamed("INSERT INTO profile_questionnaire_header(code, description, priority, contact_type, business_relation_type_id, rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:code, :description, :priority, :contact_type, :business_relation_type_id, :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id) RETURNING id")
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		var id int64
		err = stmt.Get(&id, profileQuestionnaireHeader)
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		} else {
			profileQuestionnaireHeader.ID = &id
		}

	} else {
		stmt, _ := db.PrepareNamed("UPDATE profile_questionnaire_header SET " +
			"code = :code," +
			"description = :description," +
			"priority = :priority," +
			"contact_type = :contact_type," +
			"business_relation_type_id = :business_relation_type_id, " +
			"status = :status," +
			"version = :version + 1," +
			"rec_modified_by = :rec_modified_by, rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

		result, err := stmt.Exec(profileQuestionnaireHeader)
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		changes, err := result.RowsAffected()
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		if changes == 0 {
			return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireHeaderNotFound.Error()}}
		}
	}

	profileQuestionnaireHeader, _ = GetProfileQuestionnaireHeaderByID(*profileQuestionnaireHeader.ID)
	return profileQuestionnaireHeader, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

// GetProfileQuestionnaireHeaderByID returns the ProfileQuestionnaireHeader that the given id corresponds to. If no ProfileQuestionnaireHeader is found, an
// error is thrown.
func GetProfileQuestionnaireHeaderByID(id int64) (ProfileQuestionnaireHeader, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	profileQuestionnaireHeader := ProfileQuestionnaireHeader{}
	err = db.Get(&profileQuestionnaireHeader, "SELECT profile_questionnaire_header.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization,"+
		"COALESCE(business_relation_type.code, '') as business_relation_type_code"+
		"	FROM profile_questionnaire_header "+
		"		INNER JOIN user_profile as user_created ON profile_questionnaire_header.rec_created_by = user_created.id "+
		"		INNER JOIN user_profile as user_modified ON profile_questionnaire_header.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON profile_questionnaire_header.organization_id = organization.id "+
		"		LEFT JOIN business_relation_type as business_relation_type ON profile_questionnaire_header.business_relation_type_id = business_relation_type.id "+
		"	WHERE profile_questionnaire_header.id=$1", id)

	if err != nil && err == sql.ErrNoRows {
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireHeaderNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	//errs := profileQuestionnaireHeader.getDetails()
	//if errs != nil {
	//	return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: errs}
	//}
	return profileQuestionnaireHeader, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

// GetProfileQuestionnaireHeaderByCode returns the ProfileQuestionnaireHeader that the given id corresponds to.
// If no ProfileQuestionnaireHeader is found, an error is thrown.
func GetProfileQuestionnaireHeaderByCode(code string, orgID int64) (ProfileQuestionnaireHeader, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	org, _ := GetOrganizationByID(orgID)

	profileQuestionnaireHeader := ProfileQuestionnaireHeader{}
	err = db.Get(&profileQuestionnaireHeader, "SELECT profile_questionnaire_header.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"COALESCE(business_relation_type.code, '') as business_relation_type_code"+
		"	FROM profile_questionnaire_header "+
		"		INNER JOIN user_profile as user_created ON profile_questionnaire_header.rec_created_by = user_created.id "+
		"		INNER JOIN user_profile as user_modified ON profile_questionnaire_header.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON profile_questionnaire_header.organization_id = organization.id "+
		"		LEFT JOIN business_relation_type as business_relation_type ON profile_questionnaire_header.business_relation_type_id = business_relation_type.id "+
		"	WHERE profile_questionnaire_header.code=$1 and profile_questionnaire_header.client_id=$2", code, org.ClientID)

	if err != nil && err == sql.ErrNoRows {
		log.Error(err)
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireHeaderNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	errs := profileQuestionnaireHeader.getDetails()
	if errs != nil {
		return ProfileQuestionnaireHeader{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: errs}
	}
	return profileQuestionnaireHeader, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

func DeleteProfileQuestionnaireHeaderById(orgID int64, ids []string) TransactionalInformation {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	query, args, err := sqlx.In("DELETE FROM profile_questionnaire_header "+
		" WHERE profile_questionnaire_header.id IN (?) and profile_questionnaire_header.organization_id=?", ids, orgID)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err = db.Exec(query, args...)
	if err != nil && err == sql.ErrNoRows {
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireHeaderNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}
