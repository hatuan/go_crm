
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE profile_questionnaire_line ADD COLUMN type smallint NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

