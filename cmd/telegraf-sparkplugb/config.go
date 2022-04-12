package main

type config struct {
	Brokers []string `yaml:"brokers"`
	Username string
	Password string
	ClientId string `yaml:"client_id"`
}
