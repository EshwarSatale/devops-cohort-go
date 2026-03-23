package main

import "fmt"

type Service struct {
	Name string
	Port int
	Healthy bool
}

func main(){
	fmt.Println("Learning Struct")
	httpService:=Service{Name:"gateway", Port: 8080, Healthy: true}
	fmt.Println(httpService)
	fmt.Printf("%+v\n",httpService)
}
