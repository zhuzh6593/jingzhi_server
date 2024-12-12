package mirror

import (
	"context"

	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/config"
	"jingzhi-server/mirror/lfssyncer"
)

type LFSSyncWorker interface {
	Run()
	SyncLfs(ctx context.Context, workerID int, mirror *database.Mirror) error
}

func NewLFSSyncWorker(config *config.Config, numWorkers int) (LFSSyncWorker, error) {
	return lfssyncer.NewMinioLFSSyncWorker(config, numWorkers)
}
