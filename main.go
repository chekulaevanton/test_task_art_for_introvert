package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/chekulaevanton/test_task_art_for_introvert/config"
	"github.com/chekulaevanton/test_task_art_for_introvert/controllers"
	"github.com/chekulaevanton/test_task_art_for_introvert/handlers"
	"github.com/chekulaevanton/test_task_art_for_introvert/models"
	"github.com/chekulaevanton/test_task_art_for_introvert/server"
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

    db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", config.DB.Username, config.DB.Pass, config.DB.Hostname, config.DB.Name, config.DB.SSLMode))
    if err != nil {
        log.Fatal("No database connection:", err)
    }
    defer db.Close()

    model := models.NewCoursesPostgresModel(db)
    cache := models.NewCoursesCacheModel()

    controller := controllers.NewCoursesController(model, cache)
    controller.RunCacheSync()

    handler := handlers.NewCoursesHandler(controller)

    server := server.NewServer(fmt.Sprintf("%s:%s", config.Hostname, config.Port), handler)
    server.Run()

}
