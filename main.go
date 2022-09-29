package main

import (
    "log"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Print("No .env file found")
    }
}

func main() {
    log.Print("IT'S ALIVE!")

    config := config.NewConfig()

}
