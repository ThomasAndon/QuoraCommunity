package svc

import (
	"QuoraCommunity/application/user/rpc/internal/config"
	"QuoraCommunity/application/user/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	con := sqlx.NewMysql(c.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(con, c.CacheRedis),
	}
}
