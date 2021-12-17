
-- +migrate Up
CREATE TABLE IF NOT EXISTS order_items (
  id bigint(20) AUTO_INCREMENT,
  order_id bigint(20) NOT NULL,
  product_id bigint(20) NOT NULL,
  price INT(11),
  tax int(11),
  quantity INT(11) DEFAULT 1,
  other_person BOOLEAN DEFAULT false,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  KEY index_order_items_on_order_id (order_id)
  );
    
-- +migrate Down
DROP TABLE IF EXISTS order_items;