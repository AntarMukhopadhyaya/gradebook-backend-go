package helpers

import "gradebook-api/models"

func MigrateDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Question{}, &models.Assignment{})
}
