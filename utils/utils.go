package utils

import (
	"crypto/rand"
	// "encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/Narutchai01/Project_S-BE/entities"
	"gorm.io/gorm"

	// "github.com/gofiber/fiber/v3/client"
	"github.com/golang-jwt/jwt"
	storage_go "github.com/supabase-community/storage-go"
	"gopkg.in/gomail.v2"
	// "gorm.io/gorm"
)

// type RoboflowResponse struct {
// 	Predictions []map[string]interface{} `json:"predictions"`
// }

func CheckDirectoryExist() error  {
	dir := "./uploads"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		    return err
		}
	}
	return nil
}

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

func UpdateImage(oldFilePath string, newFilePath string, folderName string) error {
	supa_api_url := config.GetEnv("SUPA_API_URL")
	supa_api_key := config.GetEnv("SUPA_API_KEY")
	bucket_name := config.GetEnv("SUPA_BUCKET_NAME")

	file, err := os.Open("./uploads/" + newFilePath)
	if err != nil {
		return fmt.Errorf("failed to open new file: %w", err)
	}
	defer file.Close()

	storageClient := storage_go.NewClient(supa_api_url, supa_api_key, nil)

	options := storage_go.FileOptions{
		ContentType: func() *string { contentType := "image/jpeg"; return &contentType }(),
	}

	_, err = storageClient.UpdateFile(bucket_name, folderName + "/" + oldFilePath, file, options)
	if err != nil {
		return fmt.Errorf("failed to update file %w", err)
	}

	return nil
}

func DeleteImage(oldFilePath string, folderName string) error {
      supa_api_url := config.GetEnv("SUPA_API_URL")
      supa_api_key := config.GetEnv("SUPA_API_KEY")
      bucket_name := config.GetEnv("SUPA_BUCKET_NAME")

      storageClient := storage_go.NewClient(supa_api_url, supa_api_key, nil)

      _, err := storageClient.RemoveFile(bucket_name, []string{folderName + "/" + oldFilePath})
      if err != nil {
            return fmt.Errorf("failed to update file %w", err)
      }

      return nil
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

func GenerateOTP() (string, error) {
      range_number := big.NewInt(900000)
	random_big, err := rand.Int(rand.Reader, range_number)
	if err != nil {
		return "", err
	}

	random_number := int(random_big.Int64()) + 100000
	return strconv.Itoa(random_number), nil
}

func SendEmailVerification(email string, otp string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", config.GetEnv("SENDING_EMAIL"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", "OTP verification from UCARE")
	message.SetBody("text/html", fmt.Sprintf(`
	    <html>
	    <body>
		  <h1> %s </h1>
		  <p>This is your UCARE verification code.</p>
		  <p>Don't share with anyone, including UCARE.</p>
	    </body>
	    </html>
	`, otp))
  
	dialer := gomail.NewDialer(
		"smtp.gmail.com",
	    	587, 
	    	config.GetEnv("SENDING_EMAIL"),
	    	config.GetEnv("EMAIL_PASSWORD"),
	)
  
	if err := dialer.DialAndSend(message); err != nil {
	    return fmt.Errorf("failed to send email OTP verification: %w", err)
	}

	return nil
}

func ExtractToken(token string) (uint, error) {
	secretKey := []byte(config.GetEnv("JWT_SECRET_KEY"))

	extractToken, err := jwt.Parse(strings.TrimPrefix(token, "Bearer "), func(extractToken *jwt.Token) (interface{}, error) {
		if _, ok := extractToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sign method")
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := extractToken.Claims.(jwt.MapClaims); ok && extractToken.Valid {
		user_id, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("user_id not found")
		}
		log.Println("Extracted user_id: ", user_id)
		return uint(user_id), nil
	}

	return 0, fmt.Errorf("invalid token")
}

func CheckEmptyValueBeforeUpdate(newValue string, oldValue string) string {
	if newValue == "" {
		return oldValue
	}
	return newValue
}

func FacialAcneSkinAnalysis(apiKey string, modelID string, imageURL string) (entities.Result, error) {
	var result entities.Result
	// fullURL := fmt.Sprintf("https://detect.roboflow.com/%s", modelID)

	// client := client.New()
	// client.SetParams(map[string]string{
	// 	"api_key": apiKey,
	// 	"image": imageURL,
	// })
	// resp, err := client.Post(fullURL)
	// if err != nil {
	// 	return entities.Result{}, err
	// }

	// err = json.Unmarshal(resp.Body(), &result)
	// if err != nil {
	// 	return entities.Result{}, err
	// }

	result.AcneType = []entities.Acne_Facial_Result{
		{ID: 1, Count: 10},
		{ID: 2, Count: 5},
	}
	result.FacialType = []entities.Acne_Facial_Result{
		{ID: 1, Count: 10},
		{ID: 2, Count: 5},
	}
	result.SkinType = 1
	result.Skincare = []entities.Skincare{
		{Model: gorm.Model{ID: 2}},
	}
	// result.Skincare = []uint{2}
	
	return result, nil
}
