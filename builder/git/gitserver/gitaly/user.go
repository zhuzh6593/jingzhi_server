package gitaly

import (
	"context"

	"jingzhi-server/builder/git/gitserver"
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/types"
)

func (c *Client) CreateUser(u gitserver.CreateUserRequest) (user *gitserver.CreateUserResponse, err error) {
	return &gitserver.CreateUserResponse{
		GitID:    0,
		Password: "",
	}, nil
}

func (c *Client) UpdateUser(u *types.UpdateUserRequest, user *database.User) (*database.User, error) {
	return nil, nil
}

func (c *Client) UpdateUserV2(u gitserver.UpdateUserRequest) error {
	return nil
}

// Create gitea orgs for user to store different type repositories
func (c *Client) FixUserData(ctx context.Context, userName string) (err error) {
	return
}
