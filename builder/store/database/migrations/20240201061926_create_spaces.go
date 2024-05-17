package migrations

import (
	"context"

	"caict.ac.cn/llm-server/builder/store/database"
	"github.com/uptrace/bun"
)

type Space struct {
	ID           int64                `bun:",pk,autoincrement" json:"id"`
	RepositoryID int64                `bun:",notnull" json:"repository_id"`
	Repository   *database.Repository `bun:"rel:belongs-to,join:repository_id=id" json:"repository"`
	// gradio, streamlit, docker etc
	Sdk string `bun:",notnull" json:"sdk"`
	times
}

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return createTables(ctx, db, Space{})
	}, func(ctx context.Context, db *bun.DB) error {
		return dropTables(ctx, db, Space{})
	})
}
