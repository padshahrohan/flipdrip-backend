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
	r.POST("/adduser",handlers.InsertUserData)
	r.POST("/addproduct",handlers.InsertProductData)

	r.Run()
}
