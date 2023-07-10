package crud

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/tniita/ochacafe-github-actions/oke-app/backend-app/db"
	"github.com/tniita/ochacafe-github-actions/oke-app/backend-app/repo"
)

func GetItems() []repo.Items {
	dbInfo := db.GetDbInfo()
	items := []repo.Items{}
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return items
}
func GetItemById(id string) repo.Items {
	dbInfo := db.GetDbInfo()
	item := repo.Items{}
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.First(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return item
}

func UpdateItem(items repo.Items) int64 {
	dbInfo := db.GetDbInfo()
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.Save(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}

func DeleteItem(id string) int64 {
	dbInfo := db.GetDbInfo()
	item := repo.Items{}
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.Delete(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}
