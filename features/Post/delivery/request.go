package delivery

import "backend/domain"

type PostFormat struct {
	User    domain.User `gorm:"foreignKey:Username"`
	Photo   string      `json:"Photo"`
	Caption string      `json:"Caption" validate:"required"`
}

func (i *PostFormat) ToModel() domain.Post {
	return domain.Post{
		Photo:   i.Photo,
		Caption: i.Caption,
	}
}
