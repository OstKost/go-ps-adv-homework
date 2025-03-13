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

func (repository *StatRepository) AddClick(linkId string) {
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
