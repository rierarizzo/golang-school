package users

type LoginRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
}

type RegisterRequest struct {
	Username string `binding:"required" json:"username"`
	Password string `binding:"required" json:"password"`
	// TODO
}
