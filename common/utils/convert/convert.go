package convert

import (
	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/types"
)

// ToExternalSources converts database.RepositoryExternalSource slice to types.ExternalSource slice
func ToExternalSources(sources []database.RepositoryExternalSource) []types.ExternalSource {
	result := make([]types.ExternalSource, len(sources))
	for i, source := range sources {
		result[i] = types.ExternalSource{
			SourceName: source.SourceName,
			SourceURL:  source.SourceURL,
		}
	}
	return result
}
