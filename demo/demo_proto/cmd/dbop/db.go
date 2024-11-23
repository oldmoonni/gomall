package main

import (
	"github.com/joho/godotenv"
	"github.com/oldmoonni/gomall/demo/demo_proto/biz/dal"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	// Create
	// mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "123456"})

	// Update
	// mysql.DB.Model(&model.User{}).Where("Email =?", "demo@example.com").Update("password", "654321")

	// Read
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("Email =?", "demo@example.com").First(&row)
	// fmt.Printf("row: %+v\n", row)

	// Delete
	// mysql.DB.Where("Email =?", "demo@example.com").Delete(&model.User{})
	// mysql.DB.Unscoped().Where("Email =?", "demo@example.com").Delete(&model.User{})

}
