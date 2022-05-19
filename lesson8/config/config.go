package config

import (
	"log"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Address    string
	Address2   string
	Debug      bool
	Port       int
	User       string
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
}

func ReadConfig() Specification {
	var s Specification
	err1 := envconfig.Process("myapp", &s)
	if err1 != nil {
		log.Fatal(err1.Error())
	}

	_, err2 := url.Parse(s.Address)
	if err2 != nil {
		log.Fatal(err2.Error())
	}

	_, err3 := url.Parse(s.Address2)
	if err3 != nil {
		log.Fatal(err3.Error())
	}

	// address := os.Getenv("MYAPP_ADDRESS")
	// fmt.Println(address)
	// if address == "" {
	// 	log.Fatal("Adress is not filled")
	// }

	return s
}
