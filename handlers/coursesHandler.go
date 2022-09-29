package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	c "github.com/chekulaevanton/test_task_art_for_introvert/controllers"
	"github.com/chekulaevanton/test_task_art_for_introvert/models"
	"github.com/chekulaevanton/test_task_art_for_introvert/utils"
)

type CoursesHandler struct {
    controller *c.CourseController
}

func NewCoursesHandler(controller *c.CourseController) *CoursesHandler {
    return &CoursesHandler{
        controller: controller,
    }
}

func (h *CoursesHandler) HandleCourse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    switch (r.Method) {
    case http.MethodPut:
        log.Print("Handle request PUT /course/id")
        // Парсим url
        url := utils.GetPartsFromURL(r.URL.Path)
        if (len(url) < 2) {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(utils.GetErrorJSONResponse("You must specify id")))
            return
        }
        // Вытаскиваем id из url запроса
        id, err := strconv.Atoi(url[1])
        if err != nil || id < 0 {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(utils.GetErrorJSONResponse("Id must be non negative integer")))
            return
        }

        // Парсим тело запроса ради параметров курса
        var params models.Course
        err = json.NewDecoder(r.Body).Decode(&params)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(utils.GetErrorJSONResponse("JSON request can not be decoded")))
            return
        }

        // Изменяем или создаем курс
        err = h.controller.CreateEditCourse(id, params.Name, params.PriceUsd, params.PriceRub)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(utils.GetErrorJSONResponse(err.Error())))
            return
        }


        w.WriteHeader(http.StatusCreated)
        w.Write([]byte("{}"))
        return

    case http.MethodDelete:
        log.Print("Handle request DELETE /course/id")

        // Парсим url
        url := utils.GetPartsFromURL(r.URL.Path)
        if (len(url) < 2) {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(utils.GetErrorJSONResponse("You must specify id")))
            return
        }
        // Вытаскиваем id из url запроса
        id, err := strconv.Atoi(url[1])
        if err != nil || id < 0 {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(utils.GetErrorJSONResponse("Id must be non negative integer")))
            return
        }

        // Удаляем курс по id
        err = h.controller.DeleteCourse(id)
        if err != nil {
            if _, ok := err.(utils.NotFoundError); ok {
                w.WriteHeader(http.StatusNotFound)
                w.Write([]byte(utils.GetErrorJSONResponse("Course with id=" + strconv.Itoa(id) + " not found")))
                return

            } else {
                w.WriteHeader(http.StatusInternalServerError)
                w.Write([]byte(utils.GetErrorJSONResponse(err.Error())))
                return
            }
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("{}"))

    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(utils.GetErrorJSONResponse("Method " + r.Method + " is not implemented")))
        return
    }
}

func (h *CoursesHandler) HandleCourses(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    switch (r.Method) {
    case http.MethodGet:
        log.Print("Handle request GET /courses")

        courses := h.controller.GetAllCourses()
        response, err := json.Marshal(courses)
        if err != nil {
            w.Write([]byte(utils.GetErrorJSONResponse(err.Error())))
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write(response)
        return

    default:
        w.WriteHeader(http.StatusNotImplemented)
        w.Write([]byte(utils.GetErrorJSONResponse("Method " + r.Method + " is not implemented")))
        return
    }
}
