CREATE TABLE users (
  id bigint PRIMARY KEY,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  avatar varchar(255) DEFAULT NULL,
  summary varchar(255) DEFAULT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp DEFAULT NULL
);

CREATE UNIQUE INDEX idx_unique_email ON users(email);