package stat

import (
	"go-ps-adv-homework/pkg/db"
	"gorm.io/datatypes"
	"time"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repository *StatRepository) AddClick(linkId uint) {
	var stat Stat
	currentDate := datatypes.Date(time.Now())
	repository.Db.Find(&stat, "link_id = ? and date = ?", linkId, currentDate)
	if stat.ID == 0 {
		repository.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks += 1
		repository.Db.Save(&stat)
	}
}

// GetStatsByPeriod
// ************** SQL Query **************
// select to_char(date, 'YYYY-MM-DD') as period, SUM(stats.clicks) AS "clicks"
// from stats
// where date BETWEEN '2021-07-01' AND '2025-03-15'
// group by period
// order by period;
// ************** SQL Query **************
func (repository *StatRepository) GetStatsByPeriod(startDate, endDate time.Time, period string) []ClicksByPeriod {
	var stats []ClicksByPeriod
	var selectQuery string
	switch period {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, SUM(stats.clicks) as clicks"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, SUM(stats.clicks) as clicks"
	}
	repository.Db.
		Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Group("period").
		Order("period").
		Find(&stats)

	return stats
}
