package common

import (
	"errors"
	"log/slog"
	"strconv"

	"caict.ac.cn/llm-server/common/types"
	"github.com/gin-gonic/gin"
)

func GetNamespaceAndNameFromContext(ctx *gin.Context) (namespace string, name string, err error) {
	namespace = ctx.Param("namespace")
	name = ctx.Param("name")
	if namespace == "" || name == "" {
		err = errors.New("Invalid namespace or name")
		return
	}
	return
}

func GetPerAndPageFromContext(ctx *gin.Context) (perInt int, pageInt int, err error) {
	per := ctx.Query("per")
	if per == "" {
		per = "50"
	}
	perInt, err = strconv.Atoi(per)
	if err != nil {
		return
	}
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}
	pageInt, err = strconv.Atoi(page)
	if err != nil {
		return
	}
	return
}

func RepoTypeFromContext(ctx *gin.Context) types.RepositoryType {
	rawRp, exist := ctx.Get("repo_type")
	slog.Debug("get repo type from context", "repo_type", rawRp, "exists", exist)
	if !exist {
		return types.UnknownRepo
	}
	return rawRp.(types.RepositoryType)
}

func SetRepoTypeContext(ctx *gin.Context, t types.RepositoryType) {
	ctx.Set("repo_type", t)
}
