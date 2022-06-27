package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"zoop/one/models"
	"zoop/one/services"

	"github.com/gin-gonic/gin"
)

//EndpointList list downs every endpoints that we register through the /urls/register route
var EndpointList []*models.Endpoint

//HealthyEndpoints puts endpoints which are healthy from the EndpointList slice
var HealthyEndpoints []*models.Endpoint

// this fuction is responosible for forwarding the requst to the endpoint that receives priorty with
//round robin algorithm
func GetData(c *gin.Context) {
	var respBody []byte
	index := 0

	//we will loop through the healthy endpoints and send request to a end point with
	//round robin algorithm

	for index = 0; index < len(HealthyEndpoints); index++ {
		if HealthyEndpoints[index].Blocked {
			continue
		}
		client := &http.Client{}

		req, err := http.NewRequest(c.Request.Method, HealthyEndpoints[index].Url, c.Request.Body)
		if err != nil {
			fmt.Println("error")
		}

		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("Failed while making a request to the", HealthyEndpoints[index].Url)
			continue
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Failed while reading body of the response of", HealthyEndpoints[index].Url)
			continue
		}

		respBody = body
		fmt.Println("sending response from", HealthyEndpoints[index].Url)
		HealthyEndpoints[index].Blocked = true
		break
	}
	if len(HealthyEndpoints)-1 == index {
		for i := range HealthyEndpoints {
			(*HealthyEndpoints[i]).Blocked = false
		}
	}
	//payload := strings.NewReader(string(bodyData))
	if respBody == nil {
		c.JSON(503, gin.H{})
		return
	}
	c.JSON(200, string(respBody))

}

//RegisterEndPoints registers endpoint. this fuction takes a endpoint from the request body and
//puts it in the EndpointList slice
func RegisterEndPoints(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("failed while reading data from json body")
		c.JSON(500, gin.H{"status": "failed"})
		return
	}

	var jsonResp models.JsonResponse
	err = json.Unmarshal([]byte(jsonData), &jsonResp)
	if err != nil {
		fmt.Println("failed while getting json data")
		c.JSON(500, gin.H{"status": "failed"})
		return
	}

	ep := models.Endpoint{
		Url: jsonResp.Url,
	}

	EndpointList = append(EndpointList, &ep)
	go services.CheckHealthFirstTime(&EndpointList, &ep)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
