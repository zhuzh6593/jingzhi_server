package migrations

import (
	"context"
	"time"

	"jingzhi-server/builder/store/database"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return createTables(ctx, db, RepositoryExternalSource{})

	}, func(ctx context.Context, db *bun.DB) error {
		return dropTables(ctx, db, RepositoryExternalSource{})
	})
}

type RepositoryExternalSource struct {
	ID           int64                `bun:",pk,autoincrement"`
	RepositoryID int64                `bun:",notnull"`
	Repository   *database.Repository `bun:"rel:belongs-to,join:repository_id=id"`
	SourceName   string               `bun:",notnull"`
	SourceURL    string               `bun:",notnull"`
	CreatedAt    time.Time            `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time            `bun:",nullzero,notnull,default:current_timestamp"`
}
