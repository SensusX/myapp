package main

import (
	"fmt"
	"myapp/internal/config"
)

func main() {
	_, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

}
