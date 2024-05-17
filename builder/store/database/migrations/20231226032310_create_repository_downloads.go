package migrations

import (
	"context"

	"caict.ac.cn/llm-server/builder/store/database"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return createTables(ctx, db, database.RepositoryDownload{})
	}, func(ctx context.Context, db *bun.DB) error {
		return dropTables(ctx, db, database.RepositoryDownload{})
	})
}
