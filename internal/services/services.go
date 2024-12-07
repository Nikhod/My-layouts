package services

import (
	"Nikcase/internal/repositories"
	"Nikcase/pkg/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

const (
	signingKey = "grkjk#4#%35FSFJLja#4353KSFjH"
	lifetime   = 3 * time.Hour
)

type Service struct {
	Repository *repositories.Repository
}

func NewService(repository *repositories.Repository) *Service {
	return &Service{Repository: repository}
}

// Проверяет пользователя на Существование логина и правильности пароля
func (s *Service) ValidatePassAndLogin(login, pass string) (*models.Users, error) {
	user, err := s.Repository.ValidateLogin(login)
	if err != nil {
		log.Println("Error in func 'ValidatePassAndLogin'")
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		log.Println("Пароль неверный!")
		return nil, err
	}
	return user, nil
}

//____________________________________________________________________________________

// Проверяет валидность данных пользователя при регистрации (Длина введённых данных)
func (s *Service) IsValidDataForRegistration(user *models.Users) error {
	if len(user.Name) < 3 || len(user.Login) < 8 || len(user.Password) < 8 {
		if len(user.Name) > 20 || len(user.Login) > 20 || len(user.Password) > 20 {
			err := errors.New("Неправильно Введеные данные Имени/Логина/Пароля Для регистрации")
			return err
		}
	}
	return nil
}

// Регистрирует Пользователя генерируя ему Хэшированный пароль
func (s *Service) RegistrationUser(user *models.Users) error {
	isFree := s.Repository.IsLoginFree(user.Login)
	if isFree == false {
		err := errors.New("Логин Занят другим пользователем!")
		return err
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = s.Repository.RegistrateUserToDB(user, hashPass)
	if err != nil {
		return err
	}
	return nil
}

// Возвращает сгенерированный Токен по логину и паролю
func (s *Service) GenerateToken(login, password string) (string, error) {
	user, err := s.Repository.GetUser(login)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifetime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: user.Login,
	})
	token, err := tok.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	err = s.Repository.AddTokenToDB(user.Id, token)
	if err != nil {
		return "", err
	}
	return token, nil
}
