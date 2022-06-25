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

var EndpointList []models.Endpoint
var HealthyEndpoints []models.Endpoint

func GetData(c *gin.Context) {
	// do the robinson algo on healthy end points
	var respBody []byte
	fmt.Println(len(HealthyEndpoints))
	for i := 0; i < len(HealthyEndpoints); i++ {
		if HealthyEndpoints[i].Blocked {
			continue
		}
		client := &http.Client{}

		req, err := http.NewRequest(c.Request.Method, HealthyEndpoints[i].Url, c.Request.Body)
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
		HealthyEndpoints[i].Blocked = true
		//fmt.Println(v.Url, HealthyEndpoints[i].Blocked, "wtf boro")
		break
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

	EndpointList = append(EndpointList, models.Endpoint{
		Url: jsonResp.Url,
	})
	go services.CheckHealthFirstTime(&EndpointList, &HealthyEndpoints)

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
