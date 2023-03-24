package rbac

import (
	"context"

	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"personal-website-golang/service/internal/model/po"
	"personal-website-golang/service/internal/thirdparty/logger"
)

type IAuthClient interface {
}

type AuthClient struct {
	*casbin.Enforcer
}

func newAuth(sysLogger logger.ILogger, db *gorm.DB, prefix string) IAuthClient {
	cas, err := casbin.NewEnforcer("config/rbac_model.conf")
	if err != nil {
		panic(err)
	}

	adapter, err := gormAdapter.NewAdapterByDBUseTableName(db, prefix, "casbin_rule")
	if err != nil {
		panic(err)
	}
	cas.SetAdapter(adapter)
	cas.EnableAutoSave(true)

	adminInstance := &AuthClient{
		cas,
	}
	if err := adminInstance.initData(db); err != nil {
		panic(err)
	}

	sysLogger.Info(context.Background(), "RBAC init success")

	return adminInstance
}

func (ac *AuthClient) initData(db *gorm.DB) error {
	if err := ac.Enforcer.LoadPolicy(); err != nil {
		return err
	}

	var features []*po.AccessFeature
	if err := db.Find(&features).Error; err != nil {
		return err
	}
	for _, feature := range features {
		if _, err := ac.AddPolicy(feature.Id, feature.Path, feature.Method); err != nil {
			return err
		}
	}

	var rules []*po.AuthRule
	if err := db.Find(&rules).Error; err != nil {
		return err
	}
	for _, rule := range rules {
		if _, err := ac.AddGroupingPolicy(rule.NodeId, rule.FeatureId); err != nil {
			return err
		}
	}

	return nil
}
