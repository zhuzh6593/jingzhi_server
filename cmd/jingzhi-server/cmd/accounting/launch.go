package accounting

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
	"jingzhi-server/accounting/consumer"
	"jingzhi-server/accounting/router"
	"jingzhi-server/api/httpbase"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
	"jingzhi-server/mq"
)

var launchCmd = &cobra.Command{
	Use:     "launch",
	Short:   "Launch accounting server",
	Example: serverExample(),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}
		slog.Debug("config", slog.Any("data", cfg))
		// Check APIToken length
		if len(cfg.APIToken) < 128 {
			return fmt.Errorf("API token length is less than 128, please check")
		}
		dbConfig := database.DBConfig{
			Dialect: database.DatabaseDialect(cfg.Database.Driver),
			DSN:     cfg.Database.DSN,
		}
		database.InitDB(dbConfig)

		mqHandler, err := mq.Init(cfg)
		if err != nil {
			return fmt.Errorf("fail to build message queue handler: %w", err)
		}

		// Do metering
		meter := consumer.NewMetering(mqHandler, cfg)
		meter.Run()

		r, err := router.NewAccountRouter(cfg)
		if err != nil {
			return fmt.Errorf("failed to init router: %w", err)
		}
		slog.Info("http server is running", slog.Any("port", cfg.Accounting.Port))
		server := httpbase.NewGracefulServer(
			httpbase.GraceServerOpt{
				Port: cfg.Accounting.Port,
			},
			r,
		)
		server.Run()

		return nil
	},
}

func serverExample() string {
	return `
# for development
jingzhi-server accounting launch
`
}
