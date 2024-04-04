CREATE TABLE tags (
  id bigint PRIMARY KEY,
  user_id bigint NOT NULL,
  name varchar(255) NOT NULL,
  description text DEFAULT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp DEFAULT NULL
);
