package mysql

import (
	"context"

	"gorm.io/gorm"

	"personal-website-golang/service/internal/model/po"
)

type IMysqlUserAdminAuthRole interface {
	List(ctx context.Context, db *gorm.DB, cond *po.AuthNode) ([]*po.AuthNode, error)
}

type mysqlUserAdminAuthRole struct {
	in repositoryIn
}

func newMysqlUserAdminAuthRole(in repositoryIn) IMysqlUserAdminAuthRole {
	return &mysqlUserAdminAuthRole{
		in: in,
	}
}

func (repo *mysqlUserAdminAuthRole) List(ctx context.Context, db *gorm.DB, cond *po.AuthNode) ([]*po.AuthNode, error) {
	var nodes []*po.AuthNode

	if err := db.
		Scopes(repo.listWhereScope(cond)).
		Find(&nodes).
		Error; err != nil {
		return nil, err
	}

	return nodes, nil
}

func (repo *mysqlUserAdminAuthRole) listWhereScope(cond *po.AuthNode) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if cond.Id != "" {
			db = db.Where("id = ?", cond.Id)
		}

		if cond.Name != "" {
			db = db.Where("name = ?", cond.Name)
		}

		if cond.Parent != "" {
			db = db.Where("parent = ?", cond.Parent)
		}

		return db.Model(&po.AuthNode{})
	}
}
