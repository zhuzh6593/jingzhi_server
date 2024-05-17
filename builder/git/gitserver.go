package git

import (
	"errors"

	"caict.ac.cn/llm-server/builder/git/gitserver"
	"caict.ac.cn/llm-server/builder/git/gitserver/gitea"
	"caict.ac.cn/llm-server/common/config"
)

func NewGitServer(config *config.Config) (gitserver.GitServer, error) {
	if config.GitServer.Type == "gitea" {
		gitServer, err := gitea.NewClient(config)
		return gitServer, err
	}

	return nil, errors.New("undefined git server type")
}
