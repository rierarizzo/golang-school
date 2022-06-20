package students

type StudentRequest struct {
	Name    string `binding:"required" json:"name"`
	Surname string `binding:"required" json:"surname"`
}