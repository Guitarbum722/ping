package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guitarbum722/ping"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("please provide a host to ping")
	}

	summary, err := ping.Ping(os.Args[1])
	if err != nil {
		log.Fatalln("err pinging host : ", err.Error())
	}
	fmt.Println(summary)
}
