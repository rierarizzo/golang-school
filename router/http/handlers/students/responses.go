package students

type StudentResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type ListResponse struct {
	Data []StudentResponse `json:"data"`
}
