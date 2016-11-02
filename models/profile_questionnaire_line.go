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
	"github.com/shopspring/decimal"
)

type ProfileQuestionnaireLine struct {
	ID                             string          `db:"id"`
	ProfileQuestionnaireHeaderID   string          `db:"profile_questionnaire_header_id"`
	ProfileQuestionnaireHeaderCode string          `db:"profile_questionnaire_header_code"`
	LineNo                         int64           `db:"line_no"`
	Description                    string          `db:"description"`
	MultipleAnswers                int8            `db:"multiple_answers"`
	AutoContactClassification      int8            `db:"auto_contact_classification"`
	Priority                       int8            `db:"priority"`
	CustomerClassField             int8            `db:"customer_class_field"`
	VendorClassField               int8            `db:"vendor_class_field"`
	ContactClassField              int8            `db:"contact_class_field"`
	StartingDateFormula            string          `db:"starting_date_formula"`
	EndingDateFormula              string          `db:"ending_date_formula"`
	ClassificationMethod           int8            `db:"classification_method"`
	SortingMethod                  int8            `db:"sorting_method"`
	FromValue                      decimal.Decimal `db:"from_value"`
	ToValue                        decimal.Decimal `db:"to_value"`
	RecCreatedByID                 string          `db:"rec_created_by"`
	RecCreatedByUser               string          `db:"rec_created_by_user"`
	RecCreated                     *Timestamp      `db:"rec_created_at"`
	RecModifiedByID                string          `db:"rec_modified_by"`
	RecModifiedByUser              string          `db:"rec_modified_by_user"`
	RecModified                    *Timestamp      `db:"rec_modified_at"`
	Status                         int8            `db:"status"`
	Version                        int16           `db:"version"`
	ClientID                       string          `db:"client_id"`
	OrganizationID                 string          `db:"organization_id"`
	Organization                   string          `db:"organization"`
}

// ErrProfileQuestionnaireLineNotFound indicates there was no ProfileQuestionnaireLine
var ErrProfileQuestionnaireLineNotFound = errors.New("ProfileQuestionnaireLine not found")

// ErrProfileQuestionnaireLineDescriptionNotSpecified indicates there was no description given by the user
var ErrProfileQuestionnaireLineDescriptionNotSpecified = errors.New("ProfileQuestionnaireLine's description not specified")

// ErrProfileQuestionnaireLineCodeNotSpecified indicates there was no code given by the user
var ErrProfileQuestionnaireLineCodeNotSpecified = errors.New("ProfileQuestionnaireLine's code not specified")

// ErrProfileQuestionnaireLineCodeDuplicate indicates there was duplicate of code given by the user
var ErrProfileQuestionnaireLineCodeDuplicate = errors.New("ProfileQuestionnaireLine's code is duplicate")

// ErrProfileQuestionnaireLineFatal indicates there was fatal error
var ErrProfileQuestionnaireLineFatal = errors.New("ProfileQuestionnaireLine has fatal error")

// ErrProfileQuestionnaireLineValidate indicates there was validate error
var ErrProfileQuestionnaireLineValidate = errors.New("ProfileQuestionnaireLine has validate error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *ProfileQuestionnaireLine) Validate() map[string]InterfaceArray {
	validationErrors := make(map[string]InterfaceArray)

	if c.Description == "" {
		validationErrors["Description"] = append(validationErrors["Description"], ErrProfileQuestionnaireLineDescriptionNotSpecified.Error())
	}
	return validationErrors
}

func GetProfileQuestionnaireLines(orgID string, searchCondition string, infiniteScrollingInformation InfiniteScrollingInformation) ([]ProfileQuestionnaireLine, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT profile_questionnaire_line.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization" +
		" profile_questionnaire_header.code as profile_questionnaire_header_code" +
		" FROM profile_questionnaire_line " +
		" INNER JOIN \"user\" as user_created ON profile_questionnaire_line.rec_created_by = user_created.id " +
		" INNER JOIN \"user\" as user_modified ON profile_questionnaire_line.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON profile_questionnaire_line.organization_id = organization.id " +
		" INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON profile_questionnaire_line.profile_questionnaire_header_id = profile_questionnaire_header.id "

	sqlWhere := " WHERE profile_questionnaire_line.organization_id = $1"
	if len(searchCondition) > 0 {
		sqlWhere += fmt.Sprintf(" AND %s", searchCondition)
	}

	var sqlOrder string
	if len(infiniteScrollingInformation.SortDirection) == 0 || infiniteScrollingInformation.SortDirection == "ASC" {
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s ASC", "profile_questionnaire_line."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	} else { //sort DESC
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s DESC", "profile_questionnaire_line."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	}

	sqlLimit := ""
	if len(infiniteScrollingInformation.FetchSize) > 0 {
		sqlLimit += fmt.Sprintf(" LIMIT %s ", infiniteScrollingInformation.FetchSize)
	}
	sqlString += sqlWhere + sqlOrder + sqlLimit
	log.Debug(sqlString)

	ProfileQuestionnaireLines := []ProfileQuestionnaireLine{}
	err = db.Select(&ProfileQuestionnaireLines, sqlString, orgID)

	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireLines, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return ProfileQuestionnaireLines, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(ProfileQuestionnaireLines)) + " records found"}}
}

