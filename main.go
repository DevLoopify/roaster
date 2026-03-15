package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {

}

func GetEnvVariable(variable string) string {
	godotenv.Load()
	result := os.Getenv(variable)
	return result
}
