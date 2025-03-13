package link

import (
	"go-ps-adv-homework/pkg/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		database: db,
	}
}

func (repository *LinkRepository) Create(link *Link) (*Link, error) {
	result := repository.database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repository *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repository.database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repository *LinkRepository) Update(link *Link) (*Link, error) {
	result := repository.database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repository *LinkRepository) Delete(id uint) error {
	result := repository.database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repository.database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repository *LinkRepository) GetActiveList(url string, limit, offset int) ([]Link, error) {
	var links []Link
	result := repository.database.DB.
		Table("links").
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Where("url LIKE ?", "%"+url+"%").
		Limit(limit).
		Offset(offset).
		Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}

func (repository *LinkRepository) Count(url string, limit, offset int) (int64, error) {
	var count int64
	result := repository.database.DB.
		Table("links").
		Where("deleted_at IS NULL").
		Where("url LIKE ?", "%"+url+"%").
		Limit(limit).
		Offset(offset).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
