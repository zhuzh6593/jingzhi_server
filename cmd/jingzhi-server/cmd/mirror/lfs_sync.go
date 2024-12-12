package mirror

import (
	"github.com/spf13/cobra"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
	"jingzhi-server/mirror"
)

var lfsSyncCmd = &cobra.Command{
	Use:     "lfs-sync",
	Short:   "Start the repoisotry lfs files sync server",
	Example: lfsSyncExample(),
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

		lfsSyncWorker, err := mirror.NewLFSSyncWorker(cfg, cfg.Mirror.WorkerNumber)
		if err != nil {
			return err
		}
		lfsSyncWorker.Run()

		return nil
	},
}

func lfsSyncExample() string {
	return `
# for development
jingzhi-server mirror lfs-sync
`
}
