package students

import "time"

type Student struct {
	Id        uint `gorm:"primary_key;auto_increment"`
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
