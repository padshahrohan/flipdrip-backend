package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=rain.db.elephantsql.com user=nxtmliar password=vkPgteDZ5vuOSBQGq5ayqJ24Zd_Crgnz dbname=nxtmliar port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil{
		log.Fatal("Failed To Connect To DB")
	}
}