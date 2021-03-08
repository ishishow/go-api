package db

import (
	um "20dojo-online/pkg/domain/model/user"
)

// UserRepository データ永続化のために抽象化したUserデータ更新周りの処理
type UserRepository interface {
	Create(ID, authToken, name string) error
	SelectByAuthToken(authToken string) (*um.User, error)
	SelectByPrimaryKey(userID string) (*um.User, error)
	SelectAllPlayingUsers() ([]*um.User, error)
	Update(user *um.User) error
}
