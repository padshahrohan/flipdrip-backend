package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klipfart/handlers"
	"github.com/klipfart/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/transaction",handlers.InsertTransData)
	r.POST("/user/register",handlers.InsertUserData)
	r.POST("/user/login",handlers.Login)
	//login user
	//buyer
	//prouduct/list
	//product/buy
	//product/showLoyalty
	//seller
	//user/register and add to approvalList
	//product/add
	//getApprovalListOfBuyers
	//coinsApprovedForBuyers
	//Admin
	//getApprovalListOfSellers
	//coinsApprovedForSeller
	r.POST("/addproduct",handlers.InsertProductData)
	r.POST("/addseller",handlers.InsertAdminApprovalData)
	r.GET("/getAllProductData",handlers.GetAllProductData)
	r.POST("/insertReward",handlers.InsertLoyaltyPointsData)
	r.POST("/coinApproval",handlers.CoinApproval)
	r.POST("/sellerApproval",handlers.SellerApproval)
	r.Run()
}
