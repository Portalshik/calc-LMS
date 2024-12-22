# Calculator Web Service

Это веб-сервис для вычисления арифметических выражений. Сервис принимает арифметическое выражение через HTTP POST запрос, вычисляет результат и возвращает его в ответе.
Описание

Сервис реализует один endpoint: /api/v1/calculate. Он принимает POST запрос с телом, содержащим арифметическое выражение, и возвращает результат вычисления этого выражения или ошибку, если выражение некорректно.
API
URL:

## POST /api/v1/calculate

Тело запроса:
```
{
    "expression": "выражение"
}
```

где "expression" — это строка с арифметическим выражением, которое нужно вычислить.
Ответ:
1. Успешный ответ (код 200):

```
{
    "result": "результат"
}
```

где "result" — результат вычисления выражения.
2. Ошибка 422 (Неверное выражение):

Если выражение не соответствует требованиям приложения (например, содержит недопустимые символы, такие как буквы), сервер возвращает:
```
{
    "error": "Expression is not valid"
}
```
3. Ошибка 500 (Внутренняя ошибка сервера):

В случае непредвиденной ошибки (например, ошибка в процессе вычисления или обработке запроса):
```
{
    "error": "Internal server error"
}
```

Пример использования

    Пример успешного запроса: Используя curl, отправляем запрос на сервер:
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
Ответ:
```
{
    "result": "6.000000"
}
```
Пример запроса с ошибкой 422 (некорректное выражение):
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2a"
}'
```
Ответ:
```
{
    "error": "Expression is not valid"
}
```
Пример ошибки 500: Это может произойти при внутренних сбоях, например, ошибке при вычислении:
```
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1/0"
}'
```
Ответ:
```
{
    "error": "Internal server error"
}
```
Запуск проекта

    Склонируйте репозиторий:
```
git clone https://github.com/Portalshik/calc-LMS.git
cd calc-lms
```
Убедитесь, что у вас установлен Go версии 1.18 и выше.

Запустите сервер с помощью команды:
- Для Linux:
    ```
    sudo go run cmd/calc/main.go 
    ```
- Для Windows:
    ```
    go run cmd/calc/main.go 
    ```
Сервер будет доступен по адресу localhost:8080. Вы можете отправлять запросы на /api/v1/calculate.
Требования

    Go версии 1.18 и выше.
    Установленный пакет curl для отправки запросов (или использование других HTTP клиентов).

Структура проекта
```
calc-lms/
├── cmd
│   └── calc
│       └── main.go
├── go.mod
├── internal
│   ├── api
│   │   └── v1
│   │       └── api.go
│   ├── calculator
│   │   └── calculator.go
│   └── server
│       └── server.go
└── README.md
```