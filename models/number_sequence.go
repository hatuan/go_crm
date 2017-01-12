package models

import (
	"biovegi/log"
	"biovegi/settings"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type NumberSequence struct {
	ID                string     `db:"id"`
	Code              string     `db:"code"`
	Description       string     `db:"description"`
	CurrentNo         int        `db:"current_no"`
	StartingNo        int        `db:"starting_no"`
	EndingNo          int        `db:"ending_no"`
	FormatNo          string     `db:"format_no"`
	IsDefault         bool       `db:"is_default"`
	Manual            bool       `db:"manual"`
	NoSequenceName    string     `db:"no_seq_name"`
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

// ErrNumberSequenceNotFound indicates there was no NumberSequence
var ErrNumberSequenceNotFound = errors.New("NumberSequence not found")

// ErrNumberSequenceDescriptionNotSpecified indicates there was no name given by the user
var ErrNumberSequenceDescriptionNotSpecified = errors.New("NumberSequence's description not specified")

// ErrNumberSequenceCodeNotSpecified indicates there was no code given by the user
var ErrNumberSequenceCodeNotSpecified = errors.New("NumberSequence's code not specified")

// ErrNumberSequenceCodeDuplicate indicates there was duplicate of code given by the user
var ErrNumberSequenceCodeDuplicate = errors.New("NumberSequence's code is duplicate")

// ErrNumberSequenceFatal indicates there was fatal error
var ErrNumberSequenceFatal = errors.New("NumberSequence has fatal error")

// ErrNumberSequenceValidate indicates there was validate error
var ErrNumberSequenceValidate = errors.New("NumberSequence has validate error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *NumberSequence) Validate() map[string]InterfaceArray {
	validationErrors := make(map[string]InterfaceArray)

	if c.Code == "" {
		validationErrors["Code"] = append(validationErrors["Code"], ErrNumberSequenceCodeNotSpecified.Error())
	}
	if c.Description == "" {
		validationErrors["Description"] = append(validationErrors["Description"], ErrNumberSequenceDescriptionNotSpecified.Error())
	}
	if c.Code != "" {
		db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
		if err != nil {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrNumberSequenceFatal.Error())
		}
		defer db.Close()
		var otherID string
		ID := EmptyUUID
		if c.ID != "" {
			ID = c.ID
		}
		err = db.Get(&otherID, "SELECT id FROM number_sequence WHERE code = $1 AND id != $2 AND client_id = $3", c.Code, ID, c.ClientID)
		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
			validationErrors["Fatal"] = append(validationErrors["Fatal"], ErrNumberSequenceFatal.Error())
		}
		if otherID != "" && err != sql.ErrNoRows {
			validationErrors["Code"] = append(validationErrors["Code"], ErrNumberSequenceCodeDuplicate.Error())
		}
	}
	return validationErrors
}

func GetNumberSequences(orgID string, searchCondition string, infiniteScrollingInformation InfiniteScrollingInformation) ([]NumberSequence, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT number_sequence.*, user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, organization.name as organization" +
		" FROM number_sequence " +
		" INNER JOIN \"user\" as user_created ON number_sequence.rec_created_by = user_created.id " +
		" INNER JOIN \"user\" as user_modified ON number_sequence.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON number_sequence.organization_id = organization.id "

	sqlWhere := " WHERE number_sequence.organization_id = $1"
	if len(searchCondition) > 0 {
		sqlWhere += fmt.Sprintf(" AND %s", searchCondition)
	}

	var sqlOrder string
	if len(infiniteScrollingInformation.SortDirection) == 0 || infiniteScrollingInformation.SortDirection == "ASC" {
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s ASC", "number_sequence."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	} else { //sort DESC
		if len(infiniteScrollingInformation.SortExpression) > 0 {
			sqlOrder = fmt.Sprintf(" ORDER BY %s DESC", "number_sequence."+strings.ToLower(infiniteScrollingInformation.SortExpression))
		}
	}
	sqlLimit := ""
	if len(infiniteScrollingInformation.FetchSize) > 0 {
		sqlLimit += fmt.Sprintf(" LIMIT %s ", infiniteScrollingInformation.FetchSize)
	}
	sqlString += sqlWhere + sqlOrder + sqlLimit
	log.Debug(sqlString)

	numberSequences := []NumberSequence{}
	err = db.Select(&numberSequences, sqlString, orgID)

	if err != nil {
		log.Error(err)
		return numberSequences, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return numberSequences, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(numberSequences)) + " records found"}}
}

