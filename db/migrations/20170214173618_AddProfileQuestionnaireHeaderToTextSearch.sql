
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TRIGGER textsearch_udate
  AFTER INSERT OR UPDATE OR DELETE
  ON profile_questionnaire_header
  FOR EACH ROW
  EXECUTE PROCEDURE textsearch_udate_trigger();

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TRIGGER textsearch_udate ON profile_questionnaire_header;
