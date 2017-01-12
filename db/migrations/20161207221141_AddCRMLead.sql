
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS lead
(
    id uuid NOT NULL,
    number_sequence_id uuid NOT NULL,
    no character varying NOT NULL,
    description character varying NOT NULL,
    contact_type smallint NOT NULL, /*1 - Company, 2 - Person*/
    contact_name character varying NOT NULL,
    contact_person_name character varying NOT NULL, 
    street character varying NOT NULL,
    city_id character varying NOT NULL,
    county_id uuid NOT NULL,
    country_region_id uuid NOT NULL,
    state_id uuid NOT NULL,
    zip_postal_id uuid NOT NULL,
    phone character varying NOT NULL,
    phone_extension character varying NOT NULL,
    mobile character varying NOT NULL,
    sms character varying NOT NULL,
    telex character varying NOT NULL,
    fax character varying NOT NULL,
    email character varying NOT NULL,
    url character varying NOT NULL,
    pager character varying NOT NULL,
    latitude double precision NOT NULL,
    longtude double precision NOT NULL,
    timezone character varying NOT NULL,
    address_master_id uuid NOT NULL,
    date_open timestamp with time zone NOT NULL,
    date_close timestamp with time zone NOT NULL,
    user_owner_id uuid NOT NULL,
    user_open_by_id uuid NOT NULL,
    user_close_by_id uuid NOT NULL,
    priority smallint NOT NULL,
    sale_unit_id uuid NOT NULL,
    source_type_id uuid NOT NULL,
    source_id character varying NOT NULL, /*uuid or free text*/
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_lead PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_lead_no
    ON lead USING btree
    (client_id, no)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS address_master
(
    id uuid NOT NULL,
    name character varying NOT NULL, /*gia tri cua contact_name, contact_person_name*/
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_address_master PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS address
(
    id uuid NOT NULL,
    table_name character varying NOT NULL,
    table_rec_id uuid NOT NULL,
    primary_address smallint NOT NULL,
    name character varying NOT NULL,
    street character varying NOT NULL,
    city_id character varying NOT NULL,
    county_id uuid NOT NULL,
    country_region_id uuid NOT NULL,
    state_id uuid NOT NULL,
    zip_postal_id uuid NOT NULL,
    phone character varying NOT NULL,
    phone_extension character varying NOT NULL,
    mobile character varying NOT NULL,
    sms character varying NOT NULL,
    telex character varying NOT NULL,
    fax character varying NOT NULL,
    email character varying NOT NULL,
    url character varying NOT NULL,
    pager character varying NOT NULL,
    latitude double precision NOT NULL,
    longtude double precision NOT NULL,
    timezone character varying NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_address PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS profile_questionnaire_answer
(
    id uuid NOT NULL,
    table_name character varying NOT NULL,
    table_rec_id uuid NOT NULL,
    profile_questionnaire_line_id uuid NOT NULL,
    profile_questionnaire_priority smallint NOT NULL,
    profile_questionnaire_line_priority smallint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_profile_questionnaire_answer PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_profile_questionnaire_answer
    ON profile_questionnaire_answer USING btree
    (table_name, table_rec_id, profile_questionnaire_line_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS business_relation
(
    id uuid NOT NULL,
    number_sequence_id uuid NOT NULL,
    no character varying NOT NULL,
    contact_type smallint NOT NULL, /*1 - Company, 2 - Person*/
    contact_name character varying NOT NULL, /*neu type la person se co chi tiet first, last, middle ... ben duoi*/
    street character varying NOT NULL,
    city_id character varying NOT NULL,
    county_id uuid NOT NULL,
    country_region_id uuid NOT NULL,
    state_id uuid NOT NULL,
    zip_postal_id uuid NOT NULL,
    phone character varying NOT NULL,
    phone_extension character varying NOT NULL,
    mobile character varying NOT NULL,
    sms character varying NOT NULL,
    telex character varying NOT NULL,
    fax character varying NOT NULL,
    email character varying NOT NULL,
    url character varying NOT NULL,
    pager character varying NOT NULL,
    latitude double precision NOT NULL,
    longtude double precision NOT NULL,
    timezone character varying NOT NULL,
    address_master_id uuid NOT NULL,
    user_owner_id uuid NOT NULL,
    sale_unit_id uuid NOT NULL,
    source_type_id uuid NOT NULL,
    source_id character varying NOT NULL, /*uuid or free text*/
    salutation_id uuid NOT NULL, /*cach xung ho*/
    language_id uuid NOT NULL, 
    currency_id uuid NOT NULL,
    vat_no character varying NOT NULL,
    sale_district_id uuid NOT NULL, /*vung, mien (territory)*/ 
    person_first_name character varying NOT NULL,
    person_middle_name character varying NOT NULL,
    person_last_name character varying NOT NULL,
    person_birthday timestamp with time zone NOT NULL,
    person_dayofbirthday smallint NOT NULL,
    person_monthofbirthday smallint NOT NULL,
    person_yearofbirthday smallint NOT NULL,
    stopped_segment smallint NOT NULL, 
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_business_relation PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_business_relation_no
    ON business_relation USING btree
    (client_id, no)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS contact_person
(
    id uuid NOT NULL,
    table_name character varying NOT NULL,
    table_rec_id uuid NOT NULL,
    number_sequence_id uuid NOT NULL,
    no character varying NOT NULL,
    name character varying NOT NULL, 
    street character varying NOT NULL,
    city_id character varying NOT NULL,
    county_id uuid NOT NULL,
    country_region_id uuid NOT NULL,
    state_id uuid NOT NULL,
    zip_postal_id uuid NOT NULL,
    phone character varying NOT NULL,
    phone_extension character varying NOT NULL,
    mobile character varying NOT NULL,
    sms character varying NOT NULL,
    telex character varying NOT NULL,
    fax character varying NOT NULL,
    email character varying NOT NULL,
    url character varying NOT NULL,
    pager character varying NOT NULL,
    latitude double precision NOT NULL,
    longtude double precision NOT NULL,
    timezone character varying NOT NULL,
    address_master_id uuid NOT NULL,
    salutation_id uuid NOT NULL, /*cach xung ho*/
    language_id uuid NOT NULL, 
    sale_district_id uuid NOT NULL, /*vung, mien (Territory)*/ 
    person_first_name character varying NOT NULL,
    person_middle_name character varying NOT NULL,
    person_last_name character varying NOT NULL,
    person_birthday timestamp with time zone NOT NULL,
    person_dayofbirthday smallint NOT NULL,
    person_monthofbirthday smallint NOT NULL,
    person_yearofbirthday smallint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    status smallint NOT NULL,
    version bigint NOT NULL,
    rec_modified_by uuid NOT NULL,
    rec_created_by uuid NOT NULL,
    client_id uuid NOT NULL,
    organization_id uuid NOT NULL,
    CONSTRAINT pk_contact_person PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_contact_person_no
    ON contact_person USING btree
    (client_id, no)
    TABLESPACE pg_default;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE lead;
DROP TABLE address_master;
DROP TABLE address;
DROP TABLE profile_questionnaire_answer;
DROP TABLE business_relation;
DROP TABLE contact_person;
