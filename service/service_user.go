package service

import (
	"api_red_social/model"
	"api_red_social/repository"
	"api_red_social/utils/password"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: userRepo,
	}
}

func (service *UserService) GetUserProfile(id string) (*model.Users, error) {
	user, err := service.Repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) Login(email, password string) (string, *model.Users, error) {
	user, err := service.Repo.FindUserByEmail(email)
	if err != nil {
		return "", nil, errors.New("no existe el usuario")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("no te has identificado correctamente")
	}

	token, err := service.CreateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (service *UserService) CreateToken(user *model.Users) (string, error) {
	claims := jwt.MapClaims{
		"id":   user.User_id,
		"name": user.Name,
		"nick": user.Nick,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte("Fuerzaabasto1@"))
}

func (service *UserService) CreateNewUser(users *model.Users) (primitive.ObjectID, error) {
	hashedPassword, err := password.Hash(users.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	users.Password = string(hashedPassword)
	users.User_id = primitive.NewObjectID()
	users.CreatedAt = time.Now()

	return service.Repo.CreateNewUser(users)
}
