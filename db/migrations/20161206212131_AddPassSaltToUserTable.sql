
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

/*random between 10000-99999*/
/*UPDATE "user" SET salt = floor(random()*(99999-1000+1))+10000;*/

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
