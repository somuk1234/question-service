package managers

import (
	"questions-keeper-service/internal/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database is the object uses by the managers for accessing
// database tables and executing queries.
var Database *gorm.DB

func NewPostgresSQL() *gorm.DB {
	return Database
}
func SetPostgreSQL() {
	var err error
	Database, err = gorm.Open("postgres", "host=localhost port=5432 user=somu dbname=my_db password=9110SOMU@ sslmode=disable")

	if err != nil {
		panic(err)
	}
	if !Database.HasTable(&models.Question{}) {
		Database.CreateTable(&models.Question{})
	}
	// set this to 'true' to see sql logs
	Database.LogMode(true)
}
