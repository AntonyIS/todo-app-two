package app

type UserServiceInterface interface {
	CreateUser(u *UserModel) error
	User(id string) (*UserModel, error)
	Users() (*[]UserModel, error)
	UpdateUser(u *UserModel) error
	DeleteUser(id string) error
}
