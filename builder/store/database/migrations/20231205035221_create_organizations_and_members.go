package migrations

import (
	"context"

	"caict.ac.cn/llm-server/builder/store/database"
	"github.com/uptrace/bun"
)

var orgTables = []any{
	database.Organization{},
	database.Member{},
}

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return createTables(ctx, db, orgTables...)
	}, func(ctx context.Context, db *bun.DB) error {
		return dropTables(ctx, db, orgTables...)
	})
}
