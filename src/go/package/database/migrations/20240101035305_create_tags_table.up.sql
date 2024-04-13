CREATE TABLE tags (
  id bigint PRIMARY KEY,
  user_id bigint DEFAULT NULL,
  name varchar(255) NOT NULL,
  is_show int DEFAULT 1,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp DEFAULT NULL
);

COMMENT ON COLUMN tags.is_show IS '1: show, 2: not show';
