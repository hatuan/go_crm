package models

import (
	"github.com/shopspring/decimal"
)

type Rating struct {
	ID                                   *int64          `db:"id" json:",string"`
	Type                                 int8            `db:"type"`
	ProfileQuestionnaireHeaderID         int64           `db:"profile_questionnaire_header_id" json:",string"`
	ProfileQuestionnaireHeaderCode       string          `db:"profile_questionnaire_header_code"`
	ProfileQuestionnaireLineID           int64           `db:"profile_questionnaire_line_id" json:",string"`
	RatingProfileQuestionnaireHeaderID   int64           `db:"rating_profile_questionnaire_header_id" json:",string"`
	RatingProfileQuestionnaireHeaderCode string          `db:"rating_profile_questionnaire_header_code"`
	RatingProfileQuestionnaireLineID     int64           `db:"rating_profile_questionnaire_line_id" json:",string"`
	Points                               decimal.Decimal `db:"points" json:",string"`
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
