
package main


import (
	"fmt"
	"github.com/EshwarSatale/devops-cohort-go/healthcheck/models"
	"github.com/EshwarSatale/devops-cohort-go/healthcheck/checker"
)

func main() {

	fmt.Println("Devops Healthcheck!")

	services := []models.Service{
   	   {Name: "gateway",  Port: 8080  , Healthy: true },
    	   {Name: "postgres",  Port: 5432 , Healthy: false },
    	   {Name: "frontend",  Port:  443 , Healthy:  true },
       }

	for _, svc  := range services {

		checker.PrintStatus(svc)
	}
}
