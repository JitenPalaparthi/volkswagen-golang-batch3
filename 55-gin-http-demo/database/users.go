package database

import (
	"demo/models"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) IUserDB {
	return &UserDB{db}
}

type IUserDB interface {
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
	GetBy(id string) (*models.User, error)
	DeleteBy(id string) error
}

func (u *UserDB) Create(user *models.User) (*models.User, error) {
	err := u.DB.AutoMigrate(&models.User{}) // tables are created or altered
	if err != nil {
		return nil, err
	}

	tx := u.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) Update(user *models.User) (*models.User, error) {
	tx := u.DB.Save(user) // unsert
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
func (u *UserDB) GetBy(id string) (*models.User, error) {
	user := new(models.User)
	tx := u.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) DeleteBy(id string) error {
	tx := u.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// https://www.youtube.com/playlist?list=PLJE7PIP1qj_Rn9vq4V4jGJbj5KqEIWSUc
