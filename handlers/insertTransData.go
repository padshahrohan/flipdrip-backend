package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/klipfart/initializers"
	"github.com/klipfart/models"
)
func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func InsertTransData(c *gin.Context) {
	var body models.Trans
	var _ gin.H

		//Binding the object
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Fatal("Not Able to bind the object")
		return
	}

	trans:=models.Trans{BuyerId: body.BuyerId,SellerId: body.SellerId,ProductId: body.ProductId}
	result:=initializers.DB.Create(&trans)
	if result.Error!=nil{
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	c.JSON(200,gin.H{"result":trans})
}

func InsertUserData(c * gin.Context){

	var user models.Users
		//Binding the object
		if err := c.ShouldBindJSON(&user); err != nil {
			log.Fatal("Not Able to bind the object")
			return
		}
		user_entry:=models.Users{UserName: user.UserName,UserType:user.UserType,Userusername:user.UserName,UserPassword:user.UserPassword,WalletAddress: user.WalletAddress}
		result_entry:=initializers.DB.Create(&user_entry)
		if result_entry.Error!=nil{
			c.Status(400)
			log.Fatal("Getting Error while fetching data from db")
			return
		}
		c.JSON(200,gin.H{"result":user_entry})
	
}

func InsertProductData(c * gin.Context){

	var product models.Product
		//Binding the object
		if err := c.ShouldBindJSON(&product); err != nil {
			log.Fatal("Not Able to bind the object")
			return
		}
		product_entry:=models.Product{ProductDescription: product.ProductDescription,ProductPrice:product.ProductPrice,SellerId: product.SellerId}
		result_entry:=initializers.DB.Create(&product_entry)
		if result_entry.Error!=nil{
			c.Status(400)
			log.Fatal("Getting Error while fetching data from db")
			return
		}
		c.JSON(200,gin.H{"result":product_entry})
	
}