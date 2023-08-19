package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/klipfart/handlers"
	"github.com/klipfart/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin,access-control-allow-headers"},
	}))

	//Common
	r.POST("/user/register", handlers.InsertUserData)
	r.POST("/user/login", handlers.Login)

	//Buyer
	r.POST("/product/buy", handlers.InsertTransData)
	r.GET("/product/showLoyalty", handlers.ShowLoyalty)
	//Redeem
	
	//Seller
	r.POST("/product/add", handlers.InsertProductData)
	r.GET("/seller/getApprovalListOfBuyers", handlers.GetAllApprovalListOfSellers)
	r.POST("/seller/buyersTokensApproved", handlers.CoinApproval)
	r.GET("/product/list", handlers.GetAllProductData)

	//Admin
	r.GET("/admin/getApprovalListOfSellers", handlers.GetAllApprovalListOfSellers)
	r.PUT("/admin/sellerTokensApproved", handlers.SellerApproval)

	// r.POST("/insertReward", handlers.InsertLoyaltyPointsData)
	
	//Get Wallet Address
	r.POST("/getWalletAddress", handlers.GetWalletAddress)

	//Wallet Address
	r.POST("/getUsersForWalletAddresses", handlers.GetUsersForWalletAddresses)

	r.Run()
}
