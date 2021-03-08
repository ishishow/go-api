package handler

import (
	"log"

	"golang.org/x/net/context"

	"20dojo-online/pkg/interfaces/rpc/user"
	"20dojo-online/pkg/usecase/interactor"
)

// UserService ユーザーサービス
type UserService struct {
	UserUseCase interactor.UserUseCase
}

// NewUserService Userデータに関するServiceを生成
func NewUserService(userUseCase interactor.UserUseCase) UserService {
	return UserService{
		UserUseCase: userUseCase,
	}
}

// CreateUser generates authToken response
func (u *UserService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	authToken, err := u.UserUseCase.Create(req.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.CreateUserResponse{AuthToken: authToken}, nil
}

// ReadUser generates user response
func (u *UserService) ReadUser(ctx context.Context, req *user.ReadUserRequest) (*user.ReadUserResponse, error) {
	authToken, _ := ctx.Value("key").(string)
	readuser, err := u.UserUseCase.SelectByAuthToken(authToken)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(readuser)
	return &user.ReadUserResponse{User: &user.User{
		Id:        readuser.ID,
		Name:      readuser.Name,
		HighScore: readuser.HighScore,
		Coin:      readuser.Coin,
	}}, nil
}

// UpdateUser generates user_update_result response
func (u *UserService) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	authToken, _ := ctx.Value("key").(string)
	return &user.UpdateUserResponse{}, u.UserUseCase.UpdateName(authToken, req.Name)
}
