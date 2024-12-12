package git

import (
	"errors"

	"jingzhi-server/builder/git/gitserver"
	"jingzhi-server/builder/git/gitserver/gitaly"
	"jingzhi-server/builder/git/gitserver/gitea"
	"jingzhi-server/common/config"
	"jingzhi-server/common/types"
)

func NewGitServer(config *config.Config) (gitserver.GitServer, error) {
	if config.GitServer.Type == types.GitServerTypeGitea {
		gitServer, err := gitea.NewClient(config)
		return gitServer, err
	} else if config.GitServer.Type == types.GitServerTypeGitaly {
		gitServer, err := gitaly.NewClient(config)
		return gitServer, err
	}

	return nil, errors.New("undefined git server type")
}
