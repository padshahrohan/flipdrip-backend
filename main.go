package main

import (
	"github.com/gin-gonic/gin"
	"github.com/klipfart/handlers"
	"github.com/klipfart/initializers"
	//"github.com/gin-contrib/cors"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
// 	r.Use(cors.New(cors.Config{
//         AllowOrigins: []string{"*"},
//         AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
//         AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
//     }))
	r.POST("/transaction",handlers.InsertTransData)
	//user register
	r.POST("/user/register",handlers.InsertUserData)
	//user login
	r.POST("/user/login",handlers.Login)
	//login user
	//buyer
	//prouduct/list
	//product/buy
	//product/showLoyalty
	r.GET("/product/showLoyalty",handlers.ShowLoyalty)
	//seller
	//user/register and add to approvalList
	//product/add
	//getApprovalListOfBuyers
	//coinsApprovedForBuyers
	//Admin
	//getApprovalListOfSellers
	r.GET("/admin/getApprovalListOfSellers",handlers.GetAllApprovalListOfSellers)
	//coinsApprovedForSeller

	//product Add
	r.POST("/product/add",handlers.InsertProductData)
	//Not Required
	//r.POST("/addseller",handlers.InsertAdminApprovalData)
	//Product List
	r.GET("/product/list",handlers.GetAllProductData)
	r.POST("/insertReward",handlers.InsertLoyaltyPointsData)
	r.POST("/coinApproval",handlers.CoinApproval)
	r.POST("/sellerApproval",handlers.SellerApproval)
	r.Run()
}
