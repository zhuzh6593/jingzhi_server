package git

import (
	"jingzhi-server/builder/git/membership"
	"jingzhi-server/builder/git/membership/gitea"
	"jingzhi-server/common/config"
)

func NewMemberShip(config config.Config) (membership.GitMemerShip, error) {
	c, err := gitea.NewClient(&config)
	return c, err
}
