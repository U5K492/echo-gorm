package migration

import "echogorm/infrastructure/database"

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Migration() {
	database.Connect()
	database.DB.AutoMigrate(&User{})
}
