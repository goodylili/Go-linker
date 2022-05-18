package models

import (
	"crypto/rand"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func (LinkMethod *LinkModel) GenerateIdentifier() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("There's an environment variable error", err)
	}
	chars := os.Getenv("CHARACTERS")
	bytes := make([]byte, 10) // makes an array of bytes'=O
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("There was a problem with rand read bytes")
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	LinkMethod.Identifier = string(bytes)
}
