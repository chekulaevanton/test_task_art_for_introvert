package utils

import "errors"

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

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{
		Err: errors.New(err),
	}
}

func (e NotFoundError) Error() string {
	return e.Err.Error()
}
