package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	todo "github.com/toolnado/rest-todo-app"
	"github.com/toolnado/rest-todo-app/pkg/repository"
)

const (
	salt       = "wefjefkek24234235kqklwshmnmkjk"
	signingKey = "wefaef#$%1q2431543643rhGGESC24234235kqklw"
	tockenTTl  = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenKlaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenKlaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tockenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
