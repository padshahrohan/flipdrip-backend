package main

import (
	"github.com/klipfart/initializers"
	"github.com/klipfart/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Trans{})
	initializers.DB.AutoMigrate(&models.Reward{})
	initializers.DB.AutoMigrate(&models.Product{})
	initializers.DB.AutoMigrate(&models.Users{})
	initializers.DB.AutoMigrate(&models.AdminApproval{})
}