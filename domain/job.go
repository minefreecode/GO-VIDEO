package domain

import (
	"time"

	"github.com/asaskevich/govalidator" //Валидатор
	uuid "github.com/satori/go.uuid"    //Идентификатор
)

// Тип, представляющий Job
type Job struct {
	ID               string    `json:"job_id" valid:"uuidv4" gorm:"type:uuid;primary_key"`           //Идентификатор Job
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"`                           //Вывод пакета джоб
	Status           string    `json:"status" valid:"notnull, in(pending|ongoing|completed|failed)"` //Статус Job
	Video            *Video    `json:"video" valid:"-"`                                              //Видео
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`         //Идентификатор видео
	Error            string    `valid:"-"`                                                           //Ошибка если возникла
	CreatedAt        time.Time `json:"created_at" valid:"-"`                                         //Время создания
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`                                         //Время обновления
}

func init() {
	//Активизируем валидатор
	govalidator.SetFieldsRequiredByDefault(true)
}

// NewJob Создание новой Job и валидация
func NewJob(outputBucketPath string, video *Video) (*Job, error) {
	//Создаем новую Job
	job := Job{
		OutputBucketPath: outputBucketPath,
		Video:            video,
		Status:           "pending",
	}

	job.prepare()

	err := job.Validate()
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// Заполнение Job дополнительными условиями
func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}

// Validate Добавление к типу данных метода валидации
func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job) //Валидация
	if err != nil {
		return err
	}
	return nil
}
