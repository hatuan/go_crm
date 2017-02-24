package models

import (
	"biovegi/log"
	"biovegi/settings"
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Rating struct {
	ID                                   *int64          `db:"id" json:",string"`
	Type                                 int8            `db:"type"`
	ProfileQuestionnaireHeaderID         int64           `db:"profile_questionnaire_header_id" json:",string"`
	ProfileQuestionnaireHeaderCode       string          `db:"profile_questionnaire_header_code"`
	ProfileQuestionnaireLineID           *int64          `db:"profile_questionnaire_line_id" json:",string"`
	RatingProfileQuestionnaireHeaderID   int64           `db:"rating_profile_questionnaire_header_id" json:",string"`
	RatingProfileQuestionnaireHeaderCode string          `db:"rating_profile_questionnaire_header_code"`
	RatingProfileQuestionnaireLineID     int64           `db:"rating_profile_questionnaire_line_id" json:",string"`
	Points                               decimal.Decimal `db:"points"`
	RecCreatedByID                       int64           `db:"rec_created_by" json:",string"`
	RecCreatedByUser                     string          `db:"rec_created_by_user"`
	RecCreated                           *Timestamp      `db:"rec_created_at"`
	RecModifiedByID                      int64           `db:"rec_modified_by" json:",string"`
	RecModifiedByUser                    string          `db:"rec_modified_by_user"`
	RecModified                          *Timestamp      `db:"rec_modified_at"`
	Status                               int8            `db:"status"`
	Version                              int16           `db:"version"`
	ClientID                             int64           `db:"client_id" json:",string"`
	OrganizationID                       int64           `db:"organization_id" json:",string"`
	Organization                         string          `db:"organization"`
}

// ErrRatingNotFound indicates there was no Rating
var ErrRatingNotFound = errors.New("Rating not found")

// ErrRatingValidate indicates there was validate error
var ErrRatingValidate = errors.New("Rating has validate error")

// Validate checks to make sure there are no invalid fields in a submitted
func (c *Rating) Validate() map[string]InterfaceArray {
	validationErrors := make(map[string]InterfaceArray)

	//if c.Points.Cmp(decimal.Zero) < 0 {
	//	validationErrors["Points"] = append(validationErrors["Points"], ErrRatingValidate.Error())
	//}
	return validationErrors
}

func GetRatingsByHeaderID(headerID int64) ([]Rating, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT rating.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization, " +
		" profile_questionnaire_header.code as profile_questionnaire_header_code, " +
		" rating_profile_questionnaire_header.code as rating_profile_questionnaire_header_code" +
		" FROM rating " +
		" INNER JOIN user_profile as user_created ON rating.rec_created_by = user_created.id " +
		" INNER JOIN user_profile as user_modified ON rating.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON rating.organization_id = organization.id " +
		" INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON rating.profile_questionnaire_header_id = profile_questionnaire_header.id " +
		" INNER JOIN profile_questionnaire_line as profile_questionnaire_line ON rating.profile_questionnaire_line_id = profile_questionnaire_line.id " +
		" INNER JOIN profile_questionnaire_header as rating_profile_questionnaire_header ON rating.rating_profile_questionnaire_header_id = rating_profile_questionnaire_header.id " +
		" INNER JOIN profile_questionnaire_line as rating_profile_questionnaire_line ON rating.rating_profile_questionnaire_line_id = rating_profile_questionnaire_line.id " +
		" WHERE rating.profile_questionnaire_header_id = $1" +
		" ORDER BY profile_questionnaire_header.code, profile_questionnaire_line.line_no, rating_profile_questionnaire_header.code, rating_profile_questionnaire_line.line_no"

	ratings := []Rating{}
	err = db.Select(&ratings, sqlString, headerID)

	if err != nil {
		log.Error(err)
		return ratings, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return ratings, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(ratings)) + " records found"}}
}

