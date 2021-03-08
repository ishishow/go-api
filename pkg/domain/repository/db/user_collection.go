package db

import (
	ucm "20dojo-online/pkg/domain/model/user_collection_item"
)

// UserCollectionItemRepository データ永続化のために抽象化したUserCollectionItemデータ更新周りの処理
type UserCollectionItemRepository interface {
	SelectUserCollectionItemsByUserID(userID string) ([]*ucm.UserCollectionItem, error)
}
