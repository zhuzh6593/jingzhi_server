package git

import (
	"errors"

	"jingzhi-server/builder/git/mirrorserver"
	"jingzhi-server/builder/git/mirrorserver/gitea"
	"jingzhi-server/common/config"
	"jingzhi-server/common/types"
)

func NewMirrorServer(config *config.Config) (mirrorserver.MirrorServer, error) {
	if !config.MirrorServer.Enable {
		return nil, nil
	}
	if config.MirrorServer.Type == types.GitServerTypeGitea {
		mirrorServer, err := gitea.NewMirrorClient(config)
		return mirrorServer, err
	}

	return nil, errors.New("undefined mirror server type")
}
