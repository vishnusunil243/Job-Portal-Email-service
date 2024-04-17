package otp

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"gopkg.in/gomail.v2"
)

var redisClient *redis.Client

func init() {
	fmt.Println("hii from init redis")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: "",
		DB:       0,
	})
}

func generateOTP() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}
func SendOTP(email string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("SMTP_USER"))
	message.SetHeader("To", email)
	message.SetHeader("Subject", "OTP Verification")
	otp := generateOTP()
	message.SetBody("text/plain", "THIS WILL EXPIRE IN 5 MINUTES \n YOUR OTP IS : "+otp)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	otpKey := fmt.Sprintf("otp:%s", email)
	err := redisClient.Set(otpKey, otp, 300*time.Second).Err()
	if err != nil {
		log.Println("failed to store otp in redis")
		log.Println("error is : ", err)
		return err
	}
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	fmt.Println("otp successfully sent to : ", email)
	return nil
}
func GetStoredOTP(email string) (string, error) {
	otpKey := fmt.Sprintf("otp:%s", email)
	otp, err := redisClient.Get(otpKey).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("OTP not found")
	} else if err != nil {
		return "", err
	}
	return otp, nil
}
func VerifyOTP(email, otp string) bool {
	storedotp, err := GetStoredOTP(email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if otp == storedotp {
		return true
	}
	return false
}
