
-- +migrate Up
ALTER TABLE users
  ADD family_user_id bigint(20) AFTER reason_for_coming,
  ADD family_relationship int(11) AFTER family_user_id;

-- +migrate Down
ALTER TABLE users DROP COLUMN family_user_id, family_relationship;