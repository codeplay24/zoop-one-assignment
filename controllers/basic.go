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

var EndpointList []*models.Endpoint
var HealthyEndpoints []*models.Endpoint

func GetData(c *gin.Context) {
	// do the robinson algo on healthy end points
	var respBody []byte
	fmt.Println(len(HealthyEndpoints))
	index := 0
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
			fmt.Println(err)
			continue
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		respBody = body
		fmt.Println("sending response from", (*HealthyEndpoints[index]).Url)
		HealthyEndpoints[index].Blocked = true
		//fmt.Println(v.Url, HealthyEndpoints[i].Blocked, "wtf boro")
		break
	}
	if len(HealthyEndpoints)-1 == index {
		for i := range HealthyEndpoints {
			(*HealthyEndpoints[i]).Blocked = false
		}
	}
	//payload := strings.NewReader(string(bodyData))
	c.JSON(http.StatusOK, respBody)

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

	ep := models.Endpoint{
		Url: jsonResp.Url,
	}

	EndpointList = append(EndpointList, &ep)
	go services.CheckHealthFirstTime(&EndpointList, &ep)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func Getval(c *gin.Context) {
	for i := range EndpointList {
		fmt.Println(EndpointList[i].IsReadyToServe, EndpointList[i].Url)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}
