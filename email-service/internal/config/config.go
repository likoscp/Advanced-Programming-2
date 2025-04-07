package config

import (
    "log"
    "os"
    "strconv"
)

type Config struct {
    Addr        string
    MongoURI    string
    DBName      string
    Collection  string
    SMTPHost    string
    SMTPPort    int
    SMTPEmail   string
    SMTPPassword string
}

func LoadConfig() (*Config, error) {
    addr := os.Getenv("ADDR")
    mongoURI := os.Getenv("MONGO")
    dbName := os.Getenv("DBname")
    collection := os.Getenv("COLLECTION")
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPortStr := os.Getenv("SMTP_PORT")
    smtpEmail := os.Getenv("SMTP_EMAIL")
    smtpPassword := os.Getenv("SMTP_PASSWORD")

    if addr == "" || mongoURI == "" || dbName == "" || collection == "" || smtpHost == "" || smtpPortStr == "" || smtpEmail == "" || smtpPassword == "" {
        log.Fatal("Missing required environment variables")
    }

    smtpPort, err := strconv.Atoi(smtpPortStr)
    if err != nil {
        log.Fatalf("Invalid SMTP_PORT value: %v", err)
    }

    return &Config{
        Addr:        addr,
        MongoURI:    mongoURI,
        DBName:      dbName,
        Collection:  collection,
        SMTPHost:    smtpHost,
        SMTPPort:    smtpPort,
        SMTPEmail:   smtpEmail,
        SMTPPassword: smtpPassword,
    }, nil
}