package start

import (
	"fmt"

	"github.com/spf13/cobra"
	"jingzhi-server/api/httpbase"
	"jingzhi-server/api/router"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
)

var rproxyCmd = &cobra.Command{
	Use:     "rproxy",
	Short:   "Start the reverse proxy server",
	Example: rproxyExample(),
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
		r, err := router.NewRProxyRouter(cfg)
		if err != nil {
			return fmt.Errorf("failed to init router: %w", err)
		}
		server := httpbase.NewGracefulServer(
			httpbase.GraceServerOpt{
				Port: cfg.Space.RProxyServerPort,
			},
			r,
		)
		server.Run()

		return nil
	},
}

func rproxyExample() string {
	return `
# for development
jingzhi-server start rproxy
`
}
