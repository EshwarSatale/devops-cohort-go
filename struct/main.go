package main

import (
	"fmt"
	"github.com/fatih/color"
)

type Service struct {
	Name string
	Port int
	Healthy bool
}

func printStatus(s Service){
	status := "Healthy"

	if s.Healthy == false {
		status = "Unhealthy"
		msg := fmt.Sprintf("Name: %s  |  Port: %d  |  %s \n", s.Name, s.Port, status)
		color.Red(msg)
	}else {
		msg := fmt.Sprintf("Name: %s  |  Port: %d  |  %s \n", s.Name, s.Port, status)
		color.HiGreen(msg)
	}
}

func main(){
	fmt.Println("Learning Struct")
	services := []Service{
		{Name:"gateway", Port: 8080, Healthy: true},
		{Name:"postgres", Port: 5432, Healthy: false},
		{Name:"frontend", Port: 443, Healthy: true},
	}
	
	for _, s:=range services {
		printStatus(s)
	}
}
