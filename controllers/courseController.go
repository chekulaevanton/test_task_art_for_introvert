package controllers

import (
	"log"
	"time"

	"github.com/chekulaevanton/test_task_art_for_introvert/models"
)

type CourseController struct {
    data models.CoursesModel
    cache *models.CoursesCacheModel
}

func NewCoursesController(dataModel models.CoursesModel, cacheModel *models.CoursesCacheModel) *CourseController {
    return &CourseController{
        data: dataModel,
        cache: cacheModel,
    }
}

func (c *CourseController) CreateEditCourse(id int, name string, price_usd int, price_rub int) (err error) {
    exists, err := c.data.CheckCourse(id)
    if err != nil {
        return err
    }

    if exists {
        err = c.data.EditCourse(id, name, price_usd, price_rub)
    } else {
        err = c.data.CreateCourse(id, name, price_usd, price_rub)
    }

    if err != nil {
        return err
    }

    c.cache.AddToQueue(id, models.StatusCreated)
    return nil
}

func (c *CourseController) DeleteCourse(id int) (err error) {
    err = c.data.DeleteCourse(id)
    if err != nil {
        return err
    }

    c.cache.AddToQueue(id, models.StatusDeleted)
    return nil
}

func (c *CourseController) GetAllCourses() []models.Course {
    return c.cache.GetAllCourses()
}

func (c *CourseController) RunCacheSync() {
    go func() {
        courses, err := c.data.GetAllCourses()
        if err != nil {
            log.Print(err)
            return
        }
        for _, course := range courses {
            c.cache.AddCourse(course)
        }


        ticker := time.NewTicker(5 * time.Second)
        for range ticker.C {
            log.Print("Sync started")

            queue := c.cache.GetQueue()

            for id, status := range queue {
                switch status {
                case models.StatusDeleted:
                    c.cache.DeleteCourse(id)

                case models.StatusCreated:
                    course, err := c.data.GetCourse(id)
                    if err != nil {
                        log.Print(err)
                        continue
                    }
                    c.cache.AddCourse(course)
                }
            }
        }
    }()
}
