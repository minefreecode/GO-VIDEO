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

// Тест на вставку новых видео в репозиторий
func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDatabaseTest() //Создаем тестовую БД
	defer db.Close()                 //Ожидание конца выполнения функции и закрытие БД

	fmt.Println("Database instance created")

	video := domain.NewVideo()       //Создаем новое видео
	video.ID = uuid.NewV4().String() //Создаем идентификатор
	video.FilePath = "path"          //Устанавливаем тестовый путь на видео
	video.CreatedAt = time.Now()     // Создаем время

	repo := repositories.VideoRepositoryDb{Db: db} //Создаем новый репозиторий
	repo.Insert(video)                             //Вставка нового видео

	v, err := repo.Find(video.ID) //Поиск видео

	require.NotEmpty(t, v.ID)        //Проверка что не пусто
	require.Nil(t, err)              // Проверка что не было ошибок
	require.Equal(t, video.ID, v.ID) //Проверка что идентфиикатор найденного видео совпадает с идентификатором поиска
}
