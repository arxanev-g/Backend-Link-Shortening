package stat

import (
	"time"

	"github.com/arxanev/adv/pkg/db"
	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentData := datatypes.Date(time.Now())
	repo.Db.Find(&stat, "link_id = ? and data = ?", linkId, datatypes.Date(time.Now()))
	if stat.ID == 0 {
		repo.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Data:   currentData,
		})
	} else {
		stat.Clicks += 1
		repo.Db.Save(&stat)
	}
}

func (repo *StatRepository) GetSTats(by string, from, to time.Time) []GetStatResponse {
	var stats []GetStatResponse
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMounth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	repo.DB.Table("stats").Select(selectQuery).Where("data BETWEEN ? AND ?", from, to).Group("period").Order("period").Scan(&stats)
	return stats
}
