package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("snet-ecosystem-contracts/snet/contracts/resources/abi/Registry.json")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(file))
}
