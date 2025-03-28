package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	LinkId uint           `json:"link_id"`
	Clicks int            `json:"clicks"`
	Data   datatypes.Date `json:"data"`
}
