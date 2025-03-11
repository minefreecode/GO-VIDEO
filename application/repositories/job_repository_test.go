package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"fmt"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDatabaseTest() //Создается тестовая БД
	defer db.Close()                 // Дожидается окончания функции и закрывает БД

	fmt.Println("Database instance created") //Сообщение что БД создана

	video := domain.NewVideo()       // Создаем новое видео
	video.ID = uuid.NewV4().String() // Создаем идентификатор для видео
	video.FilePath = "path"          //Задаем путь для видео
	video.CreatedAt = time.Now()     //Добавляем время создания

	repo := repositories.VideoRepositoryDb{Db: db} //Создаем репозиторий
	repo.Insert(video)                             //Вставляем видео в БД

	job, err := domain.NewJob("output_path", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, job.ID, j.ID)
	require.Equal(t, video.ID, j.VideoID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDatabaseTest()
	defer db.Close()

	fmt.Println("Database instance created")

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "completed"
	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, job.Status, j.Status)
}
