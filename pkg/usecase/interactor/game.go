package interactor

import (
	ur "20dojo-online/pkg/domain/repository/db"
)

type gameUseCase struct {
	userRepository ur.UserRepository
}

// GameUseCase is
type GameUseCase interface {
	GameFinish(authToken string, score int32) (coin int32, err error)
}

// NewGameUseCase Userデータに関するユースケースを生成
func NewGameUseCase(userRepo ur.UserRepository) GameUseCase {
	return &gameUseCase{
		userRepository: userRepo,
	}
}

func (uu gameUseCase) GameFinish(authToken string, score int32) (int32, error) {
	user, err := uu.userRepository.SelectByAuthToken(authToken)
	if err != nil {
		return 0, err
	}
	// 得点が最高値かどうか判別
	if user.HighScore < score {
		user.HighScore = score
	}
	// 実行userにコインを付与
	user.Coin += score
	_ = uu.userRepository.Update(user)
	return user.Coin, nil
}
