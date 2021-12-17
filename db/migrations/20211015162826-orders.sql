
-- +migrate Up
CREATE TABLE IF NOT EXISTS orders (
  id bigint(20) AUTO_INCREMENT,
  user_id bigint(20) NOT NULL,
  status int(11),
  date_of_visit DATETIME,
  date_of_exit DATETIME,
  sub_total_price int(11),
  total_price int(11),
  employee_id bigint(20),
  payment_method int(11),
  discount_type INT(11),
  discount_method INT(11),
  discount_amount INT(11),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  KEY index_orders_on_user_id (user_id)
  );
    
-- +migrate Down
DROP TABLE IF EXISTS orders;