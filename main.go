package main

import (
	"fmt"
	"igaming-service/server"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("logs/"); os.IsNotExist(err) {
		if err = os.Mkdir("logs", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.OpenFile("logs/logs "+string(time.Now().Format("2006-01-02T15:04:05"))+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}

	log.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8888
	}
	fmt.Printf("Server started on port: %v\n", port)
	server.Run(port)
}
