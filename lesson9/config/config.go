package config

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type AppConfig1 struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Image       string   `json:"image"`
	MountDir    string   `json:"mount_dir"`
	Website     string   `json:"website"`
	Repository  string   `json:"repository"`
}

type AppConfig2 struct {
	Adress1_url string
	Adress2_url string
	Debug       bool
	Port        int
	User        string
	Users       []string
	Rate        float32
	Timeout     time.Duration
}

func ReadConfig() (conf1 AppConfig1, conf2 AppConfig2) {

	// Создаем файловый дескриптор для файла, в котором хрнаится json-конфигурация
	file1, err := os.Open("conf1.json")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	// Не забываем закрыть файл при выходе из функции
	defer func() {
		err := file1.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	// Задаем переменную, в которую будем считывать конфиг
	conf1 = AppConfig1{}

	// Задаем декодер из файла и сразу же вызываем его
	err = json.NewDecoder(file1).Decode(&conf1)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	_, err = url.Parse(conf1.Website)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = url.Parse(conf1.Repository)
	if err != nil {
		log.Fatal(err.Error())
	}

	//////////////////////////

	file2, err := os.ReadFile("./conf2.yaml")
	if err != nil {
		log.Fatalf("Не могу прочитать файл: %v", err)
	}

	conf2 = AppConfig2{}

	err = yaml.Unmarshal(file2, &conf2)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = url.Parse(conf1.Website)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = url.Parse(conf1.Repository)
	if err != nil {
		log.Fatal(err.Error())
	}

	// d, err := yaml.Marshal(&conf2)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- t dump:\n%s\n\n", string(d))

	// m := make(map[interface{}]interface{})

	// err = yaml.Unmarshal([]byte(data), &m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m:\n%v\n\n", m)

	// d, err = yaml.Marshal(&m)
	// if err != nil {
	// 	log.Fatalf("error: %v", err)
	// }
	// fmt.Printf("--- m dump:\n%s\n\n", string(d))

	return
}
