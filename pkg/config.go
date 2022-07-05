package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const port = ":8080"

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
		Port = port
	}

}
