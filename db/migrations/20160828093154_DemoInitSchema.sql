
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
DO $$
DECLARE
    client_id varchar = '';
    org_root_id varchar = '';
    org_hq_id varchar = '';
    org_hno_id varchar = '';
    org_hcm_id varchar = ''; 
    user_admin_id varchar = '';
    user_demo_id varchar = '';
BEGIN
    INSERT INTO user_profile(password_question, password_answer, password, salt, organization_id, name, last_modified_date, last_login_ip, last_login_date, last_locked_out_reason, last_locked_out_date, is_locked_out, is_acvicated, full_name, email, created_date, comment, client_id, culture_ui_id)   
    VALUES
    ('') 
END
$$;
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

