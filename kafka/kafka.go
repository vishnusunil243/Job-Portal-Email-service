package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	email "github.com/vishnusunil243/Job-Portal-Email-service/internal/helper/EmailSend"
)

type ShortlistedUser struct {
	Email   string `json:"Email"`
	UserID  string `json:"UserID"`
	JobID   string `json:"JobID"`
	Company string `json:"company"`
}

func StartConsuming() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "ShortlistConsumers",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		var shortlistedUser ShortlistedUser
		err = json.Unmarshal(msg.Value, &shortlistedUser)
		if err != nil {
			log.Println(err)
			continue
		}

		go func(user ShortlistedUser) {
			message := fmt.Sprintf("Congratulations! You have been shortlisted for the %s position at %s. Please visit [link to application portal] to proceed.", user.JobID, user.Company)
			if err := email.SendEmail(user.Email, message); err != nil {
				log.Println(err)
			}
		}(shortlistedUser)
	}
}
