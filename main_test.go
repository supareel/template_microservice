package main

import (
	"gomicro/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_USERNAME", "root")
  os.Setenv("DB_PASS","secret")
  os.Setenv("DB_SERVER", "localhost")
  os.Setenv("DB_PORT", "5432")
  os.Setenv("DB_NAME", "simple_bank")
  os.Setenv("DB_SSL", "disable")
	
	database.ConnectToDB()
}