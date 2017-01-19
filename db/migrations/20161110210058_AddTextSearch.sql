
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE business_relation_sector RENAME COLUMN name TO description;
ALTER TABLE business_relation_type RENAME COLUMN name TO description;

CREATE TABLE IF NOT EXISTS textsearch
(
    id bigint NOT NULL DEFAULT id_generator(),
    textsearch_object character varying NOT NULL,
    textsearch_value tsvector NOT NULL,
    client_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    CONSTRAINT pk_textsearch PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_textsearch_object
    ON textsearch USING btree
    (textsearch_object, id);

CREATE INDEX IF NOT EXISTS idx_textsearch_object_organization_id
    ON textsearch USING btree
    (textsearch_object, organization_id, client_id);

CREATE INDEX IF NOT EXISTS idx_textsearch_value
    ON textsearch USING gist
    (textsearch_value);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION textsearch_udate_trigger()
  RETURNS trigger AS
$BODY$
DECLARE
BEGIN
    CASE TG_TABLE_NAME
        WHEN 'xxxx' THEN
            RETURN NEW;
        ELSE
            IF (TG_OP = 'DELETE') THEN
                DELETE FROM textsearch WHERE textsearch_object = TG_TABLE_NAME AND id = OLD.id;
                RETURN OLD;
            ELSE
                INSERT INTO textsearch(id, textsearch_object, textsearch_value, client_id, organization_id) VALUES
                (NEW.id, TG_TABLE_NAME, to_tsvector(coalesce(NEW.code, '') || ' ' || coalesce(NEW.description, '')), NEW.client_id, NEW.organization_id)
                ON CONFLICT ON CONSTRAINT pk_textsearch DO UPDATE SET textsearch_value = to_tsvector(coalesce(NEW.code, '') || ' ' || coalesce(NEW.description, ''));   
                
                RETURN NEW;
            END IF;    
    END CASE;
	RETURN NULL; -- result is ignored since this is an AFTER trigger
END

$BODY$
  LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER textsearch_udate
  AFTER INSERT OR UPDATE OR DELETE
  ON business_relation_sector
  FOR EACH ROW
  EXECUTE PROCEDURE textsearch_udate_trigger();

CREATE TRIGGER textsearch_udate
  AFTER INSERT OR UPDATE OR DELETE
  ON business_relation_type
  FOR EACH ROW
  EXECUTE PROCEDURE textsearch_udate_trigger();

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE business_relation_sector RENAME COLUMN description TO name;
ALTER TABLE business_relation_type RENAME COLUMN description TO name;
DROP TRIGGER textsearch_udate ON business_relation_sector;
DROP TRIGGER textsearch_udate ON business_relation_type;
DROP FUNCTION textsearch_udate_trigger();
DROP TABLE textsearch;
