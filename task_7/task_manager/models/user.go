package models

type User struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Role     string `json:"role" bson:"role"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PromoteRequest struct {
	Username string `json:"username" binding:"required"`
}

