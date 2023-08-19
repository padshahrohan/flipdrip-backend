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

func GetAllApprovalListOfSellers(c *gin.Context) {

	var adminApproval []models.Users
	//initializers.DB.Find(&adminApproval)
	initializers.DB.
		Model(models.Users{}).
		Joins("INNER JOIN admin_approvals ON admin_approvals.seller_id = users.ID").
		Find(&adminApproval)
	c.JSON(200, gin.H{"result": adminApproval})

}
func ShowLoyalty(c *gin.Context) {

	var products []models.Product

	sellerId := c.Query("SellerId") // Get the SellerId query parameter from the request
	buyerId := c.Query("BuyerId")

	if sellerId != "" {
		// If SellerId is provided, filter products based on the sellerId
		initializers.DB.Where("seller_id = ? AND buyer_id = ?", sellerId, buyerId).Find(&products)
	} else {
		// If no SellerId is provided, fetch all products
		initializers.DB.Where("buyer_id = ?", buyerId).Find(&products)
	}

	c.JSON(200, gin.H{"result": products})
}

// func GetWalletAddress(c *gin.Context) {

// 	var user models.Users

// 	id := c.Query("Id")

// 	if id != "" {
// 		// If SellerId is provided, filter products based on the sellerId
// 		initializers.DB.Where("id= ?", id).Find(&user)
// 	}

// 	c.JSON(200, gin.H{"result": user})
// }
func GetWalletAddress(c *gin.Context) {
	var users []models.Users
	var input struct {
		Ids []string `json:"ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}
	if len(input.Ids) > 0 {
		initializers.DB.Where("id IN (?)", input.Ids).Find(&users)
	}
	c.JSON(200, gin.H{"result": users})
}
func GetUsersForWalletAddresses(c *gin.Context) {
	var users []models.Users
	var input struct {
		WalletAdd []string `json:"WalletAdd"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input data"})
		return
	}
	if len(input.WalletAdd) > 0 {
		initializers.DB.Where("wallet_address IN (?)", input.WalletAdd).Find(&users)
	}

	// Create a map to store the result
	resultMap := make(map[string]string)
	for _, user := range users {
		resultMap[user.WalletAddress] = user.Name
	}

	c.JSON(200, gin.H{"result": resultMap})
}

func GetApprovalListOfBuyers(c *gin.Context) {

	type BuyerDetailsResponse struct {
		models.Users
		Coins int
	}

	var buyerDetails []BuyerDetailsResponse
	sellerID := c.Query("SellerId")

	initializers.DB.
		Model(models.Users{}).
		Select("users.*, rewards.coins AS coins").
		Joins("INNER JOIN rewards ON rewards.buyer_id = users.ID").
		Where("rewards.seller_id = ? AND rewards.coins > ?", sellerID, 0).
		Find(&buyerDetails)

	c.JSON(200, gin.H{"buyerDetails": buyerDetails})

}
