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
	User          User   `gorm:"foreignkey:UserID"`
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

// ListLineUsers List nonces from store
func (repo *LineUserRepository) ListLineUsers() ([]LineUser, error) {
	var nonces []LineUser
	err := repo.db.Table(LineUserTableName).Find(&nonces).Error
	return nonces, err
}
