
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS business_relation_sector
(
    code character varying NOT NULL,
    name character varying NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    id uuid NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_business_relation_sector PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_business_relation_sector_code
    ON business_relation_sector USING btree
    (code, client_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS business_relation_type
(
    code character varying NOT NULL,
    name character varying NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    id uuid NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_business_relation_type PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_business_relation_type_code
    ON business_relation_type USING btree
    (code, client_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS profile_questionnaire_header
(
    id uuid NOT NULL,
    code character varying NOT NULL,
    description character varying NOT NULL,
    priority smallint NOT NULL,
    contact_type smallint NOT NULL,
    business_relation_type_id uuid NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_profile_questionnaire_header PRIMARY KEY (id)
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_profile_questionnaire_header_code
    ON profile_questionnaire_header USING btree
    (code, client_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS profile_questionnaire_line
(
    id uuid NOT NULL,
    profile_questionnaire_header_id uuid NOT NULL,
    line_no bigint NOT NULL,
    description character varying NOT NULL, 
    multiple_answers smallint NOT NULL,
    auto_contact_classification smallint NOT NULL,
    priority smallint NOT NULL,
    customer_class_field smallint NOT NULL,
    vendor_class_field smallint NOT NULL,
    contact_class_field smallint NOT NULL,
    starting_date_formula character varying NOT NULL,
    ending_date_formula character varying NOT NULL,
    classification_method smallint NOT NULL,
    sorting_method smallint NOT NULL,
    from_value numeric(38, 20) NOT NULL,
    to_value numeric(38, 20) NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_profile_questionnaire_line PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_profile_questionnaire_line_no
    ON profile_questionnaire_line USING btree
    (profile_questionnaire_header_id, line_no)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS rating
(
    id uuid NOT NULL,
    profile_questionnaire_line_id uuid NOT NULL,
    rating_profile_questionnaire_line_id uuid NOT NULL,
    points numeric(38, 20) NOT NULL, 
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_rating PRIMARY KEY (id)    
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_rating_line_id
    ON rating USING btree
    (profile_questionnaire_line_id, rating_profile_questionnaire_line_id)
    TABLESPACE pg_default;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE business_relation_sector;
DROP TABLE business_relation_type;
DROP TABLE profile_questionnaire_header;
DROP TABLE profile_questionnaire_line;
DROP TABLE rating;
