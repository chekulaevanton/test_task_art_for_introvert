package models

/* TODO: Сделать реализацию модели для MongoDB */

type CoursesMongoModel struct {

}

func NewCoursesMongoModel() *CoursesMongoModel {
    return &CoursesMongoModel{}
}

func (m *CoursesMongoModel) CheckCourse(id int) (bool, error) {
    return false, nil
}

func (m *CoursesMongoModel) CreateCourse(id int, name string, price_usd int, price_rub int) error {
    return nil
}

func (m *CoursesMongoModel) EditCourse(id int, name string, price_usd int, price_rub int) error {
    return nil
}

func (m *CoursesMongoModel) DeleteCourse(id int) error {
    return nil
}

func (m *CoursesMongoModel) GetCourse(id int) (Course, error) {
    return Course{}, nil
}

func (m *CoursesMongoModel) GetAllCourses() ([]Course, error) {
    return make([]Course, 0), nil
}
