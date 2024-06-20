CREATE TABLE package_tags (
  id SERIAL PRIMARY KEY NOT NULL,
  package_id bigint NOT NULL,
  tag_id bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(package_id, tag_id)
);
