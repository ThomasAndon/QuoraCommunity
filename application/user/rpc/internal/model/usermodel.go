package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindByMobNoCache(ctx context.Context, mob string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (c customUserModel) FindByMobNoCache(ctx context.Context, mob string) (*User, error) {
	var user User
	err := c.QueryRowNoCacheCtx(ctx, &user, fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, c.table), mob)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}
