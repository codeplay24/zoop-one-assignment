package services

import (
	"fmt"
	"net/http"
	"time"
	"zoop/one/models"
)

//this function will filter out unhealty endpoints
func FilterHealtyEndpoints(urlsPtr *[]*models.Endpoint, healthyEndpoints *[]*models.Endpoint) {
	for _, v := range *urlsPtr {
		prevStatuses := v.PreviousStatuses
		successfulStatusCount := 0

		for i := len(prevStatuses) - 1; i >= 0; i-- {
			statusCode := prevStatuses[i].StatusCode
			timeOfTheStatus := prevStatuses[i].Time

			if statusCode == 200 {
				successfulStatusCount++
			}

			if successfulStatusCount == 3 {
				diffInSeconds := time.Since(timeOfTheStatus).Seconds()
				if diffInSeconds < 15 {
					// if the request alredy exits in the healthy end point
					flag := false
					for _, Ep := range *healthyEndpoints {
						if v.Url == Ep.Url {
							flag = true
							break
						}
					}
					if !flag {
						*healthyEndpoints = append(*healthyEndpoints, v)
					}

					break
				} else {
					for j, Ep := range *healthyEndpoints {
						if v.Url == Ep.Url {
							*healthyEndpoints = append((*healthyEndpoints)[:j], (*healthyEndpoints)[j+1:]...)
							break
						}
					}

					break
				}
			}
		}
	}
}

// this function checks healt when ever it has time
func GetHealth(urlsPtr *[]*models.Endpoint, healtyEndpoints *[]*models.Endpoint) {
	flag := false
	for {
		for i, v := range *urlsPtr {
			if !v.IsReadyToServe {
				continue
			}
			flag = true
			statusCode, err := MakeARequest(v.Url)
			if err != nil {
				*urlsPtr = (*urlsPtr)[0 : len(*urlsPtr)-1]
				continue
			}

			(*urlsPtr)[i].PreviousStatuses = append((*urlsPtr)[i].PreviousStatuses, models.Status{
				Time:       time.Now(),
				StatusCode: statusCode,
			})
		}
		if flag {
			FilterHealtyEndpoints(urlsPtr, healtyEndpoints)
		}
		time.Sleep(1 * time.Second)
	}
}

// this function checks health for the first time when registering the url
func CheckHealthFirstTime(urlsPtr *[]*models.Endpoint, endpoint *models.Endpoint) {
	t := time.Now()
	for {
		diffInSeconds := time.Since(t).Seconds()
		if diffInSeconds >= 15 {
			(*endpoint).IsReadyToServe = true
			break
		}

		statusCode, err := MakeARequest((*endpoint).Url)
		if err != nil {
			(*urlsPtr)[len(*urlsPtr)-1] = nil
			break
		}

		//fmt.Println(statusCode, (*urlsPtr)[len(*urlsPtr)-1].Url, "request sent from first time checking health")
		//fmt.Println()

		(*endpoint).PreviousStatuses = append((*endpoint).PreviousStatuses, models.Status{
			Time:       time.Now(),
			StatusCode: statusCode,
		})

	}
	//FilterHealtyEndpoints(urlsPtr, healtyEndpoints)

}

//this function makes request to the url
func MakeARequest(url string) (int, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	defer res.Body.Close()

	return res.StatusCode, nil
}
