package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user-api/cmd/api/repository"
)

type UserRequestBody struct {
	Id string `json:"id"`
}

type UserSuccessResponse struct {
	Data interface{} `json:"data,omitempty"`
}

// var ctx = context.Background()

func GetMemberById(c *gin.Context) {
	var requestBody UserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Id == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "JSON must contain a phone field",
		})
		return
	}

	var u gin.H
	var userRepo, err = repository.GetMemberById(requestBody.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.Unmarshal([]byte(userRepo), &u)

	c.JSON(200, UserSuccessResponse{
		Data: u,
	})
}
