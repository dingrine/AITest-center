package repository

import (
	"examCenter/internal/chaindata"
	"examCenter/internal/repository"
	"examCenter/internal/server/client"
)

type Repositories struct {
	RepoExamData repository.ExamDataRepository
	// RepoCacheExamInfo repository.TokenCache[*redis.NewExamInfo]
}

func New(client *client.Clients) *Repositories {
	var repo Repositories
	repo.RepoExamData = chaindata.NewRepo(client.DB)
	return &repo
}
