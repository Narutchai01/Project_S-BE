package db

import (
	"fmt"
	"log"
	"os"
	"time"

	//"log"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/Narutchai01/Project_S-BE/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {

	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	pass := config.GetEnv("DB_PASS")
	name := config.GetEnv("DB_NAME")

	//log.Fatalln(host + " " + port + " " + user + " " + pass + " " + name)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok", host, port, user, pass, name)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entities.Admin{}, &entities.Skincare{}, &entities.User{}, &entities.Recovery{}, &entities.Facial{}, &entities.Acne{}, &entities.Skin{}, &entities.Result{}, &entities.Thread{}, &entities.ThreadDetail{}, entities.Bookmark{}, entities.Comment{}, entities.FavoriteComment{})

	return db, nil
}
