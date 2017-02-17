package models

import (
	"database/sql"
	"erpvietnam/crm/log"
	"erpvietnam/crm/settings"
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type ProfileQuestionnaireLine struct {
	ID                             *int64          `db:"id" json:",string"`
	Type                           int8            `db:"type"`
	ProfileQuestionnaireHeaderID   int64           `db:"profile_questionnaire_header_id" json:",string"`
	ProfileQuestionnaireHeaderCode string          `db:"profile_questionnaire_header_code"`
	LineNo                         int64           `db:"line_no" json:",string"`
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
	FromValue                      decimal.Decimal `db:"from_value" json:",string"`
	ToValue                        decimal.Decimal `db:"to_value" json:",string"`
	Ratings                        []Rating        `db:"-"`
	RecCreatedByID                 int64           `db:"rec_created_by" json:",string"`
	RecCreatedByUser               string          `db:"rec_created_by_user"`
	RecCreated                     *Timestamp      `db:"rec_created_at"`
	RecModifiedByID                int64           `db:"rec_modified_by" json:",string"`
	RecModifiedByUser              string          `db:"rec_modified_by_user"`
	RecModified                    *Timestamp      `db:"rec_modified_at"`
	Status                         int8            `db:"status"`
	Version                        int16           `db:"version"`
	ClientID                       int64           `db:"client_id" json:",string"`
	OrganizationID                 int64           `db:"organization_id" json:",string"`
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

func (c *ProfileQuestionnaireLine) GetRatings() error {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		c.Ratings = []Rating{}
		return err
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
	err = db.Select(&ratings, sqlString, c.ProfileQuestionnaireHeaderID, c.ID)

	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		c.Ratings = []Rating{}
		return err
	}
	c.Ratings = ratings
	return nil
}

func (c *ProfileQuestionnaireLine) GetNextQuestion() (ProfileQuestionnaireLine, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT question_line.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization, " +
		" profile_questionnaire_header.code as profile_questionnaire_header_code" +
		" FROM profile_questionnaire_line as question_line" +
		" INNER JOIN user_profile as user_created ON question_line.rec_created_by = user_created.id " +
		" INNER JOIN user_profile as user_modified ON question_line.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON question_line.organization_id = organization.id " +
		" INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON question_line.profile_questionnaire_header_id = profile_questionnaire_header.id " +
		" WHERE question_line.organization_id = $1 " +
		" 	AND question_line.profile_questionnaire_header_id = $2 " +
		"	AND question_line.ln > $3 " +
		"	AND question_line.type = 1 "

	var next_question ProfileQuestionnaireLine
	err = db.Select(&next_question, sqlString, c.OrganizationID, c.ProfileQuestionnaireHeaderID, c.LineNo)

	if err == sql.ErrNoRows {
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{err.Error()}}
	} else if err != nil {
		log.Error(err)
		return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return next_question, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Record Found"}}

}

func (c *ProfileQuestionnaireLine) GetAnswers() ([]ProfileQuestionnaireLine, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	nextQuestion, transInfo := c.GetNextQuestion()
	if !transInfo.ReturnStatus {
		return []ProfileQuestionnaireLine{}, transInfo
	}

	sqlString := "SELECT profile_questionnaire_line.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization, " +
		" profile_questionnaire_header.code as profile_questionnaire_header_code" +
		" FROM profile_questionnaire_line " +
		" 	INNER JOIN user_profile as user_created ON profile_questionnaire_line.rec_created_by = user_created.id " +
		" 	INNER JOIN user_profile as user_modified ON profile_questionnaire_line.rec_modified_by = user_modified.id " +
		" 	INNER JOIN organization as organization ON profile_questionnaire_line.organization_id = organization.id " +
		" 	INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON profile_questionnaire_line.profile_questionnaire_header_id = profile_questionnaire_header.id " +
		" WHERE profile_questionnaire_line.organization_id = $1 " +
		" 	AND profile_questionnaire_line.profile_questionnaire_header_id = $2 " +
		"	AND profile_questionnaire_line.line_no > $3"

	if nextQuestion.ID != nil {
		sqlString += "	AND profile_questionnaire_line.line_no < $4 "
	}
	sqlString += " ORDER BY profile_questionnaire_line.line_no"

	profileQuestionnaireLines := []ProfileQuestionnaireLine{}
	if nextQuestion.ID != nil {
		err = db.Select(&profileQuestionnaireLines, sqlString, c.OrganizationID, c.ProfileQuestionnaireHeaderID, c.LineNo, nextQuestion.LineNo)
	} else {
		err = db.Select(&profileQuestionnaireLines, sqlString, c.OrganizationID, c.ProfileQuestionnaireHeaderID, c.LineNo)
	}

	if err != nil {
		log.Error(err)
		return profileQuestionnaireLines, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	return profileQuestionnaireLines, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(profileQuestionnaireLines)) + " records found"}}
}

