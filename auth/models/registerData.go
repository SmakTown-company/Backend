package models

// Структура данных хранения данных для регистрации пользователя
type RegisterData struct {
	UserName string `json:"username" binding:"required,min=4"`
	// Должен быть валидный email
	Email string `json:"email" binding:"required"`
	// Должен быть валидный phone
	Phone string `json:"phone" binding:"required,min=12"`
	// Требуется валидный пароль
	Password string `json:"password" binding:"required,min=8"`
}
