package env

import (
	"github.com/amiralii/goEchoExample/config/response"

	"github.com/joho/godotenv"
)

func Init() (err error) {

	err = godotenv.Load()
	if err != nil {
		return response.Error("env not found or not readable", 1001)
	}
	return
}
