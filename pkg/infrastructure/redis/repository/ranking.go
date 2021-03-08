package repository

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"20dojo-online/pkg/constant"
	"20dojo-online/pkg/domain/repository/cache"

	rm "20dojo-online/pkg/domain/model/ranking"
	um "20dojo-online/pkg/domain/model/user"
	re "20dojo-online/pkg/infrastructure/redis"

	"github.com/go-redis/redis"
)

type rankingRepositoryImpl struct {
	re.CacheHandler
}

// NewRankingRepositoryImpl Userに関するDB更新処理を生成
func NewRankingRepositoryImpl(cacheHandler re.CacheHandler) cache.RankingRepository {
	return &rankingRepositoryImpl{
		cacheHandler,
	}
}

// GetRankingList はソート済みマップから指定範囲のランキングを取得する
func (rri rankingRepositoryImpl) GetRankingList(start int64) ([]*rm.RankInfo, error) {
	redisZList, err := rri.Client.ZRangeByScoreWithScores(constant.RankingKey, redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: start,
		Count:  constant.RankingUserLimit,
	}).Result()
	if err != nil {
		return nil, err
	}
	return convertToRankInfo(redisZList, start)
}

// SaveRankingListはRankingUseCaseインスタンス生成時にソート済みランキングマップを作成する
func (rri rankingRepositoryImpl) SaveRankingList(users []*um.User) error {
	for _, user := range users {
		userJSON, _ := json.Marshal(user)
		if err := rri.Client.ZAdd(constant.RankingKey, redis.Z{
			Score:  float64(user.HighScore),
			Member: userJSON,
		}).Err(); err != nil {
			log.Printf("ZAdd userdata is failed: %v", err)
			return err
		}
	}
	return nil
}

func convertToRankInfo(Zlist []redis.Z, start int64) ([]*rm.RankInfo, error) {
	rankInfoList := make([]*rm.RankInfo, 0, constant.RankingUserLimit)
	for i, Z := range Zlist {
		user := um.User{}
		userJSON, ok := Z.Member.(string)
		if !ok {
			log.Println("casting is failed ... not string")
			return nil, errors.New("casting is failed")
		}

		dec := json.NewDecoder(strings.NewReader(userJSON))
		if err := dec.Decode(&user); err != nil {
			log.Printf("Decode error :%v", err)
			return nil, err
		}

		rankInfoList = append(rankInfoList, &rm.RankInfo{
			UserID:   user.ID,
			UserName: user.Name,
			Rank:     int32(i + int(start)),
			Score:    user.HighScore,
		})
	}
	return rankInfoList, nil
}
