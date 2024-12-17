SET statement_timeout = 0;

--bun:split

ALTER TABLE repository_external_sources
DROP CONSTRAINT IF EXISTS fk_repository_external_sources_repository;
