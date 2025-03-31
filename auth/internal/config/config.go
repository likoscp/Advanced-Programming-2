package config

import "github.com/joho/godotenv"

type Config struct {
	Addr     string
	MongoUri string
	DBname   string
	CollectionName string
}

func NewConfig() (*Config, error) {
	c, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		Addr: c["ADDR"],
		MongoUri: c["MONGO"],
		DBname:   c["DBname"],
		CollectionName: c["COLLECTION"],}, 
		nil
}
