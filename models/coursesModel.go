package models

type CoursesModel interface {
    CheckCourse(id int) (bool, error)
    CreateCourse(id int, name string, priceUsd int, priceRub int) error
    EditCourse(id int, name string, priceUsd int, priceRub int) error
    DeleteCourse(id int) error
    GetCourse(id int) (Course, error)
    GetAllCourses() ([]Course, error)
}

type Course struct {
    Id        int    `json:"id"`
    Name      string `json:"name"`
    PriceUsd  int    `json:"price_usd"`
    PriceRub  int    `json:"price_rub"`
}

const (
    StatusCreated = 0
    StatusDeleted = 1
)
