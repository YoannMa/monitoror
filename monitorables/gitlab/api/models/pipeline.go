package models

import (
	"time"

	coreModels "github.com/monitoror/monitoror/models"
)

type Pipeline struct {
	ID         int
	Url        string
	Branch     string
	Author     coreModels.Author
	Status     string
	StartedAt  *time.Time
	FinishedAt *time.Time
}
