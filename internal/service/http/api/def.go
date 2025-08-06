package api

import (
	"examCenter/internal/core"
	"examCenter/internal/repository"

	"github.com/CrazyThursdayV50/pkgo/log"
)

type Service struct {
	logger   log.Logger
	RepoData repository.ExamDataRepository
	triggers *core.Triggers
}
