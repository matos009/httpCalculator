

HTTP Calculator

HTTP Calculator — это простой веб-сервис для вычисления арифметических выражений, предоставляемых в формате JSON.

📋 Возможности
	•	Принимает арифметические выражения через HTTP POST запрос.
	•	Возвращает результат в формате JSON.
	•	Обрабатывает ошибки (например, некорректный ввод, деление на ноль и т. д.).

🚀 Как запустить проект

Для запуска проекта выполните следующие шаги:
	1.	Клонируйте репозиторий:

git clone https://github.com/matos009/httpCalculator.git
cd httpCalculator


	2.	Убедитесь, что у вас установлен Go (версии 1.20 или выше). Если не установлен, скачайте и установите с официального сайта.
	3.	Запустите сервис:

go run ./cmd/calc_service/...



Сервис будет доступен по адресу http://localhost:8080.

📖 Примеры использования

Успешный запрос

Запрос:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'

Ответ:

{
  "result": 6
}

Ошибка 422: Некорректное выражение

Запрос:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2a"
}'

Ответ:

{
  "error": "Something wrong in expression"
}

Ошибка 500: Деление на ноль

Запрос:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "1/0"
}'

Ответ:

{
  "error": "Something wrong in expression"
}

🛠 Структура проекта

httpCalculator/
├── cmd/
│   └── calc_service/
│       └── main.go         # Точка входа
├── internal/
│   └── calculate/
│       └── calculate.go    # Логика вычислений
├── go.mod                  # Модуль Go
├── go.sum                  # Зависимости
└── README.md               # Документация

🔍 Как работает
	•	Запрос: Пользователь отправляет POST-запрос на /api/v1/calculate с JSON-телом:

{
  "expression": "2+2*2"
}


	•	Ответ: В случае успешного вычисления, сервис возвращает результат:

{
  "result": 6
}


	•	Если ввод некорректен или происходит ошибка, сервис возвращает соответствующий код ошибки и сообщение.

💡 Примечания
	•	Сервис поддерживает базовые арифметические операции: +, -, *, /.
	•	Скобки и десятичные числа также поддерживаются.

Если у вас есть вопросы или предложения, пожалуйста, свяжитесь со мной! 😊

