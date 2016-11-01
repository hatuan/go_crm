
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS "user" (
  password_question character varying NOT NULL,
  password_answer character varying NOT NULL,
  password character varying NOT NULL,
  organization_id uuid,
  name character varying NOT NULL,
  last_modified_date timestamp(6) with time zone NOT NULL,
  last_login_ip character varying NOT NULL,
  last_login_date timestamp(6) with time zone NOT NULL,
  last_locked_out_reason character varying NOT NULL,
  last_locked_out_date timestamp(6) with time zone NOT NULL,
  is_locked_out boolean NOT NULL,
  is_activated boolean NOT NULL,
  id uuid NOT NULL,
  full_name character varying NOT NULL,
  email character varying NOT NULL,
  created_date timestamp(6) with time zone NOT NULL,
  comment character varying NOT NULL,
  client_id uuid,
  culture_ui_id character varying NOT NULL,
  CONSTRAINT pk_usr PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON "user" USING btree (email);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_name ON "user" USING btree (name);

CREATE TABLE IF NOT EXISTS role
(
  client_id uuid NOT NULL,
  description character varying NOT NULL,
  id uuid NOT NULL,
  name character varying NOT NULL,
  organization_id uuid NOT NULL,
  rec_created_by uuid NOT NULL,
  rec_modified_by uuid NOT NULL,
  status smallint NOT NULL,
  version bigint NOT NULL,
  rec_created_at timestamp(6) with time zone NOT NULL,
  rec_modified_at timestamp(6) with time zone NOT NULL,
  CONSTRAINT pk_role PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_in_role
(
  id uuid NOT NULL,
  role_id uuid NOT NULL,
  user_id uuid NOT NULL,
  version bigint NOT NULL,
  CONSTRAINT pk_user_in_role PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS "idx_userinrole_roleid" ON user_in_role USING btree (role_id);
CREATE INDEX IF NOT EXISTS "idx_userinrol_userid"  ON user_in_role USING btree (user_id);

CREATE TABLE IF NOT EXISTS client
(
  version bigint NOT NULL,
  rec_created_by uuid NOT NULL,
  name character varying NOT NULL,
  is_activated boolean NOT NULL,
  id uuid NOT NULL,
  culture_id character varying NOT NULL,
  amount_decimal_places smallint NOT NULL,
  amount_rounding_precision numeric NOT NULL,
  "unit-amount_decimal_places" smallint NOT NULL,
  "unit-amount_rounding_precision" numeric NOT NULL,
  currency_lcy_id uuid,
  rec_modified_by uuid NOT NULL,
  rec_created_at timestamp(6) with time zone NOT NULL,
  rec_modified_at timestamp(6) with time zone NOT NULL,
  CONSTRAINT pk_client PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_client_currency_lcy ON client USING btree (currency_lcy_id);

CREATE TABLE IF NOT EXISTS organization
(
  version bigint NOT NULL,
  status smallint NOT NULL,
  rec_modified_by uuid NOT NULL,
  rec_created_by uuid NOT NULL,
  name character varying NOT NULL,
  id uuid NOT NULL,
  code character varying NOT NULL,
  client_id uuid NOT NULL,
  rec_created_at timestamp(6) with time zone NOT NULL,
  rec_modified_at timestamp(6) with time zone NOT NULL,
  CONSTRAINT pk_organization PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_organization_client_id ON organization USING btree (client_id);

CREATE UNIQUE INDEX IF NOT EXISTS idx_organization_code ON organization USING btree (code, client_id);

CREATE TABLE IF NOT EXISTS number_sequence
(
  client_id uuid NOT NULL,
  code character varying NOT NULL,
  current_no integer NOT NULL,
  ending_no integer NOT NULL,
  format_no character varying NOT NULL,
  id uuid NOT NULL,
  is_default boolean NOT NULL,
  manual boolean NOT NULL,
  name character varying NOT NULL,
  no_seq_name character varying NOT NULL,
  organization_id uuid NOT NULL,
  rec_created_by uuid NOT NULL,
  rec_modified_by uuid NOT NULL,
  starting_no integer NOT NULL,
  status smallint NOT NULL,
  version bigint NOT NULL,
  rec_created_at timestamp(6) with time zone NOT NULL,
  rec_modified_at timestamp(6) with time zone NOT NULL,
  CONSTRAINT pk_number_sequence PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_number_sequence_code ON number_sequence USING btree (code, client_id);

CREATE TABLE IF NOT EXISTS currency
(
  version bigint NOT NULL,
  status smallint NOT NULL,
  rec_modified_by uuid NOT NULL,
  rec_created_by uuid NOT NULL,
  organization_id uuid NOT NULL,
  name character varying NOT NULL,
  id uuid NOT NULL,
  code character varying NOT NULL,
  client_id uuid NOT NULL,
  rec_created_at timestamp(6) with time zone NOT NULL,
  rec_modified_at timestamp(6) with time zone NOT NULL,
  CONSTRAINT pk_currency PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_currency_code ON currency USING btree (code, client_id);

CREATE TABLE IF NOT EXISTS currency_convert_rate
(
  client_id uuid NOT NULL,
  currency_id uuid NOT NULL,
  id uuid NOT NULL,
  organization_id uuid NOT NULL,
  rec_created_by uuid NOT NULL,
  rec_modified_by uuid NOT NULL,
  status smallint NOT NULL,
  valid_from date NOT NULL,
  version bigint NOT NULL,
  exchange_rate_amount bigint NOT NULL,
  relational_exch_rate_amount bigint NOT NULL,
  currency_relational_id uuid NOT NULL,
  rec_created_at timestamp(6) with time zone NOT NULL,
  rec_modified_at timestamp(6) with time zone NOT NULL,
  CONSTRAINT pk_currency_convert_rate PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS session
(
    working_date date NOT NULL,
    warehouse_id uuid,
    user_id uuid NOT NULL,
    organization_id uuid,
    last_time timestamp with time zone NOT NULL,
    id uuid NOT NULL,
    expire boolean,
    client_id uuid NOT NULL,
    CONSTRAINT pk_session PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_session_user ON session USING btree (user_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "user";
DROP TABLE role;
DROP TABLE user_in_role;
DROP TABLE client;
DROP TABLE organization;
DROP TABLE number_sequence;
DROP TABLE currency;
DROP TABLE currency_convert_rate;
DROP TABLE session;