package po

type AuthNode struct {
	Id      string `gorm:"column:id"`      // Id
	Name    string `gorm:"column:name"`    // 節點名稱
	Parent  string `gorm:"column:parent"`  // 父節點
	Type    int    `gorm:"column:type"`    // 0=節點 1=功能
	Sort    int    `gorm:"column:sort"`    // UI排序
	Enabled int    `gorm:"column:enabled"` // 狀態 0: 禁用 1: 啟用
}

func (AuthNode) TableName() string {
	return "auth_node"
}

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
