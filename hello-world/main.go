package main
import "fmt"
func main(){
	fmt.Println("Hello, DevOps Bootcamp!")
	name := "prometheus"
	port := 9090
	isActive := true

	var service string = "grafana"

	fmt.Printf("%s is hosted on port %d and active status is %t which uses %s for virtualization\n", name, port, isActive, service)
}
