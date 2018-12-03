package main

import (
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("GO_ENV")

	if "" == env {
		env = "development"
	}
	
	if "test" != env {
		godotenv.Load(".env." + env + ".local")
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}