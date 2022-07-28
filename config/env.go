package config

import (
	"log"
	"os"
	"strconv"
)

type EnvConfiguration struct {
	APP_PORT int
  DB_USERNAME string
  DB_PASS string
  DB_SERVER string
  DB_PORT int
  DB_NAME string
  DB_SSL string
  GIN_MODE string
}

var EnvConfig EnvConfiguration

func LoadEnvConfig() {
  appport, err := strconv.ParseInt(os.Getenv("APP_PORT"), 10, 32)
  if err != nil {
    log.Fatalf("\nUnable to parse port value :: %s\nRecieved value :: %s\n", err, os.Getenv("DB_PORT"))
  }
  dbport, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
  if err != nil {
    log.Fatalf("\nUnable to parse port value :: %s\nRecieved value :: %s\n", err, os.Getenv("DB_PORT"))
  }
  // --------------------------------------------------------------------------------------
  // --------------------------------------------------------------------------------------
	EnvConfig.APP_PORT = int(appport)
  EnvConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
  EnvConfig.DB_PASS = os.Getenv("DB_PASS")
  EnvConfig.DB_SERVER = os.Getenv("DB_SERVER")
  
  EnvConfig.DB_PORT = int(dbport)
  EnvConfig.DB_NAME = os.Getenv("DB_NAME")
  EnvConfig.DB_SSL = os.Getenv("DB_SSL")
  EnvConfig.GIN_MODE = os.Getenv("GIN_MODE")
}