package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	email "github.com/vishnusunil243/Job-Portal-Email-service/internal/helper/EmailSend"
)

type ShortlistedUser struct {
	Email   string `json:"Email"`
	UserID  string `json:"UserID"`
	JobID   string `json:"JobID"`
	Company string `json:"company"`
}

func StartConsuming() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("ShortlistUser", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			var shortlistedUser ShortlistedUser
			err := json.Unmarshal(msg.Value, &shortlistedUser)
			fmt.Println("message recieved")
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			go func(user ShortlistedUser) {
				message := fmt.Sprintf("Congratulations! You have been shortlisted for the %s position at %s. Please visit [link to application portal] to proceed.", user.JobID, user.Company)
				if err := email.SendEmail(user.Email, message); err != nil {
					log.Println(err)
				}
			}(shortlistedUser)

		case err := <-partitionConsumer.Errors():
			log.Printf("Error consuming message: %v", err)
		}
	}
}
