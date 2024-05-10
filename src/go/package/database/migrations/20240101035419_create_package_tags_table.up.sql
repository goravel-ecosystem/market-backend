CREATE TABLE package_tags (
  id bigint AUTO_INCREMENT PRIMARY KEY,
  package_id bigint NOT NULL,
  tag_id bigint NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE(package_id, tag_id)
);
