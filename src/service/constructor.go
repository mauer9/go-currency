package service

import (
	"github.com/go-kit/log"
	"test/src/repository"
	"test/src/rpc"
)

type service struct {
	rpc       rpc.RPC
	logger    log.Logger
	mainStore repository.MainStore
}

func NewService(rpc rpc.RPC, logger log.Logger, store repository.MainStore) MainService {
	return &service{
		rpc:       rpc,
		logger:    logger,
		mainStore: store,
	}
}
