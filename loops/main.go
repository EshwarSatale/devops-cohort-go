package main

import "fmt"

func main(){
	fmt.Println("learnign loops")

	//classic loop
	for i:=0; i<=3; i++ {
		fmt.Println(i)
	}
	
	var servers []string
	servers = []string{"server1", "server2", "server3"}

	//print using classic loop
	fmt.Println("Print using classic loop")
	for i:=0; i < len(servers); i++ {
		fmt.Println("index:", i, "server:", servers[i])
	}

	fmt.Println("Print using range:")

	for i, s := range servers {
		fmt.Printf("[%d] => %s\n",i,s)
	}

	//skip index
	for _, s:=range servers {
		fmt.Println(s)
	}

	fmt.Println("Using while")
	retries := 0
	for retries<5 {
		if retries == 3{
			retries++
			continue
		}
		fmt.Println(retries)
		retries++
	}
}
