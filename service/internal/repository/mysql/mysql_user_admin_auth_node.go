package mysql

import (
	"context"

	"gorm.io/gorm"

	"personal-website-golang/service/internal/model/po"
)

type IMysqlUserAdminAuthNode interface {
	List(ctx context.Context, db *gorm.DB, cond *po.AuthNode) ([]*po.AuthNode, error)
}

type mysqlUserAdminAuthNode struct {
	in repositoryIn
}

func newMysqlUserAdminAuthNode(in repositoryIn) IMysqlUserAdminAuthNode {
	return &mysqlUserAdminAuthNode{
		in: in,
	}
}

func (repo *mysqlUserAdminAuthNode) List(ctx context.Context, db *gorm.DB, cond *po.AuthNode) ([]*po.AuthNode, error) {
	var nodes []*po.AuthNode

	if err := db.
		Scopes(repo.listWhereScope(cond)).
		Find(&nodes).
		Error; err != nil {
		return nil, err
	}

	return nodes, nil
}

func (repo *mysqlUserAdminAuthNode) listWhereScope(cond *po.AuthNode) func(db *gorm.DB) *gorm.DB {
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
