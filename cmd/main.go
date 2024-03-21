package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/vishnusunil243/Job-Portal-Email-service/db"
	"github.com/vishnusunil243/Job-Portal-Email-service/initializer"
	"github.com/vishnusunil243/Job-Portal-Email-service/kafka"
	"github.com/vishnusunil243/Job-Portal-proto-files/pb"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf(err.Error())
	}
	addr := os.Getenv("DB_KEY")
	DB, err := db.InitMongoDB(addr)
	if err != nil {
		log.Fatal("error connecting to database")

	}
	listener, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatal("failed to listen on port 8087")
	}
	fmt.Println("email service listening on port 8087")
	services := initializer.Initializer(DB)
	server := grpc.NewServer()
	pb.RegisterEmailServiceServer(server, services)
	go kafka.StartConsuming()
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to listen on port 8087")
	}

}
