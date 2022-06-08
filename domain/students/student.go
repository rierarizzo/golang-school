package students

import "time"

type Student struct {
	Id        uint
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
