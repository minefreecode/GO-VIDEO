package repositories // Package repositories Основной реализующий пакет проекта

import (
	"encoder/domain" // библиотека для кодирования
	"fmt"            //Библиотека для форматированного ввода/вывода

	"github.com/jinzhu/gorm"         //Библиотека для ORM, для работы с Базами Данных
	uuid "github.com/satori/go.uuid" //Библиотека реализующая UUID
)

// VideoRepository Функции для работы с репозиториями Баз Данных. В данном случае он не применяется
type VideoRepository interface {
	// Insert Функция для вставки видео
	Insert(video *domain.Video) (*domain.Video, error)
	// Find  Функция для поиска видео
	Find(id string) (*domain.Video, error)
}

// VideoRepositoryDb ORM Базы Данных для работы с репозиториями
type VideoRepositoryDb struct {
	Db *gorm.DB //Указатель на ORM Базы Данных
}

// Insert Добавляем в структуру VideoRepositoryDb функцию Insert
func (repo VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	// Если идентификатор не задан задаем
	if video.ID == "" {
		video.ID = uuid.NewV4().String() //Идентификатор UUID
	}

	// Добавляем видео в Базу Данных. Если возникла ошибка считываем его
	err := repo.Db.Create(video).Error

	//Если при обращении к Базе Данных возникла ошибка, возвращаем ошибку
	if err != nil {
		return nil, err
	}

	// Результатом является видео
	return video, nil
}

// Find Добавляем в структуру VideoRepositoryDb функцию Find
func (repo VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	// Присваиваем переменной video значение из объекта domain
	var video domain.Video
	//Загрузка Джобов из Баз Данных по условию, что id совпадает с id видео
	repo.Db.Preload("Jobs").First(&video, "id = ?", id)

	//Нсли видео не найдено сообщаем об этом возвращая значение из функции
	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	// Результатом является ссылка видео
	return &video, nil
}
