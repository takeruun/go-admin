
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id bigint(20) AUTO_INCREMENT,
  family_name VARCHAR(255),
  given_name VARCHAR(255),
  family_name_kana VARCHAR(255),
  given_name_kana VARCHAR(255),
  postal_code VARCHAR(255),
  email VARCHAR(255),
  prefecture_id INT(11),
  address1 VARCHAR(255),
  address2 VARCHAR(255),
  address3 VARCHAR(255),
  phone_number VARCHAR(255),
  home_phone_number VARCHAR(255),
  gender INT(11),
  birthday DATE,
  occupation INT(11),
  first_visit_date DATETIME,
  dm_forwarding_flg BOOLEAN,
  memo TEXT,
  reason_for_coming TEXT,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY index_users_on_email (email)
  );
    
-- +migrate Down
DROP TABLE IF EXISTS users;