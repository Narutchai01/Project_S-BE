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
			LogLevel:                  logger.Warn, // Log level
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

	db.AutoMigrate(entities.CommunityType{}, entities.Community{}, entities.CommunityImage{}, entities.User{}, entities.Skincare{}, entities.Admin{}, entities.SkincareCommunity{}, entities.Comment{}, entities.Favorite{}, entities.Bookmark{}, entities.FaceProblemType{}, entities.FaceProblem{}, entities.Recovery{}, entities.Follower{}, entities.Result{}, entities.SkincareResult{})

	return db, nil
}

func Seeds(db *gorm.DB) {
	// initialize the database vactor
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS vector").Error; err != nil {
		log.Printf("Error creating vector extension: %v", err)
		// Decide whether to panic or continue
	}
	type_community := []entities.CommunityType{
		{Type: "thread"},
		{Type: "review"},
		{Type: "comment"},
	}

	type_face_problem := []entities.FaceProblemType{
		{Name: "acne"},
		{Name: "facial"},
		{Name: "skin"},
	}

	for _, communityType := range type_community {
		var existing entities.CommunityType
		if err := db.Where("type = ?", communityType.Type).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&communityType)
			}
		}
	}

	for _, faceProblemType := range type_face_problem {
		var existing entities.FaceProblemType
		if err := db.Where("name = ?", faceProblemType.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&faceProblemType)
			}
		}
	}
}

func ManageOTP(db *gorm.DB) {
	// Start a goroutine that runs every 5 minutes
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			expireTime := time.Now().Add(-5 * time.Minute)
			result := db.Where("created_at < ?", expireTime).Delete(&entities.Recovery{})
			if result.Error != nil {
				log.Printf("Error deleting expired OTP records: %v", result.Error)
			} else if result.RowsAffected > 0 {
				log.Printf("Deleted %d expired OTP records", result.RowsAffected)
			}
		}
	}()

	// Also run once immediately on startup
	expireTime := time.Now().Add(-5 * time.Minute)
	result := db.Where("created_at < ?", expireTime).Delete(&entities.Recovery{})
	if result.Error != nil {
		log.Printf("Error deleting expired OTP records: %v", result.Error)
	} else if result.RowsAffected > 0 {
		log.Printf("Deleted %d expired OTP records", result.RowsAffected)
	}
}
