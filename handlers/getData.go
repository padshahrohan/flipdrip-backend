package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/klipfart/initializers"
	"github.com/klipfart/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}
func GetAllProductData(c * gin.Context){
	var products []models.Product
	initializers.DB.Find(&products)

	c.JSON(200,gin.H{"result":products})


}