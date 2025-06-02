package initializers

import "ExpenseAPI/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to sync User database")
	}

	err = DB.AutoMigrate(&models.Expense{})
	if err != nil {
		panic("Failed to sync Expense database")
	}
}
