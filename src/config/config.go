package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	// database string connection
	StringConnection = ""
	// api port
	Port             = 0
)

func LoadDotEnv() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}
	StringConnection = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
}
