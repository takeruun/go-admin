
-- +migrate Up
ALTER TABLE products ADD category_id bigint(20) NOT NULL AFTER product_type;

-- +migrate Down
ALTER TABLE products DROP COLUMN category_id;