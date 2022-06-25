package main

import (
	"zoop/one/controllers"
	"zoop/one/routes"
	"zoop/one/services"
)

func main() {
	r := routes.SetUpRouter()
	go services.GetHealth(&controllers.EndpointList, &controllers.HealthyEndpoints)
	r.Run(":8080")
}
