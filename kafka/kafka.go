package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	email "github.com/vishnusunil243/Job-Portal-Email-service/internal/helper/EmailSend"
)

type ShortlistedUser struct {
	Email       string `json:"Email"`
	Designation string `json:"Designation"`
	JobID       string `json:"JobID"`
	Company     string `json:"company"`
}

func StartConsumingShortlist() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	// config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("ShortlistUser", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var shortlistedUser ShortlistedUser
			err := json.Unmarshal(msg.Value, &shortlistedUser)
			fmt.Println("message received")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			go func(user ShortlistedUser) {
				message := fmt.Sprintf("Congratulations! You have been shortlisted for the %s position at %s.", user.Designation, user.Company)
				if err := email.SendEmail(user.Email, message); err != nil {
					log.Println(err)
				}
			}(shortlistedUser)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}

type InterviewSchedule struct {
	Email       string `json:"Email"`
	Designation string `json:"Designation"`
	Date        string `json:"Date"`
	Company     string `json:"Company"`
	RoomId      string `json:"RoomId"`
}

func StartConsumingInterviewSchedule() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	// config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("InterviewSchedule", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var interviewSchedule InterviewSchedule
			err := json.Unmarshal(msg.Value, &interviewSchedule)
			fmt.Println("message received from interview schedule")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			go func(interviewSchedule InterviewSchedule) {
				message := fmt.Sprintf("your interview for the position of %s has been scheduled by the company %s at %s your room id is : %s", interviewSchedule.Designation, interviewSchedule.Company, interviewSchedule.Date, interviewSchedule.RoomId)
				if err := email.SendEmail(interviewSchedule.Email, message); err != nil {
					log.Println(err)
				}
			}(interviewSchedule)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}

type Hired struct {
	Designation string `json:"Designation"`
	Email       string `json:"Email"`
	Company     string `json:"Company"`
	Date        string `json:"Date"`
}

func StartConsumingHired() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	// config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("Hired", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var hired Hired
			err := json.Unmarshal(msg.Value, &hired)
			fmt.Println("message received from interview schedule")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			go func(hired Hired) {
				message := fmt.Sprintf("In light of the interview conducted on %s we are pleased to inform you that the company %s have decided to hire you for the position of %s, you will be recieving an offer letter soon", hired.Date, hired.Company, hired.Designation)
				if err := email.SendEmail(hired.Email, message); err != nil {
					log.Println(err)
				}
			}(hired)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}

type Warning struct {
	UserName     string `json:"UserName"`
	Designation  string `json:"Designation"`
	Company      string `json:"Company"`
	CompanyId    string `json:"CompanyId"`
	Date         string `json:"Date"`
	UserEmail    string `json:"UserEmail"`
	CompanyEmail string `json:"CompanyEmail"`
	RoomId       string `json:"RoomId"`
}

func StartConsumingWarning() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("Warning", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var warning Warning
			err := json.Unmarshal(msg.Value, &warning)
			fmt.Println("message received warning")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}
			go func(warning Warning) {
				message := fmt.Sprintf("This is a remainder message for your interview for the job opening %s for the user %s on %s with roomId %s", warning.Designation, warning.UserName, warning.Date, warning.RoomId)
				if err := email.SendEmail(warning.CompanyEmail, message); err != nil {
					log.Println(err)
				}
				userMsg := fmt.Sprintf("This is a remainder for your interview for the job opening %s at %s which is scheduled on %s with roomId %s", warning.Designation, warning.Company, warning.Date, warning.RoomId)
				if err := email.SendEmail(warning.UserEmail, userMsg); err != nil {
					log.Println(err)
				}
			}(warning)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}

type SubscriptionEnding struct {
	Email string `json:"Email"`
}

func StartConsumingSubscriptionEnding() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("SubscriptionEnding", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var sub SubscriptionEnding
			err := json.Unmarshal(msg.Value, &sub)
			fmt.Println("message received subscription ending")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}
			go func(sub SubscriptionEnding) {
				message := "This is a remainder regarding your subscription which is ending tomorrow so please subscribe to continue using our services"
				if err := email.SendEmail(sub.Email, message); err != nil {
					log.Println(err)
				}

			}(sub)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}

type Subscribed struct {
	Email    string `json:"Email"`
	Duration string `json:"Duration"`
}

func StartConsumingSubscribed() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = true
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition("Subscribed", 0, sarama.OffsetNewest)
	fmt.Println("offset ", sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var sub Subscribed
			err := json.Unmarshal(msg.Value, &sub)
			fmt.Println("message received subscribed")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}
			go func(sub Subscribed) {
				message := "You have subscribed to our service for the duration of " + sub.Duration
				if err := email.SendEmail(sub.Email, message); err != nil {
					log.Println(err)
				}

			}(sub)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
