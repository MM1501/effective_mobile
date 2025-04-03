# Effective Mobile REST API

Простой REST API сервис для управления данными о людях с обогащением информации из внешних API.

## ⚙️ Установка

1.  **Предварительные требования:**

    Убедитесь, что у вас установлено следующее:

    *   Go 1.21+
    *   PostgreSQL 12+

2.  **Клонируйте репозиторий:**

    ```bash
    git clone https://github.com/ваш-логин/effective-mobile.git
    cd effective-mobile
    ```

    *Замените `ваш-логин` на ваше имя пользователя на GitHub.*

3.  **Установите зависимости:**

    ```bash
    go mod download
    ```

4.  **Настройте базу данных (PostgreSQL):**

    Выполните следующие SQL-команды в PostgreSQL:

    ```sql
    CREATE DATABASE effective_mobile;
    \c effective_mobile
    CREATE TABLE people (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255),
        surname VARCHAR(255),
        patronymic VARCHAR(255),
        age INT,
        gender VARCHAR(10),
        nationality VARCHAR(50)
    );
    ```

5.  **Создайте файл `.env`:**

    Создайте файл `.env` в корневой директории проекта и добавьте строку подключения к базе данных:

    ```ini
    DB_CONNECTION_STRING=host=localhost port=5432 user=ваш_пользователь dbname=effective_mobile password=ваш_пароль sslmode=disable
    ```

    *Замените `ваш_пользователь` и `ваш_пароль` на соответствующие значения.*  Обязательно сохраните этот файл вне системы контроля версий (добавьте его в `.gitignore`).

## 🚀 Запуск

1.  **Сгенерируйте Swagger-документацию:**

    ```bash
    swag init
    ```

2.  **Запустите сервер:**

    ```bash
    go run main.go
    ```

    API будет доступно по адресу `http://localhost:8080`.

## 📚 Эндпоинты

*   `POST /people` - Добавить человека.
*   `GET /people` - Получить список людей (можно фильтровать по `age`, `gender`, `nationality`).
*   `PUT /people/:id` - Обновить данные человека (по `id`).
*   `DELETE /people/:id` - Удалить запись о человеке (по `id`).

    **Документация Swagger:** `http://localhost:8080/swagger/index.html`

## 📝 Пример запроса (POST /people)

```json
{
    "name": "Ivan",
    "surname": "Petrov",
    "patronymic": "Sergeevich"
}
