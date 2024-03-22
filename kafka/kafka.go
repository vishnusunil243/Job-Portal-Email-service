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
				message := fmt.Sprintf("your interview for the position of %s has been scheduled by the company %s at %s", interviewSchedule.Designation, interviewSchedule.Company, interviewSchedule.Date)
				if err := email.SendEmail(interviewSchedule.Email, message); err != nil {
					log.Println(err)
				}
			}(interviewSchedule)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
