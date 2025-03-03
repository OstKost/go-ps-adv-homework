package link

import (
	"go-ps-adv-homework/pkg/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (repository *LinkRepository) Create(link *Link) (*Link, error) {
	result := repository.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repository *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repository.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repository *LinkRepository) Update(link *Link) (*Link, error) {
	result := repository.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repository *LinkRepository) Delete(id uint) error {
	result := repository.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repository.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repository *LinkRepository) GetActiveList(url string) (*[]Link, error) {
	var links []Link
	result := repository.Database.DB.Where("deleted_at IS NULL").Where("url LIKE ?", "%"+url+"%").Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return &links, nil
}
