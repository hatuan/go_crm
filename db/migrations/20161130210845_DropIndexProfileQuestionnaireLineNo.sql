
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP INDEX IF EXISTS idx_profile_questionnaire_line_no;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