func GetRatingsByLineID(headerID int64, lineID int64) ([]Rating, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT rating.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization, " +
		" profile_questionnaire_header.code as profile_questionnaire_header_code, " +
		" rating_profile_questionnaire_header.code as rating_profile_questionnaire_header_code" +
		" FROM rating " +
		" INNER JOIN user_profile as user_created ON rating.rec_created_by = user_created.id " +
		" INNER JOIN user_profile as user_modified ON rating.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON rating.organization_id = organization.id " +
		" INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON rating.profile_questionnaire_header_id = profile_questionnaire_header.id " +
		" INNER JOIN profile_questionnaire_line as profile_questionnaire_line ON rating.profile_questionnaire_line_id = profile_questionnaire_line.id " +
		" INNER JOIN profile_questionnaire_header as rating_profile_questionnaire_header ON rating.rating_profile_questionnaire_header_id = rating_profile_questionnaire_header.id " +
		" INNER JOIN profile_questionnaire_line as rating_profile_questionnaire_line ON rating.rating_profile_questionnaire_line_id = rating_profile_questionnaire_line.id " +
		" WHERE rating.profile_questionnaire_header_id = $1" +
		" AND rating.profile_questionnaire_line_id = $2" +
		" ORDER BY profile_questionnaire_header.code, profile_questionnaire_line.line_no, rating_profile_questionnaire_header.code, rating_profile_questionnaire_line.line_no"

	ratings := []Rating{}
	err = db.Select(&ratings, sqlString, headerID, lineID)

	if err != nil {
		log.Error(err)
		return ratings, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return ratings, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(ratings)) + " records found"}}
}

func PostRatingsWithLineID(headerID int64, lineID int64, ratings []Rating) ([]Rating, TransactionalInformation) {
	validationErrors := make(map[string]InterfaceArray)
	_ids := []int64{}

	for key, rating := range ratings {
		if lineValidateErrs := rating.Validate(); len(lineValidateErrs) != 0 {
			validationErrors["Rating"+strconv.Itoa(key)] = append(validationErrors["Rating"+strconv.Itoa(key)], lineValidateErrs)
		}
	}
	if len(validationErrors) != 0 {
		return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrRatingValidate.Error()}, ValidationErrors: validationErrors}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()
	tx := db.MustBegin()

	for _, rating := range ratings {
		stmt, _ := tx.PrepareNamed("INSERT INTO rating AS rating(id, " +
			" profile_questionnaire_header_id, " +
			" profile_questionnaire_line_id," +
			" rating_profile_questionnaire_header_id," +
			" rating_profile_questionnaire_line_id," +
			" points," +
			" rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:id, " +
			" :profile_questionnaire_header_id, " +
			" :profile_questionnaire_line_id," +
			" :rating_profile_questionnaire_header_id," +
			" :rating_profile_questionnaire_line_id," +
			" :points," +
			" :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id) " +
			" ON CONFLICT ON CONSTRAINT pk_rating DO UPDATE SET " +
			" profile_questionnaire_header_id = EXCLUDED.profile_questionnaire_header_id, " +
			" profile_questionnaire_line_id = EXCLUDED.profile_questionnaire_line_id, " +
			" rating_profile_questionnaire_header_id = EXCLUDED.rating_profile_questionnaire_header_id, " +
			" rating_profile_questionnaire_line_id = EXCLUDED.rating_profile_questionnaire_line_id, " +
			" points = EXCLUDED.points, " +
			" status = EXCLUDED.status, " +
			" version = :version + 1, " +
			" rec_modified_by = EXCLUDED.rec_modified_by, " +
			" rec_modified_at = EXCLUDED.rec_modified_at WHERE rating.version = :version RETURNING id")

		var id int64
		err := stmt.Get(&id, rating)

		if err != nil {
			tx.Rollback()
			return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrRatingNotFound.Error()}}
		}
		_ids = append(_ids, id)

	}
	if len(_ids) == 0 { // neu xoa het point phai gan gia tri mac dinh cho _ids => xoa het trong bang rating
		_ids = append(_ids, 0)
	}
	query, args, err := sqlx.In("DELETE FROM rating WHERE profile_questionnaire_header_id = ? AND profile_questionnaire_line_id = ? AND id NOT IN (?)", headerID, lineID, _ids)
	if err != nil {
		tx.Rollback()
		return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err = tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return []Rating{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	tx.Commit()

	ratings, _ = GetRatingsByLineID(headerID, lineID)
	return ratings, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}
