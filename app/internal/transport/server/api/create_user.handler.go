package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

const BillingHost = "http://billing-service"

// CreateUserHandler handles request to create user and account in billing service
func CreateUserHandler() func(c *gin.Context) {
	return func(c *gin.Context) {

		userId := uuid.New()
		endpoint := fmt.Sprintf("%s/account", BillingHost)
		data := map[string]interface{}{
			"user_id": userId,
		}
		body, _ := json.Marshal(data)

		request, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		client := &http.Client{}
		response, err := client.Do(request)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("Problem with creating account. Error = %s", err.Error()),
				"data":    gin.H{},
			})
			return
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err.Error())
			}
		}(response.Body)

		if response.StatusCode != http.StatusOK {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("Problem with creating account. Status = %v", response.Status),
				"data":    gin.H{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data": gin.H{
				"user_id": userId,
			},
		})
	}
}
