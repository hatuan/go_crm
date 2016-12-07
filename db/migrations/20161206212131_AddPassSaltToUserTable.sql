
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE "user" ADD COLUMN salt character varying;
UPDATE "user" SET salt = '';
ALTER TABLE "user" ALTER COLUMN salt SET NOT NULL;

/*random between 10000-99999*/
/*UPDATE "user" SET salt = floor(random()*(99999-1000+1))+10000;*/

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE "user" DROP COLUMN salt;
