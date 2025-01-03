package handlers

import (
	"auth/database"
	"auth/models"
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerUserHandler(ctx *gin.Context) {
	var user models.User
	var registerData models.RegisterData

	// Парсим JSON тело запроса в структуру RegisterData
	if err := ctx.ShouldBind(&registerData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Хэшируем пароль
	hashedPassword, err := utils.HashPassword(registerData.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
		return
	}

	// Проверяем валидность телефона
	validPhone, err := utils.IsValidPhone(registerData.Phone)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный номер телефона"})
		return
	}

	// Проверяем, существует ли уже пользователь с таким номером телефона
	var anExistingUserOnThePhone models.User
	if err := database.DB.Where("phone = ?", validPhone).First(&anExistingUserOnThePhone).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Номер телефона уже зарегистрирован"})
		return
	}

	// Проверяем валидность email
	validEmail, err := utils.IsValidEmail(registerData.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный email"})
		return
	}
	// Проверяем, существует ли уже пользователь с таким email
	var anExistingUserOnTheEmail models.User
	if err := database.DB.Where("email = ?", validEmail).First(&anExistingUserOnTheEmail).Error; err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Этот email уже зарегистрирован"})
		return
	}

	// Если номер телефона и email уникален, создаем нового пользователя
	user.UserName = registerData.UserName
	user.Phone = validPhone
	user.Email = validEmail
	user.Hash = hashedPassword

	// Сохраняем пользователя в базу данных
	result := database.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить пользователя"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Регистрация пользователя прошла успешно"})
}
