package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	FullName string `bson:"full_name"     json:"full_name"     validate:"required,min=3,max=45"`
	Email    string `bson:"email"         json:"email"         validate:"required,email"`
	Password string `bson:"password"      json:"password"      validate:"required"`
}

type RegisterResponse struct {
}
