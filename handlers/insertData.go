package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/klipfart/initializers"
	"github.com/klipfart/models"
	"gorm.io/gorm"
)

func init() {
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

	trans := models.Trans{BuyerId: body.BuyerId, SellerId: body.SellerId, ProductId: body.ProductId}
	result := initializers.DB.Create(&trans)
	if result.Error != nil {
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	var existingReward models.Reward

	result_entry := initializers.DB.Model(&existingReward).
		Where("buyer_id = ? AND seller_id = ?", body.BuyerId, body.SellerId).
		First(&existingReward)

	if result_entry.Error != nil {
		if result_entry.Error == gorm.ErrRecordNotFound {
			loyaltyRewardEntry := models.Reward{SellerId: body.SellerId, BuyerId: body.BuyerId, Coins: 1, Count: 1}
			resultEntry := initializers.DB.Create(&loyaltyRewardEntry)
			if resultEntry.Error != nil {
				return
			}
			c.JSON(200, gin.H{"result": loyaltyRewardEntry})
			return
		} else {
			c.Status(400)
			log.Fatal("Error While Fetching DB Query")
			return
		}
	} else {
		count := existingReward.Count + 1

		var loyaltyPointsDtl models.LoyaltyPointsDtl
		err := initializers.DB.Model(&loyaltyPointsDtl).
			Where("starting_range <= ? AND ending_range >= ?", count, count).
			First(&loyaltyPointsDtl)

		if err != nil {
			c.Status(400)
			log.Fatal("Error while fetching LoyaltyPointsDtl")
			return
		}
		fmt.Println(loyaltyPointsDtl.Coins)
		coins := existingReward.Coins + loyaltyPointsDtl.Coins

		loyaltyRewardEntry := models.Reward{
			SellerId: existingReward.SellerId,
			BuyerId:  existingReward.BuyerId,
			Coins:    coins,
			Count:    count,
		}

		result := initializers.DB.Model(&models.Reward{}).
			Where("seller_id = ? AND buyer_id = ?", body.SellerId, body.BuyerId).
			Updates(map[string]interface{}{
				"coins": coins,
				"count": count,
			})

		if result.Error != nil {
			c.Status(400)
			log.Fatal("Getting Error while updating data in db")
			return
		}

		c.JSON(200, gin.H{"insertedRecord": loyaltyRewardEntry})
		return
	}

	//c.JSON(200, gin.H{"result": trans})
}

func InsertUserData(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	var user models.Users
	//Binding the object
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatal("Not Able to bind the object")
		return
	}
	user_entry := models.Users{UserName: user.UserName, Role: user.Role, UserPassword: user.UserPassword, WalletAddress: user.WalletAddress,Name:user.Name}
	result_entry := initializers.DB.Create(&user_entry)
	if result_entry.Error != nil {
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	if user.Role == "Seller" {
		adminApproval := models.AdminApproval{SellerId: int16(user_entry.ID), UserName: user.UserName}
		resultAdminApproval_entry := initializers.DB.Create(&adminApproval)
		if resultAdminApproval_entry.Error != nil {
			c.Status(400)
			log.Fatal("Getting Error while fetching data from db")
			return
		}
	}
	c.JSON(200, gin.H{"result": user_entry})

}

func InsertProductData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Methods", "*")
	// c.Header("Access-Control-Allow-Headers", "*")

	var product models.Product
	//Binding the object
	if err := c.ShouldBindJSON(&product); err != nil {
		log.Fatal("Not Able to bind the object")
		return
	}

	product_entry := models.Product{ProductDescription: product.ProductDescription, ProductPrice: product.ProductPrice, SellerId: product.SellerId, Tokens: product.Tokens, ProductName: product.ProductName}
	result_entry := initializers.DB.Create(&product_entry)
	if result_entry.Error != nil {
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	c.JSON(200, gin.H{"result": product_entry})

}

func InsertAdminApprovalData(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	var sellerID int16
	var userUsername string
	var userName string
	var role string
	var userPassword string
	var walletAddress string

	// Parse the JSON request body
	if err := c.ShouldBindJSON(&gin.H{
		"SellerId":      &sellerID,
		"Userusername":  &userUsername,
		"UserName":      &userName,
		"Role":          &role,
		"UserPassword":  &userPassword,
		"WalletAddress": &walletAddress,
	}); err != nil {
		log.Fatal("Not Able to bind the object")
		return
	}

	// var adminApproval models.AdminApproval
	// //var user models.Users

	// 	//Binding the object
	// 	if err := c.ShouldBindJSON(&adminApproval); err != nil {
	// 		log.Fatal("Not Able to bind the object")
	// 		return
	// 	}
	//fmt.Println("sahdfahfbhsdfbhsbdsfbzshfdbcisdcdbcuedi")
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	log.Fatal("Not Able to bind the object")
	// 	return
	// }
	// sellerid:=
	//Inserting in AdminApproval Table
	adminApproval_entry := models.AdminApproval{SellerId: sellerID, UserName: userName}
	resultAdminApproval_entry := initializers.DB.Create(&adminApproval_entry)
	if resultAdminApproval_entry.Error != nil {
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	//InsertUserData(c)
	//Inserting in Users Table
	user_entry := models.Users{UserName: userName, Role: role, UserPassword: userPassword, WalletAddress: walletAddress}
	resultUser_entry := initializers.DB.Create(&user_entry)
	if resultUser_entry.Error != nil {
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	c.JSON(200, gin.H{"resultAdminApproval": adminApproval_entry, "resultUserData": "user_entry"})

}
func CoinApproval(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	var reward models.Reward
	//Binding the object
	if err := c.ShouldBindJSON(&reward); err != nil {
		c.Status(400)
		log.Fatal("Not Able to bind the object")
		return
	}
	loyaltyReward_entry := models.Reward{SellerId: reward.SellerId, BuyerId: reward.BuyerId, Coins: 0}
	result := initializers.DB.Model(&models.Reward{}).
		Where("seller_id = ? AND buyer_id = ?", reward.SellerId, reward.BuyerId).
		Updates(map[string]interface{}{
			"coins": 0,
		})
	if result.Error != nil {
		c.Status(400)
		log.Fatal("Getting Error while fetching data from db")
		return
	}
	c.JSON(200, gin.H{"insertedRecord": loyaltyReward_entry})
	return

}

func InsertLoyaltyPointsData(c *gin.Context) {

	var reward models.Reward

	// Binding the object
	if err := c.ShouldBindJSON(&reward); err != nil {
		c.Status(400)
		log.Fatal("Not Able to bind the object")
		return
	}

	var existingReward models.Reward

	result := initializers.DB.Model(&existingReward).
		Where("buyer_id = ? AND seller_id = ?", reward.BuyerId, reward.SellerId).
		First(&existingReward)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			loyaltyRewardEntry := models.Reward{SellerId: reward.SellerId, BuyerId: reward.BuyerId, Coins: 1, Count: 1}
			resultEntry := initializers.DB.Create(&loyaltyRewardEntry)
			if resultEntry.Error != nil {
				return
			}
			c.JSON(200, gin.H{"result": loyaltyRewardEntry})
			return
		} else {
			c.Status(400)
			log.Fatal("Error While Fetching DB Query")
			return
		}
	} else {
		count := existingReward.Count + 1

		var loyaltyPointsDtl models.LoyaltyPointsDtl
		err := initializers.DB.Model(&loyaltyPointsDtl).
			Where("starting_range <= ? AND ending_range >= ?", count, count).
			First(&loyaltyPointsDtl)

		if err != nil {
			c.Status(400)
			log.Fatal("Error while fetching LoyaltyPointsDtl")
			return
		}
		fmt.Println(loyaltyPointsDtl.Coins)
		coins := existingReward.Coins + loyaltyPointsDtl.Coins

		loyaltyRewardEntry := models.Reward{
			SellerId: existingReward.SellerId,
			BuyerId:  existingReward.BuyerId,
			Coins:    coins,
			Count:    count,
		}

		result := initializers.DB.Model(&models.Reward{}).
			Where("seller_id = ? AND buyer_id = ?", reward.SellerId, reward.BuyerId).
			Updates(map[string]interface{}{
				"coins": coins,
				"count": count,
			})

		if result.Error != nil {
			c.Status(400)
			log.Fatal("Getting Error while updating data in db")
			return
		}

		c.JSON(200, gin.H{"insertedRecord": loyaltyRewardEntry})
		return
	}
}

func Login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
	var user models.Users
	//Binding the object
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Fatal("Not Able to bind the object")
		return
	}
	fmt.Println(user.WalletAddress)
	user = models.Users{UserName: user.UserName, UserPassword: user.UserPassword, WalletAddress: user.WalletAddress, Role: user.Role}
	result := initializers.DB.Model(&user).Where("user_name = ? AND user_password = ? AND wallet_address LIKE ?", user.UserName, user.UserPassword, user.WalletAddress).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(200, gin.H{"result": "Login Details Not Found"})
			return
		} else {
			c.Status(400)
			log.Fatal("Error While Fetching DB Query")
			return
		}
	}
	user = models.Users{UserName: user.UserName, UserPassword: user.UserPassword, WalletAddress: user.WalletAddress, Role: user.Role}
	c.JSON(200, gin.H{"result": user})

}
