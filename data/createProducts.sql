-- SQLite

-- Create Tables

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS  products;
DROP TABLE IF EXISTS  discount_rules;

CREATE TABLE products
(
  SKU TEXT PRIMARY KEY NOT NULL,
  Name TEXT NOT NULL,
  Price NUMERIC NOT NULL,
  Inventory_Qty INT
);

CREATE TABLE
(
  "id" integer primary key autoincrement,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  "username" varchar(255) NOT NULL,
  "firstname" varchar(255) NOT NULL,
  "lastname" varchar(255) NOT NULL,
  "hash" varchar(255) NOT NULL 
);

CREATE TABLE IF NOT EXISTS discount_rules(
  id integer PRIMARY key AUTOINCREMENT,
  condition_qty INTEGER NOT NULL,
  condition_item TEXT NOT NULL,
  FOREIGN KEY(condition_item) REFERENCES products(SKU)
);

-- Add data

INSERT INTO users
VALUES(1, '2020-09-28 13:51:20.7550326+02:00', '2020-09-28 13:51:20.9140391+02:00', 'Angelo', 'Perera', 'Paya', '$2a$04$8/dc2MobAHst255vM2ulTuibUhxLz2uCLqZrGNeVxQsAI.vJ/CGBa');

INSERT INTO products
  (SKU,Name,Price,Inventory_Qty)
VALUES
  ('120P90', 'Google Home', 49.99, 10 );
INSERT INTO products
  (SKU,Name,Price,Inventory_Qty)
VALUES
  ('43N23P', 'MacBook Pro', 5399.99, 5 );
INSERT INTO products
  (SKU,Name,Price,Inventory_Qty)
VALUES
  ('A304SD', 'Alexa Speaker', 109.5, 10 );
INSERT INTO products
  (SKU,Name,Price,Inventory_Qty)
VALUES
  ('234234', 'Raspberry Pi B', 30, 2 );

INSERT INTO discount_rules
  (condition_qty,condition_item) VALUES
  (   1,      '43N23P');