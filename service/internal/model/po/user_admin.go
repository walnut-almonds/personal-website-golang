package po

import "time"

type UserAdmin struct {
	ID            int64      `gorm:"column:id"`              // 會員ID
	Level         int        `gorm:"column:level"`           // 會員級別
	Status        int        `gorm:"column:status"`          // 狀態
	Password      string     `gorm:"column:password"`        // 密碼Hash
	Salt          string     `gorm:"column:salt"`            // 密碼鹽
	LastLoginTime *time.Time `gorm:"column:last_login_time"` // 最後登入時間
	Remark        *string    `gorm:"column:remark"`          // 備注
	CreateTime    time.Time  `gorm:"column:create_time"`     // 創建時間
	UpdateTime    time.Time  `gorm:"column:update_time"`     // 修改時間
}

func (UserAdmin) TableName() string {
	return "user_admin"
}
