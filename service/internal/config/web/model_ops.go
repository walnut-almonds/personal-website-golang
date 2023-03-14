package admin

type MySQLOps struct {
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
}

type FileServerOps struct {
	TaskB266      FileServerDetail `mapstructure:"task_b_266"`
	TaskB277      FileServerDetail `mapstructure:"task_b_277"`
	TaskB277Fixed FileServerDetail `mapstructure:"task_b_277_fixed"`
}

type DWHOps struct {
	Address string `mapstructure:"address"`
	User    string `mapstructure:"user"`
	Passwd  string `mapstructure:"passwd"`
}

type FileServerDetail struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type MongoOps struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Dbname   string `mapstructure:"dbname"`
}
