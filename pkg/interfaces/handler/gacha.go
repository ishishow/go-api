package handler

import (
	"errors"
	"log"

	"golang.org/x/net/context"

	"20dojo-online/pkg/interfaces/rpc/gacha"
	"20dojo-online/pkg/usecase/interactor"
)

// GachaService ユーザーサービス
type GachaService struct {
	GachaUseCase interactor.GachaUseCase
}

// NewGachaService Gachaデータに関するServiceを生成
func NewGachaService(GachaUseCase interactor.GachaUseCase) GachaService {
	return GachaService{
		GachaUseCase: GachaUseCase,
	}
}

// GachaDraw generates gacha-results response
func (g *GachaService) GachaDraw(ctx context.Context, req *gacha.GachaDrawRequest) (*gacha.GachaDrawResponse, error) {
	if req.Times <= 0 {
		log.Println("times is not positive")
		return nil, errors.New("times-parameters is invalid")
	}

	authToken, _ := ctx.Value("key").(string) //TODO:
	gachaResults, err := g.GachaUseCase.GachaDraw(ctx, authToken, req.Times)
	if err != nil {
		log.Printf("GachaDraw is failed: %v", err)
		return nil, err
	}

	return &gacha.GachaDrawResponse{Results: toResults(gachaResults)}, nil
}

func toResults(results []*interactor.GachaResult) []*gacha.Result {
	res := make([]*gacha.Result, len(results))
	for i, result := range results {
		res[i] = &gacha.Result{
			CollectionId: result.CollectionID,
			Name:         result.Name,
			Rarity:       result.Rarity,
			IsNew:        result.IsNew,
		}
	}
	return res
}
