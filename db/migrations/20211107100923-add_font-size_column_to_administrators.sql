
-- +migrate Up
ALTER TABLE administrators 
  ADD font_size int(11) DEFAULT 100 AFTER password;

-- +migrate Down
ALTER TABLE administrators DROP COLUMN font_size;