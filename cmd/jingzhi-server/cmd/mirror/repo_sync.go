package mirror

import (
	"github.com/spf13/cobra"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
	"jingzhi-server/mirror"
)

var repoSyncCmd = &cobra.Command{
	Use:     "repo-sync",
	Short:   "Start the repoisotry sync server",
	Example: repoSyncExample(),
	RunE: func(*cobra.Command, []string) (err error) {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		dbConfig := database.DBConfig{
			Dialect: database.DatabaseDialect(cfg.Database.Driver),
			DSN:     cfg.Database.DSN,
		}
		database.InitDB(dbConfig)

		repoSYncer, err := mirror.NewRepoSyncWorker(cfg, cfg.Mirror.WorkerNumber)
		if err != nil {
			return err
		}

		repoSYncer.Run()

		return nil
	},
}

func repoSyncExample() string {
	return `
# for development
jingzhi-server mirror repo-sync
`
}