func GetProfileQuestionnaireLinesByHeaderID(headerID int64) ([]ProfileQuestionnaireLine, TransactionalInformation) {
	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()

	sqlString := "SELECT profile_questionnaire_line.*, " +
		" user_created.name as rec_created_by_user, " +
		" user_modified.name as rec_modified_by_user, " +
		" organization.name as organization, " +
		" profile_questionnaire_header.code as profile_questionnaire_header_code" +
		" FROM profile_questionnaire_line " +
		" INNER JOIN user_profile as user_created ON profile_questionnaire_line.rec_created_by = user_created.id " +
		" INNER JOIN user_profile as user_modified ON profile_questionnaire_line.rec_modified_by = user_modified.id " +
		" INNER JOIN organization as organization ON profile_questionnaire_line.organization_id = organization.id " +
		" INNER JOIN profile_questionnaire_header as profile_questionnaire_header ON profile_questionnaire_line.profile_questionnaire_header_id = profile_questionnaire_header.id " +
		" WHERE profile_questionnaire_line.profile_questionnaire_header_id = $1" +
		" ORDER BY profile_questionnaire_line.line_no"

	profileQuestionnaireLines := []ProfileQuestionnaireLine{}
	err = db.Select(&profileQuestionnaireLines, sqlString, headerID)

	if err != nil {
		log.Error(err)
		return profileQuestionnaireLines, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}

	//for _, profileQuestionnaireLine := range profileQuestionnaireLines {
	//	if err = profileQuestionnaireLine.GetRatings(); err != nil {
	//		return profileQuestionnaireLines, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	//	}
	//}

	return profileQuestionnaireLines, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{strconv.Itoa(len(profileQuestionnaireLines)) + " records found"}}
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

	if profileQuestionnaireLine.ID == nil {
		profileQuestionnaireLine.Version = 1
		stmt, _ := db.PrepareNamed("INSERT INTO profile_questionnaire_line(type, profile_questionnaire_header_id, line_no," +
			" description, multiple_answers, auto_contact_classification, priority," +
			" customer_class_field, vendor_class_field, contact_class_field," +
			" starting_date_formula, ending_date_formula, classification_method, sorting_method, from_value, to_value," +
			" rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:type, :profile_questionnaire_header_id, :line_no," +
			" :description, :multiple_answers, :auto_contact_classification, :priority," +
			" :customer_class_field, :vendor_class_field, :contact_class_field," +
			" :starting_date_formula, :ending_date_formula, :classification_method, :sorting_method, :from_value, :to_value, " +
			" :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id) RETURNING id")
		var id int64
		err := stmt.Get(&id, profileQuestionnaireLine)
		if err != nil {
			log.Error(err)
			return ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		} else {
			profileQuestionnaireLine.ID = &id
		}

	} else {
		stmt, _ := db.PrepareNamed("UPDATE profile_questionnaire_line SET " +
			" type = :type, " +
			" code = :code, " +
			" profile_questionnaire_header_id = :profile_questionnaire_header_id, " +
			" decription = :decription, " +
			" multiple_answers = :multiple_answers, " +
			" auto_contact_classification = :auto_contact_classification, " +
			" priority = :priority, " +
			" customer_class_field = :customer_class_field, " +
			" vendor_class_field = :vendor_class_field, " +
			" contact_class_field = :contact_class_field, " +
			" starting_date_formula = :starting_date_formula, " +
			" ending_date_formula = :ending_date_formula, " +
			" classification_method = :classification_method, " +
			" sorting_method = :sorting_method, " +
			" from_value = :from_value, " +
			" to_value = :to_value, " +
			" status = :status," +
			" version = :version + 1," +
			" rec_modified_by = :rec_modified_by," +
			" rec_modified_at = :rec_modified_at WHERE id = :id AND version = :version")

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

	profileQuestionnaireLine, _ = GetProfileQuestionnaireLineByID(*profileQuestionnaireLine.ID)
	return profileQuestionnaireLine, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

func PostProfileQuestionnaireLines(headerID int64, profileQuestionnaireLines []ProfileQuestionnaireLine) ([]ProfileQuestionnaireLine, TransactionalInformation) {
	validationErrors := make(map[string]InterfaceArray)
	_ids := []int64{}

	for key, profileQuestionnaireLine := range profileQuestionnaireLines {
		if lineValidateErrs := profileQuestionnaireLine.Validate(); len(lineValidateErrs) != 0 {
			validationErrors["ProfileQuestionnaireLine"+strconv.Itoa(key)] = append(validationErrors["ProfileQuestionnaireLine"+strconv.Itoa(key)], lineValidateErrs)
		}
	}
	if len(validationErrors) != 0 {
		return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireLineValidate.Error()}, ValidationErrors: validationErrors}
	}

	db, err := sqlx.Connect(settings.Settings.Database.DriverName, settings.Settings.GetDbConn())
	if err != nil {
		log.Error(err)
		return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
	}
	defer db.Close()
	tx := db.MustBegin()

	for _, profileQuestionnaireLine := range profileQuestionnaireLines {
		stmt, _ := tx.PrepareNamed("INSERT INTO profile_questionnaire_line AS profile(type, profile_questionnaire_header_id, line_no," +
			" description, multiple_answers, auto_contact_classification, priority," +
			" customer_class_field, vendor_class_field, contact_class_field," +
			" starting_date_formula, ending_date_formula, classification_method, sorting_method, from_value, to_value," +
			" rec_created_by, rec_created_at, rec_modified_by, rec_modified_at, status, version, client_id, organization_id)" +
			" VALUES (:type, :profile_questionnaire_header_id, :line_no," +
			" :description, :multiple_answers, :auto_contact_classification, :priority," +
			" :customer_class_field, :vendor_class_field, :contact_class_field," +
			" :starting_date_formula, :ending_date_formula, :classification_method, :sorting_method, :from_value, :to_value, " +
			" :rec_created_by, :rec_created_at, :rec_modified_by, :rec_modified_at, :status, :version, :client_id, :organization_id) " +
			" ON CONFLICT ON CONSTRAINT pk_profile_questionnaire_line DO UPDATE SET " +
			" type = EXCLUDED.type, " +
			" profile_questionnaire_header_id = EXCLUDED.profile_questionnaire_header_id, " +
			" line_no = EXCLUDED.line_no, " +
			" description = EXCLUDED.description, " +
			" multiple_answers = EXCLUDED.multiple_answers, " +
			" auto_contact_classification = EXCLUDED.auto_contact_classification, " +
			" priority = EXCLUDED.priority, " +
			" customer_class_field = EXCLUDED.customer_class_field, " +
			" vendor_class_field = EXCLUDED.vendor_class_field, " +
			" contact_class_field = EXCLUDED.contact_class_field, " +
			" starting_date_formula = EXCLUDED.starting_date_formula, " +
			" ending_date_formula = EXCLUDED.ending_date_formula, " +
			" classification_method = EXCLUDED.classification_method, " +
			" sorting_method = EXCLUDED.sorting_method, " +
			" from_value = EXCLUDED.from_value, " +
			" to_value = EXCLUDED.to_value, " +
			" status = EXCLUDED.status, " +
			" version = :version + 1, " +
			" rec_modified_by = EXCLUDED.rec_modified_by, " +
			" rec_modified_at = EXCLUDED.rec_modified_at WHERE profile.version = :version RETURNING id")

		var id int64
		err := stmt.Get(&id, profileQuestionnaireLine)

		if err != nil {
			tx.Rollback()
			return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{ErrProfileQuestionnaireLineNotFound.Error()}}
		}

		if profileQuestionnaireLine.Ratings != nil {
			for _, rating := range profileQuestionnaireLine.Ratings {
				if rating.ProfileQuestionnaireLineID == nil {
					rating.ProfileQuestionnaireLineID = &id
				}
			}
			_, traninfo := PostRatingsWithLineID(profileQuestionnaireLine.ProfileQuestionnaireHeaderID, *profileQuestionnaireLine.ID, profileQuestionnaireLine.Ratings)

			if !traninfo.ReturnStatus {
				tx.Rollback()
				return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: traninfo.ReturnMessage, ValidationErrors: traninfo.ValidationErrors}
			}
		}
		_ids = append(_ids, id)
	}

	if len(_ids) != 0 {
		query, args, err := sqlx.In("DELETE FROM profile_questionnaire_line WHERE profile_questionnaire_header_id = ? AND id NOT IN (?)", headerID, _ids)
		if err != nil {
			tx.Rollback()
			return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		query = sqlx.Rebind(sqlx.DOLLAR, query)
		_, err = tx.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}

		query, args, err = sqlx.In("DELETE FROM rating WHERE profile_questionnaire_header_id = ? AND profile_questionnaire_line_id  NOT IN (?)", headerID, _ids)
		if err != nil {
			tx.Rollback()
			return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
		query = sqlx.Rebind(sqlx.DOLLAR, query)
		_, err = tx.Exec(query, args...)
		if err != nil {
			tx.Rollback()
			return []ProfileQuestionnaireLine{}, TransactionalInformation{ReturnStatus: false, ReturnMessage: []string{err.Error()}}
		}
	}
	tx.Commit()

	profileQuestionnaireLines, _ = GetProfileQuestionnaireLinesByHeaderID(headerID)
	return profileQuestionnaireLines, TransactionalInformation{ReturnStatus: true, ReturnMessage: []string{"Updated/Created successfully"}}
}

// GetProfileQuestionnaireLineByID returns the ProfileQuestionnaireLine that the given id corresponds to. If no ProfileQuestionnaireLine is found, an
// error is thrown.
func GetProfileQuestionnaireLineByID(id int64) (ProfileQuestionnaireLine, TransactionalInformation) {
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
		"		INNER JOIN user_profile as user_created ON profile_questionnaire_line.rec_created_by = user_created.id "+
		"		INNER JOIN user_profile as user_modified ON profile_questionnaire_line.rec_modified_by = user_modified.id "+
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

func DeleteProfileQuestionnaireLineById(orgID string, ids []int64) TransactionalInformation {
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
