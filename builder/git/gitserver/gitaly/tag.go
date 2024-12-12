package gitaly

import (
	"context"

	"jingzhi-server/builder/git/gitserver"
	"jingzhi-server/common/types"
)

func (c *Client) GetRepoTags(ctx context.Context, req gitserver.GetRepoTagsReq) (tags []*types.Tag, err error) {
	return
}
