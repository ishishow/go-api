package interactor

import ur "20dojo-online/pkg/domain/repository/db"

// CollectionItemUseCase is
type CollectionItemUseCase interface {
	SelectALLUserCollectionItems(userID string) ([]*Collection, error)
}

type collectionItemUseCase struct {
	collectionRepository     ur.CollectionItemRepository
	userCollectionRepository ur.UserCollectionItemRepository
}

// NewCollectionItemUseCase Userデータに関するユースケースを生成
func NewCollectionItemUseCase(collectionRepo ur.CollectionItemRepository,
	userCollectionRepo ur.UserCollectionItemRepository) CollectionItemUseCase {
	return &collectionItemUseCase{
		collectionRepository:     collectionRepo,
		userCollectionRepository: userCollectionRepo,
	}
}

func (uu collectionItemUseCase) SelectALLUserCollectionItems(userID string) ([]*Collection, error) {

	collectionItems, err := uu.collectionRepository.SelectAllCollectionItems()
	if err != nil {
		return nil, err
	}

	// userIDに紐づくUserCollectionItemを全件取得
	userCollectionItems, err := uu.userCollectionRepository.SelectUserCollectionItemsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// CollectionItem_IDでkey判別するためのMAPを作成
	userCollectionItemMapByCollectionItemID := make(map[string]struct{}, len(collectionItems))
	for _, userCollectionItem := range userCollectionItems {
		userCollectionItemMapByCollectionItemID[userCollectionItem.CollectionItemID] = struct{}{}
	}

	// response作成
	collections := make([]*Collection, 0, len(collectionItems))
	for _, collectionItem := range collectionItems {
		_, hasItem := userCollectionItemMapByCollectionItemID[collectionItem.ID]
		collections = append(collections, &Collection{
			CollectionID: collectionItem.ID,
			Name:         collectionItem.Name,
			Rarity:       collectionItem.Rarity,
			HasItem:      hasItem,
		})
	}

	return collections, nil
}

// Collection コレクションアイテム情報詳細
type Collection struct {
	CollectionID string `json:"collectionID"`
	Name         string `json:"name"`
	Rarity       int32  `json:"rarity"`
	HasItem      bool   `json:"hasItem"`
}
