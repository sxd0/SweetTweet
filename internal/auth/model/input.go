package model

type RegisterInput struct {
    Email    string `json:"email" example:"user@example.com"`
    Password string `json:"password" example:"123456"`
}

type UserResponse struct {
    ID    int    `json:"id" example:"1"`
    Email string `json:"email" example:"user@example.com"`
}

type ErrorResponse struct {
    Message string `json:"message" example:"Invalid input"`
}

type LoginInput struct {
    Email    string `json:"email" example:"user@example.com"`
    Password string `json:"password" example:"123456"`
}

type LoginResponse struct {
    Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type MeResponse struct {
    UserID int `json:"user_id" example:"1"`
}
