package models

import (
	"github.com/dragonsecurity/ptaas-api/pkg/models/document"
	"github.com/dragonsecurity/ptaas-api/pkg/models/project"
	"github.com/dragonsecurity/ptaas-api/pkg/models/track"
	"github.com/dragonsecurity/ptaas-api/pkg/models/user"

	"gorm.io/gorm"
)

// Interface manages the models interfaces
type Interface struct {
	Documents document.Interface
	Projects  project.Interface
	Users     user.Interface
	Tracks    track.Interface
}

func New(db *gorm.DB) *Interface {
	return &Interface{
		Documents: document.New(db),
		Projects:  project.New(db),
		Users:     user.New(db),
		Tracks:    track.New(db),
	}
}
