package model

import (
	"github.com/jinzhu/gorm"
)

// LineUserRepository repository model
type LineUserRepository struct {
	db *gorm.DB
}

// LineUser data model
type LineUser struct {
	ID            int    `gorm:"column:id;type:integer;primary_key;not null"`
	User          User   `gorm:"foreignkey:ID;association_foreignkey:UserID"`
	UserID        int    `gorm:"column:user_id;type:integer;not null"`
	LineID        string `gorm:"column:line_id;type:text;not null"`
	LinkedAccount string `gorm:"column:linked_account;type:text;not null"`
}

// NewLineUserRepository create nonce repository
func NewLineUserRepository(db *gorm.DB) *LineUserRepository {
	repo := new(LineUserRepository)
	repo.db = db
	return repo
}

// CreateLineUser Create nonce
func (repo *LineUserRepository) CreateLineUser(nonce *LineUser) error {
	err := repo.db.Table(LineUserTableName).Create(nonce).Error
	return err
}

// ListLineUsers List users from store
func (repo *LineUserRepository) ListLineUsers() ([]LineUser, error) {
	var users []LineUser
	err := repo.db.Table(LineUserTableName).Find(&users).Error
	return users, err
}

// ListLineUsersWithUser List users from store
func (repo *LineUserRepository) ListLineUsersWithUser() ([]LineUser, error) {
	var users []LineUser
	err := repo.db.Preload("User").Table(LineUserTableName).Find(&users).Error
	return users, err
}

// FindLineUserByID Find user by user id
func (repo *LineUserRepository) FindLineUserByID(id int) (LineUser, error) {
	var lineUser LineUser
	if err := repo.db.Table(LineUserTableName).First(&lineUser, id).Error; gorm.IsRecordNotFoundError(err) {
		return lineUser, ErrRecordNotFound
	}
	return lineUser, nil
}

// FindLineUserWithUserByID Find user by user id
func (repo *LineUserRepository) FindLineUserWithUserByID(id int) (LineUser, error) {
	var lineUser LineUser
	if err := repo.db.Preload("User").Table(LineUserTableName).First(&lineUser, id).Error; gorm.IsRecordNotFoundError(err) {
		return lineUser, ErrRecordNotFound
	}
	return lineUser, nil
}
