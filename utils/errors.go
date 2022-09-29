package utils

import (
	"errors"
	"fmt"
)

/* Value Error
 * Описывает состояние, когда поля структуры заданы не верно.
 */
type ValueError struct {
    Err error
}

func NewValueError(err string) ValueError {
    return ValueError{
        Err: errors.New(err),
    }
}

func (e ValueError) Error() string {
    return e.Err.Error()
}

/* Not Found Error
 * Описывает состояние, когда запрошенный ресурс не был найден,
 * эквивалентра HTTP коду 404.
 */
type NotFoundError struct {
    Err error
}

func NewNotFoundError(id int) NotFoundError {
    return NotFoundError{
        Err: fmt.Errorf("Course with id=%v not found", id),
    }
}

func (e NotFoundError) Error() string {
    return e.Err.Error()
}

/* Already Exists Error
 * Описывает состояние, когда запрошенный ресурс требуется создать,
 * но уже существует.
 */
 type AlreadyExistsError struct {
    Err error
}

func NewAlreadyExistsError(id int) AlreadyExistsError {
    return AlreadyExistsError{
        Err: fmt.Errorf("Course with id=%v already exits", id),
    }
}

func (e AlreadyExistsError) Error() string {
    return e.Err.Error()
}
