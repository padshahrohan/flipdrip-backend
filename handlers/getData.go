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
// func GetAllProductData(c * gin.Context){
// 	var products []models.Product
// 	initializers.DB.Find(&products)

// 	c.JSON(200,gin.H{"result":products})


// }
func GetAllProductData(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	var products []models.Product

	sellerId := c.Query("SellerId") // Get the SellerId query parameter from the request

	if sellerId != "" {
		// If SellerId is provided, filter products based on the sellerId
		initializers.DB.Where("seller_id = ?", sellerId).Find(&products)
	} else {
		// If no SellerId is provided, fetch all products
		initializers.DB.Find(&products)
	}

	c.JSON(200, gin.H{"result": products})
}

func GetAllApprovalListOfSellers(c * gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	var adminApproval []models.AdminApproval
	initializers.DB.Find(&adminApproval)

	c.JSON(200,gin.H{"result":adminApproval})


}
func ShowLoyalty(c * gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	var products []models.Product

	sellerId := c.Query("SellerId") // Get the SellerId query parameter from the request
	buyerId:=c.Query("BuyerId")

	if sellerId != "" {
		// If SellerId is provided, filter products based on the sellerId
		initializers.DB.Where("seller_id = ? AND buyer_id = ?", sellerId,buyerId).Find(&products)
	} else {
		// If no SellerId is provided, fetch all products
		initializers.DB.Where("buyer_id = ?",buyerId).Find(&products)
	}

	c.JSON(200, gin.H{"result": products})
}