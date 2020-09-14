package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/student-management-service/models"
	"github.com/student-management-service/util"
)

func UserOnboardRepository(any interface{}) error {
	db, err := gorm.Open("mysql", util.DbConfig)

	if err != nil {
		fmt.Print(err)
		// panic function to cause a run time error
		panic("failed to connect database")
	}
	// defer which schedules a function call to be run after the function completes
	defer db.Close()

	// Migrate the schema
	keyOnboard := models.Key{}
	db.AutoMigrate(keyOnboard)

	// Create

	db.Create(any)
	fmt.Println("repository")
	fmt.Print(any)

	return nil

}

func RetriveKeysFromDb(userId string, keyId string) models.Key {
	// get guid fro session id
	db, err := gorm.Open("mysql", util.DbConfig)
	if err != nil {
		fmt.Print(err)
		// panic function to cause a run time error
		panic("failed to connect database")
	}
	// defer which schedules a function call to be run after the function completes
	defer db.Close()
	var key models.Key
	// Get all matched records
	db.Where("user_name = ? AND key_id = ?", userId, keyId).Find(&key)

	return key
}

// RetriveLatestPublicKeyAccordingToUser from "key" table
func RetriveLatestPublicKeyAccordingToUser() []models.Key {
	db, err := gorm.Open("mysql", util.DbConfig)
	if err != nil {
		fmt.Print(err)
		// panic function to cause a run time error
		panic("failed to connect database")
	}
	// defer which schedules a function call to be run after the function completes
	defer db.Close()
	var key []models.Key

	db.Where("user_name = ?", "chalani+19@entrusttitle.com").Find(&key)

	return key
}
