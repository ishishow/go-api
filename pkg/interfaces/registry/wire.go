// +build wireinject

package registry

import (
	"github.com/google/wire"

	dbr "20dojo-online/pkg/infrastructure/mysql/repository"
	cacher "20dojo-online/pkg/infrastructure/redis/repository"

	"20dojo-online/pkg/infrastructure/mysql"
	"20dojo-online/pkg/infrastructure/redis"
	"20dojo-online/pkg/interfaces/handler"
	"20dojo-online/pkg/usecase/interactor"
)

func InitializeUserService() handler.UserService {

	wire.Build(
		handler.NewUserService,
		interactor.NewUserUseCase,
		mysql.NewSQLHandler,
		dbr.NewUserRepositoryImpl,
	)
	return handler.UserService{}
}

func InitializeGameService() handler.GameService {

	wire.Build(
		handler.NewGameService,
		interactor.NewGameUseCase,
		mysql.NewSQLHandler,
		dbr.NewUserRepositoryImpl,
	)
	return handler.GameService{}
}

func InitializeGachaService() handler.GachaService {

	wire.Build(
		handler.NewGachaService,
		interactor.NewGachaUseCase,
		mysql.NewSQLHandler,
		dbr.NewUserRepositoryImpl,
		dbr.NewCollectionItemRepositoryImpl,
		dbr.NewUserCollectionItemRepositoryImpl,
		dbr.NewGachaProbabilityRepositoryImpl,
	)
	return handler.GachaService{}
}

func InitializeRankingService() handler.RankingService {

	wire.Build(
		handler.NewRankingService,
		interactor.NewRankingUseCase,
		mysql.NewSQLHandler,
		dbr.NewUserRepositoryImpl,
		cacher.NewRankingRepositoryImpl,
		redis.NewCacheHandler,
	)
	return handler.RankingService{}
}

func InitializeCollectionService() handler.CollectionService {

	wire.Build(
		handler.NewCollectionService,
		interactor.NewCollectionItemUseCase,
		mysql.NewSQLHandler,
		dbr.NewCollectionItemRepositoryImpl,
		dbr.NewUserCollectionItemRepositoryImpl,
	)
	return handler.CollectionService{}
}
