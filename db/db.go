package db

import (
	"fmt"
	//"log"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/Narutchai01/Project_S-BE/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	pass := config.GetEnv("DB_PASS")
	name := config.GetEnv("DB_NAME")

	//log.Fatalln(host + " " + port + " " + user + " " + pass + " " + name)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", host, port, user, pass, name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Admin{}, &entities.Skincare{}, &entities.User{}, &entities.Recovery{}, &entities.Facial{}, &entities.Acne{}, &entities.Skin{}, &entities.Result{}, &entities.Thread{}, &entities.ThreadDetail{})

	return db, nil
}
