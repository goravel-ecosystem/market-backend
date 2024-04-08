CREATE TABLE tags (
  id bigint PRIMARY KEY,
  user_id bigint DEFAULT NULL,
  name varchar(255) NOT NULL,
  description text DEFAULT NULL,
  /* 1: show, 2: not show */
  is_show int DEFAULT 1,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp DEFAULT NULL
);
