package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

// JobRepository Методы репозитория
type JobRepository interface {
	// Insert Вставить
	Insert(job *domain.Job) (*domain.Job, error)
	// Find Найти
	Find(id string) (*domain.Job, error)
	// Update Обновить
	Update(job *domain.Job) (*domain.Job, error)
}

// JobRepositoryDb Репозиторий БД
type JobRepositoryDb struct {
	Db *gorm.DB
}

// Insert Добавляем метод вставки Job в БД
func (repo JobRepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {

	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

// Find Добавляем метод поиска в БД
func (repo JobRepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.Db.Preload("Video").First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &job, nil
}

// Update Добавляем метод обновления
func (repo JobRepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