func PostProfileQuestionnaireLine(profileQuestionnaireLine ProfileQuestionnaireLine) (ProfileQuestionnaireLine, TransactionalInformation) {
	if validateErrs := profileQuestionnaireLine.Validate(); len(validateErrs) != 0 {
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireLineValidate.Error()}, ValidationErrors: validateErrs}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	if profileQuestionnaireLine.ID == "" {
		profileQuestionnaireLine.ID = uuid.NewV4().String()
		profileQuestionnaireLine.Version = 1
		stmt, _ := db.PrepareNamed("INSERT INTO profile_questionnaire_line(id, profile_questionnaire_header_id, line_no," +
			" description, multiple_answers, auto_contact_classification, priority," +
			" customer_class_field, vendor_class_field, contact_class_field," +
			" starting_date_formula, ending_date_formula, classification_method, sorting_method, from_value, to_value," +
			" rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:id, :profile_questionnaire_header_id, :line_no," +
			" :description, :multiple_answers, :auto_contact_classification, :priority," +
			" :customer_class_field, :vendor_class_field, :contact_class_field," +
			" :starting_date_formula, :ending_date_formula, :classification_method, :sorting_method, :from_value, :to_value, " +
			" :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id)")
		_, err := stmt.Exec(profileQuestionnaireLine)
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}

	} else {
		stmt, _ := db.PrepareNamed("UPDATE profile_questionnaire_line SET " +
			"code = :code," +
			"profile_questionnaire_header_id = :profile_questionnaire_header_id," +
			"decription = :decription," +
			"multiple_answers = :multiple_answers," +
			"auto_contact_classification = :auto_contact_classification," +
			"priority = :priority," +
			"customer_class_field = :customer_class_field," +
			"vendor_class_field = :vendor_class_field," +
			"contact_class_field = :contact_class_field," +
			"starting_date_formula = :starting_date_formula, " +
			"ending_date_formula = :ending_date_formula, " +
			"classification_method = :classification_method, " +
			"sorting_method = :sorting_method, " +
			"from_value = :from_value, " +
			"to_value = :to_value, " +
			"status = :status," +
			"version = :version + 1," +
			"rec_modified_by = :rec_modified_by," +
			"rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

		result, err := stmt.Exec(profileQuestionnaireLine)
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		changes, err := result.RowsAffected()
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		if changes == 0 {
			return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireLineNotFound.Error()}}
		}
	}

	profileQuestionnaireLine, _ = GetProfileQuestionnaireLineByID(profileQuestionnaireLine.ID)
	return profileQuestionnaireLine, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

// GetProfileQuestionnaireLineByID returns the ProfileQuestionnaireLine that the given id corresponds to. If no ProfileQuestionnaireLine is found, an
// error is thrown.
func GetProfileQuestionnaireLineByID(id string) (ProfileQuestionnaireLine, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	profileQuestionnaireLine := ProfileQuestionnaireLine{}
	err = db.Get(&profileQuestionnaireLine, "SELECT profile_questionnaire_line.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization,"+
		"profile_questionnaire_header.code as profile_questionnaire_header_code"+
		"	FROM profile_questionnaire_line "+
		"		INNER JOIN \"user\" as user_created ON profile_questionnaire_line.rec_created_by = user_created.id "+
		"		INNER JOIN \"user\" as user_modified ON profile_questionnaire_line.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON profile_questionnaire_line.organization_id = organization.id "+
		"		INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON profile_questionnaire_line.profile_questionnaire_header_id = profile_questionnaire_header.id "+
		"	WHERE profile_questionnaire_line.id=$1", id)

	if err != nil && err == sql.ErrNoRows {
		log.Error(err)
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireLineNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return profileQuestionnaireLine, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

func DeleteProfileQuestionnaireLineById(orgID string, ids []string) TransactionalInformation {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	query, args, err := sqlx.In("DELETE FROM profile_questionnaire_line "+
		" WHERE profile_questionnaire_line.id IN (?) and profile_questionnaire_line.organization_id=?", ids, orgID)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err = db.Exec(query, args...)
	if err != nil && err == sql.ErrNoRows {
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireLineNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}
