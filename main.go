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
	r.POST("/user/register", handlers.InsertUserData) //checked
	r.POST("/user/login", handlers.Login)             //checked

	//Buyer
	r.POST("/product/buy", handlers.InsertTransData) //Checked
	r.GET("/product/showLoyalty", handlers.ShowLoyalty)
	//Redeem

	//Seller
	r.POST("/product/add", handlers.InsertProductData)                         //checked
	r.GET("/seller/getApprovalListOfBuyers", handlers.GetApprovalListOfBuyers) //checked
	r.POST("/seller/buyersTokensApproved", handlers.CoinApproval)              //checked -> response to be finalised
	r.GET("/product/list", handlers.GetAllProductData)                         //checked

	//Admin
	r.GET("/admin/getApprovalListOfSellers", handlers.GetAllApprovalListOfSellers) //Checked
	r.PUT("/admin/sellerTokensApproved", handlers.SellerApproval)                  //Checked

	r.POST("/insertReward", handlers.InsertLoyaltyPointsData)

	//Get Wallet Address
	r.POST("/getWalletAddress", handlers.GetWalletAddress)

	//Wallet Address
	r.PUT("/getUsersForWalletAddresses", handlers.GetUsersForWalletAddresses) //checked

	r.Run()
}
