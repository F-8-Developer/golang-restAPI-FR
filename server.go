package main
import (
	"fmt"

	"golang-restAPI-FR/Config"
	"golang-restAPI-FR/Database"
	"golang-restAPI-FR/Core/Router"
	"golang-restAPI-FR/Core/Models"
)

// Api server start from here. router is define your api router and public it.
func main() {
	// GORM DATABASE
	Database.Mysql, Database.Err = Database.ConnectToDB("main")
	if Database.Err != nil {
		fmt.Println("status error : ", Database.Err)
		return
	} else {
		fmt.Println("database connected")
	}
	defer Database.Mysql.Close()
	// auto migrate
	Database.Mysql.AutoMigrate(&Models.User{})
	Database.Mysql.AutoMigrate(&Models.Friend{})

	app_env := Config.GoDotEnvVariable("APP_ENV")
	Router.Start(app_env)
}
