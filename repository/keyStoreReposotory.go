package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/student-management-service/models"
	"github.com/student-management-service/util"
)

// RetriveLatestPublicKeyAccordingToUser from "key_stores" table
func GetAllPublicKeysOfUser(userId string) []models.KeyStore {
	// get guid fro session id
	db, err := gorm.Open("mysql", util.DbConfig)
	if err != nil {
		fmt.Print(err)
		// panic function to cause a run time error
		panic("failed to connect database")
	}
	// defer which schedules a function call to be run after the function completes
	defer db.Close()

	var mobileKeys []models.KeyStore
	db.Where("user_id = ?", userId).Find(&mobileKeys)
	return mobileKeys
}

func GetAllPublicKeyAccordingToKeyId(KeyId string) models.KeyStore {
	// get guid fro session id
	db, err := gorm.Open("mysql", util.DbConfig)
	if err != nil {
		fmt.Print(err)
		// panic function to cause a run time error
		panic("failed to connect database")
	}
	// defer which schedules a function call to be run after the function completes
	defer db.Close()

	var mobileKeys models.KeyStore
	db.Where("key_id = ?", KeyId).Find(&mobileKeys)
	return mobileKeys
}
