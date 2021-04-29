package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"ip_changer/internal/app/api"
	"ip_changer/internal/app/notifier"
	"log"
)

func main() {
	config := api.NewConfig()
	yamlFile, err := ioutil.ReadFile("data/config.yml")
	if err != nil {
		log.Println("Error while reading yaml file: ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Println("Error while decoding yaml file: ", err)
	}

	sc := notifier.SlackClient{
		WebHookUrl: config.SlackWebhook,
		UserName: config.SlackUsername,
		Channel: config.SlackChannel,
	}
	config.SlackClient = sc

	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}
