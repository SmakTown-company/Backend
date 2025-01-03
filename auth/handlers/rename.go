package handlers

import (
	"auth/database"
	"auth/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"username" json:"username"`
	Email    string `gorm:"not nul;unique" json:"email"`
	Phone    string `gorm:"not nul;unique" json:"phone"`
	Hash     string `gorm:"hash" json:"-"`
}

type renameData struct {
	UserName string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) UpdateUserName(newUserName string) {
	u.UserName = newUserName
}

func (u *User) UpdateEmail(newEmail string) (string, error) {
	validEmail, err := utils.IsValidEmail(newEmail)
	if err != nil {
		return "", errors.New("Невалидный email")
	}
	u.Email = validEmail
	return "", nil
}

func (u *User) UpdatePhone(newPhone string) (string, error) {
	validPhone, err := utils.IsValidPhone(newPhone)
	if err != nil {
		return "", errors.New("Невалидный номер телефона")
	}
	u.Phone = validPhone
	return "", nil
}

func (u *User) UpdatePassword(newPassword string) error {
	if len(newPassword) < 8 {
		return errors.New("Пароль должен содержать не меньше 8 символов")
	}
	newHashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("Не удалось хэшировать пароль")
	}
	u.Hash = newHashedPassword
	return nil
}

func renameUserHandler(ctx *gin.Context) {
	// Получаем email пользователя из запроса
	email := ctx.Param("email")

	// Извлекаем пользователя из базы данных
	var user User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	var requestData renameData

	// Декодируем данные из тела запроса в структуру renameData
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		// Если ошибка при парсинге запроса, возвращаем ошибку 400
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Обновляем имя пользователя, если оно было передано в запросе
	if requestData.UserName != "" {
		user.UpdateUserName(requestData.UserName)
	}

	// Обновляем email, если он был передан в запросе
	if requestData.Email != "" {
		_, err := user.UpdateEmail(requestData.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Обновляем номер телефона, если он был передан в запросе
	if requestData.Phone != "" {
		_, err := user.UpdatePhone(requestData.Phone)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Обновляем пароль, если он был передан в запросе
	if requestData.Password != "" {
		if err := user.UpdatePassword(requestData.Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Сохраняем обновленного пользователя в базу данных
	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить обновленные данные"})
		return
	}

	// Возвращаем успешный ответ с обновленными данными пользователя
	ctx.JSON(http.StatusOK, gin.H{
		"username": user.UserName,
		"email":    user.Email,
		"phone":    user.Phone,
		"hash":     user.Hash,
	})
}
