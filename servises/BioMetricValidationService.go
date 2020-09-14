package servises

import (
	"time"

	"github.com/student-management-service/models"
	repo "github.com/student-management-service/repository"
	viewModel "github.com/student-management-service/viewModels"
)

func OnboardToDigitalSignature(onboard viewModel.DigitalSignatureOnboard) (bool, string) {

	// insert to db
	var keyTable models.Key
	keyTable.UserName = onboard.Username
	keyTable.KeyId = ""
	keyTable.CreatedTime = time.Now()
	keyTable.Type = onboard.KeyType

	repo.UserOnboardRepository(keyTable)

	return true, keyTable.KeyId
}
