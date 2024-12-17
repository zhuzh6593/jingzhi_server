SET statement_timeout = 0;

--bun:split

ALTER TABLE repository_external_sources
ADD CONSTRAINT fk_repository_external_sources_repository
FOREIGN KEY (repository_id)
REFERENCES repositories(id)
ON DELETE CASCADE;
