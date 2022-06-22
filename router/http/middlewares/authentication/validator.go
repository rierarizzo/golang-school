package authentication

type LoginValidator struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}
