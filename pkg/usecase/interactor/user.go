package interactor

import (
	"github.com/google/uuid"

	um "20dojo-online/pkg/domain/model/user"
	ur "20dojo-online/pkg/domain/repository/db"
)

// UserUseCase Userにおけるユースケースのインターフェース
type UserUseCase interface {
	Create(name string) (authToken string, err error)
	Get(userID string) (user *um.User, err error)
	SelectByAuthToken(authToken string) (*um.User, error)
	UpdateName(userID string, name string) error
	SelectAllPlayingUsers() ([]*um.User, error)
}

type userUseCase struct {
	repository ur.UserRepository
}

// NewUserUseCase Userデータに関するユースケースを生成
func NewUserUseCase(userRepo ur.UserRepository) UserUseCase {
	return &userUseCase{
		repository: userRepo,
	}
}

// CreateUser Userを新規作成するためのユースケース
func (uu userUseCase) Create(name string) (string, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	authToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	if err := uu.repository.Create(userID.String(), authToken.String(), name); err != nil {
		return "", err
	}

	return authToken.String(), nil
}

// SelectByAuthToken Userをトークンから取得するためのユースケース
func (uu userUseCase) SelectByAuthToken(authToken string) (*um.User, error) {
	return uu.repository.SelectByAuthToken(authToken)
}

func (uu userUseCase) Get(userID string) (*um.User, error) {
	return uu.repository.SelectByPrimaryKey(userID)
}

func (uu userUseCase) SelectAllPlayingUsers() ([]*um.User, error) {
	return uu.repository.SelectAllPlayingUsers()
}

func (uu userUseCase) UpdateName(userID string, name string) error {
	user, err := uu.repository.SelectByAuthToken(userID)
	if err != nil {
		return err
	}
	user.Name = name
	return uu.repository.Update(user)
}
