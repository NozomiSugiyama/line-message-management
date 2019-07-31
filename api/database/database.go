package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Initialize initializes the database
func Initialize(host string, port string, user string, dbname string, password string, sslMode string) (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslMode)

	fmt.Println("======= Database config ========")
	fmt.Println(config)
	fmt.Println("================================")

	db, err := gorm.Open("postgres", config)

	db.LogMode(true)
	return db, err
}
