package po

type AccessFeature struct {
	Id     string `gorm:"column:id"`     // 序號
	Name   string `gorm:"column:name"`   // 名稱
	Path   string `gorm:"column:path"`   // 路徑
	Method string `gorm:"column:method"` // 方法
}

func (AccessFeature) TableName() string {
	return "access_feature"
}

type AuthRule struct {
	Id        string `gorm:"column:id"`         // 序號
	NodeId    string `gorm:"column:node_id"`    // AuthNode.Id
	FeatureId string `gorm:"column:feature_id"` // AccessFeature.Id
}

func (AuthRule) TableName() string {
	return "auth_rule"
}
