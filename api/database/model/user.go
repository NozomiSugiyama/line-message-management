package model

import (
	"github.com/jinzhu/gorm"
)

// UserRepository repository model
type UserRepository struct {
	db *gorm.DB
}

// User data model
type User struct {
	ID       int    `gorm:"column:id;type:integer;primary_key;not null"`
	Name     string `gorm:"column:name;type:text;not null"`
	Email    string `gorm:"column:email;type:text;not null;unique"`
	Password string `gorm:"column:password;type:text;not null"`
}

// NewUserRepository create user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	repo := new(UserRepository)
	repo.db = db
	return repo
}

// CreateUser Create user
func (repo *UserRepository) CreateUser(user *User) error {
	repo.db.Table(UserTableName).Create(user)
	return nil
}

// ListUsers List users from store
func (repo *UserRepository) ListUsers() ([]User, error) {
	var users []User
	err := repo.db.Table(UserTableName).Find(&users).Error
	return users, err
}

// FindUserByID Find user by user id
func (repo *UserRepository) FindUserByID(id int) (User, error) {
	var user User
	if err := repo.db.Table(UserTableName).First(&user, id).Error; gorm.IsRecordNotFoundError(err) {
		return user, ErrRecordNotFound
	}
	return user, nil
}

// FindUserByEMail Find user by email
func (repo *UserRepository) FindUserByEMail(email string) (User, error) {
	var user User
	if err := repo.db.Table(UserTableName).Where("email = ?", email).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return user, ErrRecordNotFound
	}
	return user, nil
}
