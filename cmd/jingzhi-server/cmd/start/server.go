package start

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"jingzhi-server/api/httpbase"
	"jingzhi-server/api/router"
	"jingzhi-server/builder/deploy"
	"jingzhi-server/builder/event"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
	"jingzhi-server/common/types"
	"jingzhi-server/docs"
	"jingzhi-server/mirror"
)

var enableSwagger bool

func init() {
	serverCmd.Flags().BoolVar(&enableSwagger, "swagger", false, "Start swagger help docs")
}

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start the API server",
	Example: serverExample(),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		enableSwagger = enableSwagger || cfg.EnableSwagger

		if enableSwagger {
			//	@securityDefinitions.apikey ApiKey
			//	@in                         header
			//	@name                       Authorization
			//	@description                Bearer token
			docs.SwaggerInfo.Title = "Jingzhi Server API"
			docs.SwaggerInfo.Description = "Jingzhi Server API."
			docs.SwaggerInfo.Version = "1.0"
			docs.SwaggerInfo.Host = cfg.APIServer.PublicDomain
			docs.SwaggerInfo.BasePath = "/api/v1"
			docs.SwaggerInfo.Schemes = []string{"http", "https"}
		}

		// Check APIToken length
		if len(cfg.APIToken) < 128 {
			return fmt.Errorf("API token length is less than 128, please check")
		}
		dbConfig := database.DBConfig{
			Dialect: database.DatabaseDialect(cfg.Database.Driver),
			DSN:     cfg.Database.DSN,
		}
		database.InitDB(dbConfig)
		err = event.InitEventPublisher(cfg)
		if err != nil {
			return fmt.Errorf("fail to initialize message queue, %w", err)
		}
		deploy.Init(deploy.DeployConfig{
			ImageBuilderURL:         cfg.Space.BuilderEndpoint,
			ImageRunnerURL:          cfg.Space.RunnerEndpoint,
			MonitorInterval:         10 * time.Second,
			InternalRootDomain:      cfg.Space.InternalRootDomain,
			SpaceDeployTimeoutInMin: cfg.Space.DeployTimeoutInMin,
			ModelDeployTimeoutInMin: cfg.Model.DeployTimeoutInMin,
			ModelDownloadEndpoint:   cfg.Model.DownloadEndpoint,
			PublicRootDomain:        cfg.Space.PublicRootDomain,
		})
		r, err := router.NewRouter(cfg, enableSwagger)
		if err != nil {
			return fmt.Errorf("failed to init router: %w", err)
		}
		server := httpbase.NewGracefulServer(
			httpbase.GraceServerOpt{
				Port: cfg.APIServer.Port,
			},
			r,
		)

		// Initialize mirror service
		mirrorService, err := mirror.NewMirrorPriorityQueue(cfg)
		if err != nil {
			return fmt.Errorf("failed to init mirror service: %w", err)
		}

		if cfg.MirrorServer.Enable && cfg.GitServer.Type == types.GitServerTypeGitaly {
			mirrorService.EnqueueMirrorTasks()
		}

		server.Run()

		return nil
	},
}

func serverExample() string {
	return `
# for development
jingzhi-server start server
`
}
