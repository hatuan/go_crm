
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
    
END
$$;
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

