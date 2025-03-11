package domain_test

// Импортирование
import (
	"encoder/domain" //Библиотека для кодирования
	"testing"        //Библиотека для автоматизированного тестирования Go пакетов
	"time"           //Библиотека для измерения и показа времени

	uuid "github.com/satori/go.uuid"      //Библиотека обеспечивает реализацию UUID
	"github.com/stretchr/testify/require" //Библиотека реализовывает некоторые assert
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceID = "abc"
	video.FilePath = "abc"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "abc"
	video.FilePath = "abc"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
