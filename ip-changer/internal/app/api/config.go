package api

import "ip_changer/internal/app/notifier"

type Config struct {
	SlackWebhook string `yaml:"slack_webhook"`
	SlackUsername string `yaml:"slack_username"`
	SlackChannel string `yaml:"slack_channel"`
	Port string `yaml:"port"`
	LogLevel string `yaml:"log_level"`
	DockerNetwork string `yaml:"docker_network"`
	SlackClient notifier.SlackClient
	Token string `yaml:"token"`
}

func NewConfig() *Config {
	return &Config{
		Port: "9999",
		LogLevel: "debug",
	}
}