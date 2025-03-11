package database

import (
	"encoder/domain"
	"log"

	"github.com/jinzhu/gorm"                   // ОРМ
	_ "github.com/jinzhu/gorm/dialects/sqlite" //Sqlite
	_ "github.com/lib/pq"                      //Postgresql
)

// Database Описание Баз Данных
type Database struct {
	Db            *gorm.DB //Ссылка на ORM
	Dsn           string   //путь к БД прода
	DsnTest       string   //путь к тестовой БД
	DbType        string   //Тип базы прода
	DbTypeTest    string   //Тип тестовой БД
	Debug         bool     //Включать ли отладочный режим
	AutoMigrateDb bool     //Делать ли автоматические миграции
	Env           string   // test или prod
}

// NewDatabase Создать новую Базу Данных
func NewDatabase() *Database {
	return &Database{}
}

// NewDatabaseTest Новая тестовая База Данных в sqlite
func NewDatabaseTest() *gorm.DB {
	dbInstance := NewDatabase()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3" //Тип БД
	dbInstance.DsnTest = ":memory:"   //В памяти
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	return connection
}

// Connect Метод соединения с базой данных
func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	//Если задано окружение для тестовой базы используем тестовую базу, иначе используем базу для прода
	if d.Env == "test" {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest) //Открываем тестовую БД
	} else {
		d.Db, err = gorm.Open(d.DbType, d.Dsn) // Открываем БД прода
	}

	if err != nil {
		return nil, err
	}

	// Включаем логи в отладочном режиме
	if d.Debug {
		d.Db.LogMode(true)
	}

	//Если автоматичесукие миграции включены, делаем миграции
	if d.AutoMigrateDb {
		//Делаем миграции Video и Job
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		//Добавляем внешние ключи в БД
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos(id)", "CASCADE", "CASCADE")
	}

	return d.Db, nil
}
