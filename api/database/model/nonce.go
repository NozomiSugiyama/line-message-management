package model

import (
	"github.com/jinzhu/gorm"
)

// NonceRepository repository model
type NonceRepository struct {
	db *gorm.DB
}

// Nonce data model
type Nonce struct {
	gorm.Model
	Text   string `sql:"type:text;"`
	User   User   `gorm:"foreignkey:UserID"`
	UserID uint
}

// NewNonceRepository create nonce repository
func NewNonceRepository(db *gorm.DB) *NonceRepository {
	repo := new(NonceRepository)
	repo.db = db
	return repo
}

// CreateNonce Create nonce
func (repo *NonceRepository) CreateNonce(nonce *Nonce) error {
	repo.db.Table(NonceTableName).Create(nonce)
	return nil
}

// ListNonces List nonces from store
func (repo *NonceRepository) ListNonces() ([]Nonce, error) {
	var nonces []Nonce
	repo.db.Find(&nonces)
	return nonces, nil
}
