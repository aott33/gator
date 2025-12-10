package main

import (
	"fmt"

	"github.com/aott33/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Read Error: %v\n", err)
	}

	cfg.SetUser("lane")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Read Error: %v\n", err)
	}	

	fmt.Printf("%+v\n",cfg)
}
