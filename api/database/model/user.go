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
	ID   string `json:"id"`
	Name string `json:"name"`
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
	repo.db.Find(&users)
	return users, nil
}

// FindUserByID Find user by user id
func (repo *UserRepository) FindUserByID(id string) (User, error) {
	var user User
	repo.db.First(&user, id)
	return user, nil
}
