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
	ID            int    `gorm:"column:id;type:integer;primary_key;not null"`
	User          User   `gorm:"foreignkey:UserID"`
	UserID        int    `gorm:"column:user_id;type:integer;not null"`
	Nonce         string `gorm:"column:nonce;type:text;not null"`
	LinkedAccount string `gorm:"column:linked_account;type:text;not null"`
}

// NewNonceRepository create nonce repository
func NewNonceRepository(db *gorm.DB) *NonceRepository {
	repo := new(NonceRepository)
	repo.db = db
	return repo
}

// CreateNonce Create nonce
func (repo *NonceRepository) CreateNonce(nonce *Nonce) error {
	err := repo.db.Table(NonceTableName).Create(nonce).Error
	return err
}

// FindNonceByNonce Find nonce by nonce
func (repo *NonceRepository) FindNonceByNonce(nonceValue string) (Nonce, error) {
	var nonce Nonce
	if err := repo.db.Table(NonceTableName).First(&nonce, nonceValue).Error; gorm.IsRecordNotFoundError(err) {
		return nonce, ErrRecordNotFound
	}
	return nonce, nil
}

// ListNonces List nonces from store
func (repo *NonceRepository) ListNonces() ([]Nonce, error) {
	var nonces []Nonce
	err := repo.db.Table(NonceTableName).Find(&nonces).Error
	return nonces, err
}
