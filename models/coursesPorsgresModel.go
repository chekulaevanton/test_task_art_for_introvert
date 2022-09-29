package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/chekulaevanton/test_task_art_for_introvert/utils"
	_ "github.com/lib/pq"
)

type CoursesPostgresModel struct {
    db *sql.DB
}

func NewCoursesPostgresModel(postgresDB *sql.DB) *CoursesPostgresModel {
    return &CoursesPostgresModel{
        db: postgresDB,
    }
}

func (m *CoursesPostgresModel) CheckCourse(id int) (bool, error) {
    if id <= 0 {
        return false, nil
    }

    rows, err := m.db.Query("SELECT id FROM courses WHERE id=$1", id)
    if err != nil {
        return false, err
    }
    defer rows.Close()

    if rows.Next() {
        return true, nil
    }

    return false, nil
}

func (m *CoursesPostgresModel) CreateCourse(id int, name string, priceUsd int, priceRub int) error {
    exists, err := m.CheckCourse(id)
    if err != nil {
        return err
    }
    if exists {
        return utils.NewAlreadyExistsError(id)
    }

    err = m.checkValues(id, name, priceUsd, priceRub)
    if err != nil {
        return err
    }

    _, err = m.db.Exec("INSERT INTO courses (id, name, price_usd, price_rub) VALUES ($1, $2, $3, $4);", id, name, priceUsd, priceRub)
    if err != nil {
        return err
    }

    return nil
}

func (m *CoursesPostgresModel) EditCourse(id int, name string, priceUsd int, priceRub int) error {
    exists, err := m.CheckCourse(id)
    if err != nil {
        return err
    }
    if !exists {
        return utils.NewNotFoundError(id)
    }

    var params []string
    if name != "" {
        params = append(params, fmt.Sprintf("name = '%s'", name))
    }
    if priceUsd > 0 {
        params = append(params, fmt.Sprintf("price_usd = %v", priceUsd))
    } else if priceUsd < 0 {
        return utils.NewValueError("Price USD must be non-negative integer")
    }
    if priceRub > 0 {
        params = append(params, fmt.Sprintf("price_rub = %v", priceRub))
    } else if priceRub < 0{
        return utils.NewValueError("Price RUB must be non-negative integer")
    }

    if len(params) == 0 {
        return nil
    }

    _, err = m.db.Exec("UPDATE courses SET " + strings.Join(params, ", ") + "WHERE id = $1;", id)
    if err != nil {
        return err
    }

    return nil
}

func (m *CoursesPostgresModel) DeleteCourse(id int) error {
    exists, err := m.CheckCourse(id)
    if err != nil {
        return err
    }
    if !exists {
        return utils.NewNotFoundError(id)
    }

    _, err = m.db.Exec("DELETE FROM courses WHERE id = $1", id)
    if err != nil {
        return err
    }

    return nil
}

func (m *CoursesPostgresModel) checkValues(id int, name string, priceUsd int, priceRub int) error {
    if id <= 0 {
        return utils.NewValueError("Id must be non-negative integer")
    }

    if name == "" {
        return utils.NewValueError("Name must not be empty")
    }

    if priceUsd <= 0 {
        return utils.NewValueError("Price USD must be non-negative integer")
    }

    if priceRub <= 0 {
        return utils.NewValueError("Price RUB must be non-negative integer")
    }

    return nil
}

func (m *CoursesPostgresModel) GetCourse(id int) (Course, error) {
    rows, err := m.db.Query("SELECT id, name, price_usd, price_rub FROM courses WHERE id = $1", id)
    if err != nil {
        return Course{}, err
    }

    if rows.Next() {
        var course Course
        rows.Scan(&course.Id, &course.Name, &course.PriceUsd, &course.PriceRub)
        return course, nil

    } else {
        return Course{}, utils.NewNotFoundError(id)
    }
}

func (m *CoursesPostgresModel) GetAllCourses() ([]Course, error) {
    rows, err := m.db.Query("SELECT id, name, price_usd, price_rub FROM courses")
    if err != nil {
        return make([]Course, 0), err
    }

    courses := make([]Course, 0)
    for rows.Next() {
        var course Course
        rows.Scan(&course.Id, &course.Name, &course.PriceUsd, &course.PriceRub)
        courses = append(courses, course)
    }

    return courses, nil
}
