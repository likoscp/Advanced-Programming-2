package config

import "github.com/joho/godotenv"

type Config struct {
	Addr            string
	MongoUri        string
	DBname          string
	Collection      string
	SECRET          string
	AdminCollection string
	MailURI         string
}

func NewConfig() (*Config, error) {
	c, err := godotenv.Read(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
			Addr:            c["ADDR"],
			MongoUri:        c["MONGO"],
			DBname:          c["DBname"],
			Collection:      c["COLLECTION"],
			SECRET:          c["SECRET"],
			AdminCollection: c["ADMIN_COLLECTION"],
			MailURI:         c["MAIL_URI"]},

		nil
}
