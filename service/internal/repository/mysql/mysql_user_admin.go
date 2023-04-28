package mysql

import (
	"context"

	"gorm.io/gorm"

	"personal-website-golang/service/internal/model/po"
)

type IMysqlUserAdmin interface {
	List(ctx context.Context, db *gorm.DB, cond *po.UserAdmin) ([]*po.UserAdmin, error)
}

type mysqlUserAdmin struct {
	in repositoryIn
}

func newMysqlUserAdmin(in repositoryIn) IMysqlUserAdminAuthNode {
	return &mysqlUserAdmin{
		in: in,
	}
}

func (repo *mysqlUserAdmin) List(ctx context.Context, db *gorm.DB, cond *po.UserAdmin) ([]*po.UserAdmin, error) {
	var nodes []*po.UserAdmin

	if err := db.
		Scopes(repo.listWhereScope(cond)).
		Find(&nodes).
		Error; err != nil {
		return nil, err
	}

	return nodes, nil
}

func (repo *mysqlUserAdmin) Insert(ctx context.Context, db *gorm.DB, user *po.UserAdmin) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
