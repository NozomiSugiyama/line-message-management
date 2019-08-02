package model

import (
	"github.com/jinzhu/gorm"
)

// LineUsersRepository repository model
type LineUsersRepository struct {
	db *gorm.DB
}

// LineUser data model
type LineUser struct {
	ID            int    `gorm:"column:id;type:integer;primary_key;not null"`
	User          User   `gorm:"foreignkey:UserID"`
	UserID        int    `gorm:"column:user_id;type:integer;not null"`
	LineID        string `gorm:"column:line_id;type:text;not null"`
	LinkedAccount string `gorm:"column:linked_account;type:text;not null"`
}

// NewLineUsersRepository create nonce repository
func NewLineUsersRepository(db *gorm.DB) *LineUsersRepository {
	repo := new(LineUsersRepository)
	repo.db = db
	return repo
}

// CreateLineUser Create nonce
func (repo *LineUsersRepository) CreateLineUser(nonce *LineUser) error {
	err := repo.db.Table(LineUserTableName).Create(nonce).Error
	return err
}

// ListLineUsers List nonces from store
func (repo *LineUsersRepository) ListLineUsers() ([]LineUser, error) {
	var nonces []LineUser
	err := repo.db.Table(LineUserTableName).Find(&nonces).Error
	return nonces, err
}
