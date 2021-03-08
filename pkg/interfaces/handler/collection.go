package handler

import (
	"log"

	"golang.org/x/net/context"

	"20dojo-online/pkg/interfaces/rpc/collection"
	"20dojo-online/pkg/usecase/interactor"
)

// CollectionService ユーザーサービス
type CollectionService struct {
	CollectionUseCase interactor.CollectionItemUseCase
}

// NewCollectionService Collectionデータに関するServiceを生成
func NewCollectionService(CollectionUseCase interactor.CollectionItemUseCase) CollectionService {
	return CollectionService{
		CollectionUseCase: CollectionUseCase,
	}
}

// CollectionList generates user_collectionList response
func (c *CollectionService) CollectionList(ctx context.Context, req *collection.CollectionListRequest) (*collection.CollectionListResponse, error) {
	authToken, _ := ctx.Value("key").(string)
	// TODO: errorハンドリング あとkey
	collections, err := c.CollectionUseCase.SelectALLUserCollectionItems(authToken)
	if err != nil {
		log.Printf("cannot get userCollectionItems: %v", err)
		return nil, err
	}
	return &collection.CollectionListResponse{Collections: toCollections(collections)}, nil
}

func toCollections(collections []*interactor.Collection) []*collection.Collection {
	res := make([]*collection.Collection, len(collections))
	for i, c := range collections {
		res[i] = &collection.Collection{
			CollectionId: c.CollectionID,
			Name:         c.Name,
			Rarity:       c.Rarity,
			HasItem:      c.HasItem,
		}
	}
	return res
}
