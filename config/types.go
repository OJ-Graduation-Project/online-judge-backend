package config

type ApplicationConfig struct {
	Server   ServerConfig   `bson:"server,omitempty" json:"server,omitempty"`
	Mongo    MongoConfig    `bson:"mongo,omitempty" json:"mongo,omitempty"`
	Redis    RedisConfig    `bson:"redis,omitempty" json:"redis,omitempty"`
	Frontend FrontendConfig `bson:"frontend,omitempty" json:"frontend,omitempty"`
}

type FrontendConfig struct {
	Host string `bson:"host,omitempty" json:"host,omitempty"`
	Port string `bson:"port,omitempty" json:"port,omitempty"`
}

type ServerConfig struct {
	Host string `bson:"host,omitempty" json:"host,omitempty"`
	Port string `bson:"port,omitempty" json:"port,omitempty"`
}

type MongoConfig struct {
	Host string `bson:"host,omitempty" json:"host,omitempty"`
	Port string `bson:"port,omitempty" json:"port,omitempty"`
}

type RedisConfig struct {
	Host string `bson:"host,omitempty" json:"host,omitempty"`
	Port string `bson:"port,omitempty" json:"port,omitempty"`
}
