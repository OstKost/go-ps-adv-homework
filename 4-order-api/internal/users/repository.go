package users

import (
	"go-ps-adv-homework/pkg/db"
	"gorm.io/gorm/clause"
	"time"
)

type UserRepository struct {
	database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (repository *UserRepository) GetOrCreateUser(phone string) (*User, error) {
	var user User
	result := repository.database.DB.
		Where(User{Phone: phone}).
		Assign(User{Name: phone, Birthdate: time.Now()}).
		FirstOrCreate(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repository *UserRepository) CreateUser(user *User) (*User, error) {
	result := repository.database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repository *UserRepository) UpdateUser(user *User) (*User, error) {
	result := repository.database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repository *UserRepository) DeleteUser(id uint) error {
	result := repository.database.DB.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepository) GetUserById(id uint) (*User, error) {
	var user User
	result := repository.database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repository *UserRepository) GetUserByPhone(email string) (*User, error) {
	var user User
	result := repository.database.DB.First(&user, "phone = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repository *UserRepository) FindUsers(name, phone string) (*[]User, error) {
	var users []User
	result := repository.database.DB.
		Where("deleted_at IS NULL").
		Where("name LIKE ? OR phone LIKE ?", "%"+name+"%", "%"+phone+"%").
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}
