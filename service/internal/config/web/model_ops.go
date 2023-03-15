package admin

type MySQLOps struct {
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
}

type MongoOps struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Dbname   string `mapstructure:"dbname"`
}

type FileServerOps struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type SlackOps struct {
	Token     string `mapstructure:"token"`
	ChannelID string `mapstructure:"channel_id"`
	BotName   string `mapstructure:"bot_name"`
}

type TelegramOps struct {
	Token string `mapstructure:"token"`
}
