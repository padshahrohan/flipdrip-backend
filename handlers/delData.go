package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/klipfart/initializers"
	"github.com/klipfart/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func SellerApproval(c * gin.Context) {

	var adminApproval models.AdminApproval
	if err := c.ShouldBindJSON(&adminApproval); err != nil {
		c.Status(400)
		log.Fatal("Not Able to bind the object")
		return
	}
	result:=initializers.DB.Where("seller_id = ?", adminApproval.SellerId).Delete(&adminApproval)
	if result.Error!=nil{
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	c.JSON(200,gin.H{"Message":"Approved Successfully!"})
	return


}