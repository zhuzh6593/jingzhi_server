package middleware

import (
	"log/slog"

	"caict.ac.cn/llm-server/common/types"
	"caict.ac.cn/llm-server/common/utils/common"
	"github.com/gin-gonic/gin"
)

func RepoType(t types.RepositoryType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slog.Debug("middleware RepoType called", "repo_type", t)
		common.SetRepoTypeContext(ctx, t)
		ctx.Next()
	}
}
