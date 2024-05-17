package deploy

import (
	"caict.ac.cn/llm-server/builder/deploy/imagerunner"
	"caict.ac.cn/llm-server/common/config"
	"github.com/spf13/cobra"
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

		s, err := imagerunner.NewHttpServer(config)
		if err != nil {
			return err
		}
		err = s.Run(8082)
		return err
	},
}
