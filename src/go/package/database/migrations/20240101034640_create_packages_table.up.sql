CREATE TABLE packages (
  id bigint PRIMARY KEY,
  user_id bigint NOT NULL,
  name varchar(255) NOT NULL,
  summary varchar(255) DEFAULT NULL,
  description text DEFAULT NULL,
  link varchar(255) DEFAULT NULL,
  cover varchar(255) DEFAULT NULL,
  version varchar(255) DEFAULT NULL,
  is_public int DEFAULT 2,
  is_approved int DEFAULT 2,
  view_count bigint DEFAULT 0,
  last_updated_at timestamp NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp DEFAULT NULL
);
