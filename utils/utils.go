package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/golang-jwt/jwt"
	storage_go "github.com/supabase-community/storage-go"
)

func UploadImage(fileName string, dir string) (string, error) {

	supa_api_url := config.GetEnv("SUPA_API_URL")
	supa_api_key := config.GetEnv("SUPA_API_KEY")

	filePath := "./uploads/" + fileName

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	storageClient := storage_go.NewClient(supa_api_url, supa_api_key, nil)

	options := storage_go.FileOptions{
		ContentType: func() *string { s := "image/jpeg"; return &s }(),
	}
	// bucketName := "public"
	bucketName := config.GetEnv("SUPA_BUCKET_NAME")
	fileName = dir + "/" + fileName
	_, err = storageClient.UploadFile(bucketName, fileName, file, options)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	url := fmt.Sprintf("%s/object/public/%s/%s", supa_api_url, bucketName, fileName)

	return url, nil
}

func CreateJWTToken(secretKey string, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func GenerateToken(userID int) (string, error) {
	secretKey := config.GetEnv("JWT_SECRET_KEY")
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	return CreateJWTToken(secretKey, claims)
}
