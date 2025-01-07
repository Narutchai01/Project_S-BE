package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/Narutchai01/Project_S-BE/config"
	"github.com/golang-jwt/jwt"
	storage_go "github.com/supabase-community/storage-go"
	"gopkg.in/gomail.v2"
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