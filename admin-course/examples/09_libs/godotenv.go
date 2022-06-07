package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//START_1 OMIT
func main() {
	err := godotenv.Load() // HL1
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET") // HL1
	secretKey := os.Getenv("SECRET_KEY")

	// now do something with s3 or whatever
}

//END_1 OMIT
