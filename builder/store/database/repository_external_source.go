package database

import (
	"context"
	"fmt"
	"time"
)

type RepositoryExternalSourceStore struct {
	db *DB
}

func NewRepositoryExternalSourceStore() *RepositoryExternalSourceStore {
	return &RepositoryExternalSourceStore{
		db: defaultDB,
	}
}

func (s *RepositoryExternalSourceStore) Create(ctx context.Context, source RepositoryExternalSource) (*RepositoryExternalSource, error) {
	err := s.db.RunInTx(ctx, func(ctx context.Context, tx Operator) error {
		source.CreatedAt = time.Now()
		source.UpdatedAt = time.Now()
		_, err := tx.Core.NewInsert().Model(&source).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to create external source: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (s *RepositoryExternalSourceStore) CreateBatch(ctx context.Context, sources []RepositoryExternalSource) error {
	if len(sources) == 0 {
		return nil
	}

	return s.db.RunInTx(ctx, func(ctx context.Context, tx Operator) error {
		now := time.Now()
		for i := range sources {
			sources[i].CreatedAt = now
			sources[i].UpdatedAt = now
		}

		_, err := tx.Core.NewInsert().Model(&sources).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to create external sources in batch: %w", err)
		}
		return nil
	})
}

func (s *RepositoryExternalSourceStore) Update(ctx context.Context, source RepositoryExternalSource) error {
	return s.db.RunInTx(ctx, func(ctx context.Context, tx Operator) error {
		source.UpdatedAt = time.Now()
		result, err := tx.Core.NewUpdate().Model(&source).WherePK().Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to update external source: %w", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("failed to get rows affected: %w", err)
		}
		if rowsAffected == 0 {
			return fmt.Errorf("external source not found")
		}
		return nil
	})
}

func (s *RepositoryExternalSourceStore) Delete(ctx context.Context, id int64) error {
	return s.db.RunInTx(ctx, func(ctx context.Context, tx Operator) error {
		result, err := tx.Core.NewDelete().Model((*RepositoryExternalSource)(nil)).Where("id = ?", id).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to delete external source: %w", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("failed to get rows affected: %w", err)
		}
		if rowsAffected == 0 {
			return fmt.Errorf("external source not found")
		}
		return nil
	})
}

func (s *RepositoryExternalSourceStore) DeleteByRepositoryID(ctx context.Context, repositoryID int64) error {
	return s.db.RunInTx(ctx, func(ctx context.Context, tx Operator) error {
		_, err := tx.Core.NewDelete().Model((*RepositoryExternalSource)(nil)).Where("repository_id = ?", repositoryID).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to delete external sources: %w", err)
		}
		return nil
	})
}

func (s *RepositoryExternalSourceStore) FindByRepositoryID(ctx context.Context, repositoryID int64) ([]RepositoryExternalSource, error) {
	var sources []RepositoryExternalSource
	err := s.db.Core.NewSelect().
		Model(&sources).
		Relation("Repository").
		Where("repository_id = ?", repositoryID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find external sources: %w", err)
	}
	return sources, nil
}

func (s *RepositoryExternalSourceStore) FindByID(ctx context.Context, id int64) (*RepositoryExternalSource, error) {
	source := new(RepositoryExternalSource)
	err := s.db.Core.NewSelect().
		Model(source).
		Relation("Repository").
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find external source: %w", err)
	}
	return source, nil
}
