
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE rating ADD COLUMN profile_questionnaire_header_id bigint;
ALTER TABLE rating ADD COLUMN rating_profile_questionnaire_header_id bigint;

UPDATE rating SET 
                profile_questionnaire_header_id = 0, 
                rating_profile_questionnaire_header_id = 0;

ALTER TABLE rating ALTER profile_questionnaire_header_id SET NOT NULL;
ALTER TABLE rating ALTER rating_profile_questionnaire_header_id SET NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE rating DROP COLUMN profile_questionnaire_header_id;
ALTER TABLE rating DROP COLUMN rating_profile_questionnaire_header_id;

