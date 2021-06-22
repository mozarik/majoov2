package repository

import (
	model "github.com/mozarik/majoov2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) ReturnCurrentUser(username string) (*model.User, error) {
	user := &model.User{}
	// err := u.db.Where("username = ?", username).First(&user).Error
	err := u.db.Preload(clause.Associations).Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func (u *userRepository) ReturnAllUser() (*model.User, error) {
	err := u.db.Find(&model.User{}).Error
	if err != nil {
		return nil, err
	}
	return &model.User{}, err
}

func (u *userRepository) Register(user *model.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) DeleteById(id uint) error {
	var user model.User
	return u.db.Delete(&user, id).Error
}

func (u *userRepository) Save(user *model.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) FindById(id uint) error {
	var user model.User

	return u.db.First(&user, id).Error
}

func (u *userRepository) UsernameIsInDb(username string) (bool, error) {
	var user model.User
	result := u.db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return true, result.Error
	}

	return false, result.Error
}

func (u *userRepository) GetPassword(username string) (string, error) {
	var user model.User
	result := u.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Password, nil
}

func (u *userRepository) GetIDByUsername(username string) (uint, error) {
	var id uint
	err := u.db.Raw("SELECT id from users WHERE username = ?", username).Find(&id)
	if err.Error != nil {
		return 0, err.Error
	}
	return id, nil
}
