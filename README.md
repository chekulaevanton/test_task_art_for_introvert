# Courses API
Тестовое задание с API для управления БД курсов

API предоставляет методы для работы с списком курсов,
создания и изменения конкретных курсов, 

## Интерфейс API
- GET /courses
    Возвращает список всех доступных курсов.
    Добавленные в данный момент курсы становятся
    доступными через определенное время.

    Возвращаемые данные:
    ```json
        {
            "courses": [
                {
                    "id": 0,
                    "name": "Course Name 1",
                    "price_usd": 99,
                    "price_rub": 5000,
                },
                {
                    "id": 1,
                    "name": "Course Name 2",
                    "price_usd": 50,
                    "price_rub": 2500,
                },
            ],
        }
    ```

- POST /course/id
    Запрос на создание нового курса.
    При использовании существующего id, возвращает
    ошибку "Course already exists".

    Принимаемые параметры:
    ```json
        {
            "name": "Course Name 1",
            "price_usd": 50,
            "price_rub": 2500,
        }
    ```
    В данном запросе каждое поле является обязательным
    для заполнения.


- PUT /course/id
    Запрос на изменение существующего курса.
    При использовании несуществующего id, возвращает
    ошибку "Course does not exist"

    Принимаемые параметры:
    ```json
        {
            "name": "Course Name 1",
            "price_usd": 50,
            "price_rub": 2500,
        }
    ```

    Каждое поле в данном запросе является необязательным.
    Если поле не указано, значение поля остается прежним.

    Возвращаемые значения:
    ```json
        {
            "status": true | false,
            "error": "Error Text",
        }
    ```

- DELETE /course/id
    Удаляет указанный курс. При несуществующем id
    возврщает ошибку "Course does not exist"
    Возвращаемые значения:
    ```json
        {
            "status": true | false,
            "error": "Error Text",
        }
    ```

