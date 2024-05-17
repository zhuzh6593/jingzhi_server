package git

import (
	"errors"

	"caict.ac.cn/llm-server/builder/git/membership"
	"caict.ac.cn/llm-server/builder/git/membership/gitea"
	"caict.ac.cn/llm-server/common/config"
)

func NewMemberShip(config config.Config) (membership.GitMemerShip, error) {
	if config.GitServer.Type == "gitea" {
		c, err := gitea.NewClient(&config)
		return c, err
	}
	return nil, errors.New("undefined git server type")
}
