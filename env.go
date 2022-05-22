package main

import (
	"fmt"
	"os"
)

var CONNECTION_URL string
var DATABASE string
var DB_USERNAME string
var DB_PASSWORD string

func GetOSEnv() {
	DATABASE = os.Getenv("DATABASE")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	
	url := os.Getenv("CONNECTION_URL")
	CONNECTION_URL = fmt.Sprintf(url, DB_USERNAME, DB_PASSWORD)
}
