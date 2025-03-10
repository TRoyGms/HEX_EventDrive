package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	amqpURL := os.Getenv("AMQP_URL")
	amqpQueue := os.Getenv("AMQP_QUEUE")
	hostURL := os.Getenv("HOST_URL")

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(amqpQueue, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	//Limita la cantidad de mensajes simultáneos
	err = ch.Qos(1, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := ch.Consume(amqpQueue, "", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to start consuming: %v", err)
	}

	var wg sync.WaitGroup
	ticker := time.NewTicker(2 * time.Second) //Cooldown de 2s
	defer ticker.Stop()

	go func() {
		for msg := range msgs {
			wg.Add(1)
			go func(msg amqp.Delivery) {
				defer wg.Done()
				<-ticker.C

				log.Printf("Processing message: %s", msg.Body)

				var order struct {
					ID          int     `json:"id"`
					Name        string  `json:"name"`
					Amount      int     `json:"amount"`
					Description *string `json:"description"`
				}

				err := json.Unmarshal(msg.Body, &order)
				if err != nil {
					log.Printf("Failed to unmarshal order message: %v", err)
					_ = msg.Nack(false, false)
					return
				}

				client := resty.New()
				resp, err := client.R().
					SetBody(order).
					Post(fmt.Sprintf("%s/payments", hostURL))

				if err != nil || resp.StatusCode() >= 400 {
					log.Printf("Error processing order: %v | Status: %d", err, resp.StatusCode())
					_ = msg.Nack(false, true)
					return
				}

				log.Printf("Order sent to API2: %s", resp.String())
				_ = msg.Ack(false)
			}(msg)
		}
	}()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Consumer running"})
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}
}
