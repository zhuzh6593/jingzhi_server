package migrations

import (
	"context"

	"caict.ac.cn/llm-server/builder/store/database"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		dropTables(ctx, db, database.Space{})
		return createTables(ctx, db, database.Space{}, database.Deploy{}, database.DeployTask{})
	}, func(ctx context.Context, db *bun.DB) error {
		return dropTables(ctx, db, database.Deploy{}, database.DeployTask{})
	})
}
