package userService

import "ransmart_pay/app/models/userModel"

type IUserService interface {
	FindByUsername(username string) (userModel.User, error)
	Create(user userModel.User) (err error)
	Update(id int, User userModel.User) (err error)
}
