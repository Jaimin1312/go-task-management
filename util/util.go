package util

import (
	"crypto/rand"
	"fmt"

	"github.com/sirupsen/logrus"
)

func RecoverPanic() {
	if r := recover(); r != nil {
		// Handle the panic here
		logrus.Info("Recovered from go routine panic:", r)
	}
}

func SetResponse(data interface{}, status int, message string) map[string]interface{} {
	response := make(map[string]interface{})
	response["data"] = nil
	if data != nil {
		response["data"] = data
	}
	response["status"] = status
	response["message"] = message
	return response
}

func SetPaginationResponse(data interface{}, total, status int, message string) map[string]interface{} {
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"info":  []int{},
			"total": 0,
		},
		"status":  status,
		"message": message,
	}
	if data != nil {
		response["data"].(map[string]interface{})["info"] = data
		response["data"].(map[string]interface{})["total"] = total
	}
	return response
}

// ID a unique identifier
type ID []byte

// NewID generate a new ID
func NewID() ID {
	ret := make(ID, 20)
	if _, err := rand.Read(ret); err != nil {
		fmt.Println(err)
	}
	return ret
}

// Response represents a standard API response structure.
type Response struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}
