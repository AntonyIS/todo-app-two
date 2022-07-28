package app

import (
	"errors"
	"os/exec"
	"time"

	er "github.com/pkg/errors"
)

var (
	ErrorInvalidUser         = errors.New("invalid user")
	ErrorUserNotFound        = errors.New("user not found")
	ErrorInternalServerError = errors.New("internal server error")
)

type userService struct {
	userRepo UserRepositoryInterface
}

func NewUserService(userRepo UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepo,
	}
}

func (svc *userService) CreateUser(u *UserModel) error {
	publicID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return er.Wrap(ErrorInternalServerError, "logic.CreateUser")
	}
	createdAt := time.Now().UTC().Unix()
	u.PublicID = string(publicID)
	u.CreatedAt = createdAt
	return svc.CreateUser(u)
}

func (svc *userService) User(id string) (*UserModel, error) {
	return svc.userRepo.User(id)
}

func (svc *userService) Users() (*[]UserModel, error) {
	return svc.userRepo.Users()
}

func (svc *userService) UpdateUser(u *UserModel) error {
	return svc.userRepo.UpdateUser(u)
}

func (svc *userService) DeleteUser(id string) error {
	return svc.userRepo.DeleteUser(id)
}
