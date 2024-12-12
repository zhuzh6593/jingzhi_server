package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"jingzhi-server/builder/store/database"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return createTables(ctx, db, database.Code{})
	}, func(ctx context.Context, db *bun.DB) error {
		return dropTables(ctx, db, database.Code{})
	})
}
