# BibTex CLI

Утилита командной строки для возможной автоматизации работы с приложением.
Общий принцип работы: каждый вызов работает отдельно, сохраняя, при необходимости,
данные в промежуточные файлы(можно задавать имена этих файлов флагами). Взаимодействие
с пользователем происходит при помощи корректирования JSON-структур.

Работа с приложением строится по следующему принципу:  
`./cli tableName cmd`, где tableName -- имя таблицы, а cmd -- одна из следующих команд:  

* add -- добавление данных из файла в таблицу
* get -- получение данных с сервера
* prototype -- получение прототипа JSON для добавления записи в заданную таблицу
