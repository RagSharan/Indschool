package services

import (
	"../entity"
	"../repository"
)

type UserService interface {
	CreateUser(user entity.User) (entity.User, error)
	FindUser(userName string) (entity.User, error)
	FindUserList() ([]*entity.User, error)
}

type services struct{}

var (
	//user     entity.User = entity.User{UserName: "golang name", TagLine: "tagline"}
	userRepo repository.UserRepo
)

func NewUserService(repo repository.UserRepo) UserService {
	userRepo = repo
	return &services{}
}

func (*services) CreateUser(user entity.User) (entity.User, error) {
	user, err := userRepo.CreateUser(user)
	// if err != nil {
	// 	return err
	// }
	return user, err
}
func (*services) FindUser(userName string) (entity.User, error) {
	return userRepo.FindUser(userName)

}
func (*services) FindUserList() ([]*entity.User, error) {

	return userRepo.FindUserList()
}
