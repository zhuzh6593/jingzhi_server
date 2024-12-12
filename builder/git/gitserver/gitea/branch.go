package gitea

import (
	"context"

	"github.com/OpenCSGs/gitea-go-sdk/gitea"
	"jingzhi-server/builder/git/gitserver"
	"jingzhi-server/common/types"
	"jingzhi-server/common/utils/common"
)

func (c *Client) GetRepoBranches(ctx context.Context, req gitserver.GetBranchesReq) ([]types.Branch, error) {
	var branches []types.Branch
	namespace := common.WithPrefix(req.Namespace, repoPrefixByType(req.RepoType))
	giteaBranches, _, err := c.giteaClient.ListRepoBranches(
		namespace,
		req.Name,
		gitea.ListRepoBranchesOptions{
			ListOptions: gitea.ListOptions{
				PageSize: req.Per,
				Page:     req.Page,
			},
		},
	)
	for _, giteaBranch := range giteaBranches {
		branches = append(branches, types.Branch{
			Name:    giteaBranch.Name,
			Message: giteaBranch.Commit.Message,
			Commit: types.RepoBranchCommit{
				ID: giteaBranch.Commit.ID,
			},
		})
	}
	return branches, err
}
