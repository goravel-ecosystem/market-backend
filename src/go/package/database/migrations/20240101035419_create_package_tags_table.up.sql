CREATE TABLE package_tags (
  id bigint PRIMARY KEY,
  package_id bigint NOT NULL,
  tag_id bigint NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);
