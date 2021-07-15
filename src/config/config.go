package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// mongo uri
	DB_URI = ""

	// api port
	Port = 0
)

// Load env variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// String to int
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 3003
	}

	DB_URI = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
}
