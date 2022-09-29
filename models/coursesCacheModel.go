package models

import (
	"sort"
	"sync"
)

type sortCourse []Course

func (a sortCourse) Len() int           { return len(a) }
func (a sortCourse) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortCourse) Less(i, j int) bool { return a[i].Id < a[j].Id }


type CoursesCacheModel struct {
    sync.Mutex
    courses map[int]Course
    queue map[int]int
}

func NewCoursesCacheModel() *CoursesCacheModel{
    return &CoursesCacheModel{
        courses: make(map[int]Course),
        queue: make(map[int]int),
    }
}


func (m *CoursesCacheModel) GetAllCourses() []Course {
    m.Lock()
    defer m.Unlock()

    courses := make([]Course, 0)
    for _, course := range m.courses {
        courses = append(courses, course)
    }

    sort.Sort(sortCourse(courses))
    return courses
}

func (m *CoursesCacheModel) AddCourse(course Course) {
    m.Lock()
    defer m.Unlock()

    m.courses[course.Id] = course
}

func (m *CoursesCacheModel) DeleteCourse(id int) {
    m.Lock()
    defer m.Unlock()

    delete(m.courses, id)
}

func (m *CoursesCacheModel) AddToQueue(id int, status int) {
    m.Lock()
    defer m.Unlock()

    m.queue[id] = status
}

func (m *CoursesCacheModel) GetQueue() map[int]int {
    m.Lock()
    defer m.Unlock()
    queue := m.queue
    m.queue = make(map[int]int)

    return queue
}
