package handler

import (
	"errors"
	"log"

	"golang.org/x/net/context"

	"20dojo-online/pkg/interfaces/rpc/game"
	"20dojo-online/pkg/usecase/interactor"
)

// GameService ユーザーサービス
type GameService struct {
	GameUseCase interactor.GameUseCase
}

// NewGameService Gameデータに関するServiceを生成
func NewGameService(GameUseCase interactor.GameUseCase) GameService {
	return GameService{
		GameUseCase: GameUseCase,
	}
}

// GameFinish generates Game-coin response
func (g *GameService) GameFinish(ctx context.Context, req *game.GameFinishRequest) (*game.GameFinishResponse, error) {
	if req.Score < 0 {
		log.Println("score is not positive")
		return nil, errors.New("score-parameter is invalid")
	}

	s, _ := ctx.Value("key").(string)
	coin, err := g.GameUseCase.GameFinish(s, req.Score)
	if err != nil {
		log.Println("score doesn't save")
		return nil, err
	}
	return &game.GameFinishResponse{Coin: coin}, nil
}
