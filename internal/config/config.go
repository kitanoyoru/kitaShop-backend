package config

import (
  "os"
  "time"
)

type Config struct {
  Env string
  Mongo MongoConfig
  HTTP HTTPConfig
}

type MongoConfig struct {
  URI string
  User string
  password string
  DatabaseName string
}

type HTTPConfig struct {
  Host string
  Port string
  ReadTimeout time.Duration
  WriteTimeout time.Duration
}

// TODO: Implement unmarshal config from yml 
