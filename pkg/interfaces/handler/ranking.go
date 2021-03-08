package handler

import (
	"errors"
	"log"
	"strconv"

	"golang.org/x/net/context"

	mr "20dojo-online/pkg/domain/model/ranking"
	"20dojo-online/pkg/interfaces/rpc/ranking"
	"20dojo-online/pkg/usecase/interactor"
)

// RankingService ランキングサービス
type RankingService struct {
	RankingUseCase interactor.RankingUseCase
}

// NewRankingService Rankingデータに関するServiceを生成
func NewRankingService(RankingUseCase interactor.RankingUseCase) RankingService {
	return RankingService{
		RankingUseCase: RankingUseCase,
	}
}

// RankingList generates ranking response
func (r *RankingService) RankingList(ctx context.Context, req *ranking.RankingListRequest) (*ranking.RankingListResponse, error) {
	start, err := strconv.Atoi(req.Start)
	if err != nil {
		log.Printf("strconv.Atoi is failed : %v", err)
		return nil, err
	}

	if start <= 0 {
		log.Println("start must be positive")
		return nil, errors.New("start-parameters is invalid")
	}

	rankList, err := r.RankingUseCase.GetRankingList(int64(start))
	if err != nil {
		log.Println("start must be positive")
		return nil, err
	}

	return &ranking.RankingListResponse{Ranks: toRanks(rankList)}, nil
}

func toRanks(rankList []*mr.RankInfo) []*ranking.Rank {
	res := make([]*ranking.Rank, len(rankList))
	for i, rank := range rankList {
		res[i] = &ranking.Rank{
			UserId:   rank.UserID,
			UserName: rank.UserName,
			Rank:     rank.Rank,
			Score:    rank.Score,
		}
	}
	return res
}
