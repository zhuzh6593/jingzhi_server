package mirror

import (
	"context"

	"jingzhi-server/common/config"
	"jingzhi-server/mirror/queue"
	"jingzhi-server/mirror/reposyncer"
)

type RepoSyncWorker interface {
	Run()
	SyncRepo(ctx context.Context, task queue.MirrorTask) error
}

func NewRepoSyncWorker(config *config.Config, numWorkers int) (RepoSyncWorker, error) {
	return reposyncer.NewLocalMirrorWoker(config, numWorkers)
}
