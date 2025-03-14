package track

import (
	"github.com/dragonsecurity/ptaas-api/pkg/enum"

	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	ProjectID   uint
	DocumentID  uint
	Service     string
	Description string
	Type        enum.TrackType
}
