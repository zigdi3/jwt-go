package launcher

import "jwt-dev/models"

func SyncDataBases() {
	DB.AutoMigrate(&models.User{})
}