package auth

type LoginRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ProfileResponse struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	GoldAmount float64 `json:"goldAmount"`
}
