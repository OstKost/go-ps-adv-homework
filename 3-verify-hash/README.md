# go-ps-adv-homework

# URL Shortener Go App

## Возможности:

- Регистрация \ Вход
- Создание короткой ссылки
- Получение статистики

## Cтруктура:

- CMD: Это папка для точки входа приложения. Содержит main-файлы, которые запускают приложение. Если у вас несколько приложений, например, веб-сайт и API, для каждого будет отдельный файл в этой папке.
- Internal: Здесь содержатся основные бизнес-модули, которые не экспортируются наружу и не могут быть использованы как внешняя зависимость. Они распределяются на отдельные модули по функциональности (например, пользователи, статистика).
- Package (pkg): В этой папке размещается общий код библиотек, который может переиспользоваться в других проектах, такие как обработчики, подключения к базам данных и другие утилиты.
- Configs: Шаблоны и модули конфигурации приложения. Здесь настроены параметры подключения к базам данных, URL и другие конфигурации.
- Migrations: Хранит миграции базы данных. Только важно учитывать, что структура может расширяться и дополняться в зависимости от потребностей проекта.