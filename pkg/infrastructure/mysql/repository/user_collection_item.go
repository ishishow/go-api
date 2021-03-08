package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"20dojo-online/pkg/domain/repository/db"
	"20dojo-online/pkg/infrastructure/mysql"

	ucm "20dojo-online/pkg/domain/model/user_collection_item"
)

type userCollectionItemRepositoryImpl struct {
	mysql.SQLHandler
}

// NewUserCollectionItemRepositoryImpl Userに関するDB更新処理を生成
func NewUserCollectionItemRepositoryImpl(sqlHandler mysql.SQLHandler) db.UserCollectionItemRepository {
	return &userCollectionItemRepositoryImpl{
		sqlHandler,
	}
}

// SelectUserCollectionItemsByUserID user_idと一致する行全て取得
func (uri userCollectionItemRepositoryImpl) SelectUserCollectionItemsByUserID(userID string) ([]*ucm.UserCollectionItem, error) {
	rows, err := uri.SQLHandler.Conn.Query("SELECT * FROM user_collection_item WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return convertToUserCollectionItems(rows)
}

// InsertUserCollectionItemsTx ctx, txを受け取り、user_collection_itemに新規データを全て挿入する
func insertUserCollectionItemsTx(ctx context.Context, tx *sql.Tx, records []*ucm.UserCollectionItem) error {
	insertValue := ""
	for _, record := range records {
		insertValue += fmt.Sprintf("('%s', '%s'),", record.UserID, record.CollectionItemID)
	}
	sqlRaw := fmt.Sprintf(`INSERT INTO user_collection_item(user_id, collection_item_id) VALUES%s`, insertValue)
	sqlRaw = strings.TrimRight(sqlRaw, ",")

	_, err := tx.ExecContext(ctx, sqlRaw)
	return err
}

// convertToUserCollectionItems rowsデータをUserCollectionItemsデータへ変換する
func convertToUserCollectionItems(rows *sql.Rows) ([]*ucm.UserCollectionItem, error) {
	var userCollectionItems []*ucm.UserCollectionItem
	for rows.Next() {
		userCollectionItem := ucm.UserCollectionItem{}
		if err := rows.Scan(&userCollectionItem.UserID, &userCollectionItem.CollectionItemID); err != nil {
			return nil, err
		}
		userCollectionItems = append(userCollectionItems, &userCollectionItem)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userCollectionItems, nil
}
