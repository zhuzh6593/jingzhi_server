package common

import (
	"testing"
	"time"

	"jingzhi-server/builder/store/database"
	"jingzhi-server/common/types"
)

func TestAddPrefixBySourceID(t *testing.T) {
	s := &database.SyncVersion{
		Version:        1,
		SourceID:       0,
		RepoPath:       "test/test",
		RepoType:       types.ModelRepo,
		LastModifiedAt: time.Now(),
		ChangeLog:      "test log",
	}
	str := AddPrefixBySourceID(s.SourceID, "test")
	if str != "Jingzhi_test" {
		t.Errorf("Expected str should be 'Jingzhi_test' but got %s", str)
	}

	s1 := &database.SyncVersion{
		Version:        1,
		SourceID:       1,
		RepoPath:       "test/test",
		RepoType:       types.ModelRepo,
		LastModifiedAt: time.Now(),
		ChangeLog:      "test log",
	}
	str1 := AddPrefixBySourceID(s1.SourceID, "test")
	if str1 != "Huggingface_test" {
		t.Errorf("Expected str should be 'Huggingface_test' but got %s", str)
	}
}

func TestTrimPrefixCloneURLBySourceID(t *testing.T) {
	s := &database.SyncVersion{
		Version:        1,
		SourceID:       0,
		RepoPath:       "test/test",
		RepoType:       types.ModelRepo,
		LastModifiedAt: time.Now(),
		ChangeLog:      "test log",
	}
	cloneURL := TrimPrefixCloneURLBySourceID(
		"https://jingzhi.com",
		"model",
		"Jingzhi_namespace",
		"name",
		s.SourceID,
	)
	if cloneURL != "https://jingzhi.com/models/namespace/name.git" {
		t.Errorf("Expected cloneURL should be 'https://jingzhi.com/models/namespace/name' but got %s", cloneURL)
	}

	s1 := &database.SyncVersion{
		Version:        1,
		SourceID:       1,
		RepoPath:       "test/test",
		RepoType:       types.ModelRepo,
		LastModifiedAt: time.Now(),
		ChangeLog:      "test log",
	}
	cloneURL1 := TrimPrefixCloneURLBySourceID(
		"https://jingzhi.com",
		"model",
		"Huggingface_namespace",
		"name",
		s1.SourceID,
	)
	if cloneURL1 != "https://jingzhi.com/models/namespace/name.git" {
		t.Errorf("Expected cloneURL should be 'https://jingzhi.com/models/namespace/name' but got %s", cloneURL1)
	}
}
