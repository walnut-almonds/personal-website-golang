package mysql

import (
	"context"

	"gorm.io/gorm"

	"personal-website-golang/service/internal/model/po"
)

type IMysqlUserAdminAuthPermission interface {
	List(ctx context.Context, db *gorm.DB, condFunc ...func(*gorm.DB) *gorm.DB) ([]*po.AuthPermission, error)
	Insert(ctx context.Context, db *gorm.DB, role *po.AuthPermission) error
	Delete(ctx context.Context, db *gorm.DB, roleId string) error
}

type mysqlUserAdminAuthPermission struct {
	in repositoryIn
}

func newMysqlUserAdminAuthPermission(in repositoryIn) IMysqlUserAdminAuthPermission {
	return &mysqlUserAdminAuthPermission{
		in: in,
	}
}

func (repo *mysqlUserAdminAuthPermission) List(ctx context.Context, db *gorm.DB, cond *po.AuthNode) ([]*po.AuthNode, error) {
	var nodes []*po.AuthNode

	if err := db.
		Scopes(repo.listWhereScope(cond)).
		Find(&nodes).
		Error; err != nil {
		return nil, err
	}

	return nodes, nil
}

func (repo *mysqlUserAdminAuthPermission) listWhereScope(cond *po.AuthNode) func(db *gorm.DB) *gorm.DB {
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
