# BibTex Server

Основной backend для всей базы данных.

Зависимости:

* `github.com/lib/pq`

Для запуска необходимо иметь auth-key для сервиса Google Books.

Команда запуска:  

`./server --googleToken=code --postgrePort=port_int`

Структура проекта:  

* resources -- информация о базе данных
* src -- исходный код
