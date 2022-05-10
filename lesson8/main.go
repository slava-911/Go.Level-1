package main

import (
	"GeekBrains/Go.Level-1/lesson8/config"
	"fmt"
)

func main() {
	sp := config.ReadConfig()
	format := "Address: %s\nAddress2: %s\nDebug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\n"
	_, err := fmt.Printf(format, sp.Address, sp.Address2, sp.Debug, sp.Port, sp.User, sp.Rate, sp.Timeout)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range sp.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range sp.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
}
