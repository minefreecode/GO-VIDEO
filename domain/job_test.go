package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateNewJob(t *testing.T) {
	//Создание видео
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	//Создание связной Job
	job, err := domain.NewJob("path", video)
	require.Nil(t, err)
	require.NotNil(t, job)
}
