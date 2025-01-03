package utils

import (
	"auth/database"
	"auth/envs"
	"auth/models"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Хэшируем пароль при помощи bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 - стоимость хэширования
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Проверяем хэш пароль при помощи bcrypt
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Функция для генерации JWT-токена шанс колизии токенов (8.63 * 10^-78)
func GenerateTokens(userID uint) (models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Срок действия токена 24 часа (exp)
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 600).Unix(), // Срок действия токена 1 месяц
	})
	signedAccessToken, _ := accessToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	signedRefreshToken, _ := refreshToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	return models.Tokens{AccessToken: signedAccessToken, RefreshToken: signedRefreshToken}, nil
}

// Функция проверки JWT-токена

func ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDValue, ok := claims["user_id"].(float64) // приведение к float64
		if !ok {
			return 0, fmt.Errorf("user_id claims is not float64")
		}
		return uint(userIDValue), nil // Конвертация float64 в uint
	} else {
		return 0, fmt.Errorf("недействительный токен")
	}
}

// Извлекаем ID пользователя из JWT AccessToken
func ExtractUserID(tokenString string) (uint, error) {
	// Отсечение 'Bearer' из заголовка
	str := strings.TrimSpace(strings.TrimPrefix(tokenString, "Bearer"))

	// Проверяем что токен валиден
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		// Убедимся что наш алгоритм соответствует 'jwt.SigningMethodHS256'
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный алгоритм подписи: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"]

		if userIDFloat, ok := userID.(float64); ok {
			return uint(userIDFloat), nil // Преобразуем float64 в uint
		}
	}
	return 0, fmt.Errorf("Невозможно извлечь user_id из токена")
}

// IsValidPhone проверяет, является ли номер телефона валидным и возвращает сам номер и ошибку
func IsValidPhone(phone string) (string, error) {
	// Если номер пустой, возвращаем ошибку
	if phone == "" {
		return "", errors.New("номер телефона не может быть пустым")
	}

	// Убираем пробелы, если они есть
	phone = strings.TrimSpace(phone)

	// Если номер начинается с "+", то он должен содержать 11 цифр
	if phone[0] == '+' {
		if len(phone) == 12 && regexp.MustCompile(`^\+7\d{10}$`).MatchString(phone) {
			return phone, nil
		}
		return "", errors.New("телефонный номер, начинающийся с '+', должен содержать ровно 12 цифр после '+'")
	}

	// Если номер начинается с цифры или открывающей скобки, то он должен содержать 10 цифр
	if phone[0] == '(' || (phone[0] >= '0' && phone[0] <= '9') {
		if len(phone) == 10 && regexp.MustCompile(`^8\d{10}$`).MatchString(phone) {
			return phone, nil
		}
		return "", errors.New("телефонный номер, начинающийся с цифры или '(', должен содержать ровно 10 цифр")
	}

	// Регулярное выражение для проверки номера с дефисами и скобками
	re := `^\+7[\s\(\)-]?\d{3}[\s\(\)-]?\d{3}[\s\(\)-]?\d{2}[\s\(\)-]?\d{2}$`
	match, _ := regexp.MatchString(re, phone)

	// Проверка на корректность формата и что номер заканчивается цифрой
	if match && phone[len(phone)-1] >= '0' && phone[len(phone)-1] <= '9' {
		return phone, nil
	}

	// Если невалидный номер, возвращаем ошибку
	return "", errors.New("неверный формат телефонного номера")
}

func IsValidEmail(email string) (string, error) {
	// Если строка email пустая - ошибка
	if email == "" {
		return "", errors.New("Пустая строка email")
	}
	// Убираем пробелы, если они есть
	email = strings.TrimSpace(email)

	// Регулярное выражение для проверки email
	re := `^[a-zA-Z0-9._%+-]{1,64}@(yandex\.ru|mail\.ru|gmail\.com)$`
	match, _ := regexp.MatchString(re, email)

	if match {
		return email, nil
	}
	return "", errors.New("Неверный формат email")
}

// Функция для получения пользователя по email
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Функция для получения пользователя по ID
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
