package config

type RabbitMqConf struct {
	Host     string `json:",optional"`
	Port     string `json:",optional"`
	Username string `json:",optional"`
	Password string `json:",optional"`
}

type Config struct {
	RabbitMq RabbitMqConf
}
