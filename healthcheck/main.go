package main

import (

 "github.com/fatih/color"
 "fmt"
)

func main(){

	color.Green("application is healthy")
        color.Red("worker service is down")
        fmt.Println("plain output")
}
