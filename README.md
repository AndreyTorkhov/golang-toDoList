# ToDo Приложение на Go

## Описание

Это простое ToDo-приложение, реализованное на языке программирования Golang. Приложение позволяет управлять задачами: добавлять новые, удалять, помечать как выполненные и получать список актуальных задач.

## Стек технологий

- **Go** - язык программирования.
- **PostgreSQL** - реляционная база данных для хранения задач.
- **Chi** - маршрутизатор для обработки HTTP-запросов.
- **pgx** - библиотека для работы с PostgreSQL в Go.
- **docker** - платформа для разработки, доставки и эксплуатации приложений.
- **swagger** - это набор инструментов для создания документации к API.

## Установка

`Перед использованием в корне проекта создайте файл .env и запишите в него переменные на прнимере файла example.env`

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/AndreyTorkhov/todo-golang.git
   cd todo-golang
   ```
2. Запуск контейнеров:

   ```bash
    docker-compose up --build
   ```

3. Остановка контейнеров:

   ```bash
    docker-compose stop
   ```

## Используемые функции

1. `GET` `/tasks` - Получить список всех задач.
2. `GET` `/tasks/{id}` - Получить список всех задач.
3. `POST` `/tasks` - Добавить новую задачу.
4. `PATCH` `/tasks/{id}/done` - Обновить задачу (пометить как выполненную).
5. `DELETE` `/tasks/{id}` - Удалить задачу по ID.
6. `GET` `/tasks/filter?done=` - Удалить задачу по ID.

## Документация

`http://localhost:8080/docs/index.html/` - Swagger UI.
