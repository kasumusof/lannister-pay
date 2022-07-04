package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

}
