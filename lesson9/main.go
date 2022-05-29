package main

import (
	"GeekBrains/Go.Level-1/lesson9/config"
	"fmt"
)

func main() {

	confJSON, confYAML := config.ReadConfig()

	fmt.Printf("--- confJSON:\n%v\n\n", confJSON)

	fmt.Printf("--- confYAML:\n%v\n\n", confYAML)

}
