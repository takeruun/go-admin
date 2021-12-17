
-- +migrate Up
CREATE TABLE IF NOT EXISTS administrators (
  id int(15) AUTO_INCREMENT,
  name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(255),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  UNIQUE KEY index_administrators_on_email (email)
  );
    
-- +migrate Down
DROP TABLE IF EXISTS administrators;