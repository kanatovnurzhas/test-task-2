# test-task-2

### Description
Проект состоит из двух CRUD микросервисов. Решил использовать фреймворк fiber, так как он очень понятен для использования, также имеют очень понятную документацию. Можно было бы использовать и gin, выбрал fiber так как хотел попробовать что то новое.

### Database
Как базу данных выбрал postgresql. Так как ранее работал с ним и реализовывал несколько CRUD на нем. Посчитал также преимуществом от noSQL баз данных он имеет строго определенную схему данных и поля должны быть переданы все тогда как в noSQL можно некоторые поля не передавать что может быть плюсом также как и минусом

### Logging
Логирование до этого не использовал. Увидел что с коробки в fiber есть логрус и решил прикрутить и его)

### Связь между микросервисами
Связь между микросервисами решил реализовать простыми запросами к друг другу. Создал нужные эндпоинты для этого и чуть чуть изменил вид эндпоинтов. Изначально передавали id курсов или студентов, так как таблицы не были связаны друг с другом решил немножно изменить этот момент. Решил искать в колонках имена студентов или курсов соответственно. Можно было бы и через id реализовать но для этого нужно было бы сходить лишний раз в бд и посмотреть существуют ли такие бд и после этого их вытаскивать.

### Запуск
Для начала нужно запустить докер образ postgresql: docker run --name postgresql-container -p 5432:5432 -e POSTGRES_PASSWORD=admin -d postgres
Далее в двух терминалах запустить два сервиса : go run app1/cmd/main.go && go run app2/cmd/main.go
