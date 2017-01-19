
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS account
(
    version bigint NOT NULL,
    status smallint NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_by bigint NOT NULL,
    parent_account_id bigint,
    organization_id bigint NOT NULL,
    name character varying NOT NULL,
    level smallint NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    detail boolean NOT NULL,
    currency_id bigint,
    code character varying NOT NULL,
    client_id bigint NOT NULL,
    ar_ap boolean NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_account PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_account_code
    ON account USING btree
    (code, client_id)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_account_parent_account_id
    ON account USING btree
    (parent_account_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS general_journal_document
(
    version bigint NOT NULL,
    transaction_type integer NOT NULL,
    total_debit_amount_lcy numeric(38, 20) NOT NULL,
    total_debit_amount numeric(38, 20) NOT NULL,
    total_credit_amount_lcy numeric(38, 20) NOT NULL,
    total_credit_amount numeric(38, 20) NOT NULL,
    status integer NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_by bigint NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    document_type integer NOT NULL,
    document_posted_date date NOT NULL,
    document_no character varying NOT NULL,
    document_created_date date NOT NULL,
    description character varying NOT NULL,
    currency_id bigint,
    client_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    number_sequence_id bigint NOT NULL,
    exchange_rate_amount numeric(38, 20) NOT NULL,
    relational_exch_rate_amount numeric NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_general_journal_document PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_general_journal_document_client
    ON general_journal_document USING btree
    (client_id)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_general_journal_document_no
    ON general_journal_document USING btree
    (document_no)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_general_journal_document_organization
    ON general_journal_document USING btree
    (organization_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS general_journal_line
(
    version bigint NOT NULL,
    transaction_type integer NOT NULL,
    general_journal_document_id bigint NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_by bigint NOT NULL,
    organization_id bigint NOT NULL,
    job_id bigint,
    id bigint NOT NULL DEFAULT id_generator(),
    document_type integer NOT NULL,
    document_posted_date date NOT NULL,
    document_no character varying NOT NULL,
    document_created_date date NOT NULL,
    description character varying NOT NULL,
    debit_amount_lcy numeric(38, 20) NOT NULL,
    debit_amount numeric(38, 20) NOT NULL,
    currency_id bigint,
    currency_exchange_rate numeric(38, 20),
    credit_amount_lcy numeric(38, 20) NOT NULL,
    credit_amount numeric(38, 20) NOT NULL,
    cor_account_id bigint NOT NULL,
    client_id bigint NOT NULL,
    business_partner_id bigint,
    account_id bigint NOT NULL,
    fix_asset_id bigint,
    line_no bigint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_general_journal_line PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_general_journal_line_client
    ON general_journal_line USING btree
    (client_id)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_general_journal_line_document_no
    ON general_journal_line USING btree
    (document_no)
    TABLESPACE pg_default;    

CREATE UNIQUE INDEX IF NOT EXISTS idx_general_journal_line_no
    ON general_journal_line USING btree
    (general_journal_document_id, line_no)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_general_journal_line_organization
    ON general_journal_line USING btree
    (organization_id)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_general_journal_line_posted
    ON general_journal_line USING btree
    (document_posted_date)
    TABLESPACE pg_default;

CREATE INDEX IF NOT EXISTS idx_gnrl_jrnl_ln_gnrl_jrnl_dcmn
    ON general_journal_line USING btree
    (general_journal_document_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS general_journal_setup
(
    client_id bigint NOT NULL,
    local_currency_id bigint NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    lcy_exchange_rate_unit integer NOT NULL,
    organization_id bigint NOT NULL,
    rec_created_by bigint NOT NULL,
    rec_modified_by bigint NOT NULL,
    version bigint NOT NULL,
    general_journal_number_sequence_id bigint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_general_journal_setup PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_general_journal_setup_organization_id
    ON general_journal_setup USING btree
    (organization_id, client_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS business_partner
(
    version bigint NOT NULL,
    vendor_account_id bigint NOT NULL,
    vat_code character varying NOT NULL,
    telephone character varying NOT NULL,
    status smallint NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_by bigint NOT NULL,
    payment_term_id bigint NOT NULL,
    organization_id bigint NOT NULL,
    name character varying NOT NULL,
    e_mail character varying NOT NULL,
    is_vendor boolean NOT NULL,
    is_employee boolean NOT NULL,
    is_customer boolean NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    home_page character varying NOT NULL,
    fax character varying NOT NULL,
    employee_account_id bigint NOT NULL,
    customer_account_id bigint NOT NULL,
    credit_limit numeric(38, 20) NOT NULL,
    contact_name character varying NOT NULL,
    comment character varying NOT NULL,
    code character varying NOT NULL,
    client_id bigint NOT NULL,
    business_partner_group_id3 bigint NOT NULL,
    business_partner_group_id2 bigint NOT NULL,
    business_partner_group_id1 bigint NOT NULL,
    amount_limit numeric(38, 20) NOT NULL,
    address character varying NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_business_partner PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS business_partner_group
(
    version bigint NOT NULL,
    status smallint NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_by bigint NOT NULL,
    organization_id bigint NOT NULL,
    name character varying NOT NULL,
    level smallint NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    code character varying NOT NULL,
    client_id bigint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_business_partner_group PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_business_partner_group_code
    ON business_partner_group USING btree
    (code, client_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS job
(
    status smallint NOT NULL,
    organization_id bigint NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    business_partner_id bigint NOT NULL,
    code character varying NOT NULL,
    client_id bigint NOT NULL,
    version bigint NOT NULL,
    account_id bigint,
    amount numeric(38, 20) NOT NULL,
    amount_lcy numeric(38, 20) NOT NULL,
    comment character varying NOT NULL,
    currency_id bigint NOT NULL,
    job_end date,
    job_group_id1 bigint,
    job_group_id2 bigint,
    job_group_id3 bigint,
    job_start date,
    name character varying NOT NULL,
    rec_created_by bigint NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_job PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_job_code
    ON job USING btree
    (code, client_id)
    TABLESPACE pg_default;

CREATE TABLE IF NOT EXISTS job_group
(
    status smallint NOT NULL,
    organization_id bigint NOT NULL,
    code character varying NOT NULL,
    client_id bigint NOT NULL,
    id bigint NOT NULL DEFAULT id_generator(),
    version bigint NOT NULL,
    level smallint NOT NULL,
    name character varying NOT NULL,
    rec_created_by bigint NOT NULL,
    rec_modified_by bigint NOT NULL,
    rec_created_at timestamp with time zone NOT NULL,
    rec_modified_at timestamp with time zone NOT NULL,
    CONSTRAINT pk_job_group PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_job_group_code
    ON job_group USING btree
    (code, client_id)
    TABLESPACE pg_default;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE account;
DROP TABLE general_journal_document;
DROP TABLE general_journal_line;
DROP TABLE general_journal_setup;
DROP TABLE business_partner;
DROP TABLE business_partner_group;
DROP TABLE job;
DROP TABLE job_group;
