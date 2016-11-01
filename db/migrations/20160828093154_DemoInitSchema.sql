
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +goose StatementBegin
DO $$
DECLARE
    client_id varchar = '28cc612c-807d-458d-91e7-f759080b0e40';
    org_root_id varchar = '4336fecf-8c21-4531-afe6-76d34603ea34';
    org_hq_id varchar = '876a2286-4907-4a7f-b841-5cf7fd4c1288';
    org_hno_id varchar = '3ff91ce8-6ad5-4734-a0d6-21495b339bcd';
    org_hcm_id varchar = 'e190641f-d3b7-4bce-b405-36f3cfb86dda'; 
    user_admin_id varchar = '4e7739e3-939a-4181-b468-c35bdbf7a7ef';
    user_demo_id varchar = '5e6af2aa-e21a-4afd-815e-0cc3dbefa08a';
BEGIN
    
END
$$;
-- +goose StatementEnd

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

