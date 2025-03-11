# Описание кодов

Начальный скрипт  `video.go`, `job.go`

Тесты  `video_test.go`, `job_test.go` - это еще примеры использования

В файле `go.mod` описан корневой модуль `encoder` так `module encoder`

Файлы `video.go`, `job.go` относятся к пакету `encoder/domain`
Файл `db.go` относится к пакету `encoder/framework/database`
Файлы `job_repository.go`, `video_repository.go` относятся к пакету `encoder/application/repositories`

Использованные внешние библиотеки
```GO
require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.1.1
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```






