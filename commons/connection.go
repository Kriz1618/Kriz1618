package commons

import (
	"log"

	"github.com/Kriz1618/crud-api-rest-go/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3360)/test")
	// Open("mysql", "root:welcome@tcp(127.0.0.1:3306)/test")
	if error != nil {
		log.Fatal(error)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrating db schema...")
	db.AutoMigrate(&models.Persona{})
}
