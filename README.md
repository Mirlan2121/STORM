Virgil Fan Site REST API (DMC)
REST API для фанатского сайта Вергилия из Devil May Cry. Этот проект предоставляет информацию о персонаже, его цитаты, галерею изображений и другие данные.

🛠 Технологии
Язык: Go (Golang)

Фреймворк: Gin (HTTP-роутер)

База данных: PostgreSQL (или SQLite для разработки)

Документация API: Swagger (используется swaggo)

📌 Функционал
Получение информации о Вергилии

Список цитат персонажа

Галерея изображений

Возможность добавлять/удалять контент (для администраторов)

🚀 Установка и запуск
Требования
Go 1.20+

PostgreSQL (или SQLite)

Запуск
Клонировать репозиторий:

sh
git clone https://github.com/yourusername/virgil-dmc-api.git
cd virgil-dmc-api
Установить зависимости:

sh
go mod download
Настроить .env (пример в .env.example):

sh
cp .env.example .env
Запустить сервер:

sh
go run main.go
Документация API
После запуска Swagger UI будет доступен по адресу:
http://localhost:8080/swagger/index.html

📁 Структура проекта
├── config/          # Конфигурация приложения  
├── controllers/     # Обработчики запросов  
├── models/          # Модели данных и работа с БД  
├── routes/          # Маршруты API  
├── static/          # Изображения и статика  
├── docs/            # Swagger документация  
├── main.go          # Точка входа  
└── README.md  
🌟 Примеры запросов
Получить информацию о Вергилии
http
GET /api/v1/vergil/info  
Получить случайную цитату
http
GET /api/v1/vergil/quotes/random  
🤝 Участие
Приветствуются пул-реквесты и предложения!

📜 Лицензия
MIT

🔥 "I need more power!" — Vergil, Devil May Cry 5
