package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	//dsn := "host=rain.db.elephantsql.com user=ykcgtjpe password=zDnWh9zc1Rgx5fjLx-kxOfk4FsT0_tDy dbname=ykcgtjpe port=5432 sslmode=disable"
	dsn := "host=localhost user=postgres password=pass123 dbname=flipdrip port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed To Connect To DB")
	}
}
