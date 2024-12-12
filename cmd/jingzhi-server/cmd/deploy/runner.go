package deploy

import (
	"github.com/spf13/cobra"
	"jingzhi-server/api/httpbase"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
	"jingzhi-server/servicerunner/router"
)

var startRunnerCmd = &cobra.Command{
	Use:   "runner",
	Short: "start space runner service",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		return
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.LoadConfig()
		if err != nil {
			return err
		}

		dbConfig := database.DBConfig{
			Dialect: database.DatabaseDialect(config.Database.Driver),
			DSN:     config.Database.DSN,
		}
		database.InitDB(dbConfig)

		s, err := router.NewHttpServer(config)
		if err != nil {
			return err
		}
		server := httpbase.NewGracefulServer(
			httpbase.GraceServerOpt{
				Port: config.Space.RunnerServerPort,
			},
			s,
		)
		server.Run()
		return nil
	},
}
