package helper

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
}

func generateOTP() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}
func SendOTP(email string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", viper.GetString("SMTP_USER"))
	message.SetHeader("Subject", "OTP Verification")
	otp := generateOTP()

	message.SetBody("text/plain", "THIS WILL EXPIRE IN 5 MINUTES \n YOUR OTP IS : "+otp)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, viper.GetString("SMTP_USER"), viper.GetString("SMTP_PASSWORD"))
	otpKey := fmt.Sprintf("otp:%s", email)
	err := redisClient.Set(otpKey, otp, 300*time.Second).Err()
	if err != nil {
		log.Println("failed to store otp in redis")
		return err
	}
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
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
