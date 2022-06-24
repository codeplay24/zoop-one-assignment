package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"zoop/one/models"

	"github.com/gin-gonic/gin"
)

var EndpointList []models.Endpoint

func GetData(c *gin.Context) {
	for _, v := range EndpointList {
		fmt.Println(v.Url)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func RegisterEndPoints(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	var jsonResp models.JsonResponse
	err = json.Unmarshal([]byte(jsonData), &jsonResp)
	if err != nil {
		fmt.Println("failed while getting json data")
	}

	EndpointList = append(EndpointList, models.Endpoint{
		Url: jsonResp.Url,
	})
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
