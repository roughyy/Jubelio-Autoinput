package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getUserCredentials() Credentials {
	return Credentials{
		Email:    os.Getenv("USER_EMAIL"),
		Password: os.Getenv("USER_PASSWORD"),
	}
}

func GetToken() (result string, err error) {
	credentials, err := json.Marshal(getUserCredentials())
	if err != nil {
		return "", err
	}

	resp, err := http.Post("https://api2.jubelio.com/login", "application/json", bytes.NewBuffer(credentials))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var resultMap map[string]interface{}
	
	if err := json.NewDecoder(resp.Body).Decode(&resultMap); err != nil {
		return "", err
	}
	// fmt.Println(resultMap)
	if token, ok := resultMap["token"].(string); ok {
		return token, nil
	}

	return "", fmt.Errorf("no token found in response")
}