func PostNumberSequence(numberSequence NumberSequence) (NumberSequence, TransactionalInformation) {
	if validateErrs := numberSequence.Validate(); len(validateErrs) != 0 {
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrNumberSequenceValidate.Error()}, ValidationErrors: validateErrs}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	if numberSequence.ID == "" {
		numberSequence.ID = uuid.NewV4().String()
		numberSequence.Version = 1
		stmt, _ := db.PrepareNamed("INSERT INTO number_sequence(id, code, description, rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:id, :code, :description, :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id)")
		_, err := stmt.Exec(numberSequence)
		if err != nil {
			log.Error(err)
			return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}

	} else {
		stmt, _ := db.PrepareNamed("UPDATE number_sequence SET " +
			"code = :code," +
			"description = :description," +
			"status = :status," +
			"version = :version + 1," +
			"rec_modified_by = :rec_modified_by, rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

		result, err := stmt.Exec(numberSequence)
		if err != nil {
			log.Error(err)
			return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		changes, err := result.RowsAffected()
		if err != nil {
			log.Error(err)
			return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		if changes == 0 {
			return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrNumberSequenceNotFound.Error()}}
		}
	}
	numberSequence, _ = GetNumberSequenceByID(numberSequence.ID)
	return numberSequence, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

// GetNumberSequenceByID returns the NumberSequence that the given id corresponds to. If no NumberSequence is found, an
// error is thrown.
func GetNumberSequenceByID(id string) (NumberSequence, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	numberSequence := NumberSequence{}
	err = db.Get(&numberSequence, "SELECT number_sequence.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM number_sequence "+
		"		INNER JOIN \"user\" as user_created ON number_sequence.rec_created_by = user_created.id "+
		"		INNER JOIN \"user\" as user_modified ON number_sequence.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON number_sequence.organization_id = organization.id "+
		"	WHERE number_sequence.id=$1", id)
	if err != nil && err == sql.ErrNoRows {
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrNumberSequenceNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return numberSequence, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

// GetNumberSequenceByCode returns the NumberSequence that the given id corresponds to.
// If no NumberSequence is found, an error is thrown.
func GetNumberSequenceByCode(code string, orgID string) (NumberSequence, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	org, _ := GetOrganizationByID(orgID)

	numberSequence := NumberSequence{}
	err = db.Get(&numberSequence, "SELECT number_sequence.*,"+
		"user_created.name as rec_created_by_user,"+
		"user_modified.name as rec_modified_by_user,"+
		"organization.name as organization"+
		"	FROM number_sequence "+
		"		INNER JOIN \"user\" as user_created ON number_sequence.rec_created_by = user_created.id "+
		"		INNER JOIN \"user\" as user_modified ON number_sequence.rec_modified_by = user_modified.id "+
		"		INNER JOIN organization as organization ON number_sequence.organization_id = organization.id "+
		"	WHERE number_sequence.code=$1 and number_sequence.client_id=$2", code, org.ClientID)

	if err != nil && err == sql.ErrNoRows {
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrNumberSequenceNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return NumberSequence{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return numberSequence, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}

func DeleteNumberSequenceById(orgID string, ids []string) TransactionalInformation {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	query, args, err := sqlx.In("DELETE FROM number_sequence "+
		" WHERE number_sequence.id IN (?) and number_sequence.organization_id=?", ids, orgID)

	query = sqlx.Rebind(sqlx.DOLLAR, query)

	_, err = db.Exec(query, args...)

	if err != nil && err == sql.ErrNoRows {
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrNumberSequenceNotFound.Error()}}
	} else if err != nil {
		log.Error(err)
		return TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	return TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Successfully"}}
}
