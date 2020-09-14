package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/student-management-service/models"
	"github.com/student-management-service/util"
)

func BioMetricMobileResponse(corrId string) models.MobileResponse {

	db, err := gorm.Open("mysql", util.DbConfig)
	if err != nil {
		fmt.Print(err)
		// panic function to cause a run time error
		panic("failed to connect database")
	}
	// defer which schedules a function call to be run after the function completes
	defer db.Close()
	var response models.MobileResponse
	// Get all matched records
	db.Where("correlation_id = ?", corrId).Find(&response)
	return response

}
