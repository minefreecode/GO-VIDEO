package domain

import (
	"time"

	"github.com/asaskevich/govalidator" // Валидатор
)

// Video Описана структура с валидациями
type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuidv4" gorm:"type:uuid;primary_key"` //Идентфикатор
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:uuid"`                     //Идентфикатор ресурса
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`               //Путь к файлу
	CreatedAt  time.Time `json:"-" valid:"-"`                                                      //Время создания
	Job        []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`                            //Ссылка на Job в Базе Данных
}

// Инициализация пакета
func init() {
	// Активирует требование проверять все поля
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	// Новый объект
	return &Video{}
}

func (video *Video) Validate() error {
	//Валидация объекта
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}

	return nil
}
