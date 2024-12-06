package models

// Структура данных хранения данных для регистрации пользователя
type RegisterData struct {
	// Должен быть валидный email
	Email string `json:"email" binding:"required,email"`
	// Требуется валидный пароль
	Password string `json:"password" binding:"required,min=8"`
}
